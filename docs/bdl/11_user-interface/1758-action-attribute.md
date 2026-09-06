---
title: "ACTION attribute"
source: "fgl-topics/c_fgl_FSFAttributes_ACTION.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > ACTION attribute"
type: "concept"
---

# ACTION attribute

> The ACTION attribute defines the action associated with the form item.

## Syntax

```
ACTION = action-name
```

1. action-name is an identifier that defines the name of
   the action to be sent.

## Usage

The `ACTION` attribute defines the name of the action to be sent to
the program when the user activates the form item.

This attribute can for example be used in a `BUTTONEDIT` field to identify the
corresponding action handle to be executed in the program when the button is pressed. The [Dialog actions](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") chapter describes how actions can be implemented in
dialogs.

The action name can be prefixed with a sub-dialog identifier and/or field name, to define a
qualified action view (see [Binding action views to action handlers](2280-binding-action-views-to-action-handlers.md "How are action views of the forms bound to action handlers in the program code?") for more
details).

## Example

```
BUTTONEDIT f001 = customer.state, ACTION = print;
```

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")
