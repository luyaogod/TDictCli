---
title: "TABLES section"
source: "fgl-topics/c_fgl_FormSpecFiles_TABLES_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > TABLES section"
type: "concept"
---

# TABLES section

> Defines the list of database tables referenced by form field definitions.

## Syntax

```
TABLES
[ alias = [database[@dbserver]:][owner.] ] table [,...]
[END]
```

1. alias represents an alias name for the given
   table.
2. table is the name of the database table.
3. database is the name of the database of the
   table (see warnings).
4. dbserver identifies the Informix® database server (INFORMIXSERVER)
5. owner is the name of the table owner (see warnings).

## Usage

The `TABLES` section lists every database table or view referenced in the form
specification file. This section is mandatory when form fields reference database columns defined in
the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").

The `TABLE` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `END` keyword
is optional.

The [`SCHEMA`
section](1710-schema-section.md "Defines the database schema file to be used to compile the form.") must also exist to define the database schema.

Field identifiers in programs or in other sections
of the form specification file can reference screen fields as column,
alias.column, or table.column.

The
same alias must also appear in screen interaction
statements of programs that reference screen fields linked to the
columns of a table that has an alias.

If
a table requires the name of an owner or of a database as
a qualifier, the `TABLES` section must also declare
an alias for the table. The alias can be the same
identifier as table.

For backward compatibility
with the Informix form
specification, the comma separator is optional and the database, dbserver
and owner specifications are ignored.

## Example

```
SCHEMA stores 
LAYOUT
GRID
{
  ...
}
END
TABLES
 customer, orders 
END
ATTRIBUTES
...
END
```

## Related links

**Related concepts**  

[Form fields](1670-form-fields.md "Form fields are form elements designed for data input and/or data display.")
