---
title: "com.WebService.saveWSDL"
source: "fgl-topics/c_gws_ComWebService_saveWSDL.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.saveWSDL"
type: "concept"
---

# com.WebService.saveWSDL

> Writes to a file the WSDL corresponding to the Web Service object.

## Syntax

```
saveWSDL(
   location STRING )
  RETURNS INTEGER
```

1. location defines
   the URL where the Web Service will be deployed.

## Usage

The `saveWSDL()` method writes the WSDL data corresponding to the Web Service
object.

The URL where the Web Service will be deployed must be specified.

The name of the file will be the name of the Web Service defined by the name
parameter passed to the [`createWebService()`](3758-com-webservice-createwebservice.md "Creates a new object to implement a Web Service.") or [`createStatefulWebService()`](3759-com-webservice-createstatefulwebservice.md "Creates a new object to implement a stateful Web service.") methods.

The method returns 0 if the file was saved, -1 in case of error.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.WebService.setComment](3768-com-webservice-setcomment.md "Defines the comment for the Web Service object.")
