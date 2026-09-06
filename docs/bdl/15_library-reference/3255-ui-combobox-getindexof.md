---
title: "ui.ComboBox.getIndexOf"
source: "fgl-topics/c_fgl_ClassCombo_getIndexOf.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.getIndexOf"
type: "concept"
---

# ui.ComboBox.getIndexOf

> Get an item position by name.

## Syntax

```
getIndexOf(
   code STRING )
  RETURNS INTEGER
```

1. code is the name of a combobox item.

## Usage

The `getIndexOf()` method takes an item name as parameter and returns the
position of the item in the drop down list.

The first item is at position 1. The method returns 0 (zero) if the item name does not
exist.

The following example checks for item existence, before adding the
item.

```
IF cb.getIndexOf("SFO") == 0 THEN
  CALL cb.addItem("SFO", "San Francisco International Airport, CA" )
END IF
```
