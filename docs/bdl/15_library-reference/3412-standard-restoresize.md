---
title: "standard.restoreSize"
source: "fgl-topics/c_fgl_frontcall_standard_restoresize.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.restoreSize"
type: "concept"
---

# standard.restoreSize

> Asks GDC to restore the stored window container size.

## Syntax

```
ui.Interface.frontCall("standard", "restoreSize",
  [delay], [result])
```

1. delay - Define the delay (in milliseconds) used to revert the window size.
   The window will smoothly shrink or grow to reach the saved size instead of having its new size
   immediately.
2. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`storeSize`" and "`restoreSize`" front calls can be used to
save temporarily the size of the [window
container](../11_user-interface/1563-containers-for-program-windows.md "Program windows are displayed in window containers by the front-end."), to restore the container to the saved size.

> **Important:**
>
> The "`storeSize`" and "`restoreSize`" front calls are only
> available with the GDC front end.

The "`restoreSize`" front call takes an optional parameter to define the delay in
milliseconds. The window container will then smoothly shrink or grow to reach the saved size.

Invoking "`restoreSize`" will have no effect, if there was no prior call to
"`storeSize`", or if "`storeSize`" was called for a different window
container (see [`desktopMultiWindow`](../11_user-interface/1652-userinterface-style-attributes.md) style attribute).

Note that the stored window size is just a hint for a desired size: The layout has always higher
priority. When restoring a window that is too small for the form layout, the window container will
not shrink to the stored size.

## Example

Test the sample as follows:

- Compile and start the program with display to GDC
- Resize the window container
- Store the current window size with the "store" button
- Resize the window container to a different size
- Restore the window size with the "restore" button

The form.per file:

```
LAYOUT
GRID
{
[f1    |f2                ]
[f3                       ]
[                         ]
[                         ]
}
END
ATTRIBUTES
EDIT f1 = FORMONLY.pkey;
EDIT f2 = FORMONLY.name;
TEXTEDIT f3 = FORMONLY.comm, STRETCH=BOTH;
END
```

The main.4gl file:

```
MAIN
    DEFINE rec RECORD
               pkey INTEGER,
               name VARCHAR(40),
               comm STRING
           END RECORD
    DEFINE ret INTEGER

    OPEN FORM f1 FROM "form"
    DISPLAY FORM f1

    INPUT BY NAME rec.*
        ON ACTION store
           CALL ui.interface.frontCall("standard","storeSize",[],[ret])
        ON ACTION restore
           CALL ui.interface.frontCall("standard","restoreSize",[500],[ret])
    END INPUT

END MAIN
```
