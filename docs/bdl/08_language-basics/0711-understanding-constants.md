---
title: "Understanding constants"
source: "fgl-topics/c_fgl_Constants_002.html"
breadcrumb: "Language basics > Constants > Understanding constants"
type: "concept"
---

# Understanding constants

> This is an introduction to constant definition.

A constant defines a read-only value identified by a name. A constant is similar to a variable,
except that its value cannot be modified by program code.

A constant is defined with an [identifier](0551-identifiers.md "A Genero BDL identifier is a sequence of characters used to identify a program entity."), and optional [data type](0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") and a [literal value](0585-literals.md "Describes the syntax of literals (constant values) to be used in sources.").

Constants are typically used to define common invariable values that will be used at several
places in a program:

```
CONSTANT PI DECIMAL(12,10) = 3.1415926,
       MAX_SIZE INT = 10000,
       ERRMSG = "PROGRAM ERROR: %1" -- type defaults to STRING
```

A good practice is to group constants that belong to the same domain in a given .4gl module.
Defining these constants as `PUBLIC`, and import the module where the constants
are needed.

## Related links

**Related concepts**  

[Importing modules](../09_advanced-features/0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.")
