---
title: "BDL 2.32 upgrade guide"
source: "fgl-topics/c_fgl_Migrate_to_232.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.32 upgrade guide"
type: "concept"
---

# BDL 2.32 upgrade guide

> These topics describe product changes you must be aware of when upgrading to version 2.32.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This is an incremental upgrade guide that
> covers only topics related to the Genero BDL version specified in the page title. Check prior
> upgrade guides if you migrate from an earlier version. Make sure to also read about the new features
> for this Genero version.

Corresponding new features page: [BDL 2.32 new features](0069-bdl-2-32-new-features.md "Features added in 2.32 releases of the Genero Business Development Language.").

Previous upgrade guide: [BDL 2.30 upgrade guide](0261-bdl-2-30-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.30.").

## Child topics

- [Front-end protocol compression disabled](0257-front-end-protocol-compression-disabled.md): GUI communication does not require protocol compression on LAN networks.
- [SQLite driver no longer needs libiconv on Windows](0258-sqlite-driver-no-longer-needs-libiconv-on-windows.md): UTF-8 string data storage in SQLite requires conversion when the application is not UTF-8.
- [Need for Informix CSDK to compile C extensions](0259-need-for-informix-csdk-to-compile-c-extensions.md): Compiling C Extensions requires now the Informix CSDK.
- [FESQLC tool removal](0260-fesqlc-tool-removal.md): The ESQL/C compiler (fesql) has been removed from the Genero BDL product.
