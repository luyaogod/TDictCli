---
title: "TIMEEDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_TIMEEDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > TIMEEDIT item type"
type: "concept"
---

# TIMEEDIT item type

> Defines a line-edit field with a clock widget to pick a time.

## TIMEEDIT item basics

The `TIMEEDIT` form item defines a field that allows the user to edit 24H time
values, or time duration (intervals), with a specific clock widget for time input.

![TIMEEDIT rendering](../_images/FormItemType_TIMEEDIT.jpg)

*TIMEEDIT form item type*

To store `TIMEEDIT` field values, use the appropriate [`DATETIME HOUR TO MINUTE`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") or [`DATETIME HOUR TO SECOND`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") data type,
depending on the target front-end.

On desktop (GDC/GBC) and browser (GAS/GBC), the time value picker is generic and common to both
front-end platforms. On mobile devices (GMA and GMI), when the dialog is an `INPUT`,
the native time picker is used to edit the value. With a `CONSTRUCT`, the generic
time picker is used to let the user enter a search criteria.

## Defining a TIMEEDIT

No specific attribute is needed to define the rendering and behavior of a
`TIMEEDIT` field. Common data validation attributes such `NOT NULL`,
`REQUIRED`, `DEFAULT` are allowed.

The time display format is automatically taken from the front-end platform settings. For example,
time values can display in the 0-12 hour clock format (with AM/PM indicators), or in the 0-24 hour
clock format.

The [`FORMAT`](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute does not apply to form fields that are bound to a
`DATETIME` or `INTERVAL` variable.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.").

Depending on the front-end platform, the widget and time picker to render a
`TIMEEDIT` field may have the following limitations:

- Time editor and time picker may not handle the seconds.
- The time value picker/editor may not allow to specify `NULL`.
- `TIMEEDIT` fields may also be used to handle `INTERVAL` values of
  the class `HOUR TO {MINUTE|SECOND}`,
  in order to input a time duration. However, the time interval pickers are limited to 24H hours and
  allow only positive values. As result, not all values allowed in an `INTERVAL HOUR TO
  MINUTE` variable (such as -86 hours 23 minutes) can be displayed by such widgets.

## Detecting TIMEEDIT modification

To inform the dialog when a date is picked from the clock widget, define an
`ON CHANGE` block for the `TIMEEDIT` field. The program can
then react immediately to user changes in the field:

```
-- Form file (grid layout)
TIMEEDIT de1 = order.ord_shiptime,
   NOT NULL;

-- Program file:
ON CHANGE ord_shiptime
   -- A new time value was picked from the clock widget
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a TIMEEDIT

A `TIMEEDIT` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [TIMEEDIT item definition](1748-timeedit-item-definition.md "Defines attributes for a line-edit with a clock widget to pick a time.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Field input length

The input length in a `TIMEEDIT` fields is defined by the
(`DATETIME`) program variable. Define an item tag wide enough to fit all time value
digits (5 positions for `HH:MM`, 8 positions for `HH:MM:SS`). For more
details, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Related links

**Related concepts**  

[DATEEDIT item type](1688-dateedit-item-type.md "Defines a line-edit with a calendar widget to pick a date.")

[DATETIMEEDIT item type](1689-datetimeedit-item-type.md "Defines a line-edit with a calendar widget to pick a datetime.")
