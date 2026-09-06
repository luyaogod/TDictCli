---
title: "INPUT ARRAY instruction configuration"
source: "fgl-topics/c_fgl_InputArray_009.html"
breadcrumb: "User interface > Dialog instructions > Editable record list (INPUT ARRAY) > Using editable record lists > INPUT ARRAY instruction configuration"
type: "concept"
description: "This section describes the options that can be specified in the ATTRIBUTES clause of the INPUT ARRAY instruction. The options of the ATTRIBUTES clause override all default attributes and temporarily ..."
---

# INPUT ARRAY instruction configuration

This section describes the options that can be specified in the `ATTRIBUTES`
clause of the `INPUT ARRAY` instruction. The options of the
`ATTRIBUTES` clause override all default attributes and temporarily override any
display attributes that the `OPTIONS` or the `OPEN WINDOW` statement
specified for these fields. With the `INPUT ARRAY` statement, the
`INVISIBLE` attribute is ignored.

## HELP option

The `HELP` clause specifies the number of a [help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to display if the user invokes the help the
`INPUT ARRAY` dialog. The predefined 'help' action is automatically created by the
runtime system. You can bind [action views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.") to the
'help' action.

The `HELP` clause overrides the `HELP` attribute.

## WITHOUT DEFAULTS option

The `WITHOUT DEFAULT` clause defines whether the
program array elements are populated (and to be displayed) when the dialog begins. Once the dialog
is started, existing rows are always handled as records to be updated in the database, while
newly-created rows are handled as records to be inserted in the database. The [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") and [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attributes defined in the
form are only used for newly-created rows.

It is unusual to implement an `INPUT ARRAY` with no `WITHOUT
DEFAULTS` option, because the data of the program array would be cleared and the list empty.
In a singular `INPUT ARRAY`, the default is `WITHOUT
DEFAULTS=FALSE`.

> **Important:**
>
> The default in `INPUT ARRAY` used inside `DIALOG` is
> `WITHOUT DEFAULTS=TRUE`, while in a singular `INPUT ARRAY` dialog, the
> default is `WITHOUT DEFAULTS=FALSE`.

For more details about `WITHOUT DEFAULTS` usage, see [Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.").

## FIELD ORDER FORM option

By default, the form tabbing order is defined by the variable list in the [binding specification](2010-variable-binding-in-input-array.md). You can control the tabbing order by
using the `FIELD ORDER FORM` attribute. When this attribute is used, the tabbing
order is defined by the [`TABINDEX`](1826-tabindex-attribute.md "The TABINDEX attribute defines the tab order for a form item.") attribute of the form items. With `FIELD ORDER
FORM`, if you jump from one field to a another with the mouse, the `BEFORE
FIELD` / `AFTER FIELD` triggers of intermediate fields are not executed
(actually, the [`Dialog.fieldOrder`](2221-dialog-configuration-with-fglprofile.md "FGLPROFILE parameters can be used to configure dialog behavior.") FGLPROFILE entry is ignored.)

If the form uses a [`TABLE`](1724-table-container.md "Defines a re-sizable table designed to display a list of records.") container, the front-end resets the tab indexes when the user moves
columns around. This way, the visual column order always corresponds to the input tabbing order. The
order of the columns in an editable list can be important; you may want to freeze the table columns
with the [`UNMOVABLECOLUMNS`](1833-unmovablecolumns-attribute.md "The UNMOVABLECOLUMNS attribute prevents the user from moving columns of a table.") attribute.

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

## COUNT option

The `COUNT` attribute defines the number of valid rows in the [static array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") to be displayed as default rows. If you do not use
the `COUNT` attribute, the runtime system cannot determine how much data to display,
so the screen array remains empty. You can also use the [SET\_COUNT()](../15_library-reference/2788-set-count.md "Defines the number of rows containing explicit data in a static array used by the next dialog.") built-in function, but it is
supported for backward compatibility only. The `COUNT` option is ignored when using a
[dynamic array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements."). If you specify the `COUNT`
attribute, the `WITHOUT DEFAULTS` option is not required because it is implicit. If
the `COUNT` attribute is greater than `MAXCOUNT`, the runtime system
will take `MAXCOUNT` as the actual number of rows. If the value of
`COUNT` is negative or zero, it defines an empty list.

## MAXCOUNT option

The `MAXCOUNT` attribute defines the maximum number of rows that can be
inserted in the program array. This attribute allows you to give an upper limit of the total number
of rows the user can enter, when using both [static or dynamic
arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.").

When binding a static array, `MAXCOUNT` is used as upper limit if it is
lower or equal to the actual declared static array size. If `MAXCOUNT` is greater
than the array size, the size of the static array is used as the upper limit. If
`MAXCOUNT` is lower than the `COUNT` attribute (or to the [`SET_COUNT()`](../15_library-reference/2788-set-count.md "Defines the number of rows containing explicit data in a static array used by the next dialog.") parameter), the
actual number of rows in the array will be reduced to `MAXCOUNT`.

When binding a dynamic array, the user can enter an infinite number of rows unless the
`MAXCOUNT` attribute is used. If `MAXCOUNT` is lower than the actual
size of the dynamic array, the number of rows in the array will be reduced to
`MAXCOUNT`.

If `MAXCOUNT` is negative or equal to zero, the user cannot insert rows.

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

## APPEND ROW option

The `APPEND ROW` attribute can be set to `FALSE` to avoid the
append default action, and prevent the user adding rows at the end of the list. If `APPEND
ROW=FALSE`, it is still possible to insert rows in the middle of the list. Use the
`INSERT ROW` attribute to disallow the user from inserting rows. To deny [automatic temporary row creation](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") if only
`APPEND ROW=FALSE` is used, set `AUTO APPEND` is not set to
`FALSE`.

## INSERT ROW option

The `INSERT ROW` attribute can be set to `FALSE` to avoid the
insert default action, and prevent the user inserting new rows in the middle of the list. However,
even if `INSERT ROW` is `FALSE`, it is still possible to append rows
at the end of the list. Use the `APPEND ROW` attribute to disallow the user from
appending rows. To deny [automatic temporary row
creation](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") if only `INSERT ROW=FALSE` is used, set `AUTO APPEND`
is not set to `FALSE`.

## DELETE ROW option

The `DELETE ROW` attribute can be set to `FALSE` to avoid the
delete default action, and prevent the user removing rows from the list.

## AUTO APPEND option

By default, an `INPUT ARRAY` controller creates a temporary row when needed (for
example, when the user deletes the last row of the list, a new row will be automatically created).
You can prevent this default behavior by setting the `AUTO APPEND` attribute to
`FALSE`. When this attribute is set to `FALSE`, the only way to create
a [new temporary row](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") is to execute the append
action.

If both the `APPEND ROW` and `INSERT ROW` attributes
are set to `FALSE`, the dialog automatically behaves
as if `AUTO APPEND` equals `FALSE`.

## KEEP CURRENT ROW option

Depending on the list container used in the form, the current row may be highlighted during the
execution of the dialog, and cleared when the instruction ends. You can change this default behavior
by using the `KEEP CURRENT ROW` attribute, to force the runtime system to keep the
current row highlighted.

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

[Syntax of INPUT ARRAY instruction](2007-syntax-of-input-array-instruction.md "The INPUT ARRAY supports data entry by users into a screen array and stores the entered data in an array of records.")
