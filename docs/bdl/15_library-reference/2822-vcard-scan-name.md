---
title: "VCard.scan_name()"
source: "fgl-topics/c_fgl_utility_functions_vcard_scan_name.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.scan_name()"
type: "concept"
---

# VCard.scan_name()

> Extracts name information from a string representing a vCard.

## Syntax

```
FUNCTION scan_name(
   source STRING)
  RETURNS VCName
```

1. source is the vCard string ([version 3.0](http://tools.ietf.org/html/rfc6350)).

## Usage

This function parses the vCard string passed as parameter to find a person's name data, and
returns name information in a record defined with the [VCName](2817-vcard-vcname-type.md "The VCName structured type holds vCard data related to the person's name.") type.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE n VCard.VCName,
         f TEXT
  LOCATE f IN FILE arg_val(1)
  CALL VCard.scan_name(f) RETURNING n.*
  DISPLAY n.*
END MAIN
```
