---
title: "om.DomDocument.removeElement"
source: "fgl-topics/c_fgl_ClassDomDocument_removeElement.html"
breadcrumb: "Library reference > Built-in packages > The om package > The DomDocument class > om.DomDocument methods > om.DomDocument.removeElement"
type: "concept"
---

# om.DomDocument.removeElement

> Remove a DomNode object and all its descendants.

## Syntax

```
removeElement(
   oldChild om.DomNode )
```

1. oldChild is the DOM node to be removed.

## Usage

Use the `removeElement()` method to remove an element and all its
descendants from DOM document.

Any reference to the removed `om.DomNode` objects becomes invalid.
