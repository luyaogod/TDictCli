---
title: "Connection parameters in FGLPROFILE when using Informix"
source: "fgl-topics/c_fgl_Migrate_to_200_dbi_informix.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Connection parameters in FGLPROFILE when using Informix®"
type: "concept"
description: "The dbi.database.* connection parameters defined in FGLPROFILE are used by the Informix driver"
---

# Connection parameters in FGLPROFILE when using Informix

> The dbi.database.* connection parameters defined in FGLPROFILE are used by the Informix® driver

Before version 2.00, the `dbi.database.*` connection parameters defined in
[FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.") are ignored by the Informix drivers.

Starting with version 2.00, the `dbi.database.*` connection parameters
defined in FGLPROFILE are used by the Informix
driver, as well as other database vendor drivers. For example, if you connect to the database
"stores", and you have the following entries defined, the driver tries to connect as "user1"
with password "alpha":

```
dbi.database.stores.username = "user1"
dbi.database.stores.password = "alpha"
```

You typically get SQL errors -387 or -329 when the wrong database login or the wrong
database name is used.

## Related links

**Related concepts**  

[Database connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program.")
