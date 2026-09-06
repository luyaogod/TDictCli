---
title: "FGLDBPATH"
source: "fgl-topics/c_fgl_EnvVariables_FGLDBPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLDBPATH"
type: "concept"
---

# FGLDBPATH

> Defines a list of paths to database schema files for compilers.

The [fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") and [fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.") compilers need
database schema files to compile source modules and forms. The path to the database
schema files can be specified with FGLDBPATH.

If FGLDBPATH is not defined, the current directory is the default path for the database
schema files. You can provide a list of paths, separated by the operating system specific
path separator. FGLDBPATH is only used in development.

FGLDBPATH must contain a list of paths, separated by the operating system specific path
separator. The path separator is **":"** on UNIX™ platforms and  **";"** on Windows® platforms.

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
