---
title: "File transfer support in GWA deployment"
source: "fgl-topics/c_fgl_gwa_file_transfer_support.html"
breadcrumb: "Genero Web applications > File transfer support in GWA deployment"
type: "concept"
---

# File transfer support in GWA deployment

> File transfer operations are supported in Genero Web Application (GWA) deployments using fgl_putfile and fgl_getfile.

The `fgl_putfile` function can be used to download a file from the virtual directory where the GWA resides in the browser. This operation is equivalent to a standard file download in any browser. For details, see [fgl\_putfile()](../15_library-reference/2774-fgl-putfile.md "Transfers a file from the virtual machine context to the front-end context.").

The `fgl_getfile` function is used for uploading: it locates a file from the user's machine and places it in the virtual directory of the GWA. For details, see [fgl\_getfile()](../15_library-reference/2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.").

For more information about the virtual directory system, see [File system](5125-file-system.md "When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory.").
