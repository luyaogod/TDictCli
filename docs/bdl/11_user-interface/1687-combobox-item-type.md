---
title: "COMBOBOX item type"
source: "fgl-topics/c_fgl_FormSpecFiles_COMBOBOX.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > COMBOBOX item type"
type: "concept"
---

# COMBOBOX item type

> Defines a line-edit with a drop-down list of values.

## COMBOBOX item basics

The `COMBOBOX` form item defines a field that can open a list of possible values
the end user can choose from.

![COMBOBOX rendering](../_images/FormItemType_COMBOBOX.jpg)

*COMBOBOX form item type*

All `COMBOBOX` items will be transmitted to the front-end. Consider limiting the
number of items to get better performances. When several hundred or thousand of items are possible,
use a [`BUTTONEDIT`](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.") instead of
a `COMBOBOX`.

## Defining a COMBOBOX

The values of the drop-down list are defined by the [`ITEMS`](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.") attribute. Define a simple list of values like
`("A","B","C","D", ... )` or a list of key/value pairs like in
`((1,"Paris"),(2,"Madrid"),(3,"London"))`. In the second case, the labels (city
names) display depending on the key value (the city number) held by the field.

```
COMBOBOX ...
   ITEMS=((1,"Paris"),(2,"Madrid"),(3,"London"));
```

Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") when
defining key/value pairs in the combobox
items:

```
COMBOBOX ...
   ITEMS=((1,%"cities.paris"),
          (2,%"cities.madrid"),
          (3,%"cities.london"));
```

The [`INITIALIZER`](1791-initializer-attribute.md "The INITIALIZER attribute allows you to specify an initialization function that will be automatically called by the runtime system to set up the form item.")
attribute allows you to define an initialization function for the `COMBOBOX`. This
function is invoked at runtime when the form is loaded, to fill the item list dynamically, for
example with database records. It is recommended that you use the `TAG` attribute, so
you can identify in the program the kind of `COMBOBOX` form item to be initialized.
The initialization function name is converted to lowercase by fglform.

```
COMBOBOX ...
   TAG = "city", INITIALIZER = city_module.cmb_init;
```

If neither `ITEMS` nor `INITIALIZER` attributes are specified, the
form compiler automatically fills the list of items with the values of the [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") attribute, when specified.
However, the item list will not automatically be populated with include range values (i.e. values
defined using the TO keyword). The `INCLUDE` attribute can be specified directly in
the form or indirectly in the schema files.

```
COMBOBOX ...
   INCLUDE=("A","B","C","D","E");
```

During an `INPUT` dialog, a `COMBOBOX` field value can only be one
of the values specified in the `ITEMS` attribute. If the field allows
`NULL` values, a `NULL` item can be placed anywhere in the combobox
item list, to satisfy end-user
preferences:

```
COMBOBOX ...
   ITEMS=((NULL,"<Undefined>"),
          (1,"Red"),
          (2,"Yellow"),
          (3,"Green"));
```

During an `INPUT` dialog, if the field allows `NULL` values, and no
`NULL` item is specified in the item list, a `NULL` item (with empty
label) will be added automatically at the end of the item list, to let the end-user set the field
value to `NULL`. If the field is defined with `NOT NULL`, no default
item for `NULL` will be added automatically.

During a `CONSTRUCT` dialog, the `COMBOBOX` drop down list allows
to select multiple items, to produce
`"value1|value2|..."` and generate the SQL
condition `"colname in
(value1,value2,...)"`. If only one item is selected,
the SQL condition will be `"colname=value"`. To
clear the query field, the user can deselect all items, or use the `DEL` key in the
drop down list. If the field allows `NULL`, automatic items will be added to specify
if a value is defined by using `!=` (to generate `"colname
is not null"`), or specify that the value is undefined by using `"="` (to
generate `"colname is null"`). If the item list contains an
explicit item for `NULL`, selecting this item during a `CONSTRUCT`
will set `=`, show the corresponding item label, and generate a
"`colname is null`" SQL condition.

Avoid to define `COMBOBOX` items for `NULL` and
`CONSTRUCT` operators `"="`, `"!="`: These are provided
by default when the field allows nulls. Item labels are localized by default, and if really needed,
they can be redefined with GBC customization.

A good practice is to deny nulls with the [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") attribute, and add a special item such as
`(0,"<Undefined>")` to identify a non-specified-value:

```
COMBOBOX ...
   NOT NULL,
   ITEMS=((0,"<Undefined>"),
          (1,"Red"),
          (2,"Yellow"),
          (3,"Green"));
```

When using the [`AUTONEXT`](1763-autonext-attribute.md "The AUTONEXT attribute forces the focus to automatically leave the current field when completed.")
attribute, the focus will automatically go to the next editable field, when an item is selected from
the combobox drop-down list:

```
COMBOBOX cb1 = customer.cust_city,
   AUTONEXT, ITEMS = ... ;
```

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [ComboBox style attributes](1634-combobox-style-attributes.md "ComboBox presentation style attributes apply to COMBOBOX elements.").

## Detecting COMBOBOX item selection

To inform the dialog when the value changes, define an `ON CHANGE` block for the
`COMBOBOX` field. The program can then react immediately to user changes in the
field:

```
-- Form file (grid layout)
COMBOBOX cb1 = customer.cust_city,
   ITEMS = ... ;

-- Program file:
ON CHANGE cust_city
   -- A new item was selected in the combobox list
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a COMBOBOX

A `COMBOBOX` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [COMBOBOX item definition](1734-combobox-item-definition.md "Defines attributes for an edit field with a drop-down list.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The size of a `COMBOBOX` widget is computed from the [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") and [`SAMPLE`](1814-sample-attribute.md "The SAMPLE attribute defines the text to be used to compute the width of a form field widget.") attributes, and by following
the layout rules as described in [Widget width inside hbox tags](1559-widget-width-inside-hbox-tags.md).

## Related links

**Related concepts**  

[Filling a COMBOBOX item list](2250-filling-a-combobox-item-list.md "The item list of COMBOBOX fields can be initialized at runtime.")
