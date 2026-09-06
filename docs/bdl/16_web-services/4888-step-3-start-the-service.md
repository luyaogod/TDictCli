---
title: "Step 3: Start the service"
source: "fgl-topics/c_gws_java_APIs_009.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Calling Java from Genero > Step 3: Start the service"
type: "concept"
description: "Compile the Java class created in Step 2: Transform the Java class in a Web service , and run it. Commands to compile and execute the service in standalone mode: $ javac BarcodeService.java $ java ..."
---

# Step 3: Start the service

Compile the Java class created in [Step 2: Transform the Java class in a Web service](4887-step-2-transform-the-java-class-in-a-web-service.md), and run it.

Commands to compile and execute the service in standalone
mode:

```
$ javac BarcodeService.java
$ java BarcodeService
```

Once the service is started, it is ready to accept requests and you can also retrieve its WSDL at
the following URL:

http://localhost:9090/BarcodeService?WSDL

In our example, we anticipate the BarcodeService is on the localhost. For a production
environment, if you want the service to be started on a Web server, you must deploy it first using
Eclipse or the Web Server deployment tools.

In the next step we generate the client stub to access the BarcodeService, [Step 4: Generate BDL stub to access the Java library](4889-step-4-generate-bdl-stub-to-access-the-java-library.md).

## Related links

**Related concepts**  

[Deploy a Web Service](4907-deploy-a-web-service.md "Genero Web services need to be deployed on a Genero Application Server that plugs into a Web server, such as Apache. Understand what you need to implement for configuration and security.")
