---
title: "BDL 3.00 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_300v.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide"
type: "concept"
---

# BDL 3.00 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 3.00.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 3.00 new features](0064-bdl-3-00-new-features.md "Features added in 3.00 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.51 upgrade guide](0226-bdl-2-51-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.51.").

## Child topics

- [Web Services changes](0207-web-services-changes.md): There are changes in support of web services in Genero 3.00.
- [Form definitions for mobile applications](0208-form-definitions-for-mobile-applications.md): Genero version 3 supports grid and stack-based layout for mobile applications.
- [Database drivers changes](0209-database-drivers-changes.md): Desupported database drivers.
- [Oracle DB NUMBER type](0210-oracle-db-number-type.md): The NUMBER/FLOAT Oracle® data type can now be extracted by fgldbsch to create .sch files.
- [Oracle DB scroll cursor emulation removal](0211-oracle-db-scroll-cursor-emulation-removal.md): The scroll cursor emulation has been removed in the Oracle® DB driver.
- [MySQL VARCHAR size limit](0212-mysql-varchar-size-limit.md): MySQL 5 VARCHAR columns can be used to store VARCHAR(N>255) values.
- [MySQL DATETIME fractional seconds](0213-mysql-datetime-fractional-seconds.md): MySQL 5.6.4 TIME and DATETIME types support fractions of seconds that can be used to store DATETIME HOUR TO FRACTION(N) or DATETIME YEAR TO FRACTION(N).
- [PostgreSQL DATETIME type mapping change](0214-postgresql-datetime-type-mapping-change.md): Conversion of DATETIME type with fractional seconds to PostgreSQL TIME(N)/TIMESTAMP(N) was invalid and has been reviewed.
- [MariaDB support](0215-mariadb-support.md): The MariaDB database is now supported by Genero 3.00.
- [FreeTDS driver supports SQL Server 2008, 2012, 2014](0216-freetds-driver-supports-sql-server-2008-2012-2014.md): The FreeTDS driver can now be used for SQL Server versions > 2005.
- [FGL_GETVERSION() built-in function](0217-fgl-getversion-built-in-function.md): The FGL_GETVERSION() function now returns the product version number (for example: 3.00.00).
- [Built-in front-end icons desupport](0218-built-in-front-end-icons-desupport.md): Image resources included in front-ends are desupported with Genero 3.00.
- [Presentation styles changes](0219-presentation-styles-changes.md): Deprecated and renamed presentation style attributes.
- [Front calls changes](0220-front-calls-changes.md): Describes changes applied to front calls.
- [SERIAL emulation with SQL Server](0221-serial-emulation-with-sql-server.md): The SERIAL and BIGSERIAL types can be emulated with triggers and sequences when using SQL Server 2012 and higher.
- [Improved compilation time](0222-improved-compilation-time.md): The fglcomp and fglform compilers have been reviewed to achieve faster compilation.
- [Preprocessor changes](0223-preprocessor-changes.md): Several bugs have been fixed in the preprocessor, that can now result in a compilation error.
- [Current system time in UTC](0224-current-system-time-in-utc.md): Use the util.Datetime.getCurrentAsUTC() method to get the current system date/time in UTC.
- [Structured ARRAYs in list dialogs](0225-structured-arrays-in-list-dialogs.md): ARRAYs with sub-records can be used in list dialogs, to simplify array definition based on database tables, requiring additional information at runtime.
