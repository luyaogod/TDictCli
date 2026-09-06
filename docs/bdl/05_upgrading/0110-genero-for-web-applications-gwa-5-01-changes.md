---
title: "Genero for Web Applications (GWA) 5.01 changes"
source: "fgl-topics/c_fgl_Migrate_to_501_gwa_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 5.01 upgrade guide > Genero for Web Applications (GWA) 5.01 changes"
type: "concept"
---

# Genero for Web Applications (GWA) 5.01 changes

> Modifications to consider when using Genero for Web Applications.

> **Note:**
>
> This topic describes features changes in the GWA 5.01 product.

## GWA 5.01 with FGLGWS 5.01

> **Important:**
>
> The GWA version 5.01 is built on FGLGWS 5.01 and therefore, strongly tied to
> this Genero BDL version.

## Changes to the /home directory

Starting from GWA 5.01.02, each GWA program with a different path on the same domain has a
distinct home directory that is mapped in [IndexedDB](https://developer.mozilla.org/en-US/docs/Web/API/IndexedDB_API) (external link) caches, where the application state can be stored even if the
application is restarted or the browser is restarted.

Previously, the `/home` directory was shared across the entire domain. Now, each
application has its own subdirectory within `/home`, such as
`/home/build_directory_basename`, where `build_directory_basename`
corresponds to the name of the program's build directory.

For example, the `os.Path.pwd()` for an application's home directory now returns
/home/build\_directory\_basename instead of just
/home as it was before.

For more information on the GWA file system in the browser, go to [File system](../18_genero-web-applications/5125-file-system.md "When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory."). For more information about managing persistence in the GWA, go
to [Manage persistence in the file system](../18_genero-web-applications/5125-file-system.md "As directories in a GWA application are in memory, this means the lifetime of those directories is the same as the lifetime of the application. Understanding the file system will help you develop a strategy for making data persistent.").

## Specifying the working directory as app

If your application relies on relative paths and expects assets to be in the current working
directory (for example, FGL demo programs), you need to set the initial working directory to the
app directory when building your GWA package.

Starting with GWA 5.01.02, the path to the working directory can be set to
app with the gwabuildtool option
`--app-dir-is-pwd`.

For more information on the GWA file system in the browser, go to [File system](../18_genero-web-applications/5125-file-system.md "When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory."). For more information about gwabuildtool
options, go to [gwabuildtool](../18_genero-web-applications/5150-gwabuildtool.md).
