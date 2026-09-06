---
title: "om.SaxAttributes.getLength"
source: "fgl-topics/c_fgl_ClassSaxAttributes_getLength.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.getLength"
type: "concept"
---

# om.SaxAttributes.getLength

> Returns the number of attributes in the list.

## Syntax

```
getLength()
  RETURNS INTEGER
```

## Usage

The `getLength()` method
returns the number of attributes in the current SAX attribute list.

Use
this method with `getName()` and `getValueByIndex()` to
retrieve attributes by position.

## Example

```
DEFINE attrs om.SaxAttributes,
       index INTEGER
...
FOR index = 1 TO attrs.getLength()
    DISPLAY attrs.getName(index), " = ", attrs.getValueByIndex(index)
END FOR
```

## Related links

**Related concepts**  

[om.SaxAttributes.getName](3344-om-saxattributes-getname.md "Returns the name of an attribute by position.")

[om.SaxAttributes.getValueByIndex](3346-om-saxattributes-getvaluebyindex.md "Returns an attribute value by position.")
