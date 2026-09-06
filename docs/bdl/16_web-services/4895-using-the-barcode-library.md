---
title: "Using the barcode library"
source: "fgl-topics/c_gws_NET_APIs_004.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Using the barcode library"
type: "concept"
description: "This section is based on using the library in Genero. In our tutorial, we create one function called buildImage . This is the C# implementation: buildImage( type : String, code : String) : byte[] ..."
---

# Using the barcode library

This section is based on using the library in Genero. In our tutorial, we create one function
called `buildImage`. This is the C# implementation:

`buildImage( type : String, code : String) :
byte[]`

```
Linear barcode = new Linear();
barcode.Data = code;
barcode.Type = GetBarcodeBuilderType(type);
barcode.AddCheckSum = true;
// save barcode image into your system
barcode.ShowText = true;
byte[] ret = barcode.drawBarcodeAsBytes();
if (ret != null) return ret;
else return null;
```

You will also need to convert the barcode type to the type expected by the library. Therefore,
you will need this function.

```
private BarcodeType GetBarcodeBuilderType(String str)
{
if (str.Equals("CODABAR")) {
      return BarcodeType.CODABAR;
   } else if (str.Equals("CODE11")) {
      return BarcodeType.CODE11;
   } else if (str.Equals("CODE128")) {
      return BarcodeType.CODE128;
   } else if (str.Equals("CODE128A")) {
      return BarcodeType.CODE128A;
   } else if (str.Equals("CODE128B")) {
      return BarcodeType.CODE128B;
   } else if (str.Equals("CODE128C")) {
      return BarcodeType.CODE128C;
   } else if (str.Equals("CODE2OF5")) {
      return BarcodeType.CODE2OF5;
   } else if (str.Equals("CODE39")) {
      return BarcodeType.CODE39;
   } else if (str.Equals("CODE39EX")) {
      return BarcodeType.CODE39EX;
   } else if (str.Equals("CODE93")) {
      return BarcodeType.CODE93;
   } else if (str.Equals("EAN13")) {
      return BarcodeType.EAN13;
   } else if (str.Equals("EAN13_2")) {
      return BarcodeType.EAN13_2;
   } else if (str.Equals("EAN13_5")) {
      return BarcodeType.EAN13_5;
   } else if (str.Equals("EAN8")) {
      return BarcodeType.EAN8;
   } else if (str.Equals("EAN8_2")) {
      return BarcodeType.EAN8_2;
   } else if (str.Equals("EAN8_5")) {
      return BarcodeType.EAN8_5;
   } else if (str.Equals("INTERLEAVED25")) {
      return BarcodeType.INTERLEAVED25;
   } else if (str.Equals("ITF14")) {
      return BarcodeType.ITF14;
   } else if (str.Equals("ONECODE")) {
      return BarcodeType.ONECODE;
   } else if (str.Equals("PLANET")) {
      return BarcodeType.PLANET;
   } else if (str.Equals("POSTNET")) {
      return BarcodeType.POSTNET;
   } else if (str.Equals("RM4SCC")) {
      return BarcodeType.RM4SCC;
   } else if (str.Equals("UPCA")) {
      return BarcodeType.UPCA;
   } else if (str.Equals("UPCE")) {
      return BarcodeType.UPCE;
   } else {
      throw new Exception();
   }
}
```
