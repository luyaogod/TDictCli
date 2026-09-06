---
title: "CONSTRUCT instruction configuration"
source: "fgl-topics/c_fgl_Construct_008.html"
breadcrumb: "User interface > Dialog instructions > Query by example (CONSTRUCT) > Using query by example > CONSTRUCT instruction configuration"
type: "concept"
description: "This section describes the options that can be specified in the ATTRIBUTES clause of the CONSTRUCT instruction. The options of the ATTRIBUTES clause override all default attributes and temporarily ..."
---

# CONSTRUCT instruction configuration

This section describes the options that can be specified in the `ATTRIBUTES`
clause of the `CONSTRUCT` instruction. The options of the `ATTRIBUTES`
clause override all default attributes and temporarily override any display attributes that the
`OPTIONS` or the `OPEN WINDOW` statement specified for these fields.
With the `CONSTRUCT` statement, the `INVISIBLE` attribute is
ignored.

## NAME option

The `NAME` attribute can be used to name the `CONSTRUCT` dialog.
This attribute is used to identify the `CONSTRUCT` dialog.

## HELP option

The `HELP` clause specifies the number of a [help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to display if the user invokes the help in
the `CONSTRUCT` dialog. The predefined 'help' action is automatically created by the
runtime system. You can bind [action views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the
'help' action.

The `HELP` clause overrides the `HELP` attribute.

## FIELD ORDER FORM option

By default, the tabbing order is defined by the [variable
binding list](1935-variable-binding-in-input.md) in the instruction description. You can control the tabbing order by using the
`FIELD ORDER FORM` attribute. When this attribute is used, the tabbing order is
defined by the [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.")
attribute of the form fields. If this attribute is used, the [Dialog.fieldOrder](2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior.") FGLPROFILE entry is ignored.

The [`OPTIONS`](../09_advanced-features/0924-options-runtime.md "The OPTIONS instruction inside program blocks controls program behavior at runtime.") instruction
can also change the behavior of the `INPUT` instruction, with the `INPUT
WRAP` or `FIELD ORDER FORM` options.

## ACCEPT option

The `ACCEPT` attribute can be set to `FALSE` to avoid the automatic
creation of the accept default action. This option can be used for example when you want to write a
specific validation procedure, by using [`ACCEPT INPUT`](2034-accept-input-instruction.md).

## CANCEL option

The `CANCEL` attribute can be set to `FALSE` to avoid the automatic
creation of the cancel default action. This is useful for example when you only need a validation
action (accept), or when you want to write a specific cancellation procedure, by using [`EXIT INPUT`](2035-exit-input-instruction.md).

If the `CANCEL=FALSE` option is set, no [close](2286-multilevel-action-conflicts.md) action will be created, and you must
write an `ON ACTION close` control block to create an explicit action.

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

[Syntax of CONSTRUCT instruction](2048-syntax-of-construct-instruction.md "The CONSTRUCT instruction provides database query by example, producing a WHERE condition for SELECT.")
