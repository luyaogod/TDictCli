---
title: "Action default attributes reference (.4ad)"
source: "fgl-topics/r_fgl_action_defaults_files_reference.html"
breadcrumb: "User interface > Form definitions > Action defaults files > Action default attributes reference (.4ad)"
type: "reference"
---

# Action default attributes reference (.4ad)

> This topic contains all attributes you can define in a .4ad action defaults file.

| Attribute | Description |
| --- | --- |
| `name = "action-name"` | This attribute identifies the action. |
| `text = "action-label"` | The default label to be displayed in action views (typically, the text of buttons).See also: [TEXT attribute](1828-text-attribute.md "The TEXT attribute defines the label associated with a form item.") |
| `comment = "action-comment"` | The default help text for this action (typically, displayed as bubble help).See also: [COMMENT attribute](1769-comment-attribute.md "The COMMENT attribute defines a hint for the user about the form element.") |
| `image = "action-icon"` | The default image file to be displayed in the action view.See also: [IMAGE attribute](1785-image-attribute.md "The IMAGE attribute defines the image resource to be displayed for the form item.") |
| `acceleratorName = "key-name"` | The default accelerator key that can trigger the action, as defined in [Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.").See also: [ACCELERATOR attribute](1754-accelerator-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action.") |
| `acceleratorName2 = "key-name"` | The second default accelerator key that can trigger the action, as defined in [Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.").See also: [ACCELERATOR2 attribute](1755-accelerator2-attribute.md "The ACCELERATOR2 is an action attribute defining the secondary accelerator key for an action.") |
| `acceleratorName3 = "key-name"` | The third default accelerator key that can trigger the action, as defined in [Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.").See also: [ACCELERATOR3 attribute](1756-accelerator3-attribute.md "The ACCELERATOR3 is an action attribute defining the third accelerator key for an action.") |
| `acceleratorName4 = "key-name"` | The fourth default accelerator key that can trigger the action, as defined in [Keyboard accelerator names](2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.").See also: [ACCELERATOR4 attribute](1757-accelerator4-attribute.md "The ACCELERATOR4 is an action attribute defining the fourth accelerator key for an action.") |
| `defaultView = {"yes"\|"no"\|"auto"}` | Defines whether the front-end must show the default action view (buttons in action panel).Values can be:`"no"` the default action view is never visible.`"yes"` the default action view is always visible, if the action is visible (`ui.Dialog.setActionHidden`).`"auto"` the default action view is visible if no other action view is explicitly defined and the action is visible (`ui.Dialog.setActionHidden`).The default is `"auto"`.See also: [DEFAULTVIEW attribute](1773-defaultview-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.") |
| `contextMenu = {"yes"\|"no"\|"auto"}` | Defines whether the front-end must render the action in the default context menu.Values can be:`"no"` the context menu option is never visible.`"yes"` the context menu option is always visible, if the action is visible (`ui.Dialog.setActionHidden`).`"auto"` the context menu option is visible if no other action view is explicitly defined and the action is visible (`ui.Dialog.setActionHidden`).The default is `"yes"`.See also: [CONTEXTMENU attribute](1768-contextmenu-attribute.md "The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action.") |
| `validate = "no"` | Defines the behavior of data validation when the action is invoked.Values can be:`"no"` no data validation is done (field text only available in input buffer).By default, data validation is driven by the dialog mode (`UNBUFFERED` or default mode).For more details, see [Data validation at action invocation](2281-data-validation-at-action-invocation.md "The validate action attribute controls field validation when an action is fired.").See also: [VALIDATE attribute](1839-validate-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.") |

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[ON ACTION block](1918-on-action-block.md "ON ACTION block")
