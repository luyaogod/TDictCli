---
title: "SQL Server drivers performance"
source: "fgl-topics/c_fgl_Migrate_to_320_sqlserver_perfs.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > SQL Server drivers performance"
type: "concept"
---

# SQL Server drivers performance

> SQL Server ODI drivers based on FreeTDS, Easysoft and MS ODBC have been reviewed to achieve better execution times.

Starting with Genero BDL 3.20, the FTM, ESM and SNC drivers for Microsoft® SQL Server use now the ODBC API
`SQLExecDirect()` for any sort of SQL statement. In prior versions, the drivers use
`SQLPrepare()` + `SQLExecute()` for common SQL statements such as
`SELECT`, `INSERT`, `UPDATE`,
`DELETE`.

Other improvements have been done, such as setting the ODBC cursor name
(`SQLSetCursorName()`) only when needed, namely for `SELECT`
statements with `FOR UDPATE` clause.

This leads to a reduction in the number of `sp_cursor*` stored procedure calls and
shows a noticeable improvement in performance, especially in a client/server configuration where
each `sp_cursor*` call results in a network round trip.

## Related links

**Related concepts**  

[Performances with SQL interruption](0168-performances-with-sql-interruption.md "With some database drivers, performances can be impacted when using OPTIONS SQL INTERRUPT ON.")
