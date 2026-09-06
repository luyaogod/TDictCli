---
title: "com.WebService.registerWSDLHandler"
source: "fgl-topics/c_gws_ComWebService_registerWSDLHandler.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.registerWSDLHandler"
type: "concept"
---

# com.WebService.registerWSDLHandler

> Registers the function to be executed when a WSDL is generated.

## Syntax

```
registerWSDLHandler(
   function STRING )
```

1. function has
   the name of a program function.

## Usage

The `registerWSDLHandler()` method registers a function to be called when the WSDL
of the current Web Service object is generated.

The callback function must be defined with a unique parameter of type [`xml.DomDocument`](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards."), and must return the
reference to this
object:

```
FUNCTION myWSDLHandler( wsdl )
    DEFINE wsdl xml.DomDocument
    ...
    RETURN wsdl
END FUNCTION
```

The callback function typically modifies the content of the WSDL DOM document object passed as
parameter.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Server handlers](../16_web-services/4650-server-handlers.md "Create and register your callback handlers (request, and response) to modify the WSDL.")
