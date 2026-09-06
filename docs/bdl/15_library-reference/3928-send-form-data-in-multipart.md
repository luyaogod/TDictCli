---
title: "Examples: Send multipart form data"
source: "fgl-topics/c_gws_ComHTTPPart_example_multipart_simple_form.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > Examples (HttpPart) > Send form data in multipart"
type: "concept"
---

# Examples: Send multipart form data

> Example of client using methods of the com.HttpPart class to send form data in multipart.

This example consists of a client sending a log-in request in a
`multipart/form-data` request body. The user name is sent in one part and the
password is sent in another part of the request body. A `Content-Disposition`
header is set to provide information for each subpart of the form. The `name`
value ("user" and "password") must match what the server is expecting.

## Client Application

```
IMPORT com

MAIN

    DEFINE res INTEGER
    DEFINE msg STRING

    CALL send_rest_post("myuser","mypassword") RETURNING res, msg

    DISPLAY SFMT("Result: %2 (%1)", res, msg)

END MAIN

FUNCTION send_rest_post(r_param)

    DEFINE url STRING
    DEFINE result, tmp STRING
    DEFINE req com.HttpRequest
    DEFINE rsp com.HttpResponse
    DEFINE prt com.HttpPart
    DEFINE r_param RECORD
             user STRING,
             password STRING
           END RECORD

    # create the URL
    LET url="https://server1/myapi/api_login.php"
   
    TRY
        # prepare request with multipart
        LET req = com.HttpRequest.Create(url)
        CALL req.setMethod("POST")
        CALL req.setHeader("Accept", "application/json, text/plain")
        CALL req.setMultipartType("form-data", NULL, NULL)
       
        # configure 1st part
        LET prt = com.HttpPart.CreateFromString(r_param.user)
        CALL prt.setHeader(
            "Content-Disposition", "form-data; name=\"user\"")
        CALL req.addPart(prt)

        # perform request
        CALL req.setHeader(
            "Content-Disposition", "form-data; name=\"password\"")
        CALL req.doTextRequest(r_param.password)

        # process response
        LET rsp=req.getResponse()
        IF rsp.getStatusCode() != 200 THEN
           RETURN rsp.getStatusCode(), rsp.getStatusDescription()
        END IF
        LET tmp = rsp.getStatusDescription()
        IF (tmp == "OK" || tmp == "no error" ) THEN
            RETURN -1, SFMT("HTTP request status description: %1 ",tmp)
        END IF
        LET result = rsp.getTextResponse()
        
    CATCH
        RETURN -2, SFMT("HTTP request error: STATUS=%1 (%2)",status,sqlca.sqlerrm)
    END TRY

    RETURN 0, result

END FUNCTION
```

## Related links

**Related concepts**  

[com.HttpPart methods](3911-httppart-methods.md "Methods for the com.HttpPart class.")

[com.HttpServiceRequest methods](3807-httpservicerequest-methods.md "Methods of the com.HttpServiceRequest class.")
