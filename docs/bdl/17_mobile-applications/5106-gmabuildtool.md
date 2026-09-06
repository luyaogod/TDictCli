---
title: "gmabuildtool"
source: "fgl-topics/c_fgl_tools_gma_app_builder.html"
breadcrumb: "Mobile applications > Deploying mobile apps > Deploying mobile apps on Android™ devices > gmabuildtool"
type: "concept"
description: "The gmabuildtool is a utility to create app packages for an Android device."
---

# gmabuildtool

> The gmabuildtool is a utility to create app packages for an Android™ device.

## Syntax

```
gmabuildtool command [option [...]] [@argfile [...]]
```

1. command can be one of the following:
   - `updatesdk`: updates the Android SDK, to download packages required by GMA.
   - `scaffold`: manages scaffold archives.
   - `build`: builds app packages (.aab and .apk).
   - `test`: deploys an APK package and launches the app on the device or
     emulator.
2. option can be a general or command-specific option, as described in Options.
3. argfile defines a file that contains an optional command and a list of
   options. The options file must use the following format:

   ```
   [command]
   option-name option-value
   [...]
   ```

   Several
   `@argfile` files can be provided.

## Options

| Option | Short option | Description |
| --- | --- | --- |
| `--android-sdk path` | `-as` | The path to the Android SDK installation directory.Default is ANDROID\_HOME. |
| `--help` | `-h` | Display options for the tool.Prefix with the command name to get command options details. |
| `--java-home path` | `-jh` | Java home path.Default is JAVA\_HOME. |
| Verbose mode (no long option name) | `-v` / `-vv` / `-vvv` | Defines the verbose mode. |
| `--version` | `-V` | Displays version information. |

| Option | Short option | Description |
| --- | --- | --- |
| `--accept-licenses` | `-al` | Silently accept Android SDK licenses when the Android SDK is updated. |
| `--no-install-extras` | `-uN` | Avoid installation of extra SDK modules. |
| `--no-interactions` | `-nI` | Silently accepts prompts when extras are installed during Android SDK installation, by answering yes to all questions asked during the installation process.This option also answers yes to prompts to accept licenses. `gmabuildtool updatesdk --no-interactions` is the same as the command `gmabuildtool updatesdk --no-interactions --accept licenses` |
| `--proxy-host host` | `-ph` | Defines the proxy host. |
| `--proxy-port port` | `-pp` | Defines the proxy port. |

| Option | Short option | Description |
| --- | --- | --- |
| `--install-plugins plugin-list` | `-ip` | Install the specified plugins in the scaffold archive.The plugin-list must be a comma-separated list of plugins. |
| `--list-plugins` | `-lp` | List the plugins available in the scaffold archive. |

| Option | Short option | Description |
| --- | --- | --- |
| `--clean` | `-c` | Cleans the intermediate build files before a rebuild.Use the `--clean` option if the previous build was interrupted or has failed.The `--clean` option does not remove and replace the scaffold. |
| `--generate-default-manifest` | `-gdm` | Generate a default AndroidManifest.xml file in the build directory (`--build-apk-outputs` option). This option takes into account manifest-related options such as `--build-app-permissions`. The build process stops at the manifest file generation. |
| `--custom-manifest-file path` | `-cmf` | Manifest file to be used to build the app packages. The basename of the provided file path must be `AndroidManifest.xml`. If this option is not used, `gmabuildtool` will generate a new manifest file. |
| `--main-app-path path` | `-map` | Defines the path to the main module of the application (42m, 42r, or xcf file), and implicitly defines the path to the application program files (42m, 42f, 42s, etc)This option is mandatory.The gmabuildtool will create a zip file for the app by taking all files from the parent directory of this path.If the `-bpn` option is not used, the default application name and package name will be constructed from the name of the parent directory of this path. |
| `--root-path path` | -rp | Defines the path to the directory containing Android app icons, and used temporary build files, also known as "root-dir".The resources found in the defaults directories such as root-dir/icons/\*.png icons will be zipped and bundled inside app packages.If not specified, defaults to the current working directory.The path defined by this option is used as base directory for other options such as `--build-project` and application icon resources options. |
| `--custom-foreground-service-type service-type` | `-cfst` | Defines the `foregroundServiceType` manifest parameter.Repeat this option to specify multiple forground service types.If this option is not specified, the default is always `specialUse`. With this service type, `--background-explanation` option is required to describe the use cases.When using `shortService`, the timeout is set to 3 minutes, starting when the app enters in background mode.This option cannot be used when building in debug mode with the `--build-mode debug`. |
| `--background-explanation explanation` | `-be` | Defines the `foregroundServiceType` as `"specialUse"`, by providing an explanation that needs to be approved by Google for app publication on the store. |
| `--build-app-icon-hdpi path` | `-bih` | Defines the path to application icon in hdpi.Default is root-dir/icons/ic\_app\_hdpi.png, where root-dir is defined by the `--root-path` option. |
| `--build-app-icon-mdpi path` | `-bim` | Defines the path to application icon in mdpi.Default is root-dir/icons/ic\_app\_mdpi.png, where root-dir is defined by the `--root-path` option. |
| `--build-app-icon-xhdpi path` | `-bixh` | Defines the path to application icon in xhdpi.Default is root-dir/icons/ic\_app\_xhdpi.png, where root-dir is defined by the `--root-path` option. |
| `--build-app-icon-xxhdpi path` | `-bixxh` | Defines the path to application icon in xxhdpi.Default is root-dir/icons/ic\_app\_xxhdpi.png, where root-dir is defined by the `--root-path` option. |
| `--build-app-package-name name` | `-bpn` | App package name.The package name implicitly defines the application name.It is recommended to format the package name as "`com.organization-name.app-name`".If not specified, the application package name defaults to `com.example.current-working-directory` |
| `--build-app-permissions permissions` | `-ba` | Android application permissions.The list of permissions is provided as a comma separated list of `android.permission.*` identifiers.For more details, see [Android permissions](5105-building-android-apps-with-genero.md). |
| `--build-app-version-code version-code` | `-bvc` | Application version code.For example: 100915The value of this option must be an integer (do not use decimal numbers). |
| `--build-app-version-name version-name` | `-bvn` | Application version name.For example: 10.09.15This will be the actual app version visible on devices. |
| `--build-apk-outputs path` | `-bo` | Defines the destination folder where the app packages must be created. |
| `--output-format package-type` | `-of` | Defines the type of app package format.Values can be:`aab`: Create an AAB package.`apk`: Create an APK package.`all`: Create all type of packages.Default is `all`. |
| `--build-cordova cordova-plugin-names` | `-bco` | Defines Cordova plugins to be embedded in the app package.When specifying multiple cordova plugins, use the comma (`,`) as separator.The name of the plugin must match the Git repository name. It is case-sensitive.To get the list of available Cordova plugins, use the gmabuildtool scaffold --list-plugins command.For further information, see [Cordova plugins](5117-cordova-plugins.md "This section describes how to use Cordova plugins."). |
| `--build-gbc-runtime gbc-dir-or-zip` | `-bgr` | Defines the GBC to be included in the app packages and used at runtime. The parameter can be a GBC ZIP archive or a GBC directory. If this option is not specified, default is [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode."), and if the FGLGBCDIR is not defined, default is [FGLDIR](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language.")/web\_utilities/gbc/gbc.For more details about GBC archive creation, see the Create a runtime zip topic in the Genero Browser Client User Guide. |
| `--build-signing-alias alias` | `-bsa` | App signing alias.This is the alias provided to the keystore utility to build the keystore file to sign the app.This option is mandatory if `-bsks` is used. When no signing option is used, default is "androiduserkey". |
| `--build-signing-keypass keypass` | `-bsk` | App signing keypass.Specifies the password used to protect the private key of the keystore entry addressed by the alias specified in the `--build-signing-alias` option.This option is mandatory if `-bsks` is used. When no signing option is used, default is "android". |
| `--build-signing-storepass storepass` | `-bss` | App signing storepass.Specifies the password that is required to access the keystore.This option is mandatory if `-bsks` is used. When no signing option is used, default is "android". |
| `--build-signing-keystore path` | `-bsks` | App signing keystore file path.This is the path to the keystore file generated by the keystore utility to sign the app.Default is `$HOME/.android/debug.keystore`. |
| `--build-mode {release\|debug}` | `-bm` | Package build mode, to build a release version or a development/debug version.Default: `release` |
| `--build-output-apk-name name` | `-ban` | Defines the prefix for the app packages names.By default, this prefix is "app".The filename of the app package is formed from:The filename prefix defined by the `--build-output-apk-name` option (by default, "app"),When building a debug version, the `-debug` suffix,The `.aab` or `.apk` file extension.For example, if the filename prefix is `MyApp` and is a debug package, the resulting package filenames will be: MyApp-debug.aab and MyApp-debug.apk. |
| `--build-project path` | `-bp` | Defines the path to the directory where the scaffold zip file will be extracted. Default is root-dir/project, where root-dir is defined by the `--root-path` option.The status of the target directory determines whether the files are extracted:If the directory is empty, the scaffold files are extracted in root-dir/project/scaffold.If the directory already has a valid scaffold installation, files are not extracted. If you need to replace the existing scaffold, you must clear the directory first.If the directory contains some files but not a valid scaffold, a warning is raised, and the scaffold files are extracted.The path to the scaffold zip file does not need to be specified as it is found based on the location of the gmabuildtool jar. |
| `--build-status-icon-hdpi path` | `-bsh` | Status icon path for hdpi (high dots per inch) size.The default path is root-dir/icons/ic\_status\_hdpi.png, where root-dir is defined by the `--root-path` option.If this option is not specified, yet you provide default files under the gma directory named like those defined for the default path, your package will use these files. If you don't provide any status icon files, the default files are used. |
| `--build-status-icon-mdpi path` | `-bsm` | Status icon path for mdpi (medium dots per inch) size.The default path is root-dir/icons/ic\_status\_mdpi.png, where root-dir is defined by the `--root-path` option.If this option is not specified, yet you provide default files under the gma directory named like those defined for the default path, your package will use these files. If you don't provide any status icon files, the default files are used. |
| `--build-status-icon-xhdpi path` | `-bsxh` | Status icon path for xhdpi (extra-high dots per inch) size.The default path is root-dir/icons/ic\_status\_xhdpi.png, where root-dir is defined by the `--root-path` option.If this option is not specified, yet you provide default files under the gma directory named like those defined for the default path, your package will use these files. If you don't provide any status icon files, the default files are used. |
| `--build-with-firebase-analytics-collection` | `-bfac` | When doing push notifications with Firebase Could Messaging, enable user data collection for analytics (For more details, see [Analytics](https://firebase.google.com/docs/analytics) in [FCM](https://firebase.google.com/docs/cloud-messaging) documentation) |
| `--build-app-name app-name` | `-bn` | Defines the application name.**Important:**This option is deprecated: The application name is constructed from the package name (`-bpn` option)If not specified, the application name defaults to the current working directory. |
| `--build-app-genero-program` | `-bgp` | Defines the path to the application program files (.42m, .42f, etc)**Important:**This option is deprecated, use `-rp`.The contents of this directory will be zipped and bundled inside app packages. |
| `--build-app-genero-program-main` | `-bgpm` | Relative path to the main module of the application (can be .42m or .42r).Defaults to main.42m**Important:**This option is deprecated, use `-map`. |
| `--build-force-scaffold-update` | `-bfsu` | Force scaffold update / recreate from default gma scaffold directory.**Important:**This option is deprecated and has no effect. |
| `--build-quietly` | `-bq` | Avoid all questions asked (answer yes by default)**Important:**This option is deprecated and has no effect. |
| `--build-app-colors color-list` | `-bc` | Define the Android color theme for the app (Android 5.0+ / SDK 21+)**Important:**The `--build-app-colors` is deprecated, use [GBC theming to define application colors](5105-building-android-apps-with-genero.md).The value must be a comma-separated list of four hexadecimal RGB colors: `#F44336,#B71C1C,#EF9A9A,...`The position of the RGB value in the color list defines its purpose:The main color used in the app.The color used for the status bar and the navigation bar.The accent color used for widgets and table lines.The foreground color for the texts in the action bar.The text color for items in the whole application.The window background color.The background color of the bottom bar.By default, the color theme is the Genero purple color. |
| `--build-jarsigner-alias alias` | `-bja` | **Important:**This option is deprecated, use `-bsa` instead. |
| `--build-jarsigner-keypass keypass` | `-bjk` | **Important:**This option is deprecated, use `-bsk` instead. |
| `--build-jarsigner-storepass storepass` | `-bjs` | **Important:**This option is deprecated, use `-bss` instead. |
| `--build-jarsigner-keystore path` | `-bjks` | **Important:**This option is deprecated, use `-bsks` instead. |

| Option | Short option | Description |
| --- | --- | --- |
| `--test-apk path` | `-ta` | Path to the APK file to deploy and launch for testing. |

## Related links

**Related concepts**  

[Building Android apps with Genero](5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.")
