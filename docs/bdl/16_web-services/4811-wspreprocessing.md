---
title: "WSPreProcessing"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSPreProcessing.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSPreProcessing"
type: "concept"
---

# WSPreProcessing

> Specify functions in a REST web service module that are used as preprocessing request handlers.

## Syntax

```
WSPreProcessing = "[ wsquery | wsheader | wsparam ]"
```

Where:

1. WSPreProcessing has options for query, header, and parameter:
   `wsquery`, `wsheader`, or `wsparam`. It is a string
   value setting.
2. When used on its own without an option, WSPreProcessing is set on a callback
   function for the body of the REST request.

`WSPreProcessing` is an attribute set on a BDL private function in a REST web
service module.

## Usage

You use this attribute to specify callback functions in the REST service module that run when an
incoming request is received and before the REST engine has processed it. The callback function must
be one of the predefined callback functions (see [Using preprocessing and postprocessing callbacks](4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.") and [Examples of callback functions](4752-examples-of-callback-functions.md "These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.")). These functions typically modify the
query (for example, validate or adjust query parameter values before the operation runs). The
WSPreProcessing attribute applies preprocessing callbacks to adjust headers,
parameters, or the body of the REST request.

Callbacks may also set `Context["HTTPProcessingErrorMsg"]` before returning a
non-zero status code. If specified, that string overrides the HTTP reason phrase sent to the
client. If the value is empty or NULL the standard phrase is used; if the key is not set a default
message ("Preprocessing done") is used. The key is ignored when the callback returns 0.

```
PRIVATE # REST Query string incoming callback
FUNCTION QueryPreCallback(query STRING, value STRING) ATTRIBUTES (WSPreProcessing="wsquery") RETURNS (STRING)
  DISPLAY "Query:",query
  DISPLAY "Value:",value
  # your code here
  RETURN value # new query value
END FUNCTION
```

## Example using WSPreProcessing

For sample callback functions, go to [Examples of callback functions](4752-examples-of-callback-functions.md "These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.").

## Related links

**Related concepts**  

[Using preprocessing and postprocessing callbacks](4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.")

[WSPostProcessing](4812-wspostprocessing.md "Specify functions in a REST web service module that are used as postprocessing request handlers.")

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")
