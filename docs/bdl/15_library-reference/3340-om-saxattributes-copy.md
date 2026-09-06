---
title: "om.SaxAttributes.copy"
source: "fgl-topics/c_fgl_ClassSaxAttributes_copy.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.copy"
type: "concept"
---

# om.SaxAttributes.copy

> Clones an existing SAX attributes object.

## Syntax

```
copy(
   attr om.SaxAttributes )
  RETURNS om.SaxAttributes
```

1. attr is a set of SAX attributes to clone.

## Usage

The `om.SaxAttributes.copy()` class method makes a clone of the
`om.SaxAttributes` object passed as reference and returns the created object.

## Example

```
DEFINE copy, orig om.SaxAttributes
...
LET copy = om.SaxAttributes.copy(orig)
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.create](3341-om-saxattributes-create.md "Create a new SAX attributes object.")
