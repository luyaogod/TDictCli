---
title: "VCard.VCAddress type"
source: "fgl-topics/c_fgl_utility_functions_vcard_vcaddress.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.VCAddress type"
type: "concept"
---

# VCard.VCAddress type

> The VCAddress structured type holds vCard address data.

## Syntax

```
TYPE VCAddress RECORD
    PostOfficeBox,
    ExtendedAddress, --  apartment or suite number
    Street,
    City,
    State,
    ZIP,
    Country STRING
    -- , CountryCode STRING -- X-ABADR:de
  END RECORD
```

## Usage

This type defines a record structure to hold vCard address information.
It is used for values returned by the [scan\_address()](2820-vcard-scan-address.md "Extracts an address from a string representing a vCard.") function.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE a VCard.VCAddress
  LET a.Street = "Sunset Bld"
END MAIN
```
