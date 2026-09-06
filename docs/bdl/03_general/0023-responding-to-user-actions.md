---
title: "Responding to user actions"
source: "fgl-topics/c_fgl_intro_BDL_014.html"
breadcrumb: "General > Introduction to Genero BDL programming > Genero BDL concepts > Responding to user actions"
type: "concept"
---

# Responding to user actions

> Clicking a form button or pressing a key triggers actions that can invoke the execution of program of code called action handlers. Form elements that can trigger actions are called action views.

Action handlers are defined in interactive statements with the `ON ACTION` clause.
The code defined in action handler blocks is executed when an action is fired. Action
objects are created and linked to action views when such `ON ACTION`
handlers are seen by the runtime system. Common action handlers, such as
`accept` (dialog validation) and `cancel` (dialog
cancellation), are created automatically in accordance with the interactive
instruction.

By configuring action defaults, you define the default decoration attributes (text,
image) and functional attributes (accelerator keys, context menu display) for the action
views associated with actions.

## Related links

**Related concepts**  

[Dialog actions](../11_user-interface/2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.")
