---
title: "json.JSONReader methods"
source: "fgl-topics/c_gws_jsonJSONReader_methods.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.JSONReader class > json.JSONReader methods"
type: "concept"
---

# json.JSONReader methods

> Methods for the json.JSONReader class.

| Name | Description |
| --- | --- |
| json.JSONReader.Create() RETURNS json.JSONReader | Constructor of a JSONReader object. |

| Name | Description |
| --- | --- |
| setInputCharset( charset STRING ) | Defines the charset used on the input stream. |

| Name | Description |
| --- | --- |
| readFrom( name STRING ) | Sets the input stream of the JSONReader object to a file or an URL. |
| readFromPipe( cmd STRING ) | Sets the input stream of the JSONReader object to a PIPE. |
| readFromText( txt TEXT ) | Sets the output stream of the JSONReader object to a TEXT large object. |
| getProperty() RETURNS STRING | Gets a property of a JSONReader object. |
| getValue() RETURNS STRING | Gets the value of a JSONReader object. |
| getBooleanValue( b BOOLEAN) | Gets a boolean value of a JSONReader object. |
| getDecimalValue( de DECIMAL) | Gets a decimal value of a JSONReader object. |
| getByteValue( by BYTE) | Gets a byte value of a JSONReader object. |
| getTextValue( txt TEXT ) | Gets a text value of the current JSON node. |
| getDateTimeValue( dt DATETIME ) | Gets a datetime value of the current JSON node. |
| getEventType() RETURNS STRING | Returns a string that indicates the type of event the cursor of the JSONReader object is pointing to. |

| Name | Description |
| --- | --- |
| hasNext() RETURNS BOOLEAN | Checks whether the JSONReader cursor can be moved to a node next to it. |
| next() | Moves the JSONReader cursor to the next node. |
| getSize() RETURNS INTEGER | Returns the number of bytes read by the JSONReader in the current stream. |
| close() | Closes the JSONReader streaming, and releases all associated resources. |
