---
title: "COMMAND [KEY] block"
source: "fgl-topics/c_fgl_dialog_COMMAND_KEY.html"
breadcrumb: "User interface > Dialog instructions > Multiple dialogs (DIALOG - inside functions) > Using multiple dialogs > DIALOG interaction blocks > COMMAND [KEY] block"
type: "concept"
description: "Use COMMAND [KEY] blocks as global procedural DIALOG action handler to execute a sequence of instructions when the user clicks on a button or presses a specific key. COMMAND defines the text and ..."
---

# COMMAND [KEY] block

Use `COMMAND [KEY]` blocks as global procedural `DIALOG` action
handler to execute a sequence of instructions when the user clicks on a button or presses a specific
key. `COMMAND` defines the text and comment decoration attributes as well as
accelerator keys for a specific action.

`COMMAND` is especially useful when writing TUI programs; however, it's legal to
use such handler when programming new GUI dialogs, to allow focusable action views
(`BUTTON` in form can take the focus if bound to a `COMMAND`
handler).

Declaring a `COMMAND` block in `DIALOG` is
similar to an `ON ACTION` block, except that `COMMAND`
defines an implicit text and comment decoration attribute.
The name of the action will be the command text converted to lowercase letters.
For example, with the following code:

```
COMMAND "Open" "Opens a new file"
```

The name of the action will be "`open`", and the
default decoration text will be "`Open`" with
a capital letter.

Unlike `ON KEY` actions, if no explicit action view is defined in the form,
the default action view will be visible for a `COMMAND` hander (i.e. the automatic
button will appear for this action on the front-end).

[Action defaults](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") will be applied by
using the action name. For explicit action views such as a `BUTTON` in the form
layout, the text/comment defined in the corresponding action default entry will overwrite the values
used in the `COMMAND` handler. When no explicit action view is defined in the form,
the text/comment defined in the program `COMMAND` clause take precedence over action
defaults, to display the [default action
view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.").

Inside `DIALOG` instruction, `COMMAND` blocks
can only be defined as global dialog actions; Sub-dialog
specific `COMMAND` handlers cannot be defined.
When binding a form `BUTTON` to a `COMMAND` handler,
the button can get the focus and will be managed in the tabbing
list, using preferably the [`FIELD
ORDER FORM`](2243-defining-the-tabbing-order.md "Control the order of tabbing through the fields with the TABINDEX attribute.") option.

When using the optional `KEY` clause, `COMMAND` defines
also an implicit accelerator key. The key name must be specified
between parentheses with `COMMAND KEY`:

```
COMMAND KEY (F5) "Open" "Opens a new file"
```

> **Note:**
>
> The `KEY()` clause allows a comma-separated list of keys. Up to four keys can be
> specified. For new developments, consider using a single key, or prefer [`ON ACTION`](1896-on-action-block.md) handlers with a single accelerator
> definition in action defaults.

When using multiple keys in an `COMMAND KEY` clause, the `DIALOG`
instruction will assign the specified keys as accelerators:

```
COMMAND KEY (F5, CONTROL-P, CONTROL-Z) "Open" "Opens a new file"
```

With the above code example, the action name will be "`open`"
and accelerators will be F5, CONTROL-P and CONTROL-Z.

> **Note:**
>
> The keys defined with the `KEY()` clause will take precedence over accelerators
> defined with action defaults corresponding to the action name.

The `COMMAND [KEY]` block specification can also
define a help number with the `HELP` clause,
to display the corresponding text of the current [help file](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.").

```
COMMAND "Open" "Opens a new file" HELP 34
```

## Related links

**Related concepts**  

[Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[BUTTON item type](1684-button-item-type.md "Defines a push-button that can trigger an action.")
