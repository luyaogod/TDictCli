---
title: "om.DomNode.getTagName"
source: "fgl-topics/c_fgl_ClassDomNode_getTagName.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomNode class > om.DomNode methods > om.DomNode.getTagName"
type: "concept"
---

# om.DomNode.getTagName

> Returns the XML tag name of a DOM node.

## Syntax

```
getTagName()
  RETURNS STRING
```

## Usage

The `getTagName()` method returns the XML tag name of the node.

Use this method to identity the type of the node.

For character nodes (created for example with the [`createChars()`](3288-om-domdocument-createchars.md "Create a new text node in the DOM document.") of a
`DomDocument` object), the `getTagName()` method returns
`"@chars"`.

## Example

```
DEFINE node om.DomNode
...
DISPLAY node.getTagName()
```
