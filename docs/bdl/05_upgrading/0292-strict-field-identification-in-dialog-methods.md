---
title: "Strict field identification in dialog methods"
source: "fgl-topics/c_fgl_Migrate_to_220_strict_field_ident.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Strict field identification in dialog methods"
type: "concept"
---

# Strict field identification in dialog methods

> Fields referenced in methods of the dialog class must exist in the current dialog, or an error is raised.

Starting with version 2.20.05, dialog class methods like
`ui.Dialog.setFieldTouched()` can now raise a runtime error [-1373](../15_library-reference/4483-genero-bdl-errors.md) if the field specified
does not match a field in the current dialog. Before version 2.20.05, these methods previously
ignored the invalid field specification, making it difficult for the programmer to quickly find
the mistake.

## Related links

**Related concepts**  

[ui.Dialog.setFieldTouched](../15_library-reference/3227-ui-dialog-setfieldtouched.md "Sets the modification flag of the specified field.")
