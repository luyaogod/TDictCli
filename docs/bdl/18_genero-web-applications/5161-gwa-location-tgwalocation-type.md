---
title: "gwa.location.TGWALocation type"
source: "fgl-topics/c_gwa_api_type_TGWALocation.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.location module > gwa.location module > gwa.location.TGWALocation type"
type: "concept"
---

# gwa.location.TGWALocation type

> The TGWALocation type defines a record for a Genero wrapper around the JavaScript window.location object.

## Syntax

```
TYPE TGWALocation RECORD
   href STRING
   protocol STRING
   host STRING
   hostname STRING
   port INTEGER
   pathname STRING
   search STRING
   hash STRING
   origin STRING
   approot STRING
END RECORD
```

1. The `href` string holds the entire URL. For example,
   https://myserver.org:6394/gwa/gwa\_dist?test=true#bang
2. The `protocol` string holds the protocol scheme of the URL, including the final
   colon (`":"`). For example, "HTTP:" or "HTTPS:".
3. The `host` string holds the host, that is the hostname, a colon (
   `":"`), and the port of the URL. For example,
   myserver.org:6394
4. The `hostname` string holds the domain of the URL. For example,
   myserver.org.
5. The `port` string holds the port number of the URL.
6. The `pathname` string holds the initial slash (`'/'`) followed by
   the path of the URL, but not including the query string or fragment. For example,
   /gwa/gwa\_dist
7. The `search` string holds a question mark (`"?"`) followed by the
   parameters or "query string" of the URL. For example, ?test=true
8. The `hash` string holds a hash or pound sign (`"#"`) followed by
   the fragment identifier of the URL. For example, #bang
9. The `origin` string holds the canonical form of the origin of the specific
   location.
10. The `approot` holds the web server path where the GWA application was started.
    For example, if the application was launched at
    https://myserver.org:6394/gwa/gwa\_dist/index.html, the
    approot value is
    https://myserver.org:6394/gwa/gwa\_dist.

## Usage

It provides an API to inspect a location URL of an application window. A variable of the type
`TGWALocation` must be defined.

For more information, refer to the [HTML
DOM API location](https://developer.mozilla.org/en-US/docs/Web/API/Location) (external link) documentation.
