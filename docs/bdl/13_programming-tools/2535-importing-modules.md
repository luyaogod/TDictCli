---
title: "Importing modules"
source: "fgl-topics/c_fgl_CompilingPrograms_003.html"
breadcrumb: "Programming tools > Compiling source files > Importing modules"
type: "concept"
---

# Importing modules

> Describes how to define module interdependence with IMPORT FGL.

With the `IMPORT FGL` instruction, module symbols such as variables, types and
constants can be referenced in the importing module.

The following source example imports the `myutils` and `account`
modules, and uses the `init()` and `set_account()` functions of the
imported modules.

The first function call is qualified with the module name - this is optional but required to
resolve ambiguities when the same function name is used by different
modules:

```
IMPORT FGL myutils 
IMPORT FGL account 
MAIN
  CALL myutils.init()
  CALL set_account("CFX4559")
  ...
END MAIN
```

For more details, see [IMPORT FGL](../09_advanced-features/0815-import-fgl.md "The IMPORT FGL instruction imports module symbols.").
