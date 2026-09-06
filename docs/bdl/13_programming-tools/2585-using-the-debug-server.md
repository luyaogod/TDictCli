---
title: "Using the debug-server"
source: "fgl-topics/c_fgl_Debugger_014.html"
breadcrumb: "Programming tools > Integrated debugger > Using the debug-server"
type: "concept"
---

# Using the debug-server

> The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger.

## Purpose of the debug-server

The Genero BDL runtime system provides an option to start [fglrun](2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") in debug-server mode, and allow the real debugger to
attach to a running process, that cannot be accessed directly.

This tool acts like a proxy for fglrun. It has been implemented especially to
debug Genero programs with the [VS Code extension](2584-debugging-with-vs-code.md "The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities.").

The debug-server solves the following cases:

- Avoids maintaining multiple launch.json configuration to run/debug programs
  with different argument sets and environment variables.
- Debug a program started by an OS user, which is not the user running the debugger. In this case,
  for security reasons, a SIGTRAP signal cannot be sent to the running process.
- Debug a program running on a remote computer where the VS Code server component is not
  available.

## Starting the debug-server

The debug-server must be started on the machine where the fglrun
program to debug executes, using the same OS user as the fglrun process.

If the fglrun process is already running, find the process id (pid) of the
program and bind the debug-server to that process with the following command, using the
`--da-pid`
option:

```
$ fglrun --da-listen [host]:port --da-pid process-id
```

If you want to start and debug a new program instance, use the `--da-run` option
with program name and program
arguments:

```
$ fglrun --da-listen [host]:port --da-run program arg1 ...
```

After starting the debug-server, go to the debugger tool, and attach to the
debug-server by using the TCP port specified. With VS Code, you need to configure a specific launch
configuration as described in Start a debugging session with a launch configuration.

> **Tip:**
>
> In order to debug an program on a remove computer, consider using ssh port
> forwarding for security reasons. In this case, the debug-server must listen to the
> localhost interface.

The debug-server is automatically stopped when the debug session is finished. For a
new debug session, restart the debug-server.

## Automatic port allocation (`--da-listen 0`)

The debug-server can be started with a TCP port zero (`fglrun --da-listen 0
...`), to automatically select a free port in the range of ephemeral TCP ports
(on Linux, the range is typically 32768 to 60999).

For example, you can use `fglrun --da-listen 0` when the 6400 TCP port is already
in use:

```
$ fglrun --da-listen 6400 --da-pid 12345
error: listen failed: Address already in use
```

By default, VS Code automatically detects whether new processes are running and forwards the
appropriate ports on which these processes are listening. With `fglrun --da-listen
0`, a new port will be forwarded and this is unwanted here.

Consequently, you want to disable VS Code automatic port forwarding, for a range of TCP ports
that can be allocated dynamically when using the `fglrun --da-listen 0` form of the
debug-server.

In the VS Code settings, find the `remote.portAttributes` parameter, go to
Remote: Ports Attributes, then Edit in settings.json
and specify a range of port with `"onAutoFormard"` set to
`"ignore"`:

```
"remote.portsAttributes": {
    "32768-60999": {
        "onAutoForward": "ignore"
    },
    ...
```

## Start a debugging session with a launch configuration

In your VS Code launch.json file, create a new configuration as follows,
using the `"connect"` property instead of `"processId"`, and specify
the host and port the debug-server is listening (if you don't specify the host, it
defaults to localhost):

```
        {
            "type": "fgldb-dap",
            "request": "attach",
            "name": "(4gl) Attach TCP port",
            "connect": { "host": "jupiter", "port": 4711 }
        },
```

Make sure that the debug-server is started as described in Starting the debug-server, and begin a debug session as follows:

1. In VS Code, select the Run and Debug option, choose (4gl)
   Attach TCP port, start debugging (F5).
2. The program should stop at the current instruction, and you can debug the source code.

This solution can be used to debug a program running on the same computer, or on a remote
computer, when the debug-server and debugger target program are executing.

## Related links

**Related concepts**  

[Visual Studio Code extension](2545-visual-studio-code-extension.md "Visual Studio Code extension")
