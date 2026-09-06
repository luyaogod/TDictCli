---
title: "Front-end protocol errors"
source: "fgl-topics/c_fgl_feconn_fe_errors.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Front-end protocol errors"
type: "concept"
description: "When the front-end receives an invalid instruction from the runtime system, it stops the application connection. The runtime system then stops and displays error -6313 with an additional message, for ..."
---

# Front-end protocol errors

When the front-end receives an invalid instruction from the runtime system, it stops the
application connection.

The runtime system then stops and displays error [-6313](../15_library-reference/4483-genero-bdl-errors.md) with an additional message,
for example:

```
Program stopped at 'myprog.4gl', line number 675.
 FORMS statement error number -6313.
 The User Interface has been destroyed: Unexpected interface version sent 
 by the runtime system.
```

When the runtime system receives an invalid AUI event from a front-end, it will raise a C
assertion and produce a core file on UNIX®
systems.

## Related links

**Related concepts**  

[Exceptions](../09_advanced-features/0848-exceptions.md "Describes exception (error) handling in the programs.")
