---
title: "Option flags"
source: "fgl-topics/c_gws_configuration_API_option_flags.html"
breadcrumb: "Web services > Reference > Configuration API functions - version 1.3 only > Option flags"
type: "concept"
description: "Table 1. Option flags Flag Client or Server Description http_invoketimeout Client Defines the maximum time in seconds a client has to wait before the client connection raises an error because the ..."
---

# Option flags

| Flag | Client or Server | Description |
| --- | --- | --- |
| http\_invoketimeout | Client | Defines the maximum time in seconds a client has to wait before the client connection raises an error because the server is not responding.A value of -1 means that it has to wait until the server responds.The default value is -1. |
| tcp\_connectiontimeout | Client | Defines the maximum time in seconds a client has to wait for the establishment of a TCP connection with a server.A value of -1 means infinite wait.The default value is 30 seconds except for Windows®, where it is 5 seconds. |
| soap\_ignoretimezone | Both | Defines if, during the marshalling and unmarshalling process of a BDL DATETIME data type, the [SOAP](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") engine should ignore the time zone information.A value of zero means false.The default value is false. |
| soap\_usetypedefinition | Both | Defines if the Web Services engine must specify the type of data in all [SOAP](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") requests. This will add an "xsi:type" attribute to each parameter of the request.A value of zero means false.The default value is false. |
| wsdl\_decimalsize | Server | Defines if, during the [WSDL](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") generation, the precision and scale of a DECIMAL variable will be taken into account. See [WSDL generation option notes](5078-wsdl-generation-option-notes.md).A value of zero means false.The default value is true. |
| wsdl\_arraysize | Server | Defines if, during the [WSDL](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") generation, the size of a BDL array will be taken into account. See [WSDL generation option notes](5078-wsdl-generation-option-notes.md).A value of zero means false.The default value is true. |
| wsdl\_stringsize | Server | Defines if, during the [WSDL](4486-introduction-to-web-services.md "Web services are a standard way of communicating between applications over an intranet or Internet.") generation, the size of a CHAR or VARCHAR variable will be taken into account. See [WSDL generation option notes](5078-wsdl-generation-option-notes.md).A value of zero means false.The default value is true. |
