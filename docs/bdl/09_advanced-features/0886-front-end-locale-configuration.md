---
title: "Front-end locale configuration"
source: "fgl-topics/c_fgl_localization_025.html"
breadcrumb: "Advanced features > Localization > Application locale > Front-end locale configuration"
type: "concept"
---

# Front-end locale configuration

> The host operating system on the front-end workstation must be able to handle the character set and fonts.

For instance, a Western-European Windows® is not
configured to handle Arabic applications. If you start an Arabic application, some graphical
problems may occur (for instance the title bar won't display Arabic characters, but unwanted
characters instead).

The GUI front-end software must support the conversion of the runtime system character set
to/from the character set used internally by the client, and must be configured with the correct
font to display the characters used by the application. For example, the default font for a
front-end installed on an English Windows system might not
be able to display Japanese characters. You must then change the font in the front-end configuration
panel. Refer to the front-end documentation to see how character set conversion and fonts can be
configured.

When using a TUI program in a terminal emulator such as Putty, XTerm or even the Windows Console, make sure the terminal is configured properly to display
the characters of the application locale. For example, on a Windows Console you can use the chcp command to change the current code
page.
