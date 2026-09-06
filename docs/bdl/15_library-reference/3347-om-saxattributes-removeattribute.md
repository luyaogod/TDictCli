---
title: "om.SaxAttributes.removeAttribute"
source: "fgl-topics/c_fgl_ClassSaxAttributes_removeAttribute.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.removeAttribute"
type: "concept"
---

# om.SaxAttributes.removeAttribute

> Delete an attribute by position.

## Syntax

```
removeAttribute(
   index INTEGER )
```

1. index is the position of the attribute in the list.

## Usage

The `removeAttribute()` method removes the attribute at the
given ordinal position.

If the attribute does not exist at the given position, the method returns
silently.

## Example

```
DEFINE attrs om.SaxAttributes
...
CALL attrs.removeAttribute( attrs.getLength() )
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.addAttribute](3339-om-saxattributes-addattribute.md "Appends a new attribute to the end of the list.")
