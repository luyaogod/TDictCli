---
title: "TTY attributes alternatives"
source: "fgl-topics/c_fgl_MigI4GL_040.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > TTY attributes alternatives"
type: "concept"
description: "IBM Informix 4GL programs use TTY attributes based on terminal capabilities, for colors and text intensity or effects."
---

# TTY attributes alternatives

> IBM® Informix® 4GL programs use TTY attributes based on terminal capabilities, for colors and text intensity or effects.

With intructions such as `DISPLAY TO / BY NAME`, Genero BDL supports TTY
attributes specification like `BLUE`, `RED`, `REVERSE`,
`BOLD`, but it is recommended that you use those instructions for TUI programs
only.

For new GUI programs, use graphical user interface possibilities. For example, a good replacement
for TTY attributes is to use presentation styles, to centralize the decoration in external
.4st files, similar to HTML/CSS.

The GBC front-end also offer solutions to change the default rendering of application forms, by
using GBC themes and GBC customization. Refer to the GBC manual for more details.

## Related links

**Related concepts**  

[Presentation styles](../11_user-interface/1607-presentation-styles.md "Use presentation styles to specify decoration attributes for window and form elements.")
