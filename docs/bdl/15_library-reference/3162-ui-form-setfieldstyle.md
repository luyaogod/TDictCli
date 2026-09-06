---
title: "ui.Form.setFieldStyle"
source: "fgl-topics/c_fgl_ClassForm_setFieldStyle.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setFieldStyle"
type: "concept"
---

# ui.Form.setFieldStyle

> Change the style of a form field.

## Syntax

```
setFieldStyle(
   name STRING,
   style STRING )
```

1. name defines the name of the form field, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).
2. style is the set of style names to be used. Several style
   names can be combined, by using a space as separator.

## Usage

Change the [style](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.") of a form field with the
`setFieldStyle()` method.

Pass the identifier of the form field, as defined in the form definition. The
form field is identified by column name, with an optional prefix (`table.column` or
`column`).

Depending on the attributes used in the style definition
(.4st), changing the style of an element back and forth at runtime can lead to
performance issues on the front-end side, when the style attribute results in complex rendering
changes. As a general rule, consider modifying the style attribute once for the current form, before
executing an interactive instruction (dialog). During a dialog, only change the style of leaf
elements in the layout structure, such as labels, or small containers such as groups. Some style
changes may not have any effect at all, when the GUI framework does not allow modifying the
decoration once the widget is created.

## Example

```
CALL DIALOG.getForm().setFieldStyle("customer.cust_name","mandatory")
```

In the above example, `"mandatory"` refers to a style name defined in your [.4st](../11_user-interface/1618-loading-presentation-styles.md "Presentation styles are defined in an XML file with a 4st extension. In order to load the presentation styles, the runtime system needs to locate the appropriate style file.") file.

## Related links

**Related concepts**  

[ui.Form.setFieldHidden](3161-ui-form-setfieldhidden.md "Show or hide a form field.")

[ui.Form.setElementStyle](3158-ui-form-setelementstyle.md "Change the style of form elements.")
