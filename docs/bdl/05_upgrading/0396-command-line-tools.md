---
title: "Command line tools"
source: "fgl-topics/c_fgl_MigI4GL_060.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Command line tools"
type: "concept"
---

# Command line tools

> I4GL and FGL provide different compiler commands with different options.

## Review makefiles and program start scripts

Since Genero BDL provides different compiler and program execution commands than IBM® Informix® 4GL, it is mandatory to
review the application compilation procedure, as well as the tools that lauch your application
programs.

The following table list between I4GL commands and Genero commands:

| Informix 4GL command | Genero BDL command |
| --- | --- |
| c4gl / fglpc | fglcomp |
| form4gl | fglform |
| mkmessage | fglmkmsg |
| fglgo | fglrun |

## The I4GL -ansi compiler option

IBM Informix 4GL
(I4GL) provides the `-ansi` option with c4gl and
fglpc compilers, in order to check all SQL statements for compliance with the
ANSI/ISO standard for SQL syntax.

The SQL standard syntax has evolved since the `-ansi` option has been introduce in
I4GL, so this option is no longer relevant.

Genero BDL provides a similar option: `-Wstdsql`, to detect SQL statements that
may be problematic when using a database engine different from Informix.

## Related links

**Related concepts**  

[Compiling 4GL to C](0346-compiling-4gl-to-c.md "Genero's use of a p-code architecture removes restrictions on which platforms you use to develop your application.")

[fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")

[fglform](../13_programming-tools/2514-fglform.md "The fglform tool compiles form specification files into XML formatted files used by programs.")
