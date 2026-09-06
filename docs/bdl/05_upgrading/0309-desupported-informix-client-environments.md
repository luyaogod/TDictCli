---
title: "Desupported Informix client environments"
source: "fgl-topics/c_fgl_Migrate_to_200_informix_drivers.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Desupported Informix® client environments"
type: "concept"
description: "Upgrade IBM Informix Client Software Development Kit (CSDK) to the most recent version."
---

# Desupported Informix client environments

> Upgrade IBM® Informix® Client Software Development Kit (CSDK) to the most recent version.

Always upgrade the IBM Informix Client Software Development Kit (CSDK) to the most recent version
supported by Genero BDL.

The database interface of Genero Business Development Language (BDL) version 2.00 was redesigned
to allow dynamic loading of database drivers. The following IBM
Informix drivers and environments have been
desupported with this redesign:

- ix210: Informix ESQL/C 2.10
- ix410: Informix ESQL/C 4.10
- ix501: Informix ESQL/C 5.01
- ix711: Informix ESQL/C 7.11
- ix720: Informix ESQL/C 7.20

If required, old IBM
Informix drivers can
be re-enabled in a next Genero BDL version. However, we strongly
recommend you to upgrade the IBM
Informix Client
Software Development Kit (CSDK) to the most recent version
supported by Genero BDL.

## Related links

**Related concepts**  

[Installation](../04_installation/0031-installation.md "This chapter contains installation and setup instructions.")
