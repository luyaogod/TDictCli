---
title: "standard.cbAdd"
source: "fgl-topics/c_fgl_frontcall_standard_cbadd.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.cbAdd"
type: "concept"
---

# standard.cbAdd

> Adds to the content of the clipboard.

## Syntax

```
ui.Interface.frontCall("standard", "cbAdd",
  [text], [result])
```

1. text - The text to be added.
2. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`cbAdd`" front call adds the text passed as parameter to
the content of the clipboard of the front-end platform.

> **Note:**
>
> In order to use clipboard front calls, you need a secure environment. For web applications
> (using the GAS), you need to reach the programs from an https scheme or with
> localhost in the URL. There is no restriction when using GDC, GMA or GMI
> front-ends.
