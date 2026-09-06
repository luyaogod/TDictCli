---
title: "Input parameter attributes"
source: "fgl-topics/c_gws_high_level_rest_input_parameters.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > Using RESTful attributes in functions > Input parameter attributes"
type: "concept"
---

# Input parameter attributes

> A RESTful web service function can have input parameters.

The type of input parameter is defined by the `ATTRIBUTES` property. Table 1 lists the types.
Additional parameters can be specified to provide a name or description; and depending on the
parameter type, other attributes might identify the parameter as optional or specify a supported
data format.

```
PUBLIC FUNCTION function-name(
    parameter-name INTEGER ATTRIBUTES(WSHeader, WSName = "id"),
    parameter-name INTEGER ATTRIBUTES(WSParam, WSDescription = "Value to be used")
    # ...
    )
    # ...
END FUNCTION
```

| Parameter type | Description | Other attributes allowed |
| --- | --- | --- |
| [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") | The parameter holds a template path. | [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") |
| [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") | The parameter holds a custom header. | [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")[WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") |
| [WSQuery](4836-wsquery.md "Maps a parameter to a query string value in the request URL.") | The parameter holds a query path in the resource URL. | [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")[WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") |
| [WSCookie](4835-wscookie.md "Maps a parameter to a cookie value in the HTTP request.") | The parameter holds a cookie. | [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")[WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") |
| Request body (without attributes `WSHeader`, `WSQuery`, `WSCookie`, or `WSParam`) | The parameter is used for transferring data in the request body. Data can be JSON, XML, text, or a file type (such as an image).Data of different MIME types can be sent in multiple parts of the same request. See [Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body."). | [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.")[WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")[WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") |
| [WSAttachment](4839-wsattachment.md "Defines file attachments in the REST message.") | The parameter holds the path to a file for attachment. | [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.")[WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") |

## Related links

**Related concepts**  

[Define your resource operations](4724-define-your-resource-operations.md "Operations are the HTTP verbs used to manipulate the resources of your Web service.")

[Set a request body](4731-set-a-request-body.md "Functions that create or update a resource need to set a request body for the incoming payload. You specify the request body in an input parameter.")
