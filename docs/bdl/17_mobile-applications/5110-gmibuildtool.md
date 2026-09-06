---
title: "gmibuildtool"
source: "fgl-topics/c_fgl_tools_gmi_app_builder.html"
breadcrumb: "Mobile applications > Deploying mobile apps > Deploying mobile apps on iOS devices > gmibuildtool"
type: "concept"
---

# gmibuildtool

> The gmibuildtool is a utility to create and test applications for an iOS devices.

## Syntax

```
gmibuildtool [options]
```

1. options are described in Options.

> **Tip:**
>
> **Options with arguments: equals (=) or space?**
>
> In options with arguments,
> you can use both equals `(=)` or space between the option and the argument. For
> example, you can invoke the verbose option
> with:
>
> ```
> gmibuildtool --verbose=yes
> ```
>
> or
>
> ```
> gmibuildtool --verbose yes
> ```
>
> In
> Options, options that take arguments are shown with an
> equals sign, but in these options you can also use a space instead.

## Options

| Option | Description |
| --- | --- |
| `--app-name=application-name` | Display name of the mobile app.This option can be specified to define the display name of the app, it sets the `CFBundleDisplayName` property in the Info.plist file.If not specified, the name defaults to "Noname". |
| `--app-version=application-version` | Defines app version visible to the users on the App Store.This option is mandatory and sets `CFBundleVersion` properties in the Info.plist file.If the `--build-number` option is not used, `--app-version` will also set the both the `CFBundleShortVersionString` property.In iTunes® Connect, you define the version of your app, that must match the `CFBundleVersion` property in the Info.plist file of the app. If these versions do not match, the app cannot be published. Once the app is visible on App Store, the version specified in iTunes Connect shows up in the "Version" section of the application page.The recommendation for the app version number is that it is a string comprised of three period-separated integers. For example: "1.4.2" |
| `--bundle-id=bundle-identifier` | Defines the Bundle Identifier (a.k.a. App Id) for the app.This option is mandatory and sets the `CFBundleIdentifier` property in the Info.plist file.A bundle identifier is the unique identifier of your app, to let iOS recognize new app versions. When developing for the simulator, you can choose your own identifier. When creating an application for the App Store, the bundle identifier must be registered with Apple.If not specified, the name defaults to "noname" (for prototyping). |
| `--build-cordova=cordova-plugin-names` | Defines Cordova plugins to be embedded in the app package.When specifying multiple cordova plugins, use the comma (`,`) as separator.The name of the plugin must match the Git repository name. It is case-sensitive.To get the list of available Cordova plugins, use the gmibuildtool --list-plugins command. |
| `--build-number=build-number` | Defines the build number used to upload a new binary of the same app version.This option must be used to distinguish different builds for the same app version. It sets the `CFBundleShortVersionString` property in the Info.plist file.The build number needs to be incremented in order to upload a new binary version of the same app version in iTunes Connect.If this option is not used, the build number defaults to the version specified with the `--app-version` option.The build number is a string comprised of three period-separated integers. For example: "1.4.2" |
| `--certificate=identity` | Name of a certificate to sign the app.This option is mandatory to build apps for a physical device or for the app store.The certificate can be found in the Keychain® access program, in the "Common Name" field of the certificate panel.The command security find-identity -v can be used to list all available certificates. |
| `--check-p-code` | This option checks if your installed `fglrun` version matches the version GMI was compiled with. |
| `--crypto={yes\|no}` | Enables GWS cryptographic APIs based on the OpenSSL library. When using this option, the OpenSSL library is embedded into the resulting .ipa file.The default is `yes`. |
| `--deploy=ipa-file` | Deploy an existing .ipa on a device simulator or physical device. This option must be used in conjunction with the `--bundle-id` option. |
| `--device={device-name\|simulator\|phone}` | Defines the name of a device or simulator.By default, when not specifying the `--device` option, a GMI.app directory is created for the simulator.When specifying the `--device simulator` option, the GMI.app directory is created.When specifying the `--device phone` option, the GMI.app directory and .ipa file are created.When specifying the `--device physical-device-name` option (with a real physical device name plugged on your Mac), the GMI.app directory and .ipa file are created.Use the xcrun xctrace list devices Xcode® command to find the list of available devices (simulators or connected devices).Use the `--install yes` option to install the app on a device. |
| `--extension-libs`=`library[,...]` | Specify the libraries to use when compiling and linking the app.This option is used when you want to provide your own C extension or custom front calls. This option can take a comma-separated list of libraries. |
| `--gbc gbc-dir-or-zip` | Defines the GBC to be included in the app packages and used at runtime. The parameter can be a GBC ZIP archive or a GBC directory. If this option is not specified, default is [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode."), and if the FGLGBCDIR is not defined, default is [FGLDIR](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")/web\_utilities/gbc/gbc.For more details about GBC archive creation, see the Create a runtime zip topic in the Genero Browser Client User Guide. |
| `-h` or `--help` | Displays options for the tool. |
| `--icons=icons-dir` | Provides the directory where the application icons are located.By default, the application icons directory is current-working-dir/gmi.The name of the app icon files must be: `icon_57x57.png, icon_72x72.png, icon_29x29.png, icon_40x40.png, icon_120x120.png, icon_152x152.png, icon_58x58.png, icon_76x76.png, icon_80x80.png` |
| `--install={yes\|no}` | The `--install` option controls if the app is installed on the simulator or device.See also [`--update`](5110-gmibuildtool.md). |
| `--install-plugins=github-url[,...]` | This option installs additional plugins in the GMI installation directory. This option can take a comma-separated list of repositories. For an example, see [Manage GMI plugins](5109-building-ios-apps-with-genero.md). |
| `--launch-images=launch-images-dir` | The directory where launch images are located.By default, the launch images directory is current-working-dir/gmi.This option is ignored if the `--storyboard` option is provided.The name of the image files must be: `Default.png, Default@2x.png Default-568h@2x.png, Default-Portrait-667h@2x.png Default-Landscape-667h@2x.png, Default-Portrait-736h@3x.png, Default-Landscape-736h@3x.png, Default-Portrait.png, Default-Landscape.png Default-Portrait@2x.png, Default-Landscape@2x.png`.Each filename corresponds to a device type (you may not need to provide all files if you target only recent iOS devices), see Apple® Developer documentation for more details about launch images. |
| `-l` or `--list-devices` | This option lists the Mac® and IOS® devices, and simulators.**Note:**It lists Mac because also on Mac you could have IOS apps compiled for iPad. |
| `--list-installed-plugins` | This option lists the Cordova plugins you install. For an example, see [Manage GMI plugins](5109-building-ios-apps-with-genero.md). |
| `-p` or `--list-plugins` | This option lists the shipped plugins and additional plugins installed in the GMI installation directory. For an example, see [Manage GMI plugins](5109-building-ios-apps-with-genero.md). |
| `--list-stock-plugins` | This option lists plugins, such as `Calendar-PhoneGap-Plugin`, `GeneroTestPlugin`, `cordova-plugin-bluetoothle`, and so on, provided in the fjs-fglgmixxx package. For an example, see [Manage GMI plugins](5109-building-ios-apps-with-genero.md). |
| `--mode={debug\|release}` | Controls the debug or release mode for the app.By default, the mode is `debug`.Note that the provisioning profile must correspond:`--mode debug`: Development provisioning profile.`--mode release`: Distribution provisioning profile. |
| `--output=ipa-file-name` | Path to output IPA and APP files to be generated.By default, a "build" directory is created, with subdirectories containing the .ipa and .app files.An IPA file is created when building an application for a physical device and the App Store. The IPA file is not needed and will not be created when building for the simulator. |
| `--plugin-info=cordova-plugin-name` | This option gives more detailed information about a cordova plugin, such as the version, the repository, and so on. May be used to get more information to support troubleshooting. |
| `--program-files=program-dir` | Path to Genero BDL program files (.42m, .42f, etc).By default, the program files directory is the current work directory.Following files are automaticlly excluded: \*.4gl, \*.per, \*.msg, \*.str, \*.sch, [Mm]akefile, \*.42d, [Mm]akefile, \*.[chdmo], \*.xib, build/ (the build directory), gmi/ (this folder is the default location of LaunchScreens and AppIcons).If the file gmiignore exists, then this file contains additional files to be ignored. |
| `--provisioning={provisioning-file\|latest}` | Path to the provisioning profile (.mobileprovision).The provisioning profile is mandatory to build apps for a physical device or for the app store.Provisioning profiles can be found in $HOME/Library/MobileDevice/Provisioning\ Profiles/ |
| `--storyboard=storyboard-file` | Path to the storyboard file, to get a splash screen to be displayed when the app starts.This file is an alternative for Launch Screens (`--launch-images` option). This option is mandatory if you do not provide launch images with the `--launch-images` option.The default storyboard is showing an empty navigation bar and an empty toolbar. If the storyboard references images, gmibuildtool searches for the images in the same directory the storyboard is in, and bundles the images with the application. |
| `--uninstall-plugin=cordova-plugin-name` | This option uninstalls the specified plugin. |
| `--update` | When installing with `--install yes` option on a simulator, the app is first un-installed and then re-installed. Use `--update` to only update and keep the "Documents" dir (`os.Path.pwd()`). When installing on a physical device, the app is always updated (needs manual un-install to cleanup) |
| `--verbose={yes\|no}` | Enable the verbose mode. |
| `-v`, `-V` or `--version` | Displays version information. |

## Related links

**Related concepts**  

[Building iOS apps with Genero](5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.")
