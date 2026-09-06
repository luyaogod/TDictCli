---
title: "ui.ComboBox methods"
source: "fgl-topics/c_fgl_ClassCombo_methods.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The ComboBox class > ui.ComboBox methods"
type: "concept"
---

# ui.ComboBox methods

> Methods of the ui.ComboBox class.

| Name | Description |
| --- | --- |
| ui.ComboBox.forName( name STRING ) RETURNS ui.ComboBox | Search for a combobox in the current form. |
| ui.ComboBox.setDefaultInitializer( initializer STRING ) | Define the default initializer for combobox form items. |
| ui.ComboBox.setDefaultInitializerFunction( initializer FUNCTION( comboBox ui.ComboBox ) RETURNS () ) | Define the default initializer for combobox form items. |

| Name | Description |
| --- | --- |
| addItem( code STRING, text STRING ) | Add an element to the item list. |
| clear() | Clear the item list of a combobox. |
| getColumnName() RETURNS STRING | Get the column name of the form field. |
| getIndexOf( code STRING ) RETURNS INTEGER | Get an item position by name. |
| getItemCount() RETURNS INTEGER | Get the number of items. |
| getItemName( index INTEGER ) RETURNS STRING | Get an item name by position. |
| getItemText( index INTEGER ) RETURNS STRING | Get the item text by position. |
| getTableName() RETURNS STRING | Get the table prefix of the form field. |
| getTag() RETURNS STRING | Get the combobox tag value. |
| getTextOf( code STRING ) RETURNS STRING | Get the item text by name. |
| removeItem( code STRING ) | Remove an item by name. |
