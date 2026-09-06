---
title: "standard.clearFileCache"
source: "fgl-topics/c_fgl_frontcall_standard_clearfilecache.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.clearFileCache"
type: "concept"
---

# standard.clearFileCache

> Clears the local file cache.

> **Important:**
>
> This feature is provided for development and
> debug purpose, and must not be used in a production environment.

## Syntax

```
ui.Interface.frontCall("standard", "clearFileCache",
  [], [result])
```

1. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`clearFileCache`" front call removes all the local resource files cached by
the front-end.

The resource files (like images, font files) transmitted by the runtime system to the front-end
are automatically cached on the front-end workstation.

This front call is only supported by the GDC front-end, and should only be used in development or
for debug purpose.

## Related links

**Related information**  

[The resource file cache of the front-end](../11_user-interface/1586-providing-the-image-resource.md "The resource file cache of the front-end")
