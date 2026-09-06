---
title: "Examples using com.HttpRequest methods"
source: "fgl-topics/c_gws_ComHTTPRequest_examples.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpRequest class > Examples (HttpRequest)"
type: "concept"
---

# Examples using com.HttpRequest methods

> These examples use methods of the com.HttpRequest class.

These examples are provided:

- HTTP GET request
- XForms HTTP POST request
- Streaming HTTP PUT request
- Asynchronous HTTP DELETE request
- HTTP request cookie management
- Streaming HTTP JSON request and response

> **Note:**
>
> **Default charset for web services:** The default character encoding for web service
> communication is UTF-8. If the request payload is in JSON format, you do not need to explicitly set
> the charset, as UTF-8 is assumed by default. For other content types (such as XML or plain text),
> you may need to explicitly set the charset to ensure correct encoding. For example: `CALL
> request.setCharset("UTF-8")`

## HTTP GET request

```
IMPORT com
 
MAIN
  DEFINE req com.HttpRequest
  DEFINE resp com.HttpResponse
  TRY
    LET req = com.HttpRequest.Create("http://localhost:8090/MyPage")
    # Set additional HTTP header with name 'MyHeader', and value 'High Priority'
    CALL req.setHeader("MyHeader","High Priority") 
    CALL req.doRequest()
    LET resp = req.getResponse()
    IF resp.getStatusCode() != 200 THEN
      DISPLAY "HTTP Error ("||resp.getStatusCode()||") ",
        resp.getStatusDescription()
    ELSE
      DISPLAY "HTTP Response is : ",resp.getTextResponse()
    END IF
  CATCH
    DISPLAY "ERROR :",status||" ("||sqlca.sqlerrm||")"
  END TRY
END MAIN
```

## XForms HTTP POST request

```
IMPORT com
IMPORT xml
 
MAIN 
  DEFINE req com.HttpRequest
  DEFINE resp com.HttpResponse
  DEFINE doc xml.DomDocument
  TRY 
    LET req = com.HttpRequest.Create("http://localhost:8090/MyProcess")
    CALL req.setMethod("POST") # Perform an HTTP POST method
      # Param1 value is 'hello', Param2 value is 'how are you ?'
    CALL req.doFormEncodedRequest("Param1=hello&Param2=how are you ?",FALSE) 
    LET resp = req.getResponse()
    IF resp.getStatusCode() != 200 THEN
      DISPLAY  "HTTP Error ("||resp.getStatusCode()||") ",
        resp.getStatusDescription()
    ELSE
       # Expect a returned content type of the form */xml 
      LET   doc = resp.getXmlResponse() 
      DISPLAY   "HTTP XML Response is : ",doc.saveToString()
    END IF 
  CATCH
    DISPLAY "ERROR :",status||" ("||sqlca.sqlerrm||")" 
  END TRY
END MAIN
```

## Streaming HTTP PUT request

```
IMPORT com
IMPORT xml
 
MAIN 
  DEFINE req com.HttpRequest
  DEFINE resp com.HttpResponse
  DEFINE writer xml.StaxWriter
  TRY 
    LET req = com.HttpRequest.Create("http://localhost:8090/MyXmlProcess")
    CALL req.setMethod("PUT") # Perform an HTTP PUT method
    CALL req.setHeader("MyHeader","Value of my header") 
      # Retrieve an xml.StaxWriter to start xml streaming 
    LET writer = req.beginXmlRequest() 
    CALL writer.startDocument("utf-8","1.0",true)
    CALL writer.comment("My first XML document sent in streaming with genero") 
    CALL writer.startElement("root")
    CALL writer.attribute("attr1","value1") 
    CALL writer.endElement()
    CALL writer.endDocument()
    CALL req.endXmlRequest(writer) # End streaming request
    LET resp = req.getResponse()
    IF resp.getStatusCode() != 201 OR resp.getStatusCode() != 204 THEN
      DISPLAY   "HTTP Error ("||resp.getStatusCode()||") ",
        resp.getStatusDescription()
    ELSE 
      DISPLAY   "XML document was correctly put on the server"
    END IF 
  CATCH
    DISPLAY "ERROR :",status||" ("||sqlca.sqlerrm||")" 
  END TRY
END MAIN
```

## Asynchronous HTTP DELETE request

```
IMPORT com
 
MAIN 
  DEFINE req com.HttpRequest
  DEFINE resp com.HttpResponse
  DEFINE url STRING
  DEFINE quit CHAR(1)
  DEFINE questionStr STRING
  DEFINE timeout INTEGER
  TRY 
    WHILE TRUE
      PROMPT "Enter http url you want to delete ? " 
         FOR url ATTRIBUTES (CANCEL=FALSE)
      LET req = com.HttpRequest.Create(url)
      CALL req.setMethod("DELETE")
      CALL req.doRequest()
      # Retrieve asynchronous response for the first time
      LET resp = req.getAsyncResponse() 
      CALL Update(resp) RETURNING questionStr,timeout
      WHILE quit IS NULL OR ( quit!="Y" AND quit!="N" )
        PROMPT questionStr FOR CHAR quit 
          ATTRIBUTES (CANCEL=FALSE,ACCEPT=FALSE,SHIFT="up")
          ON IDLE timeout
          IF resp IS NULL THEN # If no response at first try, 
                                # retrieve it again
            LET resp = req.getAsyncResponse() # as we now have time
            CALL Update(resp) RETURNING questionStr,timeout 
          END IF
        END PROMPT 
      END WHILE
      IF quit == "Y" THEN
        EXIT PROGRAM 
      ELSE
        LET quit = NULL
      END IF
    END WHILE
  CATCH
    DISPLAY "ERROR :",status,sqlca.sqlerrm
  END TRY
END MAIN

FUNCTION Update(resp)
  DEFINE resp com.HttpResponse
  DEFINE ret STRING
  IF resp IS NOT NULL THEN
    IF resp.getStatusCode() != 204 THEN
    LET   ret = "HTTP Error ("||resp.getStatusCode()||")
      :"||resp.getStatusDescription()||". Do you want to quit ? "
    ELSE 
      LET   ret = "HTTP Page deleted. Do you want to quit ? "
    END IF 
  RETURN ret, 0
  ELSE 
    LET ret = "Do you want to quit ? "
    RETURN ret, 1
  END IF
END FUNCTION
```

## HTTP request cookie management

```
IMPORT com

MAIN
  DEFINE req com.HttpRequest,
         resp com.HttpResponse,
         ret INTEGER
  CALL com.WebServiceEngine.SetOption("maximumpersistentcookies",80)
  LET req = com.HttpRequest.Create("http://cube.strasbourg.4js:com/hello/world")
  CALL req.setHeader("Cookie","MyPersonaCookie=hello") # Set a specific one
  CALL req.setAutoCookies(TRUE)
  CALL req.doRequest()
  LET resp = req.getResponse()
  IF resp.getStatusCode()!=200 THEN
      DISPLAY "Error : ",resp.getStatusDescription()
      EXIT PROGRAM 1
   ELSE
     LET ret = resp.getTextResponse()
     DISPLAY "OK :",ret
     DISPLAY "Cookie that will be handled automatically are : "
     DISPLAY resp.getHeader("Set-Cookie")
   END IF
END MAIN
```

## Streaming HTTP JSON request and response

```
IMPORT COM
IMPORT JSON
MAIN
    DEFINE req com.HttpRequest
    DEFINE resp com.HttpResponse
    DEFINE resp_body RECORD
        key STRING
    END RECORD
    TRY
        LET req = com.HttpRequest.Create("http://localhost:8090/JSONService/echo")
        CALL req.setMethod("POST")
        CALL req.setHeader("Accept", "application/json")
        CALL req.setHeader("Content-Type", "application/json")
        VAR writer = req.beginJSONRequest()
        CALL writer.startJSON() # Low-level JSON api
        CALL writer.startObject()
        CALL writer.setProperty("key")
        CALL writer.setValue("value")
        CALL writer.endObject()
        CALL writer.endJSON()
        CALL req.endJSONRequest(writer)
        LET resp = req.getResponse()
        VAR retCode = resp.getStatusCode()
        IF retcode == 200 THEN
            VAR reader = resp.beginJSONResponse()
            CALL reader.next()
            CALL json.Serializer.JSONToVariable(reader, resp_body) # High-level JSON API
            CALL resp.endJSONResponse(reader)
        ELSE
            DISPLAY "HTTP Error ("||retCode||")", resp.getStatusDescription()
        END IF
        DISPLAY "key: ",resp_body.key
    CATCH
        DISPLAY "ERROR: " ,status || " ("||sqlca.sqlerrm||")"
    END TRY
END MAIN
```
