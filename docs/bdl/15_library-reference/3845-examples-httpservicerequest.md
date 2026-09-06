---
title: "Examples using com.HttpServiceRequest methods"
source: "fgl-topics/c_gws_ComHTTPServiceRequest_examples.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The HttpServiceRequest class > Examples (HttpServiceRequest)"
type: "concept"
---

# Examples using com.HttpServiceRequest methods

> These examples use methods of the com.HttpServiceRequest class.

## Server-side example

```
IMPORT COM
IMPORT Security
IMPORT FGL WSHelper

CONSTANT C_5_MINUTES  = INTERVAL(5) MINUTE TO MINUTE

MAIN
  DEFINE  expires  DATETIME DAY TO SECOND  
  DEFINE  req      com.HttpServiceRequest
  DEFINE  query    WSHelper.WSQueryType
  DEFINE  ind      INTEGER
  DEFINE  url, host, port, path, myCookie STRING
  DEFINE  cookies  WSHelper.WSServerCookiesType
  
  CALL com.WebServiceEngine.Start()
  WHILE true
    LET req = com.WebServiceEngine.GetHTTPServiceRequest(-1)
    LET url = req.getUrl()
    DISPLAY "URL is :",url
    LET host = req.getUrlHost()
    DISPLAY "Host is :",host      
    LET port = req.getUrlPort()
    DISPLAY "Port is :",port      
    LET path = req.getUrlPath()
    DISPLAY "Path is :",path      
    CALL req.getUrlQuery(query)
    FOR ind = 1 TO query.getLength()
      DISPLAY "Query "||ind
      DISPLAY "  key =",query[ind].name
      DISPLAY "  val = ",query[ind].value
    END FOR

    # Set the response character set       
    CALL req.setResponseCharset("UTF-8") 

    LET myCookie = req.findRequestCookie("UserIDCookie")
    IF myCookie IS NULL THEN
      DISPLAY "New user"
    ELSE
      DISPLAY "User id is :",myCookie
    END IF      

    # Compute Set-Cookies
    DISPLAY "Now is :",CURRENT
    LET expires = CURRENT + C_5_MINUTES
      
    LET cookies[1].name = "AnotherCookie"
    LET cookies[1].value = security.RandomGenerator.CreateUUIDString()
    LET cookies[1].domain = ".strasbourg.4js.com"
    LET cookies[1].expires = expires
    LET cookies[1].httpOnly = TRUE
    CALL req.setResponseCookies(cookies)
    # Send response with cookie 
    CALL req.sendTextResponse(200,NULL,"IT WORKS.")
  END WHILE   
END MAIN
```
