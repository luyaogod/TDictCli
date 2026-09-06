---
title: "Genero Web Services extension"
source: "fgl-topics/c_gws_concepts_009.html"
breadcrumb: "Web services > General > Introduction to Web services > Genero Web Services extension"
type: "concept"
---

# Genero Web Services extension

> Applications providing Web services use special libraries of the Genero Business Development Language.

The Genero Web Services Extension (GWS) is an extension to the Genero
Business Development Language. It installs within the Genero Business Development Language
directory. The fglgws package includes both Genero Business Development Language and Genero Web
Services.

The Genero Application Server is required to manage your Web services in
a deployment environment. It is not required for Web services development, unless you are interested
in testing deployment issues.

When programming a Web service, your applications must include `IMPORT com` at the
top of each module. This imports the Genero Web Services Extension library named `com`:

```
IMPORT com
```
