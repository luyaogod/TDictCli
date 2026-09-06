---
title: "PACKAGE"
source: "fgl-topics/c_fgl_programs_PACKAGE.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > PACKAGE"
type: "concept"
---

# PACKAGE

> Defines the package the module belongs to.

## Syntax

```
PACKAGE package-path
```

1. package-path identifies the package for this module. This must be a
   dot-separated list of of [identifiers](../08_language-basics/0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.")
   corresponding to the directory path where the .4gl modules sources and/or
   .42m p-code modules are located, relative to a top-dir
   identifying the root of the sources tree.
2. package-path is case senstivite. The directory names in the file system must
   match the character case.

## Usage

The `PACKAGE` instruction defines the package path for the current module.

Modules can be grouped into packages, that can then be used by other modules with the
`IMPORT FGL package-path.*` instruction. Modules assigned to a
package can also be imported individually with an `IMPORT FGL
package-path.module-name` instruction.

The names in the package-path, as well as the module-name,
are case-sensitive.

The package-path must reflect the directory path where the module is located,
starting from a top-dir root directory.

Unlike directory paths using OS specific separators, the elements of a
package-path must be separated by a dot. For example, if the
cleanup.4gl module is located in the
top-dir/myutils/sqlutils directory, the
`PACKAGE` specification in cleanup.4gl must
be:

```
PACKAGE myutils.sqlutils
```

At compile time and at runtime, the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable can be used to find package modules: A
`root-path` defined in FGLLDPATH will be used to find a package
module as
`root-path/package-path/module-name[.4gl|.42m]`.
FGLLDPATH typically contains the top-dir of the package path. When compiling or
executing the main module from top-dir, no FGLLDPATH needs to be set.

## Example

In the directory top-dir/myutils/sqlutils, the source
file cleanup.4gl defines:

```
PACKAGE myutils.sqlutils
PUBLIC FUNCTION clean_temp_data()
    DISPLAY "Cleaning up temp data..."
END FUNCTION
```

Another module called find.4gl
defines:

```
PACKAGE myutils.sqlutils
PUBLIC FUNCTION find_customer_dups()
    DISPLAY "Searching customer duplicates..."
END FUNCTION
```

The using module located in top-dir can then import the packages as
follows:

```
IMPORT FGL myutils.sqlutils.*
MAIN
    CALL myutils.sqlutils.cleanup.clean_temp_data()
    CALL myutils.sqlutils.find.find_customer_dups()
END MAIN
```

If the symbol names available from a given package are not ambigious, the calling module can omit
some parts of the the path to the symbol:

```
IMPORT FGL myutils.sqlutils.*
MAIN
    CALL clean_temp_data()
    CALL find.find_customer_dups()
END MAIN
```

Compiling from top-dir:

```
$ fglcomp --verbose main.4gl
[parsing myutils/sqlutils/find.4gl]
[building myutils/sqlutils/find]
[writing myutils/sqlutils/find.42m]
[parsing myutils/sqlutils/cleanup.4gl]
[building myutils/sqlutils/cleanup]
[writing myutils/sqlutils/cleanup.42m]
[parsing main.4gl]
[building main]
[writing main.42m]
...
```

## Related links

**Related concepts**  

[Organizing modules in packages](0818-organizing-modules-in-packages.md "Modules to be imported can be grouped in packages.")

[fglcomp](../13_programming-tools/2516-fglcomp.md "The fglcomp tool compiles .4gl source files into .42m p-code modules, and does various other tasks.")
