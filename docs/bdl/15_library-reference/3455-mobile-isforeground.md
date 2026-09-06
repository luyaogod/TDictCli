---
title: "mobile.isForeground"
source: "fgl-topics/c_fgl_frontcall_mobile_isforeground.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.isForeground"
type: "concept"
---

# mobile.isForeground

> Indicates if the mobile app is in foreground mode.

## Syntax

```
ui.Interface.frontCall("mobile", "isForeground",
   [],  [result] )
```

1. result - Is set to `TRUE`, if the app is in foreground mode,
   or `FALSE` if in background mode.

## Usage

The "`mobile.isForeground`" front call is a synomym for [standard.isForeground](3405-standard-isforeground.md "Indicates if the app is in foreground mode.").
