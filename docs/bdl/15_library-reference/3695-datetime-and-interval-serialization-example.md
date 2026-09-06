---
title: "Serializing DATETIME and INTERVAL values in standard formats"
source: "fgl-topics/c_gws_json_datetime_interval_serialization.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer options > Examples > DATETIME and INTERVAL serialization example"
type: "concept"
---

# Serializing DATETIME and INTERVAL values in standard formats

> Use the datetimeSerializationMode and intervalSerializationMode options to output DATETIME values in RFC 3339 format and INTERVAL values in ISO 8601 duration format.

By default, the JSON serializer writes [`DATETIME`](../08_language-basics/0559-datetime-qual1-to-qual2.md "The DATETIME data type stores date and time data with time units from the year to fractions of a second.") and `INTERVAL` values in the Genero string representation. You can change this behavior
using the [datetimeSerializationMode](3688-datetimeserializationmode.md "Controls the JSON output format for DATETIME values during serialization.") and [intervalSerializationMode](3689-intervalserializationmode.md "Controls the JSON output format for INTERVAL values during serialization.") options.

The following example sets both options and serializes a `DATETIME YEAR TO SECOND`
and an `INTERVAL DAY TO SECOND` value to JSON. The options are reset to their
defaults at the end of the program.

```
IMPORT json

MAIN
  DEFINE d  DATETIME YEAR TO SECOND
  DEFINE iv INTERVAL DAY TO SECOND

  LET d  = DATETIME(2024-11-05 09:15:58) YEAR TO SECOND
  LET iv = INTERVAL(2 10:30:00) DAY TO SECOND

  -- DATETIME serialization modes
  CALL json.Serializer.setOption("datetimeSerializationMode", 0)
  DISPLAY "datetime mode 0 (FGL legacy)    : ", serializeDatetime(d)
  CALL json.Serializer.setOption("datetimeSerializationMode", 1)
  DISPLAY "datetime mode 1 (RFC 3339 UTC)  : ", serializeDatetime(d)
  CALL json.Serializer.setOption("datetimeSerializationMode", 2)
  DISPLAY "datetime mode 2 (RFC 3339 offset): ", serializeDatetime(d)

  -- INTERVAL serialization modes
  CALL json.Serializer.setOption("intervalSerializationMode", 0)
  DISPLAY "interval mode 0 (FGL legacy)    : ", serializeInterval(iv)
  CALL json.Serializer.setOption("intervalSerializationMode", 1)
  DISPLAY "interval mode 1 (ISO 8601)      : ", serializeInterval(iv)

  -- Restore defaults
  CALL json.Serializer.setOption("datetimeSerializationMode", 0)
  CALL json.Serializer.setOption("intervalSerializationMode", 0)
END MAIN

FUNCTION serializeDatetime(value)
  DEFINE value  DATETIME YEAR TO SECOND
  DEFINE writer json.JSONWriter
  DEFINE t      TEXT
  LOCATE t IN MEMORY
  LET writer = json.JSONWriter.Create()
  CALL writer.setOutputCharset("UTF-8")
  CALL writer.writeToText(t)
  CALL writer.startJSON()
  CALL json.Serializer.VariableToJSON(value, writer)
  CALL writer.endJSON()
  CALL writer.close()
  RETURN t
END FUNCTION

FUNCTION serializeInterval(value)
  DEFINE value  INTERVAL DAY TO SECOND
  DEFINE writer json.JSONWriter
  DEFINE t      TEXT
  LOCATE t IN MEMORY
  LET writer = json.JSONWriter.Create()
  CALL writer.setOutputCharset("UTF-8")
  CALL writer.writeToText(t)
  CALL writer.startJSON()
  CALL json.Serializer.VariableToJSON(value, writer)
  CALL writer.endJSON()
  CALL writer.close()
  RETURN t
END FUNCTION
```

The code above would generate the following output with TZ=UTC (GMT):

```
datetime mode 0 (FGL legacy)    : "2024-11-05 09:15:58"
datetime mode 1 (RFC 3339 UTC)  : "2024-11-05T09:15:58Z"
datetime mode 2 (RFC 3339 offset): "2024-11-05T09:15:58+00:00"
interval mode 0 (FGL legacy)    : "  2 10:30:00"
interval mode 1 (ISO 8601)      : "P2DT10H30M00S"
```

## Related links

**Related concepts**  

[datetimeSerializationMode](3688-datetimeserializationmode.md "Controls the JSON output format for DATETIME values during serialization.")

[intervalSerializationMode](3689-intervalserializationmode.md "Controls the JSON output format for INTERVAL values during serialization.")

[json.Serializer.setOption](3675-json-serializer-setoption.md "Sets an option on the JSON serializer.")

**Related reference**  

[json.Serializer options](3685-json-serializer-options.md "Options for controlling serialization and deserialization behavior, set using the json.Serializer class.")
