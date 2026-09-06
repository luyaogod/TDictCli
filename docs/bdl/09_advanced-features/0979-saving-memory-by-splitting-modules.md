---
title: "Saving memory by splitting modules"
source: "fgl-topics/c_fgl_optimization_013.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Saving memory by splitting modules"
type: "concept"
description: "Program modules ( .42m ) are loaded dynamically on demand. If a program only needs some independent functions of a given module, all module resources will be allocated just to call these functions. By ..."
---

# Saving memory by splitting modules

Program modules (.42m) are loaded dynamically on demand. If a program only
needs some independent functions of a given module, all module resources will be allocated just to
call these functions. By independent, we mean functions that do not use module objects such as
variables defined outside function or [SQL cursors](../10_sql-support/1148-result-set-processing.md "Shows how to fetch rows from a database query."). To
avoid unnecessary resource allocation, you can extract these independent functions into another
module and save a lot of memory at runtime.

If you are using .42x libraries, it is recommended that you create libraries
with the 42m modules that belong to the same functionality group. For example, group all accounting
modules together in an accounting library. By doing this, programmers using the
.42x libraries are not dependent from module reorganizations.

Libraries are supported for backward compatibility, it is recommended that you consider using the [`IMPORT FGL`](0813-importing-modules.md "Use the IMPORT ... instruction to import BDL, C or Java external modules in the current module.") instruction to define
module dependency and get modules loaded dynamically when needed.
