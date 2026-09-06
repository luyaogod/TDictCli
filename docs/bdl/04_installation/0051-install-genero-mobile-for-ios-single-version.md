---
title: "Install Genero Mobile for iOS (single version)"
source: "fgl-topics/t_gmi_install.html"
breadcrumb: "Installation > Install Genero Mobile for iOS (single version)"
type: "task"
---

# Install Genero Mobile for iOS (single version)

> To build and package Genero Mobile for iOS (GMI) applications, you must first install GMI. This topic explains how to install a unique GMI version/package into FGLDIR.

Before you begin:

The installation and licensing of Genero products requires you to read and
accept the End User License Agreement, which can be found on the Four Js website at
<https://4js.com/end-user-license-agreements/>.

> **Important:**
>
> The GMI and Genero BDL X.YY versions are interdependent in terms of pcode compatibility. For
> example, to build a GMI 6.00 embedded app, the source code needs to be compiled with Genero BDL
> 6.00. The Genero BDL runtime version used by GMI embedded apps can be found with the
> gmibuildtool -V command.

- Download the Genero Mobile for iOS (GMI) distribution archive from the [Four Js Web site](https://4js.com/download/products). GMI can only be
  installed on a macOS® computer.
- Install Genero Business Development Language. See [Installing Genero BDL](0048-installing-genero-bdl.md "This section provides Genero BDL installation instructions.") for
  more details.

Follow this procedure if you are using only one version of GMI for a Genero BDL
installation.
> **Important:**
>
> When re-installing a new GMI archive, remove all
> "build" directories created by the gmibuildtool.

1. Install Xcode®.

   The Xcode version must support the iOS
   versions of your mobile devices. As a general rule, update the Xcode and iOS to the latest versions.
2. Define the FGLDIR environment variable with the Genero BDL installation directory.

   ```
   $ export FGLDIR=/opt/fourjs/fglgws-3.10
   ```
3. Extract the GMI distribution archive (fjs-gmi-\*.zip) into FGLDIR.

   The GMI distribution archive contains the GMI buildtool and the needed iOS libraries and
   helper scripts.

   ```
   $ unzip -q -o -d $FGLDIR fgl-gmi-*.zip
   ```
4. If you have installed Cordova plugins, you need to re-install the plugins with the
   `--install-plugin` option of gmibuildtool. For more details, see
   [Installing Cordova plugins](../17_mobile-applications/5119-installing-cordova-plugins.md "Before usage, Cordova plugins need to be installed in the GMA or GMI development environment.").

## Related links

**Related concepts**  

[Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.")
