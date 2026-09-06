---
title: "WSTrace"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSTrace.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSTrace"
type: "concept"
---

# WSTrace

> You define the WSTrace attribute for debugging purposes.

## Syntax

```
WSTrace
```

## Usage

You use this attribute to specify the action of the HTTP TRACE verb for debugging purposes. You
set the `WSTrace` attribute in the `ATTRIBUTES()` clause of the
function.

An output message body is not allowed in the response, so returns
must be specified as headers with the [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute.

## Example using WSTrace

In this example the REST function returns a trace.
This can help identify bugs in your code. For example, you might want to check if you can reach the
service. The client calls the function using an asterisk (`*`) in the
path.

http://myhost/gas/ws/r/myGroup/myXcf/MyService/\*

```
PUBLIC FUNCTION TraceServer()
   ATTRIBUTES(WSTrace, 
             WSPath = '/*')
   RETURNS(STRING ATTRIBUTES(WSHeader, 
                            WSName = "Debug"))

   RETURN "You've reached the service"
 END FUNCTION
```
