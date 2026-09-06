---
title: "ui.Form.setElementText"
source: "fgl-topics/c_fgl_ClassForm_setElementText.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Form class > ui.Form methods > ui.Form.setElementText"
type: "concept"
---

# ui.Form.setElementText

> Change the text of form elements.

## Syntax

```
setElementText(
   name STRING,
   text STRING )
```

1. name defines the name of the node, see [Identifying elements in ui.Form methods](3164-identifying-elements-in-ui-form-methods.md).
2. text is the text to be set.

## Usage

Change the text of a form element with the `setElementText()` method, for example
to modify the text of a static label or group box during program execution.

Pass the identifier of the form element. The identifier is the element name as
defined in the form definition. All elements with this name will be affected. If you want to
distinguish form elements, use unique names in the form definition file.

## Related links

**Related concepts**  

[Example 3: Change the title of table column headers](3168-example-3-change-the-title-of-table-column-headers.md "Example 3: Change the title of table column headers")

[ui.Form.setElementImage](3157-ui-form-setelementimage.md "Change the image of form elements.")

[ui.Form.setElementStyle](3158-ui-form-setelementstyle.md "Change the style of form elements.")
