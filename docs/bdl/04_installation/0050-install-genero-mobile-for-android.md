---
title: "Install Genero Mobile for Android"
source: "fgl-topics/t_gma_install.html"
breadcrumb: "Installation > Install Genero Mobile for Android™"
type: "task"
description: "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA."
---

# Install Genero Mobile for Android

> To build and package Genero Mobile for Android™ (GMA) applications, you must first install GMA.

**Before you begin:**

The installation and licensing of Genero products requires you to read and
accept the End User License Agreement, which can be found on the Four Js website at
<https://4js.com/end-user-license-agreements/>.

> **Important:**
>
> The GMA and Genero BDL X.YY versions are interdependent in terms of pcode compatibility. For
> example, to build a GMA 6.00 embedded app, the source code needs to be compiled with Genero BDL
> 6.00. The Genero BDL runtime version used by GMA embedded apps can be found with the
> gmabuildtool -V command.

- Download Genero Mobile for Android
  (GMA) from the Four Js Web site.
- Install Genero Business Development Language. See [Installing Genero BDL](0048-installing-genero-bdl.md "This section provides Genero BDL installation instructions.") for
  more details.

1. Install the Java Development Kit (JDK).

   JDK 17 is required to build Android apps. For the latest information regarding system requirements and
   Java support, please refer to the Supported platforms and databases document, available
   on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).
2. Prepare the Android SDK installation
   directory: First install the Android
   Command Line Tools and setup the environment. The download of Android SDK packages will be done in a next step with gmabuildtool
   updatesdk:
   1. Create a directory to hold the Android SDK packages and tools.
   2. Set the ANDROID\_HOME environment variable with the new created directory.
   3. Create the cmdline-tools directory in
      $ANDROID\_HOME.
   4. Download the [commandlinetools zip file](http://developer.android.com/studio#cmdline-tools) and unzip the ZIP archive
      into the $ANDROID\_HOME/cmdline-tools directory.

      > **Important:**
      >
      > On Windows® 10, the default zip tool
      > provided by the operating system will corrupt the Android Command Line Tools zip archive. The error displayed is
      > `"Error: Could not find or load main class
      > com.android.sdklib.tool.sdkmanager.SdkManagerCli"`. To avoid this error, use another zip
      > program such as 7zip.
   5. Rename $ANDROID\_HOME/cmdline-tools/cmdline-tools, to
      "$ANDROID\_HOME/cmdline-tools/latest".

   The commandlinetools ZIP archive provides
   cmdline-tools as root directory, before renaming.

   1. Add "$ANDROID\_HOME/cmdline-tools/latest/bin" to your PATH environment
      variable, to find the sdkmanager command.
3. Install the GMA buildtool and the GMA binary archive.

   They are provided in the GMA distribution archive (fjs-gma-\*.zip).

   1. Create a directory (gma-install-dir) for the GMA development tools.

      ```
      $ mkdir /opt/fourjs/gma-4.00
      ```
   2. Extract the content of the Genero Mobile for Android package (fjs-gma-\*.zip) into this
      directory.

      ```
      $ unzip -q -o -d /opt/fourjs/gma-4.00 fjs-gma-*.zip
      ```
   3. Add the gma-install-dir directory to your PATH environment variable, in
      order to find the gmabuildtool command.
4. Execute the gmabuildtool updatesdk command.

   An internet connection is required.

   Execute the gmabuildtool updatesdk command every time a new version of the GMA
   buildtool and GMA binary archive is installed.
5. If you plan to publish your app on Google Play, [register to Google Play as a developer and create a Google Play
   project](http://developer.android.com/distribute/googleplay/start.html).
6. If you have installed Cordova plugins, you need to re-install the plugins with the
   `--install-plugin` option of gmabuildtool. For more details,
   see [Installing Cordova plugins](../17_mobile-applications/5119-installing-cordova-plugins.md "Before usage, Cordova plugins need to be installed in the GMA or GMI development environment.").

## Related links

**Related concepts**  

[Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.")
