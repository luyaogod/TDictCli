---
title: "MenuAction close no longer created by default"
source: "fgl-topics/c_fgl_Migrate_to_230_menuaction_close.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.30 upgrade guide > MenuAction close no longer created by default"
type: "concept"
---

# MenuAction close no longer created by default

> The close action is no longer created by default in MENU dialog.

Before version 2.30, a `close` MenuAction was created by default for MENU dialogs.
This action node is no longer created, except if you have a COMMAND KEY(INTERRUPT) in the MENU,
or if you have your own user action handler ON ACTION close, of course. You must take this change
into account if you are manipulating the AUI tree with om classes in MENUs.

## Related links

**Related concepts**  

[Ring menus (MENU)](../11_user-interface/1904-ring-menus-menu.md "The MENU instruction implements a list of options the end user can choose from.")

[The DomNode class](../15_library-reference/3295-the-domnode-class.md "The om.DomNode class provides methods to manipulate a DOM node of a data tree.")
