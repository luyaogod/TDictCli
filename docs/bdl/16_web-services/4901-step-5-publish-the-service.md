---
title: "Step 5: Publish the service"
source: "fgl-topics/c_gws_NET_APIs_010.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Calling .NET from Genero > Step 5: Publish the service"
type: "concept"
description: "Build the entire application, right-click on the solution, and select the publish operation. This will copy all necessary files to your IIS web server and make your application available at a URL, ..."
---

# Step 5: Publish the service

Build the entire application, right-click on the solution, and
select the publish operation. This will copy all necessary files to
your IIS web server and make your application available at a URL,
depending on where you deploy it on your IIS web server.

In our tutorial, the service will be located at the root of the server. In other words, it will
be available at http://localhost/BarCodeService.asmx and the WSDL at URL
http://localhost/BarCodeService.asmx?WSDL

![Publish Web dialog screenshot](../_images/Publish.jpg)

*Publish Web dialog*

In the next step we generate the stub file, [Step 6: Generate .4gl stub to access the .NET library](4902-step-6-generate-4gl-stub-to-access-the-net-library.md).
