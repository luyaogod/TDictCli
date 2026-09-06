---
title: "om.SaxAttributes.create"
source: "fgl-topics/c_fgl_ClassSaxAttributes_create.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.create"
type: "concept"
---

# om.SaxAttributes.create

> Create a new SAX attributes object.

## Syntax

```
create()
  RETURNS om.SaxAttributes
```

## Usage

The `om.SaxAttributes.create()` class method creates a new
`om.SaxAttributes` object and returns it.

To hold the reference to a SAX attributes object, define a variable with the type
`om.SaxAttributes` type.

## Example

```
DEFINE attrs om.SaxAttributes
...
LET attrs = om.SaxAttributes.create()
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.copy](3340-om-saxattributes-copy.md "Clones an existing SAX attributes object.")
