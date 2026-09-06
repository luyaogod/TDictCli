---
title: "Troubleshooting no matching REST operation"
source: "fgl-topics/c_gws_rest_client_troubleshooting.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful client application > Create the client application > Troubleshooting"
type: "concept"
---

# Troubleshooting no matching REST operation

> Understanding how the GWS finds an operation to use for a request can help in troubleshooting unexpected errors.

A common error you may encounter is that caused by the GWS not being able to find a matching
operation for a request. An unsupported MIME type or other issue may cause this.

## Criteria for REST operation search

The GWS scans the list of operations of a REST service for 4 criteria to find a match for the
function to use for a request:

- HTTP PATH
- HTTP VERB
- HTTP Content-Type header (if any) - corresponding to the `WSMedia` set for the input parameter of
  the operation
- HTTP Accept header (if any) - corresponding to the `WSMedia` set for the return parameter of
  the operation

> **Important:**
>
> All 4 criteria (including the HTTP headers if specified) are required to
> match the operation.

For example, your request to get a list of customers might specify the media type of the response
as XML. However, using the equivalent to this curl command, the request
fails:

```
curl -H "Accept: application/xml" http://localhost:8090/ws/r/MyService/customers
```

GWS returns this error:

```
"400 No matching Rest operation found"
```

Assuming
in this case the path (/customers) and verb (GET) is ok, the GWS did not find a
match to any operation because the operation's output type does not support XML. With reference to
the curl command above, try the request without specifying the media type of the
response to get the default output:

```
curl http://localhost:8090/ws/r/MyService/customers
```

JSON is the default media type used for input and output in REST operations from GWS 4.00
onwards.

It is recommended that you refer to the Web service OpenAPI description for details of the media
types allowed by operations. For information on getting the OpenAPI description, go to [Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.").

## Related links

**Related concepts**  

[Information needed to generate a REST stub file](4781-before-generating-the-stub-file.md "You need the OpenAPI description of the RESTful service before generating the stub file or implementing client-side calls.")

[Handling application level errors](4746-handling-application-level-errors.md "There are many situations in which you need to notify an error to a client using your REST web service.")
