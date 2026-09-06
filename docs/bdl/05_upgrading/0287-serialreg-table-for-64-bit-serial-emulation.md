---
title: "SERIALREG table for 64-bit serial emulation"
source: "fgl-topics/c_fgl_Migrate_to_220_serialreg.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > SERIALREG table for 64-bit serial emulation"
type: "concept"
---

# SERIALREG table for 64-bit serial emulation

> You must alter the SERIALREG table to do serial emulation on a BIGINT column.

The SERIALREG based serial emulation is defined by the following [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") entry:

```
dbi.database.<dbname>.ifxemul.datatype.serial.emulation = "regtable"
```

Version 2.20 introduces the [`BIGINT`](../08_language-basics/0554-bigint.md "The BIGINT data type is used for storing very large whole numbers.") data type, which is a 64-bit signed integer. You can use
`BIGSERIAL` or `SERIAL8` columns with IBM®
Informix®, and ODI drivers can emulate 64-bit serials in
other database servers. However, if you are using serial emulation based on the
`SERIALREG` table, you must redefine this table to change the
`LASTSERIAL` column data type to a `BIGINT`. If the
`BIGINT` data type is not supported by the database server, you can use a
`DECIMAL(20,0)` instead:

```
CREATE TABLE serialreg (
  tablename   VARCHAR2(50) NOT NULL,
  lastserial  BIGINT       NOT NULL,
  PRIMARY KEY ( tablename )
)
```

> **Important:**
>
> If you need to migrate an installed database using
> `SERIALREG`-based triggers, you will have to keep the current registered serials and
> use `ALTER TABLE` instead of `CREATE TABLE`. This example shows the
> `ALTER TABLE` syntax for SQL Server. Check the database server manuals for the exact
> syntax of the `ALTER TABLE`
> statement.
>
> ```
> ALTER TABLE serialreg ALTER COLUMN lastserial BIGINT NOT NULL
> ```

Additionally, all existing `SERIALREG`-based triggers must be modified, in order
to use `BIGINT` instead of `INTEGER` variables, otherwise you
will get `BIGINT` to `INTEGER` overflow errors. For example, to
modify existing triggers with SQL Server, you can use the `ALTER TRIGGER`
statement, which can be easily generated from the database browser tool (there is a
modify option in the context menu of triggers). After the existing
trigger code was generated, you must edit the code to replace the `INTEGER`
data type by `BIGINT` in the variable declarations, and execute the
`ALTER TRIGGER` statement.
