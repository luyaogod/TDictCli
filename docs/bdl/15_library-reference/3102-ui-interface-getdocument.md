---
title: "ui.Interface.getDocument"
source: "fgl-topics/c_fgl_ClassInterface_getDocument.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getDocument"
type: "concept"
---

# ui.Interface.getDocument

> Returns the DOM document of the abstract user interface tree.

## Syntax

```
ui.Interface.getDocument()
  RETURNS  om.DomDocument
```

## Usage

The `ui.Interface.getDocument()` method returns the DOM document of the
abstract user interface tree.

Define a variable with the type `om.DomDocument` to receive the result of this
method.

Consider using the `getRootNode()` method instead of getting the root DOM node of
the AUI tree directly.

## Related links

**Related concepts**  

[The DomDocument class](3281-the-domdocument-class.md "The om.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")

[User interface basics](../11_user-interface/1508-user-interface-basics.md "This section introduces to the foundation of the Genero user interface.")

[ui.Interface.getRootNode](3107-ui-interface-getrootnode.md "Get the root DOM node of the abstract user interface.")
