---
title: "SPINEDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_SPINEDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > SPINEDIT item type"
type: "concept"
---

# SPINEDIT item type

> Defines a spin box widget to enter integer values.

## SPINEDIT item basics

The `SPINEDIT` form item defines a field dedicated to numeric values. Depending on
the front-end platform, the widget may provide buttons to increment/decrement the field value.

![SPINEDIT rendering](../_images/FormItemType_SPINEDIT.jpg)

*SPINEDIT form item type*

Use a `SMALLINT` or `INTEGER` variable with a
`SPINEDIT` form item. Larger types like `BIGINT` or
`DECIMAL` are not supported.

## Defining a SPINEDIT

The increment between two values is defined by the [`STEP`](1820-step-attribute.md "The STEP attribute specifies how a value is increased or decreased in one step (by a mouse click or key up/down).") attribute:

```
SPINEDIT ...
   STEP = 5;
```

The [`VALUEMIN`](1841-valuemin-attribute.md "The VALUEMIN attribute defines a lower limit of values displayed in widgets (such as progress bars).") and [`VALUEMAX`](1842-valuemax-attribute.md "The VALUEMAX attribute defines a upper limit of values displayed in widgets (such as progress bars).") attributes define
respectively the lower and upper integer limit of the spin-edit range. There is no default minimum
or maximum value for the `SPINEDIT` widget.

This widget is not designed for `CONSTRUCT`, as you can only enter an integer
value.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [SpinEdit style attributes](1647-spinedit-style-attributes.md "SpinEdit presentation style attributes apply to an SPINEDIT element.").

## Detecting SPINEDIT modification

To inform the dialog when a value changes, define an `ON CHANGE` block for the
`SPINEDIT` field. The program can then react immediately to user changes in the
field:

```
-- Form file (grid layout)
SPINEDIT s1 = options.opts_rate,
   VALUEMIN=0, VALUEMAX=100, STEP=5;

-- Program file:
ON CHANGE opts_rate
   -- The value of the spinedit has changed
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a SPINEDIT

A `SLIDER` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [SPINEDIT item definition](1745-spinedit-item-definition.md "Defines attributes for a spin box widget to enter integer values.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Related links

**Related concepts**  

[Binding variables to form fields](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")
