---
title: "FGLAPPDIR"
source: "fgl-topics/c_fgl_EnvVariables_FGLAPPDIR.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLAPPDIR"
type: "concept"
---

# FGLAPPDIR

> Contains the path to the application directory when executing on a mobile device.

When executing on mobile devices, the FGLAPPDIR environment variable is an automatic
environment variable that contains the path to the appdir directory,
containing application program files (.42m, .42f,
and other resources).

This variable is typically used to define environment variables with
`mobile.environment` [FGLPROFILE](0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
entries, relative to the mobile appdir where application program files and resources are
located.

During development (when executing programs on a server), consider defining the FGLAPPDIR in the
shell environment, along with the other environment variables that are defined with
`mobile.environment` entries, as these are only read when executing on mobile
devices.

## Related links

**Related concepts**  

[Directory structure for GMI apps](../17_mobile-applications/5108-directory-structure-for-gmi-apps.md "Platform-specific rules need to be considered when deploying on iOS devices (GMI).")

[Directory structure for GMA apps](../17_mobile-applications/5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).")
