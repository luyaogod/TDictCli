---
title: "standard.cbPaste"
source: "fgl-topics/c_fgl_frontcall_standard_cbpaste.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.cbPaste"
type: "concept"
---

# standard.cbPaste

> Pastes the content of the clipboard to the current field.

## Syntax

```
ui.Interface.frontCall("standard", "cbPaste",
   [], [result])
```

1. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`cbPaste`" front call pastes the content of the clipboard to the current
field.

This front call takes no input parameters.

> **Note:**
>
> In order to use clipboard front calls, you need a secure environment. For web applications
> (using the GAS), you need to reach the programs from an https scheme or with
> localhost in the URL. There is no restriction when using GDC, GMA or GMI
> front-ends.
