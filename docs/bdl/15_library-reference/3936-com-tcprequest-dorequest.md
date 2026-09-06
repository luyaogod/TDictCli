---
title: "com.TcpRequest.doRequest"
source: "fgl-topics/c_gws_ComTCPRequest_doRequest.html"
breadcrumb: "Library reference > Extension packages > The com package > TCP classes > The TcpRequest class > TcpRequest methods > com.TcpRequest.doRequest"
type: "concept"
---

# com.TcpRequest.doRequest

> Performs a TCP request.

## Syntax

```
doRequest()
```

## Usage

The `doRequest()` method performs the TCP request.

The connection is shutdown for writing, to confirm that no data will be sent.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

The `int_flag` variable is checked during GWS API call to
handle program interruptions, for more details, see [Interruption handling in GWS calls (int\_flag)](../16_web-services/5065-interruption-handling-in-gws-calls-int-flag.md "Genero Web Services (GWS) tests int_flag to check if an application has been interrupted.")

## Example

```
IMPORT com
IMPORT XML

MAIN
  DEFINE url STRING
  LET url = "tcp://localhost:4242"
  CALL an_example(url)
END MAIN

FUNCTION an_example(url)
DEFINE url STRING
DEFINE req com.TcpRequest
DEFINE resp com.TcpResponse
DEFINE ret xml.DomDocument

TRY
  LET  req = com.TcpRequest.create(url)
  CALL req.doRequest()
  LET  resp = req.getResponse()
  LET  ret = resp.getXmlResponse()  
CATCH
  DISPLAY "ERROR : ", status, sqlca.sqlerrm
  EXIT PROGRAM(-1)   
END TRY

END FUNCTION
```
