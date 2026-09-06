---
title: "om.DomNode.getAttributeInteger"
source: "fgl-topics/c_fgl_ClassDomNode_getAttributeInteger.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getAttributeInteger"
type: "concept"
---

# om.DomNode.getAttributeInteger

> Returns the value of a DOM node attribute, with default integer value.

## Syntax

```
getAttributeInteger(
   name STRING,
   defaultValue STRING )
  RETURNS INTEGER
```

1. name is the name of the attribute.
2. defaultValue is the default value.

## Usage

The `getAttributeInteger()` method returns the value of the attribute
passed as parameter, as defined in the current node.

DOM node attribute names are case-sensitive.

If the attribute is not defined, the method returns the default value passed as
second parameter.

## Related links

**Related concepts**  

[om.DomNode.getAttribute](3302-om-domnode-getattribute.md "Returns the value of a DOM node attribute.")

[om.DomNode.setAttribute](3322-om-domnode-setattribute.md "Sets the value of a DOM node attribute.")
