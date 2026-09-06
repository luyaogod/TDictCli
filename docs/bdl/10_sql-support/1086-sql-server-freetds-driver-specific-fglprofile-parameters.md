---
title: "SQL Server (FreeTDS driver) specific FGLPROFILE parameters"
source: "fgl-topics/c_fgl_dbvendor_param_ftm.html"
breadcrumb: "SQL support > Database connections > Database type specific parameters in FGLPROFILE > SQL Server (FreeTDS driver) specific FGLPROFILE parameters"
type: "concept"
description: "dbi.database. dsname .ftm.logintime Connection timeout (in seconds). dbi.database.stores.ftm.logintime = 5 Set this parameter to raise an SQL error if the connection can not be established after the ..."
---

# SQL Server (FreeTDS driver) specific FGLPROFILE parameters

## `dbi.database.dsname.ftm.logintime`

Connection timeout (in seconds).

```
dbi.database.stores.ftm.logintime = 5
```

Set this parameter to raise an SQL error if the connection can not be
established after the given number of seconds.

The default is 5 seconds.

## `dbi.database.dsname.ftm.prefetch.rows`

Maximum number of rows to be pre-fetched.

```
dbi.database.stores.ftm.prefetch.rows = 50
```

Use this parameter to increase performance by defining the maximum number
of rows to be fetched into the db client buffer. However, the bigger this
parameter is, the more memory is used by each program.

The default is 10 rows.
