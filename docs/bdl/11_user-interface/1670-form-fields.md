---
title: "Form fields"
source: "fgl-topics/c_fgl_FormSpecFiles_Form_Fields.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form fields"
type: "concept"
---

# Form fields

> Form fields are form elements designed for data input and/or data display.

## Purpose of form fields

A form field is a form item dedicated to data management. It associates a form item
with a screen record field. The screen record field is used to [bind program variables](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.") in interaction
instructions (for example, dialogs). The program variables are the data models for the form
fields.

There are different types of form fields:

- [Database column fields](1671-database-column-fields.md "Form fields defined with a table and column name get data type from the database schema file.")
- [Formonly fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.")
- [Phantom fields](1673-phantom-fields.md "A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).")
- [Aggregate fields](1674-aggregate-fields.md "An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.")

Form fields are identified by the field name in programs, and are grouped in screen records
(or screen arrays in the case of list containers). The interactive instruction must mediate
between screen record fields and database columns by using program variables.

Form fields are usually related to database column, which types are defined in the database
schema file.

## Position and size of a form field

The position and size of a form field is defined with an [item tag](1679-item-tags.md "Item tags define the position and size in a grid-based container.") in the form layout, while the rendering
and behavior is defined in the [`ATTRIBUTES`](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.") section. Both parts are bound by the name of the item tag. The
item tag name is local to the .per file and is not available at runtime: It is
just the key to bind the item tag (position) with the item definition (attributes).

In the next example, the "`f1`" item tag (in the `LAYOUT` section)
is linked to the "`vehicle.num`" form field definition (in the
`ATTRIBUTES` section), which references a column of the "`vehicle`"
table, defined in the "`carstore`" database schema:

```
SCHEMA carstore 
LAYOUT
GRID
{
Number:   [f1            ]
Name:     [f2                        ]
}
END
END
TABLES
  vehicle 
END
ATTRIBUTES
  EDIT f1 = vehicle.num, STYLE="keycol";
  EDIT f2 = vehicle.name, UPSHIFT;
END
```

See also [Grid-based layout](1549-grid-based-layout.md "A form file can define a grid-based layout within a tree of layout items.").

## Related links

**Related concepts**  

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")

## Child topics

- [Database column fields](1671-database-column-fields.md): Form fields defined with a table and column name get data type from the database schema file.
- [Formonly fields](1672-formonly-fields.md): FORMONLY form fields define their data type explicitly, with or without referencing a database columns.
- [Phantom fields](1673-phantom-fields.md): A PHANTOM field defines a screen-record field which is not rendered in the layout (it acts as a hidden field).
- [Aggregate fields](1674-aggregate-fields.md): An AGGREGATE field defines a screen-record field to display summary information for a TABLE column.
