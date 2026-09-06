---
title: "Step 4: Generate BDL stub to access the Java library"
source: "fgl-topics/c_gws_java_APIs_010.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Calling Java from Genero > Step 4: Generate BDL stub to access the Java library"
type: "concept"
description: "Use the fglwsdl tool to generate the client stub to access the BarcodeService: $ fglwsdl http://localhost:9090/BarcodeService?WSDL This will create two .4gl files that must be compiled and linked into ..."
---

# Step 4: Generate BDL stub to access the Java library

Use the fglwsdl tool to generate the client stub to access the
BarcodeService:

```
$ fglwsdl http://localhost:9090/BarcodeService?WSDL
```

This will create two .4gl files that must be compiled and linked into your
BDL application in order to call the Java barcode library functions. These files contain the BDL
interface to access the Java library where you will find the two functions,
`readImage` and `buildImage`, defined in BDL.

In the final step we modify the existing application, [Step 5: Modify your BDL application](4890-step-5-modify-your-bdl-application.md).

## Related links

**Related concepts**  

[WS client stubs and handlers](4616-ws-client-stubs-and-handlers.md "To access a remote Web Service, you first must get the WSDL information from the service provider.")
