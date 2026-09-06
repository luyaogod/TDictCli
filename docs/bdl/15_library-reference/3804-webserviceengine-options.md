---
title: "WebServiceEngine options"
source: "fgl-topics/c_gws_ComWebServiceEngine_option_list.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine options"
type: "concept"
description: "The WebServiceEngine options control runtime behavior for SOAP and REST services, including timeouts, cookie handling, response limits, and data formatting. Options are set at runtime using ..."
---

# WebServiceEngine options

The WebServiceEngine options control runtime behavior for SOAP and REST services, including
timeouts, cookie handling, response limits, and data formatting. Options are set at runtime using
[com.WebServiceEngine.SetOption](3799-com-webserviceengine-setoption.md "Sets an option for the Web Service engine."). The tables below group the options by the
protocol they apply to.

| Flag | Client or Server | Description |
| --- | --- | --- |
| `SoapModuleURI` | Both | Defines the SOAP role of a Genero application with a URI to identify it along a SOAP message path.The default value is NULL. |
| `wsdl_arraysize` | Server | Defines whether the size of a BDL array will be taken into account during the [WSDL](../16_web-services/4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") generation. See [WSDL generation options notes](3803-wsdl-generation-options-notes.md "If you are planning to work with WSDL generation, you are advised to review these notes.").A value of zero means FALSE.The default is TRUE. |
| `wsdl_decimalsize` | Server | Defines whether the precision and scale of a DECIMAL variable will be taken into account during the [WSDL](../16_web-services/4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") generation. See [WSDL generation options notes](3803-wsdl-generation-options-notes.md "If you are planning to work with WSDL generation, you are advised to review these notes.").A value of zero means FALSE.The default is TRUE. |
| `wsdl_stringsize` | Server | Defines whether the size of a CHAR or VARCHAR variable will be taken into account during the [WSDL](../16_web-services/4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") generation. See [WSDL generation options notes](3803-wsdl-generation-options-notes.md "If you are planning to work with WSDL generation, you are advised to review these notes.").A value of zero means FALSE.The default is TRUE. |

| Flag | Client or Server | Description |
| --- | --- | --- |
| `server_restdefaultformat` | Server | Defines the default MIME type format used for REST operations at runtime. By default, REST supports XML and JSON. If the MIME type is not set by the [WSMedia](../16_web-services/4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.") attribute on the record definition, you can specify the format at runtime by calling `com.WebServiceEngine.SetOption("server_restdefaultformat","xml")`.Valid values are:`xml``json``both`You can retrieve the current value with `com.WebServiceEngine.GetOption`. |

| Flag | Client or Server | Description |
| --- | --- | --- |
| `autocookiesmanagement` | Client | If set to `TRUE`, it activates the automatic cookies management for any HttpRequest. It is the same as calling `req.setAutoCookies(TRUE)` for all HttpRequest to a single instance of fglrun.By setting to `TRUE`, the program does not need to call `req.setAutoCookies(true)` in its code. All cookies received from a server are automatically sent back according to their definition (path, expiration date, domain).Default value is `FALSE`. |
| `connectiontimeout` | Client | Defines the default maximum time in seconds a client, an HttpRequest and a TcpRequest have to wait for the establishment of a connection with a server.A value of -1 means infinite wait.The default is 30 seconds for non-Windows, 5 seconds for Windows®. |
| `http_invoketimeout`**Important:**This feature is deprecated, its use is discouraged although not prohibited. | Client | Defines the default maximum time in seconds a client has to wait before the client connection raises an error because the server is not responding.A value of -1 means it waits indefinitely.The default is -1.**Important:**Deprecated — use `readwritetimeout` instead. |
| `maximumpersistentcookies` | Client | Specifies the maximum number of persistent cookies that can be handled by a single fglrun process. If the limit is reached, the older cookies are deleted.Default value is 50. |
| `maximumresponselength` | Both | Defines the maximum authorized size in KBytes for a client, server, HTTP or TCP response before the engine stops processing.A value of -1 means no limit.The default is -1. |
| `readwritetimeout` | Client | Defines the default maximum time in seconds a client, an HTTP request/response or a TCP request/response waits before raising an error when the server does not accept or return data.A value of -1 means infinite wait.The default is -1. |
| `server_readwritetimeout` | Server | Defines how long a socket read or write operation can wait before raising an error. The default value is five seconds.Before this option existed, the default value was -1 (infinite) and was configurable using the accept timeout parameter of [`ProcessServices()`](3791-com-webserviceengine-processservices.md "Wait for an HTTP input request to process an operation of one of the registered SOAP or REST Web Services."). |
| `tcp_connectiontimeout`**Important:**This feature is deprecated, its use is discouraged although not prohibited. | Client | Defines the maximum time in seconds a client waits for a TCP connection with a server.A value of -1 means infinite wait.The default is 30 seconds for non-Windows, 5 seconds for Windows.**Important:**Deprecated — use `connectiontimeout` instead. |

## Related links

**Related concepts**  

[com.WebServiceEngine.GetOption](3789-com-webserviceengine-getoption.md "Returns the value of a Web Service engine option.")

[com.WebServiceEngine.SetOption](3799-com-webserviceengine-setoption.md "Sets an option for the Web Service engine.")
