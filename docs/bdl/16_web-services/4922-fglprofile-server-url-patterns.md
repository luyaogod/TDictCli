---
title: "FGLPROFILE: Server URL patterns"
source: "fgl-topics/c_gws_ssl_config_url_wildcard.html"
breadcrumb: "Web services > Reference > Web services FGLPROFILE configuration > Examples > FGLPROFILE: Server URL patterns"
type: "concept"
---

# FGLPROFILE: Server URL patterns

> FGLPROFILE entries can be used to define multiple server URLs, by using URL patterns.

## Understanding URL patterns

By using URL patterns, you can create a URL base that applies to multiple server
applications.

URLs matching the pattern can share server configuration (such as authentication and HTTPS).

FGLPROFILE entries using URL patterns apply to the following APIs:

- [`com.HttpRequest.Create()`](../15_library-reference/3857-com-httprequest-create.md "Creates a new HttpRequest object from a URL.")
- [`xml.DomDocument.load()`](../15_library-reference/3996-xml-domdocument-load.md "Loads a XML Document into a DomDocument object from a file or an URL.")
- [`xml.DomDocument.save()`](../15_library-reference/4002-xml-domdocument-save.md "Saves this xml.DomDocument object as a XML Document to a file or URL.")
- [`xml.StaxReader`](../15_library-reference/4121-the-staxreader-class.md "The StaxReader class provides methods compatible with Streaming API for XML(StAX) for reading XML documents.") methods
- [`xml.StaxWriter`](../15_library-reference/4090-the-staxwriter-class.md "The xml.StaxWriter class provides methods compatible with Streaming API for XML(StAX) for writing XML documents.") methods
- [`com.TcpRequest.create()`](../15_library-reference/3934-com-tcprequest-create.md "Creates a new TCP request object.")

## Using \* wildcards in server URLs

To create a URL base, add a wildcard (`/*`) to the end of a URL in the
FGLPROFILE file entry. A server application that starts with this URL (and that
is not explicitly defined elsewhere) shares the configuration with other applications that also
start with the same base URL. If an application has its own server configuration explicitly defined,
it uses its specific entries instead of those defined by the wildcard configuration.

Consider this excerpt from a hypothetical FGLPROFILE
file:

```
authenticate.auth.login     = "xxx"
authenticate.auth.password  = "yyy"
authenticate.auth.scheme    = "Basic"

security.sec.certificate    = "client.crt"
security.sec.privatekey     = "client.pem"

ws.myapp.url                = "http://mycompany.com/sample/*"
ws.myapp.authenticate       = "auth"
ws.myapp.security           = "sec"

ws.thirdapp.url             = "http://mycompany.com/sample/application3"
ws.thirdapp.authenticate    = "auth3"

authenticate.auth3.login    = "aaa"
authenticate.auth3.password = "bbb"
authenticate.auth3.scheme   = "Basic"
```

Based on this example:

- Requests to "`http://mycompany.com/sample/application1`" and
  `"http://mycompany.com/sample/demos/shoppingcart`" use the same authentication and
  HTTPS configuration.
- A request to "`http://mycompany.com/sample/application3`" uses its specific
  authentication "`auth3`". No security configuration is defined for this URL, nor does
  it fall back on the shared security configuration defined for the base URL.

## Using regular expressions in server URLs

To create a URL regex pattern, use the `ws.idws.regex.url` FGLPROFILE file entry, and define a regular expression that matches all server
URLs to be associated with the "idws" server configuration.

If the `ws.idws.url` is defined, the `regex.url` entry is ignored.

```
authenticate.auth.login    =  "xxx"
authenticate.auth.password =  "yyy"
authenticate.auth.scheme   =  "Basic"

ws.myident.regex.url = "http://.*\.strasbourg\.4js\.com:[0-9]{4}/.*"
ws.myident.authenticate =  "auth"
```

## Related links

**Related reference**  

[Server configuration](4916-fglprofile-entries-for-web-services.md "Server configuration")
