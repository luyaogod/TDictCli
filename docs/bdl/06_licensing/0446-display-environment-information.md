---
title: "Display environment information"
source: "genero-install-topics/t_license_controller_display_env.html"
breadcrumb: "Licensing > Manage Genero BDL local license > Displaying environment and statistics > Display environment information"
type: "task"
---

# Display environment information

> Use this procedure to display information about the environment of your machine where the license is installed.

Settings are extracted from the FGL configuration file (for example,
$FGLDIR/etc/fglprofile) and some from system
environment settings.

To display information about the environment, execute the command:

```
fglWrt -a env
```

The output shows details under the `Resource`,
`Environment` and `System` headings.
> **Note:**
>
> - If you have a local license installed and therefore **not** using a Four Js
>   License Manager (FLM) server, the `server` field will
>   be displayed as `(undefined)`.
> - If you have a local license installed, your license number and key will not be
>   displayed. License numbers that may be displayed are from the
>   fglprofile, if any had been filled in there.

![Image shows environment settings displayed by the fglWrt -a env command](../_images/fglwrt_a_env.jpg)

*Display the environment settings where license is installed.*
