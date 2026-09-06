---
title: "Step 4: Process the request"
source: "fgl-topics/c_gws_rest_server_process.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web Services server application > Step 4: Process the request"
type: "concept"
---

# Step 4: Process the request

> Having parsed the HttpServiceRequest object into its parts, you can now process the request.

In the previous step, the `m_reqInfo` record was populated with the values of the
`HttpServiceRequest` object. The application now needs to work with these values.

Each Web services server application will be different in the type of processing that is done.
For completeness, this topic will highlight some of the processing that is happening in the
calculator demo server application.

Despite having variables in the `m_reqInfo` record, four additional
`STRING` variables are defined to hold the method, the URL, the query, and the output
format. In addition, an `INTEGER` variable is defined to hold an index value, used by
some of the custom processing
functions.

```
DEFINE method STRING 
DEFINE url STRING
DEFINE qry STRING
DEFINE acc STRING
DEFINE idx INT
```

With the record populated, the server application then populates the additional
variables.

```
# Get the type of method
LET method = m_reqInfo.method
# Get the request path
LET url = m_reqInfo.path
# Get the query string
LET qry = m_reqInfo.query
# Get the output format
LET acc = m_reqInfo.outformat
```

In the demo server application, the `CASE` statement only tests for the GET
method.

```
CASE method
    WHEN "GET"
      ... processing instructions
    OTHERWISE
      CALL setError("Unknown request:\n" || url || "\n" || method)
      LET err.code = -3
      LET err.desc = ERR_METHOD
      CALL req.sendTextResponse(200,"OK",util.JSON.stringify(err))
END CASE
```

For the remainder of the processing code, the demo server application:

- Determines whether the request is for addition, subtraction, multiplication, or division
- Retrieves the values for the variables a and b
- Handles errors appropriately

For example, the source code that would process an addition
request:

```
IF url.getIndexOf("/add?",1) > 0 THEN
    LET idx = getParameterIndex("a") 
    IF idx = 0 THEN
        LET add_out.status.code = -1
        LET add_out.status.desc = ERR_PARAM_A
        CALL req.setResponseHeader("Content-Type","application/json")
        CALL req.sendTextResponse(200, "OK", util.JSON.stringify(add_out))
        EXIT CASE
    ELSE
        LET add_in.a = getParameterValue(idx)
    END IF
    LET idx = getParameterIndex("b")
    IF idx = 0 THEN
        LET add_out.status.code = -1
        LET add_out.status.desc = ERR_PARAM_B
        CALL req.setResponseHeader("Content-Type","application/json")
        CALL req.sendTextResponse(200, "OK", util.JSON.stringify(add_out))
        EXIT CASE
    ELSE
        LET add_in.b = getParameterValue(idx)
    END IF
    CALL add()
    ... send response
END IF
```

In the next step the server sends the response back to the client, [Step 5: Send response](4874-step-5-send-response.md "Having completed the processing, the server sends the response back to the client.").
