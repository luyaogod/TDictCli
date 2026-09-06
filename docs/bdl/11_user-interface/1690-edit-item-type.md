---
title: "EDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_EDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > EDIT item type"
type: "concept"
---

# EDIT item type

> Defines a simple line-edit field.

## EDIT item basics

The `EDIT` form item defines a field to enter a single-line of text, for
any data type.

![EDIT rendering](../_images/FormItemType_EDIT.jpg)

*EDIT form item type*

This item type is typically used to define character string and numeric form fields.

## Defining an EDIT

The `EDIT` item type can be used for any data type that can be converted
to editable text.

To show a hint to the user when the field has the focus, use the [`COMMENT`](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element.") attribute.

If the field is mandatory for input, combine the [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") with the [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") attribute.

The value accepted for the field can be limited with the [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") attribute.

To provide a default value, define the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attribute for the field.

Use the [`DOWNSHIFT`](1776-downshift-attribute.md "The DOWNSHIFT attribute forces character input to lowercase letters.") or [`UPSHIFT`](1838-upshift-attribute.md "The UPSHIFT attribute forces character input to uppercase letters.") attributes to force the letter case during input.

The [`FORMAT`](1779-format-attribute.md "The FORMAT attribute defines the data formatting of numeric and date fields, for input and display.") attribute can be
used to format numeric and date typed fields.

Input can be hidden (for example for password fields), with the [`INVISIBLE`](1793-invisible-attribute.md "The INVISIBLE attribute prevents field data being readable on the screen.")
attribute.

Especially for mobile devices, use the [`KEYBOARDHINT`](1798-keyboardhint-attribute.md "The KEYBOARDHINT attribute gives an indication of the kind of data the form field contains, allowing the front-end to adapt the keyboard accordingly.")
attribute to get a specific keyboard when entering values into the field.

Input completion proposals can be implemented with the [`COMPLETER`](1770-completer-attribute.md "The COMPLETER attribute enables autocompletion for the edit field.")
attribute.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [Edit style attributes](1637-edit-style-attributes.md "Edit presentation style attributes apply to an EDIT element.").

## Where to use an EDIT

An `EDIT` form item can be defined an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [EDIT item definition](1737-edit-item-definition.md "Defines attributes for a simple line-edit field.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Field input length

The input length in an `EDIT` fields is defined by the item tag and may need to
get the [`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field.") attribute. For
more details, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Related links

**Related concepts**  

[BUTTONEDIT item type](1685-buttonedit-item-type.md "Defines a line-edit with a push-button that can trigger an action.")
