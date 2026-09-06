---
title: "Elements shared by multiple modules"
source: "fgl-topics/c_fgl_optimization_005.html"
breadcrumb: "Advanced features > Optimization > Runtime system basics > Elements shared by multiple modules"
type: "concept"
description: "By definition, global variables are visible to all modules of a program, and thus shared among all modules of the program. While global variables are an easy way to share data among multiple modules, ..."
---

# Elements shared by multiple modules

By definition, global [variables](../08_language-basics/0686-variables.md "Explains how to define program variables.") are visible to all
modules of a program, and thus shared among all modules of the program. While global variables are
an easy way to share data among multiple modules, it is not recommended that you use too many global
variables.

The [data type](../08_language-basics/0553-primitive-data-types.md "Selecting the correct data type assists you in the input, storage, and display of your data.") definitions are only defined once in
memory and shared by all modules of a program instance. By data type definition we mean the type
descriptions, not the data itself. This applies only to the equivalent data types used in different
modules.
