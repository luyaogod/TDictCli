---
title: "Enabling autocompletion"
source: "fgl-topics/c_fgl_prog_dialogs_completer.html"
breadcrumb: "User interface > User interface programming > Input fields > Enabling autocompletion"
type: "concept"
---

# Enabling autocompletion

> Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.

## Introduction to autocompletion

Text input fields (like `EDIT` and `BUTTONEDIT`) can be defined
with an autocompletion feature, by combining the `COMPLETER` form field attribute with
program code providing the list of completion proposals in a dynamic array of strings.
Autocompletion is activated with the `DIALOG.setCompleterItems()` method, when the `ON CHANGE` trigger is
fired for the field.

## Defining a form field for autocompletion

In order to enable autocompletion in a text form field, you must define the
`COMPLETER` attribute:

```
EDIT f1 = FORMONLY.firstname, COMPLETER;
```

The `COMPLETER` attribute can be used for `EDIT`
and `BUTTONEDIT` fields.

## Providing the front-end with a list of completion proposals

The `DIALOG.setCompleterItems()` method must be used to provide the list of
completion proposals during dialog execution:

```
DEFINE items DYNAMIC ARRAY OF STRING
-- fill the array with items
LET items[1] = "Ann"
LET items[2] = "Anna"
LET items[3] = "Annabel"
CALL DIALOG.setCompleterItems(items)
```

> **Important:**
>
> Consider the execution time of the code creating the completion
> proposal list. For example, avoid long complex SQL queries that can take more than a
> few milliseconds to complete.

The `setCompleterItems()` method will raise error  [-8114](../15_library-reference/4483-genero-bdl-errors.md) if the list of items contains
more than 50 elements. The purpose of autocompletion is to provide a short list of completion
proposals to the user. Note that this error is not trappable with exception handlers like
`TRY/CATCH`, the code must avoid exceeding the limit.

## Detecting user input

When implementing autocompletion, you must detect when the user modifies the
field value, to adapt the list of items with the `setCompleterItems()`
method.

In order to detect user input, define the `ON CHANGE` dialog control block, and
call a custom function by passing the `DIALOG` object, and the value of the current
field as parameter, to filter the completion proposal list
accordingly:

```
INPUT BY NAME rec.firstname
   ...
   ON CHANGE firstname
      CALL fill_proposals_firstname(DIALOG, rec.firstname)
```

For text fields defined with the `COMPLETER` attribute, the `ON
CHANGE` trigger will be fired without leaving the field, each time the user types characters
in. The event is fired after a short delay, so as not to overload the UI exchanges between the
front-end and the runtime system.

The item list for a field implementing autocompletion is not permanent: The program must redefine
the autocompletion item list with `setCompleterItems()`, on every `ON
CHANGE` event.

## Go to next field on completion item selection

When combining the `COMPLETER` attribute with the [`AUTONEXT`](1763-autonext-attribute.md "The AUTONEXT attribute forces the focus to automatically leave the current field when completed.") attribute, the focus will
automatically go to the next editable field, when the user select an item in the completion
list:

```
EDIT f1 = FORMONLY.field1, COMPLETER, AUTONEXT;
```

## Example

The example below implements form field with autocompletion: Each time the
`ON CHANGE` trigger is fired, the set of completion proposals is adapted
to the current field value, to match names that start with the same characters
typed by the user.

Form file (compl.per):

```
LAYOUT
GRID
{
[f1                             ]
[f2                             ]
}
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.field1, COMPLETER;
EDIT f2 = FORMONLY.field2;
END
```

Program file
(compl.4gl):

```
DEFINE all_names DYNAMIC ARRAY OF STRING

MAIN
  DEFINE rec RECORD
             field1 STRING,
             field2 STRING
         END RECORD
  CALL fill_names()
  OPEN FORM f FROM "compl"
  DISPLAY FORM f
  OPTIONS INPUT WRAP
  INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
      ON CHANGE field1
         CALL fill_proposals(DIALOG, rec.field1)
  END INPUT
END MAIN

FUNCTION fill_names()
  DEFINE i INTEGER
  LET i=0
  LET all_names[i:=i+1] = "Amanda"
  LET all_names[i:=i+1] = "Ann"
  LET all_names[i:=i+1] = "Anna"
  LET all_names[i:=i+1] = "Annabelle"
  LET all_names[i:=i+1] = "Barbara"
  LET all_names[i:=i+1] = "Barry"
  LET all_names[i:=i+1] = "Brice"
END FUNCTION

FUNCTION fill_proposals(dlg, curr_val)
  DEFINE dlg ui.Dialog, curr_val STRING
  DEFINE curr_set DYNAMIC ARRAY OF STRING,
         i, x INTEGER
  LET x=0
  FOR i=1 TO all_names.getLength()
      IF upshift(all_names[i]) MATCHES upshift(curr_val)||"*" THEN
         LET curr_set[x:=x+1] = all_names[i]
      END IF
  END FOR
  CALL dlg.setCompleterItems(curr_set)
END FUNCTION
```

## Related links

**Related concepts**  

[Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.")

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
