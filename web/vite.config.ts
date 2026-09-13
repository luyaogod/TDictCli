import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// 前端占位:调试服务的 /api、/mcp 代理已随调试功能一并移除;
// 后续新增后端接口时在此补 server.proxy。
export default defineConfig({
  plugins: [react(), tailwindcss()],
})
