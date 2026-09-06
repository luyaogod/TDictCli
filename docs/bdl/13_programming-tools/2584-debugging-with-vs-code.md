---
title: "Debugging with VS Code"
source: "fgl-topics/c_fgl_Debugger_012.html"
breadcrumb: "Programming tools > Integrated debugger > Debugging with VS Code"
type: "concept"
---

# Debugging with VS Code

> The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities.

## VS Code extension setup

Genero BDL provides a VS Code extension to be plugged into your VS Code editor. Setup
instructions to install the VS Code extension for Genero BDL are provided in [Visual Studio Code extension](2545-visual-studio-code-extension.md).

Once the VS Code extension for Genero BDL is installed, you can debug your programs from the VS
Code editor.

Use the [`IMPORT
FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") technique in .4gl sources, to take full advantage of the
`fglcomp` compiler, fglrun integrated debugger, and VS Code Genero
extension. For new projects, we recommend organizing .4gl sources into [packages](../09_advanced-features/0786-modules-and-packages.md "Programs sources can be organized in packages of modules."), and utilizing VS Code [multi-root workspaces](https://code.visualstudio.com/docs/editing/workspaces/multi-root-workspaces).

## No-configuration debugging

Starting with FGL version 5.01.06, the VS Code extension (version 0.0.28) for Genero BDL supports
debugging without the need to [setup launch
configurations](2545-visual-studio-code-extension.md).

To begin a debug session without configuration, the fglrun command must be
executed in the VS Code terminal in the form `fglrun -d ...`, or `fglrun
--da-listen port ...`, when the [debug-server](2585-using-the-debug-server.md "The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger.") is needed.

On Unix/Linux/macOS, VS Code listens for debug requests on a Unix domain socket, and listens to a
named pipe when on Microsoft Windows. In a VS Code context with Genero BDL extension installed, the
[`VSCODE_FGLDA_IPCPATH`](../07_configuration/0540-vscode-fglda-ipcpath.md "Defines the communication channel for the FGL debugger in VS Code.") environment variable is automatically defined to the
debug communcation channel. When fglrun is started in debug mode and
`VSCODE_FGLDA_IPCPATH` is set, fglrun notifies VS Code and waits
for the VS Code debugger to attach.

To start a no-configuration debug session:

1. Open the .4gl source containing the `MAIN` block.
2. Go to the TERMINAL panel, and select a shell like bash.
3. Start the FGL VM in debug mode with the `-d`
   option:

   ```
   fglrun -d myprog.42m
   ```

   > **Tip:**
   >
   > When using the debug-server mode with port zero (fglrun
   > --da-listen=0), VS Code TCP port forwarding must be disabled. For more details, read [Automatic port allocation (--da-listen 0)](2585-using-the-debug-server.md).
4. The VS Code debug session should start.

   > **Tip:**
   >
   > If the debug session does not start, make sure that the `VSCODE_FGLDA_IPCPATH`
   > enviroment variable is defined. If it is not defined, close the terminal, ensure a
   > .4gl file is opened in the editor, open the terminal again.

## Using the Launch configuration

With VS Code having a folder opened with the source code of your project, start a launch debug
session as follows:

1. Open the .4gl source containing the `MAIN` block.
2. Select Run > Start Debugging: This will execute the selected launch configuration, by default it will be
   (4gl) Launch.
3. The program starts in debug mode, and debugger stops at the first instruction of the program. You
   can control this with the `"stopAtEntry"` property in your launch configuration
   (check tasks.json configuration file).
4. Go to the DEBUG CONSOLE panel: The prompt at the bottom allows you to
   evaluate expressions (for example, type sqlca + enter)
5. Go to the source code, and add a break point in one of the next lines.
6. Continue the program with F5, the program continues until the break
   point.
7. Right-click on the break point, and edit the break point (to add for example a condition)

## Using the Attach configuration

VS Code can attach to a running `fglrun` process.

With VS Code having a folder opened with the source code of your project, start an attach debug
session as follows:

1. Go to the TERMINAL panel or go to another terminal emulator, and start a
   program with GUI interface, using the same Genero environment.
2. In VS Code, select the Run and Debug option, choose (4gl)
   Attach, start debugging (F5), then choose a running
   fglrun process.
3. The program should stop at the current instruction, and you can debug the source code.

Several debug sessions are possible. For example, when a program starts a child program with
`RUN`, another debug session will start.

It is also possible to attach to the debug-server, when it's not possible to attach
directly to the fglrun process (for example, if fglrun is
started by a different user). For more details, read [Using the debug-server](2585-using-the-debug-server.md "The Genero BDL debug-server is a proxy for fglrun processes which can not be accessed directly by the debugger.").

## Debugging a TUI program

In order to debug programs using the TUI mode (FGLGUI=0), do the following:

1. In the environment, make sure that TERM and INFORMIXTERM variables are properly set to display
   program forms in your terminal emulator, for example:

   ```
   $ export TERM=xterm
   $ export INFORMIXTERM=terminfo
   ```
2. Start VS Code with this environment set.
3. In the launch.json file, modify (or copy and modify) the "(4gl)
   launch" configuration to set the `externalConsole` property to
   `true`:

   ```
   "externalConsole": true,
   ```

   This will enable TTY output of
   fglrun to the TERMINAL panel.
4. Select the program source to execute.
5. Select the Run and Debug option and start the modified launch
   configuration.
6. Click on Continue or press F5 to give the control
   to the runtime system.
7. The TUI program interface should display in the TERMINAL panel.

To simplify debugging of a TUI program with VS Code, consider the following tips:

- To switch to the TERMINAL panel, remember using the
  CTRL-` shortcut.
- To automatically give the focus to the TERMINAL panel when continuing with
  F5, configure a custom F5 shortcut as follows:
  1. Go to Settings > Keyboard Shortcuts.
  2. Click on the Open Keyboard Shortcuts button on the top right corner.
  3. Append new key bindings as follows:
     - When paused, to continue (or start) program execution and give the focus to the
       terminal:

       ```
       {
           "key": "f5",
           "command": "runCommands",
           "args": {
               "commands": [
                   "workbench.action.debug.continue",
                   "workbench.action.terminal.focus"
               ]
           },
           "when": "debugState==stopped && debugType==fgldb-dap && terminalIsOpen"
       }
       ```
     - When paused, to step over the current line of code and give the focus to the
       terminal:

       ```
       {
           "key": "f10",
           "command": "runCommands",
           "args": {
               "commands": [
                   "workbench.action.debug.stepOver",
                   "workbench.action.terminal.focus"
               ]
           },
           "when": "debugState==stopped && debugType==fgldb-dap && terminalIsOpen"
       }
       ```
     - When paused, to step into a function and give the focus to the terminal:

       ```
       {
           "key": "f11",
           "command": "runCommands",
           "args": {
               "commands": [
                   "workbench.action.debug.stepInto",
                   "workbench.action.terminal.focus"
               ]
           },
           "when": "debugState==stopped && debugType==fgldb-dap && terminalIsOpen"
       }
       ```
     - When paused, to step out of the current function and give the focus to the
       terminal:

       ```
       {
           "key": "shift+f11",
           "command": "runCommands",
           "args": {
               "commands": [
                   "workbench.action.debug.stepOut",
                   "workbench.action.terminal.focus"
               ]
           },
           "when": "debugState==stopped && debugType==fgldb-dap && terminalIsOpen"
       }
       ```
- To get a separate TERMINAL window:
  1. Open the command palette with CTRL-SHIFT-P
  2. Type "Create a new terminal in the editor area", select the command, this
     creates a new TERMINAL panel in the editor area.
  3. Drag & Drop the new created TERMINAL panel outside the VS Code main
     window.

## Related links

**Related concepts**  

[Visual Studio Code extension](2545-visual-studio-code-extension.md "Visual Studio Code extension")
