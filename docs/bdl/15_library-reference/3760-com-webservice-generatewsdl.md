---
title: "com.WebService.generateWSDL"
source: "fgl-topics/c_gws_ComWebService_generateWSDL.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.generateWSDL"
type: "concept"
---

# com.WebService.generateWSDL

> Creates an xml.DomDocument object with the WSDL corresponding to the Web Service object.

## Syntax

```
generateWSDL(
   location STRING )
  RETURNS xml.DomDocument
```

1. location defines
   the URL where the Web Service will be deployed.

## Usage

The `generateWSDL()` method creates a new `xml.DomDocument` object
containing the WSDL data of the Web Service object.

The URL where the Web Service will be deployed must be specified.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[com.WebService.setComment](3768-com-webservice-setcomment.md "Defines the comment for the Web Service object.")
