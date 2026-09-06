---
title: "Step 2: Rename the generated files"
source: "fgl-topics/c_gws_NET_APIs_007.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Calling .NET from Genero > Step 2: Rename the generated files"
type: "concept"
description: "Rename the generated class called Service1 with an appropriate name such as \"BarCode\", and rename the file Service1.asmx to BarCodeService.asmx , for instance. The .asmx file is the file that is ..."
---

# Step 2: Rename the generated files

Rename the generated class called Service1 with an appropriate name such as
"BarCode", and rename the file Service1.asmx to
BarCodeService.asmx, for instance.

The .asmx file is the file that is accessible from the IIS web server once
the application is deployed. The .asmx file also contains a reference to the
default generated class, Service1, which must also be renamed to the new name
(BarCode in our tutorial), in case Visual Studio didn't make the change
automatically.

The class view after renaming the class:

![The class view after renaming the class:](../_images/Rename.jpg)

*Class View; BarCode selected*

The file view after renaming the asmx file:

![The file view after renaming the asmx file](../_images/RenameFile.jpg)

*File View; BarCodeService selected*

In the next step we add the barcode library as a reference, [Step 3: Add the barcode library as a reference](4899-step-3-add-the-barcode-library-as-a-reference.md).
