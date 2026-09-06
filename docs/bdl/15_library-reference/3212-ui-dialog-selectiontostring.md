---
title: "ui.Dialog.selectionToString"
source: "fgl-topics/c_fgl_ClassDialog_selectionToString.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.selectionToString"
type: "concept"
---

# ui.Dialog.selectionToString

> Serializes data of the selected rows.

## Syntax

```
selectionToString(
   name STRING )
  RETURNS STRING
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

The `selectionToString()` method
can be used to get a tab-separated value list of the [selected rows](3231-ui-dialog-setselectionmode.md "Defines the row selection mode for the specified list.").

When
multi-row selection is disabled, the method serializes the current
row.

You typically use this method along with drag &
drop to fill the buffer, by using a text/plain MIME type, to export
data to external applications.

```
ON ACTION serialize
  LET buff = DIALOG.selectionToString( "sr" )
```

Numeric and date data will be formatted based on current locale settings ([DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined."), [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.")).

The visual presentation of data is respected. The dialog will copy the rows in the sort order
specified by the user, moved columns will appear in the same positions as in the table and hidden
columns will be ignored. Note that phantom columns are not copied.

Items in the tab-separated
record will be surrounded by double-quotes if the value contains special
characters such as a newline, a double-quote, or controls characters
with ASCII code < 0x20. Double-quotes in the value will be doubled.
