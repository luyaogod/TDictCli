---
title: "Building Android apps with Genero"
source: "fgl-topics/c_fgl_gma_app_build_tool.html"
breadcrumb: "Mobile applications > Deploying mobile apps > Deploying mobile apps on Android™ devices > Building Android™ apps with Genero"
type: "concept"
description: "Genero provides a command-line tool to create applications for Android devices."
---

# Building Android apps with Genero

> Genero provides a command-line tool to create applications for Android™ devices.

## Basics

Genero mobile apps for Android are
distributed as AAB or APK packages, like any other Android app.

- AAB packages can by published on the Google Play Store. To install the app from an AAB package
  onto your development device or emulator, you need the bundletool utility.
- APK packages (the former package format) can installed onto your development device or emulator.
  APK packages cannot be uploaded to the Google Play Store.

Genero provides the gmabuildtool command line tool, to build the AAB/APK
packages for your mobile application.

For more details about Android
packages and bundles, go to <https://developer.android.com/guide/app-bundle>.

> **Important:**
>
> You need to comply with some Google conditions before publishing your app on the Play Store. This
> topic describes how the gmabuildtool tool builds Google Play Store ready
> packages.

For testing purposes, the tool can deploy APK packages and automatically launch the app on a
specific device or emulator.

The gmabuildtool has also an option to update the Android SDK and an option to manage scaffold archives.

This documentation section implies that you are familiar with Android app programming concepts and requirements. For example, you will
need the Android SDK tools to be installed
(and up to date) to build your Android
apps. For more details, visit the [Android SDK tools download page](http://developer.android.com/studio/index.html#command-tools).

## Prerequisites

Before starting the command line tool to build or deploy the app, fulfill the following prerequisites:

1. The Genero BDL development environment (FGLDIR) must be installed to compile your program
   files.
2. The Java JDK must be installed.

   JDK 17 is required to build Android apps. For the latest information regarding system requirements and
   Java support, please refer to the Supported platforms and databases document, available
   on the "Products" download page of the [Four Js Web site](https://4js.com/download/products/).
3. The GMA software components must be installed.

   The GMA buildtool and GMA binary archive are
   provided in the GMA distribution archive (fjs-gma-\*.zip).

   To set up the
   GMA installation directory, perform the following steps:
   1. Create a dedicated directory (gma-install-dir)
   2. Extract the content of the fjs-gma-\*.zip.
   3. Add the gma-install-dir directory to your PATH
      environment variable, in order to find the gmabuildtool command.

   The gma-install-dir will contain the
   gmabuildtool command, and the original GMA scaffold archive
   (gma-install-dir/artifacts).
4. The Android SDK must be installed
   (an internet connection is required to download the SDK packages). For more details, see [Install Genero Mobile for Android](../04_installation/0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA.").
5. Download/upgrade the required Android
   SDK packages, by using the gmabuildtool updatesdk ... command. Execute the
   gmabuildtool updatesdk ... command every time a new version of the GMA buildtool
   and GMA binary archive is installed. For more details, see Update the Android SDK with the GMA buildtool.
6. Android-specific app resources such as
   icons (in all required sizes) are required, along with the application program files.
7. To install AAB packages on your development device or emulator, get the
   bundletool utility (see <https://github.com/google/bundletool/releases>).
8. If you plan to publish your app on Google Play, [register to Google Play as a developer and create a Google Play project](http://developer.android.com/distribute/googleplay/start.html).

## Environment settings

Define the following environment variables before starting the command-line buildtool:

- Android SDK environment settings
  (ANDROID\_HOME, PATH)
- Java JDK environment settings (JAVA\_HOME, PATH)

## Update the Android SDK with the GMA buildtool

After a fresh installation of the GMA buildtool and GMA binary archive, upgrade the Android SDK, to download all required Android SDK packages, with the
gmabuildtool updatesdk ... command.

The Android SDK installation directory
is required for the SDK update. It can be specified by the ANDROID\_HOME environment variable or with
the `--android-sdk` option:

```
gmabuildtool updatesdk
    --android-sdk /use/local/android-sdk/r22.6.2
```

Use the `--accept-licenses` option to silently accept all Android SDK licenses.

```
gmabuildtool updatesdk
    --accept-licenses
    ...
```

If you need to specify a proxy to download the Android SDK, use the `--proxy-host` and
`--proxy-port` options:

```
gmabuildtool updatesdk
    --proxy-host amadeus --proxy-port 3232
    ...
```

Use the `--no-install-extras` option to skip the installation of extra SDK modules
such as Google's driver for Windows®
and the HAXM for Windows and OS
X:

```
gmabuildtool updatesdk
    --no-install-extras
    ...
```

On Windows and Mac® extras may be installed if
`--no-install-extras` option is not used. The installation is preceded by a prompt
asking the user to continue with the installation. Use the `--no-interactions` option
to skip user interaction. This option answers yes silently to the prompt for permission to
install.

This option also silently accepts licenses. `gmabuildtool updatesdk
--no-interactions` is the same as the command `gmabuildtool updatesdk
--no-interactions --accept licenses`

```
gmabuildtool updatesdk
    --no-interactions
    ...
```

## Manage GMA scaffold archives

The GMA scaffold archives can be managed by the gmabuildtool scaffold ...
command.

In order to get the list of plugins installed in the GMA development environment, use the
`--list-plugins` option:

```
$ gmabuildtool scaffold --list-plugins
...
cordova-plugin-device-motion.aar
cordova-plugin-media.aar
...
```

To install additional plugins in the GMA installation directory, use the
`--install-plugins` option, for example:

```
$ gmabuildtool scaffold --install-plugins path-to-plugin-sources
```

For more usage examples, see [Cordova plugins](5117-cordova-plugins.md "This section describes how to use Cordova plugins.").

## Building AAB/APK packages

The gmabuildtool build ... command creates the AAB and/or APK packages from a
set of files, using the build options passed as parameter:

```
gmabuildtool build
    ... build options ...
```

The process will be explained in details in the next sections of this topic.

For a complete description of the `build` command options, see [gmabuildtool](5106-gmabuildtool.md "The gmabuildtool is a utility to create app packages for an Android device.").

## Gradle build cache

To build Android apps, the GMA
buildtool uses the [Gradle](https://docs.gradle.org)
toolkit.

Gradle can speed up AAB/APK creation time, by reusing outputs saved in a cache produced from
previous builds. The gmabuildtool command uses the Gradle build cache.

On Unix-like platforms, the gmabuildtool Gradle build cache is in the
/tmp/gma directory. On Windows, it is in C:\tmp\gma.

> **Important:**
>
> The Gradle build cache directory used by gmabuildtool
> (/tmp/gma or C:\tmp\gma) can grow to a significant size
> and consequently fill the disk. Take care to monitor the size of this cache and manually clean out
> the cache files, to ensure that it does not become too large in size.

## Cleaning intermediate build files

The build process is optimized to avoid a complete AAB/APK package rebuild every time you invoke
the GMA buildtool.

The GMA buildtool creates intermediate archive files, that can be reused in the next build, if no
changes are detected in application program files. However, these files might be corrupted, in case
of user interruption or gradle build failure.

In this situation, you can use the `--clean` option of the gmabuildtool
build ... command, to clean up the scaffold build directory, and restart with a fresh
build:

```
gmabuildtool build --clean
    ... build options ...
```

## Specifying the GBC to be used

Use the `--build-gbc-runtime` option to specify which GBC has to be bundled with
the app
package:

```
gmabuildtool build --build-gbc-runtime gbc-archive ...
```

The gbc-archive parameter is the ZIP archive of the GBC front-end, to be
created as described in the Create a runtime zip topic in the Genero Browser Client User Guide.

The parameter for the `--build-gbc-runtime` option can also be a regular
(unzipped) GBC directory.

If the `--build-gbc-runtime` option is not used, gmabuildtool
will use the GBC found in [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode."), and if
that environment variable is not defined, it defaults to [FGLDIR](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")/web\_utilities/gbc/gbc.

## Using an options file

To simplify option specification, create a file with the list of options to be passed to the
gmabuildtool with the `@argfile`
argument:

```
$ cat myoptions.txt
--build-output-apk-name MyApp
--build-app-package-name com.example.myapp
--main-app-path /tmp/app/main.42m
...
$ gmabuildtool build @./myoptions.txt
```

Several options files can be specified and mixed with command line options, for
example:

```
$ gmabuildtool build @com_options.txt @app_options.txt -v
```

## Elements used in building the Android app

The gmabuildtool build ... command builds the Android AAB/APK packages from the following:

- The GMA binary archive, containing the GMA front-end and the FGLGWS runtime system. The original
  scaffold is provided in gma-install-dir/artifacts and
  unzipped for app package builds into a directory defined by the `--build-project`
  option.
- The app manifest file `AndroidManifest.xml` to be created/used
  (`--generate-default-manifest`, `--custom-manifest-file` options)
- The GBC to be used (`--build-gbc-runtime` option)
- The Android foreground service
  type, to publish the app on the Google App Store (`--custom-foreground-service-type`
  and `--background-explanation` options)
- The compiled application program and resource files (.42m,
  .42f, etc) (found from the `--main-app-path` option)
- The default path to Android app
  resources (`--root-path`), when not using resource dedicated options such as
  `--build-app-icon-hdpi`. The root path will also be used to create all intermediate
  build files.
- The prefix for the app package filename to be generated
  (`--build-output-apk-name` option)
- The package name and the name of the app (`--build-app-package-name` option)
- The version code of the app (`--build-app-version-code` option)
- The version name of the app (`--build-app-version-name` option)
- Android app specific resources:
  - Android app icons (all sizes)
    (`--build-app-icon*` and `--build-status-icon-*` options)
- Android app signing keystore alias and
  file, to publish the app (`--build-signing-keystore` and related options)

## The app manifest file

Every Android app must be created
with an app manifest file named AndroidManifest.xml.

The gmabuildtool provides options to create a default manifest file that can
be modified, and an option to specify the manifest file to be used when building the AAB/APK.

The [`--generate-default-manifest/-gdm`](5106-gmabuildtool.md) option can be used to generate a default
AndroidManifest.xml file, that can be modified to build the app package. The
file is generated in the build directory (`--build-apk-outputs/-bo` option). This
option takes into account manifest-related options such as
`--build-app-permissions/-ba`.

To build the app packages, a user-defined manifest file can be specified with the [`--custom-manifest-file/-cmf`](5106-gmabuildtool.md) option. The basename of the provided file path
must be `AndroidManifest.xml`. If this option is not used,
`gmabuildtool` will generate a new manifest file.

Note that the build process to produce the AAB/APK packages may add permissions coming from
external librairies, that are not present in the provided manifest file.

For more details about app manifest files, see the [Android
documentation](https://developer.android.com/guide/topics/manifest).

## Defining the Android service type

> **Tip:**
>
> If you do not plan to put your app in the Google Play Store, you can ignore this step.

> **Tip:**
>
> If you plan to put your app in the Google Play Store, first try without changing the service type
> and see whether it is accepted. If refused by Google, then try changing the default explanation or
> changing the service type to "`shortService`".

A GMA application embeds a dynamic virtual machine (DVM) that needs to run even if the GMA
application is in the background. To do so, GMA use a mechanism called foreground
service.

Starting with Android 14, to be
accepted on the Google Play Store, GMA apps using foreground services must declare a foreground
service type in the manifest. A foreground service type – such as `shortService` or
`specialUse` – is declared with the [`foregroundServiceType`](https://developer.android.com/about/versions/14/changes/fgs-types-required) manifest parameter.

The foreground service type can be controlled with the gmabuildtool options
`--custom-foreground-service-type` and `--background-explanation`:

- The [`--custom-foreground-service-type service-type`](5106-gmabuildtool.md) option
  defines a foreground service type, such as `specialUse`,
  `shortService`, `camera`, or `dataSync`. To specify
  multiple service types, repeat the option with specific values.
- If the `--custom-foreground-service-type` option is not specified, the default is
  `specialUse` and the [`--background-explanation`](5106-gmabuildtool.md) option is required if you plan to upload your app
  to the Google Play Store. Google requires the use cases for the `specialUse` service
  type to be precisely described and approved in order for the app to be accepted on the Google Play
  Store.
- The service description is specified with the `--background-explanation` option
  parameter. You can try the default background explanation to validate a package if needed; it reads
  “This application embeds a virtual machine (process) that needs to be used as a foreground service
  when the application is in the background to avoid losing any data.”
- When you specify the service type as `shortService`, the operating system kills
  the fglrun process after approximately 4 minutes (depending on factors such as battery state) after
  the GMA app goes into the background. The GMA application can use the [`enterbackground`](../11_user-interface/2257-list-of-predefined-actions.md) action to save its current state when it moves to the
  background, ensuring that this saved state is restored when the app is brought back to the
  foreground and relaunched.

  While the service type of `shortService` has the downside of being killed in
  background after 4 minutes, it will normally be accepted by Google validation without the need of
  any explanation.
- The Google validation process may ask the user to provide additional explanation regarding the
  app usage. Such requirements are subject to change and not controlled by Genero; they are determined
  by Google.
- Changing the service type to a type other than `specialUse` or
  `shortService` should be rare and only done if you know exactly what you are
  doing!

## Generate the keystore file to sign your app

In order to build an AAB package that can be deployed on Google Play Store, you need to sign your
Android app with a keystore.

> **Important:**
>
> There is critical information regarding app signing in the [Android
> documentation](http://developer.android.com/tools/publishing/app-signing.html#signing-manually). For example, in order to install new versions as updates to your app, each app
> package must use the same certificate: The signing key must not change during the lifetime of your
> app. Make sure that you read all details about Android app signing.

The keystore is used by the gmabuildtool to sign the APK with the
apksigner utility from Android SDK and to sign the AAB with the jarsigner utility from Java
SDK.

The signing credentials are passed to `gmabuildtool` with the
`--build-signing-keystore`, `--build-signing-alias`,
`--build-signing-storepass` and `--build-signing-keypass` options.

Using `--build-signing-keystore` option makes all other
`--build-signing-*` options mandatory.

Default values of all `--build-signing-*` options will only apply, if the
`--build-signing-keystore` option is not specified. This is done to avoid setting
some weak default password if user misses to use some options. For defaults, see [gmabuildtool](5106-gmabuildtool.md) command reference.

The passwords specified with `--build-signing-keypass` and
--build-signing-storepass must be 6 char long minimum (required by key generation), and must be
identical since different passwords is not supported by the PKCS12 key format. Passwords can be
different, when specifying a path to an existing key file that is in a encryption format allowing
this.

When no keystore exists, or if the alias in the key does not exists, the
gmabuildtool will automatically create a keystore of type PKCS12.

## Controlling the production of AAB/APK package files

By default, the gmabuildtool produces both AAB and APK packages. To build only
one type of Android package, use the
`--output-format` option to specify `"aab"` or `"apk"`
types:

```
gmabuildtool ... --output-format aab
```

The
`--output-format` option supports also the `"all"` value, to build all
type of packages. Uppercase AAB or APK values are accepted.

The filename of the app package is formed from:

1. The filename prefix defined by the `--build-output-apk-name` option (by default,
   "app"),
2. When building a debug version, the `-debug` suffix,
3. The `.aab` or `.apk` file extension.

## Directory structures required for the build

To build an Android app, you must
prepare the main-app-dir directory containing the program files
(.42m, .42f, ...) and resource files (images, database
files, webcomponents) needed by the Genero application. The main-app-dir is
implicitly defined by the parent directory of the [`--main-app-path`](5106-gmabuildtool.md) option argument. This directory must have a structure as
described in [Directory structure for GMA apps](5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA)."):

```
main-app-dir
|
|-- main.42m
|-- comutils.42m
|-- mainform.42f
|-- fglprofile
|-- defaults/*.42s
|   ...
```

For convenience, the GMA buildtool supports a default directory structure
(root-dir) to find Android resource files (icons) required to build the app package. The
root-dir can be specified with the [`--root-path`](5106-gmabuildtool.md) option. This root directory also defines the default directory
for scaffold and temporary files used by the build process. It is strongly recommended to use
different locations for the root directory and for the main application directory:

```
root-dir
|
|-- project
|   |-- scaffold
|   |-- build
|   |   ...
|
|-- icons
|   |-- ic_app_hdpi.png
|   |-- ic_app_mdpi.png
|   |-- ic_app_xhdpi.png
|   |-- ic_app_xxhdpi.png
|   |   ...
```

1. root-dir/project is the default directory for scaffold
   and temporary files used by the build process. You can use a different directory with the [`--build-project`](5106-gmabuildtool.md) option.
2. root-dir/project/scaffold contains the scaffold files
   copied automatically by gmabuildtool from
   gma-install-dir/artifacts.
3. root-dir/project/build us used as working directory for
   the gradle process.
4. root-dir/icons is the default directory containing the
   app icons. You can specify app icons with options such as [`--build-app-icon-hdpi`](5106-gmabuildtool.md).

## Android permissions

To use a mobile device feature such as the camera, an Android app must be created with the corresponding Android permission.

The permissions implemented in the APK/AAB package will be a merge of the permissions set by
default (typically for mobile front calls), the permissions defined for the app with the
gmabuildtool, and the permissions already implemented by external libraries
required to build the GMA app. The manifest file used by the gmabuildtool may
show less permissions than the APK/AAB package implements.

Android distinguishes "Normal
Permissions" and "Dangerous Permissions": Dangerous Permissions require a user validation at
runtime, the first time the device feature is used.

Before Android 6, Dangerous Permissions
defined by the app were set at app installation. Starting with version 6, Dangerous Permissions must
be requested by the app code on demand.

Android permissions can be specified
with the `--build-app-permissions` option of the gmabuildtool.
Define the list of permissions as a single argument, by using the comma as separator.

For example:

```
gmabuildtool build \
    ...
    --build-app-permissions android.permission.READ_CALENDAR,... \
    ...
```

Assuming that the permission was specified when building the app package, a Dangerous Permission
required for a core GMA feature (like internet access), or a built-in front call, will make the GMA
automatically ask for user confirmation. For example, if the app code makes a [`chooseContact`](../15_library-reference/3442-mobile-choosecontact.md "Lets the user choose a contact from the mobile device contact list and returns the vCard.") front call,
the GMA will automatically request the user to access the contacts database.

Other permissions (not related to core GMA features or built-in front calls) need to be specified
when building the app.

Dangerous Permissions not related to core GMA features or built-in front calls, need to be
enabled by the app code. To request the user for a specific permission, perform a [`askForPermission`](../15_library-reference/3464-android-askforpermission.md "Ask the user to enable a dangerous feature on the Android device.") front
call, before using the feature.

Several Android "Normal Permissions"
required for core GMA app features and built-in mobile front calls are automatically defined in the
app manifest file by gmabuildtool. To identify these default permissions, inspect
the manifest file of the generated APK/AAB file.

The following list shows the Android
"Dangerous Permissions", required for core GMA app features and built-in mobile front calls.
> **Important:**
>
> Dangerous Permissions listed below are not set by default: They must be specified explicitly with
> the `--build-app-permissions` option, when building the app package.

- `android.permission.ACCESS_FINE_LOCATION`: For [mobile.getGeolocation](../15_library-reference/3451-mobile-getgeolocation.md "Returns the Global Positioning System (GPS) location of a mobile device.") front call.
- `android.permission.ACCESS_COARSE_LOCATION`: For [mobile.getGeolocation](../15_library-reference/3451-mobile-getgeolocation.md "Returns the Global Positioning System (GPS) location of a mobile device.") front call.
- `android.permission.READ_CONTACTS`: For [mobile.chooseContact](../15_library-reference/3442-mobile-choosecontact.md "Lets the user choose a contact from the mobile device contact list and returns the vCard.") front call.
- `android.permission.GET_ACCOUNTS`: Only on Android 5.1 and lower (< API 23), for [mobile.chooseContact](../15_library-reference/3442-mobile-choosecontact.md "Lets the user choose a contact from the mobile device contact list and returns the vCard.") front call.
- `android.permission.CAMERA`: For [mobile.takeVideo](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier."), [mobile.takePhoto](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.")
  front calls.
- `android.permission.WRITE_EXTERNAL_STORAGE`: For [mobile.importContact](../15_library-reference/3453-mobile-importcontact.md "Creates a new contact, or merges to an existing entry, the contact details passed in a vCard string."), [mobile.takeVideo](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier."), [mobile.takePhoto](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.")
  front calls.
- `android.permission.READ_EXTERNAL_STORAGE`: For [standard.openFile](../15_library-reference/3409-standard-openfile.md "Displays a file dialog window to let the user select a single file path on the local file system."), [standard.openFiles](../15_library-reference/3410-standard-openfiles.md "Displays a file dialog window to let the user select a list of file paths on the local file system."), [mobile.chooseVideo](../15_library-reference/3444-mobile-choosevideo.md "Lets the user select a video from the mobile device's video gallery and returns a video identifier."), [mobile.takeVideo](../15_library-reference/3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier."), [mobile.choosePhoto](../15_library-reference/3443-mobile-choosephoto.md "Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier."), [mobile.takePhoto](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.") front calls.

For a complete list of Android
permissions, see [Android's Manifest permissions](http://developer.android.com/reference/android/Manifest.permission.html).

## Define app's color theme

The app colors can be defined with a GBC theme.

For example, the GBC theme variable `"theme-primary-background-color"` is used to
defined the main color of the app.
> **Important:**
>
> Theme names in GBC customization must have a valid Java-id variable name, in order to work with
> GMA.

Refer to the GBC manual for more details about GBC customization and theme colors usage.

The GBC theme is applied when the app starts, and every time the theme is changed by a [`theme.setTheme`](../15_library-reference/3429-theme-settheme.md "Activates a specific theme.") frontcall.

The GBC theme is also used to colorize some native UI elements such as the date/time picker, and
the navigation bar. When changing the theme dynamically during the app execution, the colors will
only apply to new created native widgets such as the date/time picker. Already existing native
widgets like the navigation bar will not be affected by a theme change. However, the last used theme
is saved in application settings, and all native widgets will get the expected colors when the app
is restarted.

> **Important:**
>
> The `--build-app-colors` is deprecated, use [GBC theming to define
> application colors](5105-building-android-apps-with-genero.md).

## Debug and release versions

Android apps can be generated in a
debug or release version. Release versions meet the requirements for distribution on Google Play,
while debug versions are used in development. In debug mode, the app installed on the device will
listen on the debug TCP port to allow [fgldb
-m connections](5101-debugging-a-mobile-app.md "Different solutions are available to debug a mobile app.").

Debug or release mode can be controlled with the `--build-mode` option of the
gmabuildtool command:

```
gmabuildtool build \
    --build-mode debug \
    ...
```

By default the app is built in release mode.

## Building an Android app with gmabuildtool

Follow the next steps to create an Android app:

- Prepare the root distribution directory root-dir/gma,
  with Android app resources (icons) under
  root-dir/gma. Or use specific options such as
  `--build-app-icon-hdpi` to define the location of app icons.
- Create the main application directory with Genero program files, based on the default directory structure.
- Execute the gmabuildtool command with the `build` verb and
  related build options, to create the Android packages. The location of main application directory must be specified with the
  `--main-app-path` option.

For example:

```
$ gmabuildtool build \
    --main-app-path /home/mike/my_genero_app/main.42m \
    --root-path /home/mike/my_app_project \
    --android-sdk /home/mike/android/sdk \
    --build-apk-outputs /home/mike/work/example/outputs \
    --output-format aab \
    --build-output-apk-name MyApp \
    --build-app-package-name com.example.myapp \
    --build-app-version-code 1002 \
    --build-app-version-name "10.02" \
    --build-signing-alias android_alias \
    --build-signing-keystore /home/mike/work/example/sign/android.keystore \
    --build-signing-keypass fourjs \
    --build-signing-storepass fourjs \
    --build-mode release \
    --build-app-permissions android.permission.ACCESS_WIFI_STATE,android.permission.CALL_PHONE
```

## Building an app with GMA custom extensions

The gmabuildtool build ... command supports AAB/APK creation for applications
using GMA custom extensions written in Java.

Before building the app package, create the custom GMA binary archive with your extensions, as
described in [Packaging custom Java extensions for GMA](../14_extending-the-language/2697-packaging-custom-java-extensions-for-gma.md "Custom Java extension must be integrated in the GMA to run on Android devices.").

When your custom GMA binary archive is complete, build the app package with the
gmabuildtool build ... command. Use the `--build-project` option
to specify the path to the Android Studio
project that was used to build your custom GMA binary archive:

```
$ gmabuildtool build
    ...
    --build-project /home/mike/android_project/mycustgma
    ...
```

Other options have to be specified as for a regular build using the original standard GMA binary
archive.

## Deploy and launch the app

After building the app packages, you can deploy and launch your app for testing purpose, on your
local mobile device or simulator.

There must be only one Android device
connected or running Android emulator.

Android App Bundle
(AAB/.aab) packages cannot be installed directly on your local device or
simulator (these are archives to be uploaded on the Google Play Store).

To install the app on your local device or simulator, you need first to produce an
.apks for your device or simulator, by using the [bundletool](https://developer.android.com/studio/command-line/bundletool) (assuming your device is connected or simulator is
started):

```
$ bundletool build-apks \
     --bundle=/MyApp/outputs/my_app.aab \
     --output=/MyApp/outputs/my_app.apks
```

The .apks can then be used to install the app on you local device or
simulator, with the following command:

```
bundletool install-apks --apks=/MyApp/outputs/my_app.apks
```

Alternatively, the APK/.apk packages can be installed and started with the
gmabuildtool test ... command, by using the `--test-apk` option to
specify the APK filename to be installed:

```
$ gmabuildtool test \
     --test-apk /MyApp/outputs/MyApp-arm-debug.apk
```

## Related links

**Related concepts**  

[Genero Mobile common front calls](../15_library-reference/3441-genero-mobile-common-front-calls.md "This section describes common front calls provided by all mobile front-ends.")

[Color and theming](5095-color-and-theming.md "Mobile applications must follow the platform colors and theming.")
