---
title: "fgl_getenv()"
source: "fgl-topics/c_fgl_BuiltInFunctions_FGL_GETENV.html"
breadcrumb: "Library reference > Built-in functions > Built-in functions > fgl_getenv()"
type: "concept"
---

# fgl_getenv()

> Returns the value of the environment variable.

## Syntax

```
FUNCTION fgl_getenv(
   name STRING )
  RETURNS STRING
```

1. name is the name of the environment variable.

## Usage

The argument of `fgl_getenv()` must be the name of an
environment variable.

If the requested value exists in the current user environment, the
function returns the value of that variable. If the specified environment
variable is not defined, the function returns a
[`NULL`](../08_language-basics/0572-null.md "The NULL constant defines a non-value.") value.
If the environment variable is defined but does not have a value
assigned to it, the function returns blank spaces.

## Related links

**Related concepts**  

[fgl\_setenv()](2781-fgl-setenv.md "Sets the value of an environment variable.")
