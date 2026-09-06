---
title: "The WebOperation class"
source: "fgl-topics/c_gws_ComWebOperation.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class"
type: "concept"
---

# The WebOperation class

> The com.WebOperation class provides an interface to create and manage the operations of a Genero Web Service.

This class is provided in the `com` [C-Extension](../14_extending-the-language/2703-c-extensions.md "With C-Extensions, you can bind your own C libraries in the runtime system, to call C function from the application code.") library; To use this class, import the
`com` package with:

```
IMPORT com
```

The Web Operation can be created as RPC Style or Document Style.
Both RPC/Literal and Doc/Literal Styles are WS-I compliant (standards
set by the Web Services Interoperability organization).

RPC Style Service (RPC/Literal) is generally used to execute a function, such as a service that
returns a stock option. Document Style Service (Doc/Literal) is generally used for more
sophisticated operations that exchange complex data structures, such as a service that sends an
invoice to an application, or exchanges a Word document; this is the MS.Net default. The input or
output `RECORD` cannot have `XMLNamespace` attributes set on their
members.

Calling the appropriate function to create the desired style is the only difference in your
Genero code that creates the service. The remainder of the code that describes the service is the
same, regardless of whether you want to create an RPC or Document style of service.

Do not use the `setInputEncoded()` and `setOutputEncoded()`
methods, as they will specify the RPC/Encoded Style, which is not recommended (see [Choosing a Web Service Style](../16_web-services/4673-choosing-a-web-services-style.md "Genero Web Services contains style options for creating SOAP Web services. Your choice is dependent on the type of service, (Document or RPC), and the encoding mechanism (literal or encoded) required.")).

Since release 2.0 GWS allows you to create RPC Style and Document Style operations in the same
Web service. However, we do not recommend this, as it is not WS-I compliant.

## Child topics

- [com.WebOperation methods](3771-weboperation-methods.md): Methods for the com.WebOperation class.
