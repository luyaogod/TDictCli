---
title: "FGLPROFILE for mobile apps"
source: "fgl-topics/c_fgl_fglprofile_008.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > FGLPROFILE for mobile apps"
type: "concept"
---

# FGLPROFILE for mobile apps

> The name of the FGLPROFILE file matters for mobile applications.

For non-mobile apps, there is no specific naming convention for FGLPROFILE configuration files.
You can use a filename without an extension, or use the .txt or
.prf extensions.

On mobile devices, it is not possible to define environment variables.

To specify a custom FGLPROFILE file for a mobile application, you must deploy a file with the
name "`fglprofile`" in the [appdir
directory](../09_advanced-features/0829-executing-programs.md "There are different ways to execute compiled programs, depending on the configuration and the development or production context."), along with the other application program files (.42m, .42f, and so on).

Only one custom FGLPROFILE file can be deployed for a given mobile application.

For more details, see [Deploying mobile apps](../17_mobile-applications/5102-deploying-mobile-apps.md "This section describes how to build and deploy mobile apps with Genero.").
