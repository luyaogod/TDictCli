---
title: "Step 5: Process the HTTP response"
source: "fgl-topics/c_gws_rest_client_response.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web services client application > Step 5: Process the HTTP response"
type: "concept"
---

# Step 5: Process the HTTP response

> In the final step you process the HTTP response; check for errors, and process the data.

At this stage a response to our request from the Web service has been received and is stored in
the variable, `resp`. We defined the variable `resp` as an object of
the `com.HttpResponse` class earlier, see [Step 2: Import extension packages (com, xml, util)](4857-step-2-import-extension-packages-com-xml-util.md "The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages."). Now it is time to process the
response.

```
LET resp = req.getResponse()
```

Start checking for errors. First get the HTTP status code returned in the response by calling the
`com.HttpResponse` object's getStatusCode function referenced by
the `resp` variable. Assign the status code to the "status" field of the
`info` record ( for details about this record see [Step 3: Define the records](4858-step-3-define-the-records.md "In this step you define the records you need for the HTTP Request and Response and the processing of the data.")).

```
LET info.status = resp.getStatusCode()
```

The standard HTTP code of 200
indicates a successful operation, so we check for this before looking at the
data.

```
IF info.status = 200 THEN
```

If the HTTP response was good, we save the value for Content-type response header to the
"resptype" field of the `info` record. The Content-type header defines the type
associated with the message body's byte sequence. REST APIs commonly use "application/json"or
"application/xml" to reveal the format of the message body.

```
LET info.resptype = resp.getHeader("Content-Type")
```

## Process for XML

```
IF info.resptype.getIndexOf("/xml",1) THEN
    LET doc = resp.getXmlResponse()                  
    LET node = doc.getDocumentElement()
    CALL xml.Serializer.DomToVariable(node, add_out) 
    LET info.response = node.toString()
ELSE # JSON response processing
     [...]
END IF
```

If
the Content-type indicates XML format, we store the HTTP response as data in the
`xml.DomDocument` object referenced by the `doc` variable using the
`getXmlResponse` function. We then get a reference to the root node of the
`doc` variable object by instantiating the `node` variable defined as
an object of the `xml.DomNode` class using the `getDocumentElement`
function.

Recall the `doc` and `node` variables were defined earlier, see
[Step 2: Import extension packages (com, xml, util)](4857-step-2-import-extension-packages-com-xml-util.md "The functions you need to create a REST Web Service client application are contained in the classes that make up the com package of the Genero Web Services (GWS). Use the IMPORT statement to include the required packages.").

To serialize the XML data into the `add_out` variable, we pass a reference to the
`node` object variable in a call to the `xml.Serializer.DomToVariable`
function. To convert the data referenced in the `node` object variable to string, we
use the `xml.DomNode.toString` function to store it in the "response" field of the
`info` record.

## Process for JSON

If the Content-type does not indicate XML format, then we process the response for JSON, as we
anticipate this as the only other option REST APIs use to format the message body.

We store the HTTP response as a string in the "response" field of the `info`
record with a reference to the `resp` variable, object of the
`com.HttpResponse` class, using the getTextResponse function. Then
we pass a reference to the `info` record in a call to the
`util.JSON.parse()` function to convert the HTTP response to JSON format into the
`add_out`
record.

```
IF info.resptype.getIndexOf("/xml",1) THEN
   # XML  response processing
   [...]
ELSE # JSON response processing
     LET info.response = resp.getTextResponse()  
     CALL util.JSON.parse(info.response, add_out)
END IF
```

## Outputting to the display

Finally, we output the HTTP response to the display depending on the status
code.

```
IF add_out.status.code = 0 THEN
    DISPLAY add_in.a, "+", add_in.b, "=", add_out.r # Display the response, status and result
ELSE
    DISPLAY "[", add_out.status.code, "] ", add_out.r
END IF
```

## Handling Errors

If the HTTP request failed, we handle the error by formatting the HTTP response status code and
description into the "response" field of the `info` record. Other program errors are
caught in the CATCH block and stored in "result" fields of the `info` record.

```
TRY
[...]
    IF info.status = 200 THEN 
        [...]
    ELSE 
        LET info.response = SFMT("[%1] %2",resp.getStatusCode(), resp.getStatusDescription())
    END IF
CATCH
    # Catch other runtime execution errors from the sqlca diagnostic record
     LET info.result.code = status
     LET info.result.desc = sqlca.sqlerrm
 END TRY
```

This completes the tutorial for the calculator demo client application. We have seen the basic
steps that a RESTful client application must perform in order to carry out a simple request to a Web
service. To review the complete source code, see [The RESTful calculator demo source](4876-the-restful-calculator-demo-source.md "This calculator demo provides a sample RESTful Web services, demonstrating how to create a server and client application using low-level APIs.").

Each Web services client application will be different in the type of processing that is done.
For more examples of Web service applications, see the code samples provided with the package in
demo/WebServices.
