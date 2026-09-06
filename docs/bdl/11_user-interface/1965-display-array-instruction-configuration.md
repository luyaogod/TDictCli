---
title: "DISPLAY ARRAY instruction configuration"
source: "fgl-topics/c_fgl_DisplayArray_016.html"
breadcrumb: "User interface > Dialog instructions > Record list (DISPLAY ARRAY) > Using record lists > DISPLAY ARRAY instruction configuration"
type: "concept"
description: "This section describes the options that can be specified in the ATTRIBUTES clause of the DISPLAY ARRAY instruction. The options of the ATTRIBUTES clause override all default attributes and temporarily ..."
---

# DISPLAY ARRAY instruction configuration

This section describes the options that can be specified in the `ATTRIBUTES`
clause of the `DISPLAY ARRAY` instruction. The options of the
`ATTRIBUTES` clause override all default attributes and temporarily overrides
any display attributes that the `OPTIONS` or the `OPEN WINDOW`
statement specified for these fields. With the `DISPLAY ARRAY` statement, the
`INVISIBLE` attribute is ignored.

## HELP option

The `HELP` clause specifies the number of a [help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to display if the user invokes the
help in the `DISPLAY ARRAY` dialog. The predefined 'help' action is
automatically created by the runtime system. You can bind [action views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the 'help' action.

The `HELP` clause overrides the `HELP` attribute.

## COUNT option

When using a dynamic array, the number of rows to be displayed is defined by the number
of elements in the dynamic array; the `COUNT` attribute is ignored.

When using a static array or the paged mode, the number of rows to be displayed is defined
by the `COUNT` attribute. You can also use the [`SET_COUNT()`](../15_library-reference/2788-set-count.md "Defines the number of rows containing explicit data in a static array used by the next dialog.") built-in
function, but it is supported for backward compatibility only. If you don't know the total
number of rows for the paged mode, you can specify -1 for the `COUNT`
attribute (or in the `SET_COUNT()` call before the dialog block). With
`COUNT=-1`, the dialog will ask for rows by executing `ON FILL
BUFFER` until less rows are provided than are required, or if you reset the number
of rows to a value higher than -1 with [ui.Dialog.setArrayLength()](../15_library-reference/3220-ui-dialog-setarraylength.md "Sets the number of rows in a DISPLAY ARRAY using paged mode.").

## KEEP CURRENT ROW option

Depending on the list container used in the form, the current row may be highlighted
during the execution of the dialog, and cleared when the instruction ends. You can change
this default behavior by using the `KEEP CURRENT ROW` attribute, to force
the runtime system to keep the current row highlighted.

## ACCEPT option

The `ACCEPT` attribute can be set to `FALSE` to avoid the
automatic creation of the default accept action. Use this attribute when you want to avoid
dialog validation, or if you need to write a specific validation procedure using [ACCEPT DISPLAY](1998-accept-display-instruction.md).

## CANCEL option

The `CANCEL` attribute can be set to `FALSE` to avoid the
automatic creation of the cancel default action. Use this attribute when you only need a
validation action (accept), or when you want to write a specific cancellation procedure
using [EXIT DISPLAY](1997-exit-display-instruction.md).

If the `CANCEL=FALSE` option is set, no [close](2286-multilevel-action-conflicts.md) action will be
created, and you must write an `ON ACTION close` control block to create an
explicit action.

## DOUBLECLICK option

The `DOUBLECLICK` option can be used to define the action that will
be fired when the user chooses a row from the list. Different configuration options are
available to control the row selection action of desktop and mobile devices. For more
details, see [Defining the action for a row choice](2330-defining-the-action-for-a-row-choice.md "The row choice in a TABLE can be associated with a dedicated action.").

## FOCUSONFIELD option

When the `FOCUSONFIELD` option is used, the `DISPLAY ARRAY` allows
focus at the field (or cell) level. It is then possible to implement `BEFORE FIELD`
and `AFTER FIELD` blocks, as well as using `NEXT FIELD` instructions.
However, the dialog still manages a read-only list. For more details, see [Field-level focus in DISPLAY ARRAY](2305-field-level-focus-in-display-array.md "The DISPLAY ARRAY dialog supports cell-level focus with the FOCUSONFIELD.").

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

[Syntax of DISPLAY ARRAY instruction](1961-syntax-of-display-array-instruction.md "The DISPLAY ARRAY instruction controls the display of a program array on the screen.")
