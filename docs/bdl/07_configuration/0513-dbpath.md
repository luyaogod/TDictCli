---
title: "DBPATH"
source: "fgl-topics/c_fgl_EnvVariables_DBPATH.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBPATH"
type: "concept"
---

# DBPATH

> Defines a list of paths for Genero program resource files.

For IBM® Informix®
4GL compatibility, DBPATH is used by the runtime system to find resource files such as form
definitions.

> **Important:**
>
> The DBPATH environment variable is also used by the IBM Informix SE engine and SQLite,
> to define the path list to find database files. Genero has introduced the FGLRESOURCEPATH
> environment variable to not interfere with the database DBPATH settings. Consider dedicating DBPATH
> for database configurations, and use the FGLRESOURCEPATH to define program resource path
> list.

DBPATH must contain a list of paths, separated by the operating system specific path separator.

The path separator is platform specific ( **":"** on UNIX™
platforms and **";"** on Windows® platforms).

See [FGLRESOURCEPATH](0531-fglresourcepath.md "Defines a list of paths for program resource files.") for more details
about resource files search path.
