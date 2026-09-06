---
title: "Example BDL program"
source: "fgl-topics/c_gws_NET_APIs_013.html"
breadcrumb: "Web services > How Do I ... ? > Call .NET APIs from Genero in a SOA environment > Example BDL program"
type: "concept"
description: "This program calls the buildImage function of the Barcode .NET library. GLOBALS \"BarCode_BarCodeSoap.inc\" MAIN DEFINE wsstatus INTEGER IF num_args() != 3 THEN CALL ExitHelp() END IF LET ..."
---

# Example BDL program

This program calls the `buildImage` function of the Barcode .NET library.

```
GLOBALS "BarCode_BarCodeSoap.inc"
MAIN
  DEFINE wsstatus INTEGER

  IF num_args() != 3 THEN
    CALL ExitHelp()
  END IF

  LET buildImage.type = arg_val(1)
  LET buildImage.code = arg_val(2)
  LOCATE buildImageResponse.buildImageResult IN MEMORY

  LET wsstatus = buildImage_g()
  IF wsstatus <> 0 THEN
    DISPLAY "Error ("||wsError.code||") : ",wsError.description
  ELSE
    IF buildImageResponse.buildImageResult IS NULL THEN
      DISPLAY "Encoding failed"
    ELSE
      CALL buildImageResponse.buildImageResult.writeFile(arg_val(3))
    END IF
  END IF

  FREE buildImageResponse.buildImageResult

END MAIN

FUNCTION ExitHelp()
  DISPLAY arg_val(0)||" <type> <data> <filename>"
  DISPLAY "type : barcode type such as EAN8 or CODE128"
  DISPLAY "data : data to be encoded with a barcode [0-9A-D]"
  DISPLAY "filename : resulting image filename"
  DISPLAY "example : createImage EAN8 12358723A mybarcode.jpg"
  EXIT PROGRAM (-1)
END FUNCTION
```
