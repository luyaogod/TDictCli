---
title: "Understanding FGLPROFILE"
source: "fgl-topics/c_fgl_fglprofile_002.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > Understanding FGLPROFILE"
type: "concept"
---

# Understanding FGLPROFILE

> The runtime system uses one or more configuration files in which you can define options and parameters to change the behavior of the programs.

The FGLPROFILE files define standard BDL or user-defined entries with a name and value.

Standard entries can be used to control the runtime system behavior, and user-defined entries can
be defined to configure your application (prefix user-defined entries with the name of your
application to avoid conflicts with standard BDL FGLPROFILE entries)

Multiple profile files can be specified in the [FGLPROFILE](0530-fglprofile.md "Defines the configuration files to be used by the runtime system.") environment variable.

On mobile devices,
you must deploy a file with the name "`fglprofile`" in the appdir
directory. See [FGLPROFILE for mobile apps](0490-fglprofile-for-mobile-apps.md "The name of the FGLPROFILE file matters for mobile applications.") for more details.

The list of standard core language FGLPROFILE entries can be found in [FGLPROFILE entries for core language](0486-fglprofile-entries-for-core-language.md "This is a summary of FGLPROFILE entries supported by the core BDL.").
