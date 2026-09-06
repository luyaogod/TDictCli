---
title: "Multi-valued fields"
source: "fgl-topics/c_fgl_prog_dialogs_tagedit.html"
breadcrumb: "User interface > User interface programming > Input fields > Multi-valued fields"
type: "concept"
---

# Multi-valued fields

> Multi-valued form fields allow to display and input a set of values in the same character string field.

## Introduction to multi-valued fields

Text input fields `EDIT` and `BUTTONEDIT` support multi-valued
display and entry when using the `customWidget` style attribute set to [`"tagEdit"`](1637-edit-style-attributes.md):

![Screenshot showing a multi-value field](../_images/TagEdit01_gbc.jpg)

*Multi-valued field*

The raw value of the variable bound to the tag edit field must be a `;` semi-colon
seperated list of items, for example: `"Support;Marketing;Sales;"`.

The tag edit widget can be combined with the autocompletion feature, to provide the list of items
that can be entered in the field. You need to understand how autocompletion works with the
`COMPLETER` form field attribute. For more details, read [Enabling autocompletion](2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.").

The tag edit widget enters in action when the field is controlled by an `INPUT`,
`INPUT ARRAY` or `DISPLAY ARRAY` dialog. If the dialog is a
`CONSTRUCT`, the widget will be a regular edit box, to allow any search criteria
input.

## Defining the multi-valued form field

In order to get a multi-valued editor, define a form field with a [`STYLE`](1825-style-attribute.md "The STYLE attribute specifies a presentation style for a form element.") attribute that will reference a
.4st style entry defining the `customWidget` attribute to
`"tagEdit"`, and the [`COMPLETER`](1770-completer-attribute.md "The COMPLETER attribute enables autocompletion for the edit field.") attribute, to enable autocompletion proposals.

Best practice is to force the input length to be of the size of the program
`CHAR`, `VARCHAR` or `STRING` variable ([`SCROLL`](1815-scroll-attribute.md "The SCROLL attribute can be used to enable horizontal scrolling in a character field.")) and let the field widget
stretch horizontally ([`STRETCH=X`](1821-stretch-attribute.md "The STRETCH attribute defines if the form element can grow or has a fixed size."))

For example:

```
EDIT f1 = FORMONLY.labels,
              COMPLETER, STYLE="mvfield",
              SCROLL, STRETCH=X;
```

In the .4st file, define the `customWidget` style attribute
as `"tagEdit"`. To control new items creation in the tag edit, use the [`"allowTagCreation"`](1637-edit-style-attributes.md) style attribute (default is
`"yes"`):

```
  <Style name="Edit.mvfield">
     <StyleAttribute name="customWidget" value="tagEdit" />
     <StyleAttribute name="allowTagCreation" value="yes" />
  </Style>
```

## Storage data type for multi-valued fields

In the program code, the variable [bound](2232-binding-variables-to-form-fields.md "Some dialogs need program variables to store form field values.")
to the multi-valued form field, must be of type [`CHAR`](../08_language-basics/0557-char-size.md "The CHAR data type is a fixed-length character string data type."), [`VARCHAR`](../08_language-basics/0570-varchar-size.md "The VARCHAR data type is a variable-length character string data type, with a maximum size.") or [`STRING`](../08_language-basics/0567-string.md "The STRING data type is a variable-length, dynamically allocated character string data type, without limitation.").

The data type to be used depends on the solution you choose to store the multi-valued data in
your SQL database, by controlling the [maximum
input length](2233-input-length-of-form-fields.md "Field input length defines the amount of characters the user can type in a form field.") of the field and avoid input data loss.

Suggested solutions:

1. Store the multi-valued data as is, into a dedicated SQL column defined as
   `VARCHAR(size)`, as part of the data row for the main record.

   For example:

   ```
   CREATE TABLE incident (
      incident_id INTEGER NOT NULL PRIMARY KEY,
      inc_description VARCHAR(100) NOT NULL,
      ...
      inc_labels VARCHAR(500)  -- multi-valued data
   )
   ```

   In this case, use a `VARCHAR(size)` variable corresponding to
   the SQL column, in order to limit the maximum input length, and avoid data loss when writing to the
   database. The length of the SQL column and corresponding variable must be large enough to store all
   possible multi-valued combinations.
2. Split the multi-valued data into individual items, by using the semi-colon separator, and store
   each item in a row of a secondary SQL table.

   For example:

   ```
   CREATE TABLE mylabels (
      incident_id INTEGER NOT NULL,    -- pkey of main record
      label_pos INTEGER NOT NULL,      -- position of item
      label_val VARCHAR(50) NOT NULL,  -- item value
      PRIMARY KEY(incident_id,label_pos)
   )
   ```

   In this case, since the data is stored in a additional SQL table as individual rows, the number
   of items is not limited and the program variable can be of type `STRING` without
   input length limitation.

   Some additional code is needed to split and restore the multi-valued data, before writing to, and
   after reading from the database table. This code must control the maximum length of each item, to
   make sure that it fits in the target SQL column.

## Implemeting the list of completion proposals

The raw content of the field must be a `;` semi-colon seperated list of items. If
it is complete, the last item must be terminated with a semi-colon. If the last item in the list is
not ended by a semi-colon, it is considered as incomplete entry:

- `"Support;Marketing;Sales;"`: `"Sales"` is complete.
- `"Support;Marketing;Sal"` : `"Sal"` is incomplete.

The dialog code controlling the multi-valued field must implement the `ON CHANGE`
block, using the `DIALOG.setCompleterItems()` API to provide the completion proposals that hold a list
of items for all possible additional values. Each element of the proposal list must contain the
existing values present in the field, as proposal item prefix.

For example: The field contains already the values `"Support"` and
`"Sales"`, and the user enters additional characters such as `"Ma"`.
In order to propose the new values `"Marketing"` and `"Management"`
that can be added to the set, the array passed to the `setCompleterItems()` method
must hold the following elements:

- `"Support;Sales;Marketing;"`
- `"Support;Sales;Management;"`

If the new value entered by the user does not match a reference item, leave the array of
proposals empty, and implement code to add the new values to the reference list, for further usage,
based on the following rule: If the last entered item is ended by a semi-colon, it is considered as
complete.

## Example

Styles file
(mystyles.4st):

```
<?xml version="1.0" encoding="ANSI_X3.4-1968"?>
<StyleList>
  <Style name="Edit.mvfield">
     <StyleAttribute name="customWidget" value="tagEdit" />
  </Style>    
  <Style name="Window">
     <StyleAttribute name="windowType" value="normal" />
  </Style>
</StyleList>
```

Form file (form.per):

```
LAYOUT
GRID
{
Tag edit field:
[f1                               ]
Raw value field:
[f2                               ]
}
END
END

ATTRIBUTES
EDIT f1 = FORMONLY.field1,
  COMPLETER, STYLE="mvfield",
  SCROLL, STRETCH=X;
EDIT f2 = FORMONLY.field2,
  SCROLL;
END
```

Program file (main.4gl):

```
DEFINE rec RECORD
    field1 STRING,
    field2 STRING
END RECORD

DEFINE reftags DYNAMIC ARRAY OF STRING =
    ["Development", "Management", "Marketing", "Production", "Sales", "Support"]

MAIN
    DEFINE s INTEGER
    CALL ui.Interface.loadStyles("mystyles")
    OPEN FORM f FROM "form"
    DISPLAY FORM f
    OPTIONS INPUT WRAP
    LET rec.field1 = "Support;Sales;"
    LET rec.field2 = rec.field1
    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED, WITHOUT DEFAULTS)
        ON CHANGE field1
            LET s = fill_proposals(DIALOG, rec.field1)
            IF s < 0 THEN
                ERROR "Invalid field input"
                NEXT FIELD field1
            END IF
            LET rec.field2 = rec.field1
    END INPUT
END MAIN

PRIVATE FUNCTION fill_proposals(
    dlg ui.Dialog,
    curr_tags STRING
) RETURNS INTEGER
    DEFINE proptags DYNAMIC ARRAY OF STRING
    IF taglist_to_proposals(curr_tags, reftags, proptags) < 0 THEN
        RETURN -1
    END IF
    CALL dlg.setCompleterItems(proptags)
    RETURN 0
END FUNCTION

PRIVATE FUNCTION taglist_to_proposals(
    field_value STRING,
    refs DYNAMIC ARRAY OF STRING,
    prop DYNAMIC ARRAY OF STRING
) RETURNS INTEGER
    DEFINE x, lx, n INTEGER
    DEFINE tok base.StringTokenizer
    DEFINE element STRING
    DEFINE elements DYNAMIC ARRAY OF STRING
    DEFINE candidates DYNAMIC ARRAY OF STRING
    DEFINE current_value STRING

    CALL prop.clear()

    IF length(field_value) == 0 THEN
        RETURN 0
    END IF

    -- The element to check is always the last item in the multi-value string
    --
    --    Support;Sales;      <-- "Sales" (complete element)
    --    Support;Sa          <-- "Sa"    (partial or new element)
    --
    -- If ON CHANGE completion is triggered, it's because the user is editing
    -- this last element or has validated the selection of a new element in
    -- proposals.

    -- First, split current field value in to individual elements
    LET tok = base.StringTokenizer.create(field_value, ";")
    LET x = 0
    WHILE tok.hasMoreTokens()
        LET element = tok.nextToken()
        IF element IS NOT NULL THEN
            LET elements[x := x + 1] = element
            LET lx = x
        END IF
    END WHILE

    -- Save the last element to check, and remove it from the elements array, as
    -- it will be used later to check if already used
    LET element = elements[lx]
    CALL elements.deleteElement(lx)
    FOR x = 1 TO elements.getLength()
        LET current_value = current_value, (elements[x] || ";")
    END FOR

    -- Ending separator means last element is complete / selected from proposals
    -- by the end user, so we just append it to the reference list, if is not
    -- yet in that list.
    IF field_value.getCharAt(field_value.getLength()) == ";" THEN
        -- Check for duplicated usage
        IF array_search_ignore_case(elements, element) > 0 THEN
            RETURN -1
        END IF
        -- Here you can add some business rules if you want to deny this new
        -- element creation...
        IF array_search_ignore_case(refs, element) == 0 THEN
            LET refs[refs.getLength() + 1] = element
        END IF
        -- In any case, ending sep means no proposals are required
        RETURN 0
    END IF

    -- Now search in the reference list if last element exists, or looks like
    -- and fill the candidates list
    LET x = 0
    FOR n = 1 TO refs.getLength()
        IF refs[n].toUpperCase() MATCHES (element.toUpperCase() || "*") THEN
            IF array_search_ignore_case(elements, refs[n]) == 0 THEN
                LET candidates[x := x + 1] = refs[n]
            END IF
        END IF
    END FOR

    -- Fill the proposals with existing elements, one element candidates
    FOR x = 1 TO candidates.getLength()
        LET prop[prop.getLength() + 1] = current_value, (candidates[x] || ";")
    END FOR

    RETURN 0

END FUNCTION

PRIVATE FUNCTION array_search_ignore_case(
    arr DYNAMIC ARRAY OF STRING,
    element STRING
) RETURNS INTEGER
    DEFINE x INTEGER
    FOR x = 1 TO arr.getLength()
        IF arr[x].toUpperCase() == element.toUpperCase() THEN
            RETURN x
        END IF
    END FOR
    RETURN 0
END FUNCTION
```

> **Tip:**
>
> To deny new element creations, consider to reset the field value if the last element has to be
> rejected.
>
> First, set the `allowTagCreation` style attribute to `"no"` in
> order to avoid drop down list to appear for new entered
> elements:
>
> ```
>   <Style name="Edit.mvfield">
>      <StyleAttribute name="customWidget" value="tagEdit" />
>      <StyleAttribute name="allowTagCreation" value="no" />
>   </Style>
> ```
>
> Adapt the program code implementing the field autocompletion as follows, to reset the field value
> for invalid elements:
>
> ```
> PRIVATE FUNCTION taglist_to_proposals(
>     field_value STRING,
>     refs DYNAMIC ARRAY OF STRING,
>     prop DYNAMIC ARRAY OF STRING
> ) RETURNS (INTEGER, STRING)
>
>   ...
>
>   IF field_value.getCharAt(field_value.getLength()) == ";" THEN
>       -- Check for duplicated usage
>       IF array_search_ignore_case(elements,element) > 0 THEN
>           RETURN -1, current_value
>       END IF
>       -- Check it element is in the reference list
>       IF array_search_ignore_case(refs,element) == 0 THEN
>          RETURN -2, current_value
>       END IF
>       RETURN 0, NULL
>   END IF
>
>   ...
> ```
>
> In the calling function:
>
> ```
> PRIVATE FUNCTION fill_proposals(
>     dlg ui.Dialog,
>     curr_tags STRING
> ) RETURNS (INTEGER, STRING)
>     DEFINE proptags DYNAMIC ARRAY OF STRING
>     DEFINE tmp STRING, s INTEGER
>     CALL taglist_to_proposals(curr_tags, reftags, proptags)
>          RETURNING s, tmp
>     IF s < 0 THEN
>         RETURN -1, tmp
>     END IF
>     CALL dlg.setCompleterItems(proptags)
>     RETURN 0, NULL
> END FUNCTION
> ```
>
> In the dialog code:
>
> ```
> DEFINE tmp STRING
> ...
>         ON CHANGE field1
>             CALL fill_proposals(DIALOG, rec.field1) RETURNING s, tmp
>             IF s < 0 THEN
>                 ERROR "Invalid field input"
>                 LET rec.field1 = tmp
>                 NEXT FIELD field1
>             END IF
> ...
> ```

## Related links

**Related concepts**  

[Enabling autocompletion](2247-enabling-autocompletion.md "Autocompletion allows a list of completion proposals to be displayed while the user is typing text into a field.")
