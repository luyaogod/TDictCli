---
title: "JSON utility classes"
source: "fgl-topics/c_fgl_json_utils_classes.html"
breadcrumb: "Advanced features > JSON support > JSON utility classes"
type: "concept"
---

# JSON utility classes

> Genero BDL provides utility classes to manipulate JSON formatted data.

## JSON classes from the `util` package

The following JSON utility classes provided in the `"util"` extension package can
be used to manipulate JSON data in memory:

- [`util.JSON`](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") implements basic JSON
  to/from FGL types conversion.
- [`util.JSONObject`](../15_library-reference/3590-the-util-jsonobject-class.md "The util.JSONObject class provides methods to handle an structured data object following the JSON string syntax.") implements
  detail JSON Object control.
- [`util.JSONArray`](../15_library-reference/3604-the-util-jsonarray-class.md "The util.JSONArray class provides methods to handle an array of values, following the JSON string syntax.") implements
  detail JSON Array control.

## JSON classes from the `json` package

The following JSON utility classes provided in the `"json"` extension package can
be used for streaming JSON data, especially with Genero Web Services:

- [`json.Serializer`](../15_library-reference/3662-the-json-serializer-class.md "The json.Serializer class provides methods to serialize Genero BDL variable to JSON objects and vice versa.") implements
  methods to serialize BDL variable to/from JSON objects.
- [`json.JSONWriter`](../15_library-reference/3618-the-json-jsonwriter-class.md "The json.JSONWriter class provides an interface compatible with JSON streaming that writes data in a JSON format to an output source.") provides an interface
  that writes data in a JSON format to an output source.
- [`json.JSONReader`](../15_library-reference/3639-the-json-jsonreader-class.md "The json.JSONReader class provides an interface compatible with JSON streaming that reads data in a JSON format from an input source.") provides an interface
  that reads data in a JSON format from an input source.
