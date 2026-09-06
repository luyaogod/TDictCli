---
title: "RADIOGROUP item type"
source: "fgl-topics/c_fgl_FormSpecFiles_RADIOGROUP.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item types > RADIOGROUP item type"
type: "concept"
---

# RADIOGROUP item type

> Defines a mutual exclusive set of options field.

## RADIOGROUP item basics

The `RADIOGROUP` form item defines a field that provides several options that
the user can make a selection from. Checking one radio button unchecks any previously checked
button within the same group.

![RADIOGROUP rendering](../_images/FormItemType_RADIOGROUP.jpg)

*RADIOGROUP form item type*

## Defining a RADIOGROUP

A `RADIOGROUP` defines a set of radio buttons where each button is associated
with a value defined in the [`ITEMS`](1795-items-attribute.md "The ITEMS attribute defines a list of possible values that can be used by the form item.")
attribute.

The text associated with each item value will be used as the label of the corresponding radio
button, for example: `ITEMS=((1,"Beginner"), (2,"Normal"), (3,"Expert"))` will
create three radio buttons with the texts `Beginner`, `Normal` and
`Expert`, respectively.

```
RADIOGROUP ...
   ITEMS=((1,"Beginner"),(2,"Normal"),(3,"Expert"));
```

Consider using [localized strings](../09_advanced-features/0901-localized-strings.md "Localized strings provide a means of writing applications in which the text of strings can be customized on site.") when
defining key/value pairs in the radio group items:

```
RADIOGROUP ...
   ITEMS=((1,%"skills.beginner"),
          (2,%"skills.normal"),
          (3,%"skills.expert"));
```

If the `ITEMS` attribute is not specified, the form compiler automatically fills
the list of items with the values of the [`INCLUDE`](1790-include-attribute.md "The INCLUDE attribute defines a list of possible values for a field.") attribute, when specified.
However, the item list will not automatically be populated with include range values (that is values
defined using the `TO` keyword). The `INCLUDE` attribute can be
specified directly in the form or indirectly in the schema files.

During an `INPUT`, a `RADIOGROUP` field value can only be one of
the values specified in the `ITEMS` attribute. During a `CONSTRUCT`,
a `RADIOGROUP` field allows all items to be unchecked (even if the field is
`NOT NULL`), to let the user clear the search condition.

If one of the items is explicitly defined with `NULL` and the `NOT NULL`
attribute is omitted, in `INPUT`, selecting the corresponding radio button sets the
field value to null. In `CONSTRUCT`, selecting the radio button corresponding to null
will be equivalent to the equals (`=`) query operator, which will generate a
"`colname is null`" SQL condition.

Use the [`ORIENTATION`](1805-orientation-attribute.md "The ORIENTATION attribute defines whether an element displays vertically or horizontally.")
attribute, to define if the radio group items must be arranged vertically or
horizontally:

```
RADIOGROUP ...
   ITEMS=(...),
   ORIENTATION = HORIZONTAL;
```

A radio group can adapt its orientation for [responsive layout](1541-responsive-layout.md "Forms can be designed to adapt to the front-end screen possibilities."), if you specify a `@screen-size`
modifier for the `ORIENTATION` attribute:

```
RADIOGROUP ...
   ORIENTATION@SMALL = VERTICAL, -- no really needed: default is vertical
   ORIENTATION@MEDIUM = HORIZONTAL,
   ORIENTATION@LARGE = HORIZONTAL;
```

Front-ends support different presentation and behavior options, which can be controlled by a
[`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute. For more
details, see [Style attributes common to all elements](1629-style-attributes-common-to-all-elements.md "Common presentation style attributes apply to any graphical element, such as windows, layout containers, or form items.").

## Detecting RADIOGROUP item selection

To inform the dialog when a value change, define an `ON CHANGE` block for the
`RADIOGROUP` field. The program can then react immediately to user changes in the
field:

```
-- Form file (grid layout)
RADIOGROUP rg1 = user.user_skill,
   ITEMS = ... ;

-- Program file:
ON CHANGE user_skill
   -- An new item was selected in the radiogroup
```

For more details, see [Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.").

## Where to use a RADIOGROUP

A `RADIOGROUP` form item can be defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") and a [RADIOGROUP item definition](1742-radiogroup-item-definition.md "Defines attributes for a mutually-exclusive set of option fields.") in a [`GRID`](1692-grid-item-type.md "Defines a layout area based on a grid of cells."), [`SCROLLGRID`](1700-scrollgrid-item-type.md "Defines a scrollable grid view widget.") and [`TABLE`](1703-table-item-type.md "Defines a list view widget.")/[`TREE`](1706-tree-item-type.md "Defines a tree view widget.").
