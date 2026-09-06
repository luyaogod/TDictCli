---
title: "standard.cbClear"
source: "fgl-topics/c_fgl_frontcall_standard_cbclear.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.cbClear"
type: "concept"
---

# standard.cbClear

> Clears the content of the clipboard.

## Syntax

```
ui.Interface.frontCall("standard", "cbClear",
  [], [result])
```

1. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`cbClear`" front call clears the content of the clipboard. This front call
takes no input parameters.

> **Note:**
>
> In order to use clipboard front calls, you need a secure environment. For web applications
> (using the GAS), you need to reach the programs from an https scheme or with
> localhost in the URL. There is no restriction when using GDC, GMA or GMI
> front-ends.
