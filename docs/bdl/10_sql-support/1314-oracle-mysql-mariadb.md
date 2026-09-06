---
title: "Oracle MySQL / MariaDB"
source: "fgl-topics/c_fgl_odiagmys_001.html"
breadcrumb: "SQL support > SQL database guides > Oracle® MySQL / MariaDB"
type: "concept"
description: "Supported versions Genero BDL supports the following Oracle® MySQL versions: Oracle MySQL 8.4 LTS with dbmmys_8_4 ODI driver. Note: The MySQL 8.2 and 8.3 versions are Innovation Releases (IR), while ..."
---

# Oracle MySQL / MariaDB

## Supported versions

Genero BDL supports the following Oracle® MySQL versions:

- Oracle MySQL 8.4 LTS with
  `dbmmys_8_4` ODI driver.
  > **Note:**
  >
  > The MySQL 8.2 and 8.3 versions are Innovation Releases (IR), while MySQL versions 8.0 and 8.4 are
  > Long-Term Support releases (LTS). The `dbmmys_8_2` driver provided in previous Genero
  > versions has been desupported in favor to MySQL 8.4 with the new `dbmmys_8_4` ODI
  > driver, requiring libmysqlclient.so.24. If you need to access a MySQL 8.2 or
  > 8.3 server, install MySQL Connector client with libmysqlclient.so.24 and use
  > the `dbmmys_8_4` ODI driver.
- Oracle MySQL 9.7 LTS with
  `dbmmys_9_7` ODI driver.

Genero BDL supports the following MariaDB versions:

- MariaDB 10.2+, 11 and 12 with `dbmmdb_10_2` ODI driver.

## Child topics

- [Purpose of the Oracle MySQL / MariaDB SQL guide](1315-purpose-of-the-oracle-mysql-mariadb-sql-guide.md)
- [Installation (Runtime Configuration)](1316-installation-runtime-configuration.md): Oracle MySQL related installation topics.
- [Database concepts](1319-database-concepts.md): Oracle MySQL related database concepts topics.
- [Data dictionary](1324-data-dictionary.md): Oracle MySQL related data dictionary topics.
- [Data manipulation](1337-data-manipulation.md): Oracle MySQL related data manipulation topics.
- [BDL programming](1347-bdl-programming.md): Oracle MySQL related programming topics.
