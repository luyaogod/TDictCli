---
title: "Server handlers"
source: "fgl-topics/c_gws_handlers_server_handlers.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > WS server stubs and handlers > Server handlers"
type: "concept"
---

# Server handlers

> Create and register your callback handlers (request, and response) to modify the WSDL.

The [COM](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") library allows you to intercept high-level
Web services operation on the server side. You can now define three BDL functions via the following
methods of the Web service class. They will be executed at different steps of a web service request
processing in order to modify the SOAP request, response or the generated WSDL document before or
after the SOAP engine has processed it. This helps handle **WS-\*** specifications not supported
in the web service API.

- Method `registerWSDLHandler()`
- Method `registerInputRequestHandler()`
- Method `registerOutputRequestHandler()`

All three kinds of BDL callback functions must conform to this
prototype:

```
FUNCTION CallbackHandler( doc xml.DomDocument )
   RETURNING xml.DomDocument
```

## Example 1: Modify the generation of a WSDL

Register your handler
with:

```
CALL serv.registerWsdlHandler("WSDLHandler")
```

where `serv` is of class `com.WebService` and
`WSDLHandler` is the following
function:

```
IMPORT XML

FUNCTION WSDLHandler(wsdl)
  DEFINE wsdl Xml.DomDocument
  DEFINE node Xml.DomNode
  DEFINE list Xml.DomNodeList
  DEFINE ind  INTEGER
  DEFINE name STRING
  # Add a comment
  LET node = wsdl.createComment(
              "First modified WSDL via a BDL callback function")
  CALL wsdl.prependDocumentNode(node)
  # Rename input and output parameter in UPPERCASE
  LET list = wsdl.selectByXPath(
              "//wsdl:definitions/wsdl:types/xsd:schema/
xsd:complexType/xsd:sequence/xsd:element/xsd:complexType/
xsd:sequence/xsd:element",NULL) 
-- first input parameter for selectByXPath above
-- one string, no spaces!
  FOR ind=1 TO list.getCount()
    LET node = list.getItem(ind)
    LET name = node.getAttribute("name")
    LET name = name.toUpperCase()
    CALL node.setAttribute("name",name)
  END FOR
  RETURN wsdl
END FUNCTION
```

If NULL is returned from the callback function,
an HTTP error will be sent and the [ProcessServices()](../15_library-reference/3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.") returns
error code -20.

## Example 2: Change the SOAP incoming request

Register your handler
with:

```
CALL serv.registerInputRequestHandler("InputRequestHandler")
```

where `serv` is of class `com.WebService` and
`InputRequestHandler` is this
function:

```
IMPORT XML

FUNCTION InputRequestHandler(in)
  DEFINE in Xml.DomDocument
  DEFINE ind  INTEGER
  DEFINE node Xml.DomNode
  DEFINE copy Xml.DomNode
  DEFINE tmp  Xml.DomNode
  DEFINE parent Xml.DomNode
  DEFINE name STRING
  DEFINE list Xml.DomNodeList
  # Change input parameter below myrecord in lower case 
  # to follow high-level web service
  LET list = in.SelectByXPath(
    "//SOAP:Envelope/SOAP:Body/fjs:EchoDOCRecordRequest/fjs:myrecord/*",
    "SOAP","http://schemas.xmlsoap.org/soap/envelope/",
    "fjs","http://www.mycompany.com/webservices")
  FOR ind = 1 TO list.getCount()
    LET node = list.getItem(ind)
    LET parent = node.getParentNode()
    LET name = node.getLocalName()
    LET copy = in.createElementNS(node.getPrefix(),
      name.toLowerCase(),node.getNamespaceURI())
    LET tmp = node.getFirstChild()
    LET tmp = tmp.clone(true)
    CALL copy.appendChild(tmp)
    CALL parent.replaceChild(copy,node)
  END FOR
  RETURN in
END FUNCTION
```

If NULL is returned from the callback function, a SOAP fault will be sent (but can be changed
from the output handler) and the [ProcessServices()](../15_library-reference/3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.") returns error code -18.

## Example 3: Modify the SOAP outgoing request

Register your handler
with:

```
CALL serv.registerOutputRequestHandler("OutputRequestHandler")
```

where serv is of class `com.WebService` and `OutputRequestHandler`
is this function:

```
IMPORT XML

FUNCTION OutputRequestHandler(out)
  DEFINE out Xml.DomDocument
  DEFINE ind  INTEGER
  DEFINE node Xml.DomNode
  DEFINE copy Xml.DomNode
  DEFINE tmp  Xml.DomNode
  DEFINE parent Xml.DomNode
  DEFINE name STRING
  DEFINE list Xml.DomNodeList
  # Change output parameter below myrecord in uppercase 
  # before sending back to the client
  LET list = out.SelectByXPath(
    "//SOAP:Envelope/SOAP:Body/fjs:EchoDOCRecordResponse/fjs:myrecord/*",
    "SOAP","http://schemas.xmlsoap.org/soap/envelope/",
    "fjs","http://www.mycompany.com/webservices")
  FOR ind = 1 TO list.getCount()
    LET node = list.getItem(ind)
    LET parent = node.getParentNode()
    LET name = node.getLocalName()
    LET copy = out.createElementNS(node.getPrefix(),name.toUpperCase(),
      node.getNamespaceURI())
    LET tmp = node.getFirstChild()
    LET tmp = tmp.clone(true)
    CALL copy.appendChild(tmp)
    CALL parent.replaceChild(copy,node)
  END FOR
  RETURN out
END FUNCTION
```

If NULL is return from the callback function, a SOAP fault will be sent and the [ProcessServices()](../15_library-reference/3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services.") returns
error code -19.
