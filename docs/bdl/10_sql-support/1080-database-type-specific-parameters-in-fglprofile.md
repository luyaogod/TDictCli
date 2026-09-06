---
title: "Database type specific parameters in FGLPROFILE"
source: "fgl-topics/c_fgl_Connections_016.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE"
type: "concept"
---

# Database type specific parameters in FGLPROFILE

> Specific connection parameters can be configured with FGLPROFILE entries.

The FGLPROFILE entries for database-server specific configuration all following the same syntax
scheme:

```
dbi.database.dsname.dbtype.param[.subparam] = "value"
```

Where dbtype identifies the database vendor type, such as
`ifx`, `ora`, `snc`.

## Child topics

- [Oracle DB specific FGLPROFILE parameters](1081-oracle-db-specific-fglprofile-parameters.md)
- [Oracle MySQL specific FGLPROFILE parameters](1082-oracle-mysql-specific-fglprofile-parameters.md)
- [MariaDB specific FGLPROFILE parameters](1083-mariadb-specific-fglprofile-parameters.md)
- [PostgreSQL specific FGLPROFILE parameters](1084-postgresql-specific-fglprofile-parameters.md)
- [SQL Server (MS ODBC) specific FGLPROFILE parameters](1085-sql-server-ms-odbc-specific-fglprofile-parameters.md)
- [SQL Server (FreeTDS driver) specific FGLPROFILE parameters](1086-sql-server-freetds-driver-specific-fglprofile-parameters.md)
- [SQL Server (Easysoft driver) specific FGLPROFILE parameters](1087-sql-server-easysoft-driver-specific-fglprofile-parameters.md)
- [Dameng specific FGLPROFILE parameters](1088-dameng-specific-fglprofile-parameters.md)
