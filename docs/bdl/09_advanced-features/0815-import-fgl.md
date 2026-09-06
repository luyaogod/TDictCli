---
title: "IMPORT FGL"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > IMPORT FGL"
type: "concept"
---

# IMPORT FGL

> The IMPORT FGL instruction imports module symbols.

## Syntax

```
IMPORT FGL { module-name [ AS alias-name ]
           | package-path.module-name [ AS alias-name ]
           | package-path.*
           }
```

1. module-name is an [identifier](../08_language-basics/0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.") defining the module to be imported (without the file extension).
2. package-path is a dot-separated list of [identifiers](../08_language-basics/0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.") specifying a package path, to
   import a set of modules defined with the [`PACKAGE`](0816-package.md "Defines the package the module belongs to.") instruction.
3. The identifiers specified with `IMPORT FGL` are case sensitive. The directory
   names and module file names in the file system must match the character case.
4. The alias-name after the `AS` keyword defines an alias for the
   imported module. The module symbols can be prefixed with this alias name.

## Usage

With `IMPORT FGL module-name` or `IMPORT FGL
package-path.module-name`, the symbols of the named
module can be referenced in the current module. In this form, it is possible to define an alias for
the module with the `AS` keyword.

When specifying a package path with `IMPORT FGL package-path.*`
or `IMPORT FGL package-path.module-name`, the
modules must have been marked with the [`PACKAGE`](0816-package.md "Defines the package the module belongs to.") instruction.

Using `IMPORT FGL package-path.*`, is equivalent to importing
all modules that belong to this package. Avoid using `IMPORT FGL
package-path.*` when the package contains a lot of modules: The compiler
will have to process all imported modules.

Packages can be organized hierarchically in a tree of directories. However, when using
`IMPORT FGL package-path.*`, only modules that belong to this
package are imported. If package directories exist below package-path, they will
not be imported recursively.

At runtime, the imported modules are only loaded on demand, when the program flow reaches an
instruction that uses an element of the imported module. For example, when calling a function or
when assigning a (public) module variable of the imported module.

The names specified after the `IMPORT FGL` instruction are case-sensitive.

The imported module symbols that can be referenced are:

- Public [functions](../08_language-basics/0761-functions.md "Describes user defined functions.")
- Public [constants](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values.")
- Public [types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.")
- Public module [variables](../08_language-basics/0686-variables.md "Explains how to define program variables.")

## Example

Module "account.4gl":

```
PRIVATE DEFINE current_account VARCHAR(20)

PUBLIC FUNCTION set_account(id)
  DEFINE id VARCHAR(20)
  LET current_account = id 
END FUNCTION
```

Module "myutils.4gl":

```
PRIVATE DEFINE initialized BOOLEAN

PUBLIC TYPE t_prog_info RECORD
        name STRING,
        version STRING,
        author STRING
     END RECORD
  
PUBLIC FUNCTION init()
  LET initialized = TRUE
END FUNCTION
  
PUBLIC FUNCTION fini()
  LET initialized = FALSE
END FUNCTION

PUBLIC FUNCTION tokenize(s STRING,
                         a DYNAMIC ARRAY OF STRING)
  DEFINE tok base.StringTokenizer,
         x INTEGER
  LET tok = base.StringTokenizer.create(s," \t\n")
  CALL a.clear()
  LET x=0
  WHILE tok.hasMoreTokens()
      LET x=x+1
      LET a[x] = tok.nextToken()
  END WHILE
END FUNCTION
```

Module "program.4gl":

```
IMPORT FGL myutils
IMPORT FGL account
DEFINE filename STRING
DEFINE proginfo t_prog_info  -- Type is defined in myutils 
MAIN
  DEFINE arr DYNAMIC ARRAY OF STRING
  LET proginfo.name = "program"
  LET proginfo.version = "0.99"
  LET proginfo.author = "scott"
  CALL myutils.init()  -- with module prefix
  CALL set_account("CFX4559")  -- without module prefix 
  CALL tokenize("aaa bbb ccc",arr)
  DISPLAY arr[2]
END MAIN
```
