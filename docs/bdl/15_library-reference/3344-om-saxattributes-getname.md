---
title: "om.SaxAttributes.getName"
source: "fgl-topics/c_fgl_ClassSaxAttributes_getName.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.getName"
type: "concept"
---

# om.SaxAttributes.getName

> Returns the name of an attribute by position.

## Syntax

```
getName(
   index INTEGER )
  RETURNS STRING
```

1. index is the position of the attribute in the list.

## Usage

The `getName()` method returns the name of the attribute at
the specified ordinal position in the list.

If the attribute does not exist at the given position, the method returns
`NULL`.

## Example

```
DEFINE attrs om.SaxAttributes
...
DISPLAY attrs.getName(3)
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.getLength](3343-om-saxattributes-getlength.md "Returns the number of attributes in the list.")

[om.SaxAttributes.getValueByIndex](3346-om-saxattributes-getvaluebyindex.md "Returns an attribute value by position.")
