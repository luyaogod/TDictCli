---
title: "Command tools changes"
source: "fgl-topics/c_fgl_Migrate_to_600_fgl_tools.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 6.00 upgrade guide > Command tools changes"
type: "concept"
---

# Command tools changes

> Modifications to consider regarding command line tools.

This topic describes changes that may need code review.

See also [new 6.00 features of commands](../02_what-s-new-in-6-00/0002-what-s-new-in-6-00.md).

## Compilation error on comparing different record types

Before version 6.00.00, the fglcomp compiler was displaying the warning -8442,
if the code tries to compare records of different types. Starting with 6.00.00, the
fglcomp compiler now produces an error -8442:

```
DEFINE r1 RECORD
             id INTEGER, name VARCHAR(50)
        END RECORD
DEFINE r2 RECORD
             id INTEGER, name VARCHAR(100)
        END RECORD
MAIN
    IF r1.* == r2.* THEN
        DISPLAY "Identical"
    END IF
END MAIN
```

Compilation command:

```
$ fglcomp -M main.4gl
main.4gl:8:8:8:19:error:(-8442) Incompatible types: 'RECORD <anonymous>'
   and 'RECORD <anonymous>' are not comparable.
```

## `[DESCRIBE] DATABASE` detected by `-W stdsql`

Starting with version 6.00.03, the [fglcomp -W
stdsql](../13_programming-tools/2516-fglcomp.md) warning option now detects old legacy `[DESCRIBE] DATABASE`
instructions specifying a .sch database schema for `DEFINE … LIKE
…` instructions.

When `[DESCRIBE] DATABASE` is used in the module containing the
`MAIN` block, an implicit database connection is performed at runtime when the
MAIN block is executed. This can lead to runtime errors, when deploying your
application code on a production server where the database name used in `[DESCRIBE]
DATABASE` does not exist.

Since early versions of Genero BDL, it is recommended to specify the database schema with the
`SCHEMA` instruction, which does not do an implicit database connection in
`MAIN`.

Read [SCHEMA](../09_advanced-features/0793-schema.md "Defines the database schema files to be used for compilation.") for more details.

## Compiler warning for risky SQL cursor usage

Starting with version 6.00.03, the [fglcomp -W
fragile-cursor](../13_programming-tools/2516-fglcomp.md) warning option can be used to detect SQL cursor instructions that
have an undefined behavior at runtime:

- `DECLARE c CURSOR` with host variables or `INTO` variables having
  local function scope, and one of the following statements is used outside the scope of the function:
  `OPEN`, `FETCH`, `FOREACH`, `PUT`
  without a `USING`, `INTO` or `FROM` clause to re-bind
  local variables
- `OPEN c USING` followed by `OPEN` or `FOREACH`
  without `USING` clause, (Informix OPTOFC usage)

Read [Using program variables in static SQL](../10_sql-support/1117-using-program-variables-in-static-sql.md "Static SQL syntax supports the usage of program variables as SQL parameters.") for more details.

## Changes in earlier versions

Make sure to check the upgrade notes of earlier versions, to not miss changes introduced in
maintenance releases.

## Related links

**Related concepts**  

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
