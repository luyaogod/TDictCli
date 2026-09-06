---
title: "Key labels versus action defaults"
source: "fgl-topics/c_fgl_Mig0000_015.html"
breadcrumb: "Upgrading > Migrating from Four Js BDS to Genero BDL > User interface topics > Key labels versus action defaults"
type: "concept"
---

# Key labels versus action defaults

> Genero BDL introduces the ON ACTION block to define actions. While the ON KEY block is still supported, it comes with limitations.

In Four Js Business Development Suite (BDS), labels can be defined for keys such as
`accept`, `F10` or `Control-Z`. With this feature, it
is possible to easily decorate `ON KEY` or `COMMAND KEY` blocks with a
button in the action panel.

Key labels can be specified at different levels:

- at a global level with FGLPROFILE settings,
- at the program level with the `fgl_setkeylabel()` function,
- at the form level with the `KEYS` section,
- at dialog level with the `fgl_dialog_setkeylabel()` function, and
- at the field level with the `KEY="label"` attribute.

For more details, see [Setting action key labels](../11_user-interface/2296-setting-action-key-labels.md "Labels can be defined to decorate buttons controlled by ON KEY / COMMAND KEY action handlers.").

With Genero Business Development Language (BDL), interaction statements can define actions with
the `ON ACTION` blocks. These action handlers are more abstract than `ON
KEY`: You identify an action by a name, while decoration is defined in form files (
`ACTION DEFAULTS` section) or in global configuration files
(.4ad files).

When adapting your code for Genero, you are free to use the traditional `ON KEY`
blocks or the new `ON ACTION` blocks. Genero still supports the key label settings as
in Four Js BDS. However, key label settings will overwrite action defaults settings. Additionally,
if the name of the key specified in the `ON KEY` clause does not only contain
alphanumeric characters (such as `Control-Z`), it will not be possible to define
action defaults attributes for these action handlers, as action names must be simple identifiers.
This is also true for Menu `COMMAND` labels, for example with `COMMAND "Exit
program"`.

## Related links

**Related concepts**  

[Binding action views to action handlers](../11_user-interface/2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?")

[ACTION DEFAULTS section](../11_user-interface/1711-action-defaults-section.md "The ACTION DEFAULTS section defines local action view default attributes for the form elements.")

[User interface basics](../11_user-interface/1508-user-interface-basics.md "This section introduces to the foundation of the Genero user interface.")
