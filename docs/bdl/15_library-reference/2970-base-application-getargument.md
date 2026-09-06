---
title: "base.Application.getArgument"
source: "fgl-topics/c_fgl_ClassApplication_getArgument.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods > base.Application.getArgument"
type: "concept"
---

# base.Application.getArgument

> Returns the command line argument by position.

## Syntax

```
base.Application.getArgument(
   index INTEGER )
  RETURNS STRING
```

1. index is the index of the program argument.

## Usage

The index is the program argument
position. The first program argument is identified by the position
1. Argument number zero is the program name.

Returns `NULL` if
there is no argument provided at the position.

## Example

```
MAIN
  DEFINE i INTEGER
  FOR i=1 TO base.Application.getArgumentCount()
    DISPLAY base.Application.getArgument(i)
  END FOR
END MAIN
```
