import { app, BrowserWindow, ipcMain, dialog, Menu, MenuItemConstructorOptions } from 'electron'
import path from 'node:path'
import { spawn } from 'node:child_process'
import fs from 'node:fs'

process.env.DIST = path.join(__dirname, '../dist')
process.env.VITE_PUBLIC = app.isPackaged ? process.env.DIST : path.join(__dirname, '../public')

let win: BrowserWindow | null

function createWindow() {
  win = new BrowserWindow({
    width: 1200,
    height: 900,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      nodeIntegration: false,
      contextIsolation: true,
    },
  })

  // --- NATIVE MENU ---
  const template: MenuItemConstructorOptions[] = [
    {
      label: 'File',
      submenu: [
        { label: 'New File', accelerator: 'CmdOrCtrl+N', click: () => win?.webContents.send('menu-new-file') },
        { label: 'Open File...', accelerator: 'CmdOrCtrl+O', click: () => win?.webContents.send('menu-open-file') },
        { label: 'Save', accelerator: 'CmdOrCtrl+S', click: () => win?.webContents.send('menu-save-file') },
        { type: 'separator' },
        { label: 'Run', accelerator: 'CmdOrCtrl+R', click: () => win?.webContents.send('menu-run') },
        { type: 'separator' },
        { role: 'quit' }
      ]
    },
    {
      label: 'Edit',
      submenu: [
        { label: 'Undo', accelerator: 'CmdOrCtrl+Z', click: () => win?.webContents.send('menu-undo') },
        { label: 'Redo', accelerator: 'CmdOrCtrl+Y', click: () => win?.webContents.send('menu-redo') },
        { type: 'separator' },
        { label: 'Cut', accelerator: 'CmdOrCtrl+X', role: 'cut' },
        { label: 'Copy', accelerator: 'CmdOrCtrl+C', role: 'copy' },
        { label: 'Paste', accelerator: 'CmdOrCtrl+V', role: 'paste' },
        { type: 'separator' },
        { label: 'Select All', accelerator: 'CmdOrCtrl+A', click: () => win?.webContents.send('menu-select-all') }
      ]
    },
    { role: 'viewMenu' },
    { role: 'windowMenu' }
  ]
  const menu = Menu.buildFromTemplate(template)
  Menu.setApplicationMenu(menu)

  win.webContents.on('did-finish-load', () => {
    win?.webContents.send('main-process-message', (new Date).toLocaleString())
  })

  if (process.env.VITE_DEV_SERVER_URL) {
    win.loadURL(process.env.VITE_DEV_SERVER_URL)
  } else {
    win.loadFile(path.join(process.env.DIST!, 'index.html'))
  }
}

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit()
    win = null
  }
})

app.whenReady().then(createWindow)

// --- IPC HANDLERS ---

ipcMain.handle('open-file', async () => {
  const { canceled, filePaths } = await dialog.showOpenDialog({
    properties: ['openFile'],
    filters: [{ name: 'MiniGo Files', extensions: ['go', 'mg'] }]
  });
  if (canceled) return null;
  return {
    path: filePaths[0],
    content: fs.readFileSync(filePaths[0], 'utf-8')
  };
});

ipcMain.handle('save-file', async (event, content: string, existingPath: string | null) => {
  let filePath = existingPath;
  if (!filePath) {
    const { canceled, filePath: dialogPath } = await dialog.showSaveDialog({
      filters: [{ name: 'MiniGo Files', extensions: ['go', 'mg'] }]
    });
    if (canceled || !dialogPath) return null;
    filePath = dialogPath;
  }
  fs.writeFileSync(filePath, content);
  return filePath;
});

ipcMain.handle('compile', async (event, code: string) => {
  const tempDir = path.join(app.getPath('userData'), 'temp');
  if (!fs.existsSync(tempDir)) fs.mkdirSync(tempDir, { recursive: true });
  
  const tempFile = path.join(tempDir, 'temp_ide.go');
  fs.writeFileSync(tempFile, code);

  const projectRoot = path.resolve(app.getAppPath(), '..');
  const compilerPath = path.join(projectRoot, 'backend', 'minigo');
  const workingDir = tempDir;
  
  return new Promise((resolve) => {
    const compiler = spawn(compilerPath, [tempFile], { cwd: workingDir });
    let compilerOutput = '';
    
    compiler.stdout.on('data', (data) => compilerOutput += data.toString());
    compiler.stderr.on('data', (data) => compilerOutput += data.toString());
    
    compiler.on('error', (err) => {
        resolve({ success: false, output: `Failed to start compiler: ${err.message}\nPath: ${compilerPath}`, phase: 'Spawn' });
    });

    compiler.on('close', (exitCode) => {
      if (exitCode !== 0) {
        resolve({ success: false, output: compilerOutput, phase: 'Compilation' });
        return;
      }
// 2. Run generated binary (if compilation succeeded)
const binaryPath = path.join(workingDir, 'output');
if (!fs.existsSync(binaryPath)) {
    resolve({ success: false, output: compilerOutput + "\nError: executable not found", phase: 'Linker' });
    return;
}

// Small delay to prevent ETXTBSY on Linux
setTimeout(() => {
  const runner = spawn(binaryPath, [], { cwd: workingDir });
  let runOutput = '';
  runner.stdout.on('data', (data) => runOutput += data.toString());
  runner.stderr.on('data', (data) => runOutput += data.toString());

  runner.on('error', (err) => {
      resolve({ success: false, output: `Failed to start executable: ${err.message}`, phase: 'Execution Spawn' });
  });

  runner.on('close', (runCode) => {
    resolve({
      success: runCode === 0,
      output: runOutput,
      compilerLog: compilerOutput,
      phase: 'Execution'
    });
  });
}, 150);
});
});
});
