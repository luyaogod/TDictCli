---
title: "Controlling web application state (#anchor)"
source: "fgl-topics/c_fgl_prog_dialogs_app_state.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Controlling web application state (#anchor)"
type: "concept"
---

# Controlling web application state (#anchor)

> With GBC in a web browser, the context/state of a program can be managed with URL # anchors.

## Purpose of `#anchor` in URLs

Single-page web applications (SPA) manage the state of an app by using `#` anchors
in the URL, such as in `"http://myserver/ua/r/myApp#myState"`.

URL anchors allow the end user to bookmark the application in a given state.

When reloading an app by using an URL with a `#` anchor marker, end users expect
to land in the required application state.

Genero BDL offers an API and a special action to control and detect
`#` anchors usage in URLs.
> **Important:**
>
> The URL `#anchor` management is only available in browser
> context as there is no way for the end-user to bookmark an application in GDC, GMA or GMI.

## Limitations

URL anchors management is limited to the main running application. Any application spawned using
the `RUN` instruction will not be able to execute these front calls or get the
`applicationstatechanged` special action.

## Detecting `#anchor` changes

Changing manually the anchor in the browser address bar widget will trigger a predefined
[special action](2257-list-of-predefined-actions.md) named `applicationstatechanged`.

In the program dialog code, you can then define an action handler to detect such change, and
redirect the code to the corresponding context /
state:

```
ON ACTION applicationstatechanged
   -- get the current #anchor
   -- call the corresponding function/dialog
```

When using the [`UNBUFFERED`](2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
attribute in the input dialog, you can control the field validation with the [`VALIDATE`](2277-validate-action-attribute.md "The VALIDATE action attribute defines the data validation level for a given action.") action attribute. The
`applicationstatechanged` action is typically similar to a `cancel`
action, and the field validation attribute should be set to `NO`.

## Controlling the URL `#anchor`

Genero BDL provides two front calls, to get and set the `#` anchor of the current
URL displayed in the browser address bar:

1. [browser.getApplicationState](../15_library-reference/3434-browser-getapplicationstate.md "Gets the # anchor of the current URL in the browser address bar.")
2. [browser.setApplicationState](../15_library-reference/3433-browser-setapplicationstate.md "Sets the # anchor of the URL in the browser address bar.")

The application code should start by a call to the `getApplicationState`
frontcall, to check for an existing `#anchor`. If an anchor is
present, the program flow should go to the corresponding function/dialog. When returning from this
function, the code should execute a dialog as when no anchor was provided.

In the initial/main dialog (and if needed, in subsequent dialogs), implement an `ON
ACTION` handler for the `applicationstatechanged` action name, to detect when
the end user changes the `#anchor` part of the URL. If this action
is triggered, query the anchor with the `getApplicationState` frontcall and follow
the same work flow as when the application starts.

In all interactive code parts that correspond to an application state which can be reached
directly by using a #anchor in the URL, set the current `#anchor`
with the `setApplicationState` frontcall, before starting a new dialog
instruction.

## Example

In the following example, the initial state of the application is a regular
`MENU`, as well as en `INPUT BY NAME` and `CONSTRUCT`
dialog, that can be reached with the #input and #construct anchors in the URL.

Form file: form.per

```
LAYOUT
GRID
{
Num:  [f1         ]
Name: [f2                        ]
}
END
END

ATTRIBUTES
EDIT f1 = FORMONLY.num TYPE INTEGER;
EDIT f2 = FORMONLY.name TYPE VARCHAR;
END
```

Program file: main.4gl

```
MAIN
    DEFINE anchor STRING
    OPTIONS INPUT WRAP
    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1
    CALL ui.Interface.frontCall("browser",
            "getApplicationState",[],[anchor])
    ERROR "Initial anchor: ", NVL(anchor,"none")
    LET int_flag = FALSE
    WHILE NOT int_flag
       LET anchor = switch_context(anchor)
    END WHILE
END MAIN

FUNCTION switch_context(anchor STRING) RETURNS STRING
    DEFINE next_anchor STRING
    CASE anchor
    WHEN "input"
       CALL do_input()
       LET next_anchor = NULL
    WHEN "construct"
       CALL do_construct()
       LET next_anchor = NULL
    OTHERWISE
       IF anchor IS NULL THEN
          LET next_anchor = do_menu()
       ELSE
          ERROR "Invalid application state: ", anchor
       END IF
    END CASE
    RETURN next_anchor
END FUNCTION

FUNCTION do_menu() RETURNS STRING
    DEFINE anchor STRING
    CALL ui.Interface.frontCall("browser",
            "setApplicationState",[NULL],[])
    MENU "test"
        ON ACTION input
           CALL do_input()
           LET anchor = NULL
           EXIT MENU
        ON ACTION construct
           CALL do_construct()
           LET anchor = NULL
           EXIT MENU
        -- During this menu dialog, user can switch to another app state
        ON ACTION applicationstatechanged
           CALL ui.Interface.frontCall("browser",
                   "getApplicationState",[],[anchor])
           ERROR "State change: ", anchor
           EXIT MENU
        ON ACTION quit
           LET int_flag=TRUE
           EXIT MENU
        ON ACTION close
           LET int_flag=TRUE
           EXIT MENU
    END MENU
    RETURN anchor
END FUNCTION

FUNCTION do_input() RETURNS ()
    DEFINE rec RECORD
               num INTEGER,
               name VARCHAR(50)
           END RECORD
    CALL ui.Interface.frontCall("browser",
            "setApplicationState",["input"],[])
    INPUT BY NAME rec.* ATTRIBUTES(UNBUFFERED)
    LET int_flag=FALSE
END FUNCTION

FUNCTION do_construct() RETURNS ()
    DEFINE sqlcond STRING
    CALL ui.Interface.frontCall("browser",
            "setApplicationState",["construct"],[])
    CONSTRUCT BY NAME sqlcond ON num, name
    LET int_flag=FALSE
END FUNCTION
```

## Related links

**Related concepts**  

[Predefined actions](2255-predefined-actions.md "Genero predefines some action names for common operations of interactive instructions.")
