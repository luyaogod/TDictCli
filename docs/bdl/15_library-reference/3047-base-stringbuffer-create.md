---
title: "base.StringBuffer.create"
source: "fgl-topics/c_fgl_ClassStringBuffer_create.html"
breadcrumb: "Library reference > Built-in packages > The base package > The StringBuffer class > base.StringBuffer methods > base.StringBuffer.create"
type: "concept"
---

# base.StringBuffer.create

> Create a string buffer object.

## Syntax

```
base.StringBuffer.create()
  RETURNS base.StringBuffer
```

## Usage

Use the `base.StringBuffer.create()` class method to create
a string buffer object.

The created object must be assigned to a program variable defined with the
`base.StringBuffer` type.

## Example

```
DEFINE buf base.StringBuffer
LET buf = base.StringBuffer.create()
...
```

For a complete example, see [Example 2: Modify a StringBuffer with a function](3070-example-2-modify-a-stringbuffer-with-a-function.md).
