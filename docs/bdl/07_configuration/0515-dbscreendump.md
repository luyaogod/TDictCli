---
title: "DBSCREENDUMP"
source: "fgl-topics/c_fgl_EnvVariables_DBSCREENDUMP.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBSCREENDUMP"
type: "concept"
---

# DBSCREENDUMP

> Defines the output filename for text screen shots.

The DBSCREENDUMP environment variable defines the output filename for text screen shots
when pressing Ctrl-P.

When using the TUI mode, if the user pressed the Ctrl-P key, the runtime system will dump
the current screen into the file defined by this variable.

Unlike DBSCREENOUT, the output of DBSCREENDUMP includes the escape sequences of TTY
attributes, which makes it less readable.

## Related links

**Related concepts**  

[DBSCREENOUT](0516-dbscreenout.md "Defines the output filename for text screen shots.")
