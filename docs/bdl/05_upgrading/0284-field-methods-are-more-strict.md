---
title: "Field methods are more strict"
source: "fgl-topics/c_fgl_Migrate_to_220_strict_field_methods.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Field methods are more strict"
type: "concept"
---

# Field methods are more strict

> Dialog class methods are more strict regarding form field names.

Starting with Genero 2.20 (or when using multiple dialogs in 2.11.08 and higher), DIALOG class
methods such as `setFieldActive()` need the correct field specification with the
screen-record name prefix, if the field was explicitly bound with the `FROM` clause
of `INPUT` or `INPUT ARRAY`.

In prior versions, the field was found by these methods even if the prefix was invalid.
(Actually, the prefix was just ignored and only the fieldname was used.)

## Related links

**Related concepts**  

[Identifying fields in ui.Dialog methods](../15_library-reference/3239-identifying-fields-in-ui-dialog-methods.md "Identifying fields in ui.Dialog methods")
