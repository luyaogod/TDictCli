---
title: "ui.Dialog.getArrayLength"
source: "fgl-topics/c_fgl_ClassDialog_getArrayLength.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.getArrayLength"
type: "concept"
---

# ui.Dialog.getArrayLength

> Returns the total number of rows in the specified list.

## Syntax

```
getArrayLength(
   name STRING )
  RETURNS INTEGER
```

1. name is the name of the screen record, see [Identifying screen-arrays in ui.Dialog methods](3240-identifying-screen-arrays-in-ui-dialog-methods.md).

## Usage

The `getArrayLength()` method returns the total number of rows of an
`INPUT ARRAY` or `DISPLAY ARRAY` list. The name of the
screen array is passed as parameter to identify the list.

## Example

```
MAIN
  DEFINE custlist DYNAMIC ARRAY OF RECORD
            pkey INT, name VARCHAR(50)
         END RECORD

  OPEN FORM f1 FROM "form1"
  DISPLAY FORM f1

  DISPLAY ARRAY custlist TO sa_custlist.*
     BEFORE DISPLAY
         MESSAGE "Row count: " || DIALOG.getArrayLength("sa_custlist")
  END DISPLAY

END MAIN
```
