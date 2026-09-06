---
title: "SQL table definition"
source: "fgl-topics/c_fgl_odiagora_019.html"
breadcrumb: "SQL support > SQL database guides > Oracle® Database > Data dictionary > SQL table definition"
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

## ORACLE

ORACLE supports primary key, unique, foreign key, default and check constraints.

## Constraint naming syntax

The constraint naming clause must be placed
before the constraint specification.

The database interface does not convert constraint
naming expressions when creating tables from BDL programs. Review the database creation scripts to
adapt the constraint naming clauses for ORACLE.

## Primary keys

Like Informix, ORACLE creates an index to enforce
`PRIMARY KEY` constraints (some RDBMS do not create indexes for constraints). Using
`CREATE UNIQUE INDEX` to define unique constraints is obsolete (use primary keys or a
secondary key instead).

## Unique constraints

Like Informix, ORACLE creates an index to enforce
`UNIQUE` constraints (some RDBMS do not create indexes for constraints).

When using a unique constraint, Informix allows only
one row with a `NULL` value, while ORACLE allows several rows with
`NULL`! Consider using `NOT NULL` constraint with
`UNIQUE`.

## Foreign keys

Both Informix and
ORACLE support the `ON DELETE CASCADE` option. To defer constraint checking, Informix provides the `SET CONSTRAINT` command
while ORACLE provides the `ENABLE` and `DISABLE` clauses.

## Check constraints

The check condition may be any valid expression that can be evaluated to `TRUE` or
`FALSE`, including functions and literals. You must verify that the expression is not
Informix specific.

## Null constraints

Informix and
ORACLE support `NOT NULL` constraints, but Informix does not allow you to give a name to `NOT NULL`"
constraints.

## Related links

**Related concepts**  

[Data definition statements](1007-data-definition-statements.md "It is recommended to avoid use of DDL in programs.")
