---
title: "Formonly fields"
source: "fgl-topics/c_fgl_FormSpecFiles_FORMONLY_Form_Fields.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form fields > Formonly fields"
type: "concept"
---

# Formonly fields

> FORMONLY form fields define their data type explicitly, with or without referencing a database columns.

## Syntax

```
item-type item-tag = FORMONLY.field-name
   [ TYPE
      { LIKE [table.]column
      | datatype [NOT NULL] }
   ]
     [ , attribute-list ]  ;
```

where datatype is one of:

```
{ CHAR
| DECIMAL [(p[,s])]
| SMALLFLOAT
| REAL
| FLOAT
| MONEY [(p[,s])]
| INTEGER
| SMALLINT
| DATE
| VARCHAR
| TEXT
| BYTE
| INTERVAL interval-qualifier
| DATETIME datetime-qualifier
| BIGINT
| BOOLEAN
}
```

1. table is the name or alias of a table, synonym, or view, as declared in the
   `TABLES` section.
2. column is the name of a database column.
3. field-name is the identifier that will be used in programs to handle the
   field.
4. interval-qualifier is an `INTERVAL` qualification clause such
   as `HOUR(5) TO SECOND`.
5. datetime-qualifier is a `DATETIME` qualification clause such
   as `DAY TO SECOND`.

## Usage

Form fields can be specified with the `FORMONLY` prefix, when there is no
corresponding database column, or when the field must be defined with another name to that of the
database column.

> **Important:**
>
> The data type of a form field is only used by the `CONSTRUCT`
> interactive statement to do database queries. When using the form field with an
> `INPUT`, `INPUT ARRAY` or `DISPLAY ARRAY` dialog, the
> type of the program variable defines the data type of the form field.

When using the `LIKE [table.]column` syntax, the form field gets the
data type of the specific table column as defined in the [database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions."). The table name must be specified in the [`TABLES`](1726-tables-section.md "Defines the list of database tables referenced by form field definitions.") section.

When using the `TYPE datatype` clause, you explicitly specify
the type of the field.

For `CHAR`/`VARCHAR` data types, the size is defined by the item
tag length in the layout.

If no data type is specified, and no database column is referenced, the default data type
is `CHAR`.

Specifying a data type followed by the `NOT NULL` keyword is equivalent to
the [`NOT NULL`](1803-not-null-attribute.md "The NOT NULL attribute specifies that the field does not accept NULL values.") attribute.

The `STRING` data type is not supported in `FORMONLY` form
field definitions.

The definition of `FORMONLY` fields can be completed by using the [`DISPLAY LIKE`](1774-display-like-attribute.md "The DISPLAY LIKE attribute applies column attributes defined in the database schema files (.att) to a field.") and
[`VALIDATE LIKE`](1840-validate-like-attribute.md "The VALIDATE LIKE attribute applies column attributes defined in the .val database schema files to a field.")
attributes, to get the display and validation attributes from the .att
and .val database schema files.

## Example

```
LAYOUT
GRID
{
 [f001             ]
 [f002             ]
 ...
}
END
END
ATTRIBUTES
EDIT f001 = FORMONLY.total TYPE DECIMAL(10,2), NOENTRY ;
EDIT f002 = FORMONLY.name TYPE LIKE customer.cust_name,
            VALIDATE LIKE customer.cust_name ;
```

## Related links

**Related concepts**  

[Database column fields](1671-database-column-fields.md "Form fields defined with a table and column name get data type from the database schema file.")
