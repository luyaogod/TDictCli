---
title: "VCard.VCName type"
source: "fgl-topics/c_fgl_utility_functions_vcard_vcname.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.VCName type"
type: "concept"
---

# VCard.VCName type

> The VCName structured type holds vCard data related to the person's name.

## Syntax

```
TYPE VCName RECORD
    FirstName,
    LastName,
    MiddleName,
    Prefix,
    Suffix STRING
    --, FormattedName STRING
  END RECORD
```

## Usage

This type defines a record structure to hold vCard information related to the person's name. It
is used for values returned by the [scan\_name()](2822-vcard-scan-name.md "Extracts name information from a string representing a vCard.") function.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE n VCard.VCName
  LET n.FirstName = "Hans"
  LET n.LastName = "Mustermann"
END MAIN
```
