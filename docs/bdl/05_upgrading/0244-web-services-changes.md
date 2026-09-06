---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_240_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.40 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 2.40.

## Modification to server location at runtime on client side

> **Important:**
>
> It is recommended to regenerate all client stubs in your application using the [fglwsdl](../13_programming-tools/2523-fglwsdl.md "The fglwsdl tool produces web services stub files for client or server programs (from WSDL / XSD).") tool.

If you have modified the server location at runtime via the generated global variable in your
client application, you MUST apply following modification:

- Prior to version 2.40, you had something similar to the
  following:

  ```
  LET Calculator_CalculatorPortTypeLocation = "http://host:port/Calculator"
  ```
- Starting with version 2.40, you must have something similar to the
  following:

  ```
  LET Calculator_CalculatorPortTypeEndPoint.Address.Uri =
   "http://host:port/Calculator"
  ```

See [Change client behavior
at runtime](../16_web-services/4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.").
