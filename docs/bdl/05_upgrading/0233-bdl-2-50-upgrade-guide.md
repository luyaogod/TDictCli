---
title: "BDL 2.50 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_250.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide"
type: "concept"
---

# BDL 2.50 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.50.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.50 new features](0066-bdl-2-50-new-features.md "Features added in 2.50 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.41 upgrade guide](0242-bdl-2-41-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.41.").

## Child topics

- [Web Services changes](0234-web-services-changes.md): There are changes in support of web services in Genero 2.50.
- [Database drivers changes](0235-database-drivers-changes.md): Desupported database drivers.
- [TEXT/BYTE support with FTM/ESM database drivers](0236-text-byte-support-with-ftm-esm-database-drivers.md): FTM and ESM database drivers TEXT/BYTE type mapping has changed.
- [Presentation styles changes](0237-presentation-styles-changes.md): Deprecated and renamed presentation style attributes.
- [Floating point to string conversion](0238-floating-point-to-string-conversion.md): The default formatting of a DECIMAL(P), SMALLFLOAT and FLOAT adapts to the significant digits of the value.
- [Implicit creation of certificates for HTTPS](0239-implicit-creation-of-certificates-for-https.md): Certificates for HTTPS are now created implicitly when nothing is specified in FGLPROFILE.
- [PostgreSQL schema extraction needs namespace](0240-postgresql-schema-extraction-needs-namespace.md): To extract a database schema from PostgreSQL, the fgldbsch tool now requires db namespace specification.
- [Client stubs managing multipart changes](0241-client-stubs-managing-multipart-changes.md): You must update client programs that call client stubs managing multipart.
