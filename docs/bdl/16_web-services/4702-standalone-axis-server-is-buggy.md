---
title: "Standalone Axis server is buggy"
source: "fgl-topics/c_gws_i4gl_migration_guide_020.html"
breadcrumb: "Web services > SOAP Web Services > How Do I ... ? > How to migrate I4GL web service to Genero > Migrate an I4GL web service consumer to Genero > Standalone Axis server is buggy"
type: "concept"
---

# Standalone Axis server is buggy

> Describes a bug you can expect using the I4GL standalone axis server.

The I4GL standalone axis server adds an extra CR LF after the body of the SOAP HTTP post response
what leads the Genero client to return the error message : `Body content bigger than
expected`. This is not allowed as defined in HTTP [[RFC2616]](http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.1).

> **Important:**
>
> Axis works as expected if loaded from Apache server.
