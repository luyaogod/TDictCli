---
title: "TMPDIR, TMP, TEMP"
source: "fgl-topics/c_fgl_EnvVariables_TMPDIR_TEMP_TMP.html"
breadcrumb: "Configuration > Environment variables > Operating system environment variables > TMPDIR, TMP, TEMP"
type: "concept"
---

# TMPDIR, TMP, TEMP

> Defines the directory for temporary files.

The TMPDIR, TEMP and TMP environment variables define the directory where temporary files are
created by the operating system and by some other software (TMPDIR is typically used on UNIX™ platforms, TEMP and TMP are used on Windows™)

On desktop and server platforms, consider using [DBTEMP](0517-dbtemp.md "Defines the directory for temporary files.") to define the temp file directory for runtime system temporary files.

On mobile devices, there is no need to define the TMPDIR (or DBTEMP) environment variable:
The runtime system will automatically use the appropriate temporary directory within the app sandbox
file system.
