---
title: "DBTEMP"
source: "fgl-topics/c_fgl_EnvVariables_DBTEMP.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > DBTEMP"
type: "concept"
---

# DBTEMP

> Defines the directory for temporary files.

The DBTEMP environment variable defines the directory for temporary files created by the runtime
system.

If the DBTEMP variable is not defined, the runtime system uses the temporary directory as defined
on the operating system. Depending on the platform, the [TMPDIR, TMP, TEMP](0503-tmpdir-tmp-temp.md "Defines the directory for temporary files.") environment variables, or
the default system temp directory will be used.

> **Important:**
>
> The DBTEMP environment variable is also used by the IBM® Informix® database client and
> server for temporary files.

The temporary directory is used to create temporary files for:

1. [`TEXT`](../08_language-basics/0569-text.md "The TEXT data type stores large text data.") or [`BYTE`](../08_language-basics/0555-byte.md "The BYTE data type stores any type of binary data, such as images or sounds.") data located in a temporary file
   (`LOCATE IN FILE` without filename specification).
2. Temporary files of [emulated scrollable
   cursors](../10_sql-support/1021-scrollable-cursors.md "How scrollable cursors can be supported on different databases.") when the database engine does not support this feature.
3. Temporary filename generation with [`os.Path.makeTempName()`](../15_library-reference/3730-os-path-maketempname.md "Generates a new file path to be used to create a temporary file or directory.").
4. Temporary files created by the Web Services API, such as [com.HttpResponse.getFileResponse](../15_library-reference/3895-com-httpresponse-getfileresponse.md "Returns the entire HTTP response to a file on the disk.").

On mobile devices, do not set DBTEMP environment variable: The runtime system will automatically
use the appropriate temporary directory within the app sandbox file system.
