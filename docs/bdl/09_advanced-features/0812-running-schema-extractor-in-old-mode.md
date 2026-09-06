---
title: "Running schema extractor in old mode"
source: "fgl-topics/c_fgl_DatabaseSchema_016.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options > Running schema extractor in old mode"
type: "concept"
description: "The fgldbsch program can be executed in old mode by specifying the -om option as first parameter, followed by the database source. You can pass the -c and -r options after the database source: ..."
---

# Running schema extractor in old mode

The fgldbsch program can be executed in old mode by specifying the
`-om` option as first parameter, followed by the database source. You can
pass the `-c` and `-r` options after the database source:

```
fgldbsch -om test1 -c -r
```

> **Important:**
>
> Use the fgldbsch -om mode for IBM® Informix® databases
> only.

The `-c` option is equivalent to `-cv BBBBBBBBB` in
the default mode: Columns defined with an SQL type that is not a native
Genero type will be converted to an equivalent type (see [-cv and -ct](0802-data-type-conversion-control.md) options for
more details).

If the `-r` option is specified, the schema extractor will ignore columns defined
with unsupported SQL types. Unsupported types have no equivalent FGL type to store and handle
the value, such as BLOB or CLOB for example. Understand that unlike the `-ie`
option, which skips the whole table definition, `-r` will exclude table columns
with unsupported types, but the other columns defined with supported types will be written to
the .sch file. Thus, a record declared with `DEFINE RECORD
rec LIKE table.*` (from a partial schema definition of a table)
cannot be used in a `SELECT * INTO rec.*` statement, because the number
of columns in the database table is different from the record definition.

When using the fgldbsch -om mode, fgldbsch will extract
system catalog tables (`informix.sys*`) for IBM
Informix databases.

## Related links

**Related concepts**  

[fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.")
