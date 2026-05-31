import { app, BrowserWindow, ipcMain } from 'electron';
import * as path from 'path';
import { spawn } from 'child_process';
import * as fs from 'fs';

function createWindow() {
  const win = new BrowserWindow({
    width: 1000,
    height: 800,
    webPreferences: {
      nodeIntegration: true,
      contextIsolation: false, // For simplicity in this project
    },
  });

  win.loadFile(path.join(__dirname, '../src/index.html'));
  // win.webContents.openDevTools();
}

app.whenReady().then(createWindow);

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') app.quit();
});

// IPC Handler for compilation
ipcMain.handle('compile', async (event, code: string) => {
  const tempFile = path.join(__dirname, '../../backend/temp_ide.go');
  fs.writeFileSync(tempFile, code);

  const compilerPath = path.join(__dirname, '../../backend/minigo');
  
  return new Promise((resolve) => {
    const child = spawn(compilerPath, [tempFile]);
    let output = '';
    
    child.stdout.on('data', (data) => output += data.toString());
    child.stderr.on('data', (data) => output += data.toString());
    
    child.on('close', (code) => {
      resolve({
        success: code === 0,
        output: output
      });
    });
  });
});
