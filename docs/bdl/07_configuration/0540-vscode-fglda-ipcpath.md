---
title: "VSCODE_FGLDA_IPCPATH"
source: "fgl-topics/c_fgl_EnvVariables_VSCODE_FGLDA_IPCPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > VSCODE_FGLDA_IPCPATH"
type: "concept"
---

# VSCODE_FGLDA_IPCPATH

> Defines the communication channel for the FGL debugger in VS Code.

The VSCODE\_FGLDA\_IPCPATH environment variable defines the Unix domain socket or
Windows named pipe for the FGL debugger in VS Code.

There is no need to set this environment variable: It is automatically initialized in the context
of VS Code with Genero extension, for debugging purpose.

For more details, read [No-configuration debugging](../13_programming-tools/2584-debugging-with-vs-code.md).
