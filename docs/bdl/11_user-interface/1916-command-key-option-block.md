---
title: "COMMAND [KEY()] \"option\" block"
source: "fgl-topics/c_fgl_menus_014.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU interaction blocks > COMMAND [KEY()] \"option\" block"
type: "concept"
description: "Syntax COMMAND [ KEY( key-name ) ] \" option-name \" instruction [...] Usage The COMMAND [ KEY( key-name ) ] \" option-name \" clause defines a menu action handler with a set of instructions to be ..."
---

# COMMAND [KEY()] "option" block

## Syntax

```
COMMAND [KEY(key-name)] "option-name"
   instruction [...]
```

## Usage

The `COMMAND [KEY(key-name)]
"option-name"` clause defines a menu action handler with a set of
instructions to be executed when an action is invoked. The option text
(`option-name`), converted to lowercase letters, defines the name
of the action.

For example, when defining:

```
COMMAND "Hello"
```

The name of the action will be "`hello`" (not "`Hello`" with a
capital H).

When used with the `KEY()` clause, the command specifies both `accelerator keys` and an option
text.

> **Note:**
>
> The `KEY()` clause allows a comma-separated list of keys. Up to four keys can be
> specified. For new developments, consider using a single key, or prefer [`ON ACTION`](1896-on-action-block.md) handlers with a single accelerator
> definition in action defaults.

Action defaults will be applied by using the action name defined by the option text (converted to
lowercase).

In TUI mode, actions created with `COMMAND [KEY]` do not get accelerators from
action defaults; only actions defined with `ON ACTION` will get accelerators of
action defaults.

Explicit action views defined in the form (`BUTTON` in layout,
`TOPMENU` or `TOOLBAR` items) will get all action defaults associated
to the menu command, while [default action
views](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") will be decorated with the menu option text and comment specified in the program (this
means that the `TEXT` and `COMMENT` attributes of the corresponding
action defaults entry are not used for the default action views). However, other attributes such as
the `IMAGE` will also be applied to default action views.

For example, when defining:

```
COMMAND "Hello" "This is the Hello option"
```

The
name of the action will become "`hello`", the default action view button text will be
"`Hello`", and the button hint will be "`This is the Hello option`",
even if an action default defines a different text or comment for the "`hello`"
action. If the corresponding action default defines an `IMAGE` icon, it will display
in the default action view button.

> **Note:**
>
> The keys defined with the `KEY()` clause will take precedence over accelerators
> defined with action defaults corresponding to the action name.

The first letter of the display text of a `COMMAND` menu clause can be used as
default accelerator, if no other accelerator is defined by a `KEY()` clause or by
action defaults:

- When the menu option is not rendered as a [default view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.") and is bound to an action
  view like a toolbar button, the first letter of the `COMMAND` option will define the
  [default accelerator](2265-defining-keyboard-accelerators-for-actions.md), except if
  several menu `COMMAND` options start with the same letter, or if the letter is used
  as accelerator by another command.
- When the menu option is rendered as a [default view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it."), and the first letter of a menu `COMMAND` option is not used by
  other menu option labels, pressing the key corresponding to that letter will execute that action. If
  several menu `COMMAND` options start with the same letter and the user presses that
  key, the front-end shifts the focus from button to button. To fire the action, the user must press
  ENTER.

In a `MENU` instruction, the alternative to `COMMAND [KEY]
"option-name"` can be the `ON ACTION action-name` clause,
to write abstract code without having the decoration (option name, comment and accelerators) in the
program code.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")
