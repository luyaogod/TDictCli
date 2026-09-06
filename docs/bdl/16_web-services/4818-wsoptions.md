---
title: "WSOptions"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSOptions.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSOptions"
type: "concept"
---

# WSOptions

> In order to retrieve the HTTP request methods supported by the server or the resource, you define the WSOptions attribute.

## Syntax

```
WSOptions
```

## Usage

You use this attribute to specify the action of the HTTP verb OPTIONS to return the
methods allowed by the server or resource. You set the `WSOptions`
attribute in the `ATTRIBUTES()` clause of the function.

An output message body is not allowed in the response, so returns
must be specified as headers with the [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute.

## Example using WSOptions to return the methods allowed for a resource

In this example the REST function returns the supported methods in the
`Access-Control-Allow-Methods` header.

```
PUBLIC FUNCTION OptionsUsers() 
   ATTRIBUTES(WSOptions, 
             WSPath = '/users')
   RETURNS(STRING ATTRIBUTES(WSHeader, 
                            WSName = "Access-Control-Allow-Methods")) 
   RETURN "GET, HEAD, OPTIONS, TRACE"
END FUNCTION
```

## Example using WSOptions to return the methods allowed for Web service

In this example the function returns the methods allowed for the entire Web
service. The client calls the function using an asterisk (`*`) in the
path.

http://myhost/gas/ws/r/myGroup/myXcf/MyService/\*

```
PUBLIC FUNCTION OptionsServer() 
   ATTRIBUTES(WSOptions, WSPath = '/*')
   RETURNS STRING ATTRIBUTES(WSHeader, 
           WSName = "Access-Control-Allow-Methods")
   RETURN "GET, PUT, POST, DELETE, PATCH, HEAD, OPTIONS, TRACE"
END FUNCTION
```
