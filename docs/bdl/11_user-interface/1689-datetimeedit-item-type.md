---
title: "DATETIMEEDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_DATETIMEEDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > DATETIMEEDIT item type"
type: "concept"
---

# DATETIMEEDIT item type

> Defines a line-edit with a calendar widget to pick a datetime.

## DATETIMEEDIT item basics

The `DATETIMEEDIT` form item defines a field that can open a calendar to ease
date-time input.

![DATETIMEEDIT rendering](../_images/FormItemType_DATETIMEEDIT.jpg)

*DATETIMEEDIT form item type*

To store `DATETIMEEDIT` field values, use a [`DATETIME YEAR TO MINUTE`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") or [`DATETIME YEAR TO SECOND`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") program
variable, depending on the target front-end.

On desktop (GDC/GBC) and browser (GAS/GBC), the date/time value picker is generic and common to
both front-end platforms. On mobile devices (GMA and GMI), when the dialog is an
`INPUT`, the native date/time picker is used to edit the value. With a
`CONSTRUCT`, the generic date/time picker is used to let the user enter a search
criteria.

## Defining a DATETIMEEDIT

The `DATETIMEEDIT` form item type allows the user to edit date-time values with a
specific widget for date-time input. A `DATETIMEEDIT` field typically provides a
calendar and clock widget, to let the end user pick a date and time from it.

When using a `DATETIME` variable as recommended, with desktop (GDC/GBC) and
browser (GAS/GBC) front-ends, the format of `DATETIMEEDIT` fields is defined by the
[DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") environment variable. On mobile platforms,
the date format is defined by the device OS language settings. It is recommended to use the same
DBDATE setting as in the device settings.

The [`FORMAT`](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute does not apply to form fields that are bound to a
`DATETIME` or `INTERVAL` variable.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [DateTimeEdit style attributes](1636-datetimeedit-style-attributes.md "DateEdit presentation style attributes apply to DATEEDIT elements.").

Depending on the front-end platform, the widget and time picker to render a
`DATETIMEEDIT` field may have the following limitations:

- Date-time editor and calendar may not handle seconds and some front-ends may deny data types
  different from `DATETIME YEAR TO {MINUTE|SECOND}`.
- The date-time value picker/editor may not allow to specify `NULL`.

## Detecting DATETIMEEDIT calendar selection

To inform the dialog when a date-time is picked from the calendar widget, define an `ON
CHANGE` block for the `DATETIMEEDIT` field. The program can then react
immediately to user changes in the field:

```
-- Form file (grid layout)
DATETIMEEDIT dt1 = order.ord_shipdate,
   NOT NULL;

-- Program file:
ON CHANGE ord_shipdate
   -- A new date-time value was picked from the calendar
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a DATETIMEEDIT

A `DATETIMEEDIT` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [DATETIMEEDIT item definition](1736-datetimeedit-item-definition.md "Defines attributes for a line-edit field with a calendar widget to pick a datetime.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The size of a `DATETIMEEDIT` widget is computed by following the layout rules as
described in [Widget width inside hbox tags](1559-widget-width-inside-hbox-tags.md).

## Field input length

The input length in a `DATETIMEEDIT` fields is defined by the
(`DATETIME`) program variable. Define an item tag with enough positions to be able to
display dates with 4 year digits. For more details, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Related links

**Related concepts**  

[TIMEEDIT item type](1705-timeedit-item-type.md "Defines a line-edit field with a clock widget to pick a time.")
