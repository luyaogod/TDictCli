---
title: "Using a style"
source: "fgl-topics/c_fgl_presentation_styles_007.html"
breadcrumb: "User interface > Form definitions > Presentation styles > Using presentation styles > Using a style"
type: "concept"
---

# Using a style

> To apply a specific style, set the style-name in the style attribute of the node representing the graphical element in the abstract user interface tree.

There are different ways to set the `style` attribute of an element:

- As a form element attribute, with a [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute in the form specification file.
- In the `ATTRIBUTES` clause of instructions such as [`OPEN WINDOW`](1572-open-window.md "Creates and displays a new window."), [`MESSAGE`](1885-message.md "The MESSAGE instruction displays a message to the user."), [`ERROR`](1886-error.md "The ERROR instruction displays an error message to the user.").
- Dynamically by program code, using the [ui.Form.setElementStyle()](../15_library-reference/3158-ui-form-setelementstyle.md "Change the style of form elements.") method.

The names used to define the `STYLE` attribute must be a
style-name only, these must not contain the element-type
that is typically used to define the style in a .4st file (as
`CheckBox.important` for example)

For example, to define a style in a form file for an input
field:

```
EDIT f001 = customer.fname, STYLE = "info";
```

Several styles [can be combined](1615-combining-styles-and-style-attributes.md "Several styles can be combined, and style attributes can be mixed."), by using a
space as separator:

```
EDIT f001 = customer.fname, STYLE = "info important";
```

## Related links

**Related concepts**  

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
