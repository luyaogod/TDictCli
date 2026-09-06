---
title: "json.Serializer methods"
source: "fgl-topics/c_gws_jsonSerializer_methods.html"
breadcrumb: "Library reference > Extension packages > The json package > The streaming API for JSON classes > The json.Serializer class > json.Serializer methods"
type: "concept"
---

# json.Serializer methods

> Methods for the json.Serializer class.

| Name | Description |
| --- | --- |
| JSONToVariable( json json.JSONReader, var AnyType ) | Serializes entries contained in a JSONReader object into a Genero BDL variable. |
| VariableToJSON( var AnyType, json json.JSONWriter ) | Serializes a Genero BDL variable into JSON using a JSONWriter object. |
| getOption( name STRING) RETURNS INTEGER | Returns the current value of a JSON serializer option. |
| setOption( name STRING, value INTEGER ) | Sets an option on the JSON serializer. |
| getLastErrorDescription() RETURN STRING | Retrieves the most recent error generated during serialization. |
