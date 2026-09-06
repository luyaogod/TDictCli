---
title: "ui.ComboBox.getTag"
source: "fgl-topics/c_fgl_ClassCombo_getTag.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.getTag"
type: "concept"
---

# ui.ComboBox.getTag

> Get the combobox tag value.

## Syntax

```
getTag()
  RETURNS STRING
```

## Usage

The `getTag()` method returns the value define by the
[`TAG`](../11_user-interface/1827-tag-attribute.md "The TAG attribute can be used to identify the form item with a specific string.") attribute.

Use the tag to mark `COMBOBOX` form items with your own flags, in
order to adapt the configuration of the combobox dynamically by program. For
example, if `TAG` contains the token "short", fill the drop down list
with short names, otherwise fill with long names. The same code can then be used for
different `COMBOBOX` form fields.
