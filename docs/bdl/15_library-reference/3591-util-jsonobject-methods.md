---
title: "util.JSONObject methods"
source: "fgl-topics/c_fgl_ext_util_JSONObject_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONObject class > util.JSONObject methods"
type: "concept"
---

# util.JSONObject methods

> Methods for the util.JSONObject class.

| Name | Description |
| --- | --- |
| util.JSONObject.create() RETURNS util.JSONObject | Creates a new JSON object. |
| util.JSONObject.fromFGL( record record-type ) RETURNS util.JSONObject | Creates a new JSON object from a `RECORD`. |
| util.JSONObject.parse( s STRING ) RETURNS util.JSONObject | Parses a JSON string and creates a JSON object from it. |

| Name | Description |
| --- | --- |
| get( name STRING ) RETURNS result-type | Returns the value corresponding to the specified entry name. |
| getLength() RETURNS INTEGER | Returns the number of name-value pairs in the JSON object. |
| getType( name STRING ) RETURNS STRING | Returns the type of a JSON object element. |
| has( name STRING ) RETURNS BOOLEAN | Checks if the JSON object contains a specific entry name. |
| name( index INTEGER ) RETURNS STRING | Returns the name of a JSON object entry by position. |
| put( name STRING, value value-type ) | Sets a name-value pair in the JSON object. |
| remove( name STRING ) | Removes the specified element in the JSON object. |
| toFGL( recordRef record-type ) | Fills a record variable with the entries contained in the JSON object. |
| toString() RETURNS STRING | Builds a JSON string from the values contained in the JSON object. |
