---
title: "Web Services support"
source: "fgl-topics/c_fgl_MigI4GL_029.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Web Services support"
type: "concept"
---

# Web Services support

> Both I4GL and Genero BDL support Web services, albeit with different implementations.

Starting with IBM® Informix® 4GL version **7.50**,
I4GL functions can be deployed as Web Services. The published functions
can be subscribed from programs that run on a Web client in another
programming language.

Web Services support was introduced in Genero Business Development
Language before I4GL 7.50 was released. Each implementation is
quite different, but the basic principles are the same: publishing
4gl functions as Web Services, by handling WS requests and supporting
easy input and output parameter conversions between WS data formats
and 4gl program variables.

## Related links

**Related concepts**  

[Web services](../16_web-services/4484-web-services.md "Create a web service client or server with Genero BDL.")
