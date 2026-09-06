---
title: "Set an HTTP cookie"
source: "fgl-topics/c_gws_client_behavior_007.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Change WS client behavior at runtime > Set an HTTP cookie"
type: "concept"
---

# Set an HTTP cookie

> Manage the cookie binding at runtime using the global endpoint record parameter.

To send an HTTP cookie to the service, set the record `Cookie` member with the
cookie value. If you leave the variable unset, the client won't send any cookie.

Example:

```
LET Calculator_CalculatorPortTypeEndpoint.Binding.Cookie = "MyCookie=AValue"
```

Unset that variable if you don't need the cookie to be sent anymore.
