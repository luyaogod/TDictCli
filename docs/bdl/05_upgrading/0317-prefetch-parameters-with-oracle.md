---
title: "Prefetch parameters with Oracle"
source: "fgl-topics/c_fgl_Migrate_to_200_oracle_prefetch.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Prefetch parameters with Oracle"
type: "concept"
description: "Prefetch parameters allow an application to automatically fetch rows from the Oracle database when opening a cursor."
---

# Prefetch parameters with Oracle

> Prefetch parameters allow an application to automatically fetch rows from the Oracle® database when opening a cursor.

Before version 2.00, the default prefetch parameters are 50 rows and 65535 bytes for the prefetch
buffer. Some customers experienced a huge memory usage with those default values, when
using a lot of cursors: It appears that the Oracle client is allocating a buffer of prefetch.memory (64
Kbytes) for each cursor.

Starting with version 2.00, the default is 10 rows and 0 (zero) bytes for the prefetch buffer
(memory), meaning that memory is not included in computing the number of rows to
prefetch.

## Related links

**Related concepts**  

[Database connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program.")
