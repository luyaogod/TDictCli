---
title: "Change the HTTP protocol version"
source: "fgl-topics/c_gws_client_behavior_006.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Change the HTTP protocol version"
type: "concept"
---

# Change the HTTP protocol version

> The version parameter allows you to change the HTTP version binding if required.

To communicate with a service that communicates only on a given version of HTTP, set the record
`Version` member with the desired value. If you leave the variable unset, the client
will communicate in HTTP 1.1.

Example:

```
LET Calculator_CalculatorPortTypeEndpoint.Binding.Version = "1.0"
```

If you do not want the request to be split into chunks, set the HTTP protocol version to 1.0.
