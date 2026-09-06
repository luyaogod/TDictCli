---
title: "om.SaxAttributes.getValueByIndex"
source: "fgl-topics/c_fgl_ClassSaxAttributes_getValueByIndex.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.getValueByIndex"
type: "concept"
---

# om.SaxAttributes.getValueByIndex

> Returns an attribute value by position.

## Syntax

```
getValueByIndex(
   index INTEGER )
  RETURNS STRING
```

1. index is the position of the attribute in the list.

## Usage

The `getValueByIndex()` method returns the value of the
attribute at the specified ordinal position in the list.

If the attribute does not exist at the given position, the method returns
`NULL`.

## Example

```
DEFINE attrs om.SaxAttributes
...
DISPLAY attrs.getValueByIndex(3)
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.getLength](3343-om-saxattributes-getlength.md "Returns the number of attributes in the list.")

[om.SaxAttributes.getName](3344-om-saxattributes-getname.md "Returns the name of an attribute by position.")
