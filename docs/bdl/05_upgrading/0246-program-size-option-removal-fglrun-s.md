---
title: "Program size option removal (fglrun -s)"
source: "fgl-topics/c_fgl_Migrate_to_240_prog_size_option.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Program size option removal (fglrun -s)"
type: "concept"
---

# Program size option removal (fglrun -s)

> The -s option of fglrun is no longer available.

Before version 2.30 the `-s` option of [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.") was
used to compute the size of program elements such as global and module variables, p-code and
structured data types. Starting with version 2.30, this option reported a size of zero. With version
2.40 the `-s` option is now fully desupported.

The `-s` option was mainly implemented for internal use. Regarding the amount of
memory used by a program, it is recommended that you consider the memory allocated dynamically at
runtime: If you fill large dynamic arrays, or leave a lot of SQL cursors open without freeing them,
the memory footprint of a program can be much larger than the actual size of static elements that
may be reported by the `-s` option.

## Related links

**Related concepts**  

[Optimization](../09_advanced-features/0966-optimization.md "Programming tips and tricks to make your programs run faster.")

[Optimize your programs](../09_advanced-features/0972-optimize-your-programs.md "This section contains programming tips to optimize the execution of your application.")
