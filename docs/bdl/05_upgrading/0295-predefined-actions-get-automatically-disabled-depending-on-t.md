---
title: "Predefined actions get automatically disabled depending on the context"
source: "fgl-topics/c_fgl_Migrate_to_220_predef_actions.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.20 upgrade guide > Predefined actions get automatically disabled depending on the context"
type: "concept"
---

# Predefined actions get automatically disabled depending on the context

> Dialogs will automatically disable some predefined actions, if it makes no sense to trigger the action in the current context.

Starting with version 2.20, (or with version 2.10 when `FGL_USENDIALOG=1`), the
dialogs will automatically disable some predefined actions if it makes no sense to trigger the
action in the current context. For example, during an `INPUT ARRAY`, if there are no
rows to remove, the predefined delete action will be disabled automatically. Similarly, the insert
and append actions get disabled when the array is full (this can happen with static arrays or when
using the `MAXCOUNT` attribute). The predefined actions will also be disabled if you
overwrite them with your own `ON ACTION` handler.

## Related links

**Related concepts**  

[Predefined actions](../11_user-interface/2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")
