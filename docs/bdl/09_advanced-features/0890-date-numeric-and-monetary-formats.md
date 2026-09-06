---
title: "Date, numeric and monetary formats"
source: "fgl-topics/c_fgl_localization_formats.html"
breadcrumb: "Advanced features > Localization > Application locale > Date, numeric and monetary formats"
type: "concept"
---

# Date, numeric and monetary formats

> This section describes how Genero BDL handles date, time, numeric and monetary formats.

Dates, numbers and monetary values must be displayed and entered in a format used in the
country/region. These formats can be defined with the [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") and [DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.") environment
variables.

Date and numeric format settings matter for data display and data input. For example, when
displaying a `DATE` value to a form field, it will implicitly be formatted by
the DBDATE setting. When the user enters a date in a form field bound to a `DATE`
variable, the entered digits will be interpreted by the DBDATE format.

The default values of these environment variables depend on the type of platform where the
program executes:

- On desktop/server platforms, the default formats are set for the United States of America:
  - Dates are formatted as `mm/dd/yyyy`.
  - The decimal separator is a dot.
  - The currency symbol is the dollar sign (`$`).
- On mobile platforms, the default formats are set from the regional settings defined on the device.
  - Dates are formatted depend on the regional settings.
  - The decimal separator is defined depend on the regional settings.
  - The currency symbol is not defined. No currency symbol will display.

  While it is possible to define environment settings for date and numeric formats with [FGLPROFILE entries](../07_configuration/0495-setting-environment-variables-in-fglprofile-mobile.md), it is strongly recommended
  to leave the defaults, so as to get the expected formats if the user changes the regional settings
  on the mobile device.

When using the `FORMAT` field attribute or the `USING` operator to
format dates with abbreviated day and month names- by using `ddd` /
`mmm` markers - the system uses English-language based texts for the conversion. This
means, day (`ddd`) and month (`mmm`) abbreviations are not localized
depending on the locale settings, they will always be in English.

## Related links

**Related concepts**  

[Formatting DATE values](../08_language-basics/0582-formatting-date-values.md "Date values must be formatted when converted to strings.")

[Formatting DATETIME values](../08_language-basics/0583-formatting-datetime-values.md "Date-time values must be formatted when converted to strings.")

[Formatting numeric values](../08_language-basics/0581-formatting-numeric-values.md "Numeric values must be formatted when converted to strings.")
