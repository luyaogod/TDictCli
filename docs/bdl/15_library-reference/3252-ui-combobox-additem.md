---
title: "ui.ComboBox.addItem"
source: "fgl-topics/c_fgl_ClassCombo_addItem.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.addItem"
type: "concept"
---

# ui.ComboBox.addItem

> Add an element to the item list.

## Syntax

```
addItem(
   code STRING,
   text STRING )
```

1. code is the unique key that identifies the item.
2. text is the text to be displayed in the drop down list.

## Usage

The `addItem()` method adds an item to the end of the drop down list of the
`COMBOBOX`.

The first parameter is the value that can be set in the form field. The second
parameter is the label to be displayed in the drop down list. If the second
parameter is [NULL](../08_language-basics/0572-null.md "The NULL constant defines a non-value."),the runtime system
automatically uses the first parameter as the display value.

Uniqueness is not checked by the runtime system. Make sure that the items created are
unique, regarding the value key and the display label.

Trailing spaces must be avoided for the key parameter, because values get truncated when field
validation occurs, and the resulting value (without trailing spaces) will no longer match the
`COMBOBOX` item name when these contain trailing spaces. Additionally, trailing
spaces in the label parameter may cause the `COMBOBOX` to be much wider than
expected. To avoid such problems, use the `CLIPPED` operator, to trim the trailing
whitespaces of `CHAR`, `VARCHAR` or `STRING` values
passed to the `addItem()` method.

Note that best practice is to have an integer column defined as primary key in your SQL
tables.

## Related links

**Related concepts**  

[Example Using the INITIALIZER attribute in the form file](3265-example-using-the-initializer-attribute-in-the-form-file.md "Example Using the INITIALIZER attribute in the form file")
