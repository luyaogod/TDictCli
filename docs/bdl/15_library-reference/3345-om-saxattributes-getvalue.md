---
title: "om.SaxAttributes.getValue"
source: "fgl-topics/c_fgl_ClassSaxAttributes_getValue.html"
breadcrumb: "Library reference > Built-in packages > The om package > The SaxAttributes class > om.SaxAttributes methods > om.SaxAttributes.getValue"
type: "concept"
---

# om.SaxAttributes.getValue

> Returns the value of an attribute by name.

## Syntax

```
getValue(
   name STRING )
  RETURNS STRING
```

1. name is the name of an attribute.

## Usage

The `getValue()` method
returns the value of the attribute identified by the name passed
as parameter.

If the attribute does not exist, the method returns `NULL`.

## Example

```
DEFINE attrs om.SaxAttributes
...
DISPLAY attrs.getValue("name")
```

For a complete example, see [Example 2: Creating a SAX attributes object](3351-example-2-creating-a-sax-attributes-object.md).

## Related links

**Related concepts**  

[om.SaxAttributes.getValueByIndex](3346-om-saxattributes-getvaluebyindex.md "Returns an attribute value by position.")
