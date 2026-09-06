---
title: "Choosing a web services style"
source: "fgl-topics/c_gws_styles_001.html"
breadcrumb: "Web services > SOAP Web Services > Writing a Web Services server application > Choosing a web services style"
type: "concept"
---

# Choosing a web services style

> Genero Web Services contains style options for creating SOAP Web services. Your choice is dependent on the type of service, (Document or RPC), and the encoding mechanism (literal or encoded) required.

Genero Web Services 2.0 allows you to create SOAP Web services operations in the following
styles:

| Web Services Style | Description |
| --- | --- |
| Remote Procedure Call (RPC) Style Service (RPC/Literal) | Generally used to execute a function, such as a service that returns a stock option. |
| Document Style Service (Doc/Literal) | Generally used for more sophisticated operations that exchange complex data structures, such as a service that sends an invoice to an application, or exchanges a Word document; this is the MS.Net default.Both RPC/Literal and Doc/Literal Styles are Web Services Interoperability (WS-I) organization compliant. |
| RPC Style Service (RPC/Encoded)**Important:**This feature is deprecated, its use is discouraged although not prohibited. | Provided only for backward compatibility with older versions of web services already published.**Important:**This style is deprecated by the WS-I organization, and is not recommended, as most Web Service implementations won't support it in the future. |

The style of service to be created is specified in the Genero application for the Web Service,
using the following methods of the [WebOperation](../15_library-reference/3770-the-weboperation-class.md "The com.WebOperation class provides an interface to create and manage the operations of a Genero Web Service.") class
from the ([Web Services COM
Library](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.")). The parameters are the same for both methods:

1. The name of the BDL function that is executed to process the Web service operation
2. The name you wish to assign to the Web service operation
3. The [input](4642-define-the-input-parameters.md "Define a record for the input message of the Web function.") record defining the input
   parameters of the operation (or NULL if there is none)
4. The [output](4643-define-the-output-parameters.md "Define a record for the output message of the function.") record defining the output
   parameters of the operation (or NULL if there is none)

```
LET op = com.WebOperation.CreateRPCStyle("add","Add",
           add_in,add_out)
LET op = com.WebOperation.CreateDOCStyle("checkInvoice",
           "CheckInvoice",invoice_in,invoice_out)
```

Calling the appropriate function for the desired style is the only
difference in your Genero code that creates the service. The remainder
of the code that describes the service is the same, regardless of
whether you want to create an RPC or Document style of service.

> **Important:**
>
> Do not use the `setInputEncoded()` and
> `setOutputEncoded()` methods of the [WebService](../15_library-reference/3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services.") class from the [Web
> Services COM Library](../15_library-reference/3752-the-com-package.md "The Genero Web Services com package provides classes and methods that allow you to perform tasks associated with creating Services and Clients, and managing the services.") (`com`), as they apply only to RPC/Encoded Style, which
> is not recommended.

If you add headers to your RPC Style service, choose the Literal serialization mechanism
by setting the `encoded` parameter of the `createHeader()`
method to
FALSE:

```
CALL serv.createHeader(var,FALSE)
```

GWS release 2.0 allows you to create RPC Style and Document Style operations in the same
Web service. However, we do not recommend this, as it is not WS-I compliant.
