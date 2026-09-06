---
title: "Database concepts"
source: "fgl-topics/c_fgl_odiagpgs_010.html"
breadcrumb: "SQL support > SQL database guides > PostgreSQL > Database concepts > Database concepts"
type: "concept"
description: "Like Informix® servers, PostgreSQL can handle multiple database entities. Tables created by a user can be accessed without the owner prefix by other users as long as they have access privileges to ..."
---

# Database concepts

Like Informix® servers,
PostgreSQL can handle multiple database entities. Tables created by
a user can be accessed without the owner prefix by other users as
long as they have access privileges to these tables.

> **Tip:**
>
> If you have several Informix database entities,
> migrating from the Informix database to another database it is a good opportunity to centralize all
> tables in a single database. To avoid conflicts with table names, use a prefix when needed.

## Solution

Create a PostgreSQL database for each Informix
database.
