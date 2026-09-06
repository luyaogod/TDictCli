---
title: "Step 6: Generate .4gl stub to access the .NET library"
source: "fgl-topics/c_gws_NET_APIs_011.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Calling .NET from Genero > Step 6: Generate .4gl stub to access the .NET library"
type: "concept"
description: "Use the fglwsdl tool to generate the client stub to access the BarcodeService, as follows: $ fglwsdl http://localhost/BarCodeService.asmx?WSDL This will create two .4gl files, which must be compiled ..."
---

# Step 6: Generate .4gl stub to access the .NET library

Use the fglwsdl tool to generate the client stub to access the BarcodeService,
as follows:

```
$ fglwsdl http://localhost/BarCodeService.asmx?WSDL
```

This will create two .4gl files, which must be compiled and linked into your
BDL application in order to call the .NET barcode library functions. These files contain the BDL
interface to access the .NET library where you will find the function `buildImage`,
defined in BDL.

In the final step we modify the existing application, [Step 7: Modify your BDL application](4903-step-7-modify-your-bdl-application.md).

## Related links

**Related concepts**  

[WS client stubs and handlers](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.")
