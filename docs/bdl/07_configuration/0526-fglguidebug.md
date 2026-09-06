---
title: "FGLGUIDEBUG"
source: "fgl-topics/c_fgl_EnvVariables_FGLGUIDEBUG.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLGUIDEBUG"
type: "concept"
---

# FGLGUIDEBUG

> Defines the debug level in GUI mode.

The FGLGUIDEBUG environment variable defines the debug level, when the [GUI mode](../11_user-interface/1516-genero-user-interface-modes.md "User interface modes allow you to adapt the application form rendering to different types of displays.") is used by the program.

By setting FGLGUIDEBUG to 1, the runtime system will display AUI protocol exchanges in
the stderr output of the console running the program on the server.

The runtime system displays detailed information about user interface events that occur
during program execution.

This debug log is to be used in development context only. The output format can change in next
product releases.

## Related links

**Related concepts**  

[FGLGUI](0525-fglgui.md "Defines the user interface mode to be used by the program.")
