---
title: "ui.Form.setElementImage"
source: "fgl-topics/c_fgl_ClassForm_setElementImage.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setElementImage"
type: "concept"
---

# ui.Form.setElementImage

> Change the image of form elements.

## Syntax

```
setElementImage(
   name STRING,
   image STRING )
```

1. name defines the name of the node, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).
2. image is the image to be set.

## Usage

Change the image/icon of a form element with the `setElementImage()` method.

Pass the identifier of the form element. The identifier is the element name as
defined in the form definition. All elements with this name will be affected. If you want to
distinguish form elements, use unique names in the form definition file.

## Related links

**Related concepts**  

[ui.Form.setElementText](3159-ui-form-setelementtext.md "Change the text of form elements.")

[ui.Form.setElementStyle](3158-ui-form-setelementstyle.md "Change the style of form elements.")
