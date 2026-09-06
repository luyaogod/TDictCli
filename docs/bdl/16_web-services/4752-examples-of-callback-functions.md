---
title: "Examples of callback functions"
source: "fgl-topics/c_gws_restful_high_level_callback_examples.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Using preprocessing and postprocessing callbacks > Examples of callback functions"
type: "concept"
---

# Examples of callback functions

> These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.

These samples are intended as templates for your callback functions. Copy those you need to your
REST web service module and adapt them to your requirements. Callbacks let you inspect and modify
requests before the REST operation runs (preprocessing) or adjust responses before returning them to
the client (postprocessing).

## Preprocessing callbacks

Use these to validate or modify the query, parameters, header, and body of the REST request. They
are set with the `WSPreProcessing` attribute. The function name does not matter. Names with
`*PreCallback` are conventions for clarity, not used for
selection.

The `HTTPPreCallback` function runs first in the [callback order](4751-using-preprocessing-and-postprocessing-callbacks.md). It can return `0` to continue processing or an HTTP return
code to stop and return an error to the client. When returning a non-zero code the callback may also
set `Context["HTTPProcessingErrorMsg"]` to supply a custom HTTP reason phrase; if
the key is empty or not set a default phrase is used.

> **Tip:**
>
> In your preprocessing and postprocessing functions, you can inspect the request using the [`WSContext`](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.")
> dictionary, which contains details like the resource path, HTTP verb, and function name. This lets
> you check whether the request meets your expectations before proceeding.

## Examples of preprocessing BDL functions

**Before using these examples**, note that they only display values and return
`0`. In a real implementation, add logic to validate requests before the REST
operation runs. If validation fails in any of these callbacks, return an HTTP status code (such as
`401 Unauthorized` or `400 Bad Request`) instead of
`0`.

```
# REST Web service module
# Examples of preprocessing BDL functions

IMPORT util
IMPORT XML

PRIVATE DEFINE Context DICTIONARY OF STRING

PRIVATE # HTTP incoming request callback 
FUNCTION HTTPPreCallback() ATTRIBUTES (WSPreProcessing) RETURNS (INTEGER)
  DISPLAY "Doing HTTP Pre processing"
  DISPLAY Context["RequestOPERATION"] # Display current REST service
  # your code here
  # if something is wrong, return an HTTP code and optionally set a custom message
  # LET Context["HTTPProcessingErrorMsg"] = "Invalid request" 
  RETURN 0 # HTTP ret code, or 0 to continue standard REST service processing
END FUNCTION

PRIVATE # REST Query string incoming callback
FUNCTION QueryPreCallback(query STRING, value STRING) ATTRIBUTES (WSPreProcessing="wsquery") RETURNS (STRING)
  DISPLAY "Query:",query
  DISPLAY "Value:",value
  # your code here
  RETURN value # new query value
END FUNCTION

PRIVATE # REST Parameter incoming callback
FUNCTION ParamPreCallback(param STRING, value STRING) ATTRIBUTES (WSPreProcessing="wsparam") RETURNS (STRING)
  DISPLAY "Param:",param
  DISPLAY "Value:",value
  # your code here
  RETURN value # new parameter value
END FUNCTION

PRIVATE # REST Header incoming callback
FUNCTION HeaderPreCallback(header STRING, value STRING) ATTRIBUTES (WSPreProcessing="wsheader") RETURNS (STRING)
  DISPLAY "Header:",HEADER
  DISPLAY "Value:",value
  # your code here
  RETURN value # new header value
END FUNCTION

PRIVATE # REST JSONObject incoming callback
FUNCTION JSONObjectPreCallback(obj util.JSONObject) ATTRIBUTES (WSPreProcessing) RETURNS (util.JSONObject)
  DISPLAY "Doing JSON pre processing"
  # your code here
  RETURN obj# new JSONObject value
END FUNCTION

PRIVATE # REST JSONArray incoming callback
FUNCTION JSONArrayPreCallback(obj util.JSONArray) ATTRIBUTES (WSPreProcessing) RETURNS (util.JSONArray)
  # your code here
  RETURN obj # modified JSONArray value
END FUNCTION

PRIVATE # REST XML incoming callback
FUNCTION XMLPreCallback(xml xml.DomDocument) ATTRIBUTES (WSPreProcessing) RETURNS (xml.DomDocument)
  # your code here
  return xml # modified XML value
END FUNCTION

PRIVATE # REST TXT incoming callback
FUNCTION TextPreCallback(s STRING) ATTRIBUTES (WSPreProcessing) RETURNS (STRING)
  # your code here
  RETURN s # modified STRING value
END FUNCTION
```

## Postprocessing callbacks

Use these to process the return header and body of the REST response. They are set with the
`WSPostProcessing`
attribute. The function name does not matter. Names with
`*PostCallback` are conventions for clarity, not used for
selection.

The `HTTPPostCallback` function runs first after the REST operation executes. It
can return `0` to continue processing or an HTTP return code to return to the
client. As with preprocessing, setting `Context["HTTPProcessingErrorMsg"]` when
returning a non-zero code customizes the HTTP reason phrase.

## Examples of postprocessing BDL functions

**Before using these examples**, note that they only display values and return
`0`. In a real implementation, add logic to format or transform the response before
returning it to the client.

```
# REST Web service module
# Examples of postprocessing BDL functions

IMPORT util
IMPORT XML

PRIVATE DEFINE Context DICTIONARY OF STRING

PRIVATE # HTTP response callback
FUNCTION HTTPPostCallback() ATTRIBUTES (WSPostProcessing) RETURNS (INTEGER)
  DISPLAY "Doing HTTP Post processing"
  DISPLAY Context["RequestOPERATION"] # Display current REST service operation
  # your code here
  # return non‑zero to stop and optionally set a custom message:
  # LET Context["HTTPProcessingErrorMsg"] = "Postprocessing failed"
  RETURN 0 # HTTP ret code, or 0 to continue standard REST service processing
END FUNCTION

PRIVATE # REST Header outgoing callback
FUNCTION HeaderPostCallback(header STRING, value STRING) ATTRIBUTES (WSPostProcessing="wsheader") RETURNS (STRING)
  DISPLAY "Header:",HEADER
  DISPLAY "Value:",value
  # your code here
  RETURN value # new header value
END FUNCTION

PRIVATE # REST JSONObject outgoing callback
FUNCTION JSONObjectPostCallback(obj util.JSONObject) ATTRIBUTES (WSPostProcessing) RETURNS (util.JSONObject)
  DISPLAY "Doing JSON Post processing"
  # your code here
  RETURN obj # modified JSONObject value
END FUNCTION

PRIVATE # REST JSONArray outgoing callback
FUNCTION JSONArrayPostCallback(obj util.JSONArray) ATTRIBUTES (WSPostProcessing) RETURNS (util.JSONArray)
  # your code here
  RETURN obj # modified JSONArray value
END FUNCTION

PRIVATE # REST XML outgoing callback
FUNCTION XMLPostCallback(xml xml.DomDocument) ATTRIBUTES (WSPostProcessing) RETURNS (xml.DomDocument)
  # your code here
  RETURN xml # modified XML value
END FUNCTION

PRIVATE # REST TXT outgoing callback
FUNCTION TextPostCallback(s STRING) ATTRIBUTES (WSPostProcessing) RETURNS (STRING)
    # your code here
    RETURN s # modified STRING value
END FUNCTION
```

## Related links

**Related concepts**  

[Using preprocessing and postprocessing callbacks](4751-using-preprocessing-and-postprocessing-callbacks.md "Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.")

[WSPreProcessing](4811-wspreprocessing.md "Specify functions in a REST web service module that are used as preprocessing request handlers.")

[WSPostProcessing](4812-wspostprocessing.md "Specify functions in a REST web service module that are used as postprocessing request handlers.")
