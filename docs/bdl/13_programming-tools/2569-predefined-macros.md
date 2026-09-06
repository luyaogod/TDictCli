---
title: "Predefined macros"
source: "fgl-topics/c_fgl_preprocessor_009.html"
breadcrumb: "Programming tools > Source preprocessor > Predefined macros"
type: "concept"
---

# Predefined macros

> A set of predefined preprocessor macros are available.

The preprocessor predefines 2 macros:

1. `__LINE__ expands to the current line number. Its definition
   changes with each new line of the code.`
2. `__FILE__ expands to the name of the current file as a string constant. For example:
   "subdir/file.inc"`

These macros are often used to generate error messages.

An `&include` directive changes the values of `__FILE__` and `__LINE__` to
correspond to the included file.
