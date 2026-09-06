---
title: "Writing timestamp information in p-code modules"
source: "fgl-topics/c_fgl_Migrate_to_211_pcode_timestamp.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.11 upgrade guide > Writing timestamp information in p-code modules"
type: "concept"
---

# Writing timestamp information in p-code modules

> A compilation timestamp is no longer automatically written to p-code files, when the source code is not modified.

In version 2.10, the 42m p-code files were by default stamped with a compilation timestamp. This
information changed after every compilation, even if the source code was not modified.

Starting with version 2.11, the timestamp information is no longer written to p-code files by
default, allowing 42m file comparison, checksum creation, or storage of 42m file in versioning
tools.

To force the compler to write a timestamp in p-code modules, use the [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") `--timestamp` option:

```
$ fglcomp --timestamp mymodule.4gl
$ fglrun -b mymodule.42m
2008-12-24 11:22:33 2.11.05-1169.84 /home/devel/stores/mymodule.4gl 15
```

## Related links

**Related concepts**  

[Compiling program code files (.4gl)](../13_programming-tools/2534-compiling-program-code-files-4gl.md "The .4gl source files must be compiled to .42m p-code files, in order to be loaded by the runtime system.")
