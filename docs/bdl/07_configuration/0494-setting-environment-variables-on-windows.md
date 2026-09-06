---
title: "Setting environment variables on Windows"
source: "fgl-topics/c_fgl_EnvVariables_windows.html"
breadcrumb: "Configuration > Environment variables > Setting environment variables on Windows™"
type: "concept"
description: "On Windows™ platforms, environment variables can be set by one of the following methods: In a command window, with the SET command. In the registry, for the current user in HKEY_CURRENT_USER or a ..."
---

# Setting environment variables on Windows

On Windows™ platforms, environment
variables can be set by one of the following methods:

- In a command window, with the `SET` command.
- In the registry, for the current user in `HKEY_CURRENT_USER` or a global setting
  in `HKEY_LOCAL_MACHINE`.

For more details, refer to the documentation of your Windows system.

On Windows, double quotes
do not have the same meaning as on UNIX™ systems.
For example, if you set a variable with the command `SET VAR="abc"`,
the value of the variable will be `"abc"` (with double
quotes), and not `abc`.

When using Informix®,
some variables related to the database engine must be set using the
SETNET32 utility.
