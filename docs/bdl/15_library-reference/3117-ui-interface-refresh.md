---
title: "ui.Interface.refresh"
source: "fgl-topics/c_fgl_ClassInterface_refresh.html"
breadcrumb: "Library reference > Built-in packages > The ui package > The Interface class > ui.Interface methods > ui.Interface.refresh"
type: "concept"
---

# ui.Interface.refresh

> Synchronize the user interface with the front-end.

## Syntax

```
ui.Interface.refresh()
```

> **Important:**
>
> The AUI tree is automatically synchronized by
> the runtime system, when dialog intruction gives the control back to the end user. The
> `ui.Interface.refresh()` method must only be used in specific cases, to refresh the
> display while processing. For example, to show a "Please wait" message, or to implement a progress
> dialog window with a [`PROGRESSBAR`](../11_user-interface/1698-progressbar-item-type.md "Defines a progress indicator field."). The `ui.Interface.refresh()` method should not
> be called more often than once in one second.

## Usage

The `ui.Interface.refresh()` class method forces a synchronization of the
abstract user interface tree with the front-end. This means that the end user will immediately see
the recent form modifications made by the program.

By default, during an interactive instruction like `DIALOG`, the AUI tree is
refreshed automatically, when the runtime system gets the control back after user code execution.
Consequently, there is no need to call the refresh method in regular code.

With the [`UNBUFFERED`](../11_user-interface/2235-the-buffered-and-unbuffered-modes.md "The buffered and unbuffered mode control the synchronization of program variables and form fields.")
dialog attribute, after the execution of the current dialog block, the values of the program
variables bound to the dialog will be automatically displayed in the corresponding form fields (or
table column cells). The `ui.Interface.refresh()` method only sends the current AUI
tree content to the front-end; it does not synchronize form field content with the values of program
variables, even when using the `UNBUFFERED` attribute. Consequently, to immediately
show new form field values during the execution of a dialog block, the code must explicitly
synchronize the form field with the program variable values, for example with a `DISPLAY BY
NAME`, followed by the refresh method
call:

```
INPUT BY NAME rec.* ATTRIBUTES( UNBUFFERED, WITHOUT DEFAULTS )
  ON ACTION increment
    LET rec.counter = rec.counter + 1
    DISPLAY BY NAME rec.counter -- Display value to form field
    CALL ui.Interface.refresh() -- Send AUI tree to front end
    SLEEP 3 -- Simulate some long processing
END INPUT
```

Similarly, to clear the rows of a `TABLE` container and immediately show the empty
record list during a dialog block, clear the array with `deleteAllRows()`, use
`CLEAR SCREEN ARRAY` to clean up the screen array, and call the refresh
method:

```
DISPLAY ARRAY custlist TO sr_custlist.* ATTRIBUTES(UNBUFFERED)
  ON ACTION new_query
    CALL DIALOG.deleteAllRows("sr_custlist")
    CLEAR SCREEN ARRAY sr_custlist.*
    CALL ui.Interface.refresh()
    SLEEP 3 -- Simulate some long processing
END DISPLAY
```

For backward compatibility, FGL also implements the `fgl_refresh()` function. The
`fgl_refresh()` function has exactly the same behavior than
`ui.Interface.refresh()`.

## Example

```
MAIN
    DEFINE i INT
    FOR i = 1 TO 10
        DISPLAY SFMT("Please wait, doing step #%1", i) AT 1,1
        CALL ui.Interface.refresh()
        SLEEP 1
    END FOR
END MAIN
```

## Related links

**Related concepts**  

[Refreshing the display when processing](../11_user-interface/2224-refreshing-the-display-when-processing.md "This topic explains when to use the ui.Interface.refresh() method.")

[User interruption handling](../11_user-interface/2225-user-interruption-handling.md "Allow the end user to cancel a dialog or a long running procedure.")
