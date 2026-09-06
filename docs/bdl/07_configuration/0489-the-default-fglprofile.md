---
title: "The default FGLPROFILE"
source: "fgl-topics/c_fgl_fglprofile_007.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > The default FGLPROFILE"
type: "concept"
---

# The default FGLPROFILE

> The Genero BDL package ships a default FGLPROFILE file as $FGLDIR/etc/fglprofile.

Do NOT modify the $FGLDIR/etc/fglprofile default configuration file: This
file will be overwritten by a new installation and your changes will be lost.

Make a copy of the default profile file, change the entries in your private configuration file,
and specify this private file in the [FGLPROFILE](0530-fglprofile.md "Defines the configuration files to be used by the runtime system.") environment variable.
