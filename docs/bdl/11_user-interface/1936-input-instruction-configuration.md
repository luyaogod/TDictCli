---
title: "INPUT instruction configuration"
source: "fgl-topics/c_fgl_record_input_007.html"
breadcrumb: "User interface > Dialog instructions > Record input (INPUT) > Using simple record inputs > INPUT instruction configuration"
type: "concept"
description: "This section describes the options that can be specified in the ATTRIBUTES clause of the INPUT instruction. The options of the ATTRIBUTES clause override all default attributes and temporarily ..."
---

# INPUT instruction configuration

This section describes the options that can be specified in the `ATTRIBUTES`
clause of the `INPUT` instruction. The options of the `ATTRIBUTES`
clause override all default attributes and temporarily override any display attributes that the
`OPTIONS` or the `OPEN
WINDOW` statement specified for these fields. With the `INPUT`
statement, the `INVISIBLE` attribute is ignored.

## NAME option

The `NAME` attribute can be used to name the `INPUT`
dialog. This attribute is used to identify the `INPUT` dialog.

## HELP option

The `HELP` clause specifies the number of the help message to display if the
user invokes the help while the focus is in any field used by the instruction. The predefined
'help' action is automatically created by the runtime system. You can bind [action views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the 'help' action.

The `HELP` clause overrides the `HELP` attribute.

## WITHOUT DEFAULTS option

The `WITHOUT DEFAULTS` option indicates if the
fields controlled by `INPUT` must be filled (`FALSE`) or not
(`TRUE`) with the [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") values defined in the form specification file. The runtime system
assumes that the field satisfies the [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") attribute, when `WITHOUT DEFAULTS` is used. If the
`WITHOUT DEFAULTS` option is not used, all fields defined with the
`REQUIRED` attribute must be visited and modified.

For more details about `WITHOUT DEFAULTS` usage,, see [Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.").

## FIELD ORDER FORM option

By default, the tabbing order is defined by the [variable binding list](1935-variable-binding-in-input.md) in the instruction description. You can control the tabbing order by
using the `FIELD ORDER FORM` attribute. When this attribute is used, the tabbing
order is defined by the [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.") attribute of the form fields. If this attribute is used, the [Dialog.fieldOrder](2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior.") FGLPROFILE entry is ignored.

The [`OPTIONS`](../09_advanced-features/0924-options-runtime.md "The OPTIONS instruction inside program blocks controls program behavior at runtime.") instruction
can also change the behavior of the `INPUT` instruction, with the `INPUT
WRAP` or `FIELD ORDER FORM` options.

## UNBUFFERED option

Indicates that the dialog must be sensitive to program variable changes. When using this
option, you bypass the traditional `BUFFERED` mode.

When using the traditional "buffered" mode, program variable changes are not automatically
displayed to form fields; You need to execute a `DISPLAY TO` or `DISPLAY BY
NAME`. Additionally, if an action is triggered, the value of the current field is not
validated and is not copied into the corresponding program variable. The only way to get the text of
the current field is to use [`GET_FLDBUF()`](../08_language-basics/0671-get-fldbuf-function.md "The GET_FLDBUF() operator returns as character strings the current values of the specified fields.").

If the "unbuffered" mode is used, program variables and form fields are automatically synchronized.
You don't need to display explicitly values with a `DISPLAY TO` or `DISPLAY
BY NAME`. When an action is triggered, the value of the current field is validated and is
copied into the corresponding program variable.

See also [The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.").

## ACCEPT option

The `ACCEPT` attribute can be set to `FALSE` to avoid the automatic
creation of the accept default action. This option can be used for example when you want to write a
specific validation procedure, by using [`ACCEPT
INPUT`](1950-input-control-instructions.md).

## CANCEL option

The `CANCEL` attribute can be set to `FALSE` to avoid the automatic
creation of the cancel default action. This is useful for example when you only need a validation
action (accept), or when you want to write a specific cancellation procedure, by using [`EXIT INPUT`](1950-input-control-instructions.md).

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

[Syntax of the INPUT instruction](1932-syntax-of-the-input-instruction.md "The INPUT statement supports data entry in fields of the current form.")
