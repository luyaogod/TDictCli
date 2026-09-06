---
title: "ui.ComboBox.getTextOf"
source: "fgl-topics/c_fgl_ClassCombo_getTextOf.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.getTextOf"
type: "concept"
---

# ui.ComboBox.getTextOf

> Get the item text by name.

## Syntax

```
getTextOf(
   code STRING )
  RETURNS STRING
```

1. code is the name of a combobox item.

## Usage

The `getTextOf()` method returns the display
label of the item identified by the name passed as
parameter.

The method returns `NULL` if the item name
does not exist.
