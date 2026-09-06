---
title: "DATE, DATETIME, and INTERVAL conversions"
source: "fgl-topics/r_gws_json_temporal_formats.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > OpenAPI types mapping: BDL > DATE, DATETIME, and INTERVAL conversions"
type: "reference"
---

# DATE, DATETIME, and INTERVAL conversions

> GWS serializes DATE, DATETIME, and INTERVAL values as JSON strings using either the Genero BDL native format or standard RFC 3339 / ISO 8601 formats.

The output format depends on how the client stub is generated or how the serializer is
configured. Three output modes apply to `DATETIME`, two apply to
`INTERVAL`, and `DATE` is unaffected by the mode.

The `iso8601-local` mode includes the time zone offset of the server process at
runtime. As a result, the same `DATETIME` value can produce different output on
servers configured with different time zones.

| BDL type | `native` (BDL default) | `iso8601` (RFC 3339 UTC) | `iso8601-local` (RFC 3339 local offset) |
| --- | --- | --- | --- |
| `DATE` | `"2024-11-05"` | `"2024-11-05"` | `"2024-11-05"` |
| `DATETIME YEAR TO DAY` | `"2024-11-05"` | `"2024-11-05"` | `"2024-11-05"` |
| `DATETIME YEAR TO SECOND` | `"2024-11-05 09:15:58"` | `"2024-11-05T08:15:58Z"` | `"2024-11-05T09:15:58+01:00"` |
| `DATETIME YEAR TO FRACTION(3)` | `"2024-11-05 09:15:58.123"` | `"2024-11-05T08:15:58.123Z"` | `"2024-11-05T09:15:58.123+01:00"` |
| `INTERVAL YEAR TO MONTH` | `" 2-03"` | `"P2Y03M"` | `"P2Y03M"` |
| `INTERVAL DAY(9) TO FRACTION(5)` | `" 2 10:30:00.00000"` | `"P2DT10H30M00.00000S"` | `"P2DT10H30M00.00000S"` |

## DATE and DATETIME

OpenAPI represents date and time values as strings. The format keyword distinguishes between
date-only and date-time values.

## DATE

A `DATE` value is represented as a string with the date format:

```
schema
type         "string"
format       "date"
```

In this case, fglrestful generates a Genero BDL `DATE` type in
the client stub.

## DATETIME

For `DATETIME` values, GWS sets the `format` keyword based on the
qualifier:

- A `DATETIME YEAR TO DAY` is mapped
  to:

  ```
  schema
  type         "string"
  format       "date"
  ```

  because the value has no time component.
- A `DATETIME` ending at `MINUTE`, `SECOND`, or
  `FRACTION(n)` is mapped to:

  ```
  schema
  type         "string"
  format       "date-time"
  ```
- A `DATETIME` with other qualifiers (for example, `YEAR TO MONTH`
  or `HOUR TO MINUTE`) is mapped to a plain string with no format, because no RFC 3339
  representation is available:

  ```
  schema
  type         "string"
  ```

By default, fglrestful generates a `DATETIME YEAR TO SECOND`
type in the client stub.

You can control the fractional precision by setting the `--datetime-fraction`
option in the fglrestful command. For example:

```
--datetime-fraction=3
```

This generates a `DATETIME YEAR TO FRACTION(3)` type.

fglrestful generates the data type in the client stub, mapping it to a Genero
BDL `DATETIME YEAR TO SECOND` type.

## INTERVAL

OpenAPI represents durations as strings using the duration format:

```
schema
type         "string"
format       "duration"
```

GWS maps `INTERVAL` values (both `DAY TO SECOND` and `YEAR TO
MONTH`) to this representation.

By default, fglrestful generates a `INTERVAL DAY(9) TO
FRACTION(5)` type in the client stub.

You can override this behavior by using the `--interval` option to select a
different qualifier or to restore the previous mapping to `STRING`.

The duration format is defined in the OpenAPI Format Registry (JSON Schema 2019-09 / OpenAPI
3.1). Tools that do not support this format treat the value as a plain string.
