---
title: "WSPath"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSPath.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set at the function level > WSPath"
type: "concept"
---

# WSPath

> Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.

## Syntax

```
WSPath = "/{ path-element | value-template } [/...]"
```

where path-element is an
identifier:

```
identifier
```

and where value-template is an identifier enclosed in curly
brackets:

```
{identifier}
```

1. path-element is a slash-separated list of path elements.
2. value-template elements are place holders for a slash-separated list of
   variable values.

`WSPath` is an optional attribute.

If there is no `WSPath` specified, the function name (case sensitive) is the
path to the resource in the URI.

## Usage

You use this attribute to set the resource identifier or endpoint of the URI to the resource. For
example, a path to the "users" URI resource is specified as follows:

```
WSPath="/users"
```

The `WSPath` string begins with a slash (`/`), and there is a
slash between each path element and/or template value.

If you need to specify a resource within a collection, such as a specific user identified by ID,
and when a pattern could match many similar resources, it is specified with a path
template.

```
WSPath=/users/{id}
```

The path template is an identifier enclosed in curly brackets `{}`. Zero, one, or
several path templates can be specified. Values are substituted when a GWS client makes a call to
the resource.

Path template values are provided at runtime in function parameters which have `WSParam` attributes. As many template
values as path templates must be specified. In other words there must be a function parameter with a
`WSParam` attribute to match each template path. fglcomp checks to
ensure that there is one template value per parameter, otherwise compilation [error-9111](../15_library-reference/4483-genero-bdl-errors.md) is thrown.

You set the `WSPath` attribute in the `ATTRIBUTES()` clause of the
function.

## Example path templating with WSParam

In this sample REST function a user resource is returned. In the function's
`p_user_id` parameter the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the user to
return.

The `WSPath` attribute sets the resource identifier or endpoint of the URI
to the resource. It contains the function's `p_user_id` parameter enclosed in
curly brackets `{}` as the path template.

The client application needs to provide an appropriate parameter value when making a call
to the function. For example, the variable part of the path (`p_user_id`) is
replaced by the integer value `/4` in the URL:

http://myhost/gas/ws/r/myGroup/myXcf/MyService/accounts/4

```
TYPE accountType RECORD
         user_id INTEGER,
         user_name VARCHAR(50)
END RECORD

PUBLIC FUNCTION getAccountById(p_user_id INTEGER ATTRIBUTES(WSParam))
  ATTRIBUTES(WSGet,
             WSPath = "/accounts/{p_user_id}",
             WSDescription = "Returns an account record",
             WSThrows = "404:not found"
            )
  RETURNS accountType ATTRIBUTES(WSName = "Account",WSMedia = "application/json") 
    DEFINE p_accountRec accountType
    
    SELECT * INTO p_accountRec.* FROM users WHERE id = p_user_id

    RETURN p_accountRec
END FUNCTION
```

## Related links

**Related concepts**  

[Set resource path with WSParam and WSPath](4729-set-resource-path-with-wsparam-and-wspath.md "Path parameters allow you to specify variables in the resource URL.")

[Set query, header, or cookie parameters](4730-set-query-header-or-cookie-parameters.md "Define the WSQuery, WSHeader, or WSCookie parameters in your function if the resource needs data passed as a query, cookie, or header.")
