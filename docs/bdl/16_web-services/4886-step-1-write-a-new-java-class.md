---
title: "Step 1: Write a new java class"
source: "fgl-topics/c_gws_java_APIs_007.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Calling Java from Genero > Step 1: Write a new java class"
type: "concept"
description: "Instead of writing the functions in 4GL, you simply need to write them in a Java class with the methods you want to use in 4GL. In our example, the two functions are buildImage and readImage . And of ..."
---

# Step 1: Write a new java class

Instead of writing the functions in 4GL, you simply need to write them in a
Java class with the methods you want to use in 4GL. In our example, the two
functions are **buildImage** and **readImage**. And of course, don't forget to import the
necessary Java import
instructions.

```
import com.barcodelib.barcodereader.BarcodeReader;
import com.barcodelib.barcode.Barcode;
import java.io.*;

import javax.jws.*;
import javax.jws.soap.SOAPBinding;
import javax.xml.ws.Endpoint;

public class BarcodeService {
  public byte [] buildImage(String type, String data)
  {
  /*BUILDIMAGE IMPLEMENTATION CODE DESCRIBED ABOVE*/
  }
  public String readImage(String type,byte[] img)
  { { { {
  /*READIMAGE IMPLEMENTATION CODE DESCRIBED ABOVE*/
  } }
}
```

If you want the service to run standalone, you must also add the following main method to tell
the system the port number on which the service will
run:

```
public static void main(String[] args)
{
  String endpointUri = "http://localhost:9090/";
  Endpoint.publish(endpointUri, new BarcodeService ());
  System.out.println("BarcodeService started at " + endpointUri);
}
```

In the next step we transform the Java class in a Web Service, [Step 2: Transform the Java class in a Web service](4887-step-2-transform-the-java-class-in-a-web-service.md).
