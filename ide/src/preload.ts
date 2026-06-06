import { ipcRenderer, contextBridge, IpcRendererEvent } from 'electron'

// Map to keep track of original listeners and their wrapped versions
const listenerMap = new WeakMap<Function, (event: IpcRendererEvent, ...args: any[]) => void>();

// --------- Expose some API to the Renderer process ---------
contextBridge.exposeInMainWorld('ipcRenderer', {
  on(channel: string, listener: (...args: any[]) => void) {
    const wrappedListener = (_event: IpcRendererEvent, ...args: any[]) => listener(...args);
    listenerMap.set(listener, wrappedListener);
    ipcRenderer.on(channel, wrappedListener);
  },
  off(channel: string, listener: (...args: any[]) => void) {
    const wrappedListener = listenerMap.get(listener);
    if (wrappedListener) {
      ipcRenderer.off(channel, wrappedListener);
      listenerMap.delete(listener);
    }
  },
  removeAllListeners(channel: string) {
    ipcRenderer.removeAllListeners(channel);
  },
  send(channel: string, ...args: any[]) {
    ipcRenderer.send(channel, ...args);
  },
  invoke(channel: string, ...args: any[]) {
    return ipcRenderer.invoke(channel, ...args);
  },
})
