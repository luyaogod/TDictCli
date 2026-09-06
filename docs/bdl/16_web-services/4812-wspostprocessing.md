---
title: "WSPostProcessing"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSPostProcessing.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the module level > WSPostProcessing"
type: "concept"
---

# WSPostProcessing

> Specify functions in a REST web service module that are used as postprocessing request handlers.

## Syntax

```
WSPostProcessing = "[ wsheader ]"
```

Where:

1. WSPostProcessing has one option, `wsheader`. It is a string
   value setting.
2. When used on its own without an option, WSPostProcessing is set on a callback
   function for the body of the REST request.

`WSPostProcessing` is an attribute set on a BDL private function in a REST web
service module.

## Usage

You use this attribute to specify callback functions in the REST service module that run when an
incoming REST request is executed and before the REST response is forwarded to the client. The
callback function must be one of the predefined callback functions (see [Using preprocessing and postprocessing callbacks](4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.") and [Examples of callback functions](4752-examples-of-callback-functions.md "These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.")). These functions typically modify the
return header (for example, rename header names and values to include additional metadata). The
`WSPostProcessing` attribute applies postprocessing callbacks to adjust the response
body format, such as renaming JSON keys, normalizing XML tags, or transforming text content.

Postprocessing callbacks can also set `Context["HTTPProcessingErrorMsg"]` when
returning a non-zero code; the value becomes the HTTP reason phrase. Defaulting and behavior is the
same as for preprocessing callbacks; postprocessing messages override preprocessing ones when both
fail.

```
PRIVATE # REST Header outgoing callback
FUNCTION HeaderPostCallback(header STRING, value STRING) ATTRIBUTES (WSPostProcessing="wsheader") RETURNS (STRING)
  DISPLAY "Header:",HEADER
  DISPLAY "Value:",value
  # your code here
  RETURN value # new header value
END FUNCTION
```

## Example using WSPostProcessing

For details, go to [Examples of callback functions](4752-examples-of-callback-functions.md "These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.").

## Related links

**Related concepts**  

[Using preprocessing and postprocessing callbacks](4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.")

[WSPreProcessing](4811-wspreprocessing.md "Specify functions in a REST web service module that are used as preprocessing request handlers.")

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")
