---
title: "Change WS client behavior at runtime"
source: "fgl-topics/c_gws_client_behavior_001.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime"
type: "concept"
---

# Change WS client behavior at runtime

> Various aspects of access to a Web Service may be implemented on the client side at runtime using global endpoint records.

Genero Web Services generates a global record called [tGlobalEndpointType](4608-global-endpoint-type-definition.md "The client stub references a global endpoint user-defined type, WSHelper.tGlobalEndpointType.") to change the client
behavior at runtime without the need for modifying any generated client stub.

If WS-Addressing 1.0 is enabled, the global generated record is called [tWSAGlobalEndpointType](4609-ws-addressing-1-0-global-endpoint-type-definition.md "The client stub references a global endpoint type for WS-Addressing, WSHelper.tGlobalWSAEndpointType.").

If needed, you can also access the HTTP layer via the `Request` and
`Response` record of the binding record. See [Access HTTP request and response headers for a service](4615-access-http-request-and-response-headers-for-a-service.md "Configure additional headers for requests and responses by adding them to the global endpoint record.").

## Child topics

- [Global Endpoint type definition](4608-global-endpoint-type-definition.md): The client stub references a global endpoint user-defined type, WSHelper.tGlobalEndpointType.
- [WS-Addressing 1.0 Global Endpoint type definition](4609-ws-addressing-1-0-global-endpoint-type-definition.md): The client stub references a global endpoint type for WS-Addressing, WSHelper.tGlobalWSAEndpointType.
- [Change server location](4610-change-server-location.md): If the Web service server location changes, you can update the address in the global endpoint record.
- [Change the HTTP protocol version](4611-change-the-http-protocol-version.md): The version parameter allows you to change the HTTP version binding if required.
- [Set an HTTP cookie](4612-set-an-http-cookie.md): Manage the cookie binding at runtime using the global endpoint record parameter.
- [Set the connection timeout for a service](4613-set-the-connection-timeout-for-a-service.md): A default connection timeout value is set, but you can configure the value at runtime, if required.
- [Set the read and write timeout for a service](4614-set-the-read-and-write-timeout-for-a-service.md): Set the ReadWriteTimeout value to indicate how long you are willing to wait for the server to respond to requests.
- [Access HTTP request and response headers for a service](4615-access-http-request-and-response-headers-for-a-service.md): Configure additional headers for requests and responses by adding them to the global endpoint record.
