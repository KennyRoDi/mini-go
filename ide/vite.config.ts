import { defineConfig } from 'vite'
import path from 'node:path'
import react from '@vitejs/plugin-react'
import electron from 'vite-plugin-electron/simple'

export default defineConfig({
  plugins: [
    react(),
    electron({
      main: {
        entry: 'src/main.ts',
      },
      preload: {
        input: path.join(__dirname, 'src/preload.ts'),
      },
      renderer: {},
    }),
  ],
})
