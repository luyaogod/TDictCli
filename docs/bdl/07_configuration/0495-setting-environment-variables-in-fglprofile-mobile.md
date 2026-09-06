---
title: "Setting environment variables in FGLPROFILE (mobile)"
source: "fgl-topics/c_fgl_EnvVariables_in_fglprofile.html"
breadcrumb: "Configuration > Environment variables > Setting environment variables in FGLPROFILE (mobile)"
type: "concept"
description: "When executing applications on mobile devices, you can configure environment settings with FGLPROFILE entries. Setting an environment variable with an FGLPROFILE entry is equivalent to setting the ..."
---

# Setting environment variables in FGLPROFILE (mobile)

When executing applications on mobile devices, you can configure environment settings with
FGLPROFILE entries. Setting an environment variable with an FGLPROFILE entry is equivalent to
setting the environment variable before running the fglrun VM process on a
server.

Environment variables set in an FGLPROFILE file are only read when the deployed
application runs on the mobile device. They are not read during development mode (that is when the
VM runs on the development machine and the mobile client displays on the device). The FGLPROFILE
environment variable settings are only for the VM component and are ignored by the GMA/GMI front-end
component.

FGLPROFILE environment variables settings can be used to define [DBDATE](0508-dbdate.md "Defines the default display and input format for DATE values.") and [DBFORMAT](0511-dbformat.md "Defines the characters to be used for the currency symbol, decimal and thousands separators for numeric values."), if the default regional settings on the
mobile must be ignored for date and numeric value formatting. Note that defining [DBMONEY](0512-dbmoney.md "Defines the characters to be used for the currency symbol and decimal separator for numeric values, when DBFORMAT is not defined.") will have no effect, because DBFORMAT is
defined automatically by the GMI or GMA front-end component before starting the VM component. Since
DBFORMAT takes precedence over DBMONEY, setting DBMONEY in FGLPROFILE is pointless.

> **Important:**
>
> C-runtime library variables such as [LANG/LC\_ALL](0497-lc-all-or-lang.md "Defines the current application locale on UNIX platforms.") cannot be set with FGLPROFILE entries,
> because the C-runtime library is (and must be) initialized before reading FGLPROFILE files.

The syntax is:

```
mobile.environment.env_name = "env_value"
```

where:

1. env\_name is the name of the environment variable to be set.
2. env\_value is the value for the env\_name environment
   variable.

For example:

```
mobile.environment.MY_ENV_VAR = "my value"
```

The value specified in a `mobile.environment` entry can contain
`$NAME` placeholders, that will be replaced by the actual value of
the `NAME` environment variable. The
`NAME` environment variable will typically be set by the front-end
component, before starting the runtime system component, for example to define [FGLDIR](0523-fgldir.md "Defines the installation directory of Genero Business Development Language.") and [FGLAPPDIR](0519-fglappdir.md "Contains the path to the application directory when executing on a mobile device.") values.

If the environment variable contains directory or file paths, use the UNIX path notation with `/` slashes as directory
name separator, and the `:` colon as path separator.

This example defines the [FGLIMAGEPATH](0527-fglimagepath.md "Defines a list of paths and filenames for image resources.")
environment variable for the mobile app, using FGLAPPDIR and FGLDIR predefined environment variables:

```
mobile.environment.FGLIMAGEPATH =
 "$FGLAPPDIR/myimages:$FGLAPPDIR/icons/myimage2font.txt:$FGLDIR/lib/image2font.txt"
```

During development (when executing programs on a server), consider defining environment variables
such as FGLAPPDIR in the shell environment, along with the other environment variables that are
defined with `mobile.environment` entries, as these are only read when executing on
mobile devices.

## Related links

**Related concepts**  

[Mobile applications](../17_mobile-applications/5081-mobile-applications.md "These topics cover programming subjects about mobile applications")
