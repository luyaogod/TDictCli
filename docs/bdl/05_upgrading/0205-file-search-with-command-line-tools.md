---
title: "File search with command line tools"
source: "fgl-topics/c_fgl_Migrate_to_310_tools_filename_ext.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > File search with command line tools"
type: "concept"
---

# File search with command line tools

> Filename extensions are kept by command line tools.

Starting with version 3.10.00 (fix FGL-4929), the BDL command line tools [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks."), [fglmkmsg](../13_programming-tools/2518-fglmkmsg.md "The fglmkmsg tool compiles .msg message files into a binary version used by programs."), [fglmkstr](../13_programming-tools/2522-fglmkstr.md "The fglmkstr tool compiles .str localized string resource files."), [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") and [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") will no longer try to replace an
invalid file extension by the expected file extension, in order to find the file.

Before this version, if the specified filename was not found, the tools were removing the
filename part after the dot (assuming it is a file extension), and replacing it by
.4gl, .msg, .str,
.per or .42r/.42m, before retrying to
open the file.

As a result, it was not possible to use the dot character as part of the filename, and the behavior
could lead to unexpected
results:

```
fglrun foo.bar       =>   loads foo.42m if the file exists
fglrun foo.42r       =>   loads foo.42m if the file exists
fglrun foo.42m       =>   loads foo.42r if the file exists
```

Starting with 3.10.00, for example, the fglrun command is more strict, but
allows dots in the
filename:

```
fglrun foo.42m       =>   loads foo.42m if the file exists
fglrun foo.42r       =>   fails if only foo.42m exists
fglrun foo.bar       =>   fails if only foo.42m exists
fglrun foo.bar.42m   =>   loads foo.bar.42m if the file exists
fglrun foo.bar       =>   loads foo.bar.42m if the file exists
-- Consider using no file extension with fglrun:
fglrun foo           =>   loads foo.42m if the file exists
```

## Related links

**Related concepts**  

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
