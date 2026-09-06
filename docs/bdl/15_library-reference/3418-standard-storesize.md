---
title: "standard.storeSize"
source: "fgl-topics/c_fgl_frontcall_standard_storesize.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.storeSize"
type: "concept"
---

# standard.storeSize

> Asks GDC to store the size of the current window container.

## Syntax

```
ui.Interface.frontCall("standard", "storeSize",
  [], [result])
```

1. result - The execution status (`TRUE`=success,
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

See `restoreSize` topic for a code example.

## Related links

**Related concepts**  

[standard.restoreSize](3412-standard-restoresize.md "Asks GDC to restore the stored window container size.")
