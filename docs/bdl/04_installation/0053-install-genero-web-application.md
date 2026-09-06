---
title: "Install Genero Web Application"
source: "fgl-topics/t_fgl_gwa_install.html"
breadcrumb: "Installation > Install Genero Web Application"
type: "task"
---

# Install Genero Web Application

> To build and package Genero Web applications, you must first install Genero Web Application (GWA).

**Before you begin:**

The installation and licensing of Genero products requires you to read and
accept the End User License Agreement, which can be found on the Four Js website at
<https://4js.com/end-user-license-agreements/>.

> **Important:**
>
> The Genero Web Application and Genero BDL X.YY versions are interdependent in terms of pcode
> compatibility: To build a GWA 6.00 embedded app, the source code needs to be compiled with Genero
> BDL 6.00.

- Download Genero Web Application from the [Four Js Web site](https://4js.com/download/products/).

  The Genero Web Application product is
  delivered as a zip that can be unzipped either to the FGLDIR or to a separate directory. The
  recommended installation is to unzip it directly over FGLDIR. Examples of both installation methods
  are given here.

  The zip file will contain all tools to build the final GWA application. Notice
  that those tools are written in Genero and require a valid Genero BDL 5.01 or greater to be
  installed.
- Install Genero Business Development Language. See [Installing Genero BDL](0048-installing-genero-bdl.md "This section provides Genero BDL installation instructions.") for
  more details.

The preferred installation is to unzip the Genero Web Application package
(fjs-gwa-\*.zip) directly over FGLDIR.

This has two advantages:

- The GWA tools, `gwabuildtool`, `gwarun`, `gwadb`,
  and `gwasrv`, are immediately in the PATH if `$FGLDIR/bin` is in the
  PATH already.
- The GWA API is immediately available for `IMPORT FGL` without the need to set
  `FGLLDPATH` explicitly.

Install the GWA over FGLDIR.

Extract the content of the Genero Web Application package (fjs-gwa-\*.zip)
into FGLDIR.

```
$ unzip -q -o -d $FGLDIR fjs-gwa-*.zip
```

## Install GWA in its own directory

If you need to test several versions of GWA against a single installation of Genero BDL, you will
extract each GWA distribution archive to a separate directory.

Install GWA in a directory separate to FGLDIR.

1. Create a directory (gwa-install-dir) for the GWA installation.

   ```
   $ mkdir /opt/fourjs/gwa-6.00
   ```
2. Extract the contents of the Genero Web Application package (fjs-gwa-\*.zip)
   into this directory.

   ```
   $ unzip -q -o -d /opt/fourjs/gwa-6.00 fjs-gwa-6.00.01-build-*.zip
   ```
3. Add the gwa-install-dir/bin directory to your PATH
   environment variable, in order to find the GWA tools in the path.

   ```
   export PATH=/opt/fourjs/gwa-6.00/bin:$PATH
   ```
4. Add the gwa-install-dir/lib directory to your FGLLDPATH
   environment variable.

   You may need the GWA package library during your development, for example, to get the location
   of windows, or to reload the application.

   ```
   export FGLLDPATH=/opt/fourjs/gwa-6.00/lib:$FGLLDPATH
   ```

## Related links

**Related concepts**  

[Overview](../18_genero-web-applications/5124-overview.md "Genero Web Application (GWA) is a browser-based solution that combines Genero Browser Client, p-code modules, and a WebAssembly-based runtime for executing applications offline as progressive web apps. It supports full Genero GUI features, requires HTTPS for secure deployment, and is ideal for scalable use cases, including offline functionality and large user bases.")
