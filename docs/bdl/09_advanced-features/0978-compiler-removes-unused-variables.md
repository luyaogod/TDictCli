---
title: "Compiler removes unused variables"
source: "fgl-topics/c_fgl_optimization_012.html"
breadcrumb: "Advanced features > Optimization > Optimize your programs > Compiler removes unused variables"
type: "concept"
description: "When declaring a large static array without any reference to that variable in the rest of the module, you will not see the memory grow at runtime. The compiler has removed its definition from the 42m ..."
---

# Compiler removes unused variables

When declaring a large [static array](../08_language-basics/0732-static-arrays.md "Static arrays have a predefined and limited size.") without any
reference to that variable in the rest of the module, you will not see the memory grow at runtime.
The compiler has removed its definition from the 42m module.

To get the defined variable in the .42m module, you must at least use it
once in the source (for example, with a [`LET`](../08_language-basics/0701-let.md "The LET statement assigns values to variables.") statement). Note that memory might only be allocated when reaching the
lines using the variable.
