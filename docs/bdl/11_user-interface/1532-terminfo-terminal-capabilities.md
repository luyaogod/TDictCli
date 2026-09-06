---
title: "TERMINFO terminal capabilities"
source: "fgl-topics/c_fgl_DynamicUI_031.html"
breadcrumb: "User interface > User interface basics > Using a text terminal > TERMINFO terminal capabilities"
type: "concept"
description: "When the INFORMIXTERM environment variable is set to terminfo , the runtime system will use the ncurses or curses library of the UNIX™ system to display and interact with the terminal device defined ..."
---

# TERMINFO terminal capabilities

When the INFORMIXTERM environment variable is set to `terminfo`, the runtime
system will use the ncurses or curses library of the UNIX™
system to display and interact with the terminal device defined by the TERM environment
variable.

Make sure that the Curses library is installed on your UNIX
operating system. Check [operating system installation
requirements](../04_installation/0035-system-packages.md "Some Genero BDL features need specific operating system packages.") for more details.

The TERMINFO environment variable can be used to define a different terminal capabilities
database as the default. If your UNIX system is
properly configured, there is no need to set the TERMINFO environment variable.

## Related links

**Related concepts**  

[INFORMIXTERM](../07_configuration/0539-informixterm.md "Defines terminal control library to be used.")

[TERM](../07_configuration/0500-term.md "Defines the type of terminal on UNIX platforms.")
