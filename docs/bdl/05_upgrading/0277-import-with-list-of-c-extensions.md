---
title: "IMPORT with list of C-Extensions"
source: "fgl-topics/c_fgl_Migrate_to_221_import_cext.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide > IMPORT with list of C-Extensions"
type: "concept"
---

# IMPORT with list of C-Extensions

> The IMPORT instruction for C extensions denies a comma-separated syntax.

Before version 2.21.00, the `IMPORT` instruction for C
extensions was documented as allowing a comma-separated list of
libraries:

```
IMPORT lib1, lib2
```

This compiled, but at runtime only the first library was found.
Using elements of the other libraries raised a runtime error.

With 2.21.00 and the new .42m module importation support, the compiler is
now more strict and denies the comma-separated syntax. You must specify every library, Java class
or .4gl module in separate
lines:

```
IMPORT lib1
IMPORT JAVA myclass
IMPORT FGL mymodule
```

## Related links

**Related concepts**  

[Importing modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")

[Content of a .4gl module](../09_advanced-features/0788-content-of-a-4gl-module.md "A module defines a set of program elements.")
