---
title: "standard.isForeground"
source: "fgl-topics/c_fgl_frontcall_standard_isforeground.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.isForeground"
type: "concept"
---

# standard.isForeground

> Indicates if the app is in foreground mode.

## Syntax

```
ui.Interface.frontCall("standard", "isForeground",
   [],  [result] )
```

1. result - Is set to `TRUE`, if the app is in foreground mode,
   or `FALSE` if in background mode.

## Usage

The "`isForeground`" front call checks if the app is currently in foreground or
background mode.

```
DEFINE fg BOOLEAN
CALL ui.Interface.frontCall("standard", "isForeground", []. [fg] )
```

Use the `isForeground` front call in conjunction with the [`enterforeground` and `enterbackground`](../09_advanced-features/0829-executing-programs.md) predefined
actions.

## Related links

**Related concepts**  

[Background/foreground modes](../17_mobile-applications/5094-background-foreground-modes.md "Describes how to handle background or foreground modes in mobile apps.")

[mobile.isForeground](3455-mobile-isforeground.md "Indicates if the mobile app is in foreground mode.")
