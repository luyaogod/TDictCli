---
title: "Action and field activation"
source: "fgl-topics/c_fgl_Migrate_to_130_action_field_activation.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 1.30 upgrade guide > Action and field activation"
type: "concept"
---

# Action and field activation

> Dialog methods can be used to control action and field activation.

Version 1.30 provides dialog methods to control action and field activation:

- [ui.Dialog.setActionActive(
  action-name, TRUE/FALSE )](../15_library-reference/3213-ui-dialog-setactionactive.md "Enabling and disabling dialog actions.")
- [ui.Dialog.setFieldActive(
  field-name, TRUE/FALSE )](../15_library-reference/3226-ui-dialog-setfieldactive.md "Enable and disable form fields.")

Previous versions allowed you to modify directly the 'active' attribute of the underlying DOM
node in the AUI tree. This is now forbidden: it is mandatory to use the methods to enable/disable
action or fields. The dialog will synchronize the 'active' attribute in the AUI tree based on the
value passed to the methods and depending on the context (some actions or fields can be
automatically disabled).
