---
title: "VCard.format_person()"
source: "fgl-topics/c_fgl_utility_functions_vcard_format_person.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.format_person()"
type: "concept"
---

# VCard.format_person()

> Converts a VCPerson record to a vCard string representation vCard.

## Syntax

```
FUNCTION format_person(
   person VCPerson )
  RETURNS STRING
```

1. person is a VCPerson record.

## Usage

This function converts a record defined with the [VCPerson](2818-vcard-vcperson-type.md "The VCPerson structured type holds vCard data.") type, to a string
representing a vCard.

The returned value is a [version 3.0](http://tools.ietf.org/html/rfc6350) vCard formatted string.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE p VCard.VCPerson
  LET p.FirstName = "Hans"
  LET p.LastName = "Mustermann"
  LET p.email[1].VALUE = "hans@nomail.com"
  LET p.phone[1].TYPE = "HOME"
  LET p.phone[1].number = "+49 123 4567 8901"
  LET p.phone[2].TYPE = "WORK"
  LET p.phone[2].number = "+49 123 9876 5431"
  DISPLAY VCard.format_person(p.*)
END MAIN
```

Output:

```
BEGIN:VCARD
VERSION:3.0
N:Hans;Mustermann;;;
FN:Hans Mustermann
TEL;TYPE=HOME:+49 123 4567 8901
TEL;TYPE=WORK:+49 123 9876 5431
EMAIL:hans@nomail.com
END:VCARD
```
