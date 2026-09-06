---
title: "SQLite"
source: "fgl-topics/c_fgl_Connections_034.html"
breadcrumb: "SQL support > Database connections > Database client environment > SQLite"
type: "concept"
description: "The SQLite database driver includes the SQLite library, except on systems where that library is commonly available, like Linux® distributions, MacOS® and mobile devices. Database locale: The SQLite ..."
---

# SQLite

1. The SQLite database driver includes the SQLite library, except on systems where that library is
   commonly available, like Linux® distributions,
   MacOS® and mobile devices.
2. Database locale: The SQLite library uses UTF-8. If the current application character set
   (LANG/LC\_ALL) is not UTF-8, like plain ASCII or UTF-8, the SQLite database driver will make
   appropriate character set conversions.
3. You can make a connection test with the sqlite3 command line tool.

## Related links

**Related tasks**  

[Prepare the runtime environment - connecting to the database](1470-prepare-the-runtime-environment-connecting-to-the-database.md "Prepare the runtime environment - connecting to the database")
