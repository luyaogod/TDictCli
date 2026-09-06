---
title: "Example program"
source: "fgl-topics/c_gws_java_APIs_012.html"
breadcrumb: "Web services > How Do I ... ? > Call Java APIs from Genero in a SOA environment > Example program"
type: "concept"
description: "This program calls the buildImage function of the Barcode Java library. GLOBALS \"BarcodeService_BarcodePort.inc\" MAIN DEFINE wsstatus INTEGER IF num_args() != 3 THEN CALL ExitHelp() END IF LET ..."
---

# Example program

This program calls the `buildImage` function of the Barcode Java library.

```
GLOBALS "BarcodeService_BarcodePort.inc"

MAIN

DEFINE wsstatus INTEGER

IF num_args() != 3 THEN
   CALL ExitHelp()
END IF

LET ns1buildImage.arg0 = arg_val(1)
LET ns1buildImage.arg1 = arg_val(2)
LOCATE ns1buildImageResponse.return IN MEMORY

LET wsstatus = buildImage_g()
IF wsstatus <> 0 THEN
  DISPLAY "Error ("||wsError.code||") : ",wsError.description
ELSE
  IF ns1buildImageResponse.return IS NULL THEN
    DISPLAY "Encoding failed" 
  ELSE
    CALL ns1buildImageResponse.return.writeFile(arg_val(3))
  END IF
END IF

FREE ns1buildImageResponse.return

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
