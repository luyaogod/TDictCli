---
title: "Using preprocessing and postprocessing callbacks"
source: "fgl-topics/c_gws_restful_high_level_preprocessing_postprocessing_handling.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Using preprocessing and postprocessing callbacks"
type: "concept"
---

# Using preprocessing and postprocessing callbacks

> Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.

Preprocessing runs before the REST operation executes, letting you validate or adjust incoming
data (such as query parameters, headers, or body content). For example, you might modify input
values if a database field changed and the query needs updating. Postprocessing runs after the
operation completes, enabling you to transform or format the response before sending it back to the
client—for instance, resizing an image or converting a data format based on content type.

## How to implement preprocessing and postprocessing

Define **private** callback functions in your REST module and associate them with the [WSPreProcessing](4811-wspreprocessing.md "Specify functions in a REST web service module that are used as preprocessing request handlers.") and [WSPostProcessing](4812-wspostprocessing.md "Specify functions in a REST web service module that are used as postprocessing request handlers.") attributes:

- The function name does not matter. Names like `HeaderPreCallback` are just
  conventions for clarity, not used for selection.
- When choosing which function to call, the runtime looks at the data type of the value being
  processed and calls the function whose parameter type matches that data type. For example:
  - A `STRING` value uses a callback with `(s STRING) RETURNS
    (STRING)`.
  - A JSON object uses a callback with `(obj util.JSONObject) RETURNS
    (util.JSONObject)`.
  - An XML document uses `(xml xml.DomDocument) RETURNS (xml.DomDocument)`.This type-based selection is similar to method overloading in Java.
- Callbacks must return the same kind of value they receive.
- For sample callback functions, go to [Examples of callback functions](4752-examples-of-callback-functions.md "These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.").

Callbacks run in the order shown below, from HTTP preprocessing to response body
postprocessing:

| Order | Function | Description | Callback type |
| --- | --- | --- | --- |
| 1 | `HTTPPreCallback()` | Runs first. Returns 0 to continue or an HTTP code to stop immediately. | General |
| 2 | `QueryPreCallback(query, value)` | Validate or adjust query parameter values before the operation runs. | WSQuery |
| 3 | `ParamPreCallback(param, value)` | Validate or adjust parameter values before the operation runs. | WSParam |
| 4 | `HeaderPreCallback(header, value)` | Rename or validate header names and values before the operation runs. | WSHeader |
| 5 | `JSONObjectPreCallback(obj)` `JSONArrayPreCallback(a)` `XMLPreCallback(xml)` `TextPreCallback(s)` | Adjust the request body format. The callback is selected based on body content type. | General (body) |
| 6 | REST operation | Execute the standard REST operation. | N/A |
| 7 | `HTTPPostCallback()` | Runs first after the operation. Change the HTTP return code if needed. | General |
| 8 | `HeaderPostCallback(header, value)` | Rename outgoing header names and values before sending the response. | WSHeader |
| 9 | `JSONObjectPostCallback(obj)` `JSONArrayPostCallback(a)` `XMLPostCallback(xml)` `TextPostCallback(s)` | Adjust the response body format. The callback is selected based on body content type. | General (body) |

## Related links

**Related concepts**  

[WSPreProcessing](4811-wspreprocessing.md "Specify functions in a REST web service module that are used as preprocessing request handlers.")

[WSPostProcessing](4812-wspostprocessing.md "Specify functions in a REST web service module that are used as postprocessing request handlers.")

## Child topics

- [Examples of callback functions](4752-examples-of-callback-functions.md): These examples show the structure of preprocessing and postprocessing callback functions you can use in your REST web service module.
