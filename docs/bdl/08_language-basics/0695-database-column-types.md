---
title: "Database column types"
source: "fgl-topics/c_fgl_variables_006.html"
breadcrumb: "Language basics > Variables > Database column types"
type: "concept"
---

# Database column types

> Simple variables and record structures can be defined from database columns types.

Variables defined with the `LIKE` keyword get the same data type as the table
column of a [database
schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").

> **Important:**
>
> Database schema files are generated with the fgldbsch
> tool. Column data types are read from the schema file during compilation. Make sure that your schema
> files correspond exactly to the production database.

For example:

```
SCHEMA stores
DEFINE cust_name LIKE customer.cust_name
MAIN
  DEFINE rec_cust RECORD LIKE customer.*
  ...
END MAIN
```

A `SCHEMA` statement must define the database name identifying the database schema
files to be used.

Alternatively, specify the database schema file followed by a colon, before the table name. This
allows you to use several database schemas at the same
time:

```
DEFINE rec_city      RECORD LIKE base:city.*
DEFINE rec_country   RECORD LIKE base:country.*
DEFINE rec_customer  RECORD LIKE orders:customer.*
DEFINE rec_item      RECORD LIKE stock:item.*
```

At runtime, a program typically connects to a single database source. Using multiple database
schemas is a programming feature.

The database schema files must exist and must be located in one of the directories specified in
the FGLDBPATH environment variable.

For `CHAR/VARCHAR` types, the size in the `.sch` file, the size in
the .sch file is expressed in character units, and is then interpreted by the
compiler as a number of bytes or characters, depending on the FGL\_LENGTH\_SEMANTICS environment
variable as when using directly `CHAR(N)` or `VARCHAR(N)`.
For more details, see [Extracting database schemas](../09_advanced-features/0881-length-semantics-settings.md).

When using database views, the column cannot be based on an aggregate function like
`SUM()`.

If `LIKE` references a `SERIAL` column, the variable will be
defined with the `INTEGER` data type. If `LIKE` references an
`INT8`, `SERIAL8` or `BIGSERIAL` column, the variable
will be defined with the `BIGINT` data type.

The table qualifier must specify owner if table.column is
not a unique column identifier within its database, or if the database is ANSI-compliant and any
user of your application is not the owner of table.

## Related links

**Related concepts**  

[Primitive Data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
