---
title: "Declaration context"
source: "fgl-topics/c_fgl_variables_004.html"
breadcrumb: "Language basics > Variables > Declaration context"
type: "concept"
---

# Declaration context

> A variable can be declared in different contexts, which defines its visibility.

In a module, the context of a variable declaration determines where this variable can be
referenced (by other modules and language instructions), and when the memory is allocated for the
variable.

Variables can be defined as:

- global variables (`DEFINE` in `GLOBALS` block)
- module variables (`DEFINE` at module level)
- local variables (`DEFINE` or `VAR` in
  `FUNCTION` block)

Variable declaration contexts:

1. Withing a `FUNCTION`, the [`VAR`](0689-var.md "The VAR instruction declares a program variable within a code block.") instruction defines a local variable for the code block
   where it is declared. Unlike `DEFINE`, the `VAR` instruction can be
   used anywhere in the function body. Memory for `VAR` variables is allocated on the
   runtime stack, when the instruction is reached.
2. At the top of the body of a `FUNCTION`, `MAIN`, or
   `REPORT` block, `DEFINE` declares
   local variables, and causes memory to be allocated on the runtime stack when the
   function is called. These `DEFINE` declarations of local variables must precede any
   procedural statements within the same program block. The scope of reference of a local variable is
   restricted to the same program block. The variable is not visible elsewhere. Functions can be called
   recursively, and each recursive entry creates its own set of local variables. The variable is unique
   to that invocation of its program block. Each time the block is entered, a new copy of the variable
   is created.
3. Outside any `FUNCTION`, `REPORT`, or `MAIN` program
   block, `DEFINE` declares module
   variables. Module variables have a persistent state during program execution. Memory for
   module variables is allocated when the module is loaded. Module variable declarations
   (`DEFINE`) must appear before any program blocks. By default, the scope of reference
   is the whole module (module variables are private to the module), but it can be extended to the
   whole program when the variable is declared with the `PUBLIC` qualifier.
4. Inside a `GLOBALS` block, `DEFINE` declares global variables that are visible to the whole program. Global
   variables have a persistent state during program execution. Memory for global variables is allocated
   when the program starts. Multiple `GLOBALS` blocks can be defined for a given module.
   Use one module to declare all global variables and reference that module within other modules by
   using the `GLOBALS "filename.4gl"` statement as the first statement in the
   module, outside any program block.

A compile-time error occurs if you declare the same name for two variables that have the same
scope. You can, however, declare the same name for variables that differ in their scope. For
example, you can use the same identifier to reference different local variables in different program
blocks.

You can also declare the same name for two or more variables whose
scopes of reference are different but overlapping. Within their intersection,
the compiler interprets the identifier as referencing the variable
whose scope is smaller, and therefore the variable whose scope is
a superset of the other is not visible.

If a local (function) variable has the same name as a module or global variable, then the local
variable takes precedence inside the program block in which it is declared. Elsewhere in the
program, the identifier references the module or global
variable:

```
DEFINE myvar INT -- Module variable
  
MAIN
    LET myvar = 0 -- Module variable
    DISPLAY "main:  myvar=", myvar
    CALL func()
    DISPLAY "main:  myvar=", myvar
END MAIN

FUNCTION func()
    DEFINE myvar INT -- Local variable
    LET myvar = 1
    DISPLAY "func:  myvar=", myvar
END FUNCTION
```

A module variable can have the same name as a global variable. In such case, the module variable
takes precedence:

```
GLOBALS
    DEFINE myvar TEXT -- Hidden by module variable
END GLOBALS

DEFINE myvar INT -- Module variable

MAIN
    LET myvar = 0 -- Module variable
    DISPLAY "main:  myvar=", myvar
END MAIN
```

If a variable needs to be persistent during program execution and must be accessed by
different modules, instead of using global variables, define that variable in the module it belongs
to, by specifying the `PUBLIC` modifier.

## Related links

**Related concepts**  

[Example 1: Local function variables](0704-example-1-local-function-variables.md "Example 1: Local function variables")

[Example 2: PRIVATE module variables](0705-example-2-private-module-variables.md "Example 2: PRIVATE module variables")

[Example 3: PUBLIC module variables](0706-example-3-public-module-variables.md "Example 3: PUBLIC module variables")

[Example 4: Global variables](0707-example-4-global-variables.md "Example 4: Global variables")
