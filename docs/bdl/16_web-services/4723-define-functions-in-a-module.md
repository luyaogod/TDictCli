---
title: "Define functions in a module"
source: "fgl-topics/c_gws_restful_high_level_server_module.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module"
type: "concept"
---

# Define functions in a module

> A GWS REST service is defined in a module.

You write a Genero BDL function for each service function you wish to make public. You identify
functions as RESTful functions by using [RESTful attributes](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.") in the [function attributes](../08_language-basics/0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values.") clause of the function and its
parameter and return values.

- HTTP [operation attributes (Verbs)](4814-http-operation-attributes-verbs.md "Attributes that map an HTTP operation or verb action in a function to a REST resource.")
  define the function operation (such as the HTTP GET, or POST operation).
- Resource path [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") parameters (with
  `WSParam`, `WSPath` attributes) define the location of a resource. The
  path parameter value is directly related to a resource and is part of the resources URL path.
- Attributes set on [Input
  parameters and return values](4832-attributes-set-on-parameters-and-returns.md "Attributes that affect the input parameters and return values of the function.")  define data sent in the request and response. For example, a
  parameter's RESTful attribute can specify if data is carried in the body or header, as a cookie, as
  a query, or as multipart.

## Ensure resources are unique

As function names have to be unique in your module, the Web service resource identified by
functions must also be unique. In the OpenAPI specification the combination of resource path and an
HTTP verb defines a unique resource.

The same resource path can support several operations, for example, `GET /users`
to get a list of users and `POST /users` to add a new user, etc. However, only one
operation is allowed per function and therefore you must provide separate functions for these
operations.

Two GET or two POST functions for the same path are not allowed; even if they have different
parameters. Input parameters or return values do not effect the uniqueness of the resource.

See the examples for defining your REST Web service functions.

## Related links

**Related concepts**  

[RESTful web services](4703-restful-web-services.md "Create RESTful Web service applications (server and/or client) with Genero Web Services. RESTful Web services conform to the REST architectural style.")

**Related reference**  

[High-level RESTful Web service attributes](4802-high-level-restful-web-service-attributes.md "Attributes for high-level RESTful Genero Web Services.")

## Child topics

- [Define your resource operations](4724-define-your-resource-operations.md): Operations are the HTTP verbs used to manipulate the resources of your Web service.
- [Set resource path with WSParam and WSPath](4729-set-resource-path-with-wsparam-and-wspath.md): Path parameters allow you to specify variables in the resource URL.
- [Set query, header, or cookie parameters](4730-set-query-header-or-cookie-parameters.md): Define the WSQuery, WSHeader, or WSCookie parameters in your function if the resource needs data passed as a query, cookie, or header.
- [Set a request body](4731-set-a-request-body.md): Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.
- [Set a response body and header](4732-set-a-response-body-and-header.md): You specify a response body in a return parameter without an attribute. Other return values can be sent in headers, using the WSHeader attribute.
- [Retrieve HTTP headers](4733-retrieve-http-headers.md): You can retrieve HTTP headers in your REST operation. There are two methods for doing this.
- [Handling file attachments with REST](4734-handling-file-attachments-and-data-transfer.md): The Genero REST high-level framework provides two mechanisms for handling attachments.
- [Set data format with WSMedia](4741-set-data-format-with-wsmedia.md): It is important to set the correct MIME type for a Web service request or response. You can specify the data format via the WSMedia attribute.
- [Multipart requests or responses](4743-multipart-requests-or-responses.md): In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.
- [Handling application level errors](4746-handling-application-level-errors.md): There are many situations in which you need to notify an error to a client using your REST web service.
- [Handling security](4749-handling-security.md): You handle security in GWS high-level REST via scopes.
- [Using preprocessing and postprocessing callbacks](4751-using-preprocessing-and-postprocessing-callbacks.md): Preprocessing and postprocessing in Genero REST web services allow you to customize how requests and responses are handled.
