---
title: "Web Services on GMI (iOS)"
source: "fgl-topics/c_gws_mobile_gmi_limitations.html"
breadcrumb: "Web services > General > Platform-specific notes > Web Services on GMI (iOS)"
type: "concept"
---

# Web Services on GMI (iOS)

> Requirements to use Web services on iOS platforms (GMI).

## Web services configuration options

GWS configuration [FGLPROFILE entries related to SSL/TLS keys](4916-fglprofile-entries-for-web-services.md) (`security.*`)
are not supported (uses iOS native SSL/TLS).

## Long running HTTP request

When executing an HTTP request (for example with `com.HttpRequest.doRequest()`),
if the request takes a long time to complete (for example, several minutes), the app will go into
background mode. If the user taps the app icon to return to foreground mode, the program will get
the runtime error -15553.

## The XSLTTransformer class

The [`XSLTTransformer`](../15_library-reference/4342-xslttransformer-methods.md "Methods for the xml.XSLTtransfomer class.")
class is not available on iOS.
