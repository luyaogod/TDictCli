---
title: "Database concepts"
source: "fgl-topics/c_fgl_odiagmsv_012.html"
breadcrumb: "SQL support > SQL database guides > Microsoft™ SQL Server > Database concepts > Database concepts"
type: "concept"
description: "As in Informix®, an SQL SERVER engine can manage multiple database entities. When creating a database object like a table, Microsoft™ SQL SERVER allows you to use the same object name in different ..."
---

# Database concepts

As in Informix®, an SQL
SERVER engine can manage multiple database entities. When creating
a database object like a table, Microsoft™ SQL
SERVER allows you to use the same object name in different databases.

> **Tip:**
>
> If you have several Informix database entities,
> migrating from the Informix database to another database it is a good opportunity to centralize all
> tables in a single database. To avoid conflicts with table names, use a prefix when needed.

Microsoft SQL SERVER Transac-SQL allows to change the
database context with the `USE database_name` command. For
example, a program can connect to the `"master"` database, and then switch to another
database entity. The Genero BDL ODI driver for SQL SERVER is not prepared for this, especially if
the new database is using a different collation.
