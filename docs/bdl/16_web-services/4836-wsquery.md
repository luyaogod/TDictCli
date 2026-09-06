---
title: "WSQuery"
source: "fgl-topics/c_gws_high_level_rest_api_attributes_WSQuery.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > High-level RESTful Web service attributes > Attributes set on parameters and returns > Attributes set on input parameters only > WSQuery"
type: "concept"
---

# WSQuery

> Maps a parameter to a query string value in the request URL.

## Syntax

```
WSQuery
```

This attribute supports a query in the URL. Zero, one, or several parameters can be specified.

`WSQuery` is an optional attribute.

## Usage

You set the `WSQuery`
attribute in the `ATTRIBUTES()` clause of an input parameter. It is typically used to
filter collections. The client provides query parameter values in the URL, for example:

Http://myhost/gas/ws/r/myGroup/myXcf/MyService/users?firstname=john

## WSQuery on records and arrays

`WSQuery` can be applied to parameters defined as primitive types, records, or
arrays.

The GWS and fglrestful serialize and deserialize query parameters according to
the OpenAPI defaults.

`WSQuery` supports:

- `style: form`
- `explode: true`

## OpenAPI serialization rules

The following table shows how the query string is serialized depending on parameter type:

| style | explode | Primitive value id = 5 | Array id = [3, 4, 5] | Record = {"role": "admin", "lastname": "Smith"} |
| --- | --- | --- | --- | --- |
| form | true | /users?id=5 | /users?id=3&id=4&id=5 | /users?role=admin&lastname=Smith |

For further information on OpenAPI serialization, see the [Parameter
serialization](http://swagger.io/docs/specification/serialization/) page.

The GWS and fglrestful supports
serialization and deserialization of parameters according to defaults set by the OpenAPI
specification. If a value contains special characters as defined by [RFC3986](http://tools.ietf.org/html/rfc3986#section-2.2) (for example, `:/?#[]@!$&'()*+,;=`), the GWS engine
serializes them with percent-encoding in the parameter. For
example:

```
DEFINE b DYNAMIC ARRAY OF STRING = ["O,ne", "Two", "Three", NULL, "Five"]
```

The parameter contains:
`O%2Cne,Two,Three,,Five`. The server and the client must deserialize
"`O%2Cne`" to "`O,ne`" for insertion in Genero
BDL.

## Optional parameters

If you do not need to query every field in a record, add `WSOptional`. All fields
become optional.

See the example Example WSQuery set as optional on a parameter defined as record

## Example WSQuery set as optional on parameters

This example uses
`WSQuery` and `WSOptional` on three input parameters.

The client
may supply any combination of query
parameters:

http://myhost/gas/ws/r/myGroup/myXcf/MyService/accounts?firstname=john&lastname=Smith

or:

http://myhost/gas/ws/r/myGroup/myXcf/MyService/accounts?lastname=Smith

The
SQL query uses `COALESCE` to support optional parameters. If no parameter is
provided, all accounts are returned.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

TYPE accountType RECORD
         userid VARCHAR(80),
         firstname VARCHAR(80),
         lastname VARCHAR(80),
         email VARCHAR(80)
   END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getAccountRecords(
    resourceId STRING ATTRIBUTES(WSQuery, WSOptional, WSName = "id"),
    fname STRING ATTRIBUTES(WSQuery, WSOptional, WSName = "firstname"),
    lname STRING ATTRIBUTES(WSQuery, WSOptional, WSName = "lastname"))
  ATTRIBUTES (WSGet,
              WSPath = "/accounts",
              WSDescription = "Fetches the accounts resource with the optional filter value(s) applied.",
        WSScope = "officestore.user",
        WSThrows = "404:Not Found,406:@userError,500:Internal Server Error")
  RETURNS (DYNAMIC ARRAY ATTRIBUTES(WSName = "accounts") OF
        accountType ATTRIBUTES(XMLName = "account", WSMedia = "application/json, application/xml"))
  
    DEFINE recList DYNAMIC ARRAY OF accountType
    DEFINE i INTEGER = 1
     TRY
       DECLARE c5 CURSOR FOR SELECT * FROM account
                               WHERE account.userid = COALESCE(resourceId,account.userid)
                               AND account.firstname = COALESCE(fname,account.firstname)
                               AND account.lastname = COALESCE(lname,account.lastname)
                               ORDER BY account.userid ASC           
       # COALESCE function is used to produce a query that supports the optional parameters
       FOREACH c5 INTO recList[i].*
         LET i = i+1
       END FOREACH
       CALL recList.deleteElement(recList.getLength())
       # Remove the empty element implied by reference in FOREACH loop 
     CATCH
       LET userError.message = SFMT("Error in SQL execution: %1 [%2]", sqlca.sqlcode, SQLERRMESSAGE )
       CALL com.WebServiceEngine.SetRestError(406,userError)
     END TRY
    RETURN recList
END FUNCTION
```

## Example WSQuery set as optional on a parameter defined as record

In this example, the parameter `accRec` is a record with optional fields. The
client may provide values for any field. For example:

http://myhost/gas/ws/r/myGroup/myXcf/MyService/accounts?category=admin&lastname=Smith

The GWS deserializes the fields into the Genero BDL record. The SQL uses
`COALESCE` to support optional values.

[WSThrows](4829-wsthrows.md "Defines the list of error codes the REST function may return.") is set to handle errors. In the `TRY/CATCH` block, the `sqlca` record is checked after
the execution of the SQL query. The `SQLERRMESSAGE` is set to the
`message` field of the `userError` variable, and a call to
`SetRestError()` returns the message defined in `WSThrows`
for the error.

```
IMPORT com

TYPE accountType RECORD
         userid VARCHAR(80),
         firstname VARCHAR(80),
         lastname VARCHAR(80),
         email VARCHAR(80),
         category VARCHAR(10)
  END RECORD

PUBLIC DEFINE userError RECORD ATTRIBUTE(WSError = "User error")
  message STRING
END RECORD

PUBLIC FUNCTION getAccountRecord(accRec
  RECORD ATTRIBUTES(WSQuery, WSOptional) category STRING, lastname STRING END RECORD  )
  ATTRIBUTES (WSGet, 
              WSPath = "/accounts/rec",
              WSDescription = "Fetches account records for query value(s) optionally supplied.",
        WSScope = "officestore.user",
        WSThrows = "404:Not Found,406:@userError,500:Internal Server Error")
  RETURNS (DYNAMIC ARRAY ATTRIBUTES(WSName = "Users") OF accountType )
  
    DEFINE recList DYNAMIC ARRAY OF accountType
    DEFINE i INTEGER = 1
      TRY
        DECLARE c8 CURSOR FOR SELECT * FROM users
                               WHERE users.category = COALESCE(accRec.category,users.category)
                               AND users.lname = COALESCE(accRec.lastname,users.lname)
                               ORDER BY users.id ASC                     
        # COALESCE function is used to produce a query that supports the optional parameters
        FOREACH c8 INTO recList[i].*
          LET i = i+1
        END FOREACH
        CALL recList.deleteElement(recList.getLength())
        # Remove the empty element implied by reference in FOREACH loop    
      CATCH
        LET userError.message = SFMT("Error in SQL execution: %1 [%2]", sqlca.sqlcode, SQLERRMESSAGE )
        CALL com.WebServiceEngine.SetRestError(406,userError)
      END TRY
    RETURN recList
END FUNCTION
```
