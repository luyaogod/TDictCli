---
title: "Executing sub-programs from a parent program with RUN"
source: "fgl-topics/c_fgl_gwa_program_execution.html"
breadcrumb: "Genero Web applications > Executing programs with RUN"
type: "concept"
---

# Executing sub-programs from a parent program with RUN

> Sub-programs can be executed from a parent program with the RUN instruction.

There can be limitations, as the GWA runs in the browser:

1. **GUI type program**: A sub-program typically should be a graphical (UI) application because
   browsers often kill non-GUI background processes that run for several seconds. If a non-GUI
   sub-program runs longer than about 5 seconds it may be terminated, so you should present a GUI
   progress indicator and refresh it regularly (at least once per second) to keep the browser from
   stopping the process.
2. **Environment inheritance**: The sub-program will inherit the environment of the parent
   program.
3. **File system sharing** : The child programs can access some files in the parent app's [filesystem](5125-file-system.md "When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory."): /app,
   /tmp, and /home/packdir directory
   (which can be queried via `fgl_getenv(HOME)`). However, directories created by the
   parent program with files under a new root, for example `os.Path.mkdir("/mypath")`
   are not shared.

The [`RUN`](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") instruction takes a shell
command as parameter. To start another BDL program from the current program, specify [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs."), followed by the program name and
arguments.

As the GWA runs in the browser, there is no real shell, so the `RUN` instruction
can only start an `fglrun`
command:

```
RUN "fglrun prog-name [arg [,...]]" 
 [ RETURNING variable | WITHOUT WAITING ]
```

For more details, see the reference page of the [`RUN`](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") instruction.

## Related links

**Related concepts**  

[Program execution](../09_advanced-features/0828-program-execution.md "This section describes program execution and language instructions related to program execution.")

[Start a GAS app from GWA](5146-start-a-gas-app-from-gwa.md "The runOnServer front call allows you to start an application via the Genero Application Server (GAS) from a Genero Web application (GWA).")

[Genero environment variables](../07_configuration/0506-genero-environment-variables.md "Genero environment variables")

[Operating system environment variables](../07_configuration/0496-operating-system-environment-variables.md "Describes some well-known system environment variables that are used by Genero software components.")
