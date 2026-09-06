---
title: "ui.ComboBox.getTableName"
source: "fgl-topics/c_fgl_ClassCombo_getTableName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.getTableName"
type: "concept"
---

# ui.ComboBox.getTableName

> Get the table prefix of the form field.

## Syntax

```
getTableName()
  RETURNS STRING
```

## Usage

The `getTableName()` method returns the name of the form field table
prefix. The form field table prefix can be `NULL` if not defined at
the form field level.

This allows for the identification of a `COMBOBOX` field in your program, for
example to fill the drop down list with the appropriate items.

## Related links

**Related concepts**  

[ui.ComboBox.getColumnName](3254-ui-combobox-getcolumnname.md "Get the column name of the form field.")
