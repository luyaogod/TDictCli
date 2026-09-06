---
title: "DBSCREENOUT"
source: "fgl-topics/c_fgl_EnvVariables_DBSCREENOUT.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBSCREENOUT"
type: "concept"
---

# DBSCREENOUT

> Defines the output filename for text screen shots.

The DBSCREENOUT environment variable defines the output filename for text screen shots
when pressing Ctrl-P.

When using the TUI mode, if the user pressed the Ctrl-P key, the runtime system will dump
the current screen into the file defined by this variable.

Unlike DBSCREENDUMP, the output of DBSCREENOUT excludes the escape sequences of TTY
attributes.

## Related links

**Related concepts**  

[DBSCREENDUMP](0515-dbscreendump.md "Defines the output filename for text screen shots.")
