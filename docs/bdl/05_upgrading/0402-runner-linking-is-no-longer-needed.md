---
title: "Runner linking is no longer needed"
source: "fgl-topics/c_fgl_Mig0000_004.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > Installation and setup topics > Runner linking is no longer needed"
type: "concept"
---

# Runner linking is no longer needed

> With Four Js Business Development Suite (BDS), you need to create the fglrun binary with the fglmkrun tool, by specifying the type of the database driver and C extensions libraries to be linked with the runtime system. Since Genero Business Development Language version 2.00, you no longer need need to link the runtime system.

The database drivers are provided as shared libraries ready to use; you just need to
specify the driver to be loaded.

However, C extensions must be provided shared libraries for Genero BDL. To easy migration, the
runtime system loads automatically the `userextension` share library (or
DLL).

## Related links

**Related concepts**  

[Database connections](../10_sql-support/1059-database-connections.md "Explains how to manage database connections in a program.")

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
