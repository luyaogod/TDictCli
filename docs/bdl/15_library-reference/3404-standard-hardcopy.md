---
title: "standard.hardCopy"
source: "fgl-topics/c_fgl_frontcall_standard_hardcopy.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.hardCopy"
type: "concept"
---

# standard.hardCopy

> Prints a screenshot of the current window

## Syntax

```
ui.Interface.frontCall("standard", "hardCopy",
 [pgsize], [result])
```

1. pgsize - Pass "1" to adapt the screenshot to the page size.
2. result - The execution status (`TRUE`=success,
   `FALSE`=error).

## Usage

The "`hardCopy`" front call allows you to print a screenshot of the current
window.

The pgsize parameter is optional; either leave out, or enter "1" to
indicate that the screenshot must be adapted to the page size.
