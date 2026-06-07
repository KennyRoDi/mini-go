import React, { useState, useRef, useEffect, useCallback } from 'react';
import Editor, { useMonaco } from '@monaco-editor/react';
import { Files, FileCode, X, Plus } from 'lucide-react';

interface TerminalLine {
  text: string;
  type: 'prompt' | 'error' | 'success' | 'default';
}

interface FileTab {
  id: string;
  name: string;
  content: string;
  path: string | null;
}

const App: React.FC = () => {
  const [tabs, setTabs] = useState<FileTab[]>([
    { id: '1', name: 'untitled.go', content: 'package main;\n\nfunc main() {\n    print("Hello Mini-GO");\n};\n', path: null }
  ]);
  const [activeTabId, setActiveTabId] = useState<string>('1');
  const [terminalLines, setTerminalLines] = useState<TerminalLine[]>([{ text: 'Waiting for input...', type: 'default' }]);
  const [isCompiling, setIsCompiling] = useState(false);
  const monaco = useMonaco();
  const editorRef = useRef<any>(null);

  const activeTab = tabs.find(t => t.id === activeTabId);

  const appendTerminalLine = useCallback((text: string, type: TerminalLine['type']) => {
    if (!text) return;
    setTerminalLines(prev => [...prev, { text, type }]);
  }, []);

  const clearTerminal = useCallback(() => {
    setTerminalLines([{ text: 'minigo-ide $', type: 'prompt' }]);
  }, []);

  const handleNewFile = useCallback(() => {
    const newId = Date.now().toString();
    setTabs(prev => [...prev, { id: newId, name: 'untitled.go', content: '', path: null }]);
    setActiveTabId(newId);
  }, []);

  const handleOpenFile = useCallback(async () => {
    const file = await window.ipcRenderer.invoke('open-file');
    if (file) {
      const fileName = file.path.split(/[\\/]/).pop();
      const newId = Date.now().toString();
      setTabs(prev => [...prev, { id: newId, name: fileName, content: file.content, path: file.path }]);
      setActiveTabId(newId);
      if (editorRef.current && monaco) {
        monaco.editor.setModelMarkers(editorRef.current.getModel(), 'minigo', []);
      }
    }
  }, [monaco]);

  const handleSaveFile = useCallback(async () => {
    if (!activeTab) return;
    const savedPath = await window.ipcRenderer.invoke('save-file', activeTab.content, activeTab.path);
    if (savedPath) {
      const fileName = savedPath.split(/[\\/]/).pop();
      setTabs(prev => prev.map(t => t.id === activeTab.id ? { ...t, name: fileName, path: savedPath } : t));
      appendTerminalLine('File saved successfully', 'success');
    }
  }, [activeTab, appendTerminalLine]);

  const highlightErrors = useCallback((output: string) => {
    if (!editorRef.current || !monaco) return;
    const markers: any[] = [];
    const errorRegex = /(?:ERROR|SYNTAX ERROR).*?\[linea:(\d+)\s*-\s*columna:(\d+)\]/g;
    let match;

    while ((match = errorRegex.exec(output)) !== null) {
      const line = parseInt(match[1]);
      const col = parseInt(match[2]);
      markers.push({
        severity: monaco.MarkerSeverity.Error,
        message: match[0],
        startLineNumber: line,
        startColumn: col,
        endLineNumber: line,
        endColumn: col + 5
      });
    }
    monaco.editor.setModelMarkers(editorRef.current.getModel(), 'minigo', markers);
  }, [monaco]);

  const handleRun = useCallback(async () => {
    if (!activeTab) return;
    setIsCompiling(true);
    clearTerminal();
    appendTerminalLine('Compiling and running...', 'default');
    
    if (editorRef.current && monaco) {
        monaco.editor.setModelMarkers(editorRef.current.getModel(), 'minigo', []);
    }

    const result = await window.ipcRenderer.invoke('compile', activeTab.content);
    setIsCompiling(false);
    
    if (result.success) {
      appendTerminalLine(result.output, 'default');
      appendTerminalLine('Execution finished successfully.', 'success');
    } else {
      appendTerminalLine(`Error in ${result.phase}:`, 'error');
      appendTerminalLine(result.output, 'error');
      highlightErrors(result.output);
    }
  }, [activeTab, monaco, clearTerminal, appendTerminalLine, highlightErrors]);

  useEffect(() => {
    const onNew = () => handleNewFile();
    const onOpen = () => handleOpenFile();
    const onSave = () => handleSaveFile();
    const onRun = () => handleRun();
    const onUndo = () => editorRef.current?.trigger('keyboard', 'undo', null);
    const onRedo = () => editorRef.current?.trigger('keyboard', 'redo', null);
    const onSelectAll = () => editorRef.current?.trigger('keyboard', 'editor.action.selectAll', null);

    window.ipcRenderer.removeAllListeners('menu-new-file');
    window.ipcRenderer.removeAllListeners('menu-open-file');
    window.ipcRenderer.removeAllListeners('menu-save-file');
    window.ipcRenderer.removeAllListeners('menu-run');
    window.ipcRenderer.removeAllListeners('menu-undo');
    window.ipcRenderer.removeAllListeners('menu-redo');
    window.ipcRenderer.removeAllListeners('menu-select-all');

    window.ipcRenderer.on('menu-new-file', onNew);
    window.ipcRenderer.on('menu-open-file', onOpen);
    window.ipcRenderer.on('menu-save-file', onSave);
    window.ipcRenderer.on('menu-run', onRun);
    window.ipcRenderer.on('menu-undo', onUndo);
    window.ipcRenderer.on('menu-redo', onRedo);
    window.ipcRenderer.on('menu-select-all', onSelectAll);

    return () => {
      window.ipcRenderer.off('menu-new-file', onNew);
      window.ipcRenderer.off('menu-open-file', onOpen);
      window.ipcRenderer.off('menu-save-file', onSave);
      window.ipcRenderer.off('menu-run', onRun);
      window.ipcRenderer.off('menu-undo', onUndo);
      window.ipcRenderer.off('menu-redo', onRedo);
      window.ipcRenderer.off('menu-select-all', onSelectAll);
    };
  }, [handleNewFile, handleOpenFile, handleSaveFile, handleRun]);

  const updateActiveContent = (val: string | undefined) => {
    setTabs(prev => prev.map(t => t.id === activeTabId ? { ...t, content: val || '' } : t));
  };

  const closeTab = (e: React.MouseEvent, id: string) => {
    e.stopPropagation();
    setTabs(prev => {
      const filtered = prev.filter(t => t.id !== id);
      if (activeTabId === id && filtered.length > 0) {
        setActiveTabId(filtered[0].id);
      } else if (filtered.length === 0) {
        setActiveTabId('');
      }
      return filtered;
    });
  };

  return (
    <div className="flex h-screen flex-col bg-[#181818] overflow-hidden text-sm">
      <div className="flex flex-1 overflow-hidden">
        {/* Activity Bar */}
        <div className="w-12 border-r border-[#333] flex flex-col items-center py-4 gap-6 bg-[#181818] flex-shrink-0">
          <Files className="text-white cursor-pointer w-6 h-6 border-l-2 border-[#5fb3b3] pl-1 -ml-1" />
        </div>

        {/* Sidebar */}
        <div className="w-60 bg-[#252526] border-r border-[#333] hidden md:flex flex-col flex-shrink-0">
          <div className="p-3 text-[11px] font-bold text-[#858585] uppercase tracking-wider flex justify-between items-center">
            <span>Explorer</span>
            <Plus className="w-4 h-4 cursor-pointer hover:text-white" onClick={handleNewFile} />
          </div>
          <div className="flex-1 overflow-y-auto">
            {tabs.map(tab => (
              <div 
                key={tab.id}
                onClick={() => setActiveTabId(tab.id)}
                className={`flex items-center px-4 py-2 gap-2 text-sm cursor-pointer hover:bg-[#2a2d2e] ${activeTabId === tab.id ? 'bg-[#37373d] text-white' : 'text-[#cccccc]'}`}
              >
                <div className={`w-2 h-2 rounded-full ${activeTabId === tab.id ? 'bg-[#5fb3b3]' : 'bg-transparent'}`}></div>
                <FileCode size={14} className="text-[#5fb3b3]" />
                <span className="truncate flex-1">{tab.name}</span>
                <X className="w-4 h-4 opacity-50 hover:opacity-100" onClick={(e) => closeTab(e, tab.id)} />
              </div>
            ))}
          </div>
        </div>

        {/* Main Stage */}
        <div className="flex-1 flex flex-col min-w-0 bg-[#1e1e1e]">
          {/* Tabs Bar */}
          <div className="flex bg-[#252526] border-b border-[#333] overflow-x-auto items-center flex-shrink-0">
            {tabs.map(tab => (
              <div 
                key={tab.id}
                onClick={() => setActiveTabId(tab.id)}
                className={`flex items-center gap-2 px-4 py-2 cursor-pointer border-r border-[#333] min-w-[120px] max-w-[200px] ${activeTabId === tab.id ? 'bg-[#1e1e1e] text-white border-t-2 border-t-[#5fb3b3]' : 'bg-[#2d2d2d] text-[#858585] border-t-2 border-t-transparent'}`}
              >
                <FileCode size={12} className="text-[#5fb3b3]" />
                <span className="truncate flex-1 text-xs">{tab.name}</span>
                <X className="w-3 h-3 hover:bg-[#333] rounded" onClick={(e) => closeTab(e, tab.id)} />
              </div>
            ))}
            <Plus size={14} className="mx-4 text-[#858585] cursor-pointer hover:text-white flex-shrink-0" onClick={handleNewFile} />
            <div className="ml-auto px-3 flex-shrink-0">
              <button
                onClick={handleRun}
                disabled={isCompiling || !activeTab}
                className={`flex items-center gap-1 px-3 py-1 rounded text-xs font-bold transition-colors ${isCompiling || !activeTab ? 'bg-[#333] text-[#555] cursor-not-allowed' : 'bg-[#3a7c3a] text-white hover:bg-[#4a9c4a] cursor-pointer'}`}
              >
                {isCompiling ? <span className="animate-pulse">Compiling...</span> : <span>▶ Run</span>}
              </button>
            </div>
          </div>

          {/* Editor */}
          <div className="flex-1 relative overflow-hidden border-b border-[#333]">
            {activeTab ? (
              <Editor
                height="100%"
                language="go"
                theme="vs-dark"
                value={activeTab.content}
                onChange={updateActiveContent}
                onMount={(editor) => { editorRef.current = editor }}
                options={{
                  fontSize: 15,
                  minimap: { enabled: false },
                  automaticLayout: true,
                  padding: { top: 20 }
                }}
              />
            ) : (
              <div className="flex items-center justify-center h-full text-[#858585]">
                No file is open.
              </div>
            )}
          </div>

          {/* Bottom Panel */}
          <div className="h-64 bg-[#1e1e1e] flex flex-col flex-shrink-0">
            <div className="flex border-b border-[#333] bg-[#252526]">
               <div className="px-4 py-2 text-[11px] font-bold text-white border-b-2 border-[#5fb3b3] bg-[#1e1e1e]">TERMINAL</div>
            </div>
            <div className="flex-1 p-4 font-mono text-sm overflow-y-auto whitespace-pre-wrap text-[#ccc]">
              {terminalLines.map((line, i) => (
                <div key={i} className={
                  line.type === 'error' ? 'text-[#f48771]' : 
                  line.type === 'success' ? 'text-[#89d185]' : 
                  line.type === 'prompt' ? 'text-[#5fb3b3]' : ''
                }>
                  {line.type === 'prompt' && <span className="mr-2">minigo-ide $</span>}
                  {line.text}
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
