---
title: "User interface topics"
source: "fgl-topics/c_fgl_MigI4GL_010.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics"
type: "concept"
---

# User interface topics

> When migrating from I4GL to Genero BDL, review the differences between how windows and form content is rendered between the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.


## Child topics

- [Smooth migration with traditional UI mode](0348-smooth-migration-with-traditional-ui-mode.md): IBM® Informix® 4GL (I4GL) and Genero Business Development Language (BDL) handle windows and form content rendering differently.
- [Review application ergonomics](0349-review-application-ergonomics.md): Genero BDL no longer limits programs to executing a single interactive instruction, and provides additional GUI concepts such as drag-and-drop and tree views.
- [Terminal screen display handling](0350-terminal-screen-display-handling.md): Compared to I4GL, Genero BDL uses an optimized and specific design to handle terminal display, that introduces some differences.
- [Refreshing the user interface](0351-refreshing-the-user-interface.md): Genero BDL only refreshes the user interface when the runtime waits for user interaction. In certain scenarios, this can result in information displayed in an I4GL application not being displayed when running as a Genero BDL application.
- [SCREEN versus LAYOUT section](0352-screen-versus-layout-section.md): When writing new programs for GUI applications, it is recommended that you use a LAYOUT section instead of SCREEN. However, the SCREEN section is still supported to be used to design TUI mode forms.
- [MENU rendering and behavior](0353-menu-rendering-and-behavior.md): This topic describes differences between I4GL and FGL MENU instruction in TUI mode.
- [Migrating screen arrays to tables](0354-migrating-screen-arrays-to-tables.md): Tables in Genero BDL display using a real table widget, providing a more robust display and interaction than the I4GL screen array.
- [ON KEY / COMMAND vs ON ACTION](0355-on-key-command-vs-on-action.md): IBM® Informix® 4GL applications use the ON KEY and COMMAND [KEY] clauses to handle user actions.
- [TTY attributes alternatives](0356-tty-attributes-alternatives.md): IBM® Informix® 4GL programs use TTY attributes based on terminal capabilities, for colors and text intensity or effects.
- [Moving screen array rows with SCROLL](0357-moving-screen-array-rows-with-scroll.md): IBM® Informix® 4GL programs can use the SCROLL instruction to move screen array rows up and down, by preserving the TTY attributes.
- [The default SCREEN window](0358-the-default-screen-window.md): When the first interactive instruction is reached in a Genero BDL program, a default window named SCREEN is created.
- [Specifying WINDOW position and size](0359-specifying-window-position-and-size.md): With Genero BDL in GUI mode, window position and sizes are ignored; in TUI mode, window position and sizes are respected.
- [Right justified field labels](0360-right-justified-field-labels.md): I4GL forms that specify right-justified labels should be reviewed for update to LABEL form items.
- [Reduce multiple text screens](0361-reduce-multiple-text-screens.md): Moving beyond the 80x25 dimensions of a display may require a review of your dumb-terminal-oriented programs.
- [Positions of repeated form field tags](0362-positions-of-repeated-form-field-tags.md): The Genero Abstract User Interface definition (AUI tree) does not support misaligned repeating form field tags.
- [Subscripted form fields](0363-subscripted-form-fields.md): Subscripted form fields must be located and redefined when moving to Genero BDL.
- [WORDWRAP field attribute](0364-wordwrap-field-attribute.md): Use a TEXTEDIT field to replace repeated multi-line input fields.
- [Ignored form definition attributes](0365-ignored-form-definition-attributes.md): Field attributes inherited from the Informix® SQL PERFORM syntax should be reviewed for necessity and handling.
- [INPUT ARRAY behavior](0366-input-array-behavior.md): This topic describes INPUT ARRAY differences between I4GL and Genero BDL.
- [SCREEN RECORD fields specification](0367-screen-record-fields-specification.md): Genero BDL requires a comma separator in the field list of a screen record definition.
- [BEFORE/AFTER FIELD field prefix](0368-before-after-field-field-prefix.md): Recent I4GL versions deny the specification of a screen-record prefix in BEFORE FIELD and AFTER FIELD clauses, Genero BDL allows this for backward compatibility.
- [Form field input differences](0369-form-field-input-differences.md): This topic describes differences in field input between I4GL and Genero BDL.
- [Mixing COLOR and REVERSE attributes](0370-mixing-color-and-reverse-attributes.md): I4GL ignores the COLOR attribute when REVERSE is used, while Genero BDL mixes both attributes.
- [PROMPT rendering and behavior](0371-prompt-rendering-and-behavior.md): This topic describes differences between I4GL and FGL PROMPT instruction in TUI mode.
- [fgl_lastkey() function](0372-fgl-lastkey-function.md): In a given context and for a given key press, FGL fgl_lastkey() API does not return the same key numbers than I4GL fgl_lastkey().
