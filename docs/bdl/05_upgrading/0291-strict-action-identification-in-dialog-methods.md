---
title: "Strict action identification in dialog methods"
source: "fgl-topics/c_fgl_Migrate_to_220_strict_action_ident.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Strict action identification in dialog methods"
type: "concept"
---

# Strict action identification in dialog methods

> Actions referenced in methods of the dialog class must exist in the current dialog, or an error is raised.

Starting with version 2.20.00, dialog class methods like
`ui.Dialog.setActionActive()` can now raise a runtime error [-8089](../15_library-reference/4483-genero-bdl-errors.md) if the action name is
invalid. Before version 2.20, the method ignored the invalid action name, which made it difficult
for the programmer to debug.

## Related links

**Related concepts**  

[ui.Dialog.setActionActive](../15_library-reference/3213-ui-dialog-setactionactive.md "Enabling and disabling dialog actions.")
