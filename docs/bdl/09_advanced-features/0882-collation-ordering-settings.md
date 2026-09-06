---
title: "Collation ordering settings"
source: "fgl-topics/c_fgl_localization_036.html"
breadcrumb: "Advanced features > Localization > Application locale > Defining the application locale > Collation ordering settings"
type: "concept"
description: "The runtime system supports a sorting functionality in tables. To sort the data rows, the runtime systems uses the standard C library functions to order character strings. The environment variable ..."
---

# Collation ordering settings

The runtime system supports a sorting functionality in tables.
To sort the data rows, the runtime systems uses the standard
C library functions to order character strings.

The environment variable LC\_COLLATE can be used to control sort order in Genero. You can for
example define this variable as "C" or "POSIX" to get a binary sort order.

When using LC\_COLLATE, set the LANG environment variable to define the global locale, if you use
LC\_ALL, it will overwrite all other LC\_\* variables defined.

## Related links

**Related concepts**  

[LC\_ALL (or LANG)](../07_configuration/0497-lc-all-or-lang.md "Defines the current application locale on UNIX platforms.")
