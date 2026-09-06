---
title: "com.WebService.registerInputRequestHandler"
source: "fgl-topics/c_gws_ComWebService_registerInputRequestHandler.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.registerInputRequestHandler"
type: "concept"
---

# com.WebService.registerInputRequestHandler

> Registers the function to be executed on incoming SOAP requests.

## Syntax

```
registerInputRequestHandler(
   function STRING )
```

1. function has
   the name of a program function.

## Usage

The `registerInputRequestHandler()` method registers a function to be called when
an incoming SOAP request is received and before the SOAP engine has processed it.

The callback function must be defined with a unique parameter of type [`xml.DomDocument`](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards."), and must return the
reference to this object, or
`NULL`:

```
FUNCTION myRequestInputHandler( in )
    DEFINE in xml.DomDocument
    ...
    RETURN in
END FUNCTION
```

The input callback function typically modifies the content of the SOAP input request DOM document
object passed as parameter.

When returning `NULL` from the input callback function, the output callback
function will be called with the default SOAP fault node, which can then be modified.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Server handlers](../16_web-services/4650-server-handlers.md "Create and register your callback handlers (request, and response) to modify the WSDL.")

[com.WebService.registerOutputRequestHandler](3765-com-webservice-registeroutputrequesthandler.md "Registers the function to be executed just before the SOAP response is forwarded to the client.")
