---
title: "SCHEMA section"
source: "fgl-topics/c_fgl_FormSpecFiles_SCHEMA_section.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure > SCHEMA section"
type: "concept"
---

# SCHEMA section

> Defines the database schema file to be used to compile the form.

Each form specification file can begin with a `SCHEMA` section
identifying the database schema (if any) on which the form is based.
This can be any database schema that is defined with a database schema
file. Form field data types can be automatically extracted from the
schema file if you specify the table and column name in the form
field definition (see `ATTRIBUTES` section).

## Syntax 1

```
SCHEMA { database[@dbserver] | string | FORMONLY }
```

1. database is the name of the database schema to be used for the form compilation.
2. dbserver identifies the Informix®
   database server (INFORMIXSERVER).
3. string can be a string literal containing the database name.

## Syntax 2: (supported for backward compatibility)

```
DATABASE {  database[@dbserver] | string | FORMONLY } [ WITHOUT NULL INPUT ]
```

The `DATABASE` syntax
is supported for compatibility with Informix 4gl;
using `SCHEMA` is recommended.

1. database is the name of the database schema to be used for the form compilation.
2. dbserver identifies the Informix
   database server (INFORMIXSERVER)
3. string can be a string literal containing the database name.

## Usage

The `SCHEMA` (or `DATABASE`) defines the database schema to be used
to resolve data types for [database
column-based fields](1671-database-column-fields.md "Form fields defined with a table and column name get data type from the database schema file.").

The `DATABASE` instruction is supported for backward compatibility, we recommend
using `SCHEMA` instead.

The `SCHEMA` section must appear in the sequence described in [form file structure](1709-form-file-structure.md "A form specification file is defined by a set of sections.").

The `SCHEMA` section is optional; if you do not specify it, database schema
specification defaults to `SCHEMA FORMONLY`.

You can create a form that is not related to any database schema by using the `FORMONLY`
keyword after `SCHEMA`/`DATABASE`. When using this option, you must
omit the `TABLES` section and define field data types explicitly in the
`ATTRIBUTES` section.

The database and dbserver specifications are supported
(but ignored) for backward compatibility with Informix
form specifications.

When using a specific database schema, the field data types are taken from the schema file
during compilation. Make sure that the database schema file of the [development database](../03_general/0009-general-terms-used-in-this-documentation.md "This documentation uses general terms that must be clarified for a good understanding.") corresponds to the [production database](../03_general/0009-general-terms-used-in-this-documentation.md "This documentation uses general terms that must be clarified for a good understanding."); otherwise the form fields defined in the
compiled version of your forms will not match the table structures of the production database.

The use of the `WITHOUT NULL INPUT` option in the `DATABASE`
syntax is supported for backward compatibility, but is ignored.

## Example

```
SCHEMA stores
LAYOUT
 ...
```

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")

[TABLES section](1726-tables-section.md "Defines the list of database tables referenced by form field definitions.")

[ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
