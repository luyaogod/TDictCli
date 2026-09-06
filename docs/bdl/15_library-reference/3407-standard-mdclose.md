---
title: "standard.mdClose"
source: "fgl-topics/c_fgl_frontcall_standard_mdclose.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.mdClose"
type: "concept"
---

# standard.mdClose

> Unloads a DLL or shared library front call module.

## Syntax

```
ui.Interface.frontCall("standard", "mdClose",
 [name], [result])
```

1. name - The name of the module to be closed.
2. result - The result status (0 = success, -1 = module not found, -2 = cannot
   unload (busy)).

## Usage

Front call modules are loaded on demand. After calling a function of a specific module, you can
use the "`mdClose`" front call to unload the shared library and save resources.
