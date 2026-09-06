---
title: "Understanding windows and forms"
source: "fgl-topics/c_fgl_windows_and_forms_002.html"
breadcrumb: "User interface > Form definitions > Windows and forms > Understanding windows and forms"
type: "concept"
---

# Understanding windows and forms

> This is an introduction to Genero windows and forms.

Programs manipulate windows and forms, to define display and input areas controlled by
interactive instructions such as the `INPUT` dialog.

A default window is always created with the name `SCREEN` where the main form of
the program can be displayed with `OPEN FORM` / `DISPLAY FORM`.
Secondary forms are typically displayed in secondary windows with `OPEN WINDOW + WITH
FORM`. A program can switch between windows with the `CURRENT WINDOW IS`
instruction.

When a dialog is started, it uses the form associated with the current window.

Forms define the layout and presentation of areas used by the dialogs (`INPUT`),
to display or input data.

Forms are defined in .42f compiled form files and are loaded at runtime to
be displayed in windows.

## Related links

**Related concepts**  

[What are dialog controllers?](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.")
