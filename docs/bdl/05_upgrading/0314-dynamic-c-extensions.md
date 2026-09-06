---
title: "Dynamic C extensions"
source: "fgl-topics/c_fgl_Migrate_to_200_dynamic_cext.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.00 upgrade guide > Dynamic C extensions"
type: "concept"
---

# Dynamic C extensions

> Dynamic C extensions are automatically loaded with IMPORT instructions.

Prior to version 2.00, you had to use [FGLPROFILE](../07_configuration/0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.")
entries to specify Dynamic C extensions to be loaded at runtime.

Starting with version 2.00, Dynamic C extensions are automatically loaded with `IMPORT` instructions. The FGLPROFILE entries
are no longer used.

> **Important:**
>
> Global variables (userData) can no longer be shared between the runtime
> system and the C extensions. You must use functions to pass global variable values.

There is no longer any need to define the FGL\_API\_MAIN macro in the extension interface file.

All C data type definitions are now centralized in the fglExt.h header file;
header files such as Date.h, MyDecimal.h have been removed
from the distribution.

## Related links

**Related concepts**  

[C-Extensions](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.")

[Example 4: Global variables](../08_language-basics/0707-example-4-global-variables.md "Example 4: Global variables")
