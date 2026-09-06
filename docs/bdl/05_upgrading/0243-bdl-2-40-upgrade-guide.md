---
title: "BDL 2.40 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_240.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide"
type: "concept"
---

# BDL 2.40 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.40.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.40 new features](0068-bdl-2-40-new-features.md "Features added in 2.40 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.32 upgrade guide](0256-bdl-2-32-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.32.").

## Child topics

- [Web Services changes](0244-web-services-changes.md): There are changes in support of web services in Genero 2.40.
- [Database drivers changes](0245-database-drivers-changes.md): Desupported database drivers.
- [Program size option removal (fglrun -s)](0246-program-size-option-removal-fglrun-s.md): The -s option of fglrun is no longer available.
- [Informix SERIAL emulation with SQL Server](0247-informix-serial-emulation-with-sql-server.md): SERIAL type emulation has been enhanced for SQL Server.
- [SIZEPOLICY attribute removal for containers](0248-sizepolicy-attribute-removal-for-containers.md): The SIZEPOLICY attribute is no longer available for layout containers like TABLE / GRID.
- [The LVARCHAR type in IBM Informix databases](0249-the-lvarchar-type-in-ibm-informix-databases.md): Native LVARCHAR type of Informix is now mapped by default to a large VARCHAR in schema file.
- [Right-trim collation for character types in SQLite](0250-right-trim-collation-for-character-types-in-sqlite.md): CHAR and VARCHAR columns in SQLite need to be defined with a TRIM collation to ignore trailing spaces in comparisons.
- [Message files support now 4-bytes integer message numbers](0251-message-files-support-now-4-bytes-integer-message-numbers.md): 2-byte .msg message number limitation was removed.
- [MySQL client library version change in MySQL 5.5.11](0252-mysql-client-library-version-change-in-mysql-5-5-11.md): Shared library version number of the MySQL client library must match the library used to link the ODI driver.
- [New compiler warning to avoid action shadowing](0253-new-compiler-warning-to-avoid-action-shadowing.md): Prevent the same action name at different levels of ON ACTION handlers in a dialog.
- [Runtime error raised when report dimensions are invalid](0254-runtime-error-raised-when-report-dimensions-are-invalid.md): Report page length checking error -4375 might occur at compile time or runtime.
- [Linker checks all referenced functions](0255-linker-checks-all-referenced-functions.md): The linker checks definition of all functions referenced in all modules provided in the link command.
