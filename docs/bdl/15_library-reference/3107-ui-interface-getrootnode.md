---
title: "ui.Interface.getRootNode"
source: "fgl-topics/c_fgl_ClassInterface_getRootNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.getRootNode"
type: "concept"
---

# ui.Interface.getRootNode

> Get the root DOM node of the abstract user interface.

## Syntax

```
ui.Interface.getRootNode()
  RETURNS om.DomNode
```

## Usage

The `ui.Interface.getRootNode()` method returns the root DOM node of the
abstract user interface tree.

Define a variable of the type `om.DomNode` to receive the result of this method.

```
DEFINE rn om.DomNode
LET rn = ui.Interface.getRootNode()
-- use d to inspect/change the AUI tree
```

## Example

```
MAIN
    DEFINE rn om.DomNode
    MENU "Test"
       COMMAND "Display AUI"
          LET rn = ui.Interface.getRootNode()
          DISPLAY rn.toString()
    END MENU
END MAIN
```

## Related links

**Related concepts**  

[The DomDocument class](3281-the-domdocument-class.md "The om.DomDocument class provides methods to manipulate a data tree, following the DOM standards.")

[User interface basics](../11_user-interface/1508-user-interface-basics.md "This section introduces to the foundation of the Genero user interface.")
