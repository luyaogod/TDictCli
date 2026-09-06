---
title: "ui.ComboBox.removeItem"
source: "fgl-topics/c_fgl_ClassCombo_removeItem.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.removeItem"
type: "concept"
---

# ui.ComboBox.removeItem

> Remove an item by name.

## Syntax

```
removeItem(
   code STRING )
```

1. code is the name of a combobox item.

## Usage

The `removeItem()` method deletes an item from the list. The item to be
removed is identified by the name passed as a parameter. If the item does not exist, the
method returns without error.

The following example fills a combobox with a few items, and then removes one by using the item
name:

```
DEFINE cb ui.ComboBox
...
LET cb = ui.ComboBox.forName("formonly.airport")
CALL cb.addItem("CDG", "Paris-Charles de Gaulle, France")
CALL cb.addItem("LCY", "London-City Airport, UK")
CALL cb.addItem("LHR", "London-Heathrow, UK")
...
CALL cb.removeItem("CDG")
```

## Related links

**Related concepts**  

[Example Get a ComboBox form field view and fill the item list](3264-example-get-a-combobox-form-field-view-and-fill-the-item-lis.md "Example Get a ComboBox form field view and fill the item list")
