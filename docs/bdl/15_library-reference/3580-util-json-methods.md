---
title: "util.JSON methods"
source: "fgl-topics/c_fgl_ext_util_JSON_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSON class > util.JSON methods"
type: "concept"
---

# util.JSON methods

> Methods for the util.JSON class.

| Name | Description |
| --- | --- |
| util.JSON.format( s STRING ) RETURNS STRING | Formats a JSON string with indentation. |
| util.JSON.parse( s STRING, variableRef any-type ) | Parses a JSON string and fills program variables with the values. |
| util.JSON.proposeType( s STRING ) RETURNS STRING | Describes the record structure that can hold a given JSON data string. |
| util.JSON.setDatetimeSerializationMode( s STRING ) | Defines the JSON formatting mode for `DATETIME` values. |
| util.JSON.setIntervalSerializationMode( s STRING ) | Defines the JSON formatting mode for `INTERVAL` values. |
| util.JSON.stringify( value any-type ) RETURNS STRING | Produces a JSON formatted string from the input, by including empty records and empty arrays. |
| util.JSON.stringifyOmitNulls( value any-type ) RETURNS STRING | Produces a JSON formatted string from the input, by excluding empty records and empty arrays. |
