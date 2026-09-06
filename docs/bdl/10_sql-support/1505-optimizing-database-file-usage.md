---
title: "Optimizing database file usage"
source: "fgl-topics/c_fgl_odiagsqt_vacuum.html"
breadcrumb: "SQL support > SQL database guides > SQLite > BDL programming > Optimizing database file usage"
type: "concept"
description: "SQLite By default, when deleting a large amount of data in an SQLite database, it leaves behind empty space, causing the database file to be larger than stricly necessary. This might be an issue with ..."
---

# Optimizing database file usage

## SQLite

By default, when deleting a large amount of data in an SQLite database, it leaves behind empty
space, causing the database file to be larger than stricly necessary.

This might be an issue with some mobile applications, when the disk space of the mobile device is
limited.

## Solution

Execute the `VACUUM` SQL command, to truncate the database file and reduce the
disk usage.

According to the application, the `VACUUM` command can be executed:

- when starting the application,
- after doing a large db operation (like a synchronization with a central db),
- as a manual option that the user can trigger.

Note that SQLite also supports "`PRAGMA auto_vacuum`", but it appears that it's
not as efficient as the `VACUUM` command, regarding page fragmentation.

Pay attention to the fact that `VACUUM` needs twice the disk space of the actual
database file, because it rebuilds totally the db file.

`VACUUM` is not Informix SQL syntax, use `EXECUTE IMMEDIATE` to
perform this SQL statement:

```
EXECUTE IMMEDIATE "VACUUM"
```
