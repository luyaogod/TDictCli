---
title: "Sharing GLOBALS with C Extensions"
source: "fgl-topics/c_fgl_Migrate_to_320_CExt_globals.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.20 upgrade guide > Sharing GLOBALS with C Extensions"
type: "concept"
---

# Sharing GLOBALS with C Extensions

> Sharing of global variables with a C Extension is no longer supported.

Before Genero BDL version 3.20, it was possible to produce a .h and
.c files with the `-G` option of fglcomp to
share global variables with C Extensions:

```
fglcomp -G myglobals.4gl
```

This command produced the myglobals.h and myglobals.c
files, which could be used to build the C Extension library and share global variables defined in
the myglobals.4gl file.

New features of the BDL language have required to review internals and deny global variable
sharing with C Extensions. In addition, this practice was not recommended as stated in previous
versions of the documentation.

To migrate existing code, implement a C Extension module that defines static variables in the C
module, and set/get functions to assign or read the value of these variables.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")
