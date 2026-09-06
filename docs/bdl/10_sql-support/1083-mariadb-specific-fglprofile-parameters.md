---
title: "MariaDB specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_mdb.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > MariaDB specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .mdb.config Defines an explicit configration to read MariaDB options from. dbi.database.stores.mdb.config = \"/opt/myapp/etc/my.cnf\" This parameter will be passed to the MariaDB ..."
---

# MariaDB specific FGLPROFILE parameters

## `dbi.database.dsname.mdb.config`

Defines an explicit configration to read MariaDB options from.

```
dbi.database.stores.mdb.config = "/opt/myapp/etc/my.cnf"
```

This parameter will be passed to the MariaDB API function `mysql_options((MYSQL*),
MYSQL_READ_DEFAULT_FILE, filename )`.

It can be used to bypass reading the default MariaDB configuration files, to define database
client settings in the `[client]` group, such as the client character set with the
`default-character-set` option.

On Microsoft™ Windows® platforms, the configuration file must be in DOS format.
