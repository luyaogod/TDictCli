---
title: "Step 1: Generate the client stub from your WS-Addressing stateful service"
source: "fgl-topics/c_gws_stateful_services_021.html"
breadcrumb: "Web services > Concepts > Stateful SOAP Web services > WS-Addressing 1.0 stateful services > Client side > Step 1: Generate the client stub"
type: "concept"
description: "Use the fglwsdl tool as usual. It will detect that the service returns a W3CEndpointReference and generate the appropriate code. The WSDL imports the WS-Addressing 1.0 schema, so the fglwsdl tool ..."
---

# Step 1: Generate the client stub from your WS-Addressing stateful service

Use the fglwsdl tool as usual. It will detect that the service returns a
W3CEndpointReference and generate the appropriate code.

The WSDL imports the WS-Addressing 1.0 schema, so the fglwsdl tool requires
access to the W3C server. Use the option `-proxy` if you need to connect via a proxy
server.

For
example:

```
$ fglwsdl -o ws_stub http://localhost:8090/StatefulWSAddressingService?WSDL
```

The generated .inc file contains a variable of type [tWSAGlobalEndpointType](4607-change-ws-client-behavior-at-runtime.md "Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.") to be used to transmit the
WS-Addressing 1.0 reference parameters.

## Example of a global variable name

```
DEFINE StatefulWSAddressingService_StatefulWSAddressingServicePortTypeEndpoint
   tGlobalWSAEndpointType
```

## Related links

**Related concepts**  

[Step 2: Create the MAIN application](4528-step-2-create-the-main-application.md "Step 2: Create the MAIN application")
