---
title: "JSON to BDL type conversion rules"
source: "fgl-topics/c_fgl_json_utils_conv_json_to_bdl.html"
breadcrumb: "Advanced features > JSON support > JSON to BDL type conversion rules"
type: "concept"
---

# JSON to BDL type conversion rules

> Specific type conversion rules apply when parsing a JSON string to fill a BDL variable.

| JSON source string | Target Genero BDL type |
| --- | --- |
| A JSON object in the form:{ "record-element-name" : json-value [,...] } | `RECORD ... END RECORD` |
| A JSON array in the form:[ json-value [,...] ] | `DYNAMIC ARRAY OF ...` |
| A JSON object in the form:{ "dictionary-key" : json-value [,...] } | `DICTIONARY OF ...` |
| A JSON value can be `null`, `true` or `false`.If the JSON value is a number or a string, the language conversion rules from number/string to `BOOLEAN` apply. | `BOOLEAN` |
| A JSON number.The JSON number can be assigned to any language numeric type. The limits of the target type cause potential overflows errors. On error the target variable will be initialized to `NULL`, the parser continues without an error. Some JSON classes have limitation based on binary64, see [JSON number limitations](0957-bdl-json-conversion-basics.md). | `TINYINT`, `SMALLINT`, `INTEGER`, `BIGINT`, `SMALLFLOAT`, `FLOAT`, `DECIMAL`, `MONEY` |
| A JSON string representing a date in the form `"YYYY-MM-DD"`. | `DATE` |
| A JSON string representing a date/time or a JSON number.If the value is a JSON number, it is interpreted as UNIX™ time (seconds since the Epoch 00:00:00 UTC, January 1, 1970).If the value is a JSON string, it must be formatted as `"YYYY-MM-DD hh:mm:ss.fffff"`, or represent an ISO 8601 formatted date-time, in UTC (with `Z` indicator) or with a timezone offset (`+/-hh[:mm]`). For example: `"2013-02-21T15:18:44.456Z"`, `"2013-02-21T20:18:44.456+02:00"`.Note that the ISO 8601 standard requires a `T` separator between the date and time part. However, it is possible that date/time values with time zone offset might deviate from ISO 8601, by using a space instead of the `T` separator.When the JSON date/time string does not specify a time zone offset, it is considered as a local time and assigned to the target variable without any conversion. When the JSON date/time string contains a time zone offset, this offset is applied to convert the JSON date/time string to a local time in the time zone of the system.See also [Character string to date time types](../08_language-basics/0578-data-type-conversion-reference.md)The `util.JSON` parser also supports Microsoft JSON AJAX old-style date/time strings in the `\/Date(nnn[Z\|[+\|-]hhmm])\/` form, where nnn is a number of milliseconds since epoch (always in UTC), `Z` indicates UTC and `+/-hhmm` defines a timezone offset. The UTC milliseconds since epoch are converted to a local time in the time zone of the system (the AJAX-style timezone offset is ignored). The JSON raw data string literal must contain a single backslash before the leading and trailing slash: `"\/Date(76523465)\/"`Date/time values can be converted from/to UTC by using the [`util.Datetime`](../15_library-reference/3487-util-datetime-methods.md "Methods for the util.Datetime class.") methods. | `DATETIME YEAR TO FRACTION(n)`, `DATETIME YEAR TO SECOND`, `DATETIME YEAR TO MINUTE`, `DATETIME YEAR TO HOUR` |
| A JSON string representing a year-month interval formatted as "`[-]nnnnnnnnn-nn`" such as "`-9834 11`", or in ISO 8601 duration format "`[-]PnnnnnnnnnYnnM`", such as "`-P9834Y10M`". | `INTERVAL YEAR TO MONTH`, `INTERVAL YEAR TO YEAR` |
| A JSON string representing a day-second interval formatted as "`[-]nnnnnnnnn nn:nn:nn[.nnnnn]`" such as "`510 11:34:56`", or in ISO 8601 duration format "`[-]PnnnnnnnnnDTnnHnnMnn[.nnnnn]S`", such as "`-P555DT14H55M59S`". | `INTERVAL DAY TO FRACTION(n)`, `INTERVAL DAY TO SECOND`, `INTERVAL DAY TO MINUTE`, `INTERVAL DAY TO HOUR` |
| A JSON string value encoded in Base64.The Base64 encoding is described in [[RFC4648]](http://www.ietf.org/rfc/rfc4648.txt). | `BYTE` (see note) |
| A JSON string or number.If the JSON value is a number, the resulting BDL string value uses the locale specific decimal point.If the JSON value is a string: Any character in the Basic Multilingual Plane (U+0000 through U+FFFF) may be escaped: `\u` followed by exactly 4 hexadecimal digits (`[0-9a-fA-F]`). The hexadecimal digits encode the code point. Characters outside the Basic Multilingual Plane may be escaped by there UTF-16 surrogate pairs. For example, the representation of the G clef character (`U+1D11E`) is "`\uD834\uDD1E`". | `TEXT` (see note), `CHAR`, `VARCHAR`, `STRING` |

> **Note:**
>
> When parsing a JSON string to fill a `TEXT` or
> `BYTE` variable, if the data storage for the LOB variable has not been defined with
> the [`LOCATE`](../08_language-basics/0699-locate-for-text-byte.md "The LOCATE statement specifies where to store data of TEXT and BYTE variables.") instruction, the JSON
> methods will automatically locate the `TEXT` or `BYTE` in memory. This
> applies also to `TEXT` and `BYTE` elements of records, arrays and
> dictionaries.
