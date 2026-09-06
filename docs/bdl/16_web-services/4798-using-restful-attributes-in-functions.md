---
title: "Using RESTful attributes in functions"
source: "fgl-topics/c_gws_high_level_rest_using_restful_attributes.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > Using RESTful attributes in functions"
type: "concept"
---

# Using RESTful attributes in functions

> RESTful attributes define functions for your RESTful web service.

You declare certain things about functions in your Web service program which identifies the
function as a high level RESTful web service and helps the compiler check your code accordingly.

The keyword `ATTRIBUTE` or `ATTRIBUTES` allows you to specify
special RESTful attributes (identified with `WS*` prefix) when declaring a function.
This keyword is followed by an attribute specification inside parentheses.

If you have a function you want to publish as a RESTful function of a Web service, you implement
the function in the following way:

**Function attributes**

You must specify the HTTP operation, such as GET, POST, PUT, or DELETE, you want to perform on
the resource in the function's attribute property. You must also specify the path to the Web service
resource using the `WSPath` RESTful attribute here. This is a minimum requirement for
a RESTful function. See the prototype function example in [HTTP verbs and attributes](4799-http-verbs-and-attributes.md "HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.").

**Input parameters**

If the resource needs data to be passed in as queries, cookies, or headers in order to perform
the operation, you must declare input parameters and specify the required attribute in the
attributes property. Or if data must be sent in the body of the request, you specify an input
parameter to receive the data. This parameter does not require any special RESTful attribute. See
the prototype function example in [Input parameter attributes](4800-input-parameter-attributes.md "A RESTful web service function can have input parameters.").

**Return values**

A RESTful web service function must include an explicit
`RETURNS()` clause, even if the function returns no value — use an empty
`RETURNS()` to indicate that; otherwise, you get runtime error [-9158](../15_library-reference/4483-genero-bdl-errors.md). This enforces consistent function signatures and prevents ambiguous web service
definitions.

You handle returns in the function as you would in a normal function, except when you
want to send data in the header response. Then you must specify this in the returns value attribute
property. If the data must go in the body of the response, you specify a return value to receive the
data. This return value does not require any special RESTful attribute. See the prototype function
example in [Return value attributes](4801-return-value-attributes.md "A RESTful web service function can have return values.").

## Related links

**Related concepts**  

[REST function syntax with RESTful attributes](4797-rest-function-syntax-with-restful-attributes.md "A RESTful FUNCTION definition is specified with a set of input parameter attributes, function definition attributes, and return attributes that define it as an operation for a REST web service.")

## Child topics

- [HTTP verbs and attributes](4799-http-verbs-and-attributes.md): HTTP verbs are defined by the high-level RESTful attributes. Some verbs have requirements for request or response body and others do not.
- [Input parameter attributes](4800-input-parameter-attributes.md): A RESTful web service function can have input parameters.
- [Return value attributes](4801-return-value-attributes.md): A RESTful web service function can have return values.
