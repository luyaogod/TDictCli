---
title: "Create a program directory"
source: "fgl-topics/c_fgl_gwa_programdir.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero > Create a program directory"
type: "concept"
---

# Create a program directory

> To build a GWA application, you must create a directory.

In the directory referred to as the programdir, copy your
compiled program files (such as .42m, .42f,
fglprofile, and so on) and resource files (including images, database files,
and web components) required by the Genero Web application.

In the following tree view, you can see the expected structure of this program directory. This
directory may have a structure similar to that described for mobile applications in [Directory structure for GMA apps](../17_mobile-applications/5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).").

```
programdir/
|-- main.42m                            --
|-- *.42m                                 |
|-- *.42f                                 |
|-- fglprofile                            |
|   ...                                   |Program files
|-- *.42s                                 |
|-- ... other resource files/dirs ...     |
|-- gwa                                   |
|   |-- gwa.webmanifest                   |Customized files
|   |-- index_tpl.html                    |
|   |-- locales            |
|       |-- de/                           | 
|           |-- *.42s                     |Local files
|       |-- fr/                           |
|           |-- *.42s                     |
|-- webcomponents                         |
|   |-- component-type                    |
|       |-- component-type.html           |
|       |-- other-web-comp-resource        |
|   ...                                 --
```

Within the programdir directory, you should include the
following essential files:

## Program main module

The main Genero module file for launching your GWA is typically named
main.42m.
> **Tip:**
>
> You can specify the desired module as the main entry point by using the [gwabuildtool
> --main-module](5150-gwabuildtool.md) command during the application build process.

For additional information and other available options, go to [Specify the start-up module](5131-specify-the-start-up-module.md "Specify the module you want to use as the main entry point.")

## fglprofile

The GWA app reads the default fglprofile provided in
$FGLDIR/etc/fglprofile. If using a custom fglprofile file
for you GWA app, it must be located in the root of the
programdir directory, beside the main program module. For
details on what environment variables you may need to set for your GWA application, go to [Set environment with fglprofile](5128-set-environment-with-fglprofile.md "Configure environment settings with FGLPROFILE entries.").

## Web Application Manifest

Every GWA application must include a web application manifest file when deployed. If the
`gwabuildtool` does not find
programdir/gwa/gwa.webmanifest, it creates a default
manifest file with Genero icons in the distribution directory. When you are using a customized
manifest file, it must be named gwa.webmanifest and be located in the
programdir/gwa directory of your project. For more details,
go to [gwa.webmanifest file](5153-gwa-webmanifest-file.md "Example of a customized gwa.webmanifest file, providing information about a Genero Web Application (GWA).").

## Language directories for localized strings

For each language supported by your application, a directory must exist under
programdir/gwa, with a name including the locale codes.

For more details on setting locale and customizing translation text, go to [Internationalize your app](5135-internationalize-your-app.md "GWA applications can be translated and modified for international markets.").

## Related links

**Related concepts**  

[Set environment with fglprofile](5128-set-environment-with-fglprofile.md "Configure environment settings with FGLPROFILE entries.")

[Specify the GBC](5129-specify-the-gbc.md "A GBC must be bundled with the GWA.")

[Package web components](5130-package-web-components.md "Web components must be packaged with the GWA application because they are preloaded and must also be available offline.")

[Specify the start-up module](5131-specify-the-start-up-module.md "Specify the module you want to use as the main entry point.")

[Customizing GWA apps](5132-customizing-gwa-apps.md "This section describes ways you may customize your GWA apps, such as adding a favicon, customizing the index page, and internationalizing your app.")
