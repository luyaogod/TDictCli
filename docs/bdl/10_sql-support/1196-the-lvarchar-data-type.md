---
title: "The LVARCHAR data type"
source: "fgl-topics/c_fgl_odiagifx_009.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Partially supported IBM® Informix® SQL features > The LVARCHAR data type"
type: "concept"
description: "IBM® Informix® supports the LVARCHAR type as a \"large\" VARCHAR type. The LVARCHAR type was introduced to bypass the 255 bytes size limitation of the standard VARCHAR type. Starting with IDS version ..."
---

# The LVARCHAR data type

IBM®
Informix® supports the LVARCHAR type as a "large" VARCHAR
type. The LVARCHAR type was introduced to bypass the 255 bytes size limitation of the standard
VARCHAR type. Starting with IDS version 9.4, the LVARCHAR size limit is 32739 bytes. In older
versions the limit was 2048 bytes.

Genero BDL does not support the LVARCHAR type natively, but it has the VARCHAR type which can
hold up to 65535 bytes. IBM
Informix LVARCHAR values can be inserted or fetched by
using the BDL VARCHAR type.

Static SQL statements such as CREATE TABLE can include the LVARCHAR column type.

When extracting a schema with [fgldbsch](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages."),
LVARCHAR(N) columns will by default be converted to VARCHAR2(N) in the schema file. VARCHAR2 is a
Genero BDL-only pseudo type identified with the [type code 201](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.") that allows for
VARCHAR variables with a size that can be greater than 255 bytes to be defined.

## Related links

**Related concepts**  

[Primitive Data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.")
