---
title: "ui.Form.findNode"
source: "fgl-topics/c_fgl_ClassForm_findNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.findNode"
type: "concept"
---

# ui.Form.findNode

> Search for a child node in the form.

## Syntax

```
findNode(
   tagName STRING,
   name STRING )
  RETURNS om.DomNode
```

1. tagName defines the type of the node.
2. name defines the name of the node.

## Usage

The `findNode()` method allows you to search for a specific DOM node in the
abstract representation of the form. You search for a child node by giving its type and the name of
the element (that is the tagName and the value of the 'name' attribute).

The method returns the first element found matching the specified type
(tagName) and node name. Form element names must be unique for
the same type of nodes, if you want to distinguish all elements.

## Example

```
MAIN
  DEFINE f ui.Form
  DEFINE n, c om.DomNode
  DEFINE rec RECORD
             custid INTEGER,
             custname VARCHAR(40)
         END RECORD
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  INPUT BY NAME rec.*
      BEFORE INPUT
         LET f = DIALOG.getForm()
         LET n = f.findNode("FormField", "formonly.custname")
         LET c = n.getFirstChild()
         DISPLAY c.getAttribute("shift")
  END INPUT
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[ui.Dialog.getForm](3200-ui-dialog-getform.md "Returns the current form used by the dialog.")

[The abstract user interface tree](../11_user-interface/1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")
