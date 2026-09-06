---
title: "Customize XML serialization with WSName and XMLName"
source: "fgl-topics/c_gws_restful_high_level_customize_xml_serialization.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Set data format with WSMedia > Customize XML data output"
type: "concept"
---

# Customize XML serialization with WSName and XMLName

> Use XML serialization attributes to customize serialization at runtime and improve the readability of the XML output.

## Formatting XML output with WSName and XMLName

This sample REST function returns a list of users to the client in XML format. The
function's return value defines a dynamic array of `profileType` that is set
with the following attributes:

- `WSName="All_users_list"` sets the root node for the output to
  `<All_users_list>`.
- `WSMedia = "application/xml"` sets the output to XML.
- `XMLName="User"` maps each element of the `profileType`
  record with the `<User>` tag.
- The `XMLOptional` attribute specifies that `<User>`
  elements with NULL values are omitted from the XML document. For other options of
  working with NULL values, see the [XMLNillable](5023-xmlnillable.md)
  attribute.

If you do not use the `WSName` and `XMLName`
attributes to format the XML output, names derived from DVM names such as "rv0", "rv1", and
so on are used.

```
TYPE profileType RECORD
     id INTEGER,
     name VARCHAR(100),
     email VARCHAR(255),
     ccode VARCHAR(3)
     # ...
   END RECORD

PUBLIC FUNCTION getUsersList()
  ATTRIBUTES (WSGet,
              WSPath = "/users/xml",
              WSDescription="Get list of users and output in XML format"
              )
  RETURNS (
    DYNAMIC ARRAY ATTRIBUTES(WSName = "All_users_list",
                            WSMedia = "application/xml") OF profileType
                            ATTRIBUTES(XMLName = "User", XMLOptional)
          )

    DEFINE usersList DYNAMIC ARRAY OF profileType
    DEFINE i INTEGER = 1

    # code to get users
     DECLARE userCurs CURSOR FOR SELECT * FROM users ORDER BY users.name
     FOREACH userCurs INTO arr[i].*
       LET i = i+1
     END FOREACH
     CALL arr.deleteElement(arr.getLength())
     # Remove the empty element implied by reference in FOREACH loop

  RETURN usersList
END FUNCTION
```

The output in the XML document contains the tags specified by the attributes
[WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") and [XMLName](5037-xmlname.md): If the output is viewed in an XML viewer, it
would display as shown.

```
<All_users_list>
  <User>
    ... record content
  </User>
  <User>
    ... record content
  </User>
  ...
</All_users_list>
```

## Related links

**Related concepts**  

[Serialize XML from a dynamic array](5046-serialize-xml-from-a-dynamic-array.md "The XMLList and XMLName attributes can be used to serialize dynamic arrays to XML and vice versa.")
