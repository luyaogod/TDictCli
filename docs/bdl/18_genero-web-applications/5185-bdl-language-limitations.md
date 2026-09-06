---
title: "BDL language limitations"
source: "fgl-topics/c_fgl_gwa_bdl_limitations.html"
breadcrumb: "Genero Web applications > BDL language limitations"
type: "concept"
---

# BDL language limitations

> Genero language features not supported or with limited support on GWA.

> **Important:**
>
> This topic is provided as a reference for Genero Business Development Language limitations in GWA
> applications.

## Features with limited support

You can embed a Genero application as web application with the following constraints:

| Feature | Comment |
| --- | --- |
| Web services using the [RESTful Web services (high-level framework)](../16_web-services/4705-restful-web-services-high-level-framework.md "These topics give you the information you need to begin working with RESTful Web services applications using BDL function with support for attributes.") are supported. | When using [fglrestful](../13_programming-tools/2524-fglrestful.md "The fglrestful tool produces REST web services stub files for client programs using an OpenAPI specification.") to generate REST stub files:Only the JSON API [The util.JSON class](../15_library-reference/3579-the-util-json-class.md "The util.JSON class provides a basic interface to convert program variable values to/from JSON data.") is supported. Therefore, you must use option `--legacyJSONApi` of the fglrestful tool.Only requests with JSON or plain text are supported.For XML handling, use the `om.DomDocument` API. |
| Web services using the SOAP protocol are not supported. | No additional information. |
| GWS REST low-level API is partially supported. | The GWA supports a minimum subset of Genero APIs for creating web services. The `COM` API used is based on the JavaScript [`XMLHttpRequest`](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest) (external link) object (excluding the specific XML API). |
| `RUN cmd` and `RUN cmd WITHOUT WAITING` instructions are supported. | There are some limitations because the GWA runs in the browser. For details, go to [Executing sub-programs from a parent program with RUN](5141-executing-programs-with-run.md "Sub-programs can be executed from a parent program with the RUN instruction.") |
| Database support. | Only the [SQLite](../10_sql-support/1466-sqlite.md) database is supported at this time. |
| Some methods in the [The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.") are not fully working. | In the [The Channel class](../15_library-reference/2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions."), the following socket and pipe methods are not supported and will return an error:[base.Channel.openClientSocket](../15_library-reference/2989-base-channel-openclientsocket.md "Open a TCP client socket channel.")[base.Channel.openPipe](../15_library-reference/2991-base-channel-openpipe.md "Opens a pipe channel to a subprocess.")[base.Channel.openServerSocket](../15_library-reference/2992-base-channel-openserversocket.md "Open a TCP server socket channel.")The [base.Channel.openFile](../15_library-reference/2990-base-channel-openfile.md "Opens a file channel.") is supported |

## Unsupported features

The following language features are not supported:

- All the XML classes in the [The xml package](../15_library-reference/3957-the-xml-package.md "The Genero Web Services XML package provides classes and methods to handle any kind of XML documents, including documents with namespaces.") (referenced with
  `IMPORT xml`), including the [The Encryption class](../15_library-reference/4309-the-encryption-class.md "The xml.Encryption class provides methods to encrypt and decrypt XML documents, nodes or symmetric keys."), are not supported.
  > **Note:**
  >
  > For XML handling, use the `om.DomDocument` API.
- The security classes in the [The security package](../15_library-reference/4407-the-security-package.md "The Genero Web Services security package provides classes and methods to support basic cryptographic features.") (referenced with
  `IMPORT security`) are not supported.
- The JSON streaming classes in the [The json package](../15_library-reference/3616-the-json-package.md "The Genero Web Services JSON package provides classes and methods to process JSON documents.") (referenced with
  `IMPORT json`) are not supported.
