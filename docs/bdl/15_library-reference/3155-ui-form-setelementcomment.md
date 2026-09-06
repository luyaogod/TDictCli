---
title: "ui.Form.setElementComment"
source: "fgl-topics/c_fgl_ClassForm_setElementComment.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setElementComment"
type: "concept"
---

# ui.Form.setElementComment

> Set the comment/hint of form elements.

## Syntax

```
setElementComment(
   name STRING,
   comment STRING )
```

1. name defines the name of the node, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).
2. comment the text of the element comment.

## Usage

Change the comment attribute of a form element with the `setElementComment()` method.

The comment is used to display a hint text, for example with a tooltip when mouse hovers a
toolbar item.

Pass the identifier of the form element. The identifier is the element name as
defined in the form definition. All elements with this name will be affected. If you want to
distinguish form elements, use unique names in the form definition file.

## Related links

**Related concepts**  

[ui.Form.setElementImage](3157-ui-form-setelementimage.md "Change the image of form elements.")

[ui.Form.setFieldComment](3160-ui-form-setfieldcomment.md "Set the comment/hint of a form field.")
