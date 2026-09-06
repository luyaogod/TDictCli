---
title: "RECORD"
source: "fgl-topics/c_fgl_records_002.html"
breadcrumb: "Language basics > Records > RECORD"
type: "concept"
---

# RECORD

> The RECORD keyword defines a structured type or variable.

## Syntax 1 (explicit record definition)

```
RECORD [ attributes-list ]
  member type-specification
  [,...] 
END RECORD
```

1. member is an identifier for a record field, that must follow the convention
   for [identifiers](0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity.").
2. type-specification can be one of:
   - A [primitive type](0690-primitive-type-specification.md "Type definitions using a primitive data type define a primitive type.")
   - A [record definition](0717-record.md "The RECORD keyword defines a structured type or variable.")
   - An [array definition](0731-array.md "An array defines a vector variable with a list of elements.")
   - A [dictionary definition](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")
   - A [function type definition](0770-function-references.md "Function can be referenced and invoked dynamically in a CALL instruction, or in an expression.")
   - The name of a [user defined type](0753-type.md "Types define a synonym for a base or structured data type.")
   - The name of a [built-in class](../15_library-reference/2908-built-in-packages.md "These topics cover the built-in classes provided by the Genero Business Development Language.")
   - The name of an [imported extension
     class](../15_library-reference/3480-extension-packages.md "Several utility classes and functions are provided in additional packages.")
   - The name of an [imported Java class](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.")
3. attributes-list is a comma-separated list of
   name = value pairs or name attributes, and defines [attributes for the record type](0718-attributes-on-record-definitions.md "Records can be defined with attributes, to complete the type description.").

## Syntax 2 (database column based record)

```
RECORD [ attributes-list ] LIKE [dbname:]tabname.*
```

1. dbname identifies a specific [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").
2. tabname.\* references the structure of a complete table defined in the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").
3. attributes-list is a comma-separated list of
   name = value pairs or name attributes, and defines [attributes for the record type](0718-attributes-on-record-definitions.md "Records can be defined with attributes, to complete the type description.").

## Usage

A record defines an ordered set of variables called members. Each record member is defined with a
specific type or in turn, structured type.

> **Tip:**
>
> Consider defining a [user type](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") for
> records, to avoid repeating the record definition for each variable.

Records whose members correspond in number, order, and data type compatibility to a database
table can be useful for transferring data from the database to the screen, to reports, or to
functions.

In the first form (Syntax 1), record members are defined
explicitly:

```
DEFINE rec RECORD
           cust_id INT,
           cust_name VARCHAR(50),
           cust_address VARCHAR(100),
           ...
       END RECORD
```

In the second form (Syntax 2), record members are created implicitly from the table definition
found in the database schema file specified by the [`SCHEMA`](../09_advanced-features/0793-schema.md "Defines the database schema files to be used for compilation.")
instruction:

```
SCHEMA stock
...
DEFINE rec RECORD LIKE customer.*
```

> **Important:**
>
> When using the `LIKE` clause, the data types are taken from the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") during compilation. Make sure that
> the database schema file of the development database corresponds to the production database. If
> these differ, the records defined in the your programs will not match the SQL table structures of
> the production database.
>
> For example, with an SQL statement such as `SELECT * INTO record.* FROM table` ,
> the columns of the SQL table represented by the `*` star in the select list will no
> longer match the list of fields represented by `record.*`. This can lead to missing
> values if columns are been removed from the SQL table, or result in decorrelated values if SQL
> columns have been re-ordered in the SQL table. When the SQL column type does not match the target
> variable type, conversions error will occur and the `SELECT` will fail.
>
> To make sure that the SQL columns match the target record variables, list all SQL columns
> explicitly in the `SELECT` statement; as in `SELECT cust_name, cust_addr,
> cust_state INTO record.* FROM ...`
