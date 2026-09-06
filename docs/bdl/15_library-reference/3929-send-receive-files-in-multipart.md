---
title: "Examples: Client exchanging files with the server"
source: "fgl-topics/c_gws_ComHTTPPart_example_multipart_exchange.html"
breadcrumb: "Library reference > Extension packages > The com package > HTTP classes > The HttpPart class > Examples (HttpPart) > Send/receive files in multipart"
type: "concept"
---

# Examples: Client exchanging files with the server

> Example of client and server applications using methods of the com.HttpPart class to send and receive files in multipart.

This example consists of two applications: a client and server exchanging an XML document in
multipart with an image as an attachment.

## Client Application

```
IMPORT com
IMPORT xml

CONSTANT SERVER_URL = "http://localhost:8090/MultipartMixed/Sample"                   

MAIN
 
  DEFINE req      com.HttpRequest
  DEFINE resp     com.HttpResponse
  DEFINE doc      xml.DomDocument
  DEFINE p        com.HttpPart
  DEFINE type     STRING
  DEFINE ind      INTEGER
  
  LET req = com.HttpRequest.Create(SERVER_URL)
  CALL req.setMethod("POST")
  CALL req.setHeader("MyClientHeader","Hello")
  TRY
    # Set multipart type
    CALL req.setMultipartType("mixed",NULL,NULL)
    # Add filename as part
    LET p = com.HttpPart.CreateAttachment("my_picture.png")
    # Set attachment Content-Type 
    CALL p.setHeader("Content-Type","image/png")
    # Add part to the request
    CALL req.addPart(p)
    # Perform XML request              
    LET doc = xml.DomDocument.CreateDocument("MyXmlDocument")
    CALL req.doXmlRequest(doc)
    # Check response
    LET resp=req.getResponse()
    IF resp.getStatusCode() != 200 THEN
      DISPLAY  "HTTP Error ("||resp.getStatusCode()||") ",
        resp.getStatusDescription()
      EXIT PROGRAM (-1)
    END IF
    IF resp.getStatusDescription() != "OK" THEN
      DISPLAY  "HTTP Error ("||resp.getStatusCode()||") ",
        resp.getStatusDescription()
      EXIT PROGRAM (-1)
    END IF
    # Check whether multipart response or not
    LET type = resp.getMultipartType()
    IF type IS NULL THEN
      DISPLAY  "Failed : Expected multipart in response"
      EXIT PROGRAM (-1)
    ELSE
      DISPLAY "Response is multipart of :",type
    END IF
    # Check response
    LET doc = resp.getXmlResponse()
    IF doc IS NULL THEN
      DISPLAY  "Expected XML document as response"
      EXIT PROGRAM (-1)
    ELSE
      DISPLAY "Response is : ",doc.saveToString()
    END IF  
    # Process additional parts
    FOR ind = 1 TO resp.getPartCount()
      LET p = resp.getPart(ind)
      IF p.getAttachment() IS NOT NULL THEN
        DISPLAY "Attached file at :",p.getAttachment()
      ELSE
        DISPLAY "Attached part is :",p.getContentAsString()
      END IF
    END FOR
  CATCH
    DISPLAY "unexpected exception :",status," ("||sqlca.sqlerrm||")"          
    EXIT PROGRAM (-1)
  END TRY
END MAIN
```

## Server Application

```
IMPORT com
IMPORT xml

MAIN

  DEFINE req        com.HttpServiceRequest
  DEFINE url        STRING
  DEFINE method     STRING
  DEFINE doc        xml.DomDocument
  DEFINE type       STRING
  DEFINE ind        INTEGER
  DEFINE p          com.HttpPart
  
  CALL com.WebServiceEngine.Start()
  
  LET req = com.WebServiceEngine.GetHTTPServiceRequest(-1)
  LET url = req.getUrl()
  IF url IS NULL  THEN
    DISPLAY "Failed: url should not be null"
    EXIT PROGRAM (-1)    
  END IF
  LET method = req.getMethod()
  IF method IS NULL OR method != "POST" THEN
    DISPLAY "Failed: method should be POST"
    EXIT PROGRAM (-1)    
  END IF
  # Check multipart type
  LET type = req.getRequestMultipartType()
  IF type IS NULL THEN
    DISPLAY "Failed: expected multipart in request"
    EXIT PROGRAM (-1)    
  END IF
  TRY
    LET doc = req.readXMLRequest()
    DISPLAY "Request is :", doc.saveToString()
  CATCH
    DISPLAY "Failed: unexpected error :", status
    EXIT PROGRAM (-1)    
  END TRY
  # Process additional parts
  FOR ind = 1 TO req.getRequestPartCount()
    LET p = req.getRequestPart(ind)
    IF p.getAttachment() IS NOT NULL THEN
      DISPLAY "Attached file at :",p.getAttachment()
    ELSE
      DISPLAY "Attached part is :",p.getContentAsString()
    END IF
  END FOR
  # Set multipart response type
  CALL req.setResponseMultipartType("mixed",NULL,NULL)
  # Add XML Part
  LET p = com.HttpPart.CreateAttachment("my_other_picture.jpg")
  CALL p.setHeader("Content-Type","image/jpg")
  CALL req.addResponsePart(p)
  LET doc = xml.DomDocument.CreateDocument("MyResponse")
  CALL req.sendXmlResponse(200,NULL,doc)
END MAIN
```

## Related links

**Related concepts**  

[com.HttpPart methods](3911-httppart-methods.md "Methods for the com.HttpPart class.")

[com.HttpServiceRequest methods](3807-httpservicerequest-methods.md "Methods of the com.HttpServiceRequest class.")
