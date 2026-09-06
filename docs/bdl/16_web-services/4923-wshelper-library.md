---
title: "WSHelper library"
source: "fgl-topics/c_gws_wshelper.html"
breadcrumb: "Web services > Reference > WSHelper library"
type: "concept"
---

# WSHelper library

> The WSHelper library provides a set of functions to help with web services.

The WSHelper library is delivered in the FGLGWS package. It is located in
$FGLDIR/lib. To use a WSHelper function, declare the module with the [`IMPORT FGL`](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
instruction:

```
IMPORT FGL WSHelper
#...
   CALL WSHelper.SplitURL("https://cube.strasbourg.4js.com:3128/GWS-492/TestSplitURL/test1")
```

It
is recommended that the WSHelper.42m library file is imported into every Genero
Web Services server or client program.
