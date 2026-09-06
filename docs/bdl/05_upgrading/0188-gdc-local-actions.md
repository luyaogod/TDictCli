---
title: "(GDC) Local Actions"
source: "fgl-topics/c_fgl_Migrate_to_310_local_actions.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > (GDC) Local Actions"
type: "concept"
---

# (GDC) Local Actions

> The concept of "Local Actions" is now deprecated.

Starting with Genero 3.10, the concept of "Local Actions" is deprecated.

Local Actions are available in the GDC front-end only, to bind local operations to action
views.

For example, the typical copy/cut/paste operations can be bound to ToolBar items by using the
`editcopy`, `editcut`, `editpaste` action names.

Local Actions can be considered as an over-engineered feature. Business applications do not have
to implement basic copy/cut/paste editor actions. The user interface must focus on application
functions (create new record, print order, etc).

## Related links

**Related concepts**  

[Dialog actions](../11_user-interface/2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.")

[Toolbars](../11_user-interface/1855-toolbars.md "Toolbars define a bar of buttons that appears at the top of application forms.")

[Configuring actions](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
