---
title: "BDL 3.20 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_320.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide"
type: "concept"
---

# BDL 3.20 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 3.20.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 3.20 new features](0062-bdl-3-20-new-features.md "Features added in 3.20 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 3.10 upgrade guide](0178-bdl-3-10-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.10.").

## Child topics

- [Web Services changes](0158-web-services-changes.md): There are changes in support of web services in Genero 3.20.
- [Genero Mobile for Android (GMA) 1.40 changes](0159-genero-mobile-for-android-gma-1-40-changes.md): Modifications to consider when using the Genero Mobile for Android™.
- [Genero Mobile for iOS (GMI) 1.40 changes](0160-genero-mobile-for-ios-gmi-1-40-changes.md): Modifications to consider when using Genero Mobile for iOS.
- [Presentation styles changes](0161-presentation-styles-changes.md): Modifications to consider when using presentation styles.
- [Front calls changes](0162-front-calls-changes.md): Modifications to consider when using front calls.
- [Web components changes](0163-web-components-changes.md): Modifications to consider when using web components.
- [Database drivers changes](0164-database-drivers-changes.md): New and desupported database drivers.
- [DATETIME SQL type mappings](0165-datetime-sql-type-mappings.md): For some databases, the type mapping for DATETIME HOUR TO MINUTE has changed.
- [FreeTDS 1.00 for SQL Server](0166-freetds-1-00-for-sql-server.md): Genero 3.20 requires FreeTDS version 1.00+ to connect to SQL Server.
- [SQL Server drivers performance](0167-sql-server-drivers-performance.md): SQL Server ODI drivers based on FreeTDS, Easysoft and MS ODBC have been reviewed to achieve better execution times.
- [Performances with SQL interruption](0168-performances-with-sql-interruption.md): With some database drivers, performances can be impacted when using OPTIONS SQL INTERRUPT ON.
- [PostgreSQL 12 notes](0169-postgresql-12-notes.md): This topics contains notes about PostgreSQL 12 changes that affect Genero applications.
- [Sharing GLOBALS with C Extensions](0170-sharing-globals-with-c-extensions.md): Sharing of global variables with a C Extension is no longer supported.
- [Dynamic array assignment with .* notation](0171-dynamic-array-assignment-with-notation.md): The .* notation to assign dynamic arrays is discouraged.
- [Record copy without .* notation](0172-record-copy-without-notation.md): The .* notation to assign records is discouraged.
- [Circular dependency with IMPORT FGL](0173-circular-dependency-with-import-fgl.md): The compiler allows that two modules reference each other with IMPORT FGL.
- [Case insensitive names with UI methods](0174-case-insensitive-names-with-ui-methods.md): Methods of built-in classes using user interface object names are now case insensitive.
- [FLOAT/SMALLFLOAT to string conversion](0175-float-smallfloat-to-string-conversion.md): New FGLPROFILE entry fglrun.floatToCharScale2 for FLOAT/SMALLFLOAT types.
- [ORACLE rowid in sqlca.sqlerrm](0176-oracle-rowid-in-sqlca-sqlerrm.md): With ORACLE, the rowid of the last affected row is available in sqlca.sqlerrm.
- [Command tools changes](0177-command-tools-changes.md): Modifications to consider regarding command line tools.
