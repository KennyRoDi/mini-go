import React, { useState, useRef, useEffect, useCallback } from 'react';
import Editor, { useMonaco } from '@monaco-editor/react';
import { Files, FileCode, X, Plus, AlertCircle, Play, CheckCircle2, ChevronRight } from 'lucide-react';

interface TerminalLine {
  text: string;
  type: 'command' | 'output' | 'error' | 'success' | 'info';
}

interface FileTab {
  id: string;
  name: string;
  content: string;
  path: string | null;
  isDirty: boolean;
}

interface Problem {
  line: number;
  col: number;
  message: string;
  detail: string;
  kind: 'type' | 'syntax';
}

const App: React.FC = () => {
  const [tabs, setTabs] = useState<FileTab[]>([
    {
      id: '1',
      name: 'untitled.go',
      content: 'package main;\n\nfunc main() {\n    println("Hello Mini-Go!");\n};\n',
      path: null,
      isDirty: false,
    },
  ]);
  const [activeTabId, setActiveTabId] = useState<string>('1');
  const [terminalLines, setTerminalLines] = useState<TerminalLine[]>([]);
  const [isCompiling, setIsCompiling] = useState(false);
  const [activePanel, setActivePanel] = useState<'terminal' | 'problems'>('terminal');
  const [problems, setProblems] = useState<Problem[]>([]);
  const [cursorPos, setCursorPos] = useState({ line: 1, col: 1 });
  const [showSidebar, setShowSidebar] = useState(true);
  const monaco = useMonaco();
  const editorRef = useRef<any>(null);
  const terminalEndRef = useRef<HTMLDivElement>(null);

  const activeTab = tabs.find(t => t.id === activeTabId);
  const errorCount = problems.length;

  useEffect(() => {
    terminalEndRef.current?.scrollIntoView({ behavior: 'auto' });
  }, [terminalLines]);

  const appendLine = useCallback((text: string, type: TerminalLine['type']) => {
    setTerminalLines(prev => [...prev, { text, type }]);
  }, []);

  const clearTerminal = useCallback(() => setTerminalLines([]), []);

  const parseProblems = useCallback((output: string): Problem[] => {
    const results: Problem[] = [];
    const typeRx = /ERROR DE TIPO: (.+?) \((.+?)\) \[linea:(\d+)\s*-\s*columna:(\d+)\]/g;
    const syntaxRx = /SYNTAX ERROR: (.+?) \[linea:(\d+)\s*-\s*columna:(\d+)\]/g;
    let m;
    while ((m = typeRx.exec(output)) !== null)
      results.push({ kind: 'type', message: m[1], detail: m[2], line: +m[3], col: +m[4] });
    while ((m = syntaxRx.exec(output)) !== null)
      results.push({ kind: 'syntax', message: m[1], detail: '', line: +m[2], col: +m[3] });
    return results;
  }, []);

  const applyMarkers = useCallback((parsedProblems: Problem[]) => {
    if (!editorRef.current || !monaco) return;
    const markers = parsedProblems.map(p => ({
      severity: monaco.MarkerSeverity.Error,
      message: p.detail ? `${p.message}: ${p.detail}` : p.message,
      startLineNumber: p.line,
      startColumn: p.col,
      endLineNumber: p.line,
      endColumn: p.col + 8,
    }));
    monaco.editor.setModelMarkers(editorRef.current.getModel(), 'minigo', markers);
  }, [monaco]);

  const handleNewFile = useCallback(() => {
    const id = Date.now().toString();
    setTabs(prev => [...prev, { id, name: 'untitled.go', content: '', path: null, isDirty: false }]);
    setActiveTabId(id);
  }, []);

  const handleOpenFile = useCallback(async () => {
    const file = await window.ipcRenderer.invoke('open-file');
    if (!file) return;
    const name = file.path.split(/[\\/]/).pop();
    const id = Date.now().toString();
    setTabs(prev => [...prev, { id, name, content: file.content, path: file.path, isDirty: false }]);
    setActiveTabId(id);
    if (editorRef.current && monaco)
      monaco.editor.setModelMarkers(editorRef.current.getModel(), 'minigo', []);
  }, [monaco]);

  const handleSaveFile = useCallback(async () => {
    if (!activeTab) return;
    const savedPath = await window.ipcRenderer.invoke('save-file', activeTab.content, activeTab.path);
    if (savedPath) {
      const name = savedPath.split(/[\\/]/).pop();
      setTabs(prev => prev.map(t => t.id === activeTab.id ? { ...t, name, path: savedPath, isDirty: false } : t));
    }
  }, [activeTab]);

  const handleRun = useCallback(async () => {
    if (!activeTab) return;
    setIsCompiling(true);
    setProblems([]);
    clearTerminal();
    if (editorRef.current && monaco)
      monaco.editor.setModelMarkers(editorRef.current.getModel(), 'minigo', []);

    appendLine(`$ minigo ${activeTab.name}`, 'command');

    const result = await window.ipcRenderer.invoke('compile', activeTab.content);
    setIsCompiling(false);

    if (result.success) {
      const lines = (result.output || '').split('\n');
      if (lines[lines.length - 1] === '') lines.pop();
      lines.forEach((l: string) => appendLine(l, 'output'));
      appendLine('✓ Exited with code 0', 'success');
      setActivePanel('terminal');
    } else {
      const parsed = parseProblems(result.output || '');
      setProblems(parsed);
      applyMarkers(parsed);
      appendLine(
        `✗ ${result.phase === 'compilation' ? 'Compilation' : 'Build'} failed — ${parsed.length} error${parsed.length !== 1 ? 's' : ''}`,
        'error'
      );
      setActivePanel('problems');
    }
  }, [activeTab, monaco, clearTerminal, appendLine, parseProblems, applyMarkers]);

  const jumpToLine = useCallback((line: number, col: number) => {
    if (!editorRef.current) return;
    editorRef.current.revealLineInCenter(line);
    editorRef.current.setPosition({ lineNumber: line, column: col });
    editorRef.current.focus();
  }, []);

  useEffect(() => {
    const handlers: [string, () => void][] = [
      ['menu-new-file', handleNewFile],
      ['menu-open-file', handleOpenFile],
      ['menu-save-file', handleSaveFile],
      ['menu-run', handleRun],
      ['menu-undo', () => editorRef.current?.trigger('keyboard', 'undo', null)],
      ['menu-redo', () => editorRef.current?.trigger('keyboard', 'redo', null)],
      ['menu-select-all', () => editorRef.current?.trigger('keyboard', 'editor.action.selectAll', null)],
    ];
    handlers.forEach(([ev]) => window.ipcRenderer.removeAllListeners(ev));
    handlers.forEach(([ev, fn]) => window.ipcRenderer.on(ev, fn));
    return () => handlers.forEach(([ev, fn]) => window.ipcRenderer.off(ev, fn));
  }, [handleNewFile, handleOpenFile, handleSaveFile, handleRun]);

  const updateActiveContent = (val: string | undefined) => {
    setTabs(prev => prev.map(t => t.id === activeTabId ? { ...t, content: val ?? '', isDirty: true } : t));
  };

  const closeTab = (e: React.MouseEvent, id: string) => {
    e.stopPropagation();
    setTabs(prev => {
      const next = prev.filter(t => t.id !== id);
      if (activeTabId === id) setActiveTabId(next[0]?.id ?? '');
      return next;
    });
  };

  return (
    <div className="flex h-screen flex-col overflow-hidden text-[13px]" style={{ background: '#1e1e1e', color: '#cccccc', fontFamily: 'system-ui, sans-serif' }}>
      <div className="flex flex-1 overflow-hidden">

        {/* Activity Bar */}
        <div style={{ width: 48, background: '#181818', borderRight: '1px solid #252526' }} className="flex flex-col items-center py-2 gap-1 flex-shrink-0">
          <button
            onClick={() => setShowSidebar(s => !s)}
            title="Explorer"
            style={{ color: showSidebar ? '#cccccc' : '#6c6c6c' }}
            className="p-2.5 rounded hover:text-white transition-colors"
          >
            <Files size={20} />
          </button>
        </div>

        {/* Sidebar */}
        {showSidebar && (
          <div style={{ width: 220, background: '#252526', borderRight: '1px solid #1e1e1e' }} className="flex flex-col flex-shrink-0">
            <div style={{ fontSize: 11, color: '#bbb', letterSpacing: '0.1em' }} className="px-4 py-2 font-bold uppercase">
              Explorer
            </div>
            <div className="flex-1 overflow-y-auto">
              {tabs.map(tab => (
                <div
                  key={tab.id}
                  onClick={() => setActiveTabId(tab.id)}
                  style={{
                    background: activeTabId === tab.id ? '#37373d' : 'transparent',
                    color: activeTabId === tab.id ? '#ffffff' : '#adadad',
                  }}
                  className="group flex items-center px-4 py-1.5 gap-2 cursor-pointer hover:bg-[#2a2a2a] hover:text-white"
                >
                  <FileCode size={14} style={{ color: '#75bfbf', opacity: activeTabId === tab.id ? 1 : 0.7, flexShrink: 0 }} />
                  <span className="truncate flex-1" style={{ fontSize: 13 }}>{tab.name}</span>
                  {tab.isDirty && (
                    <span style={{ width: 7, height: 7, borderRadius: '50%', background: '#e2c08d', flexShrink: 0 }} className="group-hover:hidden" />
                  )}
                  <X
                    size={14}
                    className={`flex-shrink-0 rounded hover:text-white ${tab.isDirty ? 'opacity-0 group-hover:opacity-60' : 'opacity-0 group-hover:opacity-60'}`}
                    onClick={(e) => closeTab(e, tab.id)}
                  />
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Main Area */}
        <div className="flex-1 flex flex-col min-w-0">

          {/* Tab Bar */}
          <div style={{ background: '#252526', borderBottom: '1px solid #1e1e1e' }} className="flex items-stretch flex-shrink-0">
            <div className="flex items-stretch flex-1 overflow-x-auto">
              {tabs.map(tab => {
                const active = activeTabId === tab.id;
                return (
                  <div
                    key={tab.id}
                    onClick={() => setActiveTabId(tab.id)}
                    style={{
                      background: active ? '#1e1e1e' : '#2d2d2d',
                      color: active ? '#fff' : '#888',
                      borderRight: '1px solid #1e1e1e',
                      minWidth: 110,
                      maxWidth: 180,
                      position: 'relative',
                      paddingTop: 7,
                      paddingBottom: 7,
                      paddingLeft: 14,
                      paddingRight: 10,
                    }}
                    className="flex items-center gap-2 cursor-pointer group flex-shrink-0 hover:text-[#ccc]"
                  >
                    {active && (
                      <div style={{ position: 'absolute', top: 0, left: 0, right: 0, height: 1, background: '#75bfbf' }} />
                    )}
                    <FileCode size={12} style={{ color: '#75bfbf', flexShrink: 0 }} />
                    <span className="truncate flex-1" style={{ fontSize: 12 }}>{tab.name}</span>
                    {tab.isDirty ? (
                      <>
                        <span style={{ width: 7, height: 7, borderRadius: '50%', background: '#e2c08d', flexShrink: 0 }} className="group-hover:hidden" />
                        <X size={12} style={{ flexShrink: 0 }} className="hidden group-hover:block opacity-60 hover:opacity-100 rounded hover:bg-[#444]" onClick={(e) => closeTab(e, tab.id)} />
                      </>
                    ) : (
                      <X size={12} style={{ flexShrink: 0 }} className="opacity-0 group-hover:opacity-60 hover:opacity-100 rounded hover:bg-[#444]" onClick={(e) => closeTab(e, tab.id)} />
                    )}
                  </div>
                );
              })}
              <button
                onClick={handleNewFile}
                title="New File"
                className="px-3 text-[#666] hover:text-[#ccc] transition-colors flex-shrink-0"
              >
                <Plus size={14} />
              </button>
            </div>

            {/* Run Button */}
            <div style={{ borderLeft: '1px solid #1e1e1e' }} className="flex items-center px-3 flex-shrink-0">
              <button
                onClick={handleRun}
                disabled={isCompiling || !activeTab}
                title="Run  Ctrl+R"
                style={{
                  background: isCompiling || !activeTab ? '#2a2a2a' : '#1c6e1c',
                  color: isCompiling || !activeTab ? '#555' : '#fff',
                  cursor: isCompiling ? 'wait' : !activeTab ? 'not-allowed' : 'pointer',
                  borderRadius: 4,
                  padding: '4px 12px',
                  fontSize: 12,
                  fontWeight: 600,
                  display: 'flex',
                  alignItems: 'center',
                  gap: 6,
                  border: 'none',
                  transition: 'background 0.15s',
                }}
                onMouseEnter={e => { if (!isCompiling && activeTab) (e.currentTarget as HTMLButtonElement).style.background = '#237a23'; }}
                onMouseLeave={e => { if (!isCompiling && activeTab) (e.currentTarget as HTMLButtonElement).style.background = '#1c6e1c'; }}
              >
                {isCompiling ? (
                  <>
                    <span style={{ width: 8, height: 8, borderRadius: '50%', background: '#888', animation: 'pulse 1s infinite' }} />
                    Running...
                  </>
                ) : (
                  <>
                    <Play size={11} fill="currentColor" />
                    Run
                  </>
                )}
              </button>
            </div>
          </div>

          {/* Editor */}
          <div className="flex-1 relative overflow-hidden">
            {activeTab ? (
              <Editor
                height="100%"
                language="go"
                theme="vs-dark"
                value={activeTab.content}
                onChange={updateActiveContent}
                onMount={(editor) => {
                  editorRef.current = editor;
                  editor.onDidChangeCursorPosition((e: any) => {
                    setCursorPos({ line: e.position.lineNumber, col: e.position.column });
                  });
                }}
                options={{
                  fontSize: 14,
                  fontFamily: '"JetBrains Mono", "Fira Code", "Cascadia Code", Consolas, monospace',
                  fontLigatures: true,
                  minimap: { enabled: false },
                  automaticLayout: true,
                  padding: { top: 12 },
                  scrollBeyondLastLine: false,
                  renderLineHighlight: 'line',
                  cursorBlinking: 'smooth',
                  smoothScrolling: true,
                  lineNumbersMinChars: 3,
                  overviewRulerBorder: false,
                }}
              />
            ) : (
              <div className="flex flex-col items-center justify-center h-full gap-3" style={{ color: '#555' }}>
                <FileCode size={48} strokeWidth={1} />
                <span style={{ fontSize: 13 }}>Open a file to start editing</span>
                <button
                  onClick={handleNewFile}
                  style={{ color: '#75bfbf', fontSize: 12, background: 'none', border: 'none', cursor: 'pointer' }}
                  onMouseEnter={e => (e.currentTarget.style.textDecoration = 'underline')}
                  onMouseLeave={e => (e.currentTarget.style.textDecoration = 'none')}
                >
                  Create new file
                </button>
              </div>
            )}
          </div>

          {/* Bottom Panel */}
          <div style={{ height: 220, background: '#1e1e1e', borderTop: '1px solid #252526' }} className="flex flex-col flex-shrink-0">

            {/* Panel Tab Headers */}
            <div style={{ background: '#252526' }} className="flex items-stretch flex-shrink-0">
              {(['terminal', 'problems'] as const).map(panel => (
                <button
                  key={panel}
                  onClick={() => setActivePanel(panel)}
                  style={{
                    fontSize: 11,
                    fontWeight: 700,
                    letterSpacing: '0.08em',
                    textTransform: 'uppercase',
                    padding: '6px 16px',
                    border: 'none',
                    borderBottom: activePanel === panel ? '1px solid #75bfbf' : '1px solid transparent',
                    background: activePanel === panel ? '#1e1e1e' : 'transparent',
                    color: activePanel === panel ? '#fff' : '#888',
                    cursor: 'pointer',
                    display: 'flex',
                    alignItems: 'center',
                    gap: 6,
                    transition: 'color 0.1s',
                  }}
                  onMouseEnter={e => { if (activePanel !== panel) (e.currentTarget as HTMLButtonElement).style.color = '#ccc'; }}
                  onMouseLeave={e => { if (activePanel !== panel) (e.currentTarget as HTMLButtonElement).style.color = '#888'; }}
                >
                  {panel === 'problems' && errorCount > 0 && (
                    <span style={{
                      background: '#f14c4c',
                      color: '#fff',
                      fontSize: 9,
                      fontWeight: 700,
                      padding: '1px 5px',
                      borderRadius: 10,
                      lineHeight: 1.6,
                    }}>
                      {errorCount}
                    </span>
                  )}
                  {panel === 'terminal' ? 'Terminal' : 'Problems'}
                </button>
              ))}
            </div>

            {/* Terminal */}
            {activePanel === 'terminal' && (
              <div
                className="flex-1 overflow-y-auto"
                style={{ padding: '10px 16px', fontFamily: '"JetBrains Mono", "Fira Code", Consolas, monospace', fontSize: 13, lineHeight: 1.6 }}
              >
                {terminalLines.length === 0 ? (
                  <span style={{ color: '#555', fontStyle: 'italic' }}>Press Run to compile and execute</span>
                ) : (
                  terminalLines.map((line, i) => (
                    <div
                      key={i}
                      style={{
                        color:
                          line.type === 'command' ? '#75bfbf' :
                          line.type === 'error'   ? '#f47272' :
                          line.type === 'success' ? '#89d185' :
                          line.type === 'output'  ? '#d4d4d4' :
                                                    '#888',
                      }}
                    >
                      {line.text}
                    </div>
                  ))
                )}
                <div ref={terminalEndRef} />
              </div>
            )}

            {/* Problems */}
            {activePanel === 'problems' && (
              <div className="flex-1 overflow-y-auto">
                {problems.length === 0 ? (
                  <div style={{ padding: '10px 16px', color: '#888', fontSize: 13, display: 'flex', alignItems: 'center', gap: 8 }}>
                    <CheckCircle2 size={14} style={{ color: '#89d185' }} />
                    No problems detected
                  </div>
                ) : (
                  problems.map((p, i) => (
                    <div
                      key={i}
                      onClick={() => jumpToLine(p.line, p.col)}
                      style={{ borderBottom: '1px solid #252526', cursor: 'pointer' }}
                      className="flex items-start gap-3 px-4 py-2 hover:bg-[#2a2d2e] group"
                    >
                      <AlertCircle size={14} style={{ color: '#f47272', flexShrink: 0, marginTop: 1 }} />
                      <div className="flex-1 min-w-0">
                        <span style={{ color: '#d4d4d4', fontSize: 13 }}>{p.message}</span>
                        {p.detail && (
                          <span style={{ color: '#888', fontSize: 12, marginLeft: 8 }}>{p.detail}</span>
                        )}
                      </div>
                      <div
                        style={{ fontSize: 11, flexShrink: 0, display: 'flex', alignItems: 'center', gap: 3 }}
                        className="text-[#666] group-hover:text-[#75bfbf]"
                      >
                        <span>{activeTab?.name ?? 'file'}:{p.line}:{p.col}</span>
                        <ChevronRight size={10} />
                      </div>
                    </div>
                  ))
                )}
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Status Bar */}
      <div
        style={{
          height: 24,
          background: '#0a5f5f',
          color: '#fff',
          fontSize: 11,
          padding: '0 12px',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          flexShrink: 0,
          userSelect: 'none',
        }}
      >
        <div style={{ display: 'flex', alignItems: 'center', gap: 16 }}>
          {errorCount > 0 ? (
            <button
              onClick={() => setActivePanel('problems')}
              style={{ background: 'none', border: 'none', color: '#fff', cursor: 'pointer', fontSize: 11, display: 'flex', alignItems: 'center', gap: 4, padding: 0 }}
            >
              <AlertCircle size={11} />
              {errorCount} error{errorCount !== 1 ? 's' : ''}
            </button>
          ) : (
            <span style={{ opacity: 0.7, display: 'flex', alignItems: 'center', gap: 4 }}>
              <CheckCircle2 size={11} />
              No errors
            </span>
          )}
        </div>
        <div style={{ display: 'flex', alignItems: 'center', gap: 16, opacity: 0.9 }}>
          <span>Mini-Go</span>
          <span>Ln {cursorPos.line}, Col {cursorPos.col}</span>
        </div>
      </div>
    </div>
  );
};

export default App;
