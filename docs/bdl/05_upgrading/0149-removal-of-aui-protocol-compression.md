---
title: "Removal of AUI protocol compression"
source: "fgl-topics/c_fgl_Migrate_to_400_gui_protocol_compression.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > Removal of AUI protocol compression"
type: "concept"
---

# Removal of AUI protocol compression

> The AUI protocol compression is desupported.

Starting with version 4.00, the [AUI protocol](../11_user-interface/1522-the-front-end-protocol.md)
compression is desupported.

The following FGLPROFILE entry is no longer
supported:

```
gui.protocol.format = { "default" | "zlib" }
```

## Related links

**Related concepts**  

[GUI front-end connection](../11_user-interface/1520-gui-front-end-connection.md "This section explains runtime to front-end connection in its simplest form.")
