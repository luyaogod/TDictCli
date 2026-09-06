---
title: "COMMAND KEY() block"
source: "fgl-topics/c_fgl_menus_015.html"
breadcrumb: "User interface > Dialog instructions > Ring menus (MENU) > Using ring menus > MENU interaction blocks > COMMAND KEY() block"
type: "concept"
description: "Syntax COMMAND KEY ( key-name ) instruction [...] Usage The COMMAND KEY( key-name ) block (without an option text) defines a menu action handler with a set of instructions to be executed when an ..."
---

# COMMAND KEY() block

## Syntax

```
COMMAND KEY (key-name)
   instruction [...]
```

## Usage

The `COMMAND KEY(key-name)` block (without an option text)
defines a menu action handler with a set of instructions to be executed when an action is
invoked.

The `KEY()` clause defines one to four `accelerator keys` separated by a
comma. The specified key name must be one of [the
virtual keys](2293-on-key-virtual-keys.md "Virtual keys are the key names that can be used in program instructions such as ON KEY and COMMAND KEY.").

> **Note:**
>
> The `KEY()` clause allows a comma-separated list of keys. Up to four keys can be
> specified. For new developments, consider using a single key, or prefer [`ON ACTION`](1896-on-action-block.md) handlers with a single accelerator
> definition in action defaults.

While a `COMMAND KEY(key-name)
"option-name"` (with option text) defines the name of the action with the
option text (converted to lowercase), a `COMMAND KEY(key-name)`
(without option text), defines the action name from the last key in the
`KEY()` list, converted to lowercase letters. For example, with `COMMAND
KEY(F10,F12,Control-Z)`, the name of the action will be "`control-z`".

Action defaults will be applied by using the key name of the `KEY()` clause. With
a list of keys, the last key name will be used to apply action defaults, because it defines the
action name.

> **Note:**
>
> The keys defined with the `KEY()` clause will take precedence over accelerators
> defined with action defaults corresponding to the action name.

By default, `COMMAND KEY(key-name)` actions are not decorated
with a [default action view](2258-default-action-views.md "A default action view is created to render an action handler, when no explicit action view exists for it.").
However, by defining the `TEXT` attribute within action defaults, the default action
view button will be visible. This allows you to decorate existing `COMMAND
KEY(key-name)` clauses with graphical buttons without changing the
program code.

To write abstract code without decoration in your programs, use the `ON ACTION` clause instead of `COMMAND
[KEY]`, except if the action view must get the focus.

In TUI mode, actions created with `COMMAND [KEY]` do not get accelerators from
action defaults; Only actions defined with `ON ACTION` will get accelerators of
action defaults.

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")
