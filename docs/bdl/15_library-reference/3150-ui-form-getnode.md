---
title: "ui.Form.getNode"
source: "fgl-topics/c_fgl_ClassForm_getNode.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.getNode"
type: "concept"
---

# ui.Form.getNode

> Get the DOM node of the form.

## Syntax

```
getNode()
  RETURNS om.DomNode
```

## Usage

The `getNode()` method returns the DOM node containing the abstract
representation of the window/form.

After loading and displaying a form with `OPEN FORM / DISPLAY FORM` or with
`OPEN WINDOW ... WITH FORM`, get the form object for example with [`ui.Dialog.getForm()`](3200-ui-dialog-getform.md "Returns the current form used by the dialog.") and use
the `getNode()` method to query the DOM node corresponding to the form.

## Example

```
MAIN
  DEFINE f ui.Form
  DEFINE n om.DomNode
  DEFINE rec RECORD
             custid INTEGER,
             custname VARCHAR(40)
         END RECORD
  OPEN FORM f1 FROM "customer"
  DISPLAY FORM f1
  INPUT BY NAME rec.*
      BEFORE INPUT
         LET f = DIALOG.getForm()
         LET n = f.getNode()
         DISPLAY n.toString()
  END INPUT
END MAIN
```

## Related links

**Related concepts**  

[The DomNode class](3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")

[OPEN WINDOW](../11_user-interface/1572-open-window.md "Creates and displays a new window.")

[OPEN FORM](../11_user-interface/1578-open-form.md "Declares a compiled form in the program.")

[DISPLAY FORM](../11_user-interface/1579-display-form.md "Displays and associates a form with the current window.")
