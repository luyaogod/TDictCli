---
title: "om.SaxAttributes.addAttribute"
source: "fgl-topics/c_fgl_ClassSaxAttributes_addAttribute.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.addAttribute"
type: "concept"
---

# om.SaxAttributes.addAttribute

> Appends a new attribute to the end of the list.

## Syntax

```
addAttribute(
   name STRING,
   value STRING )
```

1. name is the name of the attribute.
2. value is the value of the attribute.

## Usage

The `addAttribute()` method appends a new attribute with name and value
at the end of the list.

Attribute names are case-sensitive.

Make sure that the strings passed to the method do not contain illegal XML
characters. Illegal XML characters will be silently ignored. Illegal XML characters are any
character below space (ASCII 32), except \r (ASCII 13), \n (ASCII 10) and \t (ASCII 9).

## Example

```
DEFINE attrs om.SaxAttributes
...
CALL attrs.addAttribute("name","jo")
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.setAttributes](3348-om-saxattributes-setattributes.md "Clears the list and copies the attributes passed.")
