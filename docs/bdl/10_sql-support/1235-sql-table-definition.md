---
title: "SQL table definition"
source: "fgl-topics/c_fgl_odiagdmg_016.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Data dictionary > SQL table definition"
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

## Dameng®

Dameng supports primary key, unique, foreign key,
default and check constraints.

## Constraint naming

The constraint naming clause must be placed before the constraint
specification:

```
CREATE TABLE customer ( cust_id INTEGER CONSTRAINT pk1 PRIMARY KEY, ... )
```

The database interface does not convert constraint naming expressions when creating tables from
BDL programs. Review the database creation scripts to adapt the constraint naming clauses for
Dameng.

## Related links

**Related concepts**  

[Data definition statements](1007-data-definition-statements.md "It is recommended to avoid use of DDL in programs.")
