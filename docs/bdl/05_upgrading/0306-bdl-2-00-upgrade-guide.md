---
title: "BDL 2.00 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_200.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide"
type: "concept"
---

# BDL 2.00 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.00.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.00 new features](0077-bdl-2-00-new-features.md "Features added in 2.00 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 1.33 upgrade guide](0327-bdl-1-33-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 1.33.").

## Child topics

- [Web Services changes](0307-web-services-changes.md): There are changes in support of Web services in Genero 2.00.
- [Runner creation is no longer needed](0308-runner-creation-is-no-longer-needed.md): Starting with version 2.00, you no longer need to recompile/build a runner.
- [Desupported Informix client environments](0309-desupported-informix-client-environments.md): Upgrade IBM® Informix® Client Software Development Kit (CSDK) to the most recent version.
- [Database drivers changes](0310-database-drivers-changes.md): Desupported database drivers.
- [fglmkrtm tool removed](0311-fglmkrtm-tool-removed.md): The fglmkrtm tool has been removed, as database drivers are loaded dynamically.
- [fglinstall tool removed](0312-fglinstall-tool-removed.md): The fglinstall tool has been removed from the distribution.
- [Linking the utility functions library](0313-linking-the-utility-functions-library.md): All utility functions are in the libfgl4js.42x library, up until 2.21.
- [Dynamic C extensions](0314-dynamic-c-extensions.md): Dynamic C extensions are automatically loaded with IMPORT instructions.
- [WANTCOLUMNSANCHORED is desupported](0315-wantcolumnsanchored-is-desupported.md): Use UNMOVABLECOLUMNS to specify that table columns cannot be moved around by the user.
- [PIXELWIDTH / PIXELHEIGHT are deprecated](0316-pixelwidth-pixelheight-are-deprecated.md): Use the WIDTH and HEIGHT attributes to specify the size of an image.
- [Prefetch parameters with Oracle](0317-prefetch-parameters-with-oracle.md): Prefetch parameters allow an application to automatically fetch rows from the Oracle® database when opening a cursor.
- [Preprocessor directive syntax changed](0318-preprocessor-directive-syntax-changed.md): The preprocessor directives use an ampersand character (&) instead of a hash (#) character.
- [Static SQL cache is removed](0319-static-sql-cache-is-removed.md): The Static SQL Cache has been removed.
- [SQL directive set removed](0320-sql-directive-set-removed.md): The SQL directive set specification has been removed.
- [Connection database schema specification](0321-connection-database-schema-specification.md): Changes with FGLPROFILE entries to define the database schema at runtime.
- [Schema extraction tool changes](0322-schema-extraction-tool-changes.md): The fgldbsch schema extractor is recommended, and has been enhanced.
- [Connection parameters in FGLPROFILE when using Informix](0323-connection-parameters-in-fglprofile-when-using-informix.md): The dbi.database.* connection parameters defined in FGLPROFILE are used by the Informix® driver
- [Inconsistent USING clauses](0324-inconsistent-using-clauses.md): Having data types changing at each execute is no longer supported.
- [Usage of RUN IN FORM MODE](0325-usage-of-run-in-form-mode.md): RUN ... IN LINE MODE is recommended to run interactive applications.
- [TTY and COLOR WHERE attribute](0326-tty-and-color-where-attribute.md): All types of fields now allow TTY attributes and the conditional COLOR WHERE attribute.
