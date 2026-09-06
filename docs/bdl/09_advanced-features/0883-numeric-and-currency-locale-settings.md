---
title: "Numeric and currency locale settings"
source: "fgl-topics/c_fgl_localization_018.html"
breadcrumb: "Advanced features > Localization > Application locale > Defining the application locale > Numeric and currency locale settings"
type: "concept"
description: "The environment variables LC_MONETARY and LC_NUMERIC are ignored . To perform decimal to/from string conversions, the runtime system uses the DBMONEY or DBFORMAT environment variables. These variables ..."
---

# Numeric and currency locale settings

The environment variables LC\_MONETARY and LC\_NUMERIC are ignored.

To perform decimal to/from string conversions, the runtime system uses the DBMONEY or DBFORMAT
environment variables. These variables define hundreds / decimal separators and currency symbols for
`MONEY` data types.

## Related links

**Related concepts**  

[Date, numeric and monetary formats](0890-date-numeric-and-monetary-formats.md "This section describes how Genero BDL handles date, time, numeric and monetary formats.")

[DBFORMAT](../07_configuration/0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values.")

[DBMONEY](../07_configuration/0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.")
