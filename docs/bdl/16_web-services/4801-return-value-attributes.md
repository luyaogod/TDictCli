---
title: "Return value attributes"
source: "fgl-topics/c_gws_high_level_rest_return_values.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > Using RESTful attributes in functions > Return value attributes"
type: "concept"
---

# Return value attributes

> A RESTful web service function can have return values.

The type of return values is defined by the `ATTRIBUTES` property. Table 1 lists the return types.
Additional parameters can be specified to provide a name or description; and depending on the
parameter type, other attributes might identify the parameter as optional or specify a supported
data format.

A RESTful web service function must include an explicit
`RETURNS()` clause, even if the function returns no value — use an empty
`RETURNS()` to indicate that; otherwise, you get runtime error [-9158](../15_library-reference/4483-genero-bdl-errors.md)

```
PUBLIC FUNCTION function-name(
    #...
    parameter-name INTEGER ATTRIBUTES(WSParam),
  
    # ...
    )
    ATTRIBUTES(
       # function attributes ...
       )
    RETURNS ( INTEGER ATTRIBUTES(WSHeader),
              STRING ATTRIBUTES (WSDescription = "Greeting")
             )
        # function code
        RETURN 3, "Hello world"
END FUNCTION
```

| Return type | Description | Other attributes allowed |
| --- | --- | --- |
| [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") | The return value contains a custom header. | [WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")[WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") |
| [WSErrorHeader](4831-wserrorheader.md "Defines a custom HTTP header to include in an error response.") | The return value contains a header with an HTTP error response code. | [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")[WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.") |
| [WSAttachment](4839-wsattachment.md "Defines file attachments in the REST message.") | The return value contains a path to a file for attachment. | [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.")[WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") |
| Response body (without `WSHeader`) | Denotes the return value is used for transferring data in the response body. Data can be JSON, XML, text, or a file type (like an image).Data of different MIME types can be sent in multiple parts of the same response. See [Multipart requests or responses](4743-multipart-requests-or-responses.md "In GWS REST there is support for the standard multiple part message, in which more than one different sets of data are combined in a single body.") | [WSMedia](4841-wsmedia.md "Defines the supported media (MIME) types for a parameter or return value.")[WSDescription](4842-wsdescription.md "Describes the REST function, parameters, return values, and members of user-defined types.")[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.")If the return is serialized in XML, attributes that [customize the serialization](5021-xml-serialization-attributes.md), such as `XMLName`, may also be used. |

## Related links

**Related concepts**  

[Define your resource operations](4724-define-your-resource-operations.md "Operations are the HTTP verbs used to manipulate the resources of your Web service.")

[Set a response body and header](4732-set-a-response-body-and-header.md "You specify a response body in a return parameter without an attribute. Other return values can be sent in headers, using the WSHeader attribute.")
