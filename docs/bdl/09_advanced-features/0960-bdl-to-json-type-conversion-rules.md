---
title: "BDL to JSON type conversion rules"
source: "fgl-topics/c_fgl_json_utils_conv_bdl_to_json.html"
breadcrumb: "Advanced features > JSON support > BDL to JSON type conversion rules"
type: "concept"
---

# BDL to JSON type conversion rules

> Specific type conversion rules apply when converting a BDL variable to JSON.

| Source Genero BDL type | JSON result string |
| --- | --- |
| `RECORD .. END RECORD` | The JSON string will be a JSON object in the form:{ "record-element-name" : json-value [,...] } |
| `DYNAMIC ARRAY OF ...` | The JSON string will be a JSON array in the form:[ json-value [,...] ] |
| `DICTIONARY OF ...` | The JSON string will be a JSON object in the form:{ "dictionary-key" : json-value [,...] } |
| `BOOLEAN` | Will be serialized with the JSON values `true` or `false`.For backward compatibility, the [`TRUE`](../08_language-basics/0573-true.md "TRUE is a predefined constant to be used in boolean expressions.") and [`FALSE`](../08_language-basics/0574-false.md "FALSE is a predefined constant to be used in boolean expressions.") constants are of the `INTEGER` type (1 and 0). If you use these constants directly in an API writing JSON data, they will result in numerical values `1` or `0`. To obtain JSON's `true` and `false` boolean values, use a `BOOLEAN` variable instead. |
| `TINYINT`, `SMALLINT`, `INTEGER`, `BIGINT`, `SMALLFLOAT`, `FLOAT`, `DECIMAL`, `MONEY` | Any numeric type will be serialized to this form: an optional minus sign (`-`), a sequence of digits (`0-9`), containing an optional decimal separator (`.`), followed by an optional exponent. The exponent has the form (`e`) followed by an optional minus sign and an sequence of digits. The representation of numeric values does not depend from the current locale. The decimal separator is always a dot (`.`). `MONEY` values will be represented like `DECIMAL` values: the currency symbol will be omitted. Some JSON classes have limitation based on binary64, see [JSON number limitations](0957-bdl-json-conversion-basics.md). |
| `DATE` | The date value will be formatted as `"YYYY-MM-DD"` (with double quotes) |
| `DATETIME qual1 TO qual2` | The date-time value will be formatted as `"YYYY-MM-DD hh:mm:ss.fffff"` (with double quotes), based on the date-time type definition, or as RFC 3339 formatted date/time, depending on the formatting mode defined with [`util.JSON.setDatetimeSerializationMode()`](../15_library-reference/3584-util-json-setdatetimeserializationmode.md "Defines the JSON formatting mode for DATETIME values.").For example, in default mode, a `DATETIME HOUR TO MINUTE` will produce `"hh:mm"` formatted values. In this form, the date/time value represents a local time on the system where the JSON data is produced.With the RFC3339/TZO mode, a `DATETIME YEAR TO MINUTE/SECOND/FRACTION` is formatted as `YYYY-MM-DDThh:mm:ss[.fffff]±hh:mm` as a local time, with a timezone offset according to the current TZ system settings. Other kind of `DATETIME` types cannot be converted to RFC 3339.With RFC3339/UTC mode, a `DATETIME YEAR TO MINUTE/SECOND/FRACTION` is converted to UTC according to the current TZ system settings, and formatted as `YYYY-MM-DDThh:mm:ss[.fffff]Z`.Other kind of `DATETIME` types cannot be converted to RFC 3339 and the formatting falls back to the default FGL format.Date/time values can be converted from/to UTC by using the [`util.Datetime`](../15_library-reference/3487-util-datetime-methods.md "Methods for the util.Datetime class.") methods. |
| `INTERVAL YEAR TO MONTH`, `INTERVAL YEAR TO YEAR` | The interval value will be formatted as `"YYYY-MM"` (with double quotes), or as ISO 8601 formatted duration `[-]PnnnnnnnnnYnnM`, depending on the formatting mode defined with [`util.JSON.setIntervalSerializationMode()`](../15_library-reference/3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values."). |
| `INTERVAL DAY TO FRACTION(n)`, `INTERVAL DAY TO SECOND`, `INTERVAL DAY TO MINUTE`, `INTERVAL DAY TO HOUR` | The interval value will be formatted as `"DD hh:mm:ss.fffff"` (with double quotes), or as ISO 8601 formatted duration `[-]PnnnnnnnnnDTnnHnnMnn[.nnnnn]S`, depending on the formatting mode defined with [`util.JSON.setIntervalSerializationMode()`](../15_library-reference/3585-util-json-setintervalserializationmode.md "Defines the JSON formatting mode for INTERVAL values."). |
| `BYTE` | Will be serialized to a Base64 encoded double quoted string. The Base64 encoding is described in [[RFC4648]](http://www.ietf.org/rfc/rfc4648.txt). |
| `TEXT`, `CHAR`, `VARCHAR`, `STRING` | Character string data will be serialized as a double quoted string with backslash escaping.List of characters requiring escaping: \\ backslash U+005C \" quotation mark U+0022 \b backspace U+0008 \f form feed U+000C \n line feed U+000A \r carriage return U+000D \t tab U+0009 |
| Other | Any other type will be serialized as a double quoted (`"`) string. |
