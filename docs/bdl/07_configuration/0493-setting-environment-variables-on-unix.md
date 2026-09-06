---
title: "Setting environment variables on UNIX"
source: "fgl-topics/c_fgl_EnvVariables_unix.html"
breadcrumb: "Configuration > Environment variables > Setting environment variables on UNIX™"
type: "concept"
description: "On UNIX™ platforms, environment variables can be set through the following methods, depending on the command interpreter used: Bourne shell: VAR= value ; export VAR Korn shell: export VAR= value C ..."
---

# Setting environment variables on UNIX

On UNIX™ platforms, environment variables can be set through
the following methods, depending on the command interpreter used:

Bourne shell:

```
VAR=value; export VAR
```

Korn shell:

```
export VAR=value
```

C shell:

```
setenv VAR=value
```

For more details, refer to the documentation for your UNIX system.
