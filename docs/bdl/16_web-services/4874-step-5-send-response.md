---
title: "Step 5: Send response"
source: "fgl-topics/c_gws_rest_server_send_response.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web Services server application > Step 5: Send response"
type: "concept"
---

# Step 5: Send response

> Having completed the processing, the server sends the response back to the client.

The server application has the value or values that must be returned to the client. It now needs
to format and send the response.

In the demo server application, the `setResponseHeader()` method sets the
Content-Type to JSON.

```
CALL req.setResponseHeader("Content-Type","application/json")
```

The sendTextResponse() sends three things: a code, a description, and the
response data. In this case, the data is formatted as JSON using the
util.JSON.stringify() method, which transforms a record variable into a flat JSON
formatted
string.

```
CALL req.sendTextResponse(200, "OK", util.JSON.stringify(add_out))
```

## Related links

**Related concepts**  

[Step 6: Provide information about your service](4875-step-6-provide-information-about-your-service.md "You must provide information about the services offered by your application to the developers of the Web services client applications that will interact with your server.")
