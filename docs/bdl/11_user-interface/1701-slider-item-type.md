---
title: "SLIDER item type"
source: "fgl-topics/c_fgl_FormSpecFiles_SLIDER.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > SLIDER item type"
type: "concept"
---

# SLIDER item type

> Defines a slider form item.

## SLIDER item basics

The `SLIDER` form item defines a field where the user can set a value in a given
range, such as a typical audio volume control widget where you can grab the slider handle to change
the value.

![SLIDER rendering](../_images/FormItemType_SLIDER.jpg)

*SLIDER form item type*

Use a `SMALLINT` or `INTEGER` variable with a
`SLIDER` form item, larger types like `BIGINT` or
`DECIMAL` are not supported.

## Defining a SLIDER

A `SLIDER` field allows the user to move a handle along a horizontal or vertical
groove and translates the handle's position into a value within the legal range.

The [`VALUEMIN`](1841-valuemin-attribute.md "The VALUEMIN attribute defines a lower limit of values displayed in widgets (such as progress bars).") and [`VALUEMAX`](1842-valuemax-attribute.md "The VALUEMAX attribute defines a upper limit of values displayed in widgets (such as progress bars).") attributes define
respectively the lower and upper integer limit of the slider information. Any value outside this
range will not be displayed. The step between two marks is defined by the [`STEP`](1820-step-attribute.md "The STEP attribute specifies how a value is increased or decreased in one step (by a mouse click or key up/down).") attribute. If
`VALUEMIN` and/or `VALUEMAX` are not specified, they default
respectively to 0 (zero) and 5.

The [`ORIENTATION`](1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally.")
attribute defines whether the `SLIDER` is displayed vertically or horizontally.

This item type is not designed for `CONSTRUCT`, as the user can only select one
value.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.").

## Detecting SLIDER item selection

To inform the dialog when a value changes, define an `ON CHANGE` block for the
`SLIDER` field. The program can then react immediately to user changes in the
field:

```
-- Form file (grid layout)
SLIDER s1 = options.opts_volume,
   VALUEMIN=0, VALUEMAX=100;

-- Program file:
ON CHANGE opts_volume
   -- A value changed in the slider
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a SLIDER

A `SLIDER` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [SLIDER item definition](1744-slider-item-definition.md "Defines attributes for a slider element.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Related links

**Related concepts**  

[Binding variables to form fields](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")
