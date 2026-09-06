---
title: "Database driver internal error changed from -768 to -6319"
source: "fgl-topics/c_fgl_Migrate_to_220_db_driver_error.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Database driver internal error changed from -768 to -6319"
type: "concept"
description: "The internal error raised was changed to avoid conflicts with an IBM Informix SQL error code."
---

# Database driver internal error changed from -768 to -6319

> The internal error raised was changed to avoid conflicts with an IBM® Informix® SQL error code.

Prior to version 2.20, if an unexpected error occurred in a database driver, the driver could
return error -768, which is a real IBM
Informix SQL error that instructs the user to
call the IBM support center.

To avoid any mistake, 2.20 database drivers return now the error [-6319](../15_library-reference/4483-genero-bdl-errors.md "System error messages sorted by error number.") if an internal error
occurs, which is a Genero Business Development Language specific error message that
suggests you to set the FGLSQLDEBUG environment variable to get detailed debug
messages.

## Related links

**Related concepts**  

[FGLSQLDEBUG](../07_configuration/0534-fglsqldebug.md "Defines the debug level for tracing SQL instructions.")
