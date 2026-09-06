---
title: "SQL table definition"
source: "fgl-topics/c_fgl_odiagmsv_018.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Data dictionary > SQL table definition"
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

## Microsoft™ SQL Server

Microsoft SQL Server supports primary key, unique,
foreign key, default and check constraints.

## Constraint naming syntax

The constraint naming clause must be placed before the constraint specification.

The database interface does not convert constraint naming expressions when creating tables from
BDL programs. Review the database creation scripts to adapt the constraint naming clauses for Microsoft SQL Server.

## The NULL / NOT NULL constraint

Microsoft SQL Server creates columns as `NOT
NULL` by default, when no `NULL` constraint is specified (`colname
datatype {NULL | NOT NULL}`). A special option is provided to invert this behavior:
`ANSI_NULL_DFLT_ON`. This option can be enabled with the `SET`
command, or in the database options of SQL Server Management Studio.

Before using a database, you must check the "ANSI NULL Default" option in the database properties
if you want to have the same default NULL constraint as when using Informix.

## Related links

**Related concepts**  

[Data definition statements](1007-data-definition-statements.md "It is recommended to avoid use of DDL in programs.")
