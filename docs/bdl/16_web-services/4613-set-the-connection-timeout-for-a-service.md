---
title: "Set the connection timeout for a service"
source: "fgl-topics/c_gws_client_behavior_008.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Set the connection timeout for a service"
type: "concept"
---

# Set the connection timeout for a service

> A default connection timeout value is set, but you can configure the value at runtime, if required.

To change the default timeout value to establish a connection to the service, set the record
member `ConnectionTimeout` with the timeout value in seconds. The
`ConnectionTimeout` is a maximum time to establish the connection.

If the server is not available, you can get the response sooner than
`ConnectionTimeout`. The `ConnectionTimeout` also avoids indefinitely
waiting for a response when a server is reachable but not responding to requests.

Example:

```
LET Calculator_CalculatorPortTypeEndpoint.Binding.ConnectionTimeout = 15
```

## Related links

**Related concepts**  

[Set the read and write timeout for a service](4614-set-the-read-and-write-timeout-for-a-service.md "Set the ReadWriteTimeout value to indicate how long you are willing to wait for the server to respond to requests.")
