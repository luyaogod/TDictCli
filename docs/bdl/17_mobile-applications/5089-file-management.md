---
title: "File management"
source: "fgl-topics/c_fgl_mobile_file_mngt.html"
breadcrumb: "Mobile applications > File management"
type: "concept"
---

# File management

> Specific APIs are available to manipulate file resources in mobile apps.

Mobile devices have media and data files in their local storage unit, that can be
manipulated by Genero apps.

Pictures, videos, and other data files can be selected or created with
Genero BDL APIs.

Files on the mobile device returned by front calls such as
`standard.openFile` are identified with an opaque file identifier. This identifier
can then be passed to APIs such as `fgl_getfile()`, to be stored into a database for
example.

## Child topics

- [Handling files on Android devices](5090-handling-files-on-android-devices.md): How to manipulate file resources with GMA?
- [Handling files on iOS devices](5091-handling-files-on-ios-devices.md): How to manipulate file resources with GMI?
