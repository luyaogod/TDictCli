---
title: "ON ACTION block"
source: "fgl-topics/c_fgl_dialog_ON_ACTION_2.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT interaction blocks > ON ACTION block"
type: "concept"
description: "Syntax ON ACTION action-name instruction [...] Usage The ON ACTION action-name blocks execute a sequence of instructions when the user triggers a specific action. A typical action handler block looks ..."
---

# ON ACTION block

## Syntax

```
ON ACTION action-name
   instruction [...]
```

## Usage

The `ON ACTION action-name` blocks execute a sequence of
instructions when the user triggers a specific action.

A typical action handler block looks like this:

```
  ON ACTION action-name
     instruction
     ...
```

Action blocks are bound by name to action views (like buttons) in the current form. Action views
can be `BUTTON`, `TOOLBAR` buttons, or `TOPMENU`
options, and if no explicit action view is defined, actions are rendered with a default action view,
depending on the type of front-end.

This example defines an action block to open a typical zoom window and let the user select a
customer record:

```
   ON ACTION zoom 
      CALL zoom_customers() RETURNING st, rec.cust_id, rec.cust_name
```

In a dialog handling user input such as `INPUT`, `INPUT ARRAY`
and `CONSTRUCT`, if an action is
specific to a field, add the `INFIELD` clause to have the action automatically
enabled when the corresponding field gets the
focus:

```
   ON ACTION zoom INFIELD cust_city
      CALL zoom_cities() RETURN st, rec.cust_city
```

In most cases actions are decorated with action defaults in form files, but there can be cases
where the `ON ACTION` handler needs to define its own attributes at the program
level. This can be done by adding the `ATTRIBUTES()` clause of
`ON
ACTION`:

```
   ON ACTION custinfo ATTRIBUTES(TEXT="More details", IMAGE="info")
      CALL show_customer_info()
```

For more details about action handlers, and action configuration, see [Dialog actions](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.").

## Related links

**Related concepts**  

[Configuring actions](2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.")

[Action handling basics](2254-action-handling-basics.md "This topic describes the basic concepts of dialog actions.")
