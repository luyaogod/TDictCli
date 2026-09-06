---
title: "BDL 2.20 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_220.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide"
type: "concept"
---

# BDL 2.20 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.20.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.20 new features](0072-bdl-2-20-new-features.md "Features added in 2.20 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.11 upgrade guide](0298-bdl-2-11-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.11.").

## Child topics

- [Web Services changes](0281-web-services-changes.md): There are changes in support of web services in Genero 2.20.
- [Sort is now possible during INPUT ARRAY](0282-sort-is-now-possible-during-input-array.md): Built-in sort is available in INPUT ARRAY.
- [Cell attributes and buffered mode](0283-cell-attributes-and-buffered-mode.md): Must use the UNBUFFERED mode when setting cell attributes.
- [Field methods are more strict](0284-field-methods-are-more-strict.md): Dialog class methods are more strict regarding form field names.
- [Strict variable identification in SQL statements](0285-strict-variable-identification-in-sql-statements.md): Program variable identification in static SQL statements is more strict in version 2.20 than older versions.
- [SQL Warnings with non-Informix databases](0286-sql-warnings-with-non-informix-databases.md): SQL Warnings are now propagated for all database drivers, and can set the sqlca.sqlawarn, SQLSTATE and SQLERRMESSAGE registers.
- [SERIALREG table for 64-bit serial emulation](0287-serialreg-table-for-64-bit-serial-emulation.md): You must alter the SERIALREG table to do serial emulation on a BIGINT column.
- [Extracting the database schema with fgldbsch](0288-extracting-the-database-schema-with-fgldbsch.md): The fgldbsch database schema extraction tool has been updated to map native database types to newly-added types.
- [Database driver internal error changed from -768 to -6319](0289-database-driver-internal-error-changed-from-768-to-6319.md): The internal error raised was changed to avoid conflicts with an IBM® Informix® SQL error code.
- [Searching for image files on the application server](0290-searching-for-image-files-on-the-application-server.md): For security reasons, the image file transfer mechanism has been slightly modified in version 2.20.
- [Strict action identification in dialog methods](0291-strict-action-identification-in-dialog-methods.md): Actions referenced in methods of the dialog class must exist in the current dialog, or an error is raised.
- [Strict field identification in dialog methods](0292-strict-field-identification-in-dialog-methods.md): Fields referenced in methods of the dialog class must exist in the current dialog, or an error is raised.
- [Form compiler checking invalid layout definition](0293-form-compiler-checking-invalid-layout-definition.md): It is better to identify form layout mistakes when the form is compiled, rather than at runtime.
- [Database schema compatibility](0294-database-schema-compatibility.md): fgldbsch extracts specific type for BOOLEAN.
- [Predefined actions get automatically disabled depending on the context](0295-predefined-actions-get-automatically-disabled-depending-on-t.md): Dialogs will automatically disable some predefined actions, if it makes no sense to trigger the action in the current context.
- [BEFORE ROW no longer executed when array is empty](0296-before-row-no-longer-executed-when-array-is-empty.md): In order to trigger the BEFORE ROW block when entering an array, the array must not be empty.
- [Controlling INPUT ARRAY temporary row creation](0297-controlling-input-array-temporary-row-creation.md): Down move after last row in INPUT ARRAY creates a new temporary row.
