---
title: "Start a GAS app from GWA"
source: "fgl-topics/c_fgl_gwa_application_runonserver.html"
breadcrumb: "Genero Web applications > Start a GAS app from GWA"
type: "concept"
---

# Start a GAS app from GWA

> The runOnServer front call allows you to start an application via the Genero Application Server (GAS) from a Genero Web application (GWA).

A [`runOnServer`](../15_library-reference/3458-mobile-runonserver.md "Run an application from the Genero Application Server using the specified URL.") call
in GWA follows the same basic design as in Genero Mobile for iOS (GMI) and Android™ (GMA), but there are a few important differences to
note:

1. First, in GWA both `RUN` and `RUN ... WITHOUT WAITING` may be used
   in any combination in the application being called by `runOnServer`; there is no
   restriction that forces one or the other. All combinations to run multiple applications server side
   are possible.
2. Second, the `runOnServer` front call ends when all server side applications which
   were started by the `runOnServer` application are terminated in addition to the
   application itself. So all applications started with `RUN ... WITHOUT WAITING` need
   to be terminated in addition to the `runOnServer` application.

Finally, when supplying the GAS URL argument in GWA you may omit the full origin (for example,
`http://host:port`) and use a simple path such
as
`/ua/r/simple`:

```
CALL ui.Interface.frontCall("mobile","runonserver",["/ua/r/simple"],[])
```

If you use a full origin URL, ensure you use the same origin as the GWA application which invokes
the `runOnServer` call; otherwise, you will encounter an error. Both the GWA
application URL and the `runOnServer` application URL must have the same origin.

In contrast, GMI/GMA cannot omit the origin because the native mobile app always runs from a
different origin.

## Related links

**Related concepts**  

[mobile.runOnServer](../15_library-reference/3458-mobile-runonserver.md "Run an application from the Genero Application Server using the specified URL.")

[Executing sub-programs from a parent program with RUN](5141-executing-programs-with-run.md "Sub-programs can be executed from a parent program with the RUN instruction.")
