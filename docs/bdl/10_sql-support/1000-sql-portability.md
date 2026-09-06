---
title: "SQL portability"
source: "fgl-topics/c_fgl_sql_programming_056.html"
breadcrumb: "SQL support > SQL programming > SQL portability"
type: "concept"
---

# SQL portability

> Writing portable SQL is mandatory, to support different kind of database servers.

This section provides hints to solve SQL incompatibility problems in your programs.

In addition to this SQL portability guide, read carefully the [database-specific guides](1178-sql-database-guides.md "This section includes the SQL guides for various supported database servers.") which contain
database specific information about SQL compatibility issues.

To easily detect SQL statements with specific syntax, you can use the `-W stdsql`
option of fglcomp:

```
$ fglcomp -W stdsql orders.4gl 
module.4gl:15: SQL Statement or language instruction with specific SQL syntax.
```

This compiler option can only detect non-portable SQL syntax in static SQL statements.

## Child topics

- [Database entities](1001-database-entities.md): The database entity concept across different database engines.
- [Database users and security](1002-database-users-and-security.md): Properly identifying database users allows to use database security and audit features.
- [Creating a database from programs](1003-creating-a-database-from-programs.md): Creating a database from within a program requires special consideration.
- [Handling nested transactions](1004-handling-nested-transactions.md): You can manage nested transactions in different parts of a program.
- [Transaction blocks across connections](1005-transaction-blocks-across-connections.md): Transaction blocks manage transactions when connected to several database servers.
- [Transaction savepoints](1006-transaction-savepoints.md): Using transaction savepoints with different database engines.
- [Data definition statements](1007-data-definition-statements.md): It is recommended to avoid use of DDL in programs.
- [Using portable data types](1008-using-portable-data-types.md): Only a limited set of data types are really portable across several database engines.
- [The NULL value](1009-the-null-value.md): Several considerations need to be taken regarding SQL columns allowing NULL values.
- [The BOOLEAN data type](1010-the-boolean-data-type.md): SQL implementation of the BOOLEAN data type varies on the database type.
- [Data manipulation statements](1011-data-manipulation-statements.md): Make sure that SQL statement syntaxes are supported by all target database engines.
- [CHAR and VARCHAR types](1012-char-and-varchar-types.md): Using the CHAR and VARCHAR data types with different databases.
- [INTERVAL data types](1019-interval-data-types.md): Not all database brands support a native SQL type to store time duration.
- [The UNITS operator](1020-the-units-operator.md): The UNITS SQL operator is specific to Informix SQL and needs to be considered.
- [Scrollable cursors](1021-scrollable-cursors.md): How scrollable cursors can be supported on different databases.
- [Optimistic locking](1022-optimistic-locking.md): Implementing optimistic locking to handle access concurrently to the same database records.
- [Auto-incremented columns (serials)](1023-auto-incremented-columns-serials.md): How to implement automatic record keys.
- [IBM Informix SQL ANSI Mode](1027-ibm-informix-sql-ansi-mode.md): Understanding the impact of the SQL ANSI mode of IBM® Informix®.
- [Positioned UPDATE/DELETE](1028-positioned-update-delete.md): Using positioned updates/deletes with named database cursors.
- [Cursors WITH HOLD](1029-cursors-with-hold.md): Programming WITH HOLD cursors using SELECT with and without FOR UPDATE clause.
- [Insert cursors](1030-insert-cursors.md): Using insert cursors with non-Informix databases.
- [String literals in SQL statements](1031-string-literals-in-sql-statements.md): Single quotes is the standard for delimiting string literals in SQL.
- [String concatenation operators in SQL](1032-string-concatenation-operators-in-sql.md): The || operator is the standard to concatenate strings.
- [Date/time literals in SQL statements](1033-date-time-literals-in-sql-statements.md): Good practices for date and time handling in SQL.
- [Naming database objects](1034-naming-database-objects.md)
- [LOAD and UNLOAD instructions](1038-load-and-unload-instructions.md): The LOAD and UNLOAD instructions can produce different data formats depending on the database server type.
- [Temporary tables](1039-temporary-tables.md): Syntax for temporary table creation is not unique across all database engines.
- [Outer joins](1040-outer-joins.md): Use standard ISO outer join syntax instead of the old IBM® Informix® OUTER() syntax.
- [Substring expressions](1041-substring-expressions.md): Handle substrings expressions with different database engines.
- [Using ROWID columns](1042-using-rowid-columns.md): Automatic ROWID columns is not a common database feature.
- [MATCHES and LIKE operators](1043-matches-and-like-operators.md): Use the standard LIKE operator instead of the MATCHES operator.
- [GROUP BY clause](1044-group-by-clause.md)
- [The LENGTH() function in SQL](1045-the-length-function-in-sql.md): The semantics of the LENGTH() SQL function differs according to the database engine.
- [Large OBjects (LOBs) data types](1046-large-objects-lobs-data-types.md): Use TEXT and BYTE FGL types to store database character and binary large objects.
- [Stored procedures](1047-stored-procedures.md): Executing stored procedures with different database engine types.
- [Row limiting clause (SELECT)](1051-row-limiting-clause-select.md): How to use the right clause to limit the number of rows produced by a SELECT statement?
