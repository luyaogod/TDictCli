---
title: "Building iOS apps with Genero"
source: "fgl-topics/c_fgl_gmi_app_build_tool.html"
breadcrumb: "Mobile applications > Deploying mobile apps > Deploying mobile apps on iOS devices > Building iOS apps with Genero"
type: "concept"
---

# Building iOS apps with Genero

> Genero provides a command-line tool to build applications for iOS devices.

## Basics

Genero mobile apps for iOS are distributed as IPA packages, like any other iOS app.
Genero provides a command line tool to build the .ipa package for your mobile application, or the
.app directory for simulators.

This documentation section implies that you are familiar with iOS app programming concepts and
requirements. In order to build your apps, you must have an Apple® developer account, as well as certificates and provisioning profiles to
deploy your apps. For more details, visit the [Apple developer site](http://developer.apple.com).

## Prerequisites

Before starting the command line tool to build or deploy the app, fulfill the following
prerequisites:

1. Xcode® must be installed on your macOS® computer (utilities from Xcode toolchain are required).

   Make sure that the installed
   Xcode version supports the iOS versions of
   your mobile devices. As a general rule, update the Xcode and iOS to the latest versions.
2. The Genero BDL development environment (FGLDIR) must be installed on the Mac computer
   to compile your program files.
3. The GMI archive (and gmibuildtool) must be installed and available.

   The
   GMI archive is provided as a ZIP archive (fjs-gmi-\*.zip). Extract the archive
   into FGLDIR.

   > **Warning:**
   >
   > If the GMI archive is not extracted into FGLDIR, the Xcode project of the
   > $GMIDIR/demo/MobileDemo/userfrontcall demo does not work.

   For test purposes, the GMI zip archive can be installed elsewhere in a location of your
   choice. You may choose to do this to resolve problems resulting from several GMI versions. If you
   install elsewhere, define an environment variable [GMIDIR](../07_configuration/0538-gmidir.md "Defines the installation directory of Genero Mobile for iOS.") with
   the install location and add $GMIDIR/bin to the path.

   The GMI buildtool
   does not need GMIDIR; However, all the supplied GMI demos need GMIDIR in the Makefiles.

   Check
   that the gmibuildtool command is available (add
   GMI-install-dir/bin to PATH for convenience).
   > **Important:**
   >
   > When re-installing a new GMI archive, remove all "build" directories created
   > by the gmibuildtool.
4. Get an Apple developer account, device
   identifiers (UDID) and corresponding identifiers to sign your iOS app (certificate, bundle id,
   provisioning profile).
   > **Important:**
   >
   > The UDID is the identifier of your physical device, it
   > can be found with the xcrun xctrace list devices Xcode command when the device is plugged to the Mac. When deploying on a
   > physical device, make sure that the UDID of the device is listed in the Apple Developer account that is used to generate the
   > provisioning profiles.
5. iOS app resources such as icons and launch images (in all required sizes).

## Finding the UDID of the plugged device

In order to find the UDID of the device plugged to your Mac, execute the xcrun xctrace
list devices Xcode command, and
identify the line describing your physical device:

```
$ xcrun xctrace list devices
== Devices ==
fraise [XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX]
iPod touch (9.1) [XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX]
iPad 2 (9.0) [XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX]
...
== Simulators ==
Apple TV (14.5) (XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX)
Apple TV 4K (2nd generation) (14.5) (XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX)
Apple TV 4K (at 1080p) (2nd generation) (14.5) (XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX)
ipad (8th generation) (14.5) (XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX)
...
```

In the above output, the UDID of the iPod® is `78b7452fa9462c98c3bc7047da344314fd032004`.

## Environment settings

Before starting the command-line buildtool, make sure that Xcode tools are available. Try xcodebuild
-version from the command
line:

```
$ xcodebuild -version
Xcode 11.3.1
Build version 11C504
```

## The gmibuildtool

The gmibuildtool command line tool can build IPA packages of iOS apps written
in Genero.

In order to identify the exact product version number of each GMI component, use the
`--version` option of the
gmibuildtool:

```
$ gmibuildtool --version
GMI version:1.40.04
VM  version:3.20.05
GWS version:3.20.05
Cordova ver:gm_1.30.14
```

## Manage GMI plugins

To get a list of plugins shipped in the fjs-fglgmixxx package, use the
`--list-stock-plugins` option:

```
$ gmibuildtool --list-stock-plugins
...
Calendar-PhoneGap-Plugin
GeneroTestPlugin
cordova-plugin-bluetoothle
cordova-plugin-device-motion
cordova-plugin-file
cordova-plugin-media
...
```

To install additional plugins in the GMI installation directory, use the
`--install-plugins` option with this
format:

```
gmibuildtool --install-plugins path-to-plugin-sources
```

For
example, if you needed to install `cordova-plugin-email-composer`, you would use a
command like this with the path to the repository:

```
gmibuildtool --install-plugins ~/w/github/cordova-plugin-email-composer/
```

Now, in order to get the list of all plugins available, the stock plugins and the ones you
installed in the GMI environment, use the `--list-plugins` option:

```
$ gmibuildtool --list-plugins
...
Calendar-PhoneGap-Plugin
GeneroTestPlugin
cordova-plugin-bluetoothle
cordova-plugin-device-motion
cordova-plugin-email-composer --> this is the installed one
cordova-plugin-file
cordova-plugin-media
...
```

For more usage examples, see [Cordova plugins](5117-cordova-plugins.md "This section describes how to use Cordova plugins.").

## Creating the GMI front-end for development purpose

A self-made GMI front-end can be created with the gmibuildtool command.
For more details, see [Mobile development mode](5083-mobile-development-mode.md "Set up a development environment to display app forms on a mobile front-end.").

## Specifying the target to build and deploy the iOS app

The gmibuildtool command can build and install iOS apps for the
simulator or for physical devices.

The build and/or install action is controlled by the `--device`
option:

- By default, when not specifying the `--device` option, a
  GMI.app directory is created for the simulator.
- When specifying the `--device simulator` option, the GMI.app
  directory is created.
- When specifying the `--device phone` option, the GMI.app
  directory and .ipa file are created.
- When specifying the `--device physical-device-name` option
  (with a real physical device name plugged on your Mac), the GMI.app directory
  and .ipa file are created.

By default, the generated GMI.app directory and .ipa
archive can be found in $PWD/build subdirectories; however, you can specify the
destination IPA file with the `--output` option.

## Specifying the GBC to be used

Use the `--gbc` option, to specify which GBC has to be bundled in the iOS app
package:

```
gmibuildtool ... --gbc gbc-archive ...
```

The gbc-archive parameter is the ZIP archive of the GBC front-end, to
be created as described in the Create a runtime zip topic in the Genero Browser Client User Guide.

The parameter for the `--gbc` option can also be a regular (unzipped) GBC
directory.

If the `--gbc` option is not used, gmibuildtool will use the
GBC found in [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode."), and if that
environment variable is not defined, it defaults to [FGLDIR](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")/web\_utilities/gbc/gbc.

## Elements used to build the iOS app

The gmibuildtool command builds the iOS app package from the
following:

- The GMI binary archive, containing the GMI front-end and the FGLGWS runtime system library
  (these files are provided in the fjs-gmi-\*.zip archive),
  > **Important:**
  >
  > When re-installing a new GMI archive, remove all "build" directories created
  > by the gmibuildtool.
- The GBC to be used (`--gbc` option),
- The compiled application program and resource files (.42m,
  .42f, etc) - these application program files must include a
  main.42m or main.42r module,
- The display name of the app (`--app-name` option),
- The version of the app (-`-app-version` option),
- The debug or release mode (`--mode` option),
- The certificate (to sign the app) (`--certificate` option),
- The bundle Identifier (`--bundle-id` option),
- The app provisioning profile (.mobileprovision file)
  (`--provisioning` option),
- If needed by the app store, the updated `PrivacyInfo.xcprivacy` privacy
  manifest,
- iOS app specific resources:
  - App icons (`--icons` option),
  - Launch images (`--launch-images` option) or launch storyboard
    file (`--storyboard` option).

For a complete description of command options, see [gmibuildtool](5110-gmibuildtool.md "The gmibuildtool is a utility to create and test applications for an iOS devices.").

## Default build directory structure

For convenience, the buildtool supports a default directory structure to find all
files required to build the app:

```
top-dir
|
|-- main.42m and other program files, as described in Directory structure for GMI apps
|
|-- gmi
|   |
|   |-- Info.plist
|   |
|   |-- PrivacyInfo.xcprivacy   (custom manifest, if needed)
|   |
|   |-- LaunchScreen.storyboard  or  (default launch images)
|   |                              -- Default@2x.png
|   |                              -- Default-568h@2x.png
|   |                              -- Default-Landscape.png
|   |                              -- Default-Landscape-667h@2x.png
|   |                              -- Default-Landscape-736h@3x.png
|   |                              -- Default-Landscape@2x.png
|   |                              -- Default-Portrait.png
|   |                              -- Default-Portrait-736h@3x.png
|   |                              -- Default-Portrait-667h@2x.png
|   |                              -- Default-Portrait@2x.png
|   |   ...
|   |-- icon_29x29.png
|   |-- icon_40x40.png
|   |-- icon_57x57.png
|   |-- icon_58x58.png
|   |-- icon_72x72.png
|   |-- icon_76x76.png
|   |-- icon_80x80.png
|   |-- icon_120x120.png
|   |-- icon_152x152.png
|   |   ...
```

In the above directory structure:

1. top-dir is the top directory of the
   default structure. It will typically hold your application program files. A
   different program files directory can be specified with the
   `--program-files` option.
2. top-dir/gmi is the default directory
   containing the app resource files such as icons:
   1. Info.plist is the Information Property List File that
      will be used to build the app. Some properties will be overwritten by
      gmibuildtool options like `--app-name` and
      `--app-version`.
   2. PrivacyInfo.xcprivacy is the custom privacy manifest, only needed if the
      default privacy manifest provided in the GMI archive is not sufficient.
   3. Provide either a launch screen storyboard or default launch images:
      - LaunchScreen.storyboard is the default storyboard file
        for the app launch screen. This file can be specified with the
        gmibuildtool `--storyboard` option.
      - Default-\*.png are the app launch image files. The directory
        to find launch images can be specified with the gmibuildtool
        `--launch-images` option.
   4. icon\_\*.png are the app icon files. The directory to find
      icons can be specified with the gmibuildtool
      `--icons` option.

## Debug and release versions

iOS apps can be generated in a debug or release version. Release version are prepared
for distribution on the App Store, while debug versions are used in development.

In debug mode, the app installed on the device can listen on the debug TCP port to
allow [fgldb -m
connections](5101-debugging-a-mobile-app.md "Different solutions are available to debug a mobile app."), after enabling the debug port in the app settings.

Debug or release mode must be specified in the command line with the `--mode
debug` or `--mode release` option. Additionally, if you want to
deploy on a physical device, you need to use a provisioning profile corresponding to the
debug or release mode:

- In debug mode, the certificate must be a development certificate.
- In release mode, the certificate must be a distribution certificate.

## Defining the app version and build number

Apple distinguishes the app version number of a bundle (visible to the end user),
from the build version number of a bundle (called a release version number in Apple
docs).

You specify the app version number with the `--app-version` option of
the gmibuildtool command. This option sets the `CFBundleVersion`
property of the Info.plist file), and must match the version specified in iTunes® Connect.

In order to distinguish multiple builds (Apple's term is "releases") of the same app
version number, define the build version number of your app with the `--build-number`
option. This option sets the `CFBundleShortVersionString` property of the
Info.plist file. For a given app version, you need to increase this build
number, to be able to upload a new binary on iTunes Connect.

If you do not specify the `--build-number` option, the build version number
defaults to the app version specified with the `--app-version` option.

## Defining app properties in the ./gmi/Info.plist file

iOS app are created with a set of properties that are essential configuration information
for a bundled executable. These properties are defined in the "Information Property List File",
an XML formatted file named Info.plist by convention.

Most important `Info.plist` properties are defined with
gmibuildtool options such as `--app-name` and
`--app-version`. However, you may need to define other properties that are out
of the scope of the buildtool. For example: background modes, device capabilities, screen
orientations, permanent wifi, and so on.

In order to define specific app properties, create an Info.plist
file in the top-dir/gmi directory before executing the
gmibuildtool. Properties covered by the buildtool will be overwritten, while
any other property defined in the top-dir/gmi/Info.plist
file will be left untouched.

For more details about the Info.plist file structure, see Apple developer
site page about [Information Property List File](http://developer.apple.com/library/ios/documentation/General/Reference/InfoPlistKeyReference/Articles/AboutInformationPropertyListFiles.html).

## Customize ./gmi/PrivacyInfo.xcprivacy file

iOS apps that don’t describe their use of required reason API in their privacy manifest file
aren’t accepted by App Store Connect.

GMI provides a default PrivacyInfo.xcprivacy privacy manifest file, with
settings that should be sufficient for most apps. This default file will be automatically included
during the build process.

If the default privacy settings are not adapted to your app, make a copy of the
PrivacyInfo.xcprivacy provided in the GMI archive, add the required settings,
and put this customized privacy manifest file under the
top-dir/gmi directory, beside the
Info.plist file.

For more details about editing the private manifest file, see [Apple development documentation](https://developer.apple.com/documentation/bundleresources/privacy_manifest_files/describing_use_of_required_reason_api?language=objc).

## Building an iOS app with gmibuildtool

Follow the next steps to setup a GMI app build directory in order to create an iOS app,
based on the default directory structure:

1. Create the root distribution directory (top-dir)
2. Copy compiled program files (.42m, .42f,
   fglprofile, application images, web component files, etc) under
   top-dir.
3. Copy the default English .42s compiled string resource file under
   top-dir.
4. Create non-English language directories (fr, ge, and so on) under
   top-dir and copy the corresponding .42s
   files.
5. Copy default application data files (database file for ex) under
   top-dir.
6. Create the top-dir/gmi directory.
7. Copy iOS app resources (icons, launch screen, storyboard) under
   top-dir/gmi.
8. If needed, create an top-dir/gmi/Info.plist
   file, to define specific iOS app properties.

Once the build directory is prepared, issue the following commands:

```
$ cd top-dir
$ gmibuildtool \
    --install yes \
    --output myapp.ipa \
    --device phone \
    --app-name "My App" \
    --app-version "v3.1.6" \
    --bundle-id "com.example.mycompany.myapp" \
    --mode release \
    --certificate "iPhone Developer" \
    --provisioning "~/Downloads/myapp.mobileprovision"
```

This gmibuildtool command will build the .ipa archive and
since the `--install` option and `--device` option are specified, it
will deploy the app on the device. Once created, the .ipa file can then be
installed without a complete build, by using the `--deploy` option.

## Installing the app on a device

In order to install the .ipa on a device after creation, you need to use the
[`--install yes`](5110-gmibuildtool.md)
option:

```
$ gmibuildtool \
    --install yes \
    --output myapp.ipa \
    --device phone
    ...
```

After creating the .ipa file, if you only want to install the app on a
device simulator or physical device, you can use the [`--deploy
ipa-file`](5110-gmibuildtool.md) option, in conjunction with the
`--bundle-id` option:

```
$ gmibuildtool \
    --deploy myapp.ipa \
    --bundle-id "com.example.mycompany.myapp" \
    --device "simulator"
```

Apps can be updated, or un-installed and re-installed. When updating the app, the "Documents" dir
(`os.Path.pwd()`) is left intact. With a physical device, the app is always updated
(if needed, you can uninstall the app manually). With a simulator, the app is un-installed and
re-installed by default. If required, use the [`--update`](5110-gmibuildtool.md) option to only update the app on a device simulator.

## Building a GMI app with C extensions or custom front calls

In order to create an iOS app using C extensions written in Objective-C as in [Implementing C-Extensions for GMI](../14_extending-the-language/2717-implementing-c-extensions-for-gmi.md "This section describes how to program C-Extensions for the GMI VM."), proceed as follows (the same technique can be used to build
apps that include [custom front calls](../14_extending-the-language/2721-implement-front-call-modules-for-gmi.md "Custom front call modules for the iOS front-end are implemented by using the API for GMI front calls in Objective-C.")):

1. Build a static library from your Objective-C sources, by using the `staticlib`
   target of the $GMIDIR/lib/Makefile-gmi generic makefile (specify the library
   filename with the `USER_LIBNAME` variable). The `staticlib` makefile
   target will produce a .a library file, by using all .m and
   .c files found in the current directory. For
   example:

   ```
   $ make -f $GMIDIR/lib/Makefile-gmi USER_LIBNAME=mylib.a staticlib
   ```
2. When building the app, specify the additional libraries with the
   `--extension-libs` option. For
   example:

   ```
   $ gmibuildtool ... --extension-libs "-lz libBPush.a mylib.a" ...
   ```

Regardless of whether the `--extension-libs` option is used, the
gmibuildtool looks to see if ./gmi/\*.a exists. If they
exist, it adds these static libs to the link list.

If a file called ./gmi/link\_flags.sh exists, this file is read by
gmibuildtool to set additional link flags. For example:

```
#set LINK_FLAGS to add additional system libs or frameworks
LINK_FLAGS="-bz2 -framework MapKit"
```

In most cases the \*.a files will be sufficient, for system frameworks like MapKit, as long as the
following hint is specified in your .m source file:

```
/* this avoids using -framework MapKit and instructs the linker to link with MapKit*/
@ import MapKit;
```

Only pure C libs such as libbz2.a actually need the
link\_flags.sh file, if no `--extension-libs` option is used and
one of the static extension libs needs a system C library.

For complete examples, see $GMIDIR/demo/MobileDemo/userextension and
$GMIDIR/demo/MobileDemo/userfrontcall

## Related links

**Related concepts**  

[Mobile development mode](5083-mobile-development-mode.md "Set up a development environment to display app forms on a mobile front-end.")
