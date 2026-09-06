---
title: "Oracle MySQL specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_mys.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > Oracle® MySQL specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .mys.config Defines an explicit configuration to read MySQL options from. dbi.database.stores.mys.config = \"/opt/myapp/etc/my.cnf\" This parameter will be passed to the MySQL API ..."
---

# Oracle MySQL specific FGLPROFILE parameters

## `dbi.database.dsname.mys.config`

Defines an explicit configuration to read MySQL options from.

```
dbi.database.stores.mys.config = "/opt/myapp/etc/my.cnf"
```

This parameter will be passed to the MySQL API function
`mysql_options((MYSQL*), MYSQL_READ_DEFAULT_FILE, filename
)`.

It can be used to bypass reading the default MySQL configuration files, to define database client
settings in the `[client]` group, such as the client character set with the
`default-character-set` option.

On Microsoft™ Windows® platforms, the configuration file must be in DOS format.
