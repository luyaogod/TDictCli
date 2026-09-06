---
title: "Visual Studio Code extension"
source: "fgl-topics/c_fgl_CodeEditing_006.html"
breadcrumb: "Programming tools > Source code edition > Visual Studio Code extension"
type: "concept"
description: "The Visual Studio Code editor Microsoft® Visual Studio Code (VS Code) is a powerful source code editor available on various platforms. This topic contains instructions to set up the VS Code extension ..."
---

# Visual Studio Code extension

## The Visual Studio Code editor

Microsoft® Visual Studio Code
(VS Code) is a powerful source code editor available on various platforms.

This topic contains instructions to set up the VS Code extension for Genero BDL. For debugging
programs with VS Code, refer to [Debugging with VS Code](2584-debugging-with-vs-code.md "The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities.").

Code completion, syntax highlighting, goto definition, references, code formatting, diagnostics
and debugging is possible for Genero .4gl and .per sources
in VS Code, after installing the VS Code extension for Genero BDL.

The VS Code extension for Genero BDL is based on command-line tools fglcomp,
fglform and fglrun. Configuration settings to control FGL
tools apply also in the context of VS Code. For example, environment variables such as FGLPROFILE,
FGLSQLDEBUG, FGLLDPATH or FGLSOURCEPATH can be used.

Use the [`IMPORT
FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") technique in .4gl sources, to take full advantage of the
`fglcomp` compiler, fglrun integrated debugger, and VS Code Genero
extension. For new projects, we recommend organizing .4gl sources into [packages](../09_advanced-features/0786-modules-and-packages.md "Programs sources can be organized in packages of modules."), and utilizing VS Code [multi-root workspaces](https://code.visualstudio.com/docs/editing/workspaces/multi-root-workspaces).

## Installing VS Code

To start, you need to download and install Microsoft Visual Studio Code.

This documentation page does not explain how to use VS Code, you need to be familiar with this
editor.

See <https://code.visualstudio.com/> for more
details.

## Setup Genero environment

Before installing the VS Code extension for Genero, and before starting VS Code once Genero
extension is installed, it is mandatory to set up the Genero BDL environment in order to access
Genero tools like fglcomp and `fglform`.

For example, on a Microsoft Windows OS, if you start VS Code from the start menu, and
Genero BDL environment is not implicitly set for the user or system, the VS Code extension will not
work.

> **Tip:**
>
> If VS Code editor features such as suggestions, diagnostics, popup on mouse hover are not
> working, then it is probably because the Genero environment has not been set before starting VS
> Code. To check if the Genero environment is properly set from VSCode: Open a
> TERMINAL, and verify that the fglcomp command can be
> launched. For example, check the locale settings with fglcomp -i or compile a
> source with fglcomp source.4gl .

For more details about Genero environment setup, see [Installing Genero BDL](../04_installation/0048-installing-genero-bdl.md "This section provides Genero BDL installation instructions.").

## Installing Genero BDL extension into VS Code

The Genero BDL extension can be installed in your VS Code environment with the `code
--install-extension` command, using the .vsix file provided in the
FGLDIR/lib directory. For example:

```
$ code --install-extension $FGLDIR/lib/genero-fgl-0.0.29.vsix
Installing extensions...
Extension 'genero-fgl-0.0.28.vsix' was successfully installed.
```

Alternatively, you can start VS Code, go to the EXTENSIONS view, and load
the Genero BDL .vsix extension file from FGLDIR/lib.

When VS Code is started, you should see the Genero extension in the
EXTENSIONS view. Select the Genero extension, read the description and check
the settings.

To uninstall the Genero extensions form the command
line:

```
$ code --uninstall-extension $FGLDIR/lib/genero-fgl-0.0.29.vsix
Uninstalling fourjs.genero-fgl...
Extension 'fourjs.genero-fgl' was successfully uninstalled!
```

When upgrading to a more recent version of Genero BDL, re-install the VS Code extension file to
get bug fixes and new features of this plugin. The version number of the .vsix
file may change with a new Genero BDL release.

## Check that the extension is properly installed

After installing the Genero BDL extension in your VS Code, do the following quick checks to make
sure that the enviroment is properly set and the extension is working:

1. First make sure that the Genero BDL environment is properly set (FGLDIR, PATH). To check this,
   run `fglcomp -i` from the command line, to verify also the locale settings (LANG,
   LC\_ALL).
2. Start VS Code
3. Open a folder with .4gl sources.
4. Open a .4gl source into the code editor.
5. Start typing code, you should get auto-completion. Alternatively, you can hover with the mouse on
   a function name in a call, or expression to see a hint window popup with the function
   definition.

Syntax highlighting does not mean that all VS Code features work correctly.

## Configuring global VS Code settings

Perform the following steps to set up VS Code for Genero BDL:

1. Start VS Code
2. Configure file charset / encoding settings (recommended).

   By default, VS Code uses UTF-8 character set encoding. If your .4gl and
   .per sources are using a different encoding such as ISO-8859-15, you want to
   set the following configuration parameters:
   - Open VS Code settings
   - Select the User panel for global settings
   - Go to Text Editor > Files > Encoding and choose your default encoding. You can also go to Text Editor > Files > Auto Guess Encoding and enable this option, to let the editor guess the character set encoding when
     opening files.
3. Configure the editor to exclude `.42?` files (recommended).

   By default, the VS Code Explorer will show the .42m and
   .42f files, which is annoying.
   - Open VS Code settings
   - Select the User panel for global settings
   - Go to Text Editor > Files > Exclude
   - Add `**/*.42?` to the pattern list
4. Configure the editor to avoid suggestion selection on `Enter` (your choice).

   Suggestion items can always be selected with TAB. By default,
   ENTER can also be used for suggestion item selection: When you input a word
   at the end of the line, and you want to go to the next line by pressing
   ENTER, the word is replaced by the first item in the list of completions; you
   have not changed the line and you must press ENTER again. To use
   ENTER key for new line creation, and use only the TAB
   key for suggestion item selection:

   - Open VS Code settings
   - Select the User panel for global settings
   - Go to Text Editor > Suggestions > Accept suggestion on Enter
   - Set the option to off
5. Configure the debugger to get inline values (recommended).

   Whenever you are in an active debug session, the current values of variables can be displayed
   directly in the editor with the "Inline Values" option.

   - Open VS Code settings
   - Select the User panel for global settings
   - Go to Features > Debug > Inline Values
   - Set the option to on

## Create VS Code tasks for a project directory

Perform the following steps to set up VS Code for each of your Genero BDL projects:

1. Start VS Code
2. Open a folder with Genero BDL sources (there must be at least one .4gl
   source file in this directory in order to properly configure VS Code for Genero BDL). VS Code will
   consider this folder as the root dir for your project. Configuring VS Code tasks will automatically
   create a .vscode directory with configuration files.
3. Define a build task for Genero programs for this folder:
   - Open a .4gl source from the file explorer
   - Select Terminal > Configure Task...

     > **Tip:**
     >
     > Instead of Terminal > Configure Task..., you can use the Terminal > Configure Default Build Task... option, to create the default build task that will be automatically selected when
     > pressing the corresponding keyboard shortcut. The default build task is marked with the
     > `group.isDefault` property in the task definition:
     >
     > ```
     > "group": {
     >     "kind": "build",
     >     "isDefault": true
     > },
     > ```
   - Select a task to compile .4gl sources such as 4gl: build all 4gl
     files

     This will create the .vscode/tasks.json file if it does not exists, open the
     file for edition, and append the selected task to the file.
   - Change the `label` property, to define your own build task name, so it can be
     distinguished from the detected tasks, when selecting a build task. For example, use the name
     4gl: my program build task.
   - If needed, modify other properties of the new created build task.
   - Save the tasks.json file (CTRL-S)
4. Define a build task for Genero form files:
   - Open a .per source from the file explorer
   - Select Terminal > Configure Task...
   - Select a task to compile .per sources such as per: build active
     per file

     This will create the .vscode/tasks.json file if it does not exists, open the
     file for edition, and append the selected task to the file.
   - Change the `label` property, to define your own build task name, so it can be
     distinguished from the detected tasks, when selecting a build task. For example, use the name
     per: my form build task.
   - If needed, modify other properties of the new created build task.
   - Save the tasks.json file (CTRL-S)
   - The tasks.json file content should then look like this:

     ```
     {
     	"version": "2.0.0",
     	"tasks": [
     		{
     			"type": "genero-fgl",
     			"label": "4gl: my program build task",
     			"command": "fglcomp",
     			"args": [
     				"-M",
     				"--make",
     				"-Wall",
     				"-Wno-stdsql",
     				"-k",
     				"*.4gl"
     			],
     			"options": {
     				"cwd": "${fileDirname}"
     			},
     			"problemMatcher": [
     				"$fglcomp"
     			],
     			"group": {
     				"kind": "build",
     				"isDefault": true
     			},
     			"detail": "fglcomp -M --make -Wall -Wno-stdsql -k *.4gl"
     		},
     		{
     			"type": "genero-fgl",
     			"label": "per: my form build task",
     			"command": "fglform",
     			"args": [
     				"-M",
     				"${file}"
     			],
     			"options": {
     				"cwd": "${fileDirname}"
     			},
     			"problemMatcher": [
     				"$fglcomp"
     			],
     			"group": "build",
     			"detail": "fglform -M ${file}"
     		}
     	]
     }
     ```

     Notes:
     1. Adapt the tasks with VS Code predefined variables. For example, you may want to define the
        `"cwd"` parameter with `${workspaceFolder}` instead of
        `"${fileDirname}`.
5. To compile .4gl sources:
   - Open a .4gl source from the file explorer
   - Select Terminal > Run Build Task....
   - If no default build task is defined and multiple tasks are found by VS Code for
     .4gl sources, the task pick list opens, and you can for example select
     4gl: my program build task
   - Check the TERMINAL panel for executed commands
6. To compile .per form files:
   - Open a .per form file from the file explorer
   - Select Terminal > Run Task....
   - Select the task per: my form build task
   - Check the TERMINAL panel for executed commands
7. Define launch tasks for Genero programs:
   - Open a .4gl source file of your project.

     Click on Run and Debug arrow on the left panel
   - Click the create a launch.json file link
   - Select GeneroFgl Debug: Launch: This will automatically add the
     "(4gl) Launch" task to the new created launch.json
     file.
   - With launch.json file opened, click on the Add
     Configuration... button
   - Select GeneroFgl Debug: Attach: This will automatically add the
     "(4gl) Attach" task to the new created launch.json
     file.
   - The launch.json file content should then look like this:

     ```
     {
         "version": "0.2.0",
         "configurations": [
             {
                 "type": "fgldb-dap",
                 "request": "attach",
                 "name": "(4gl) Attach",
                 "processId": "${command:pickProcess}"
             },
             {
                 "type": "fgldb-dap",
                 "request": "launch",
                 "name": "(4gl) Launch",
                 "program": "${fileDirname}/${fileBasenameNoExtension}",
                 "args": [],
                 "stopAtEntry": true,
                 "externalConsole": false,
                 "cwd": "${fileDirname}"
             }
         ]
     }
     ```

     Notes:
     1. Depending on the source directory structure, you may want to define the
        `"program"` and `"cwd"` parameters with
        `${workspaceFolder}` instead of `${fileDirname}`.
     2. Adapt the `"stopAtEntry"` parameter to stop for debug at the first executable line
        in the `MAIN` block (`true`), or stop for debug at the first
        breakpoint (`false`).
     3. Adapt the `"externalConsole"` parameter: When set to `false`, VS
        Code runs the program to debug in it's current shell. When `true`, the program gets
        run in a new shell and will pick up the full shell environment (for example, on Linux:
        .bash\_profile and .bashrc) - this is mandatory if your
        Genero and additional tools (like JDK) environment settings are defined in the user shell
        scripts.
     4. Command-line arguments can be passed to the program with the `"args"` field. Each
        argument needs to be specified as an individual string in a JSON
        array:

        ```
        "args": ["-g", "--username", "tom", "--port", "2134"]
        ```
     5. Environment variables can be defined for the `"launch"` request, by adding an
        `"env"` object, with a set of properties in the form
        `"var-name":"value"` pairs. Use
        `null` as value, to unset an existing environment variable:

        ```
        "env":{ "TMPDIR" : "/tmp", "TMP" : null }
        ```

        > **Note:**
        >
        > Use the launch configuration option `"env"`, only if this launch target requires
        > individual environment variables. General environment variables can be set for the terminal (enter
        > settings, enter ‘terminal env’ in the search bar). A typical use case is `FGLSERVER`.
        > If the scope of the setting is `User` then this value of `FGLSERVER`
        > is used in any terminal of a VS Code remote connection.
     6. Consider adding a pre-launch task to automatically recompile your source code, before starting
        the debugger:

        ```
        "preLaunchTask": "4gl: my program build task"
        ```

     > **Tip:**
     >
     > Read also [No-configuration debugging](2584-debugging-with-vs-code.md).

## Configure VS Code for remote work

VS Code supports the Remote-Tunnels extension and the Remote-SSH extension to code, to build and
debug source Genero application code on a remote machine.

If the host with Genero development environment and application sources does not have an SSH
server, or when no SSH client is available on your workstation, it is possible to install the VS
Code [Remote-Tunnels](https://marketplace.visualstudio.com/items?itemName=ms-vscode.remote-server) extension and the VS Code [command-line interface (CLI)](https://code.visualstudio.com/docs/editor/command-line#_create-remote-tunnel) component on the server, to either connect from
a workstation where a VS Code instance executes, or from a web browser without the need to install
VS Code on the workstation. For more details, see VS Code documentation: [Developing
with Remote Tunnels](https://code.visualstudio.com/docs/remote/tunnels).

When SSH server is available on the server and an SSH client is installed on the workstation, it
is possible to use the [Remote-SSH](https://code.visualstudio.com/docs/remote/ssh) extension. Here are the steps to set up Remote-SSH for Genero FGL
on a remote machine:

1. Prerequisite: You need an OpenSSH client on the workstation, that is compatible to Remote-SSH.
   Refer to [Remote-SSH compatible OpenSSH clients](https://code.visualstudio.com/docs/remote/troubleshooting#_installing-a-supported-ssh-client).
2. Install VS Code on your workstation.
3. Install the Remote-SSH extension : Select View > Extensions, search Remote-SSH in the marketplace and install it.
4. On the remote development host:
   - Make sure that the SSH server and TCP port are reachable.
   - Check that a Genero development environment is available and licensed.
   - Write a shell script to set up the Genero environment and development environment (FGLDIR,
     PATH). Define FGLSERVER to use the TCP port that is port-forwarded by the SSH connection in VS Code.
     Source this shell script in the environment file of the user shell, like
     ~/.bashrc.
   - To check this setting, start an ssh session from the command line, and test
     fglcomp -V .
5. Follow the instructions in Remote-SSH extension help, to set up a remote ssh connection to the
   host where the source code resides:
   - Select View > Command Palette...
   - In the Command Palette prompt, type Remote-SSH: Connect to host...
   - Enter the ssh command to connect to the remote host and set up the Genero
     development environment. If you want to display application forms on a local front-end of the
     workstation in a secure way, use remote port forwarding with the `-R` option of the
     ssh command. For
     example:

     ```
     ssh mike@devhost -R 127.0.0.1:6400:127.0.0.1:6400
     ```

     with TCP port
     forwarding, fglrun processes started by VS Code on the remote server will display to the front end
     listening on a local workstation TCP port.
   - Enter remote user password when asked.
   - A new VS Code session should open.
6. Install the Genero extension for the remote connection: Select View > Extensions, click on the ... three-dots then Install from
   VSIX..., then browse the remote computer and select the Genero-FGL extension from
   FGLDIR/lib.
7. Select File > Open Folder...: This should let you browse the disk on the remote machine.
8. Make sure you have the parameter `externalConsole` set to `true` in
   the launch.json file.
9. Use local VS Code session as if the environment was local (see previous section for build tasks
   and debugger setup)

If server-side environment settings are changed (in ~/.bashrc,
~/.profile), you must restart the VS Code SSH server component
(`.vscode-server`) as follows:

1. Select View > Command Palette...
2. Find and execute Remote-SSH: Kill VS Code Server or Host...
3. Select the ssh connection to the host.

## VS Code usage tips

**Location of database schema file**

To avoid the (local or server side) definition of the [FGDBPATH](../07_configuration/0522-fgldbpath.md "Defines a list of paths to database schema files for compilers.") env var, in order to find a database schema
file: A good practice is to create a symbolic link to the .sch schema file at
the root of your source tree / VS Code project: The schema file will then be found automatically by
the background fglcomp language server process.

**Restarting the fglcomp language server process**

In some cases (for example, when the environment variables change or when a new fglcomp version
has been installed in the same environment), you may want to restart the background
fglcomp language server process.

To restart your VS Code session and the fglcomp background process:

1. Select View > Command Palette...
2. Search and execute: Developer: Reload Window

## Related links

**Related concepts**  

[Debugging with VS Code](2584-debugging-with-vs-code.md "The Genero BDL VS Code extension can be used to debug programs with VS Code debugger capabilities.")

[Source code completion](2543-source-code-completion.md "Source code completion")
