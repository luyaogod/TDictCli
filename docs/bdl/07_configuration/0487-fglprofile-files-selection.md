---
title: "FGLPROFILE files selection"
source: "fgl-topics/c_fgl_fglprofile_005.html"
breadcrumb: "Configuration > The FGLPROFILE file(s) > FGLPROFILE files selection"
type: "concept"
---

# FGLPROFILE files selection

> When multiple FGLPROFILE files can be selected, the runtime system applies an order of precedence.

There are three levels of FGLPROFILE configuration files, and these files are loaded in the
following order:

1. The runtime system reads the default configuration file provided in [$FGLDIR](0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")/etc/fglprofile. This file contains all
   supported entries, identifies the possible values for an entry, and documents default values.
   > **Important:**
   >
   > Do not modify the default configuration file FGLDIR/etc/fglprofile, as it
   > can be overwritten by a new installation.
2. If the [FGLPROFILE](0530-fglprofile.md "Defines the configuration files to be used by the runtime system.") environment variable is
   set, the runtime system reads entries from the files specified by this environment variable. A list
   of files can be provided with FGLPROFILE. Files must be separated by the operating system specific
   path separator.

   On mobile devices,
   you must deploy a file with the name "`fglprofile`" in the appdir
   directory. See [FGLPROFILE for mobile apps](0490-fglprofile-for-mobile-apps.md "The name of the FGLPROFILE file matters for mobile applications.") for more details.
3. After loading and merging the two previous levels, the runtime system checks whether the
   `fglrun.defaults` entry is set. This entry defines the *program-specific profile
   directory*. If this directory contains a file with the same name as the current program (without
   a .42r or .42m extension), the runtime system reads the
   entries from that file.

   If the value of `fglrun.defaults` starts with
   `$envvar`, the specified environment variable is expanded to its
   value in order to build the path. For example, if the value is `"$MYVAR/standard"`,
   and MYVAR contains /opt/app/config, and the program name is
   custinfo.42m (or custinfo.42r), the runtime system will
   try to read the file /opt/app/config/standard/custinfo.

   Make sure that the directory used for `fglrun.defaults` is not the same directory
   where program files are located: If a directory contains the files "myprog.42m"
   and "myprog", an "fglrun myprog" command will try to load
   the "myprog" configuration file, instead of the
   "myprog.42m" file and produce the error [-6202](../15_library-reference/4483-genero-bdl-errors.md).

The runtime system merges the different configuration files found at the three levels: If the
same entry is defined in several files, the last loaded entry wins.

This means that the order of precedence is:

1. Program-specific configuration file (if `fglrun.defaults` is defined in one of
   the other levels).
2. Configuration files defined by the FGLPROFILE environment variable, or
   appdir/fglprofile, for mobile applications.
3. The default configuration file $FGLDIR/etc/fglprofile.
