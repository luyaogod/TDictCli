---
title: "PROGRESSBAR item type"
source: "fgl-topics/c_fgl_FormSpecFiles_PROGRESSBAR.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > PROGRESSBAR item type"
type: "concept"
---

# PROGRESSBAR item type

> Defines a progress indicator field.

## PROGRESSBAR item basics

The `PROGRESSBAR` form item defines a field that shows a progress indicator.

![PROGRESSBAR rendering](../_images/FormItemType_PROGRESSBAR.jpg)

*PROGRESSBAR form item type*

Use a `SMALLINT` or `INTEGER` variable with a
`PROGRESSBAR` form item. Larger types like `BIGINT` or
`DECIMAL` are not supported.

## Defining a PROGRESSBAR

The [`VALUEMIN`](1841-valuemin-attribute.md "The VALUEMIN attribute defines a lower limit of values displayed in widgets (such as progress bars).")
and [`VALUEMAX`](1842-valuemax-attribute.md "The VALUEMAX attribute defines a upper limit of values displayed in widgets (such as progress bars).")
attributes define respectively the lower and upper integer limit of the progress information.
Any value outside this range will not be displayed. Default values are `VALUEMIN=0`
and `VALUEMAX=100`.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute.
For more details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.")
and [ProgressBar style attributes](1645-progressbar-style-attributes.md "ProgressBar presentation style attributes apply to PROGRESSBAR elements.").

## Displaying PROGRESSBAR values

The position of the progress bar indicator is defined by the value of the corresponding form
field. The value can be changed by the program using the `DISPLAY TO` instruction,
to set the value of the field, or by changing the program variable bound to the field when using
the `UNBUFFERRED` dialog mode.

Progress information is typically displayed during non-interactive program code. To show changes
to the end user in this context, you need to use the [`ui.Interface.refresh()`](../15_library-reference/3117-ui-interface-refresh.md "Synchronize the user interface with the front-end.") method to force a refresh. To provide the best
feedback to the user, consider calling the `refresh()` method regularly but not too
often, otherwise you will overload the network traffic and bring down the front-end component.

For example, if you have to process 1000 rows, define `VALUEMIN=0` and
`VALUEMAX=1000` in the `PROGRESSBAR` item, and perform a refresh every
50 rows:

```
FOR row=1 TO 1000
    ...
    IF (row MOD 50) == 0 THEN
       LET myprogbar = row
       CALL ui.Interface.refresh()
    END IF
END FOR
```

## Where to use a PROGRESSBAR

A `PROGRESSBAR` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [PROGRESSBAR item definition](1741-progressbar-item-definition.md "Defines attributes for a progress indicator field.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Related links

**Related concepts**  

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
