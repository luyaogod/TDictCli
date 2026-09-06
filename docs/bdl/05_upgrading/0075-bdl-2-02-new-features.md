---
title: "BDL 2.02 new features"
source: "fgl-topics/fgl_whatsnew_202.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.02 new features"
type: "topic"
---

# BDL 2.02 new features

> Features added in 2.02 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.02 upgrade guide](0303-bdl-2-02-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.02.").

Prior new features guide: [BDL 2.01 new features](0076-bdl-2-01-new-features.md "Features added in 2.01 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Share global variables between the Genero source and the C Extension, by using the `-G` option of `fglcomp`. | Feature is desupported in version 3.20, see [Sharing GLOBALS with C Extensions](0170-sharing-globals-with-c-extensions.md "Sharing of global variables with a C Extension is no longer supported."). |
| Customize the runtime system error messages according to the current locale. | See [Runtime system messages](../09_advanced-features/0888-runtime-system-messages.md "This section describes how to translate default English runtime system message files in a different language."). |
| New debugger commands (`ptype`).Avoid switching into debug mode with SIGTRAP (Unix) or CTRL-Break (Windows™) with the new `fglrun.ignoreDebuggerEvent` FGLPROFILE entry. | See [Integrated debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs."). |

| Overview | Reference |
| --- | --- |
| Specify a `TABINDEX` of zero to exclude the form item from the tagging list. | See [TABINDEX attribute](../11_user-interface/1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item."). |

| Overview | Reference |
| --- | --- |
| Some common SQL statements have been added to the static SQL syntax, such as `TRUNCATE TABLE`, `RENAME INDEX`, `CREATE` / `ALTER` / `DROP` / `RENAME SEQUENCE`. | See [Static SQL statements](../10_sql-support/1115-static-sql-statements.md "Describes static SQL statements supported in the language."). |
| With Oracle, specify the `SELECT` statement producing the unique session identifier which is used for temporary table names. | See [Oracle DB specific FGLPROFILE parameters](../10_sql-support/1081-oracle-db-specific-fglprofile-parameters.md). |
| To emulate Informix® temporary tables in Oracle, set the `temptables.emulation` parameter to use `GLOBAL TEMPORARY TABLES` instead of permanent tables.. | See [Using the global temporary table emulation](../10_sql-support/1392-using-the-global-temporary-table-emulation.md). |
