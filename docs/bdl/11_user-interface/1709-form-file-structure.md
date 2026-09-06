---
title: "Form file structure"
source: "fgl-topics/c_fgl_FormSpecFiles_struct.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form file structure"
type: "concept"
---

# Form file structure

> A form specification file is defined by a set of sections.

The sections of a form specification file must appear in the following order:

1. [SCHEMA section](1710-schema-section.md "Defines the database schema file to be used to compile the form.")
2. [ACTION DEFAULTS section](1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.")
3. [TOPMENU section](1712-topmenu-section.md "The TOPMENU section defines a pull-down menu with options that are bound to actions.")
4. [TOOLBAR section](1713-toolbar-section.md "The TOOLBAR section defines a toolbar with buttons that are bound to actions.")
5. [LAYOUT section](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") (or [SCREEN section](1714-screen-section.md "The SCREEN section defines the form layout for TUI mode forms.") for TUI mode)
6. [TABLES section](1726-tables-section.md "Defines the list of database tables referenced by form field definitions.")
7. [ATTRIBUTES section](1727-attributes-section.md "The ATTRIBUTES section describes properties of elements used in the form.")
8. [INSTRUCTIONS section](1751-instructions-section.md "The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties.")
9. [KEYS section](1752-keys-section.md "The KEYS section can be used to define default key labels for the current form.") (legacy)

Each section must begin with the keyword for which it is named, only the `LAYOUT`
section is mandatory.

## Related links

**Related concepts**  

[Examples](1853-examples.md "Form definition (.per) examples.")

## Child topics

- [SCHEMA section](1710-schema-section.md): Defines the database schema file to be used to compile the form.
- [ACTION DEFAULTS section](1711-action-defaults-section.md): The ACTION DEFAULTS section defines local action view default attributes for the form elements.
- [TOPMENU section](1712-topmenu-section.md): The TOPMENU section defines a pull-down menu with options that are bound to actions.
- [TOOLBAR section](1713-toolbar-section.md): The TOOLBAR section defines a toolbar with buttons that are bound to actions.
- [SCREEN section](1714-screen-section.md): The SCREEN section defines the form layout for TUI mode forms.
- [LAYOUT section](1715-layout-section.md): The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.
- [TABLES section](1726-tables-section.md): Defines the list of database tables referenced by form field definitions.
- [ATTRIBUTES section](1727-attributes-section.md): The ATTRIBUTES section describes properties of elements used in the form.
- [INSTRUCTIONS section](1751-instructions-section.md): The INSTRUCTIONS section is used to define screen arrays, non-default screen records, and global form properties.
- [KEYS section](1752-keys-section.md): The KEYS section can be used to define default key labels for the current form.
