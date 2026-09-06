---
title: "Immediate detection of user changes"
source: "fgl-topics/c_fgl_prog_dialogs_dialogtouched.html"
breadcrumb: "User interface > User interface programming > Input fields > Immediate detection of user changes"
type: "concept"
---

# Immediate detection of user changes

> This section describes the dialogtouched predefined action.

## Purpose of the `dialogtouched` action

The dialogtouched predefined action is a special action that can be used to detect
user changes immediately without leaving the current field.

The event is trapped with an `ON ACTION dialogtouched` block, to execute code in
your program.

> **Important:**
>
> **The `dialogtouched` action must be enabled/disabled in
> accordance with the needs of the dialog: When this action is enabled, the `ON ACTION
> dialogtouched` block will be invoked each time the user types characters, modifies
> the value with copy/paste actions, or uses the widget input helper (like the calendar of a
> `DATEEDIT`). In client/server mode, this can generate more network traffic as needed.
> As soon as the `dialogtouched` action is fired, it should be disabled to avoid
> un-necessary network round-trips, and it should only be re-enabled when needed.**

When the form item type allows value changes to be detected immediately, for example with
`COMBOBOX`, `CHECKBOX` or `DATEEDIT` fields, the
alternative to `dialogtouched` to detect field input changes is to use the
`ON CHANGE` trigger.

## Typical usage of `dialogtouched` action

Singular interactive instructions are typically ended with an `accept` or
`cancel` action. For example, a singular `INPUT` statement allows the
end user to enter a database record, and validate or cancel the changes. The `INPUT`
statement is then re-executed, to enter or modify another record.

Unlike singular dialogs, the `DIALOG` instruction can be used continuously for
several data operations, such as navigation, creation, or modification. Typically, the default is
the navigation mode, and as soon as the user starts to modify a field, it switches to edit mode, to
modify a record, or create a new record. Only a user-defined `close` or
`exit` action will terminate a `DIALOG` block. In such case, the
dialog must be notified when the user starts to modify the current record. This can be achieved with
the `dialogtouched` predefined action.

## What user events fire a `dialogtouched` action?

The `dialogtouched` action works for any field controlled by the current
interactive instruction, and with any type of form field.

Every time the user modifies the value of a field (without leaving the field), the `ON
ACTION dialogtouched` block will be executed, if it is enabled.

The `dialogtouched` action can be triggered by typing characters in a text editor
field, using copy/paste, clicking on a `CHECKBOX` / `RADIOGROUP`,
moving the cursor of a `SLIDER` or changing a date with the calendar of a
`DATEEDIT`.

## Field value validation when `dialogtouched` occurs

When a `dialogtouched` action occurs, the current field may contain some text that
does not represent a correct value corresponding to the field data type. For example, a form field
bound to a `DATE` variable may contain only a part of a valid date string, such as
"`12/24/`". For this reason, the target variable bound to the field cannot hold the
current text displayed on the screen when the `ON ACTION dialogtouched` code is
executed, even when using the `UNBUFFERED` mode.

To avoid data validation on action code execution, the `dialogtouched` action is
defined with [`validate="no"`](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.")
attribute in the $FGLDIR/lib/default.4ad action defaults file. This is
mandatory when using the `UNBUFFERED` mode; otherwise the runtime would try to copy
the input buffer into the program variable when a `dialogtouched` action is
invoked.

Do not define `validate="yes"` for the `dialogtouched` action,
otherwise non-string data fields will in most cases produce a conversion error, when the user enters
data.

## Programming pattern to handle a "`save`" button

By default, the `dialogtouched` action and navigation actions are enabled.

In the `ON ACTION dialogtouched` block, detect the beginning of a record
modification in a `DIALOG` block.

To prevent further `dialogtouched` action events, disable the action with a [`DIALOG.setActionActive()`](../15_library-reference/3213-ui-dialog-setactionactive.md "Enabling and disabling dialog actions.")
method, disable also navigation actions, and enable the `save` action.

Set a flag/status variable to track that the dialog is in edit mode.

When user input is validated and committed in the database, the `dialogtouched`
and navigation actions can be enabled again, and the `save` action can be
disabled.

Set a flag/status variable to indicate that the dialog is back to navigation mode.

In the action handler for the `close` or `exit` action, which can
be used to close the form, check the status flag to know if the user has started to edit the fields,
and show a warning box before leaving the dialog with `EXIT DIALOG`.

Code example:

form.per:

```
LAYOUT
GRID
{
Id:  [f1          ] Name: [f2                  ]
Addr:[f3                                       ]
}
END
END
ATTRIBUTES
EDIT f1 = FORMONLY.cust_id, NOENTRY;
EDIT f2 = FORMONLY.cust_name;
EDIT f3 = FORMONLY.cust_addr;
END
```

main.4gl:

```
PRIVATE DEFINE editing BOOLEAN

MAIN
    DEFINE r_cust RECORD
                  cust_id INTEGER,
                  cust_name VARCHAR(20),
                  cust_addr VARCHAR(100)
           END RECORD
    DEFINE sqlcond STRING

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    DIALOG ATTRIBUTES(UNBUFFERED)

        INPUT BY NAME r_cust.* ATTRIBUTES(WITHOUT DEFAULTS)
        END INPUT

        BEFORE DIALOG
           CALL setup_dialog(DIALOG,FALSE)

        ON ACTION dialogtouched
           CALL setup_dialog(DIALOG,TRUE)

        ON ACTION save
            MESSAGE "Saving record to database..."
            CALL setup_dialog(DIALOG,FALSE)

        ON ACTION query
            CONSTRUCT BY NAME sqlcond ON cust_name;

        ON ACTION close
            IF NOT check_close(DIALOG) THEN
                NEXT FIELD CURRENT
            END IF
            EXIT DIALOG

    END DIALOG

END MAIN

PRIVATE FUNCTION setup_dialog(d ui.Dialog, e BOOLEAN) RETURNS ()
   LET editing = e
   CALL d.setActionActive("dialogtouched", NOT editing)
   CALL d.setActionActive("save", editing)
   CALL d.setActionActive("query", NOT editing)
END FUNCTION

PRIVATE FUNCTION check_close(d ui.Dialog) RETURNS BOOLEAN
    IF editing THEN
       CALL setup_dialog(d,editing)
       RETURN mbox_yn("Do you want to close the form without saving changes?")
    END IF
    RETURN TRUE
END FUNCTION

PUBLIC FUNCTION mbox_yn(msg STRING) RETURNS BOOLEAN
    DEFINE r BOOLEAN
    MENU "Question" ATTRIBUTES(STYLE="dialog",COMMENT=msg)
        COMMAND "Yes" LET r = TRUE
        COMMAND "No"  LET r = FALSE
    END MENU
    RETURN r
END FUNCTION
```

## Related links

**Related concepts**  

[The buffered and unbuffered modes](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")

[Reacting to field value changes](2238-reacting-to-field-value-changes.md "This section describes the purpose of the ON CHANGE interaction block.")
