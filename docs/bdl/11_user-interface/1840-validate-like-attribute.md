---
title: "VALIDATE LIKE attribute"
source: "fgl-topics/c_fgl_FSFAttributes_VALIDATE_LIKE.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > VALIDATE LIKE attribute"
type: "concept"
---

# VALIDATE LIKE attribute

> The VALIDATE LIKE attribute applies column attributes defined in the .val database schema files to a field.

## Syntax

```
VALIDATE LIKE [table.]column
```

1. table is the optional table name to qualify the column.
2. column is the name of the column used to search for validation rules.

## Usage

Specifying the `VALIDATE LIKE` attribute is equivalent to writing in the
field definition all the attributes that are assigned to
table.column in the [.val database schema file](../09_advanced-features/0796-column-validation-file-val.md "The .val database schema file holds functional and display attributes of database table columns.").

The .val attributes are taken automatically from the schema file if the
field is linked to table.column in the field name
specification. The `VALIDATE LIKE` attribute is usually specified for
`FORMONLY` fields.

The `VALIDATE LIKE` attribute is evaluated at compile time, not at runtime.
If the database schema file changes, recompile all your forms.

Even if all of the fields in the form are `FORMONLY`,
the `VALIDATE LIKE` attribute requires the form compiler
to access the database schema file that contains the description of table.

## Example

```
EDIT f001 = FORMONLY.fullname, VALIDATE LIKE customer.custname;
```

## Related links

**Related concepts**  

[Formonly fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.")

[VALIDATE](../08_language-basics/0702-validate.md "The VALIDATE instructions checks a variable value based on database schema validation rules.")
