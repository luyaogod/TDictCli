---
title: "ATTRIBUTES section"
source: "fgl-topics/c_fgl_FormSpecFiles_ATTRIBUTES_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > ATTRIBUTES section"
type: "concept"
---

# ATTRIBUTES section

> The ATTRIBUTES section describes properties of elements used in the form.

## Syntax

```
ATTRIBUTES
 { form-field-definition
 | phantom-field-definition
 | form-item-definition }
 [...]
[END]
```

where form-field-definition
is:

```
item-type item-tag = field-name [ , attribute-list ] ;
```

where phantom-field-definition
is:

```
PHANTOM field-name ;
```

where form-item-definition
is:

```
item-type item-tag: item-name [ , attribute-list ] ;
```

1. item-type defines the type of the Form Item.
2. item-tag is the name of the screen element used in the
   `LAYOUT` section.
3. field-name defines the name of the screen record field.
4. item-name identifies the form item that is not a form field containing
   data.
5. attribute-list defines the aspect and behavior of the form item.

where attribute-list
is:

```
attribute [,...]
```

1. The attribute list is a comma-separated list of attributes.

where attribute is:

```
attribute-name [ = { value | value-list } ]
```

1. attribute identifies the attribute of the form item.

where value-list
is:

```
( { value | sub-value-list } [,...])
```

1. value is a string, date or numeric literal,
   or predefined constant like `TODAY`.
2. sub-value-list is a set of values separated
   by comma, to support subset definitions as in
   "`(1,(21,22),(31,32,33))`".

## Usage

The `ATTRIBUTES` section is required to define the attributes for the form items
used in containers of the `LAYOUT` section.

The `ATTRIBUTES` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `END` keyword is optional.

Every item-tag used in the `LAYOUT` section must get an item definition in the
`ATTRIBUTES` section.

A form item definition is associated by name to an item tag or layout tag defined in the
container.

In order to define a form field, the form item definition must use the equal sign notation to
associate a screen record field with the form item. If the form item is not associated with a screen
record field (for example, a push button), you must use the colon notation.

To match the complete structure of a database table record, additional fields can be defined as
phantom fields, when no corresponding item tag is used in the layout.

Form item definitions can optionally include an attribute-list to specify the
appearance and behavior of the item. For example, you can define acceptable input values, on-screen
comments, and default values for fields.

When no screen record is defined in the `INSTRUCTION` section, a default screen
record is built for each set of form items declared with the same table name.

The order in which you list the form items determines the order of fields in the default screen
records that the form compiler creates for each table.

To define form items as form fields, you are not required to specify table
unless the name column is not unique within the form specification. However, it
is recommended that you always specify table.column rather than the unqualified
column name. As you can refer to field names collectively through a screen record
built upon all the fields linked to the same table, your forms might be easier to work with if you
specify table for each field.

When used in a table, some widgets are rendered only when the user enters in the field. For
example `RadioGroup`, `CheckBox`, `ComboBox`,
`ProgressBar`.

## Example

```
SCHEMA game 
LAYOUT
GRID
{
  ...
}
END
TABLES
 player 
END
ATTRIBUTES
 EDIT f001 = player.name, REQUIRED,
               COMMENT="Enter player's name";
 EDIT f002 = player.ident, NOENTRY;
 COMBOBOX f003 = player.level, NOT NULL,
      ITEMS=((1,"Beginner"), (2,"Normal"),(3,"Expert"));
 CHECKBOX f004 = FORMONLY.winner,
      VALUECHECKED=1, VALUEUNCHECKED=0,
      TEXT="Winner";
 BUTTON b1: print, TEXT="Print Report";
 GROUP g1: print, TEXT="Description";
END
```

## Related links

**Related concepts**  

[Form items](1665-form-items.md "The concept of form item includes all elements used in the definition of a form.")

## Child topics

- [AGGREGATE item definition](1728-aggregate-item-definition.md): Defines screen-record fields that hold computed values to be displayed as footer cells in a TABLE container.
- [PHANTOM item definition](1729-phantom-item-definition.md): Defines a form field, that must not be displayed to the end user.
- [BUTTON item definition](1730-button-item-definition.md): Defines attributes for a push-button that can trigger an action.
- [BUTTONEDIT item definition](1731-buttonedit-item-definition.md): Defines attributes for a line-edit field with a push-button that can trigger an action.
- [CANVAS item definition](1732-canvas-item-definition.md): Defines attributes for a CANVAS drawing area.
- [CHECKBOX item definition](1733-checkbox-item-definition.md): Defines attributes for a boolean or three-state checkbox field.
- [COMBOBOX item definition](1734-combobox-item-definition.md): Defines attributes for an edit field with a drop-down list.
- [DATEEDIT item definition](1735-dateedit-item-definition.md): Defines attributes for a line-edit field with a calendar widget to pick a date.
- [DATETIMEEDIT item definition](1736-datetimeedit-item-definition.md): Defines attributes for a line-edit field with a calendar widget to pick a datetime.
- [EDIT item definition](1737-edit-item-definition.md): Defines attributes for a simple line-edit field.
- [GROUP item definition](1738-group-item-definition.md): Defines attributes for a group-box layout tag.
- [IMAGE item definition](1739-image-item-definition.md): Defines attributes for an area that can display an image resource.
- [LABEL item definition](1740-label-item-definition.md): Defines attributes for a simple text area to display a read-only value.
- [PROGRESSBAR item definition](1741-progressbar-item-definition.md): Defines attributes for a progress indicator field.
- [RADIOGROUP item definition](1742-radiogroup-item-definition.md): Defines attributes for a mutually-exclusive set of option fields.
- [SCROLLGRID item definition](1743-scrollgrid-item-definition.md): Defines attributes for a scrollgrid layout tag.
- [SLIDER item definition](1744-slider-item-definition.md): Defines attributes for a slider element.
- [SPINEDIT item definition](1745-spinedit-item-definition.md): Defines attributes for a spin box widget to enter integer values.
- [TABLE item definition](1746-table-item-definition.md): Defines attributes for a table layout tag.
- [TEXTEDIT item definition](1747-textedit-item-definition.md): Defines attributes for a multi-line edit field.
- [TIMEEDIT item definition](1748-timeedit-item-definition.md): Defines attributes for a line-edit with a clock widget to pick a time.
- [TREE item definition](1749-tree-item-definition.md): Defines attributes for a tree layout tag.
- [WEBCOMPONENT item definition](1750-webcomponent-item-definition.md): Defines attributes for a generic form field that can receive an external widget.
