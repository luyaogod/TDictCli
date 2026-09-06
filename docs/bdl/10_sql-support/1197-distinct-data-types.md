---
title: "DISTINCT data types"
source: "fgl-topics/c_fgl_odiagifx_010.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > DISTINCT data types"
type: "concept"
description: "IBM® Informix® supports DISTINCT data types as User Defined Types based on a source data type, but with different casts and functions than those on the source data type. Genero BDL partially supports ..."
---

# DISTINCT data types

IBM®
Informix® supports DISTINCT data types as User Defined
Types based on a source data type, but with different casts and functions than those on the source
data type.

Genero BDL partially supports the IBM
Informix DISTINCT data types:

The [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.") schema extractor can
extract columns defined with a distinct type and write the distinct type code in the .sch schema
file. For more details, see the list of distinct types in the [Column Definition File (.sch)](../09_advanced-features/0795-column-definition-file-sch.md "The .sch database schema file contains the data types of database table columns.")

However, there are some restrictions you must be aware of:

- It is not possible to define BDL variables explicitly with the name of a distinct type.
  Variables must be defined indirectly with the schema by using the DEFINE LIKE statement.
- The static SQL syntax does not support OPAQUE-related syntax elements:
  - The DDL statements CREATE DISTINCT TYPE, DROP TYPE, CREATE CAST, and DROP CAST are not
    allowed,
  - In CREATE TABLE / ALTER TABLE DDL statements, the data type must be a built-in type.
  - The :: cast operator is not supported.
