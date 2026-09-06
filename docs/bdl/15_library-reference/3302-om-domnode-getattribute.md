---
title: "om.DomNode.getAttribute"
source: "fgl-topics/c_fgl_ClassDomNode_getAttribute.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getAttribute"
type: "concept"
---

# om.DomNode.getAttribute

> Returns the value of a DOM node attribute.

## Syntax

```
getAttribute(
   name STRING )
  RETURNS STRING
```

1. name is the name of the attribute.

## Usage

The `getAttribute()` method
returns the value of the attribute passed as parameter, as defined
in the current node.

DOM node attribute names are case-sensitive.

If
the attribute does not exist for this node type, or if the attribute
is not set, the method returns `NULL`.

For character nodes (created for example with the [`createChars()`](3288-om-domdocument-createchars.md "Create a new text node in the DOM document.") of a
`DomDocument` object), you can get the text value by passing the
`@chars` attribute name to the `getAttribute()` method.

## Example

```
DEFINE node om.DomNode
DEFINE text_node om.DomNode
...
DISPLAY node.getAttribute("color")
DISPLAY text_node.getAttribute("@chars")
```

## Related links

**Related concepts**  

[om.DomNode.getAttributeString](3305-om-domnode-getattributestring.md "Returns the value of a DOM node attribute, with default string value.")

[om.DomNode.getAttributeInteger](3303-om-domnode-getattributeinteger.md "Returns the value of a DOM node attribute, with default integer value.")

[om.DomNode.setAttribute](3322-om-domnode-setattribute.md "Sets the value of a DOM node attribute.")

[om.DomDocument.createChars](3288-om-domdocument-createchars.md "Create a new text node in the DOM document.")
