---
title: "New compiler warning to avoid action shadowing"
source: "fgl-topics/c_fgl_Migrate_to_240_on_action_warning.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > New compiler warning to avoid action shadowing"
type: "concept"
---

# New compiler warning to avoid action shadowing

> Prevent the same action name at different levels of ON ACTION handlers in a dialog.

The fglcomp compiler of version 2.40 will now print warning [-8409](../15_library-reference/4483-genero-bdl-errors.md), if a dialog block defines ON
ACTION handlers at different levels with the same action name.

It is not good practice to use the same action names at different levels of a dialog. For
example, you can define several ON ACTION INFIELD handlers using the action name "zoom", but do
not define an ON ACTION zoom at the sub-dialog or dialog level.

If the warning occurs during compilation, modify your code in order to use specific action
names at each level, and do not forget to rename the actions of the corresponding action views in
the forms.

## Related links

**Related concepts**  

[Multilevel action conflicts](../11_user-interface/2286-multilevel-action-conflicts.md "Multilevel action conflicts")

[Command reference](../13_programming-tools/2512-command-reference.md "Command line tools provided by FGLGWS packages.")
