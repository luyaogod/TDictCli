---
title: "Database source specification (source)"
source: "fgl-topics/c_fgl_Connections_008.html"
breadcrumb: "SQL support > Database connections > Connection parameters > Database source specification (source)"
type: "concept"
description: "In database connection parameters, the source parameter identifies the data source name. If the source parameter is defined with an empty value (\"\"), the database interface connects to the default ..."
---

# Database source specification (source)

In database connection parameters, the `source` parameter identifies the
data source name.

If the `source` parameter is defined with an empty value (""), the
database interface connects to the default database server, which is usually the local
server.

If the `source` entry is not present in [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files"), direct database specification method takes
place.

| Database Type | Value of "source" entry | Description |
| --- | --- | --- |
| Dameng® | `datasource` | Database source name as defined in the dm\_svc.conf file. |
| Generic ODBC | `datasource` | ODBC Data Source |
| IBM® Informix® | `dbname[@dbserver]` | IBM Informix database specification |
| Oracle® MySQL / MariaDB | `dbname[@host[:port]]`or`dbname[@localhost~socket]` | Database Name @ Host Name: TCP PortorDatabase Name @ Local host ~ UNIX™ socket file |
| Oracle Database | `tnsname` | Oracle TNS Service name |
| PostgreSQL | `dbname[@host[:port]][?options]` | Database Name @ Host Name : TCP Port ? PostgreSQL URI-style query string options |
| Microsoft™ SQL Server | `datasource[?options]` | ODBC Data Source ? ODBC connection string parameters |
| SQLite | `filename`or`:memory:` | Path to the database file, or simple filename to be found with DBPATH, or `:memory:` to create a database in memory. |

## Related links

**Related concepts**  

[Direct database specification method](1077-direct-database-specification-method.md "Genero BDL applies direct database source specification when no FGLPROFILE entry corresponds to the database name used in programs.")
