---
title: "FGLGUI is 1 by default"
source: "fgl-topics/c_fgl_Mig0000_013.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > FGLGUI is 1 by default"
type: "concept"
---

# FGLGUI is 1 by default

> The default mode differs between Four Js BDS and Genero BDL.

With Four Js Business Development Suite (BDS), when the FGLGUI environment variable is not set,
the application starts in TUI mode (FGLGUI=0).

With Genero Business Development Language (BDL), when the FGLGUI environment variable is not
set, the application starts in GUI mode (FGLGUI=1).

Therefore, when migrating from Four Js BDS, it is recommended that you set FGLGUI=0 to run the
application in text mode as a first step.

## Related links

**Related concepts**  

[FGLGUI](../07_configuration/0525-fglgui.md "Defines the user interface mode to be used by the program.")
