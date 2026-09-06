---
title: "Step 3: Parse the request"
source: "fgl-topics/c_gws_rest_server_parse_request.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (low-level APIs) > Writing a Web Services server application > Step 3: Parse the request"
type: "concept"
---

# Step 3: Parse the request

> Using methods of the com.HttpServiceRequest class (among others), parse the details of the request.

An HTTP service request has been received and stored in an instance of a `com.HttpServiceRequest` object.

To use the details of the request, the next step is to define a record to hold the pieces of
information that come in with a request and to parse the request into the record. The
demo creates a user-defined type (`reqInfoTyp`) to reference when
defining the record.

This topic details how the calculator demo server application parses the instance of the
`HttpServiceRequest` object. While your specific HTTP service
requests may differ in the details, the parsing needs will be similar.

```
TYPE reqInfoTyp RECORD
    method      STRING,
    ctype       STRING,     # check the Content-Type
    informat    STRING,     # short word for Content Type 
    caccept     STRING,     # check which format the client accepts
    outformat   STRING,     # short word for Accept
    path        STRING,
    query       STRING,     # the query string
    items       DYNAMIC ARRAY OF RECORD
        name STRING,
        value STRING
    END RECORD
END RECORD
```

Define a variable based on the type.

```
DEFINE m_reqInfo reqInfoTyp
```

Call a function that takes the variable referencing the instance of the
`com.HttpServiceRequest` object as input. The purpose of the function will be to
parse the HTTP service request into its components.

```
CALL getReqInfo(req)
```

## The getReqInfo function

This function uses both the HttpServiceRequest and
StringTokenizer classes and methods to parse the request into the
`reqInfoTyp` record.

```
FUNCTION getReqInfo(req)
```

Define the variables needed. The variable `req` references the instance of the
HTTP service request object. The variables `str`, `val`,
`token`, and `i` are used when parsing out details with the
StringTokenizer classes and
methods.

```
DEFINE req com.HttpServiceRequest
DEFINE str, val base.StringTokenizer
DEFINE token STRING
DEFINE i INT
```

For
the remainder of this topic, you are retrieving data from the instance of the HTTP service request
object (referenced by the variable `req`) and populating the record variable designed
to hold this data (`m_reqInfo`).

Initialize `m_reqInfo` to `NULL`. This ensures that there are no
pre-existing values in the record variable.

The variable `m_reqInfo` record was defined at the top of the module,
and therefore can be used in this function.

```
INITIALIZE m_reqInfo TO NULL
```

Retrieve the value of the Content-Type request header. The Content-Type request header defines
the format of the incoming message body. REST APIs commonly use "application/json"or
"application/xml" to reveal the format of the message body. The server application anticipates the
Content-Type to be either JSON or XML. The custom getHeaderByName
function retrieves the value of the Content-Type header, and based on the return value, populates
the input format variable (`m_reqInfo.informat`) with either "XML" or
"JSON".

```
LET m_reqInfo.ctype = getHeaderByName(req,"Content-Type")
IF m_reqInfo.ctype.getIndexOf("/xml",1) THEN
    LET m_reqInfo.informat = "XML"
ELSE 
    LET m_reqInfo.informat = "JSON"
END IF
```

Retrieve the value for Accept request header. The Accept request header is where the client
application conveys its response preference. The server application expects the Accept request
header to be either JSON or XML. The custom `getHeaderByName` function retrieves the value of the Accept header, and
based on the returned value, populates the output format variable
(`m_reqInfo.outformat`) with either "XML" or
"JSON".

```
LET m_reqInfo.caccept = getHeaderByName(req,"Accept")
IF m_reqInfo.caccept.getIndexOf("/xml",1) THEN
    LET m_reqInfo.outformat = "XML"
ELSE 
    LET m_reqInfo.outformat = "JSON"
END IF
```

Parse out the HTTP method (verb) of the HTTP service request with the function
com.HttpServiceRequest.getMethod(), and populate the method variable
(`m_reqInfo.method`) with the
result.

```
LET m_reqInfo.method = req.getMethod()
```

While possible values could
include GET, POST, PUT, HEAD, and DELETE, our demo server application expects GET, and the
processing part of the server application code does not test or code for any verbs aside from
GET.

Parse out the path and the query. Retrieve the URL resource using the function
com.HttpServiceRequest.getUrl. To parse out the path and the query from the URL,
call the custom `parseUrl` function. The path is placed in the path variable
(`m_reqInfo.path)` and the query string is placed in the query variable
(`m_reqInfo.query)`.

```
CALL parseUrl(req.getUrl()) RETURNING m_reqInfo.path, m_reqInfo.query
```

Take the query variable (`m_reqInfo.query`) and parse out the value pairs using
the StringTokenizer class and methods. The value pairs of variable name and value
are stored in the `m_reqInfo.items` dynamic
array.

```
LET str = base.StringTokenizer.create(m_reqInfo.query,"&")
LET i=1
CALL m_reqInfo.items.clear()
WHILE str.hasMoreTokens()
    LET token = str.nextToken()
    LET val = base.StringTokenizer.create(token,"=")
    IF val.hasMoreTokens() THEN LET m_reqInfo.items[i].name = val.nextToken() END IF
    IF val.hasMoreTokens() THEN LET m_reqInfo.items[i].value = val.nextToken() END IF
    LET i=i+1
END WHILE
```

By the end of the custom `getReqInfo()` function, the HTTP service request object
is parsed into the variables that comprise a `reqInfoType` record. The application
can now access the values it needs to process the request.

## getHeaderByName function

The purpose of this function is to take as input two things:

- The HTTP service request object in an `com.HttpServiceRequest` variable.
- The name of the header type whose value you want to return.

In the demo application, this function is used to parse out the value for the Content-Type header
request and the Accept header request.

To do this, the function uses the API methods getRequestHeaderCount,
getRequestHeaderName, and getRequestHeaderValue against the
`com.HttpServiceRequest` object. The function first counts the number of header
requests, and then cycles through them until it finds the specific header request that is being
asked for. When the match happens, the value for that header is returned; if no match is found, then
NULL is
returned.

```
FUNCTION getHeaderByName(areq,hname)
    DEFINE areq com.HttpServiceRequest
    DEFINE hname STRING

    DEFINE aname STRING
    DEFINE iname STRING
    DEFINE i INT
    DEFINE n INT

    LET aname = hname.toLowerCase()
    LET n = areq.getRequestHeaderCount()
    FOR i=1 TO n
        LET iname = areq.getRequestHeaderName(i)
        IF aname.equals(iname.toLowerCase()) THEN
            RETURN areq.getRequestHeaderValue(i)
        END IF
    END FOR

    RETURN NULL
END FUNCTION
```

## parseUrl fuction

For this function, the URL is passed in. This function divides the URL into two parts, the path
(up to and including the question mark) and the query. The path and the query are then returned to
the calling
function.

```
FUNCTION parseUrl(url)
    DEFINE url STRING
    DEFINE i INT

    LET i = url.getIndexOf("?",1)
    IF i = 0 THEN 
        RETURN url, NULL
    ELSE
        RETURN url.subString(1,i), url.subString(i+1,url.getLength())
    END IF
END FUNCTION
```

For example, the URL being passed in by the demo client application is
"http://localhost:8090/add?a=1&b=2". The function then returns the path
("http://localhost:8090/add?") and the query
("a=1&b=2").

For more information see [com.WebServiceEngine methods](../15_library-reference/3786-webserviceengine-methods.md "Methods for the com.WebServiceEngine class.")

In the next step we process the incoming request, [Step 4: Process the request](4873-step-4-process-the-request.md "Having parsed the HttpServiceRequest object into its parts, you can now process the request.")
