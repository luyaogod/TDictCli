---
title: "REST function syntax with RESTful attributes"
source: "fgl-topics/c_gws_high_level_rest_function_definition_syntax.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > REST function syntax with RESTful attributes"
type: "concept"
---

# REST function syntax with RESTful attributes

> A RESTful FUNCTION definition is specified with a set of input parameter attributes, function definition attributes, and return attributes that define it as an operation for a REST web service.

## Syntax (RESTful Function ):

```
[PUBLIC|PRIVATE] FUNCTION function-name (
     { parameter-name type-specification [ attributes-list ]
     | record-name record-type INOUT
     }
     [,...]
  )
 [ attributes-list ]
 [ RETURNS returns-specification ]
    [ local-declaration [...] ]
    [ instruction 
    | [ RETURN expression [,...] ]
        [...] 
    ]
END FUNCTION
```

> **Note:**
>
> High-level RESTful functions require a `RETURNS` clause; otherwise, you get
> runtime error [-9158](../15_library-reference/4483-genero-bdl-errors.md).

In type specifications, the attributes-list clause is:

```
{ ATTRIBUTE | ATTRIBUTES } ( attribute [ = "value" ] [,...] )
```

1. attribute is the name of a definition
   attribute.
2. value is the value for the definition attribute, it is
   optional for boolean attributes.

where attributes-list of function parameters must be:

```
ATTRIBUTES (
{ WSParam
| WSHeader
| WSQuery
| WSCookie
| WSAttachment
| WSDescription = "parameter-description"
| WSOptional
| WSName ="alt-name"
| WSMedia = " MIME-type [,...] "
[,...]
} )
```

where attributes-list of the function must
be:

```
ATTRIBUTES ( http-verb [ , rest-function-attribute [,...] ] )
```

where http-verb must be one
of:

```
{ WSGet | WSPost | WSPut | WSDelete | WSHead
| WSOptions | WSTrace | WSPatch
}
```

where rest-function-attribute can be one
of:

```
{ WSPath = "/{ path-element | value-template } [/...]"
| WSDescription = "function-description"
| WSSummary = "function-summary"
| WSThrows = " error-definition [,...] "
| WSRetCode = "success-code:description"
| WSScope = "security-scope [,...] "
| WSVersion = "version-name [,...] "
| WSTags = "tag-name [,...]"
| WSOperationID = "id-name"
}
```

where error-definition can
be:

```
{ error-code
| error-code:description
| error-code:@variable
}
```

where returns-specification
is:

```
{ type-specification [ attributes-list ]
| ( type-specification [ attributes-list ] [,...] )
| ( )
}
```

where attributes-list of returns-specification must
be:

```
ATTRIBUTES (
{ WSHeader
| WSErrorHeader = "error-code [,...]"
| WSAttachment
| WSDescription = "return-description"
| WSName ="alt-name"
| WSMedia = " MIME-type [,...] "
[,...]
} )
```

1. path-element is a slash-separated list of path elements.
2. value-template elements are place holders for a slash-separated list of
   variable values.
3. alt-name is an alternative name for parameters or return values in a REST
   function that improves readability.
4. MIME-type is a supported data format type, such as TEXT, JSON, XML, or an
   image format such JPEG, or PNG, and so on.

## Function with attributes

Function attributes can be used to complete the definition of the function itself, for its
parameters and return values:

```
PUBLIC FUNCTION add(
    p1 INTEGER ATTRIBUTES(WSHeader, WSName = "id"),
    p2 INTEGER ATTRIBUTES(WSParam)
    )
    ATTRIBUTES(
       WSGet,
       WSPath = '/{p2}/hello',
       WSThrows="404:not found, 500:Internal Server Error",
       WSSummary = "Return customer ID",
       WSDescription="Returns an integer and a greeting",
       WSTags = "Customer"
       )
     RETURNS (INTEGER ATTRIBUTES(WSHeader), STRING ATTRIBUTES (WSDescription="Greeting") )
        DEFINE var1 INTEGER
        LET var1 = (p1 + p2)
        RETURN var1, "Hello world"
END FUNCTION
```

For more details see [Function attributes](../08_language-basics/0769-function-attributes.md "Function attributes can be used to add definition information about the function, its parameters and its return values."). See [Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.") for examples of use of the attributes
in arguments.

## Related links

**Related concepts**  

[Functions](../08_language-basics/0761-functions.md "Describes user defined functions.")

[Variables](../08_language-basics/0686-variables.md "Explains how to define program variables.")

[Expressions](../08_language-basics/0592-expressions.md "Shows the possible expressions supported in the language.")
