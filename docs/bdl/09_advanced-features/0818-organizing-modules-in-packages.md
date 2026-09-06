---
title: "Organizing modules in packages"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_packages.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Organizing modules in packages"
type: "concept"
---

# Organizing modules in packages

> Modules to be imported can be grouped in packages.

## Package basics

Structuring your sources is important for maintenance, readability, reusability and
deployment.

Modules can be grouped by packages, based on the directory structure of source files. A source
tree starts at a top-dir, with package sub-directories containing
modules:

```
top-dir/
|-- mylibs/                   -- package: mylibs
|   |-- myutils/              -- package: mylibs.myutils
|   |   |-- fileutils.4gl
|   |   |-- sqlutils.4gl
|       ...
|-- webserv/                  -- package: webserv
|   |-- stock.4gl
|   |-- shipment.4gl
|       ...
...
|-- starter.4gl               -- main programs
|-- reports.4gl
```

A sequence of directories defines a package-path. In a
package-path, package names are separated by a dot. In the above structure, the
package-path "`mylibs.myutils`" corresponds to the directory path
"mylibs/myutils", relative to "top-dir".

Use simple regular identifiers for package and module names, using ASCII characters. Do not use
names such as "1mod", "forêt",
"module-1.4gl".

## Declaring a module for a package

In order to indicate that a module belongs to a package, use the [`PACKAGE`](0816-package.md "Defines the package the module belongs to.") instruction at the top of the
module source.

In the following example, the module fileutils.4gl is declared as a module
of the `mylibs.myutils` package:

```
PACKAGE mylibs.myutils

PUBLIC FUNCTION getSubDirs(path STRING) RETURNS ...
    ...
END FUNCTION
```

## Importing package modules

Importing a module from a package is achieved by specifying the package-path
in the `IMPORT FGL` instruction, followed by a dot and the module
name:

```
IMPORT FGL mylibs.myutils.fileutils

PUBLIC FUNCTION file_browser(defpath STRING) RETURNS (INTEGER,STRING)
    DEFINE dirs DYNAMIC ARRAY OF STRING
    CALL mylibs.myutils.fileutils.getSubDirs(defpath) RETURNING dirs
    ...
END FUNCTION
```

The package name specified after `IMPORT FGL` is a
package-path, that must reflect the directory path based on the
top-dir of the sources, using dots as separators. For example, to import the
"`fileutils`" module from the
top-dir/mylibs/myutils directory,
use:

```
IMPORT FGL mylibs.myutils.fileutils
```

All modules of a package can be imported by using the `.*`
syntax:

```
IMPORT FGL mylibs.myutils.*
```

When a module under a package wants to import a sibling module, specify the full
package-path in IMPORT FGL:

This is fileutils.4gl, from the `mylibs.myutils` package, in
top-dir/mylibs/myutils:

```
PACKAGE mylibs.myutils

PUBLIC FUNCTION getSubDirs(path STRING) RETURNS ...
    ...
END FUNCTION
```

This is sqlutils.4gl, from the same `mylibs.myutils`
package/directory:

```
PACKAGE mylibs.myutils

IMPORT FGL mylibs.myutils.fileutils

PUBLIC FUNCTION findConfigFile(path STRING) RETURNS STRING
    ...
    CALL mylibs.myutils.fileutils.getSubDirs("/opt/myapp") RETURNING ...
    ...
END FUNCTION
```

## Module aliases

To simplify the source code when a package-path is long, define module aliases
with the `AS` keyword:

```
IMPORT FGL mylibs.myutils.fileutils AS fu
...
    CALL fu.getSubDirs(defpath) RETURNING dirs
```

Note that the alias provided with `IMPORT FGL … AS …` is a module alias. There is
no way to define a package-path alias.

## Finding modules with FGLLDPATH

For compilation and at runtime, a `root-path` defined in the
[FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") environment variable will be used to
find a package module as
`root-path/package-path/module-name[.4gl|.42m]`.

FGLLDPATH typically contains the top-dir of the package path.

For example, with the following package/directory structure:

```
top-dir/
|-- mylibs/
|   |-- myutils/
|   |   |-- fileutils.4gl
|   |   |-- sqlutils.4gl
|-- app1/
|   |-- main.4gl
```

FGLLDPATH can be defined to top-dir, to find modules from the
mylibs.myutils package, when compiling main.4gl in the
app1 directory:

```
$ cd top-dir
$ export FGLLDPATH=$PWD
$ cd app1
$ fglcomp -M --verbose main.4gl
[parsing top-dir/mylibs/myutils/fileutils.4gl]
[building mylibs/myutils/fileutils]
[writing top-dir/mylibs/myutils/fileutils.42m]
[parsing top-dir/mylibs/myutils/sqlutils.4gl]
[building mylibs/myutils/sqlutils]
[writing top-dir/mylibs/myutils/sqlutils.42m]
[parsing main.4gl]
[building main]
[writing main.42m]
```

## Deploying modules with package directory structure

When using packages, the directory structure of the source files can be cloned into the
production environment, by excluding the .4gl (and .per)
source files, except if the .4gl sources are required for debugging
purpose.

For example, with the following directory structure and
sources:

```
top-dir-of-sources/
|-- mylibs/
|   |-- myutils/
|   |   |-- fileutils.4gl
|   |   |-- sqlutils.4gl
|-- app1/
|   |-- main.4gl
```

With main.4gl (note that this main.4gl belongs to the
`app1` package):

```
PACKAGE app1

IMPORT FGL mylibs.myutils.sqlutils

MAIN
    DISPLAY mylibs.myutils.sqlutils.findConfigFile("/tmp")
END MAIN
```

With all forms, programs and modules
compiled:

```
$ cd top-dir-of-sources
$ tar cvzf /tmp/dep-arch.tgz --exclude "*.4gl" --exclude "*.per" .
./
./app1/
./app1/main.42m
./mylibs/
./mylibs/myutils/
./mylibs/myutils/fileutils.42m
./mylibs/myutils/sqlutils.42m

$ mkdir top-dir-of-prod-env
$ cd top-dir-of-prod-env
$ $ tar xvzf /tmp/dep-arch.tgz
./
./app1/
./app1/main.42m
./mylibs/
./mylibs/myutils/
./mylibs/myutils/fileutils.42m
./mylibs/myutils/sqlutils.42m
```

Alternatively, at compile time, you can use the `--output-dir` option of
fglcomp to deploy the .42m files directly into the
production directory used at runtime:

```
$ cd top-dir-of-sources
$ fglcomp --output-dir top-dir-of-prod-env --verbose -M app1/main.4gl
```

In order to find modules of packages at runtime with fglrun, define the
FGLLDPATH environment variable to the top-dir path of the distribution
directory:

```
$ export FGLLDPATH=top-dir-of-prod-env
$ fglrun top-dir-of-prod-env/app1/main.42m
```

## Related links

**Related concepts**  

[PACKAGE](0816-package.md "Defines the package the module belongs to.")

[IMPORT FGL](0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.")
