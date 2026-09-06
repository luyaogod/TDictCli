---
title: "Language limitations"
source: "fgl-topics/c_fgl_mobile_bdl_limitations.html"
breadcrumb: "Mobile applications > Language limitations"
type: "concept"
---

# Language limitations

> Genero language features not supported on mobile devices.

> **Important:**
>
> This topic is provided as a quick glance at Genero Business Development
> Language limitations in mobile applications. Details can be found in the BDL reference
> topics.

## Features with limited support

The following language options have limited support:

- When running an embedded app on a mobile device (or emulator), the
  `RUN cmd` and `RUN cmd WITHOUT
  WAITING` instructions are not supported. To implement apps running on a server, see the
  [`runOnServer`](../15_library-reference/3458-mobile-runonserver.md "Run an application from the Genero Application Server using the specified URL.") front
  call.

## Unsupported features

The following language features are not supported:

- The `base.Channel.openPipe` method is not supported.
- The Java interface cannot be used in apps running on iOS devices: There is
  no standard free JVM available.
