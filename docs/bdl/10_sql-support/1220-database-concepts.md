---
title: "Database concepts"
source: "fgl-topics/c_fgl_odiagdmg_010.html"
breadcrumb: "SQL support > SQL database guides > Dameng® database server > Database concepts > Database concepts"
type: "concept"
description: "Unlike Informix®, a Dameng® database server can handle only one database entity. Informix servers have an ID (INFORMIXSERVER) and databases are identified by name. Dameng instances/databases are ..."
---

# Database concepts

Unlike Informix®, a Dameng® database server can handle only one database entity. Informix servers have an ID
(INFORMIXSERVER) and databases are identified by name. Dameng instances/databases are identified by
the name defined in the dm\_svc.conf configuration file. See Dameng
documentation for more details.

> **Tip:**
>
> If you have several Informix database entities,
> migrating from the Informix database to another database it is a good opportunity to centralize all
> tables in a single database. To avoid conflicts with table names, use a prefix when needed.
