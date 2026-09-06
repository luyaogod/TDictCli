---
title: "fgldbsch"
source: "fgl-topics/c_fgl_tools_fgldbsch.html"
breadcrumb: "Programming tools > Command reference > fgldbsch"
type: "concept"
---

# fgldbsch

> The fgldbsch tool generates the database schema files from an existing database.

## Syntax

```
fgldbsch -db dbname [options]
```

1. dbname is the name of the database from which the schema is to be extracted.
2. options are described in Table 1.

## Options

| Option | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |
| `-H` | Displays long help. |
| `-v` | Enable verbose mode (display information messages). |
| `-ct` | Display data type conversion tables. |
| `-cx dbtype` | Display data type conversion table for the given database type. |
| `-db dbname` | Specify the database as dbname. This option is required to generate the schema files. |
| `-dv dbdriver` | Specify the database driver to be used. |
| `-un user` | Define the user name for database connection as user. |
| `-up pswd` | Define the user password for database connection as pswd. |
| `-ow owner` | Define the owner of the database tables as owner. |
| `-cv string` | Specify the data type conversion rules by character positions in string. |
| `-of name` | Specify output files prefix, default is database name. |
| `-tn tabname` | Extract the description of a specific table. |
| `-ie` | Ignore tables with columns having data types that cannot be converted. |
| `-cu` | Generate upper case table and column names. |
| `-cl` | Generate lower case table and column names. |
| `-cc` | Generate case-sensitive table and column names. |
| `-sc` | Extract shadow columns. |
| `-sl` | Sloppy mode:Allow extraction of table and column names not supported by Genero BDL such as `"Customer Name"`, or when table or column names are similar but different in case-sensitivity like `"Stock"` versus `"STOCK"`. |
| `-st` | Extract system tables. |
| `-om` | Run schema extractor in old fglschema mode, allowing the following compatibility options:`-c` is equivalent to `-cv BBBBBB...`, enabling Informix SQL type to FGL data type conversion.`-r` writes partial tables definitions into the .sch file, for tables having columns with unsupported types (`-ie` skips the whole table definition) |

## Usage

The fgldbsch command line tool extracts the schema description for any
database supported by the product.

The .sch schema file is mandatory to compile forms or source modules using
the `SCHEMA` instruction.

## Related links

**Related concepts**  

[Database schema](../09_advanced-features/0791-database-schema.md "Defines database table structures with column type information to be reused in program variable definitions.")
