---
title: "Scope of module symbols (PRIVATE/PUBLIC)"
source: "fgl-topics/c_fgl_programs_IMPORT_FGL_symb_scope.html"
breadcrumb: "Advanced features > Importing modules > Importing FGL modules > Scope of module symbols (PRIVATE/PUBLIC)"
type: "concept"
---

# Scope of module symbols (PRIVATE/PUBLIC)

> The PRIVATE/PUBLIC modifiers can be used to hide / publish symbols to other modules.

For backward compatibility, functions are by default public. Module variables, types and
constants are by default private.

The following example declares a module variable that can be used by other modules, and a private
function to be used only locally:

```
PUBLIC DEFINE custlist DYNAMIC ARRAY OF RECORD
  id INT,
  name VARCHAR(50),
  address VARCHAR(200)
END RECORD
...
PRIVATE FUNCTION myfunction()
...
END FUNCTION
```
