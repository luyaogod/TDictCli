---
title: "Set resource path with WSParam and WSPath"
source: "fgl-topics/c_gws_restful_high_level_path_template_example.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Set resource path with WSParam and WSPath"
type: "concept"
---

# Set resource path with WSParam and WSPath

> Path parameters allow you to specify variables in the resource URL.

In the OpenAPI specification, path templates are the parts of the path that are
replaceable with parameters. If you need to point to a specific resource within a collection (such
as a user identified by ID) and when a pattern could match many similar resources, you use path
templates in your REST function.

The OpenAPI specification denotes path templates within curly braces `{
}`.

```
/users/{id}
```

The path template `{id}` is substituted with actual values when a GWS client makes
a call to the resource. In this example, the URI provides the user id
"22":

```
http://myhost/gas/ws/r/myGroup/myXcf/Account/users/22
```

You make a request to the resource with the URI. A function in your web service with the above
path pattern processes the request.

The REST attributes [WSPath](4823-wspath.md "Specifies a path to a REST web service resource that identifies its function and allows parameters to be passed in the URL.") and [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") set in the attributes clause of the BDL
function support path templating. `WSPath` may have many templates. For each input
parameter with a `WSParam` attribute, there must be a matching template in the
`WSPath` attribute value, otherwise compilation [error-9111](../15_library-reference/4483-genero-bdl-errors.md) is thrown.

## Example: access a resource from path set with templates

In this sample function details of when a specified book was checked out by a library user
are returned.

In the function's membersid parameter the attribute [WSParam](4834-wsparam.md "Maps a parameter to a value in the resource path template.") specifies the library member
resource to return, and in its booksid parameter it specifies the book
resource to return.

The path to the resource is set with two templates; `{membersid}` and
`{booksid}`. These are set in the function attributes clause
`WSPath`
attribute:

```
ATTRIBUTES (WSGet,WSPath = "/members/{membersid}/books/{booksid}")
```

In an example endpoint URI to request the resoure, the value "48" representing the member
id and the value "3" representing the book id, replace the variable parts of the
path:

```
http://myhost/gas/ws/r/myGroup/myXcf/MyLibrary/members/48/books/3
```

The function's output parameter bookList variable of type
`checkoutType` returns the details.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

TYPE checkoutType RECORD
           booksId INT,
           title VARCHAR(100),
           author VARCHAR(100),
           m_name VARCHAR(100),
           checkout_date DATETIME YEAR TO SECOND
        END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTES(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getBooksCheckedOut(membersid INTEGER ATTRIBUTES(WSParam),
  booksid INTEGER ATTRIBUTES(WSParam) )
  ATTRIBUTES (WSGet,
              WSPath = "/members/{membersid}/books/{booksid}",
              WSDescription = "Get details when a library member checked out a specific book",
              WSThrows = "400:Invalid,406:@userError" )
 
   RETURNS (
    DYNAMIC ARRAY ATTRIBUTES(WSName = "Books_checked_Out") OF checkoutType )
    DEFINE bookList DYNAMIC ARRAY OF checkoutType
    DEFINE i INTEGER = 1
    TRY
      DECLARE c1 CURSOR FOR SELECT b.booksid, b.title, b.author, m.name, c.checkoutdate
                              FROM books b
                              INNER JOIN checkouts c ON c.booksid = b.booksid
                              INNER JOIN members m ON m.membersid = c.membersid
                              WHERE c.booksid = booksid AND c.membersid = membersid

        FOREACH c1 INTO bookList[i].*
          LET i = i+1
        END FOREACH
        CALL bookList.deleteElement(bookList.getLength())
        # Remove the empty element implied by reference in FOREACH loop
    CATCH
       LET userError.message = SFMT("Error in SQL execution: %1 [%2]", 
                                     sqlca.sqlcode, SQLERRMESSAGE)
       CALL com.WebServiceEngine.SetRestError(406,userError)
    END TRY
    RETURN bookList
END FUNCTION
```

## When to use WSParam or WSQuery

Know when to use `WSParam` or `WSQuery` attributes in your Web
service.

- Use a path parameter (with `WSParam`) to return a single entity given its
  `id` field.

  For example, `/users/{id}` specifies the entity in the
  resource path. If the entity specified by the `{id}` path parameter is not found, you
  typically get a "`404 (Not found)`" HTTP response code.
- Use a query parameter (with `WSQuery`) for any field other than a
  resource `id`. A number of entities may be returned in the response.

  With a search
  or filter, you expect to return a number of entities in the response, unlike the single entity that
  you expect to get as the response for sending `id` as path parameter.
  `WSQuery` is the logical attribute to use for filtering, searching, or sorting
  resources.

  For example, in `/users?lastname=smith`, the search criteria is
  specified in a query parameter of the URL. If no entities are found matching the search or filter
  parameters, you typically get a "`200 (OK)`" HTTP response code with an empty array
  as the response body.

## Related links

**Related concepts**  

[Using RESTful attributes in functions](4798-using-restful-attributes-in-functions.md "RESTful attributes define functions for your RESTful web service.")

[REST resource URI naming practice](4754-resource-uri-naming.md "The URI of your REST Web service using the high-level API; check that the best practice in URI naming conventions are being followed.")
