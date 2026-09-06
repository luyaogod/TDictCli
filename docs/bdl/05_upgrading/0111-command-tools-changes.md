---
title: "Command tools changes"
source: "fgl-topics/c_fgl_Migrate_to_501_fgl_tools.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.01 upgrade guide > Command tools changes"
type: "concept"
---

# Command tools changes

> Modifications to consider regarding command line tools.

This topic describes changes that may need code review.

See also [new 5.01 features of commands](0057-bdl-5-01-new-features.md).

## fglcomp -W shadow warning option

Program variables, constants and types can be defined a different scopes (global, module,
function), and can shadow symbols defined in a higher scope with the same name.

To detect symbols redefinitions in different scopes, use the `-W shadow` option of
the [`fglcomp`](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")
compiler:

```
$ cat main.4gl
GLOBALS
DEFINE trace_level INTEGER
END GLOBALS

FUNCTION trace()
    DEFINE trace_level INTEGER
    DISPLAY trace_level
END FUNCTION

$ fglcomp -W shadow main.4gl
main.4gl:6:12:6:22:warning:(-8456) The declaration of 'trace_level'
  shadows a declaration in scope global.
main.4gl:2:8:2:18:note: shadowed declaration is here.
```

## fglcomp -W keywords warning option

Most keywords can be used as variable or function identifiers without any problem. However, some
keywords will lead to problems:

- Keywords used in the grammar of expression can not be used as variable or function names.
- Keywords used as the name of primitive types can not be used as type names.

Example 1: Ambiguity:

```
DEFINE date DECIMAL(10,2) = 999.99
DISPLAY date -- displays the current date as formatted string
```

Example 2: Syntax error:

```
DEFINE date_value DATE
DEFINE day INT = 999
DISPLAY day(date_value)
DISPLAY day
DISPLAY day(date_value) -- syntax error
```

Starting with version 5.01.06, the [`fglcomp`](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.") compiler has a new "`keywords`" parameter for the
`-W` option, that checks for symbols (variables) using a language keyword name that
can lead to ambiguous semantics.

The message number is [-8455](../15_library-reference/4483-genero-bdl-errors.md).

This warning option is enabled by default, and can be disabled with the "`no`"
option prefix:

```
fglcomp -W no-keywords mysource.4gl
```

It is strongly discouraged to disable this warning.

## Changes to fgldbsch

Starting with version 5.01.00, the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool will by default deny SQL table or column names that are not
valid identifiers (like `"Customer Name"`), or when a table/column name is a
case-insensitive duplicate of other table/column names (like `"Stock"` vs
`"STOCK"`).

To bypass this control and allow invalid table or column names in the .sch
file, use the `-sl` (sloppy) option. This option has been implemented in case if an
SQL table mixes valid and invalid column names, and `DEFINE varname LIKE
tabname.colname` instructions are used with the
columns using regular identifier names.

To skip problematic tables and extract valid table defintions, use the already-existing
`-ie` (ignore errors) option.

Use the `-v` verbose option to get details of the database schema extraction
process.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases.

## Related links

**Related concepts**  

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
