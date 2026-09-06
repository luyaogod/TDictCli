---
title: "Workflow of a program"
source: "fgl-topics/c_fgl_programs_017.html"
breadcrumb: "Advanced features > Program structure > Workflow of a program"
type: "concept"
---

# Workflow of a program

> The Genero BDL language is a procedural programming language.

The program starts from the [`MAIN`
block](0789-the-main-block-function.md "The MAIN block is the starting point of the program.") where code can invoke other blocks of instructions defined as callable routines with
[`FUNCTION / END FUNCTION`](../08_language-basics/0761-functions.md "Describes user defined functions.") blocks.

The language statements are executed by the runtime system in the order that they appear in the
code:

```
MAIN
    CALL func1()
END MAIN

FUNCTION func1()
    DISPLAY "Hello from func1()!"
END FUNCTION
```

Some instructions can include other instructions. Such instructions are called compound
statements. Every compound statement of the language supports the `END
statement` keyword (where statement is the name of the compound
statement), to mark the end of the compound statement construct within the source code module.

Most compound statements also support the `EXIT statement` keywords, to
transfer control of execution to the statement that follows the `END
statement` keywords.

By definition, every compound statement can contain at least one statement block, a group of one
or more consecutive statements. In the syntax diagram of a compound statement, a statement block
always includes this element:

```
MAIN
    INPUT BY NAME rec.*
        ...
        ON ACTION quit
            EXIT INPUT
    END INPUT
END MAIN
```

Consider using the `IMPORT FGL` instruction to import module into the current
module, to benefit from all compiler and language features. With `IMPORT FGL`,
programs can be organized in a well structured tree of modules and
packages.

```
IMPORT FGL myutils
MAIN
    CALL myutils.start_debug_log("/tmp/myfile.log")
    ...
END MAIN
```

## Related links

**Related concepts**  

[Content of a .4gl module](0788-content-of-a-4gl-module.md "A module defines a set of program elements.")
