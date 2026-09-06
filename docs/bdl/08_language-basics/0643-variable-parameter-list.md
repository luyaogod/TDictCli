---
title: "Variable parameter list: [ ]"
source: "fgl-topics/c_fgl_operators_VAR_PARAM_LIST.html"
breadcrumb: "Language basics > Operators > List of expression elements > Associative syntax operators > Variable parameter list: [ ]"
type: "concept"
---

# Variable parameter list: [ ]

> Variable parameter list delimiters.

## Syntax

> **Note:**
>
> In the next(s) syntax diagram(s), the `[ ] { } |`
> symbols are part of the syntax.

```
[ var1, var2, ... ]
```

1. var1, var2, … are program [variables](0686-variables.md "Explains how to define program variables.").

## Usage

The square brackets notation in function parameters defines a variable list of arguments for
a built-in function or a built-in class method.

The elements of a variable parameter list are program variables which are passed by reference.
As result, the called function can modify the content of the passed variables, to return values
in output parameters.

It is not possible to define user functions with variable parameter lists.

For real usage examples, see the read and write methods of the `base.Channel`
class.

## Example

```
MAIN
  DEFINE id INTEGER, name STRING,
           count INTEGER, stat INTEGER
  LET id = 12345
  LET name = "Forman"
  -- Warning: This is a fake call, the function does not exist!
  -- Here, id and name are passed as input values, while count
  -- and stat are used as output parameters... 
  CALL built_in_function( [id,name], [count, stat] )
END MAIN
```

## Related links

**Related concepts**  

[base.Channel.read](../15_library-reference/2993-base-channel-read.md "Reads a list of data delimited by a separator from the channel.")

[base.Channel.write](../15_library-reference/2997-base-channel-write.md "Writes a list of data delimited by a separator to the channel.")

[ui.Interface.frontCall](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")
