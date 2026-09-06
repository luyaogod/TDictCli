---
title: "json.JSONWriter methods"
source: "fgl-topics/c_gws_jsonJSONWriter_methods.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONWriter class > json.JSONWriter methods"
type: "concept"
---

# json.JSONWriter methods

> Methods for the json.JSONWriter class.

| Name | Description |
| --- | --- |
| json.JSONWriter.Create() RETURNS json.JSONWriter | Constructor of a JSONWriter object. |

| Name | Description |
| --- | --- |
| setOutputCharset( charset STRING ) | Defines the charset used on the output stream. |

| Name | Description |
| --- | --- |
| writeTo( url STRING ) | Sets the output stream of the JSONWriter object to a file or an URL, and starts streaming data. |
| writeToPipe( cmd STRING ) | Sets the output stream of the JSONWriter object to a PIPE, and starts the streaming. |
| writeToText( txt TEXT ) | Sets the output stream of the JSONWriter object to a TEXT large object, and starts the streaming. |
| setProperty( name STRING) | Sets a property of a JSONWriter object. |
| setValue( value primitive-type) | Sets the value of a JSONWriter object. |
| setBooleanValue( value BOOLEAN) | Sets a boolean value of a JSONWriter object. |
| setNullValue() | Sets the null value of a JSONWriter object. |
| getSize() RETURNS INTEGER | Returns the number of bytes written by the JSONWriter in the current stream. |

| Name | Description |
| --- | --- |
| endJSON() | Closes any open tokens and writes the corresponding end JSON document token. |
| startJSON() | Begins JSON encoding to the JSONWriter stream. |
| endObject() | Writes a JSON end object token to the JSONWriter stream. |
| startObject() | Writes a JSON start object token to the JSONWriter stream. |
| endArray() | Writes a JSON end array tag to the JSONWriter stream. |
| startArray() | Writes a JSON start array token to the JSONWriter stream. |
| close() | Closes the JSONWriter streaming, and releases all associated resources. |
