---
title: "Set a time period for the response"
source: "fgl-topics/c_gws_client_tutorial_007.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services client application > Steps to write a GWS client > Set a time period for the response"
type: "concept"
---

# Set a time period for the response

> To protect against remote server failure or unavailability, set a timeout value that indicates how long you are willing to wait for the server to respond to your request.

Use the `SetOption()` method of the `WebServiceEngine` class to set
the `readwritetimeout` option.

For example, to wait no more than 10
seconds:

```
CALL com.WebServiceEngine.SetOption( "readwritetimeout", 10 )
```

A timeout value of **-1** means "wait forever". This is the
default value.

## Related links

**Related concepts**  

[The WebServiceEngine class](../15_library-reference/3785-the-webserviceengine-class.md "The com.WebServiceEngine class provides an interface to manage the Web Services engine.")
