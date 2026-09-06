---
title: "WSHead"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSHead.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > HTTP operation attributes (Verbs) > WSHead"
type: "concept"
---

# WSHead

> In order to retrieve resource headers, you define the WSHead attribute.

## Syntax

```
WSHead
```

## Usage

You use this attribute to specify the action of the HTTP verb HEAD to return resource
headers. An HTTP HEAD request can be made before getting a resource to check resource
size, validity, accessibility, and recent modification time. For example, if a resource
might return a large file download, you would use a HEAD request to read its
`Content-Length` header to check the filesize without actually
downloading the file.

You set the `WSHead` attribute in the `ATTRIBUTES()`
clause of the function.

An output message body is not allowed in the response, so returns
must be specified as headers with the [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute.

## Example using WSHead to return resource headers

In this example the REST function returns the resource headers for the
`users` resource. The `Accept-Language` header informs
the client that the response will be in JSON.

```
PUBLIC FUNCTION HeadUsers() 
   ATTRIBUTES(WSHead,
              WSPath = "/users",
              WSDescription = "Returns header of users",
              WSThrows = "400:Invalid,404:not found")
   RETURNS(STRING ATTRIBUTES(WSHeader,
                            WSName = "Accept-Language",
                            WSMedia = "application/json") )
   RETURN "Data returned in json"
END FUNCTION
```
