---
title: "GMIDIR"
source: "fgl-topics/c_fgl_EnvVariables_GMIDIR.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > GMIDIR"
type: "concept"
---

# GMIDIR

> Defines the installation directory of Genero Mobile for iOS.

The GMIDIR environment variable defines the installation directory of the Genero Mobile for iOS
archive, used to build iOS apps with the gmibuildtool.

When building apps with gmibuildtool, GMIDIR is used to find the GMI libraries
and resource files.

By default, if GMIDIR is not defined, the location for GMI libraries is found from the location
of the gmibuildtool. For example, if gmibuildtool is found in
/opt/genero/gmi-1.20/bin/gmibuildtool, GMIDIR will be defined as
/opt/genero/gmi-1.20.

## Related links

**Related concepts**  

[Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.")
