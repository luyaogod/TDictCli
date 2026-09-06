---
title: "Database client software"
source: "fgl-topics/c_fgl_installation_005.html"
breadcrumb: "Installation > Software requirements > Database client software"
type: "concept"
---

# Database client software

> To connect to a database server, the database client software must be installed on the system where you run the Genero BDL programs.

The Genero runtime system uses database drivers to connect to database servers, as a database
client program. Database vendor-specific client software needs to be installed on the system where
you run the Genero programs.

Genero database drivers are shipped as shared libraries and require the database vendor client
software shared library such as Informix®
Client SDK (with ESQL/C), Oracle® Client
(with OCI), Microsoft® SQL Server ODBC. The database
driver to be selected must correspond to the database client type and version.

For a detailed list of supported databases, database driver names, and operating systems on which
those database drivers are supported, refer to the System Support matrix (available on the [Products download
page](https://4js.com/download/products/) of the Four Js web site) or contact your support center. This matrix also informs you
which database drivers will no longer be supported as of the next release.

## Related links

**Related information**  

[Table 1](../10_sql-support/1073-database-driver-specification-driver.md "Table 1")
