---
title: "SQL table definition"
source: "fgl-topics/c_fgl_odiagpgs_018.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Data dictionary > SQL table definition"
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

## PostgreSQL

PostgreSQL support primary key, unique, foreign key, default and check constraints.

## Constraint naming syntax

The constraint naming clause must be placed before the constraint specification.

The database interface does not convert constraint naming expressions when creating tables from
BDL programs. Review the database creation scripts to adapt the constraint naming clauses for
PostgreSQL.

## UNIQUE constraints

When defining a `UNIQUE` constraint, Informix allows only one row with a `NULL` value, while PostgreSQL allows by
default several rows with `NULL`, since `NULL` is a non-value, two
`NULL` values are not considered equal and consequently satisfy the
`UNIQUE` constraint.

PostgreSQL version 15 introduced the `NULLS NOT DISTINCT` clause for
`UNIQUE` constraints, that will result in the same behavior as Informix. However, the
ODI driver for PostgreSQL will not automatically append this clause for `UNIQUE`
constraints.

Consider using `NOT NULL` constraint on columns with `UNIQUE`
constraint.

## Related links

**Related concepts**  

[Data definition statements](1007-data-definition-statements.md "It is recommended to avoid use of DDL in programs.")
