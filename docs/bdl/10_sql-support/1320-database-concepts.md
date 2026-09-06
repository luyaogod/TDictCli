---
title: "Database concepts"
source: "fgl-topics/c_fgl_odiagmys_011.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB > Database concepts > Database concepts"
type: "concept"
description: "Like Informix® servers, Oracle® MySQL can handle multiple database entities. Tables created by a user can be accessed without the owner prefix by other users as long as they have access privileges to ..."
---

# Database concepts

Like Informix® servers, Oracle® MySQL can handle multiple database entities.
Tables created by a user can be accessed without the owner prefix by other users as long
as they have access privileges to these tables.

> **Tip:**
>
> If you have several Informix database entities,
> migrating from the Informix database to another database it is a good opportunity to centralize all
> tables in a single database. To avoid conflicts with table names, use a prefix when needed.

## Solution

Create a MySQL database for each Informix
database.
