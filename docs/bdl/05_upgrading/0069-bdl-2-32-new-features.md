---
title: "BDL 2.32 new features"
source: "fgl-topics/fgl_whatsnew_232.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 2.32 new features"
type: "topic"
---

# BDL 2.32 new features

> Features added in 2.32 releases of the Genero Business Development Language.

> **Important:**
>
> This version of Genero BDL is desupported, use a more recent version of the product.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 2.32 upgrade guide](0256-bdl-2-32-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 2.32.").

| Overview | Reference |
| --- | --- |
| The COM library enables to intercept high-level web services operation on server side. You can now define three BDL functions via methods of the web service class.They will be executed at different steps of a web service request processing in order to modify the SOAP request, response or the generated WSDL document before or after the SOAP engine has processed it. This helps handle **WS-\*** specifications not supported in the web service API.Method `registerWSDLHandler()`Method `registerInputRequestHandler()`Method `registerOutputRequestHandler()`All three kinds of BDL callback functions must conform to the following prototype:FUNCTION CallbackHandler( doc xml.DomDocument ) RETURNING xml.DomDocument | See [The WebService class](../15_library-reference/3754-the-webservice-class.md "The com.WebService class provides an interface to create and manage Genero Web Services."). |
