---
title: "Text mode screen dump"
source: "fgl-topics/c_fgl_DynamicUI_028.html"
breadcrumb: "User interface > User interface basics > Using a text terminal > Text mode screen dump"
type: "concept"
description: "For compatibility with IBM® Informix® 4GL, Genero supports the DBSCREENDUMP and DBSCREENOUT environment variables for debugging purpose, which allows you to take a screenshot when running in TUI mode ..."
---

# Text mode screen dump

For compatibility with IBM®
Informix® 4GL, Genero supports the DBSCREENDUMP and
DBSCREENOUT environment variables for debugging purpose, which allows you to take a screenshot when
running in TUI mode and save the result in a file.

To enable TUI screenshot, set either DBSCREENDUMP or DBSCREENOUT to the name of the output file,
then run your Genero program with FGLGUI=0 set and press the Ctrl-P key to dump the current screen.
Each time you press Ctrl-P, the output file will be overwritten.

The DBSCREENDUMP variable writes the screen with escape sequences
of TTY attributes, while DBSCREENOUT writes only the characters displayed
on the screen, which makes the output more readable.

If both variables are set, the runtime will generate both output files; however, use different
file names, otherwise the output is undefined.

## Related links

**Related concepts**  

[DBSCREENDUMP](../07_configuration/0515-dbscreendump.md "Defines the output filename for text screen shots.")

[DBSCREENOUT](../07_configuration/0516-dbscreenout.md "Defines the output filename for text screen shots.")
