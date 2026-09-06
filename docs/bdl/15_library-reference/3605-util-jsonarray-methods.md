---
title: "util.JSONArray methods"
source: "fgl-topics/c_fgl_ext_util_JSONArray_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.JSONArray class > util.JSONArray methods"
type: "concept"
---

# util.JSONArray methods

> Methods for the util.JSONArray class.

| Name | Description |
| --- | --- |
| util.JSONArray.create() RETURNS util.JSONArray | Creates a new JSON array object. |
| util.JSONArray.fromFGL( array dynamic-array-type ) RETURNS util.JSONArray | Creates a new JSON array object from a `DYNAMIC ARRAY`. |
| util.JSONArray.parse( s STRING ) RETURNS util.JSONArray | Parses a JSON string and creates a JSON array object from it. |

| Name | Description |
| --- | --- |
| get( index INTEGER ) RETURNS result-type | Returns the value of a JSON array element. |
| getLength() RETURNS INTEGER | Returns the number of elements in the JSON array object. |
| getType( index INTEGER ) RETURNS STRING | Returns the type of a JSON array element. |
| put( index INTEGER, value value-type ) | Sets an element by position in the JSON array object. |
| remove( index INTEGER ) | Removes the specified entry in the JSON array object. |
| toFGL( arrayRef dynamic-array-type ) | Fills a dynamic array variable with the elements contained in the JSON array object. |
| toString() RETURNS STRING | Builds a JSON string from the elements contained in the JSON array object. |
