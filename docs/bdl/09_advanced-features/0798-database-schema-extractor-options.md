---
title: "Database schema extractor options"
source: "fgl-topics/c_fgl_DatabaseSchema_003.html"
breadcrumb: "Advanced features > Database schema > Database schema extractor options"
type: "concept"
---

# Database schema extractor options

> The fgldbsch tool extracts the schema description for an existing database.

Schema information is extracted from the database catalog tables. [fgldbsch](../13_programming-tools/2521-fgldbsch.md "The fgldbsch tool generates the database schema files from an existing database.") detects the type of database server after
connection and queries the appropriated system catalog tables.

The database system must be available and the database client environment (locale settings) must
be set properly in order to connect to the database engine and generate the schema files.

Generate the database schema files in the directory where the source code
resides or in one of the directories listed in the FGLDBPATH environment variable.

## Related links

**Related concepts**  

[FGLDBPATH](../07_configuration/0522-fgldbpath.md "Defines a list of paths to database schema files for compilers.")

## Child topics

- [Specifying the database source](0799-specifying-the-database-source.md)
- [Specifying the database driver](0800-specifying-the-database-driver.md)
- [Passing database user login and password](0801-passing-database-user-login-and-password.md)
- [Data type conversion control](0802-data-type-conversion-control.md)
- [Skip unsupported table definitions](0803-skip-unsupported-table-definitions.md)
- [Specifying the table schema/owner](0804-specifying-the-table-schema-owner.md): Providing the database schema of SQL tables is mandatory when extracting a schema from some databases.
- [Force extraction of system tables](0805-force-extraction-of-system-tables.md)
- [Specifying the output filename](0806-specifying-the-output-filename.md)
- [Extracting definition of a single table](0807-extracting-definition-of-a-single-table.md)
- [Controlling the character case](0808-controlling-the-character-case.md)
- [Using the verbose mode](0809-using-the-verbose-mode.md)
- [IBM Informix synonym tables](0810-ibm-informix-synonym-tables.md)
- [IBM Informix shadow columns](0811-ibm-informix-shadow-columns.md)
- [Running schema extractor in old mode](0812-running-schema-extractor-in-old-mode.md)
