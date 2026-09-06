---
title: "Set server endpoint"
source: "fgl-topics/c_gws_rest_high_level_client_set_endpoint.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Change REST client behavior at runtime > Set server endpoint"
type: "concept"
---

# Set server endpoint

> Set the REST Web service location from client app at runtime.

To change the server location at runtime, set the record `Uri` element with a
valid URL of another service. If you leave the variable unset, the client will connect to the server
URL defined in the stub file at code generation time.

## Set server address

In the stub file, an `Endpoint` variable of type
`tGlobalEndpointType` is defined as public. You can set the address of the Web
service via the client application with for example:

```
LET clientStub.Endpoint.Address.Uri = 
  "http://zeus:1111/mydomain/MyService"
```
