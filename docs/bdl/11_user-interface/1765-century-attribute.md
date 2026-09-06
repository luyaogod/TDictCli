---
title: "CENTURY attribute"
source: "fgl-topics/c_fgl_FSFAttributes_CENTURY.html"
breadcrumb: "User interface > Form definitions > Form specification files > Form item attributes > CENTURY attribute"
type: "concept"
---

# CENTURY attribute

> The CENTURY attribute defines expansion of the year in a DATE or DATETIME field.

## Syntax

```
CENTURY = { "R" | "C" |"F" | "P" }
```

## Usage

The CENTURY attribute specifies how to expand abbreviated one- and two-digit *year* specifications in a
[DATE](../08_language-basics/0558-date.md "The DATE data type stores calendar dates with a Year/Month/Day representation.") and [DATETIME](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.")
field.

Century expansion is based on this attribute and on the current year defined by the system clock.

The CENTURY attribute can specify any of four algorithms to expand abbreviated years into four-digit year
values that end with the same digits (or digit) that the user has entered.

CENTURY supports the same settings as the [DBCENTURY](../07_configuration/0507-dbcentury.md "Specifies the expansion for the century in DATE and DATETIME values.")
environment variable, but with a scope that is restricted to a single field.

If the CENTURY and DBCENTURY settings are different, CENTURY takes precedence.

Unlike DBCENTURY, the CENTURY attribute is not case sensitive. However, we recommend that you use uppercase
letters in the attribute.
