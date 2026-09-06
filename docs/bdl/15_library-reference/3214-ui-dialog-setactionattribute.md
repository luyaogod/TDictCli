---
title: "ui.Dialog.setActionAttribute"
source: "fgl-topics/c_fgl_ClassDialog_setActionAttribute.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Dialog class > ui.Dialog methods > ui.Dialog.setActionAttribute"
type: "concept"
---

# ui.Dialog.setActionAttribute

> Set an action attribute to configure a dynamic dialog.

## Syntax

```
setActionAttribute(
   action STRING,
   attrName STRING,
   attrValue STRING )
```

1. name is the name of the dialog action.
2. attrName is the name of the action attribute (as in [.4ad file](../11_user-interface/1603-action-default-attributes-reference-4ad.md "This topic contains all attributes you can define in a .4ad action defaults file.")).
3. attrValue is the value of the action attribute (as in [.4ad file](../11_user-interface/1603-action-default-attributes-reference-4ad.md "This topic contains all attributes you can define in a .4ad action defaults file.")).

## Usage

When creating a dynamic dialog, the `setActionAttribute()` method can be used to
define functional action attributes (such as `VALIDATE=NO`), and decoration
attributes for [default action views](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.")
(such as `TEXT`, `COMMENT`, `IMAGE`).

When defining a static dialog instruction, you can specify action options with the
`ATTRIBUTES()` clause:

```
INPUT BY NAME ...
   ON ACTION show_help ATTRIBUTES( TEXT="Help", VALIDATE=NO )
```

When creating a dynamic dialog, call the `setActionAttribute()` method after
adding an [`"ON ACTION
action-name"` trigger](3184-ui-dialog-addtrigger.md "Adds an event trigger to the dynamic dialog"), to configure the action
attributes:

```
LET d = ui.Dialog.createInputByName(fields)
...
CALL d.addTrigger("ON ACTION show_help")
CALL d.setActionAttribute("show_help","text","Help")
CALL d.setActionAttribute("show_help","validate","no")
```

| Attribute name | Possible values | Description | `ON ACTION` equivalent |
| --- | --- | --- | --- |
| `acceleratorName` | `string` | Defines the first [accelerator key](../11_user-interface/2292-keyboard-accelerator-names.md "Reference for keyboard accelerator names to be used in ACCELERATOR* attributes, and in ON KEY / COMMAND KEY clauses in source dialog code.") for the action. | [`ACCELERATOR`](../11_user-interface/2267-accelerator-action-attribute.md "The ACCELERATOR is an action attribute defining the primary accelerator key for an action.") |
| `comment` | `string` | Defines the bubble hint / tooltip help of the action. | [`COMMENT`](../11_user-interface/2271-comment-action-attribute.md "The COMMENT attribute defines hint for the user about the action.") |
| `contextMenu` | `"yes"`, `"no"`, `"auto"` | Indicates if the action must be shown in the context menu of the form. | [`CONTEXTMENU`](../11_user-interface/2272-contextmenu-action-attribute.md "The CONTEXTMENU attribute defines whether a context menu option must be displayed for an action.") |
| `defaultView` | `"yes"`, `"no"`, `"auto"` | Indicates if the [default action view](../11_user-interface/2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") must appear. | [`DEFAULTVIEW`](../11_user-interface/2273-defaultview-action-attribute.md "The DEFAULTVIEW attribute defines if a default view (a button) must be displayed for a given action.") |
| `image` | `string` | Defines the [icon](../11_user-interface/1586-providing-the-image-resource.md "There are several things you need to know about providing an image resource in a Genero program.") for the action view. | [`IMAGE`](../11_user-interface/2274-image-action-attribute.md "The IMAGE attribute defines the image resource to be displayed for the action.") |
| `text` | `string` | Defines the label for the action view. | [`TEXT`](../11_user-interface/2276-text-action-attribute.md "The TEXT attribute defines the label associated to the action.") |
| `validate` | `"yes"`, `"no"` | Indicates if the action implies [validation of current field input](../11_user-interface/2281-data-validation-at-action-invocation.md "The validate action attribute controls field validation when an action is fired."). | [`VALIDATE`](../11_user-interface/2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.") |

## Related links

**Related concepts**  

[ui.Dialog.setDialogAttribute](3225-ui-dialog-setdialogattribute.md "Set an attribute to configure a dynamic dialog.")
