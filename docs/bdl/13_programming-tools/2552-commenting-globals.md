---
title: "Commenting globals"
source: "fgl-topics/c_fgl_AutoDoc_011.html"
breadcrumb: "Programming tools > Source documentation > Adding comments to sources > Commenting globals"
type: "concept"
description: "Source comments can be added to symbols defined in a GLOBALS / END GLOBALS block. The global symbols included from an external file with GLOBALS \" filename \" will not be part of the module ..."
---

# Commenting globals

Source comments can be added to symbols defined in a `GLOBALS / END GLOBALS`
block.

The global symbols included from an external file with `GLOBALS
"filename"` will not be part of the module documentation.

In order to have fglcomp --build-doc work well,
there must be an initial #+ comment line for the module description, before any other constant,
variable, type or function documentation directives.

## Example

```
GLOBALS
    #+ a global variable
    DEFINE globalVar1 INT
    #+ a global type
    TYPE GlobalType1 RECORD
        f1, f2 INT
    END RECORD
    #+ a global constant
    CONSTANT globalConst1 = 9999
END GLOBALS
```
