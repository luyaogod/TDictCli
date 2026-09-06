---
title: "ui.Form.setElementStyle"
source: "fgl-topics/c_fgl_ClassForm_setElementStyle.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setElementStyle"
type: "concept"
---

# ui.Form.setElementStyle

> Change the style of form elements.

## Syntax

```
setElementStyle(
   name STRING,
   style STRING )
```

1. name defines the name of the node, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).
2. style is the set of style names to be used. Several style
   names can be combined, by using a space as separator.

## Usage

Change the [style](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.") of a form element with
the `setElementStyle()` method.

Pass the identifier of the form element. The identifier is the element name as
defined in the form definition. All elements with this name will be affected. If you want to
distinguish form elements, use unique names in the form definition file.

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
CALL DIALOG.getForm().setElementStyle("group1","important")
```

In the above example, `"important"` refers to a style name defined in your [.4st](../11_user-interface/1618-loading-presentation-styles.md "Presentation styles are defined in an XML file with a 4st extension. In order to load the presentation styles, the runtime system needs to locate the appropriate style file.") file.

## Related links

**Related concepts**  

[ui.Form.setElementImage](3157-ui-form-setelementimage.md "Change the image of form elements.")

[ui.Form.setElementText](3159-ui-form-setelementtext.md "Change the text of form elements.")

[ui.Form.setFieldStyle](3162-ui-form-setfieldstyle.md "Change the style of a form field.")
