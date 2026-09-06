---
title: "DISPLAY ARRAY stops if container is hidden"
source: "fgl-topics/c_fgl_Migrate_to_400_dialog_stop.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 4.00 upgrade guide > DISPLAY ARRAY stops if container is hidden"
type: "concept"
---

# DISPLAY ARRAY stops if container is hidden

> The DISPLAY ARRAY dialog will now stop, if there is no record list container visible.

Starting with version 4.00.03, which includes the fix FGL-5554, a `DISPLAY ARRAY`
will automatically terminate, if its corresponding record list container (such as a
`TABLE`, `TREE`, `SCROLLGRID`) is hidden by program
code or by a `HIDDEN@screen-size` attribute. The same behavior
exists with a `DIALOG` block containing a `DISPLAY ARRAY`.

Before this fix, the behavior of the `DISPLAY ARRAY` dialog under these conditions
was undefined.

A dialog like `DISPLAY ARRAY` is a controller that uses a
model (the `DYNAMIC ARRAY` variable), and a view (a
`TABLE` in a form or a set of form fields). If there is no view to work
with, the dialog cannot continue.

It's worth noting that in former versions that did not include this fix, a simple
`INPUT` or `CONSTRUCT` dialog would also stop, and an `INPUT
ARRAY` would produce the error -8092: Dialogs allowing field input need to have at least one
active and visible field.

If parts of a form need to be hidden during a dialog's execution, ensure that at least one field
or record list container remains available.This will allow the dialog to continue and the end user
to maintain control.

## Related links

**Related concepts**  

[The model-view-controller paradigm](../11_user-interface/2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.")

[What are dialog controllers?](../11_user-interface/2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling.")

[ui.Dialog.setFieldActive](../15_library-reference/3226-ui-dialog-setfieldactive.md "Enable and disable form fields.")

[ui.Form.setElementHidden](../15_library-reference/3156-ui-form-setelementhidden.md "Show or hide form elements.")

[ui.Form.setFieldHidden](../15_library-reference/3161-ui-form-setfieldhidden.md "Show or hide a form field.")
