---
title: "Get HTTP headers information at WS server side"
source: "fgl-topics/c_gws_access_http_headers_request_and_response.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Get HTTP headers information at WS server side"
type: "concept"
---

# Get HTTP headers information at WS server side

> To work with HTTP headers exchanged between the Genero client and a web service, you must register variables to receive and send them.

In high level (SOAP) Web services, we now give access to the HTTP headers request and
response.

The Web service can get information from the request headers and reply with custom headers and
status.

1. Declare variables to receive or send HTTP headers.
2. Register these variables to the Web service server.

## Declare variables to receive or send HTTP headers

The variable for the request headers:

```
DEFINE http_in RECORD
           verb STRING,
           url STRING,
           headers DYNAMIC ARRAY OF RECORD
                     name STRING,
                     value STRING
                   END RECORD
           END RECORD
```

After the Web service operation has been processed, the variable is set to NULL.

The variable for the response headers:

```
DEFINE http_out RECORD
           code INTEGER,
           desc STRING,
           headers DYNAMIC ARRAY OF RECORD
                     name STRING,
                     value STRING
                   END RECORD
          END RECORD
```

After the Web service operation has been processed, the variable is set to NULL.

While the variables must follow the structure shown, the variable name can be any name you
choose.

The Web service engine headers have precedence. For example, if you set the
"Content-Length" value, the one that is taken into account is the one defined by the Genero
Web Services engine.

## Register the variables to the server

This code example uses two methods, which use the defined variables:

- `com.WebService.registerInputHttpVariable(http_in)` where
  http\_in is the `RECORD` variable for the request headers.
- `com.WebService.registerOutputHttpVariable(http_out)` where
  http\_out is the `RECORD` variable for the response headers

Example

```
FUNCTION CreateService()

  DEFINE serv com.WebService # WebService
  DEFINE op com.WebOperation # Operation of a WebService

  TRY
    #
    # Create a Web Service
    #
    LET serv = com.WebService.CreateWebService("EchoHttpHeadersService",
                                               Namespace)
    #
    # Create Document Style Operations
    #
    # EchoDOCRecord
    LET op = com.WebOperation.CreateDOCStyle("echoDocRecord",
                                             "EchoDOCRecord",
                                             echoRecordDoc_in,
                                             echoRecordDoc_out)
    CALL serv.publishOperation(op,NULL)

    # Register HTTP input
    CALL serv.registerInputHttpVariable(http_in)

    # Register HTTP output
    CALL serv.registerOutputHttpVariable(http_out)

    #
    # Register service
    #
    CALL com.WebServiceEngine.RegisterService(serv)
    DISPLAY "EchoHttpHeadersService Service registered"
    CATCH
       DISPLAY "Unable to create 'EchoHttpHeadersService' Web Service : ",
               status||" ("||sqlca.sqlerrm||")"
       EXIT PROGRAM (-1)
    END TRY

END FUNCTION

FUNCTION echoDocRecord()
  DEFINE ind INTEGER
  DEFINE ok BOOLEAN

  # Check incoming VERB
  IF http_in.verb != "POST" THEN
    LET http_out.code = 400
    LET http_out.desc = "Bad request: method should be POST"
    RETURN
  END IF

  # Check incoming query string
  IF http_in.url.getIndexOF("?MyQuery=OK",1)<=0 THEN
    LET http_out.code = 400
    LET http_out.desc = "Bad request: URL should have MyQuery=OK"
    RETURN
  END IF

  # Check incoming header called MyPersonal
  LET ok = FALSE
  FOR ind = 1 TO http_in.headers.getLength()
    DISPLAY ind||"# ",http_in.headers[ind].name,
            "=",http_in.headers[ind].value
    IF http_in.headers[ind].name == "MyPersonal" THEN
      IF http_in.headers[ind].value == "Header" THEN
        LET ok = TRUE
      END IF
    END IF
  END FOR
  IF NOT ok THEN
    LET http_out.code = 400
    LET http_out.desc =
        "Bad request: expected additional header called MyPersonal"
    RETURN
  END IF

  # assign the output record
  LET echoRecordDoc_out.MyRecord.MyInt =
 echoRecordDoc_in.MyRecord.MyInt
  LET echoRecordDoc_out.MyRecord.MyFloat =
 echoRecordDoc_in.MyRecord.MyFloat

  # Add MyPersonalHeader=MyPersonalValue http headers
  LET http_out.headers[1].name = "MyPersonalHeader"
  LET http_out.headers[1].value = "MyPersonalValue"

END FUNCTION
```

## Related links

**Related concepts**  

[The WebService class](../15_library-reference/3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services.")
