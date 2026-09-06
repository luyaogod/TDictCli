---
title: "Understanding front calls"
source: "fgl-topics/c_fgl_frontcalls_intro.html"
breadcrumb: "Advanced features > Front calls > Understanding front calls"
type: "concept"
---

# Understanding front calls

> Front calls execute a native function on the front-end platform.

In your Genero program, use the `ui.Interface.frontCall()` class method to invoke
front-end functions. When calling a user function from programs, specify a module name and a
function name. Input and output parameters can be passed/returned in order to transmit/receive
values to/from the front-end. A typical example is an "open file" dialog window that allows you to
select a file from the front-end workstation file system.

> **Important:**
>
> Front calls can be specific to the platform or front-end technology. For
> example, it is not possible to perform a [shellexec](../15_library-reference/3417-standard-shellexec.md "Opens a file on the front-end platform with the program associated to the file extension.") front call with the GBC front-end, when using the GAS.

A set of front-end functions is [built-in](../15_library-reference/3386-built-in-front-calls-summary.md "Various front-end functions are implemented within Genero front-ends.") by
default in front-ends. However, it is possible to write your
[own functions](../14_extending-the-language/2718-user-defined-front-calls.md "Front-ends can be extended with custom functions to access specific features.") in order to extend the
front-end possibilities.

## Related links

**Related concepts**  

[The abstract user interface tree](../11_user-interface/1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")
