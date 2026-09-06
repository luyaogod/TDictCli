---
title: "PROMPT instruction configuration"
source: "fgl-topics/c_fgl_prompt_006.html"
breadcrumb: "User interface > Dialog instructions > Prompt for values (PROMPT) > Using simple prompt inputs > PROMPT instruction configuration"
type: "concept"
description: "HELP option The HELP clause specifies the number of a help message to display if the user invokes the help while executing the instruction. The predefined \"help\" action is automatically created by the ..."
---

# PROMPT instruction configuration

## HELP option

The `HELP` clause specifies the number of a [help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to display if the user invokes the
help while executing the instruction. The predefined "help" action is automatically created
by the runtime system. You can bind [action
views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the "help" action.

The `HELP` clause overrides the `HELP` attribute.

## ACCEPT option

The `ACCEPT` attribute can be set to `FALSE` to avoid the
automatic creation of the "accept" default action.

## CANCEL option

The `CANCEL` attribute can be set to `FALSE` to avoid the
automatic creation of the cancel default action. This is useful for example when you only
need a validation action (accept), or when you want to write a specific cancellation procedure,
by using [`EXIT INPUT`](2035-exit-input-instruction.md).

If the `CANCEL=FALSE` option is set, no [close](2286-multilevel-action-conflicts.md) action will be created, and
you must write an `ON ACTION close` control block to create an explicit action.

## Using TTY attributes

The `ATTRIBUTES` clause can define TTY attributes such as colors
(`RED`, `GREEN`), and `REVERSE`. These attributes will
be used during the dialog execution.

> **Important:**
>
> In GUI mode, form elements can also be decorated with
> presentation styles. Pay attention to the specific rules that apply when [combining TTY attributes and presentation
> styles](1619-combining-tty-and-style-attributes.md "TTY attributes can define style attribute equivalents such as the text color. Different precedence rules apply, depending on the TTY attribute specification.").

## Related links

**Related concepts**  

[Syntax of PROMPT instruction](1890-syntax-of-prompt-instruction.md "The PROMPT statement assigns a user-supplied value to a variable.")
