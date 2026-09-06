---
title: "VCard.scan_person()"
source: "fgl-topics/c_fgl_utility_functions_vcard_scan_person.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.scan_person()"
type: "concept"
---

# VCard.scan_person()

> Extracts person's data from a string representing a vCard.

## Syntax

```
FUNCTION scan_person(
   source STRING )
  RETURNS VCPerson
```

1. source is the vCard string ([version 3.0](http://tools.ietf.org/html/rfc6350)).

## Usage

This function parses the vCard string passed as parameter, extracts all information, and returns
a record defined with the [VCPerson](2818-vcard-vcperson-type.md "The VCPerson structured type holds vCard data.") type.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE p VCard.VCPerson,
         f TEXT
  LOCATE f IN FILE arg_val(1)
  CALL VCard.scan_person(f) RETURNING p.*
  DISPLAY p.*
END MAIN
```
