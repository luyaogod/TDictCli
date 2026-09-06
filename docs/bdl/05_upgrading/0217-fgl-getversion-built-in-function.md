---
title: "FGL_GETVERSION() built-in function"
source: "fgl-topics/c_fgl_Migrate_to_300_fgl_getversion.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.00 upgrade guide > FGL_GETVERSION() built-in function"
type: "concept"
---

# FGL_GETVERSION() built-in function

> The FGL_GETVERSION() function now returns the product version number (for example: 3.00.00).

Prior to Genero 3.00, the `fgl_getversion()` built-in function returned the internal build
number.

Starting with Genero 3.00, the function returns the product version number as a
string, such as 3.00.00.

## Related links

**Related concepts**  

[Built-in functions](../15_library-reference/2724-built-in-functions.md "A built-in function is a predefined function that is part of the runtime system, or provided as a library function automatically loaded when a program starts. The built-in functions are part of the language.")
