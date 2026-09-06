---
title: "Improved compilation time"
source: "fgl-topics/c_fgl_Migrate_to_300_compile_time.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > Improved compilation time"
type: "concept"
---

# Improved compilation time

> The fglcomp and fglform compilers have been reviewed to achieve faster compilation.

A Genero project can be very large, with thousands of .4gl
source files to compile. Compilation time can be an issue when the whole set of
sources needs to be compiled every day, or several times a day.

In Genero 3.00, the [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler has been improved to
deliver better performances. Depending on the content of the source file, the compiler can complete
the process twice as fast.

Loading .sch database schema files has also been improved. Using
huge schema files with several thousands lines is no longer an issue. This is especially useful when
compiling forms that define fields based on database columns in a schema file.

## Related links

**Related concepts**  

[fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.")

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
