---
title: "CHECKBOX item type"
source: "fgl-topics/c_fgl_FormSpecFiles_CHECKBOX.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > CHECKBOX item type"
type: "concept"
---

# CHECKBOX item type

> Defines a boolean or three-state checkbox field.

## CHECKBOX item basics

The `CHECKBOX` form item defines a field with a check box and a text label.

![CHECKBOX rendering](../_images/FormItemType_CHECKBOX.jpg)

*CHECKBOX form item type*

## Defining a CHECKBOX

The `TEXT` attribute defines the label to be displayed near the check box.
Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") for this
attribute.

The box shows a check mark when the form field contains the value defined in the [`VALUECHECKED`](1843-valuechecked-attribute.md "The VALUECHECKED attribute defines the value associated with a checkbox item when it is checked.") attribute (for
example: `"Y"`), and shows no check mark if the field value is equal to the value
defined by the [`VALUEUNCHECKED`](1844-valueunchecked-attribute.md "The VALUEUNCHECKED attribute defines the value associated with a checkbox item when it is not checked.")
attribute (for example: `"N"`). If you do not specify the `VALUECHECKED`
or `VALUEUNCHECKED` attributes, they respectively default to `TRUE`
(integer 1) and `FALSE` (integer 0).

By default, during an `INPUT` dialog, a `CHECKBOX` field can
have three states:

- Grayed ( `NULL` value )
- Checked ( `VALUECHECKED` value )
- Unchecked ( `VALUEUNCHECKED` value )

If the field is declared as `NOT NULL`, the initial state can be grayed if the
default value is `NULL`; once the user has changed the state of the
`CHECKBOX` field, it switches only between checked and unchecked states.

During a `CONSTRUCT`, a `CHECKBOX` field always has three
possible states (even if the field is `NOT NULL`), to allow the end user to clear
the search condition:

- Grayed (No search condition)
- Checked (Condition column = `VALUECHECKED` value)
- Unchecked (Condition column = `VALUEUNCHECKED` value)

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [CheckBox style attributes](1633-checkbox-style-attributes.md "CheckBox presentation style attributes apply to CHECKBOX elements.").

## Detecting CHECKBOX modification

To inform the dialog immediately when the value changes, define an `ON CHANGE`
block for the `CHECKBOX` field. The program can then react immediately to user
changes in the field:

```
-- Form file (grid layout)
CHECKBOX cb1 = order.ord_valid,
   VALUECHECKED="Y", VALUEUNCHECKED="N" ;

-- Program file:
ON CHANGE ord_valid
   -- The checkbox field has been modified
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a CHECKBOX

A `CHECKBOX` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [CHECKBOX item definition](1733-checkbox-item-definition.md "Defines attributes for a boolean or three-state checkbox field.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").
