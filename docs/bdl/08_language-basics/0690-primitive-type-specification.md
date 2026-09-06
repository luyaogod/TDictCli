---
title: "Primitive type specification"
source: "fgl-topics/c_fgl_variables_primitive_type.html"
breadcrumb: "Language basics > Variables > Primitive type specification"
type: "concept"
---

# Primitive type specification

> Type definitions using a primitive data type define a primitive type.

## Syntax

A primitive-type-specification
is:

```
{ BOOLEAN
| CHAR(size)
| VARCHAR(size)
| STRING
| BIGINT
| INTEGER
| SMALLINT
| TINYINT
| DECIMAL(prec [,scale])
| MONEY(prec [,scale])
| FLOAT | SMALLFLOAT
| DATE
| DATETIME qual1 TO qual2
| INTERVAL qual1 TO qual2 
| TEXT
| BYTE
| LIKE [dbname:]tabname.colname
}
   [ attributes-list ]
```

1. The type name can be one of the [primitive data types](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").
2. The type can also reference a table column in the [database schema file](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").
3. dbname identifies a specific database schema file.
4. tabname.colname references a column defined in the
   database schema file.
5. attributes-list is a comma-separated list of name = value
   pairs or name attributes, and defines [attributes for the element of this primitive
   type](0692-attributes-on-variable-definitions.md "Variables can be defined with meta-data information.").

## Usage

Primitive types can be used in:

- [Variable definitions](0688-define.md "The DEFINE instruction declares a program variable with a given type.")
- [Record definitions](0717-record.md "The RECORD keyword defines a structured type or variable.")
- [User-type definitions](0753-type.md "Types define a synonym for a base or structured data type.")
- [Array definitions](0731-array.md "An array defines a vector variable with a list of elements.")
- [Dictionary definitions](0744-dictionary.md "A dictionary defines an associative array (hash-map) of elements.")

## Example

```
DEFINE cnt INTEGER
DEFINE filename STRING
TYPE t_cust RECORD
        cust_id LIKE customer.cust_id,
        cust_id LIKE customer.cust_name
    END RECORD
```

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
