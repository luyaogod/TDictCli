---
title: "The MAIN block / function"
source: "fgl-topics/c_fgl_programs_MAIN.html"
breadcrumb: "Advanced features > Program structure > The MAIN block / function"
type: "concept"
---

# The MAIN block / function

> The MAIN block is the starting point of the program.

## Syntax 1 (MAIN / END MAIN)

```
MAIN
    [ local-declaration
        [...]
    ]
    instruction
    [...]
END MAIN
```

1. local-declaration is a `DEFINE`, `CONSTANT` or
   `TYPE` instruction.
2. instruction is a language statement.

## Syntax 2 (FUNCTION main())

```
FUNCTION main()
    [ local-declaration
        [...]
    ]
    instruction
    [...]
END FUNCTION
```

1. local-declaration is a `DEFINE`, `CONSTANT` or
   `TYPE` instruction.
2. instruction is a language statement.

## Usage

A Genero program starts in the `MAIN` block, to perform the instructions defined
in this block.

> **Important:**
>
> If a `DATABASE` instruction was specified (for the compilation
> DB schema) before the `MAIN / END MAIN` block, an implicit connection will occur in
> `MAIN`. For more details see the [`SCHEMA`](0793-schema.md "Defines the database schema files to be used for compilation.") instruction.

The `MAIN` block typically consists of:

1. The signal handling instructions [`DEFER
   INTERRUPT` and `DEFER QUIT`](0937-defer-interrupt-quit.md "The DEFER instruction defines the program behavior when interrupt or quit signals are received."),
2. The exception handling instruction [`WHENEVER ERROR
   CALL`](0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module."),
3. Global runtime configuration settings with the [`OPTIONS` instruction](0924-options-runtime.md "The OPTIONS instruction inside program blocks controls program behavior at runtime."),
4. A database connection with the [`CONNECT
   TO`](../10_sql-support/1099-connect-to.md "Opens a new database session in multi-session mode.") or [`DATABASE`](../10_sql-support/1096-database.md "Opens a new database connection in unique-session mode.")
   instruction,
5. In an interactive program, a call to a function implementing the main dialog instruction
   controlling the main form.

The `MAIN / END MAIN` block must appear before any other `FUNCTION / END
FUNCTION` block:

```
IMPORT FGL cust_module
MAIN
    DEFINE uname, upswd STRING
    DEFER INTERRUPT
    DEFER QUIT
    OPTIONS FIELD ORDER FORM, INPUT WRAP,
            SQL INTERRUPT ON, HELP FILE "myhelp"
    CALL get_login() RETURNING uname, upswd
    TRY
        CONNECT TO "stores" USER uname USING upswd
    CATCH
        IF sqlca.sqlcode < 0 THEN
           DISPLAY "Error: Could not connect to database."
           EXIT PROGRAM 1
        END IF
    END TRY
    CALL cust_module.customer_input()
END MAIN

FUNCTION show_help()
    DISPLAY "Usage: ..."
END FUNCTION
```

## FUNCTION main()

The `MAIN` block can also be defined as a regular function with `FUNCTION
main() / END FUNCTION`.

In fact a `MAIN / END MAIN` block is equivalent to `FUNCTION main() / END
FUNCTION` (returning no values), except that with a `MAIN` block, an implicit
database connection is performed, if the `DATABASE` instruction is used before
`MAIN / END MAIN`, to define the compilation database schema (the implicit database
connection does not occur, when using the `SCHEMA` instruction).

The `FUNCTION main()` can be placed after other function definitions, to make the
code mode readable:

```
TYPE t_func FUNCTION (p1 INT, p2 INT) RETURNS INT

FUNCTION add(p1 INT, p2 INT) RETURNS INT
    RETURN p1 + p2
END FUNCTION

FUNCTION sub(p1 INT, p2 INT) RETURNS INT
    RETURN p1 - p2
END FUNCTION

FUNCTION main()
    DEFINE op t_func
    LET op = FUNCTION add
    DISPLAY op( 5, 10 )
    LET op = FUNCTION sub
    DISPLAY op( 5, 10 )
END FUNCTION
```

## Defining MAIN in imported modules

When using [`IMPORT FGL`](0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.") (that
is, when not [linking programs](../13_programming-tools/2537-linking-programs.md "Describes how to link .42m modules together to build a .42r program file.")), a
`MAIN` block or `main()` function can be defined in the imported
modules.

This allows for example to write unit tests in the same source module.

File math.4gl:

```
MAIN
    DEFINE
        p1 INT = 6,
        p2 INT = 9,
        res INT

    DISPLAY "Unit testing add (", p1, " , ", p2, ")"
    LET res = add(p1, p2)
    IF res = 15 THEN
        DISPLAY " PASSED"
    ELSE
        DISPLAY " FAILED"
    END IF
END MAIN

PUBLIC FUNCTION add(p1 INT, p2 INT) RETURNS INT
    RETURN p1 + p2
END FUNCTION
```

File
main.4gl:

```
IMPORT FGL math
MAIN
    DISPLAY math.add(5, 4)
END MAIN
```

Compiling and running both
modules:

```
$ fglcomp math.4gl
$ fglrun math.42m
Unit testing add (          6 ,           9)
 PASSED
$ fglcomp main.4gl
$ fglrun main.42m
          9
```

## Related links

**Related concepts**  

[Functions](../08_language-basics/0761-functions.md "Describes user defined functions.")

[Exceptions](0848-exceptions.md "Describes exception (error) handling in the programs.")
