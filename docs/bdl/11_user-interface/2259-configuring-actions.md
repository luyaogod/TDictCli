---
title: "Configuring actions"
source: "fgl-topics/c_fgl_prog_dialogs_config_actions.html"
breadcrumb: "User interface > User interface programming > Dialog actions > Configuring actions"
type: "concept"
---

# Configuring actions

> Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.

Action attributes define attributes for actions, including decoration such as text,
icon, comment, as well as keyboard accelerator (Ctrl-?, function keys), and also semantics such as
current field validation control when an action is fired.

The action attributes can be defined in different ways:

1. Common action attributes can be centralized in a global action defaults file with the
   .4ad extension.
2. Form-specific action attributes can be defined for all action views of a form definition file,
   with the `ACTION DEFAULTS` section, including functional attributes for the dialog
   controlling the form.
3. Form-item specific action view attributes (decoration only) can be defined directly at the item
   level (labels, icons, comments).
4. Default action views attributes can be defined at dialog level in programs, with the
   `ATTRIBUTES()` clause of `ON ACTION` handlers, and changed dynamically
   with methods such as `ui.Dialog.setActionText()`. This applies only to default action
   views.

Action attributes do not only define action view decoration; it is possible to define the
semantics of an action, for example by using the `VALIDATE` action default attribute.
Functional attributes take effect for a given action when the dialog implementing the action handler
becomes active.

Action attributes are particularly important for rendering the default action view (when there
is no explicit action view defined in the form). This is typically the case when no form is
associated with the dialog.

Action attributes can be defined with action defaults. Common action defaults are defined in a
global action defaults (.4ad) file, while form specific actions are define
within the `ACTION DEFAULTS` section of form files.

If a dialog is not attached to a specific form such as an independent `MENU`,
define the action attributes with the `ATTRIBUTES` clause of `ON ACTION`
handlers, to render the default view and configure the action semantics. Attributes defined by
`ON ACTION action-name ATTRIBUTES()` will only be applied
to the default action view: The elements in the forms do not get decoration attributes defined by
dialog action handlers.

The final decoration and functional attribute values are set in this order of precedence:

1. The decoration attributes defined in the action view form element definition itself (such as a
   `BUTTON`) - For default action views, the specific attributes can be defined in the
   `ATTRIBUTES` clause of an `ON ACTION` handler, or dynamically with
   dialog methods such as `ui.Dialog.setActionText()`.
2. Functional attributes in the `ATTRIBUTES` clause of an `ON ACTION`
   handler (such as `VALIDATE=NO`).
3. Decoration and functional attributes defined for the action in the `ACTION
   DEFAULTS` section of the current form.
4. Decoration and functional attributes defined for the action in the global action defaults file
   (.4ad).

The global action defaults can be loaded at runtime with [ui.Interface.loadActionDefaults](../15_library-reference/3112-ui-interface-loadactiondefaults.md "Load the default action defaults file."), and the form-specific `ACTION
DEFAULTS` section can be loaded with [ui.Form.loadActionDefaults](../15_library-reference/3152-ui-form-loadactiondefaults.md "Load form action defaults.").
These solutions are typically used in a migration process, to get action views decoration without
modifying existing .per forms.

The syntax for defining action attributes depends on the context where the action attributes are
defined:

- In the .4ad file, the syntax follows XML standards, as defined in [Action default attributes reference (.4ad)](1603-action-default-attributes-reference-4ad.md "This topic contains all attributes you can define in a .4ad action defaults file.").
- In the .per files, the syntax follows the form specification file
  attributes, as defined in `ACTION DEFAULTS` section.
- In the .4gl files (in dialog action handlers), the syntax follows the
  language syntax, as defined in [ON ACTION block](1918-on-action-block.md).

## Example

Consider the following parts of code related to the same action definition, namely "print":

1. A `BUTTON` item defined in the form specification file:

```
ATTRIBUTES
 BUTTON b1: print, TEXT="Print item";
END
```

2. A dialog instruction with code defining the `ON ACTION`
handler with an `ATTRIBUTES` clause:

```
DIALOG ...
   ...
   ON ACTION print ATTRIBUTES(VALIDATE=NO)
      ...
```

3. The form `ACTION DEFAULTS` section
defining:

```
form.per:
ACTION DEFAULTS
  ACTION print (IMAGE="printer",
                COMMENT="Print the order",
                ACCELERATOR=Control-P,
                CONTEXTMENU=NO)
END
```

4. A global .4ad action defaults file defining:

```
<ActionDefaultList>
  <ActionDefault name="print" text="Print" image="smiley" />
</ActionDefaultList>
```

When the dialog executes, the "print" action will get the following
functional attributes:

- `validate = "no"` - from the dialog `ON ACTION` handler
- `acceleratorName = "control-p"` - from the form `ACTION DEFAULTS`
  section
- `contextMenu = "no"` - from the form `ACTION DEFAULTS` section

The form button (the action view) will get the following decoration attribute values:

- `text = "Print item"` - from the `BUTTON` form item
- `image = "printer"` - from the form `ACTION DEFAULTS` section
- `comment = "Print the order"` - from the form `ACTION DEFAULTS` section

## Related links

**Related concepts**  

[Action defaults files](1600-action-defaults-files.md "Action defaults files allow to centralize action configuration parameters such as text, icon, accelerators and behavior options in XML format.")

## Child topics

- [Action attributes context usage](2260-action-attributes-context-usage.md)
- [Using attributes of action defaults](2261-using-attributes-of-action-defaults.md)
- [Dialog action handler attributes](2262-dialog-action-handler-attributes.md)
- [Configuring default action views dynamically](2263-configuring-default-action-views-dynamically.md)
- [Text attribute shows default action view](2264-text-attribute-shows-default-action-view.md)
- [Defining keyboard accelerators for actions](2265-defining-keyboard-accelerators-for-actions.md)
- [Action attributes list](2266-action-attributes-list.md)
