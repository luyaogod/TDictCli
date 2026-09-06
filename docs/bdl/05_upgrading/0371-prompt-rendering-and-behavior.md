---
title: "PROMPT rendering and behavior"
source: "fgl-topics/c_fgl_MigI4GL_055.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > PROMPT rendering and behavior"
type: "concept"
---

# PROMPT rendering and behavior

> This topic describes differences between I4GL and FGL PROMPT instruction in TUI mode.

## ESC key validates the PROMPT

With IBM® Informix® 4GL, when you press the ESC key during a `PROMPT`, the interactive
instruction continues. User has to hit the ENTER key to validate the input.

With Genero BDL, in TUI mode, an ESC keypress will validate the `PROMPT` dialog,
like other single-dialogs such as `INPUT`.

The behavior of Genero BDL is more consistent than I4GL in this context.

## Related links

**Related concepts**  

[Prompt for values (PROMPT)](../11_user-interface/1888-prompt-for-values-prompt.md "The PROMPT instruction provides unique field input in an automatic pop-up window.")
