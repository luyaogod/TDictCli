---
title: "DISPLAY LIKE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_DISPLAY_LIKE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > DISPLAY LIKE attribute"
type: "concept"
---

# DISPLAY LIKE attribute

> The DISPLAY LIKE attribute applies column attributes defined in the database schema files (.att) to a field.

## Syntax

```
DISPLAY LIKE [table.]column
```

1. table is the optional table name to qualify the column.
2. column is the name of the column to be used to retrieve display attributes.

## Usage

Specifying this attribute is equivalent to listing all the attributes that are assigned to
table.column in the database schema file with the [.att extension](../09_advanced-features/0797-column-video-attributes-file-att.md "The .att database schema file contains the default video attributes of database table columns.").

Display attributes are automatically taken from the schema file if the field is linked to
table.column in the field name specification.

The `DISPLAY LIKE` attribute is evaluated at compile time, not at runtime.
If the database schema file changes, recompile all forms using this attribute. Even if all
of the fields in the form are `FORMONLY`, this attribute requires the form
compiler to access the database schema file that contains the description of
table.

## Example

```
EDIT f001 = FORMONLY.fullname, DISPLAY LIKE customer.custname;
```

## Related links

**Related concepts**  

[Column Video Attributes File (.att)](../09_advanced-features/0797-column-video-attributes-file-att.md "The .att database schema file contains the default video attributes of database table columns.")

[Formonly fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.")
