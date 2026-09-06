---
title: "Version number meaning"
source: "fgl-topics/c_fgl_Migrate_to_all_versions.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > General BDL upgrade guide > Version number meaning"
type: "concept"
---

# Version number meaning

> A product version number identifies a specific release of the software product.

A given product version number defines the new features and bug fixes available in that version.

The product version number can be found by executing a Genero command line tool with the
`-V` option. For example, with Genero BDL, use the fglrun command
(the product version number is highlighted):

```
$ fglrun -V
fglrun 5.01.02 rev-17288c5d
Genero virtual machine
Target l64xl228
...
```

The product version number has the format `X.YY.ZZ`, where each part identifies
the level of new features and the amount of bug fixes available in that version.

For example:

- When upgrading from 2.50.15 to 2.50.16, you get mostly bug fixes; however, you may also get some
  small enhancements that do not impact the compatibility with other products of the same version
  family.
- When upgrading from 2.50.15 to 2.51.05, you get bug fixes and medium new features.
- When upgrading from 2.40.11 to 2.50.15, you get many bug fixes and significant new
  features.
- When upgrading from 2.40.11 to 3.00.12, you upgrade to a new major version, introducing
  important new features (including sometimes strategic software architecture changes), as well as a
  lot of bug fixes, compared to the previous version.

## Related links

**Related concepts**  

[P-Code compatibility](0087-p-code-compatibility.md "P-Code incompatibility (within .42m files) may be introduced from version to version.")

[VM and front-end compatibility](0088-vm-and-front-end-compatibility.md "Always combine Genero Virtual Machine with the latest compatible Front-End.")
