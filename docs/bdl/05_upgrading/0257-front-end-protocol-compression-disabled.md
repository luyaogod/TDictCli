---
title: "Front-end protocol compression disabled"
source: "fgl-topics/c_fgl_Migrate_to_232_feprotocol_compression.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.32 upgrade guide > Front-end protocol compression disabled"
type: "concept"
---

# Front-end protocol compression disabled

> GUI communication does not require protocol compression on LAN networks.

> **Important:**
>
> Starting with version 4.00, the AUI protocol compression is desupported.
> See [Removal of AUI protocol compression](0149-removal-of-aui-protocol-compression.md "The AUI protocol compression is desupported.").

Until version 2.32.00, front-end protocol compression was enabled by default, to speed up GUI
communication on slow networks. However, on regular networks, compression is useless and can be
disabled to save processing resources. With version 2.32.00, the compression is now disabled by
default. If needed, compression can be enabled with this FGLPROFILE
entry:

```
gui.protocol.format = "zlib"
```

Note also that compression needs the zlib library to be present on the computer where
fglrun executes. Starting with 2.32.00, the product package no longer includes
the fallback zlib library ($FGLDIR/lib/libzfgl.so or
%FGLDIR%\bin\libzfgl.dll). If no standard zlib is installed on your system,
compression will not be possible.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
