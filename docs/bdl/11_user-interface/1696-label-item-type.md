---
title: "LABEL item type"
source: "fgl-topics/c_fgl_FormSpecFiles_LABEL.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > LABEL item type"
type: "concept"
---

# LABEL item type

> Defines a simple text area to display a read-only value.

## LABEL item basics

The `LABEL` form item defines a read-only text area.

![LABEL rendering](../_images/FormItemType_LABEL.jpg)

*LABEL form item type*

## Defining a LABEL

A `LABEL` form item can be defined as a form field image or as a static label.
Use a form field label when the text changes often during program execution (for example, to
display text from the database). Use a static label if the text remains the same during program
execution.

Front-ends support different presentation and behavior options, which can be controlled by
a [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.") and [Label style attributes](1642-label-style-attributes.md "Label presentation style attributes apply to LABEL elements.").

## Form field LABEL item

Use a form field label item to display values that change often during program execution, for
example if the text is stored in the database.

The label text is defined by the value of the field.

The value can be changed by the program using the `DISPLAY BY NAME / DISPLAY TO`
instruction, or just by changing the value of the program variable bound to the label field when
using the `UNBUFFERED` mode in an interactive instruction.

When defining the `LABEL` item in the form, use a field name to identify the
element in programs:

```
LABEL f001 = cars.description;
```

## Static LABEL item

Use a static label item to display text that does not change during program execution.

This kind of item is not affected by instructions such as `CLEAR FORM` or the
`DISPLAY TO` instruction.

```
LABEL lab1: label1, TEXT="Name:";
```

Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") to ease
application
internationalization:

```
LABEL ...
   TEXT = %"label.customer.name";
```

Static labels display only character text values, and therefore do not follow any justification
rule as form field labels.

## Multi-line text in LABELs

In order to display label text on several lines, the text must contain `\n`
line-feed characters:

```
LABEL lab1: label1,
      TEXT="First line.\nSecond line.";
```

## Where to use a LABEL

A `LABEL` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [LABEL item definition](1740-label-item-definition.md "Defines attributes for a simple text area to display a read-only value.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").

## Defining the widget size

The size of a `LABEL` widget can be controlled by using the [`SIZEPOLICY`](1817-sizepolicy-attribute.md "The SIZEPOLICY attribute is a sizing directive based on the content of a form item.") attribute.

By default (`SIZEPOLICY=INITIAL`), labels adapt their width to the initial text
displayed by the element.

With [static
labels](1696-label-item-type.md), the initial text is usually defined in the form file, or with a localized string, and
the size does not need to adapt once the label is displayed.

When using [form field
labels](1696-label-item-type.md), the same default rule applies. However, if the initial text displayed in the form
field is `NULL` or smaller than other texts that will be displayed, the size of the
label element will not adapt after the initial text is displayed; the label stays at the size of the
initial displayed text.

> **Important:**
>
> When using a form field label, make sure that the size of the label will be
> large enough to display all possible values. To control the label size, use the
> `SIZEPOLICY=DYNAMIC` attribute, or use `SIZEPOLICY=FIXED` and define
> the form item with a sufficient size in the `LAYOUT` section.
