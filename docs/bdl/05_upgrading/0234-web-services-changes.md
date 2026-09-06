---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_250_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.50 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 2.50.

## Client stubs managing multipart changes on client side

> **Important:**
>
> This change has been backported from version 3.00

Starting with version 2.50.25, when generating client stubs managing multipart, you get an extra
input and/or output variable called "`AnyInputParts`" and
"`AnyOutputParts`" that is a `DYNAMIC ARRAY` of
`com.HTTPPart` objects. Those variables may contain additional input and/or output
HTTP parts not specified in the WSDL. You will have to adapt your client program by handling those
dynamic arrays in any Genero functions calling such stubs.

Request example prior to 2.50.25:

```
FUNCTION xxx_g(InputHttpPart_1, ..., InputHttpPart_n)
  DEFINE InputHttpPart_1 com.HTTPPart
  ...
  DEFINE InputHttpPart_n com.HTTPPart
  ...
  RETURN wsstatus
END FUNCTION
```

Request example 2.50.25 and greater, with extra input variable
`AnyInputParts`:

```
FUNCTION xxx_g(InputHttpPart_1, ..., InputHttpPart_n)
  DEFINE InputHttpPart_1 com.HTTPPart
  ...
  DEFINE InputHttpPart_n com.HTTPPart
  DEFINE AnyInputParts DYNAMIC ARRAY OF com.HTTPPart
  ...
  RETURN wsstatus
END FUNCTION
```

## Web Services deprecated APIs

The methods listed in the following table are deprecated in version 2.50.

| Method deprecated as of 2.50 | Alternative method to use |
| --- | --- |
| `com.Util.CreateDigestString` | `security.Digest.CreateDigestString` |
| `com.Util.CreateRandomString` | `security.RandomGenerator.CreateRandomString` |
| `com.Util.CreateUUIDString` | `security.RandomGenerator.CreateUUIDString` |

## Related links

**Related concepts**  

[The HttpPart class](../15_library-reference/3910-the-httppart-class.md "The com.HttpPart class provides an interface to manage the HTTP attachment sent or received in HTTP.")

[Multipart in the client stub](../16_web-services/4626-multipart-in-the-client-stub.md "You can generate a client stub for a Web service that has multiple parts.")

[Dynamic arrays](../08_language-basics/0734-dynamic-arrays.md "Dynamic arrays")
