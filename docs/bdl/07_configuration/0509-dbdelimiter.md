---
title: "DBDELIMITER"
source: "fgl-topics/c_fgl_EnvVariables_DBDELIMITER.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBDELIMITER"
type: "concept"
---

# DBDELIMITER

> Defines the value separator for unload data files.

The DBDELIMITER environment variable defines the character to delimit data field values for [`LOAD`](../10_sql-support/1176-load.md "Inserts data from a file into an existing database table.")/[`UNLOAD`](../10_sql-support/1177-unload.md "Copies data from the database tables into a file.") instructions, and for the for the [`base.Channel`](../15_library-reference/2996-base-channel-setdelimiter.md "Define the value delimiter for a channel.") in/out API.

If DBDELIMITER is not defined, the default delimiter is a (`|`) pipe.

Only one single character can be specified. When setting the DBDELIMITER environment variable
with a string, only the first character is used.

Do not use backslash (`\`), or hex digits (`0-9`,
`A-F`, `a-f`). When using such character, the
`LOAD`/`UNLOAD` instructions will raise error [-10099](../15_library-reference/4483-genero-bdl-errors.md).

Additional formatting options can be specified with the `DELIMITER` option of
`LOAD`/`UNLOAD` and the [`base.Channel.setDelimiter()`](../15_library-reference/2996-base-channel-setdelimiter.md "Define the value delimiter for a channel.")
method.
