---
title: "Directory structure for GMA apps"
source: "fgl-topics/c_fgl_gma_app_structure.html"
breadcrumb: "Mobile applications > Deploying mobile apps > Deploying mobile apps on Android™ devices > Directory structure for GMA apps"
type: "concept"
description: "Platform-specific rules need to be considered when deploying on Android devices (GMA)."
---

# Directory structure for GMA apps

> Platform-specific rules need to be considered when deploying on Android™ devices (GMA).

## The application sandbox

On Android devices, applications are
deployed in an application sandbox. The application can access and store data outside of its space,
but then the data is also accessible by the other applications.

## Directory structure for a GMA application

Inside its application sandbox, an Android
app uses the following directory structure:

```
appdir/
|-- main.42m                            --
|-- *.42m                                 |
|-- *.42f                                 |
|-- fglprofile                            |
|   ...                                   |
|-- defaults/*.42s                        |
|-- de/                                   | Program files
|   |-- *.42s                             |
|-- fr/                                   |
|   |-- *.42s                             |
|-- zh/                                   |
|   |-- *.42s                             |
|   ...                                   |
|-- ... other resource files/dirs ...     |
|   ...                                   |
|-- webcomponents                         |
|   |-- component-type                    |
|       |-- component-type.html           |
|       |-- other-web-comp-resource       |
|   ...                                  --

|-- appdata/
    |-- ... writable app files ...
  
tmpdir/
|-- ... temporary files ...
```

## Program files directory (appdir)

Application program files (.42m, .42f, and so on) need
to be deployed in the appdir application base directory.

The program files directory can be found in programs with the `base.Application.getProgramDir` method.
> **Important:**
>
> On Android, the program files directory returned
> by the `base.Application.getProgramDir()` method is the same directory as the default
> working directory, returned by `os.Path.pwd()`.

The [FGLAPPDIR](../07_configuration/0519-fglappdir.md "Contains the path to the application directory when executing on a mobile device.") environment variable is
automatically set to the `appdir` directory.

## Program name (MAIN)

When deploying on mobile devices, the name of the program file must be
main.42m or main.42r.

When using the command-line app build scripts, the name of the program file must be
main.42?. When using Genero Studio, the packaging script
takes care of renaming this file, if you have not named it main.

As with other program files, the "MAIN" module must be located under the
appdir application program directory.

## Working directory

On Android devices, the default current
working directory is the appdir directory, and can be used for
writable files.

The current working directory can be found in programs with the `os.Path.pwd` method.

Files that need to be writable (such as SQLite database files) can be created directly under the
appdir directory. However, to better organize application files,
create sub-directories such as appdir/appdata,
keeping original files directly under the appdir directory.
For example, create the application database under os.Path.pwd() || "/database".

## Temporary directory (`tmpdir`)

A temporary directory is available for the application.

In order to find the temporary directory for the app, use the [standard.feInfo](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.") front call, with
the "`dataDirectory`" parameter.

To create a temporary file name, use the [`os.Path.makeTempName()`](../15_library-reference/3730-os-path-maketempname.md "Generates a new file path to be used to create a temporary file or directory.") method.

## Language directories for localized strings

When the app starts, the appropriate .42s string files will be loaded from the
directory corresponding to the current language settings of the mobile device. String files to be loaded
can be defined in the app's fglprofile, or you can use the main program name to
avoid fglprofile settings.

For each language supported by your application, a directory must exist under
appdir, with a name including the locale codes. Default
string files (in English for example) can be provided under
appdir/defaults, in case the regional settings of the device
do not match one of the locale directories of the app, otherwise the application will stop with
error [-8006](../15_library-reference/4483-genero-bdl-errors.md).

For example:

```
appdir/defaults/mystrings.42s
appdir/fr/mystrings.42s
appdir/de/mystrings.42s
```

For more details, see [Localized string files on mobile devices](../09_advanced-features/0909-loading-localized-strings-at-runtime.md).

## Deploying a custom fglprofile file

If you need to set fglprofile entries for your mobile application, create a file with the name
fglprofile, and deploy it under the appdir directory, along
with the other program files.

See [Understanding FGLPROFILE](../07_configuration/0484-understanding-fglprofile.md "The runtime system uses one or more configuration files in which you can define options and parameters to change the behavior of the programs.") for more details about fglprofile settings.

## Creating the initial database file

When a mobile application starts for the first time, it typically creates a new database, or copies
an existing database template file from the file directory (`base.Application.getProgramDir`) to the working directory (`os.Path.pwd`).

It is recommended that different database file names are used for the original and final
application database, as folders pointed to by `base.Application.getProgramDir()` and
`os.Path.pwd()` can be the same on Android devices.

For more details about database creation on mobile devices, see [Creating a database from programs](../10_sql-support/1003-creating-a-database-from-programs.md "Creating a database from within a program requires special consideration.").
