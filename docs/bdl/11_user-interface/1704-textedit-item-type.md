---
title: "TEXTEDIT item type"
source: "fgl-topics/c_fgl_FormSpecFiles_TEXTEDIT.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > TEXTEDIT item type"
type: "concept"
---

# TEXTEDIT item type

> Defines a multi-line edit field.

## TEXTEDIT item basics

The `TEXTEDIT` form item defines a text input field with multiple lines. This type
of element is typically used to handle large text values such as comments or addresses that would
not fit in a single-line edit field.

![TEXTEDIT rendering](../_images/FormItemType_TEXTEDIT.jpg)

*TEXTEDIT form item type*

Use a `VARCHAR(N)` or `STRING` variable to hold the data for a
`TEXTEDIT` form item.

## Defining a TEXTEDIT

Use the [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item.")
attribute to define vertical and/or horizontal scrollbars for the `TEXTEDIT` form
field. By default, when not specifying this attribute, `TEXTEDIT` fields get a
vertical scrollbar.

The [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") attribute can
be used to force the `TEXTEDIT` field to stretch when the parent container is
re-sized. Values can be `NONE`, `X`, `Y` or
`BOTH`. By default, this attribute is set to `NONE` for
`TEXTEDIT` fields.

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [TextEdit style attributes](1650-textedit-style-attributes.md "Textedit presentation style attributes apply to the TEXTEDIT element.").

## TAB and RETURN

By default, when the focus is in a `TEXTEDIT` field, the Tab key moves to the next
field, while the Return key adds a newline (ASCII 10) character in the text.

To control the user input when the Tab and Return keys are pressed, specify the [`WANTTABS`](1849-wanttabs-attribute.md "The WANTTABS attribute forces a text field to insert Tab characters in the text when the user presses the Tab key.") and [`WANTNORETURNS`](1848-wantnoreturns-attribute.md "The WANTNORETURNS attribute forces a text field to reject newline characters when the user presses the Return key.") attributes.

With `WANTTABS`, the Tab key is consumed by the `TEXTEDIT` field,
and a Tab character (ASCII 9) is added to the text. The user can still jump out of the field with
the Shift-Tab combination.

With `WANTNORETURNS`, the Return key is not intercepted or consumed by the
`TEXTEDIT` field, and the action corresponding to the Return key is triggered. The
user can still enter a newline character with Shift-Return or Ctrl-Return.

## Where to use a TEXTEDIT

A `TEXTEDIT` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [TEXTEDIT item definition](1747-textedit-item-definition.md "Defines attributes for a multi-line edit field.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The rendering of the `TEXTEDIT` widget can be controlled with the [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item."), [`STRETCH`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size.") and [`SCROLLBARS`](1816-scrollbars-attribute.md "The SCROLLBARS attribute can be used to specify scrollbars for a form item.")
attributes:

```
TEXTEDIT te1 = FORMONLY.comment, STRETCH=BOTH, SCROLLBARS=NONE;
```

## Field input length

By default, the input length in `TEXTEDIT` fields is defined by the program
variable.

There is no need to define the `SCROLL` attribute, except you explicitly specify
`SCROLLBARS=NONE`. When specifying `SCROLLBARS=NONE`, the
`TEXTEDIT` field will limit the maximum input length to the number of cells defined
by the [screen item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.").

For more details about the `SCROLL` attribute, see [Input length of form fields](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.").

## Making the TEXTEDIT read-only

Use the [`NOTEDITABLE`](1804-noteditable-attribute.md "The NOTEDITABLE attribute disables the text editor.")
attribute to prevent text modification by the user. This attribute is typically used to display a
large piece of text that the user is not required to modify (for example, to show the content of a
log file). Yet, the focus can still go to the field, if it is enabled.

## Rich Text HTML support

The `TEXTEDIT` form item can display and input HTML content, when setting the
[`textFormat`](1650-textedit-style-attributes.md) style attribute to `"html"`. With this feature,
the `TEXTEDIT` supports so-called "rich text" editing. For security reasons, consider
using the [`sanitize`](1650-textedit-style-attributes.md) style attribute when `textFormat` is
`"html"`.

![HTML format TEXTEDIT rendering](../_images/FormItemType_TEXTEDIT_HTML.jpg)

*HTML format TEXTEDIT form item type*

When using the HTML text format in `TEXTEDIT`, the [fgl\_dialog\_setcursor()](../15_library-reference/2747-fgl-dialog-setcursor.md "This function sets the position of the edit cursor in the current field.") and [fgl\_dialog\_setselection()](../15_library-reference/2756-fgl-dialog-setselection.md "Selects the text in the current field.") functions must be called carefully. Because of the rich text
format, having a corresponding cursor position / selection between displayed text and HTML
representation may be difficult, especially in the case of hidden parts.

Formatting options
are available (bold, font size, and so on) and can be controlled using an integrated toolbox.

For more details, see [Rich Text Editing in TEXTEDIT](2252-rich-text-editing-in-textedit.md "The TEXTEDIT form item provides a rich text editing feature based on HTML.").
