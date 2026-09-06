---
title: "Step 2: Transform the Java class in a Web service"
source: "fgl-topics/c_gws_java_APIs_008.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Calling Java from Genero > Step 2: Transform the Java class in a Web service"
type: "concept"
description: "To transform the Java class created in Step 1: Write a new java class into a Web Service, simply add a WebService annotation: @WebService(targetNamespace = http://www.mycompany.com/barcode \", ..."
---

# Step 2: Transform the Java class in a Web service

To transform the Java class created in [Step 1: Write a new java class](4886-step-1-write-a-new-java-class.md) into a Web
Service, simply add a WebService
annotation:

```
@WebService(targetNamespace = http://www.mycompany.com/barcode ",
  name="Barcode",
  serviceName="BarcodeService")
public class BarcodeService{
...
}
```

This defines all public and non static methods of the class as
operations of the **BarcodeService** Web Service.

In the next step we start the Web Service, [Step 3: Start the service](4888-step-3-start-the-service.md).
