---
title: "Step 4: Add the buildImage method"
source: "fgl-topics/c_gws_NET_APIs_009.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Calling .NET from Genero > Step 4: Add the buildImage method"
type: "concept"
description: "Remove the default generated HelloWorld method, and create the buildImage method. Add the three using instructions to import the barcode library, and the instruction to declare buildImage as a ..."
---

# Step 4: Add the buildImage method

Remove the default generated `HelloWorld` method, and create the
`buildImage` method.

Add the three `using` instructions to import the barcode library, and the
instruction to declare `buildImage` as a WebMethod. Use the
`GetBarcodeBuilderType()` method to convert a string to a code as expected by the
barcode library.

```
using BarcodeLib;
using BarcodeLib.Barcode;
using BarcodeLib.Barcode.Linear;

namespace BarCodeService
{
    /// <summary>
    /// Summary description for Service1
    /// </summary>
    [WebService(Namespace = "http://tempuri.org/")]
    [WebServiceBinding(ConformsTo = WsiProfiles.BasicProfile1_1)]
    [ToolboxItem(false)]
    // To allow this Web Service to be called from script, using ASP.NET
    // [System.Web.Script.Services.ScriptService]
    public class Barcode : System.Web.Services.WebService
    {

         [WebMethod]
         public byte[] buildImage(String type, String code)
         {
             try
             {
                 Linear barcode = new Linear();
                 barcode.Data = code;
                 barcode.Type = GetBarcodeBuilderType(type);
                 barcode.AddCheckSum = true;

                 // save barcode image into your system
                 barcode.ShowText = true;
                 byte[] ret = barcode.drawBarcodeAsBytes();
                 if (ret !- null) return ret;
                 else return null;
             }
             catch (Exception e)
             {
                 return null;
             }
        }
    }
}
```

![BarCodeService.BarCode screenshot showing the code replicated above.](../_images/Code.jpg)

*BarCodeService.BarCode*

In the next step we publish the service to your IIS Web server, [Step 5: Publish the service](4901-step-5-publish-the-service.md).
