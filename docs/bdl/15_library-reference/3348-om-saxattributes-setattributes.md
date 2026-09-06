---
title: "om.SaxAttributes.setAttributes"
source: "fgl-topics/c_fgl_ClassSaxAttributes_setAttributes.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.setAttributes"
type: "concept"
---

# om.SaxAttributes.setAttributes

> Clears the list and copies the attributes passed.

## Syntax

```
setAttributes(
   attr om.SaxAttributes )
```

1. attr is a reference to a list of attributes.

## Usage

The `setAttributes()` method takes an existing
`om.SaxAttributes` object reference and makes a copy
of all attributes into the current attribute list.

## Example

```
DEFINE curr, orig om.SaxAttributes
...
CALL curr.setAttributes(orig)
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.addAttribute](3339-om-saxattributes-addattribute.md "Appends a new attribute to the end of the list.")
