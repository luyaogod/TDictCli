---
title: "standard.cbGet"
source: "fgl-topics/c_fgl_frontcall_standard_cbget.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.cbGet"
type: "concept"
---

# standard.cbGet

> Gets the content of the clipboard.

## Syntax

```
ui.Interface.frontCall("standard", "cbGet",
   [], [text])
```

1. text - The text found in the clipboard.

## Usage

The "`cbGet`" front call returns the current content of the clipboard.

This front call takes no input parameters.

> **Note:**
>
> In order to use clipboard front calls, you need a secure environment. For web applications
> (using the GAS), you need to reach the programs from an https scheme or with
> localhost in the URL. There is no restriction when using GDC, GMA or GMI
> front-ends.
