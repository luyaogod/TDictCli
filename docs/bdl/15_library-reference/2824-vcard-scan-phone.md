---
title: "VCard.scan_phone()"
source: "fgl-topics/c_fgl_utility_functions_vcard_scan_phone.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.scan_phone()"
type: "concept"
---

# VCard.scan_phone()

> Extracts a phone number from a string representing a vCard.

## Syntax

```
FUNCTION scan_phone(
   source STRING,
   type STRING )
  RETURNS STRING
```

1. source is the vCard string ([version 3.0](http://tools.ietf.org/html/rfc6350)).
2. type is the type of phone number (HOME, WORK, TEXT, VOICE, FAX, CELL,
   VIDEO, PAGER, TEXTPHONE, pref).

## Usage

This function parses the vCard string passed as parameter to find phone data based on a
type, and returns the phone number in a string.

The function looks for lines starting with the "TEL" keyword.

The second parameter (type) defines the value of the "TYPE" attribute in a
"TELs" line. Values can, for example, be "HOME", "WORK", "TEXT", "VOICE", "FAX", "CELL", "VIDEO",
"PAGER", "TEXTPHONE", "pref". If this parameter is `NULL`, the phone number with
TYPE=pref will be returned. If no preferred phone number exists, the first phone number will be
returned.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE n STRING,
         f TEXT
  LOCATE f IN FILE arg_val(1)
  CALL VCard.scan_phone(f,NULL) RETURNING n
  DISPLAY n
END MAIN
```
