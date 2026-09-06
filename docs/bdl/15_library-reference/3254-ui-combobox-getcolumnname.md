---
title: "ui.ComboBox.getColumnName"
source: "fgl-topics/c_fgl_ClassCombo_getColumnName.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods > ui.ComboBox.getColumnName"
type: "concept"
---

# ui.ComboBox.getColumnName

> Get the column name of the form field.

## Syntax

```
getColumnName()
  RETURNS STRING
```

## Usage

The `getColumnName()` method returns the form field column name. The form field
column name can be `NULL` if not defined at the form field level.

Use the `getTableName()` and `getColumnName()` methods together in
order to identify the form field associated with the `COMBOBOX`. This allows for the
identification of the combobox field in your program, for example to fill the drop down list with
the appropriate items.

```
IF cb.getTableName()||"."||cb.getColumnName()
   == "customer.cust_city" THEN
      CALL cb.clear()
      CALL cb.addItem(1, "Paris" )
      CALL cb.addItem(2, "London" )
      CALL cb.addItem(3, "Madrid" )
END IF
```

## Related links

**Related concepts**  

[ui.ComboBox.getTableName](3259-ui-combobox-gettablename.md "Get the table prefix of the form field.")
