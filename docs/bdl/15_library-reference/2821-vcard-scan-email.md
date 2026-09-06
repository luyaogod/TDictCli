---
title: "VCard.scan_email()"
source: "fgl-topics/c_fgl_utility_functions_vcard_scan_email.html"
breadcrumb: "Library reference > Utility modules > VCard: VCF file format module > VCard.scan_email()"
type: "concept"
---

# VCard.scan_email()

> Extracts an email from a string representing a vCard.

## Syntax

```
FUNCTION scan_email(
   source STRING,
   type STRING )
  RETURNS STRING
```

1. source is the vCard string ([version 3.0](http://tools.ietf.org/html/rfc6350)).
2. type is the type of email (HOME, WORK, pref).

## Usage

This function parses the vCard string passed as parameter to find "EMAIL" data based on a
type, and returns the email address as a string.

The function looks for lines starting with the "EMAIL" keyword.

The second parameter (type) defines is the value of the "TYPE" attribute
in an "EMAIL" line. Values can, for example, be "HOME", "WORK", "pref". If this parameter is
`NULL`, the email with TYPE=pref will be returned. If no preferred email exists, the
first email will be returned.

## Example

```
IMPORT FGL VCard
MAIN
  DEFINE m STRING,
         f TEXT
  LOCATE f IN FILE arg_val(1)
  CALL VCard.scan_email(f,NULL) RETURNING m
  DISPLAY m
END MAIN
```
