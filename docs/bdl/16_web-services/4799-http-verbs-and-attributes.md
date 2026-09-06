---
title: "HTTP verbs and attributes"
source: "fgl-topics/c_gws_high_level_rest_verbs_summary.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > Using RESTful attributes in functions > HTTP verbs and attributes"
type: "concept"
---

# HTTP verbs and attributes

> HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.

Table 1 lists the
HTTP verbs and their corresponding RESTful high-level API attributes. The Request body required
column identifies those verbs that require an input message body, and the Response body allowed
column identifies those verbs that do not allow an output message body.

| HTTP Verb | REST API attribute | Request body required? | Response body allowed? |
| --- | --- | --- | --- |
| GET | [WSGet](4816-wsget.md "In order to retrieve a resource, you set the WSGet attribute.") | No | Yes |
| HEAD | [WSHead](4817-wshead.md "In order to retrieve resource headers, you define the WSHead attribute.") | No | No |
| POST | [WSPost](4820-wspost.md "In order to create a new resource, you set the WSPost attribute.") | Optional | Yes |
| PUT | [WSPut](4821-wsput.md "Update an existing resource with the WSPut attribute.") | Optional | Yes |
| DELETE | [WSDelete](4815-wsdelete.md "In order to remove an existing resource, you set the WSDelete attribute.") | No | Yes |
| OPTIONS | [WSOptions](4818-wsoptions.md "In order to retrieve the HTTP request methods supported by the server or the resource, you define the WSOptions attribute.") | No | No |
| TRACE | [WSTrace](4822-wstrace.md "You define the WSTrace attribute for debugging purposes.") | No | No |
| PATCH | [WSPatch](4819-wspatch.md "In order to partially update an existing resource, you define the WSPatch attribute.") | Optional | No |

In summary:

- POST, PUT, and PATCH requests are verbs that typically require an **input** message body, and
  failing to specify a message body in a request will raise [error-9106](../15_library-reference/4483-genero-bdl-errors.md) or [error-9128](../15_library-reference/4483-genero-bdl-errors.md). You can, however, define
  the input body parameter as optional by setting the [WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") attribute on the input parameter. With
  the body parameter defined as optional, an input body request is not required when a client calls
  the service.
- HEAD, TRACE, and PATCH responses do not allow an **output** message body. Specifying these
  responses with a message body will raise [error-9129](../15_library-reference/4483-genero-bdl-errors.md).

When you compile your Web service application, the fglcomp tool checks that
attributes comply with these requirements.

## Function prototype example

A function's attributes property determine the HTTP verb, the resource path, and more. Table 2 lists and
describes the attributes you can use in the `ATTRIBUTES` clause of the function.

```
PUBLIC FUNCTION function-name(
    #...
    parameter-name INTEGER ATTRIBUTES(WSParam),
    # ...
    )
    ATTRIBUTES(
       WSGet,
       WSPath = '/{parameter-name}/hello',
       WSThrows = "404:not found, 500:Internal Server Error",
       WSDescription = "Returns an integer and a greeting"
       WSVersion = "v2, v3"
       # ...
       )
     RETURNS ( INTEGER, STRING )
     # ... function code
    
END FUNCTION
```

| Attribute type | Description |
| --- | --- |
| [HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.") | An HTTP verb action specifying the operation on a REST resource. |
| [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") | Slash-separated list of path elements and/or value templates identifying the resource URL. |
| [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") | Textual description of the function. |
| [WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") | Colon-separated list of error-definitions as HTTP status code, and/or code with description. |
| [WSRetCode](4830-wsretcode.md "Sets the HTTP success status code returned by the REST function.") | Colon-separated HTTP status success code with description. |
| [WSScope (function)](4824-wsscope-function.md "Specifies the security scope required to access this REST function. Use this attribute to restrict access at the function level.") | A comma-separated list of security scopes giving access to the resource. |
| [WSVersion (function)](4825-wsversion-function.md "Specifies the version or versions in which the REST function is available.") | A comma-separated list of version names specifying the versions of the Web service in which the operation will be accessible to clients. |

## Related links

**Related concepts**  

[Define your resource operations](4724-define-your-resource-operations.md "Operations are the HTTP verbs used to manipulate the resources of your Web service.")

[Set a request body](4731-set-a-request-body.md "Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.")

[Set a response body and header](4732-set-a-response-body-and-header.md "You specify a response body in a return parameter without an attribute. Other return values can be sent in headers, using the WSHeader attribute.")
