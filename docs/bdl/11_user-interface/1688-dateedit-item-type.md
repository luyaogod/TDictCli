---
title: "DATEEDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_DATEEDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > DATEEDIT item type"
type: "concept"
---

# DATEEDIT item type

> Defines a line-edit with a calendar widget to pick a date.

## DATEEDIT item basics

The `DATEEDIT` form item defines a field that can open a calendar to ease date
input.

![DATEEDIT rendering](../_images/FormItemType_DATEEDIT.jpg)

*DATEEDIT form item type*

To store `DATEEDIT` field values, use a [`DATE`](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") program variable with this form item.

On desktop (GDC/GBC) and browser (GAS/GBC), the date value picker is generic and common to both
front-end platforms. On mobile devices (GMA and GMI), when the dialog is an `INPUT`,
the native date picker is used to edit the value. With a `CONSTRUCT`, the generic
date picker is used to let the user enter a search criteria.

## Defining a DATEEDIT

The `DATEEDIT` form item type allows the user to edit date values with a specific
widget for date input. A `DATEEDIT` field typically provides a calendar widget, to
let the end user pick a date from it.

When using a `DATE` variable as recommended, with desktop front-ends, the format
of `DATEEDIT` fields is by default defined by the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment variable. On mobile platforms, the
date format is defined by the device OS language settings. It is recommended to use the same DBDATE
setting as in the device settings.

Specific `DATE` format can be defined with the [`FORMAT`](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute, but it is
recommended to use the default date formatting.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [DateEdit style attributes](1635-dateedit-style-attributes.md "DateEdit presentation style attributes apply to DATEEDIT elements.").

Depending on the front-end platform, the widget and time picker to render a
`DATEEDIT` field may have the following limitations:

- The front-end may deny data types different from `DATE`.
- The date value picker/editor may not allow to specify `NULL`.

## Detecting DATEEDIT calendar selection

To inform the dialog when a date is picked from the calendar widget, define an
`ON CHANGE` block for the `DATEEDIT` field. The program can then react
immediately to user changes in the field:

```
-- Form file (grid layout)
DATEEDIT de1 = order.ord_shipdate,
   NOT NULL;

-- Program file:
ON CHANGE ord_shipdate
   -- A new date value was picked from the calendar
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a DATEEDIT

A `DATEEDIT` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [DATEEDIT item definition](1735-dateedit-item-definition.md "Defines attributes for a line-edit field with a calendar widget to pick a date.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The size of a `DATEEDIT` widget is computed by following the layout rules as
described in [Widget width inside hbox tags](1559-widget-width-inside-hbox-tags.md).

## Field input length

The input length in a `DATEEDIT` fields is defined by the (`DATE`)
program variable. Define an item tag with 10 positions, to be able to display dates with 4 year digits.
For more details, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Calendar configuration

A set of [presentation style
attributes for the DateEdit class](1635-dateedit-style-attributes.md "DateEdit presentation style attributes apply to DATEEDIT elements.") can be used to customize the calendar. For example, you can
define the icons of the button to open the calendar, the days off, the first day of the week, and
the type of pop-up window for the calendar.

## Related links

**Related concepts**  

[TIMEEDIT item type](1705-timeedit-item-type.md "Defines a line-edit field with a clock widget to pick a time.")
