---
title: "VCard.scan_address()"
source: "fgl-topics/c_fgl_utility_functions_vcard_scan_address.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.scan_address()"
type: "concept"
---

# VCard.scan_address()

> Extracts an address from a string representing a vCard.

## Syntax

```
FUNCTION scan_address(
   source STRING,
   type STRING )
  RETURNS VCAddress
```

1. source is the vCard string ([version 3.0](http://tools.ietf.org/html/rfc6350)).
2. type is the type of address (HOME, WORK, pref).

## Usage

This function parses the vCard string passed as parameter to find address data based on
a type, and returns address information in a record defined with the
`VCAddress` type.

The function looks for lines starting with the "ADR" keyword.

The second parameter (type) defines the value of the "TYPE" attribute in an
"ADR" line. Values can, for example, be "HOME", "WORK", "pref". If this parameter is
`NULL`, the address with TYPE=pref will be returned. If no preferred address exists,
the first address will be returned.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE a VCard.VCAddress,
         f TEXT
  LOCATE f IN FILE arg_val(1)
  CALL VCard.scan_address(f,"WORK") RETURNING a.*
  DISPLAY a.*
END MAIN
```
