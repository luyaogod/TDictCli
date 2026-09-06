---
title: "Refreshing the display when processing"
source: "fgl-topics/c_fgl_prog_dialogs_aui_refresh.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > Refreshing the display when processing"
type: "concept"
---

# Refreshing the display when processing

> This topic explains when to use the ui.Interface.refresh() method.

The [AUI tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.") on the front-end is synchronized with
the runtime system AUI tree, when a user interaction instruction (dialog) gives the control back to
the end user.

This means that the end user will not see any changes on the screen, as long as the program is
doing batch processing, until an interactive statement is reached and waits for an event triggered
by the end user.

To show something on the screen while the program is running in a batch procedure, force AUI tree
synchronization by calling the [`ui.Interface.refresh()`](../15_library-reference/3117-ui-interface-refresh.md "Synchronize the user interface with the front-end.") method.

> **Important:**
>
> The AUI tree is automatically synchronized by
> the runtime system, when dialog intruction gives the control back to the end user. The
> `ui.Interface.refresh()` method must only be used in specific cases, to refresh the
> display while processing. For example, to show a "Please wait" message, or to implement a progress
> dialog window with a [`PROGRESSBAR`](1698-progressbar-item-type.md "Defines a progress indicator field."). The `ui.Interface.refresh()` method should not
> be called more often than once in one second.

## Example

Form file form.per:

```
LAYOUT
GRID
{
[pb                              ]
         [b1           ]
}
END
END
ATTRIBUTES
PROGRESSBAR pb = FORMONLY.progress,
                 VALUEMIN=0, VALUEMAX=30;
BUTTON b1: interrupt, TEXT="Interrupt";
END
```

Program file
main.4gl:

```
MAIN
    DEFINE progress INTEGER
    DEFER INTERRUPT
    OPTIONS INPUT WRAP
    OPEN FORM f FROM "form"
    DISPLAY FORM f
    LET int_flag = FALSE
    FOR progress = 1 TO 30
        IF int_flag THEN
            EXIT FOR
        END IF
        DISPLAY BY NAME progress
        CALL ui.Interface.refresh()
        SLEEP 1
    END FOR
END MAIN
```

## Related links

**Related concepts**  

[Dynamic Dialogs](2425-dynamic-dialogs.md "Dialogs can be created at runtime with the ui.Dialog class.")
