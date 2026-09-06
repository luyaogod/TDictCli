---
title: "User-defined front calls"
source: "fgl-topics/c_fgl_frontcalls_user_ext.html"
breadcrumb: "Extending the language > User-defined front calls"
type: "concept"
---

# User-defined front calls

> Front-ends can be extended with custom functions to access specific features.

It is possible to implement custom front-end functions to interface with platform-specific
features, and use the feature from a Genero program through a front call. For example, you can
implement a front-end function module interfacing with a bar code reader, to return bar codes
to the Genero program.

This section describes how to implement your own front calls by front-end type. Because each
front-end type uses different technologies, you must use native platform APIs to implement front
calls.

## Child topics

- [Implement front call modules for GDC](2719-implement-front-call-modules-for-gdc.md): Custom front call modules for the desktop front-end are implemented by using the API for GDC front calls in C language.
- [Implement front call modules for GMA](2720-implement-front-call-modules-for-gma.md): Custom front call modules for the Android™ front-end are implemented by using the API for GMA front calls in Java.
- [Implement front call modules for GMI](2721-implement-front-call-modules-for-gmi.md): Custom front call modules for the iOS front-end are implemented by using the API for GMI front calls in Objective-C.
- [Implement front call modules for GBC](2722-implement-front-call-modules-for-gbc.md): Custom front call modules for the Genero Browser Client (GBC) front-end are implemented by using JavaScript.
