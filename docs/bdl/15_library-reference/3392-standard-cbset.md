---
title: "standard.cbSet"
source: "fgl-topics/c_fgl_frontcall_standard_cbset.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.cbSet"
type: "concept"
---

# standard.cbSet

> Set the content of the clipboard.

## Syntax

```
ui.Interface.frontCall("standard", "cbSet",
 [text], [result])
```

1. text - The text to be set.
2. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`cbSet`" front call sets the content of the clipboard with the text passed
as parameter.

> **Note:**
>
> In order to use clipboard front calls, you need a secure environment. For web applications
> (using the GAS), you need to reach the programs from an https scheme or with
> localhost in the URL. There is no restriction when using GDC, GMA or GMI
> front-ends.
