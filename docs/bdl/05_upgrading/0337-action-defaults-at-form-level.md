---
title: "Action defaults at form level"
source: "fgl-topics/c_fgl_Migrate_to_130_form_action_defaults.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 1.30 upgrade guide > Action defaults at form level"
type: "concept"
---

# Action defaults at form level

> You can define action defaults in forms.

Starting with version 1.30 it is now possible to define action defaults in forms. In previous
versions you had to define a global action default file; this works for defining common
global action attributes, but there is a need to define specific action attributes in some
forms. A typical zoom window may have search and navigation actions, while data input
windows need to define add/delete/update actions instead.

It is now possible to define an action default section in the form file, and you can also load
action defaults with [ui.Form.loadActionDefaults](../15_library-reference/3152-ui-form-loadactiondefaults.md "Load form action defaults.").

## Related links

**Related concepts**  

[ACTION DEFAULTS section](../11_user-interface/1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.")
