---
title: "What are the supported IBM Informix SQL features?"
source: "fgl-topics/c_fgl_odiagifx_007.html"
breadcrumb: "SQL support > SQL database guides > IBM® Informix® > Fully supported IBM® Informix® SQL features > What are the supported IBM® Informix® SQL features?"
type: "concept"
description: "Genero BDL was first designed for IBM® Informix® database servers. The answer to this question is: Every SQL feature that is not listed in the other sections of this chapter. The following list gives ..."
---

# What are the supported IBM Informix SQL features?

Genero BDL was first designed for IBM® Informix® database servers. The answer
to this question is: Every SQL feature that is not listed in the other
sections of this chapter.

The following list gives an idea of the IBM Informix SQL
elements you can use with Genero BDL:

- Database connection control instructions (DATABASE, CONNECT). See [Connections](1059-database-connections.md "Explains how to manage database connections in a program."), with DB user
  authentication.
- Transaction control instructions and concurrency settings (BEGIN
  WORK, SET ISOLATION). See [Transactions](1106-database-transactions.md "Database transaction concepts and handling.").
- Basic, portable data types (lNT, BIGINT, DECIMAL, CHAR, VARCHAR,
  DATE, DATETIME, TEXT, BYTE, etc). See [data types](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.").
- SERIAL, BIGSERIAL with last generated serial in sqlca.sqlerrd[2] after INSERT.
- Common Data Definition Language statements (CREATE TABLE, DROP
  TABLE, etc). See [Static SQL](1115-static-sql-statements.md "Describes static SQL statements supported in the language.").
- Common Data Manipulation Language statements (SELECT, INSERT,
  UPDATE, DELETE, etc). See [Static SQL](1115-static-sql-statements.md "Describes static SQL statements supported in the language.").
- Cursors declared with [SELECT
  ... FOR UPDATE](1158-declare-select-for-update.md "Associate a database cursor with a SELECT statement to perform positioned updates and deletes"), with or without the WITH HOLD option.
- Result set handling with cursors (DECLARE / OPEN / FETCH / CLOSE
  / FREE). See [Result Sets](1148-result-set-processing.md "Shows how to fetch rows from a database query.").
- Positioned UPDATEs and DELETEs (UPDATE/DELETE WHERE CURRENT OF).
  See [Positioned
  Updates](1156-positioned-updates-deletes.md "Describes row modification based on a FOR UPDATE cursor.").
- Cursors to insert rows (DECLARE / OPEN / PUT / FLUSH). See [Insert
  Cursors](1163-sql-insert-cursors.md "Explains how to insert a log of rows into a table efficiently.").
- Stored procedure calls. See [SQL
  Programming](0985-sql-basics.md "This section contains fundamental information to know about SQL programming with Genero BDL.").
- SQL statement interruption. See [Using SQL interruption](0993-using-sql-interruption.md "Interrupt long running SQL queries, or interrupt queries waiting for locked data.").
- SQL execution status and error messages (sqlca, SQLSTATE). See [Connections](1059-database-connections.md "Explains how to manage database connections in a program.").
- Global Language Support with single and multibyte character sets
  for CHAR/ VARCHAR data storage. See [Localization](../09_advanced-features/0863-localization.md "Localization support allows you to implement programs that follow specific language and cultural rules.").
- LOAD and UNLOAD utility statements. See [I/O SQL instructions](1175-sql-load-and-unload.md "Describes the instructions to export/import information from/to a database.").
- Database schema extraction to define program variables LIKE database
  columns. See [Database
  Schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.").
