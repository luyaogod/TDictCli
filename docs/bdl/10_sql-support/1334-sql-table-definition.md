---
title: "SQL table definition"
source: "fgl-topics/c_fgl_odiagmys_017.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Data dictionary > SQL table definition"
type: "concept"
description: "Informix® Informix supports primary key, unique, foreign key, default and check constraints. The constraint naming syntax is different in Informix and most other databases: Informix expects the ..."
---

# SQL table definition

## Informix®

Informix supports primary key, unique, foreign key,
default and check constraints.

The constraint naming syntax is different in Informix
and most other databases: Informix expects the constraint name after the constraint
definition:

```
CREATE TABLE emp (
  ...
  emp_code CHAR(10) UNIQUE CONSTRAINT pk_emp,
  ...
)
```

While other SQL database brands require to specify the constraint name
before the constraint definition:

```
CREATE TABLE emp (
   ... 
   emp_code CHAR(10) CONSTRAINT pk_emp UNIQUE, 
   ...
)
```

## Oracle® MySQL and MariaDB

MySQL and MariadDB support primary key, unique, foreign key and default constraints.

> **Important:**
>
> MySQL and MariaDB do not support `CHECK` constaints. In fact,
> the syntax is allowed but the constraint is ignored.

## Constraint naming syntax

The constraint naming clause must be placed before the constraint specification.

The database interface does not convert constraint naming expressions when creating tables from
BDL programs. Review the database creation scripts to adapt the constraint naming clauses for
MySQL.

## Primary keys

MySQL creates an index to enforce `PRIMARY KEY` constraints (some RDBMS do not
create indexes for constraints). Using `CREATE UNIQUE INDEX` to define unique
constraints is obsolete (use primary keys or a secondary key instead).

In MySQL, the name of a `PRIMARY KEY` is `PRIMARY`.

## Unique constraints

Like Informix, MySQL creates an index to enforce
`UNIQUE` constraints (some RDBMS do not create indexes for constraints).

When using a unique constraint, Informix allows only
one row with a `NULL` value, while MySQL allows several rows with
`NULL`! Using `CREATE UNIQUE INDEX` is obsolete.

## Foreign keys

Both Informix and MySQL support the `ON DELETE
CASCADE` option. In MySQL, foreign key constraints are checked immediately, so `NO
ACTION` and `RESTRICT` are the same.

## Check constraints

Check constraints are not yet supported in MySQL.

If your application tables use `CHECK` constaints, you need to implement these
constraints with triggers.

## Related links

**Related concepts**  

[Data definition statements](1007-data-definition-statements.md "It is recommended to avoid use of DDL in programs.")
