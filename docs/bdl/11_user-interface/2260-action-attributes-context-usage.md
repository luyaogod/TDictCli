---
title: "Action attributes context usage"
source: "fgl-topics/c_fgl_prog_dialogs_action_attributes_context.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions > Action attributes context usage"
type: "concept"
description: "Action attributes are used to configure functional and decoration properties of actions. The table below lists the possible action attributes and indicates in what context they can be defined. Table ..."
---

# Action attributes context usage

Action attributes are used to configure functional and decoration properties of actions.
The table below lists the possible action attributes and indicates in what context they can
be defined.

| Attribute | Context |  |  |  |  | Form action view | Dialog action handler | Form action defaults section | Global action defaults file (.4ad) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `ACCELERATOR`See [ACCELERATOR action attribute](2267-accelerator-action-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action."). | No | Yes | Yes | Yes |  |  |  |  |  |
| `ACCELERATOR2`See [ACCELERATOR2 action attribute](2268-accelerator2-action-attribute.md "The ACCELERATOR2 is an action attribute defining the secondary accelerator key for an action."). | No | No | Yes | Yes |  |  |  |  |  |
| `ACCELERATOR3`See [ACCELERATOR3 action attribute](2269-accelerator3-action-attribute.md "The ACCELERATOR3 is an action attribute defining the third accelerator key for an action."). | No | No | Yes | Yes |  |  |  |  |  |
| `ACCELERATOR4`See [ACCELERATOR4 action attribute](2270-accelerator4-action-attribute.md "The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action."). | No | No | Yes | Yes |  |  |  |  |  |
| `COMMENT`See [COMMENT action attribute](2271-comment-action-attribute.md "The COMMENT attribute defines hint for the user about the action."). | Yes | Yes | Yes | Yes |  |  |  |  |  |
| `CONTEXTMENU`See [CONTEXTMENU action attribute](2272-contextmenu-action-attribute.md "The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action."). | No | Yes | Yes | Yes |  |  |  |  |  |
| `DEFAULTVIEW`See [DEFAULTVIEW action attribute](2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action."). | No | Yes | Yes | Yes |  |  |  |  |  |
| `IMAGE`See [IMAGE action attribute](2274-image-action-attribute.md "The IMAGE attribute defines the image resource to be displayed for the action."). | Yes | Yes | Yes | Yes |  |  |  |  |  |
| `ROWBOUND`See [ROWBOUND action attribute](2275-rowbound-action-attribute.md "The ROWBOUND attribute defines if the action is related to the row context of a record list."). | No | Yes (only for list dialogs) | No | No |  |  |  |  |  |
| `TEXT`See [TEXT action attribute](2276-text-action-attribute.md "The TEXT attribute defines the label associated to the action."). | Yes | Yes | Yes | Yes |  |  |  |  |  |
| `VALIDATE`See [VALIDATE action attribute](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action."). | No | Yes (only for input dialogs) | Yes | Yes |  |  |  |  |  |
