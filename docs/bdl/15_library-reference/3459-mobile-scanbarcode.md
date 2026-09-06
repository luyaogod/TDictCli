---
title: "mobile.scanBarCode"
source: "fgl-topics/c_fgl_frontcall_mobile_scanbarcode.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.scanBarCode"
type: "concept"
---

# mobile.scanBarCode

> Allow the user to scan a barcode with a mobile device

## Syntax

```
ui.Interface.frontCall("mobile", "scanBarCode",
   [options], [code, type] )
```

1. options - Is an optional record defined with the following members:
   - outputFormat - The output format of the bar code value. Values can be
     `"base64"` or `"hexadecimal"`. Default is a UTF-8 string.
2. code - Holds a string representation of the barcode.
3. type - Holds the name of the barcode type (`EAN`,
   `UPC`, `CODE_39`, `CODE_128`, `QR`,
   `DATA_MATRIX`, etc)

## Usage

The "`scanBarCode`" front call starts the barcode scanner to let the
user scan a barcode with the device.

> **Important:**
>
> To be usable with GBC, this front call needs a secure context / front-end
> connection, in addition to mobile device permissions, when required. A secure context
> is achieved by using the GAS via HTTPS, via localhost on the same machine, or running direct via
> GDC-UR. When using GDC-UR, additional security settings need to be enabled in the GDC configuration
> panel. See GDC documentation for more details.

If the barcode scan failed, the code return parameter is set to
`NULL` and type is set to `"canceled"`.

If case of successful barcode scan, the front call returns the string representation of the
barcode (code) and the barcode symbology (type):

- If the output format is not specified, a simple UTF-8 string is returned.
- When the output format is `"base64"` or `"hexadecimal"`, you get a
  string in the specified encoding.
- If the format is not recognized, the error detail `"Format '...' is not
  supported."` will be returned with error [-6333](4483-genero-bdl-errors.md).

For details about barcode types supported by GBC/JS, see [Barcode Detection API](https://developer.mozilla.org/en-US/docs/Web/API/Barcode_Detection_API#supported_barcode_formats).

## Example

```
MAIN
  DEFINE code, type, hex STRING
  DEFINE scanOptions RECORD
    outputFormat STRING
  END RECORD
  MENU
    ON ACTION barcode
        LET scanOptions.outputFormat = arg_val(1)
        CALL ui.Interface.frontCall(
            "mobile", "scanBarCode", [scanOptions], [code, type])
        DISPLAY code
    ON ACTION EXIT
        EXIT MENU
  END MENU
END MAIN
```
