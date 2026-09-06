---
title: "CLEAR field-list"
source: "fgl-topics/c_fgl_record_display_CLEAR_field.html"
breadcrumb: "User interface > Dialog instructions > Static display (DISPLAY/ERROR/MESSAGE/CLEAR) > CLEAR field-list"
type: "concept"
---

# CLEAR field-list

> The CLEAR field-list instruction clears specific fields in the current form.

## Syntax

```
CLEAR field-list
```

where field-list
is:

```
{ field-name
| table-name.*
| table-name.field-name
| screen-array[line].*
| screen-array[line].field-name
| screen-record.*
| screen-record.field-name
}
[,...]
```

1. field-name is the identifier of a field of the current form.
2. table-name is the identifier of a database table of the current form.
3. screen-record is the identifier of a screen record of the current form.
4. screen-array is the screen array that will be used in the form.
5. line is the line number in the screen array.

## Usage

The `CLEAR field-list` instruction can be used to clear the
content of the specified form fields.

> **Important:**
>
> The `CLEAR field-list` instruction sets the
> [field modification flags](2237-input-field-modification-flag.md "Each input field controlled by a dialog instruction has a modification flag.").

The fields to be cleared can be specified individually or by referencing a [screen record or screen array](1676-screen-records-arrays.md "Form fields can be grouped in a screen record or screen array definition."), with
the `.*` notation to specify all fields.

Similar to `CLEAR
FORM`, the `CLEAR field-list` is typically used when the
program is not inside a dialog block execution controlling the form fields. For example, after a
database query with a [`CONSTRUCT`](2046-query-by-example-construct.md "The CONSTRUCT instruction implements database query criteria input in an application form.")
instruction, you might want to clear all search criteria entered by the user with this instruction,
to cleanup the form.

The `CLEAR field-list` instruction clears the field values and
resets the TTY attributes to `NORMAL`.

The `CLEAR field-list` instruction is not needed if the program
is always in the context of a [dialog controlling
the form fields](2219-the-model-view-controller-paradigm.md "The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.").

## Example

```
  CONSTRUCT BY NAME sql
     ON s_customer.*
     ...
  END CONSTRUCT
  CLEAR s_customer.*
  ...
```

## Related links

**Related concepts**  

[CLEAR FORM](1882-clear-form.md "The CLEAR FORM instruction clears all fields in the current form.")

[CLEAR SCREEN ARRAY](1883-clear-screen-array.md "The CLEAR SCREEN ARRAY instruction clears the values of all rows of the form list identified by the specified screen array.")
