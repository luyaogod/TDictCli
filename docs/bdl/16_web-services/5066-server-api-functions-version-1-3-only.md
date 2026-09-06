---
title: "Server API functions - version 1.3 only"
source: "fgl-topics/c_gws_server_API_functions.html"
breadcrumb: "Web services > Reference > Server API functions - version 1.3 only"
type: "concept"
---

# Server API functions - version 1.3 only

> Server API functions can create a Web Services server in Genero BDL.

> **Warning:**
>
> These functions are valid for backward compatibility, but they are not the preferred way to
> handle Genero Web Services. See [the com package](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") for
> the preferred classes and methods for handling Web services.

| Name | Description |
| --- | --- |
| FUNCTION fgl_ws_server_setNamespace( namespace VARCHAR ) | Defines the namespace of the service on the Web and must be called first, before all other functions of the API. |
| FUNCTION fgl_ws_server_start( tcpPort VARCHAR ) | Creates and starts the Web services server. |
| FUNCTION fgl_ws_server_publishFunction( operationName VARCHAR, inputNamespace VARCHAR, inputRecordName VARCHAR, outputNamespace VARCHAR, outputRecord VARCHAR, functionName VARCHAR) | Publishes the given BDL function as a Web-Function on the Web. |
| FUNCTION fgl_ws_server_generateWSDL( serviceName VARCHAR, serviceLocation VARCHAR, fileName VARCHAR ) RETURNS INTEGER | Generates the WSDL file based on the BDL-server program. |
| FUNCTION fgl_ws_server_process( timeout INTEGER) RETURNS INTEGER | Waits for an incoming SOAP request for a given time (in seconds) and then processes the request, or returns, if there has been no request during the given time. |
| FUNCTION fgl_ws_server_setFault( faultMessage VARCHAR ) | Return a SOAP fault string to the client at the end of the function's execution. |
| FUNCTION fgl_ws_server_getFault() RETURNS STRING | Retrieves the last fault string the user has set in a Web-Function, or an empty string if there is none. |

## Child topics

- [fgl_ws_server_setNamespace() (version 1.3)](5067-fgl-ws-server-setnamespace-version-1-3.md): Defines the namespace of the service on the Web and must be called first, before all other functions of the API.
- [fgl_ws_server_start() (version 1.3)](5068-fgl-ws-server-start-version-1-3.md): Creates and starts the Web services server.
- [fgl_ws_server_publishFunction() (version 1.3)](5069-fgl-ws-server-publishfunction-version-1-3.md): Publishes the given BDL function as a Web-Function on the Web.
- [fgl_ws_server_generateWSDL() (version 1.3)](5070-fgl-ws-server-generatewsdl-version-1-3.md): Generates the WSDL file based on the BDL-server program.
- [fgl_ws_server_process() (version 1.3)](5071-fgl-ws-server-process-version-1-3.md): Waits for an incoming SOAP request for a given time (in seconds) and then processes the request, or returns, if there has been no request during the given time.
- [fgl_ws_server_setFault() (version 1.3)](5072-fgl-ws-server-setfault-version-1-3.md): Return a SOAP fault string to the client at the end of the function's execution.
- [fgl_ws_server_getFault() (version 1.3)](5073-fgl-ws-server-getfault-version-1-3.md): Retrieves the last fault string the user has set in a Web-Function, or an empty string if there is none.
