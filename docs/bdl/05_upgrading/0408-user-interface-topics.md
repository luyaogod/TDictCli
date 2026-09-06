---
title: "User interface topics"
source: "fgl-topics/c_fgl_Mig0000_010.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics"
type: "concept"
---

# User interface topics

> When migrating from Four Js BDS to Genero BDL, review the differences between how windows and form content is rendered between the two products. Reviewing the differences allows you to plan and prepare for a smooth migration.


## Child topics

- [Smooth migration with traditional UI mode](0409-smooth-migration-with-traditional-ui-mode.md): Four Js BDS and Genero Business Development Language (BDL) handle windows and form content rendering differently.
- [Front-end compatibility](0410-front-end-compatibility.md): With Genero BDL, you must use one of the Genero front-ends
- [FGLGUI is 1 by default](0411-fglgui-is-1-by-default.md): The default mode differs between Four Js BDS and Genero BDL.
- [FGLPROFILE: GUI configuration](0412-fglprofile-gui-configuration.md): Identify the Four Js BDS FGLPROFILE GUI configuration entries that are no longer supported, and the Genero BDL equivalent (where relevant).
- [Key labels versus action defaults](0413-key-labels-versus-action-defaults.md): Genero BDL introduces the ON ACTION block to define actions. While the ON KEY block is still supported, it comes with limitations.
- [Migrating form field WIDGET="type"](0414-migrating-form-field-widget-type.md): BDS fields using the WIDGET attribute must be replaced by Genero BDL form item types.
- [SCREEN versus LAYOUT section](0415-screen-versus-layout-section.md): When writing new programs for GUI applications, it is recommended that you use a LAYOUT section instead of SCREEN. However, the SCREEN section is still supported to be used to design TUI mode forms.
- [Migrating screen arrays to tables](0416-migrating-screen-arrays-to-tables.md): A TABLE container in Genero BDL displays using a real table widget, providing a more robust display and interaction than a screen array, while the SCROLLGRID container renders a list of records in separated field cells, providing a replacement for screen arrays using the OPTIONS="-nolist"
- [Review TUI specific features](0417-review-tui-specific-features.md): Some programs use the TUI mode and often exploit all the display possibilities of the language for dumb terminals. These programs should be reviewed when redesigning the application for GUI mode.
- [The default SCREEN window](0418-the-default-screen-window.md): When the first interactive instruction is reached in a Genero BDL program, a default window named SCREEN is created.
- [Specifying WINDOW position and size](0419-specifying-window-position-and-size.md): With Genero BDL in GUI mode, window position and sizes are ignored; in TUI mode, window position and sizes are respected.
- [Front-end configuration tools](0420-front-end-configuration-tools.md): With Genero BDL, use presentation styles instead of front-end specific configuration tools to define widget aspects such as color, borders, fonts, and more.
- [Function key mapping](0421-function-key-mapping.md): When migrating to Genero BDL, special care must be taken for programs that uses function keys greater than F12.
- [Activating form items with DISPLAY](0422-activating-form-items-with-display.md): The methods for enabling or disabling fields and actions differs between Four Js BDS and Genero BDL.
- [Defining keys for WIDGET fields](0423-defining-keys-for-widget-fields.md): The method for binding a keyboard function or control key differ between Four Js BDS and Genero BDL.
- [BEFORE DISPLAY / BEFORE ROW execution order](0424-before-display-before-row-execution-order.md): Genero BDL scenarios define whether the BEFORE DISPLAY or BEFORE ROW block is executed first.
