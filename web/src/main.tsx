import React from 'react'
import ReactDOM from 'react-dom/client'
import { App } from './App'
import { editorRef } from './SourceView'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)

// 浏览器控制台调试入口
;(window as any).__ed = editorRef
