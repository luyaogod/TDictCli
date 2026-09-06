---
title: "IBM Informix shadow columns"
source: "fgl-topics/c_fgl_DatabaseSchema_015.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > IBM® Informix® shadow columns"
type: "concept"
description: "Starting with IBM® Informix® IDS version 11.50.xC1, you can create shadow columns on tables by using DDL options such as ADD VERCOLS. Shadow columns are visible in the system catalog tables and would ..."
---

# IBM Informix shadow columns

Starting with IBM® Informix® IDS version 11.50.xC1, you can create shadow columns on tables by using DDL options
such as ADD VERCOLS.

Shadow columns are visible in the system catalog tables and would be listed in the column
descriptions of the .sch schema file. However, since shadow columns are not
part of the `SELECT *` list, it is not expected to get these columns in the
.sch schema file.

By default, the [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") tool will not extract shadow columns
from an IBM Informix database.

Use the `-sc` option to force the extraction of shadow columns:

```
fgldbsch -db test1 -sc
```
