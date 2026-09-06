---
title: "Retrieve HTTP headers"
source: "fgl-topics/c_gws_restful_high_level_handling_http_headers.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Code a RESTful server application > Define functions in a module > Retrieve HTTP headers"
type: "concept"
---

# Retrieve HTTP headers

> You can retrieve HTTP headers in your REST operation. There are two methods for doing this.

You can retrieve the header using the [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") attribute as part of a function
parameter, or you can retrieve headers via [WSContext](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.").

The option you choose depends on your requirements:

- If you want the header to be part of the REST operation to implement logic on the server or
  client, you can pass data via headers you customize using [WSHeader](4838-wsheader.md "Defines a custom HTTP header for a parameter or return value.") and [WSName](4844-wsname.md "Specifies an alternative name for a parameter or return value in the REST message.") attributes in your function
  parameter. For example, the remote address can be passed in a customized HTTP header by
  specifying an input parameter with these
  attributes:

  ```
  ip_address STRING ATTRIBUTES(WSHeader, WSOptional, 
                                WSName = "X-FourJs-Environment-Variable-REMOTE_ADDR")
  ```

  You
  can display the remote address in your function using the statement `DISPLAY
  ip_address`. In case the header is not present, you must also set the [WSOptional](4845-wsoptional.md "Marks a parameter or return value as optional in the REST message.") attribute.
- If you do not want the header to be part of the REST operation, you can use the [WSContext](4806-wscontext.md "Provides a dictionary of request‑specific context values for the current REST request, available to all operations in the module.") mechanism. For example, the remote
  address and server name can be retrieved by setting a context dictionary variable at the modular
  level with the `WSContext` attribute. This allows you to retrieve all
  `X-FourJs-Environment-xxx` set by the GAS by referencing a
  dictionary key value for the environment variable.

  In the following example, the dictionary type
  variable (context) is set with the attribute `WSContext` :

  ```
  PRIVATE DEFINE context DICTIONARY ATTRIBUTES(WSContext) OF STRING

  DISPLAY context["Variable-REMOTE_ADDR"] -- displays the remote ADDR if set by the GAS
  DISPLAY context["Variable-SERVER_NAME"] -- displays the server name if set by the GAS
  ```

## Related links

**Related concepts**  

[Set query, header, or cookie parameters](4730-set-query-header-or-cookie-parameters.md "Define the WSQuery, WSHeader, or WSCookie parameters in your function if the resource needs data passed as a query, cookie, or header.")
