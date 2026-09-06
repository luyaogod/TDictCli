---
title: "Example 1: Multiple GLOBALS file"
source: "fgl-topics/c_fgl_Globals_008.html"
breadcrumb: "Advanced features > Globals > Examples > Example 1: Multiple GLOBALS file"
type: "concept"
description: "Important: Using GLOBALS is strongly discouraged. This feature is supported for backward compatibility, to compile legacy source code. Instead of GLOBALS , use PUBLIC symbols in modules to be imported ..."
---

# Example 1: Multiple GLOBALS file

> **Important:**
>
> Using `GLOBALS` is strongly discouraged. This feature is supported for backward
> compatibility, to compile legacy source code. Instead of `GLOBALS`, use
> `PUBLIC` symbols in modules to be imported with [`IMPORT FGL`](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.").
>
> Why not globals? Global variables visible and modifiable across the entire program code is an old
> concept that reduces modularity and flexibility. `GLOBALS` implements specific
> semantics that restrain source code organization; for example, it is allowed to specify a file path
> such as `GLOBALS "../../common/myglobals.4gl"`. The globals file can define functions
> after the `GLOBALS` block that are ignored when including such file.
>
> Conversely, the `IMPORT FGL` instruction brings more flexibility and allows a
> better organization of your source code, by grouping public variables in modules where other symbols
> (functions, types, contants) belong to the same domain.

Module labels.4gl: This module defines the text that is displayed on the
screen:

```
GLOBALS
  CONSTANT g_lbl_val = "Index:"
  CONSTANT g_lbl_idx = "Value:"
END GLOBALS
```

Module globals.4gl: Declares a global array and a constant containing its
size:

```
GLOBALS "labels.4gl" -- this statement could be line 2 of main.4gl

GLOBALS
  DEFINE g_idx ARRAY[100] OF CHAR(10)
  CONSTANT g_idxsize = 100
END GLOBALS
```

Module database.4gl: This module is dedicated to database
access:

```
GLOBALS "globals.4gl"

FUNCTION get_id()
  DEFINE li INTEGER
  FOR li = 1 TO g_idxsize -- this could be a FOREACH statement 
      LET g_idx[li] = g_idxsize - li 
  END FOR
END FUNCTION
```

Module main.4gl: Fill in the global array and display the
result:

```
GLOBALS "globals.4gl"

MAIN
  DISPLAY "Initializing constant values for this application..."
  DISPLAY "Filling the data from function get_idx in module database.4gl..."
  CALL get_id()
  DISPLAY "Retrieving a few values from g_idx"
  CALL display_data()
END MAIN

FUNCTION display_data()
  DEFINE li INTEGER
  LET li = 1
  WHILE li <= 10 AND li <= g_idxsize 
      DISPLAY g_lbl_idx CLIPPED || li || " " || g_lbl_val CLIPPED || g_idx[li]
      LET li = li + 1
  END WHILE
END FUNCTION
```
