---
title: "SQL table definition"
source: "fgl-topics/c_fgl_odiagsqt_029.html"
breadcrumb: "SQL support > SQL database guides > SQLite > Data dictionary > SQL table definition"
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

## SQLite

SQLite supports primary key, unique, foreign key, default and check constraints.

## Constraint naming syntax

The constraint naming clause must be placed before the constraint specification.

The database interface does not convert constraint naming expressions when creating tables from
BDL programs. Review the database creation scripts to adapt the constraint naming clauses for
SQLite.

## Primary keys

Like Informix, SQLite creates an index to enforce
`PRIMARY KEY` constraints (some RDBMS do not create indexes for constraints). Using
`CREATE UNIQUE INDEX` to define unique constraints is obsolete (use primary keys or a
secondary key instead).

## Unique constraints

Like Informix, SQLite creates an index to enforce
`UNIQUE` constraints (some RDBMS do not create indexes for constraints).

When using a unique constraint, Informix allows only
one row with a `NULL` value, while SQLite allows several rows with
`NULL`! Consider using `NOT NULL` constraint with
`UNIQUE`.

## Foreign keys

SQLite (3.6.19 and +) implements foreign key support, but this feature is not enabled by default.
In fact, it is possible to define foreign keys on tables, but when doing database operations, the
constraints are not enforced until you enable it explicitly with a `PRAGMA`
command.

To get foreign key constraint checking in SQLite, perform a `PRAGMA` command with
the `EXECUTE IMMEDIATE`
instruction:

```
EXECUTE IMMEDIATE "PRAGMA foreign_keys = ON"
```

Future releases of SQLite might change this, so that foreign key constraints enabled by
default.

## Check constraints

The check condition may be any valid expression that can be evaluated to `TRUE` or
`FALSE`, including functions and literals. You must verify that the expression is not
Informix specific. SQLite supportes
`CHECK` constraints like Informix, but you must check the syntax of the expression to
make it portable.

## Null constraints

Informix and
SQLite support not null constraints, but Informix does
not allow you to give a name to `NOT NULL`" constraints.

## Related links

**Related concepts**  

[Data definition statements](1007-data-definition-statements.md "It is recommended to avoid use of DDL in programs.")
