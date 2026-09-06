---
title: "Circular dependency with IMPORT FGL"
source: "fgl-topics/c_fgl_Migrate_to_320_circular_imports.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Circular dependency with IMPORT FGL"
type: "concept"
---

# Circular dependency with IMPORT FGL

> The compiler allows that two modules reference each other with IMPORT FGL.

Circular module references occur when several modules reference each other with `IMPORT
FGL`.

Before Genero BDL version 3.20, a circular reference was producing the compilation error
-8402.

Starting with Genero BDL version 3.20, circular references with `IMPORT FGL` are
now allowed.

## Related links

**Related concepts**  

[Circular module references](../09_advanced-features/0820-circular-module-references.md "Circular references between imported modules are allowed.")
