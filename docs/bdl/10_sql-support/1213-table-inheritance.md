---
title: "Table inheritance"
source: "fgl-topics/c_fgl_odiagifx_025.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Unsupported IBM® Informix® SQL features > Table inheritance"
type: "concept"
description: "IBM® Informix® SQL allows you to define a table hierarchy through named row types. Table inheritance allows a table to inherit the properties of the supertable in the meaning of constraints, storage ..."
---

# Table inheritance

IBM® Informix®
SQL allows you to define a table hierarchy through named row types. Table inheritance allows a
table to inherit the properties of the supertable in the meaning of constraints, storage options,
triggers. You must first create the types with CREATE ROW TYPE, then you can create the tables
with the UNDER keyword to define the hierarchy relationship.

```
CREATE ROW TYPE person_t ( name VARCHAR(50) NOT NULL,
 address VARCHAR(200), birthdate DATE )
CREATE ROW TYPE employee_t ( salary INTEGER, manager VARCHAR(50) )
CREATE TABLE person OF TYPE person_t 
CREATE TABLE employee OF TYPE employee_t UNDER person
```

A table hierarchy allows you to do SQL queries whose row scope is the supertable and its
subtables. For example, after inserting one row in the person table and another one in the
employee table, if you UPDATE the name column without a WHERE clause, it will update all
rows from both tables. To limit the set of rows affected by the statement to rows of the
supertable, you must use the ONLY keyword:

```
UPDATE ONLY(person) SET birthdate = NULL 
SELECT * FROM ONLY(person)
```

Genero BDL static SQL grammar does not include the syntax elements related to table hierarchy
management. You can however use [Dynamic
SQL](1140-dynamic-sql-management.md "Explains how to execute and manage SQL statements at runtime.") to perform such queries.

## Related links

**Related concepts**  

[Static SQL statements](1115-static-sql-statements.md "Describes static SQL statements supported in the language.")
