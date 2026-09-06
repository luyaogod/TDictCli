---
title: "Database support on mobile devices"
source: "fgl-topics/c_fgl_mobile_database_support.html"
breadcrumb: "Mobile applications > Database support on mobile devices"
type: "concept"
---

# Database support on mobile devices

> On the device, a Genero app can use SQL for data management.

## Databases supported on mobile devices

Only SQLite can be used on mobile devices. SQLite has a small footprint, is free and
readily available.

The database driver (`dbmsqt`) and the SQLite library are built into the runtime
system for mobile Genero apps. No database driver specification is required when running on
mobile.

To read more about SQLite programming, see [Using SQLite database in mobile apps](5098-using-sqlite-database-in-mobile-apps.md "On the device, Genero Mobile uses the SQLite database only.").

## Synchronizing data with a central database

When local mobile app data needs to be synchronized with a central database, you must write your
own synchronization routines using Web Services. You must implement a back-end service to collect
mobile database updates and to send central database changes back to the mobile app.

> **Important:**
>
> If you are using GMI, and the device goes into standby mode, the application
> does not run in the background and activities with the network are suspended. If you are
> synchronizing data with the server, and the device goes into standby, the synchronization is
> suspended until the device resumes from standby. If you have a long synchronization, you need to
> either disable the sleep to allow the synchronization to complete, or accept that the synchronization
> will suspend when the device goes into standby mode.

## Related links

**Related concepts**  

[SQLite](../10_sql-support/1466-sqlite.md "SQLite")

## Child topics

- [Using SQLite database in mobile apps](5098-using-sqlite-database-in-mobile-apps.md): On the device, Genero Mobile uses the SQLite database only.
