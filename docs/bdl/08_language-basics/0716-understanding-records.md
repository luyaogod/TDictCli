---
title: "Understanding records"
source: "fgl-topics/c_fgl_records_intro.html"
breadcrumb: "Language basics > Records > Understanding records"
type: "concept"
---

# Understanding records

> This is an introduction to records.

A record defines a structured [variable](0686-variables.md "Explains how to define program variables."), where each
member can be defined with a specific [data type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").

Records are used in [interactive instructions](../11_user-interface/1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation.") like
`INPUT` to control forms, and record are also used in `INSERT` and
`UPDATE` [SQL instructions](../10_sql-support/0983-sql-support.md "These topics cover SQL support in the Genero Business Development Language."), to
update the database table.

Records can contain sub-record structures, and list types like arrays and dictionaries:

```
DEFINE reader RECORD
            id INTEGER,
            name VARCHAR(100),
            birth DATE,
            address RECORD
                num VARCHAR(20),
                street VARCHAR(200),
                city_id INTEGER,
                state_id VARCHAR(5)
            END RECORD,
            book_ids DYNAMIC ARRAY OF INTEGER
       END RECORD
```

Records are typically used to store the values of a database row. Records can be based on the
column types of a database table, which is defined in a database schema
(`SCHEMA`):

```
SCHEMA stores
DEFINE cust RECORD customer.*
-- cust is defined with the column of the customer table
```

The following list summarizes record usage:

- Record variables are defined with the [`RECORD`](0717-record.md "The RECORD keyword defines a structured type or variable.") syntax block, or with a [`TYPE`](0751-types.md "Types can be defined by the programmer to centralize the definition of complex/structured variables.") defined as `RECORD` .
- Records can be defined with [attributes](0718-attributes-on-record-definitions.md "Records can be defined with attributes, to complete the type description.").
- Records can be initialized in their definitions with [record
  initializers](0719-record-initializers.md "Records can be initialized in their definition.").
- To access a record member, use the [dot notation](0720-accessing-record-members.md "Record members are accessed with the dot notation.")
  (`record-name.member-name`)
- Records can be copied to other records, if they are defined with the same type, by using the
  [`LET rec1 = rec2` assignment
  operation](0721-copying-records.md "Records can be assigned to each other with the = operator.").
- It is possible to compare all members of records of the same type with the [`rec1.* == rec2.*` comparison expression](0722-comparing-records.md "Records can be compared with the == comparison operator and the .* notation.").
- Records can be passed as [function parameters](0723-records-and-functions.md "Records can be passed as function parameters, and can be returned from functions.") by value
  or by reference (`INOUT`).
