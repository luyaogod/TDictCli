---
title: "FGLSOURCEPATH"
source: "fgl-topics/c_fgl_EnvVariables_FGLSOURCEPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLSOURCEPATH"
type: "concept"
---

# FGLSOURCEPATH

> Defines a list of paths to program source files.

The debugger needs to access the source files to display program code. By default, the current
directory and the directories defined by [FGLLDPATH](0528-fglldpath.md "Defines a list of paths to find program modules.") are used to find source files.

The FGLSOURCEPATH environment variable is provided to distinguish execution directories
(containing .42m files), from source directories (containing
.4gl files), when the sources are not located in the same directory as the
pcode files.

FGLSOURCEPATH must contain a list of paths, separated by the operating system specific path
separator. The path separator is **":"** on UNIX™
platforms and  **";"** on Windows™ platforms.

UNIX
example:

```
$ FGLSOURCEPATH="/usr/app/source:/home/scott/sources"
$ export FGLSOURCEPATH
```

Windows
example:

```
C:\> set FGLSOURCEPATH=C:\app\sources;C:\scott\sources
```

## Related links

**Related concepts**  

[Integrated debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs.")
