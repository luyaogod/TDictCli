---
title: "Refreshing the user interface"
source: "fgl-topics/c_fgl_MigI4GL_refresh.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > User interface topics > Refreshing the user interface"
type: "concept"
---

# Refreshing the user interface

> Genero BDL only refreshes the user interface when the runtime waits for user interaction. In certain scenarios, this can result in information displayed in an I4GL application not being displayed when running as a Genero BDL application.

## Genero refreshes the screen before user interaction

To optimize the display in TUI and GUI mode, the Genero runtime system only refreshes the
screen/windows, when an interactive instruction (a dialog) waits for a user interaction.

In most cases, the user interface refresh behavior of Genero BDL (where the output is delayed
until the runtime waits for a user interaction) does not introduce any display problem in legacy
programs. However, attention is required when information must be immediately visible while the
runtime continues processing code. A typical example is a `"Please wait..."` message,
displayed just before a lengthy processing.

## IBM® Informix® 4GL behavior

IBM
Informix 4GL (I4GL) only supports a TUI mode, displaying
screens on dumb terminal or terminal emulators.

When a program executes an instruction displaying information to the user, IBM
Informix 4GL refreshes the screen immediately.

For example, when doing successive `DISPLAY ... AT` instructions in a loop, I4GL
will show screen changes for each `DISPLAY`
instruction:

```
MAIN
    DEFINE i INTEGER
    MENU "Test"
        COMMAND "Start"
            FOR i=1 TO 50000
                DISPLAY i AT 5,5
            END FOR
        COMMAND "Quit"
            EXIT MENU
    END MENU
END MAIN
```

## Genero BDL in TUI mode

The application runs in text (TUI) mode when FGLGUI=0. This section refers to an application
running in TUI mode.

With the above code running in TUI mode, the screen will not show the numbers. It will only show
the final number, when the `MENU` gives the control back to the user.

To display all of the numbers to the end user, force the screen refresh with an [`ui.Interface.refresh()`](../15_library-reference/3117-ui-interface-refresh.md "Synchronize the user interface with the front-end.") API
call.

```
MAIN
    DEFINE i INTEGER
    MENU "Test"
        COMMAND "Start"
            FOR i=1 TO 50000
                DISPLAY i AT 5,5
                CALL ui.Interface.refresh()
            END FOR
        COMMAND "Quit"
            EXIT MENU
    END MENU
END MAIN
```

> **Important:**
>
> Forcing the user interface refresh as shown in this example must only
> be done when using the TUI mode. See GUI mode case for details.

## Genero BDL in GUI mode

When using the GUI mode of Genero, instead of displaying directly to the terminal screen, the
runtime system communicates with a front-end. See [The dynamic user interface](../11_user-interface/1509-the-dynamic-user-interface.md "The dynamic user interface is the base concept of the Genero user interaction components.") for more
details.

With GUI mode, if the program refreshes the user interface too often, it will produce a lot of
unnecessary network traffic. The user interface should be refreshed periodically rather than
continuously, to avoid network clogging.

In the following code example, we have added the `(i MOD 100)==0` test to reduce
the number of refresh calls.

```
MAIN
    DEFINE i INTEGER
    MENU "Test"
        COMMAND "Start"
            FOR i=1 TO 5000000
                IF (i MOD 100) == 0 THEN
                    DISPLAY i AT 5,5
                    CALL ui.Interface.refresh()
                END IF
            END FOR
        COMMAND "Quit"
            EXIT MENU
    END MENU
END MAIN
```

> **Important:**
>
> This code example is provided to explain the Genero behavior. Do not use this
> as a programming pattern! Let the runtime system refresh the user interface automatically when
> needed, and ONLY use `ui.Interface.refresh()` to update the display of messages like
> `"Please wait..."` before a lengthy processing.

## Related links

**Related concepts**  

[Refreshing the display when processing](../11_user-interface/2224-refreshing-the-display-when-processing.md "This topic explains when to use the ui.Interface.refresh() method.")
