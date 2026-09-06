---
title: "Content of a .4gl module"
source: "fgl-topics/c_fgl_programs_019.html"
breadcrumb: "Advanced features > Program structure > Content of a .4gl module"
type: "concept"
---

# Content of a .4gl module

> A module defines a set of program elements.

## Syntax

The declaration order of elements defined in a program module must follow this
syntax:

```
[ compiler-options
| package-statement
| import-statement [...]
| schema-statement
| globals-inclusion
]

[ constant-definition
| type-definition
| variable-definition
    [...] ]
]

[ MAIN-block ]

[ constant-definition
| type-definition
| variable-definition
| dialog-block
| function-block
| report-routine
    [...] ]
]
```

1. compiler-options are described in [OPTIONS (Compilation)](0922-options-compilation.md "OPTIONS outside program blocks defines semantics of the language for the compiler.").
2. package-statement defines the package this module belongs to, see [PACKAGE](0816-package.md "Defines the package the module belongs to.").
3. import-statement imports an external module, see [Importing modules](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").
4. schema-statement defines a [database schema](0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") for the
   compilation.
5. globals-inclusion includes a [globals file](0913-globals.md "Global variables can be shared among all modules of a program.").
6. MAIN-block declares the [main block of
   the program](0789-the-main-block-function.md "The MAIN block is the starting point of the program.").
7. constant-definition defines [constants](../08_language-basics/0710-constants.md "The definition of constants allows to centralize common static values.") (module-wise).
8. type-definition defines [user types](../08_language-basics/0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") (module-wise).
9. variable-definition defines [variables](../08_language-basics/0686-variables.md "Explains how to define program variables.") (module-wise).
10. dialog-block declares a [declarative dialog](../11_user-interface/2152-syntax-of-the-declarative-dialog-block.md "The declarative DIALOG block defines an interactive instruction that can be used by a parent DIALOG using the SUBDIALOG clause.").
11. function-block declares a [function](../08_language-basics/0761-functions.md "Describes user defined functions.").
12. report-routine declares a [report routine](../12_reports/2474-the-report-routine.md "The report routine implements the body of a report, with formatting instructions.").

After the `MAIN` block, or when no `MAIN` block exists in the
module, the constants, types, variables, dialog-blocks, functions, methods and report routines can
be specified in any order.

## Usage

A module defines a set of program elements that can be used by other modules when defined as
`PUBLIC`, or to be local to the current module when defined as `PRIVATE`.

Program elements are user-defined types, variables, constants, functions, report routines, and
declarative dialogs.

A module can import other modules with the [`IMPORT
FGL`](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") instruction. The imported program elements can be used in the current module.

Program modules are written as .4gl source files and are compiled to
.42m files. Compiled modules (.42m files) can be linked together
to create a program. However, linking is supported for backward compatibility only. The preferred way is
to define module dependencies with the `IMPORT FGL` instruction. For better code re-usability,
module elements can be shared by each other by qualifying module variables, constants, types, and functions
with `PRIVATE` or `PUBLIC` keywords. `PUBLIC` module elements
can be referenced in other modules.

Modules can be organized in packages.

## Example

```
OPTIONS SHORT CIRCUIT 
IMPORT FGL cust_data
SCHEMA stores

PRIVATE CONSTANT c_title = "Customer data form"

MAIN
    ...
END MAIN

PUBLIC TYPE t_cust RECORD LIKE customer.*

PRIVATE DEFINE cust_rec t_cust

DIALOG cust_input_dialog()
    INPUT BY NAME cust_rec.*
       ...
    END INPUT
END DIALOG

PRIVATE DEFINE cust_arr DYNAMIC ARRAY OF t_cust

FUNCTION do_cust_display_array()
    DISPLAY ARRAY cust_arr TO sr.* ...
        ...
    END DISPLAY
END FUNCTION

REPORT cust_report(row t_cust)
   ...
END REPORT
```

## Related links

**Related concepts**  

[Workflow of a program](0787-workflow-of-a-program.md "The Genero BDL language is a procedural programming language.")
