---
title: "com.WebService.registerOutputRequestHandler"
source: "fgl-topics/c_gws_ComWebService_registerOutputRequestHandler.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebService class > WebService methods > com.WebService.registerOutputRequestHandler"
type: "concept"
---

# com.WebService.registerOutputRequestHandler

> Registers the function to be executed just before the SOAP response is forwarded to the client.

## Syntax

```
registerOutputRequestHandler(
   function STRING )
```

1. function has
   the name of a program function.

## Usage

The `registerOutputRequestHandler()` method registers a function to be called just
after the SOAP engine has processed the request and before the SOAP response is forwarded to the
client.

The output callback function must be defined with a unique parameter of type [`xml.DomDocument`](3959-the-domdocument-class.md "The xml.DomDocument class provides methods to manipulate a data tree, following the DOM standards."), and must return the
reference to this
object:

```
FUNCTION myRequestOutputHandler( out )
    DEFINE out xml.DomDocument
    ...
    RETURN out
END FUNCTION
```

The output callback function typically modifies the content of the SOAP output request DOM
document object passed as parameter.

If `NULL` was returned from the input callback function, the output callback
function will be called with the default SOAP fault node.

When WS-Addressing is enabled, and the server side callback is triggered, the SOAP engine has
already created the SOAP:Header node with the correct WS-Addressing entries. If other headers need
to be added (for WS-Security for instance), it is recommended to add them as child nodes of the
existing SOAP header node, instead of creating a new header node. Otherwise, you may get two SOAP
headers in the same response, which is not allowed.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[Server handlers](../16_web-services/4650-server-handlers.md "Create and register your callback handlers (request, and response) to modify the WSDL.")

[com.WebService.registerInputRequestHandler](3763-com-webservice-registerinputrequesthandler.md "Registers the function to be executed on incoming SOAP requests.")
