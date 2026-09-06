---
title: "fgl_system() function"
source: "fgl-topics/c_fgl_Mig0000_029.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > 4GL Programming topics > fgl_system() function"
type: "concept"
---

# fgl_system() function

> With Genero BDL, the fgl_system() function no longer raises a terminal window by default, but some front-ends offer a workaround.

The `fgl_system()` function is still supported in Genero Business Development
Language, but it does not raise a terminal window on the front-end as with Four Js Business
Development Suite (BDS). However, some front-ends implement a workaround for this feature, based on
the detection of special strings displayed to stdout by fglrun. See the front-end documentation for
more details.

## Related links

**Related concepts**  

[fgl\_system()](../15_library-reference/2784-fgl-system.md "Runs a command on the application server.")
