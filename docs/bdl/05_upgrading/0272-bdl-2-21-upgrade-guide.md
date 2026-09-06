---
title: "BDL 2.21 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_221.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide"
type: "concept"
---

# BDL 2.21 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.21.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.21 new features](0071-bdl-2-21-new-features.md "Features added in 2.21 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.20 upgrade guide](0280-bdl-2-20-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.20.").

## Child topics

- [Web Services changes](0273-web-services-changes.md): There are changes in support of web services in Genero 2.21.
- [PostgreSQL 8.4 and INTERVAL type](0274-postgresql-8-4-and-interval-type.md): The dbmpgs84x database driver requires your database schema use the INTERVAL type, rather than a CHAR(50) type.
- [fglcomp --build-rdd compiles the module](0275-fglcomp-build-rdd-compiles-the-module.md): fglcomp --build-rdd now creates both the .42m and .rdd files.
- [Unique and primary key constraint violation](0276-unique-and-primary-key-constraint-violation.md): Unique and primary key constraint violations mostly return error -268. However, error -269 may be checked too.
- [IMPORT with list of C-Extensions](0277-import-with-list-of-c-extensions.md): The IMPORT instruction for C extensions denies a comma-separated syntax.
- [Initializing dynamic arrays to null](0278-initializing-dynamic-arrays-to-null.md): The INITIALIZE TO NULL instruction clears the dynamic array.
- [Strict screen record definition for tables](0279-strict-screen-record-definition-for-tables.md): The fglform compiler of version 2.21.00 now makes a strict checking of the fields used in the screen record definition for table containers.
