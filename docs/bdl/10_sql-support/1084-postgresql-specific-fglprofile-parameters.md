---
title: "PostgreSQL specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_pgs.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > PostgreSQL specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .pgs.prefetch.rows Maximum number of rows to be pre-fetched. dbi.database.stores.pgs.prefetch.rows = 100 Use this parameter to increase performance by defining the maximum number ..."
---

# PostgreSQL specific FGLPROFILE parameters

## `dbi.database.dsname.pgs.prefetch.rows`

Maximum number of rows to be pre-fetched.

```
dbi.database.stores.pgs.prefetch.rows = 100
```

Use this parameter to increase performance by defining the maximum number
of rows to be fetched into the db client buffer. However, the bigger this
parameter is, the more memory is used by each program. This parameter
applies to all cursors in the program.

The default is 50 rows, which is a good compromise for regular interactive OTLP applications.

When using server-side cursors (`DECLARE CURSOR`), this parameter will define the
number of rows fetched with the PostgreSQL `FETCH cursor FORWARD
nbrows` command. This parameter has no effect on static
`SELECT` statements, nor on SELECT statements executed with `PREPARE`
+ `EXECUTE INTO`.

## `dbi.database.dsname.pgs.schema`

Defines the PostgreSQL schema search path after connection is established.

```
dbi.database.stores.pgs.schema = "\"$user\",public,stock"
```

The value specified in this parameter will be used to execute the `SET search_path TO
fglprofile-value` SQL instruction.

See also [Name resolution of SQL objects](1438-name-resolution-of-sql-objects.md).

## `dbi.database.dsname.pgs.trace.file`

To enable tracing of client/server communication with PostgreSQL, use the following FGLPROFILE
entry to define the path to a trace
file:

```
dbi.database.dsname.pgs.trace.file = "/tmp/pgstrace_%(PID).log"
```

The file path can contain the `%(PID)` placeholded, that will be replaced by the
process id.

For more details about the trace file content, refere to the PostgreSQL [`PQtrace()`](https://www.postgresql.org/docs/current/libpq-control.html) API.

## `dbi.database.dsname.pgs.pgvector`

Enable pgvector extension support with the following
setting:

```
dbi.database.dsname.pgs.pgvector = true
```

For more details about pgvector support, read [PostgreSQL pgvector extension data types](1441-postgresql-pgvector-extension-data-types.md).
