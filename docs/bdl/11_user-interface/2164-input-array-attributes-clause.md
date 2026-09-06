---
title: "INPUT ARRAY ATTRIBUTES clause"
source: "fgl-topics/c_fgl_DIALOG_block_INPUT_ARRAY_ATTRIBUTES_2.html"
breadcrumb: "User interface > Dialog instructions > Declarative dialogs (DIALOG - at module level) > Using declarative dialogs > Declarative DIALOG block configuration > INPUT ARRAY ATTRIBUTES clause"
type: "concept"
---

# INPUT ARRAY ATTRIBUTES clause

> INPUT ARRAY specific attributes can be defined in the ATTRIBUTE clause of the sub-dialog header.

`INPUT ARRAY` specific attributes can be defined in the `ATTRIBUTE`
clause of the sub-dialog header.

## HELP option

The `HELP` clause specifies the number of a
[help message](1593-message-files.md "Message files centralize strings and larger texts identified by a number, that can be used in programs.") to display
if the user invokes the help the `INPUT ARRAY` dialog.
The predefined 'help' action is automatically created by the runtime system.
You can bind [action views](2253-dialog-actions.md "Describes how to program action handling when the end user triggers an action on the front-end.")
to the 'help' action.
The `HELP` clause overrides the `HELP`
attribute.

## COUNT option

The `COUNT` attribute defines the number of valid rows in
the [static array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.") to be displayed as
default rows. If you do not use the `COUNT` attribute, the
runtime system cannot determine how much data to display, so the screen array
remains empty. The `COUNT` option is ignored when using a
[dynamic array](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.").
If you specify the `COUNT` attribute, the
`WITHOUT DEFAULTS` option is not required because it is
implicit. If the `COUNT` attribute is greater than
`MAXCOUNT`, the runtime system will take `MAXCOUNT`
as the actual number of rows. If the value of `COUNT` is
negative or zero, it defines an empty list.

## MAXCOUNT option

The `MAXCOUNT` attribute defines the maximum number of rows that can be
inserted in the program array. This attribute allows you to give an upper limit of the total number
of rows the user can enter. It can be used with [static or dynamic
arrays](../08_language-basics/0729-arrays.md "Arrays (static or dynamic) allow you to handle an ordered collection of elements.").

When binding a static array, `MAXCOUNT` is used as upper limit if it is
lower or equal to the actual declared static array size. If `MAXCOUNT` is greater
than the array size, the size of the static array is used as the upper limit. If
`MAXCOUNT` is lower than the `COUNT` attribute (or to the [SET\_COUNT()](../15_library-reference/2788-set-count.md "Defines the number of rows containing explicit data in a static array used by the next dialog.") parameter when using a singular
[INPUT ARRAY](2005-editable-record-list-input-array.md "The INPUT ARRAY instruction provides always-editable record list handling in an application form.")), the actual number of rows in the array
will be reduced to `MAXCOUNT`.

When binding a dynamic array, the user can enter an infinite number of rows unless the
`MAXCOUNT` attribute is used. If `MAXCOUNT` is lower than the actual
size of the dynamic array, the number of rows in the array will be reduced to
`MAXCOUNT`.

If `MAXCOUNT` is negative or equal to zero, the user cannot
insert rows.

## APPEND ROW option

The `APPEND ROW` attribute can be set to `FALSE`
to avoid the append default action, and deny the user to add rows at the end of
the list.
If `APPEND ROW =FALSE`, it is still possible to insert rows in
the middle of the list. Use the `INSERT ROW` attribute to
disallow the user from inserting rows.
Additionally, even with `APPEND ROW=FALSE` and `INSERT
ROW=FALSE`, you can still get
[automatic temporary row creation](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.")
if `AUTO APPEND` is not set to `FALSE`.

## INSERT ROW option

The `INSERT ROW` attribute can be set to `FALSE`
to avoid the insert default action, and deny the user to insert new rows
in the middle of the list.
However, even if `INSERT ROW` is `FALSE`, it is
still possible to append rows at the end of the list.
Use the `APPEND ROW` attribute to disallow the user from
appending rows. Additionally, even with `APPEND ROW=FALSE` and
`INSERT ROW=FALSE`, you can still get
[automatic temporary row creation](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.")
if `AUTO APPEND` is not set to `FALSE`.

## DELETE ROW option

The `DELETE ROW` attribute can be set to `FALSE`
to avoid the delete default action, and deny the user to remove rows from
the list.

## AUTO APPEND option

By default, an `INPUT ARRAY` controller creates a temporary
row when needed (for example, when the user deletes the last row of the list,
an new row will be automatically created).
You can prevent this default behavior by setting the `AUTO APPEND`
attribute to `FALSE`.
When this attribute is set to `FALSE`, the only way to create
a [new temporary row](2311-appending-rows-in-input-array.md "Rows appended at the end of an editable list are temporary until they are edited.") is
to execute the append action.

If both the `APPEND ROW` and `INSERT ROW`
attributes are set to `FALSE`, the dialog automatically behaves
as if `AUTO APPEND` equals `FALSE`.

## KEEP CURRENT ROW option

Depending on the list container used in the form, the current row may be
highlighted during the execution of the dialog, and cleared when the instruction
ends. You can change this default behavior by using the
`KEEP CURRENT ROW` attribute, to force the runtime system
to keep the current row highlighted.

## WITHOUT DEFAULTS option

The `WITHOUT DEFAULT` clause defines whether the
program array elements are populated (and to be displayed) when the dialog begins. Once the dialog
is started, existing rows are always handled as records to be updated in the database, while
newly-created rows are handled as records to be inserted in the database. The [`REQUIRED`](1812-required-attribute.md "The REQUIRED attribute forces the user to modify the content of a field during an input dialog.") and [`DEFAULT`](1772-default-attribute.md "The DEFAULT attribute defines a default value to a field during data entry.") attributes defined in the
form are only used for newly-created rows.

It is unusual to implement an `INPUT ARRAY` with no `WITHOUT
DEFAULTS` option, because the data of the program array would be cleared and the list empty.
In an `INPUT ARRAY` sub-dialog, the default is `WITHOUT
DEFAULTS=TRUE`.

> **Important:**
>
> The default in `INPUT ARRAY` used inside `DIALOG` is
> `WITHOUT DEFAULTS=TRUE`, while in a singular `INPUT ARRAY` dialog, the
> default is `WITHOUT DEFAULTS=FALSE`.

For more details about `WITHOUT DEFAULTS` usage, see [Form field initialization](2236-form-field-initialization.md "Form field initialization can be controlled by the WITHOUT DEFAULTS dialog option.").
