---
title: "Database column fields"
source: "fgl-topics/c_fgl_FormSpecFiles_Database_column_Fields.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file concepts > Form items > Form fields > Database column fields"
type: "concept"
---

# Database column fields

> Form fields defined with a table and column name get data type from the database schema file.

## Syntax

```
item-type item-tag = [table.]column
    [ , attribute-list ]  ;
```

1. item-type references an item type like `EDIT`.
2. item-tag identifies the layout location of the field.
3. table is the name or alias of a table, synonym, or view, as declared in the
   `TABLES` section.
4. column is the name of a database column.
5. attribute-list is a list of field attributes.

## Usage

A form field is typically based on the definition of a database column found in the [database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") specified with the
`SCHEMA` clause at the beginning of the form file. The database column defines the
data type of the form field.

> **Important:**
>
> The data type of a form field is only used by the `CONSTRUCT`
> interactive statement to do database queries. When using the form field with an
> `INPUT`, `INPUT ARRAY` or `DISPLAY ARRAY` dialog, the
> type of the program variable defines the data type of the form field.

In order to reference database columns, the table name must be listed in the [`TABLES`](1726-tables-section.md "Defines the list of database tables referenced by form field definitions.") section of the
form.

Fields are associated with database columns only during the compilation of the form specification
file: The form compiler examines the database schema file to identify the data type of the column,
and defines the form field with this type. This technique allows form field data types in the schema
files to be centralized. If the data type of a column changes, extract the schema again and
recompile your forms to take the new type into account.

The compilers also grabs other field attributes like validation rules and video display
attributes from .val and .att schema files. However, this
is supported for backward compatibility only (formerly stored in syscolval and syscolatt database
tables). Consider reviewing programs using this feature.

After the form compiler identifies data types from the schema file, the association between
fields and database columns is broken, and the form cannot distinguish the name or synonym of a
table or view from the name of a screen record.

The programs only have access to [screen
record fields](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."), in order to display or input data using program variables. Regardless of how
you define them, there is no implicit relationship between the values of program variables, form
fields, and database columns. Even, for example, if you declare a variable `lname LIKE
customer.lname`, the changes that you make to the variable do not imply any change in the
column value. Functional relationships among these entities must be specified in the program code,
through screen interaction statements, and through SQL statements. It is up to the programmer to
determine what data a form displays and what to do with data values that the user enters into the
fields of a form. You must indicate the binding explicitly in any statement that connects variables
to forms or to database columns.

If a form field is declared with a table column using the `SERIAL`,
`SERIAL8` or `BIGSERIAL` SQL type, the field will automatically get
the [`NOENTRY`](1801-noentry-attribute.md "The NOENTRY attribute prevents data entry in the field during an input dialog.") attribute, except
if the field is defined with the `TYPE LIKE` syntax.

## Example

```
SCHEMA stores  -- Database schema
LAYOUT
GRID
{
 [f001             ]
 ...
}
END
END
TABLES
  customer  -- Database table
END
ATTRIBUTES
EDIT f001 = customer.fname,  -- DB-col form field
     REQUIRED, COMMENTS="Customer name";
...
```

## Related links

**Related concepts**  

[Formonly fields](1672-formonly-fields.md "FORMONLY form fields define their data type explicitly, with or without referencing a database columns.")
