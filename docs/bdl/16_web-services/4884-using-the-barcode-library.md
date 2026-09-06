---
title: "Using the barcode library"
source: "fgl-topics/c_gws_java_APIs_005.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Using the barcode library"
type: "concept"
description: "The barcode library is composed of two libraries: A library for building a barcode image from a numeric code A library for reading a barcode image to return the numeric code This section is based on ..."
---

# Using the barcode library

The **barcode library** is composed of two libraries:

- A library for building a barcode image from a numeric code
- A library for reading a barcode image to return the numeric code

This section is based on the library you want to use in Genero.

In our tutorial, we create two functions called **buildImage** and **readImage**.

This is the Java implementation:

`buildImage( type : String, code : String) : byte[
]`

```
try {
  Barcode builder=new Barcode();
  builder.setType(GetBarcodeBuilderType(type));
  builder.setData(data);
  builder.setAddCheckSum(true);
  ByteArrayOutputStream out=new ByteArrayOutputStream();
  if (builder.createBarcodeImage(out)) {
   byte[] ret = out.toByteArray();    
   return ret;
 } else {
   return null;
 }
} catch (Exception e) {
  return null;
}
```

`readImage( type : String, img : byte[ ]) :
String`

```
try {
  File f=new File("tmp.jpg"); 
  FileOutputStream stream=new FileOutputStream(f);
  stream.write(img);
  stream.close();  String[] datas = 
    BarcodeReader.read(f, GetBarcodeReaderType(type));
  if (datas==null) {
    return null;
  } else {
    String ret = datas[0];
    return ret;
  }  
} catch (Exception e) {
  return null;
}
```

The following two functions convert the barcode type to the type expected by the
library:

```
private int GetBarcodeBuilderType(String str) {
  if (str.equals("CODABAR")) {
    return Barcode.CODABAR;
  } else if (str.equals("CODE11")) {
    return Barcode.CODE11;  
  } else if (str.equals("CODE128")) {
    return Barcode.CODE128;  
  } else if (str.equals("CODE128A")) {
    return Barcode.CODE128A;  
  } else if (str.equals("CODE128B")) {
    return Barcode.CODE128B;  
  } else if (str.equals("CODE128C")) {
    return Barcode.CODE128C;  
  } else if (str.equals("CODE2OF5")) {
    return Barcode.CODE2OF5;  
  } else if (str.equals("CODE39")) {
    return Barcode.CODE39;  
  } else if (str.equals("CODE39EX")) {
    return Barcode.CODE39EX;  
  } else if (str.equals("CODE93")) {
    return Barcode.CODE93;  
  } else if (str.equals("CODE93EX")) {
    return Barcode.CODE93EX;  
  } else if (str.equals("EAN13")) {
    return Barcode.EAN13;  
  } else if (str.equals("EAN13_2")) {
    return Barcode.EAN13_2;  
  } else if (str.equals("EAN13_5")) {
    return Barcode.EAN13_5;  
  } else if (str.equals("EAN8")) {
    return Barcode.EAN8;  
  } else if (str.equals("EAN8_2")) {
    return Barcode.EAN8_2;  
  } else if (str.equals("EAN8_5")) {
    return Barcode.EAN8_5;  
  } else if (str.equals("INTERLEAVED25")) {
    return Barcode.INTERLEAVED25;  
  } else if (str.equals("ITF14")) {
    return Barcode.ITF14;  
  } else if (str.equals("ONECODE")) {
    return Barcode.ONECODE;  
  } else if (str.equals("PLANET")) {
    return Barcode.PLANET;  
  } else if (str.equals("POSTNET")) {
    return Barcode.POSTNET;  
  } else if (str.equals("RM4SCC")) {
    return Barcode.RM4SCC;  
  } else if (str.equals("UPCA")) {
    return Barcode.UPCA;  
  } else if (str.equals("UPCE")) {
    return Barcode.UPCE;  
  } else {
    return -1;   
  }
}

private int GetBarcodeReaderType(String str) {
  if (str.equals("CODABAR")) {
    return BarcodeReader.CODABAR;  
  } else if (str.equals("CODE11")) {
    return BarcodeReader.CODE11;  
  } else if (str.equals("CODE128")) {
    return BarcodeReader.CODE128;  
  } else if (str.equals("CODE39")) {
    return BarcodeReader.CODE39;  
  } else if (str.equals("CODE39EX")) {
    return BarcodeReader.CODE39EX;  
  } else if (str.equals("CODE93")) {
    return BarcodeReader.CODE93;  
  } else if (str.equals("DATAMATRIX")) {
    return BarcodeReader.DATAMATRIX;  
  } else if (str.equals("EAN13")) {
    return BarcodeReader.EAN13;  
  } else if (str.equals("EAN8")) {
    return BarcodeReader.EAN8;  
  } else if (str.equals("INTERLEAVED25")) {
    return BarcodeReader.INTERLEAVED25;  
  } else if (str.equals("ITF14")) {
    return BarcodeReader.ITF14;  
  } else if (str.equals("ONECODE")) {
    return BarcodeReader.ONECODE;  
  } else if (str.equals("PLANET")) {
    return BarcodeReader.PLANET;  
  } else if (str.equals("POSTNET")) {
    return BarcodeReader.POSTNET;  
  } else if (str.equals("QRCODE")) {
    return BarcodeReader.QRCODE;  
  } else if (str.equals("RM4SCC")) {
    return BarcodeReader.RM4SCC;  
  } else if (str.equals("RSS14")) {
    return BarcodeReader.RSS14;  
  } else if (str.equals("RSSLIMITED")) {
    return BarcodeReader.RSSLIMITED;  
  } else if (str.equals("UPCA")) {
    return BarcodeReader.UPCA;  
  } else if (str.equals("UPCE")) {  
    return BarcodeReader.UPCE;  
  } else {
    return -1;   
  }
}
```
