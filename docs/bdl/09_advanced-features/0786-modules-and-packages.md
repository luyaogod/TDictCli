---
title: "Modules and packages"
source: "fgl-topics/c_fgl_programs_003.html"
breadcrumb: "Advanced features > Program structure > Modules and packages"
type: "concept"
---

# Modules and packages

> Programs sources can be organized in packages of modules.

The structure of a program consists of a main module (defining the [`MAIN` block](0789-the-main-block-function.md "The MAIN block is the starting point of the program.")), and several modules, that can
be grouped in a hierarchy of packages.

A package is a collection of modules, located in directories that define a package hierarchy.

For backward compatibility, Genero BDL allows flat module organization, with a program linking
process. However, new developments should use [imported
modules and packages](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module."), to define the module dependencies as in other modern programming
languages. The import solution does not require a linking phase, and allows better error checking at
compile time.

Packages of modules must be organized in a tree of directories. The modules are named by the
filename of the .4gl source, while packages are named by their directory
name.

> **Tip:**
>
> Use lower case names for modules and packages to avoid any cross-platform issues
> regarding file system case sensitivity.

A project source tree start at the top-dir directory:

```
top-dir
|-- prog1.4gl
|-- package_1
|   |-- module_11.4gl
|   |-- module_12.4gl
|   |-- module_13.4gl
|-- package_2
|   |-- module_21.4gl
|   |-- module_22.4gl
|   |-- sub_package_22
|   |   |-- module_221.4gl
|   |   |-- module_222.4gl
...
```

Each module must define the package it belongs to, by reflecting its directory path starting from
top-dir. In the above directory structure, the
module\_221.4gl must start with the following [`PACKAGE`](0816-package.md "Defines the package the module belongs to.")
instruction:

```
PACKAGE package_2.sub_package_22
...
```

Modules using program elements of a set of modules grouped in a package can import the whole
package with `IMPORT FGL package-path.*`. For example, the
prog1.4gl main module can use these
instructions:

```
IMPORT FGL package_1.*
IMPORT FGL package_2.sub_package_22.*
...
```

However, best practice is to list the required modules, because packages can be extended with
more modules:

```
IMPORT FGL package_1.module_11
IMPORT FGL package_1.module_13
IMPORT FGL package_2.sub_package_22.module_221
...
```

## Related links

**Related concepts**  

[Content of a .4gl module](0788-content-of-a-4gl-module.md "A module defines a set of program elements.")

[Importing modules](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
