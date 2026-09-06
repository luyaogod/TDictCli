---
title: "ui.ComboBox.forName"
source: "fgl-topics/c_fgl_ClassCombo_forName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.forName"
type: "concept"
---

# ui.ComboBox.forName

> Search for a combobox in the current form.

## Syntax

```
ui.ComboBox.forName(
   name STRING )
  RETURNS ui.ComboBox
```

1. name is the name of `COMBOBOX` form field, with or without
   form field prefix (`[tabname.]colname`).

## Usage

The `ui.ComboBox.forName()` class method searches for a
`ui.ComboBox` object by form field name in the current form.

The name of the form field passed as parameter can use the same letter case as that in the form
definition file: The lookup is case-insensitive.

After loading a form with [`OPEN WINDOW WITH FORM`](../11_user-interface/1572-open-window.md "Creates and displays a new window."), use the class method to retrieve a
`ui.ComboBox` object into a variable defined as a `ui.ComboBox`.

```
DEFINE cb ui.ComboBox
LET cb = ui.ComboBox.forName("customer.cust_city")
```

Verify the function has returned an object, as the form field may not exist.

```
IF cb IS NULL THEN
  ERROR "Form field not found in current form"
  EXIT PROGRAM
END IF
```

Once instantiated, the `ui.ComboBox` object can be used,
for example to fill the items of the drop down list.

```
CALL cb.clear()
CALL cb.addItem(1,"Paris")
CALL cb.addItem(2,"London")
CALL cb.addItem(3,"Madrid")
```

## Related links

**Related concepts**  

[Example Get a ComboBox form field view and fill the item list](3264-example-get-a-combobox-form-field-view-and-fill-the-item-lis.md "Example Get a ComboBox form field view and fill the item list")
