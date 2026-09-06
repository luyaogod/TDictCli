---
title: "Environment variables"
source: "fgl-topics/c_fgl_mobile_env_vars.html"
breadcrumb: "Mobile applications > Environment variables"
type: "concept"
---

# Environment variables

> You may need to set environment variables for your app.

## Automatic environment variables

When executing on a mobile device:

- [DBDATE](../07_configuration/0508-dbdate.md "Defines the default display and input format for DATE values.") defaults to the regional settings
  defined on the device.
- [FGLAPPDIR](../07_configuration/0519-fglappdir.md "Contains the path to the application directory when executing on a mobile device.") is automatically defined to the
  application directory (`appdir`).

## Set environment variables

Setting environment variables for your app must be done in an fglprofile
file. This fglprofile file must be located in the appdir directory, beside the
main program module.

To add an environment variable for your mobile app, use the following
syntax:

```
mobile.environment.DBFORMAT="$:,:.:"
```

Any existing environment variable setting is overwritten by the value set (using
`mobile.environment.envvarname`) in the
fglprofile file.

For more details, see [Setting environment variables in FGLPROFILE (mobile)](../07_configuration/0495-setting-environment-variables-in-fglprofile-mobile.md).

Environment variables set in an FGLPROFILE file are only read when the deployed
application runs on the mobile device. They are not read during development mode (that is when the
VM runs on the development machine and the mobile client displays on the device). The FGLPROFILE
environment variable settings are only for the VM component and are ignored by the GMA/GMI front-end
component.
