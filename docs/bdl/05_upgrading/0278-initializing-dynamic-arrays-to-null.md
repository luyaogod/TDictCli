---
title: "Initializing dynamic arrays to null"
source: "fgl-topics/c_fgl_Migrate_to_221_dynarray_init.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide > Initializing dynamic arrays to null"
type: "concept"
---

# Initializing dynamic arrays to null

> The INITIALIZE TO NULL instruction clears the dynamic array.

Starting with version 2.21.00, the INITIALIZE TO NULL instruction clears the dynamic arrays (i.e.
array.getLength() returns 0). Before this version, all elements of the dynamic array were kept, and
set to null. Since the old behavior was documented, this behavior change required a migration note.
The new behavior is expected by most programmers.

## Related links

**Related concepts**  

[INITIALIZE](../08_language-basics/0698-initialize.md "The INITIALIZE instruction initializes program variables with NULL or default values.")
