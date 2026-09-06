---
title: "num_args()"
source: "fgl-topics/c_fgl_BuiltInFunctions_NUM_ARGS.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > num_args()"
type: "concept"
---

# num_args()

> Returns the number of program arguments.

## Syntax

```
FUNCTION num_args()
  RETURNS INTEGER
```

## Usage

Returns the number of arguments passed to the program.

The function returns 0 if no arguments are passed to the program.

Use the [`arg_val()`](2726-arg-val.md "Returns a command line argument by position.")
function to get a command line argument at a given position.

## Related links

**Related concepts**  

[base.Application.getArgumentCount](2971-base-application-getargumentcount.md "Returns the total number of command line arguments.")
