---
title: "Set the read and write timeout for a service"
source: "fgl-topics/c_gws_client_behavior_009.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Set the read and write timeout for a service"
type: "concept"
---

# Set the read and write timeout for a service

> Set the ReadWriteTimeout value to indicate how long you are willing to wait for the server to respond to requests.

To change the default time the application waits for a response from a service, set the
`ReadWriteTimeout` parameter with a time out value in seconds.

Example:

```
LET Calculator_CalculatorPortTypeEndpoint.Binding.ReadWriteTimeout = 5
```

## Related links

**Related concepts**  

[Set a time period for the response](4603-set-a-time-period-for-the-response.md "To protect against remote server failure or unavailability, set a timeout value that indicates how long you are willing to wait for the server to respond to your request.")

[Set the connection timeout for a service](4613-set-the-connection-timeout-for-a-service.md "A default connection timeout value is set, but you can configure the value at runtime, if required.")
