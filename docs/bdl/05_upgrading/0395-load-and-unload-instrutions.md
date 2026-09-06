---
title: "LOAD and UNLOAD instrutions"
source: "fgl-topics/c_fgl_MigI4GL_058.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > LOAD and UNLOAD instrutions"
type: "concept"
---

# LOAD and UNLOAD instrutions

> This topic describes differences between I4GL and FGL with the LOAD and UNLOAD instructions.

## LOAD requires to escape newline characters for TEXT

With IBM® Informix® 4GL, when using the [`LOAD`](../10_sql-support/1176-load.md "Inserts data from a file into an existing database table.") instruction, if a file contains [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") data, newline characters do not need to be preceeded by a backslash
(to escape the newline), while `CHAR` and `VARCHAR` data requires a
backslash to escape newline.

With Genero BDL, if a newline character is not preceded by a backslash, you get an error during
the `LOAD` instruction. This behavior is more consistent than I4GL.

Instead of using I4GL to export data from your Informix database with `UNLOAD`,
use Genero BDL to procude a correct data in the file.

## Related links

**Related concepts**  

[SQL LOAD and UNLOAD](../10_sql-support/1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.")
