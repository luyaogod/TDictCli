---
title: "Step 1: Generate the client stub from your stateful service"
source: "fgl-topics/c_gws_stateful_services_051.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > Stateful services based on HTTP cookies > Client side > Step 1: Generate the client stub"
type: "concept"
description: "Use the fglwsdl tool as usual. For example: $ fglwsdl -o ws_stub http://localhost:8090/StatefulCookieService?WSDL The generated .inc file contains a variable of type tGlobalEndpointType to be used to ..."
---

# Step 1: Generate the client stub from your stateful service

Use the fglwsdl tool as usual.

For
example:

```
$ fglwsdl -o ws_stub http://localhost:8090/StatefulCookieService?WSDL
```

The generated .inc file contains a variable of type [tGlobalEndpointType](4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.") to be
used to transmit the HTTP Cookie.

## Example of a global variable name

```
DEFINE StatefulCookieService_StatefulCookieServicePortTypeEndpoint
 tGlobalEndpointType
```

## Related links

**Related concepts**  

[Step 2: Create the MAIN application](4541-step-2-create-the-main-application.md "Step 2: Create the MAIN application")
