---
title: "Core language changes"
source: "fgl-topics/c_fgl_Migrate_to_600_core_language.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 6.00 upgrade guide > Core language changes"
type: "concept"
---

# Core language changes

> Modifications to consider regarding the core language features and syntax.

## sqlca.sqlerrd[1-6] as BIGINT

Starting with FGL 6.00.01, the [`sqlca`](../10_sql-support/0988-the-sqlca-diagnostic-record.md "The sqlca variable is a predefined record containing SQL statement execution information.") record field `sqlerrd[1-6]` is now defined as a
`BIGINT` 8-byte integer, to conform to Informix CSDK 15 sqlca for the support of large
tables.

Consider checking your code for ROWID column usage, that are now on 8 bytes with Informix 15
large tables. As a general advice, use primary keys instead of [ROWID columns](../10_sql-support/1042-using-rowid-columns.md "Automatic ROWID columns is not a common database feature.").

As a side effect of this change, if the database client provides an API to retrieve the last
generated `BIGSERIAL` column value, including BIGSERIAL emulation with non-Informix
drivers where this is possible, the `sqlca.sqlerrd[2]` register can contain 8 byte
generated serials. Read [Auto-incremented columns (serials)](../10_sql-support/1023-auto-incremented-columns-serials.md "How to implement automatic record keys.") for more details.

## Desupport of parallel dialog API

Since FGL version 4.00, parallel dialogs and splitviews are desupported on front-end side.
However, the `fglcomp` compiler was still allowing the instructions and APIs for
parallel dialog programming.

Starting with FGL 6.00.00, the compiler will now product an error, if the code uses the parallel
dialog instructions `START DIALOG`, `TERMINATE DIALOG`, or related
APIs.

For more details, read the FGL 4.00 upgrade note [Removal of parallel dialogs / splitviews](0136-removal-of-parallel-dialogs-splitviews.md "Parallel dialog and splitview features for mobile are no longer supported.").

## Desupport of FIELD form item type

Starting with FGL 6.00.00, the fglform .per form compiler
does no longer allow the `FIELD` form item type definition.

```
ATTRIBUTES
FIELD f1 = customer.cust_name;
```

This feature has been deprecated in version 2.51, see [The FIELD form item type and .val schema file](0229-the-field-form-item-type-and-val-schema-file.md "Form files using the FIELD item type and/or .val attribute definitions must be reviewed.").

Form files using this syntax will no longer compile with version 6.00.

## Desupport of WIDGET and related form attributes

Starting with FGL 6.00.00, the fglform .per form compiler
does no longer allow the `CLASS`, `WIDGET`, `CONFIG`
and `OPTIONS` attributes:

```
ATTRIBUTES
f1 = customer.cust_name, WIDGET="FIELD_BMP",
  CONFIG="XS ExtraSmall S Small M Medium L Large XL ExtraLarge";
```

This feature is considered as deprecated since the first version of Genero, see [Migrating form field WIDGET="type"](0414-migrating-form-field-widget-type.md "BDS fields using the WIDGET attribute must be replaced by Genero BDL form item types.").

Form files using this syntax will no longer compile with version 6.00.

## Desupport of GUI server autostart

Starting with FGL 6.00.01, the automatic GUI server startup feature is no longer supported.

Check FGLPROFILE settings starting with `gui.server.autostart` .

For more details about front-end connections, read [GUI front-end connection](../11_user-interface/1520-gui-front-end-connection.md "This section explains runtime to front-end connection in its simplest form.").
