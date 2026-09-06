---
title: "Loading localized strings at runtime"
source: "fgl-topics/c_fgl_localized_strings_010.html"
breadcrumb: "Advanced features > Localization > Localized strings > Loading localized strings at runtime"
type: "concept"
---

# Loading localized strings at runtime

> Understand the rules for using localized strings at runtime.

## Distributing compiled string files

The compiled string files (.42s) must be distributed with the program files
in a directory specified in the [FGLRESOURCEPATH](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.") environment variable.

## Setting the correct locale

The locale (LANG/LC\_ALL) corresponding to the encoding used in
the .42s files must be set before starting the application. If the locale is
wrong, the strings will not be loaded properly.

## How does the runtime system load the strings?

The .42s compiled string resource files are loaded in the following order of
precedence:

1. The files defined in FGLPROFILE,
2. A file having the same name as the current program (myprog.42m loads myprog.42s),
3. A file with the name "default.42s".

Each .42s string resource file is searched in several directories, as
described in [the FGLRESOURCEPATH reference
topic](../07_configuration/0531-fglresourcepath.md "Defines a list of paths for program resource files.").

## String resource file sharing

Like .42m program pcode files, the .42s
string resource files are shared by all fglrun processes running
on the computer: The string file is loaded into memory with the `mmap`
operating system function.

## Defining a list of string files in FGLPROFILE

Specify a list of compiled string files with entries in the [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files") configuration file with the
`fglrun.localization` entries.

First, define the total number of files with:

```
fglrun.localization.file.count = integer
```

For each file, define the filename (with the .42s extension),
including an index number (start index must be 1):

```
fglrun.localization.file.index.name = "filename.42s"
```

Warning switches can be specified in FGLPROFILE.

If the text of a string is not found at runtime, the runtime system can show a
warning, for development purposes.

```
fglrun.localization.warnKeyNotFound = boolean
```

By default, this warning switch is disabled.

## What happens if a 42s string file is not found?

If the 42s string file was defined with
`fglrun.localization.*` FGLPROFILE entries, it is considered as mandatory, and the
runtime system will raise [error
-8006](../15_library-reference/4483-genero-bdl-errors.md) if the file is not found. If the progname.42s
and default.42s string files are not found, no error is raised, because these
are fallback string resource files.

## What happens if a string is not defined in a resource file?

If a localized string is not defined in one of the compiled string files, the runtime system
uses the string identifier as default text.

## What happens if a string is defined more that once?

When a localized string is defined in several compiled string files, the runtime system uses
the first string found.

For example, if the string "hello" is defined in program.42s as "hello
from program", and in default.42s as "hello from default", the runtime
system will use the text "hello from program".

## Organizing .42s resource files in distribution directories

A set of .42s files using the same language and codeset is typically
copied in a distribution directory with a name identifying the locale.

For example:

```
/opt/app/resource/strings/en_US.iso8859-15  -- English strings in iso8859-15 code-set
/opt/app/resource/strings/fr_FR.iso8859-15  -- French strings in iso8859-15 code-set
/opt/app/resource/strings/jp_JP.utf8        -- Japanese strings in utf-8 code-set
```

At runtime, specify the string file search path in the DBPATH/FGLRESOURCEPATH environment
variable by adding the name of current locale as sub-directory. For example, to find the correct
string files in one of the locale-specific directories shown above, set the FGLRESOURCEPATH
variable as follows (UNIX™ shell):

```
$ echo $LC_ALL
jp_JP.utf8
$ FGLRESOURCEPATH="$FGLRESOURCEPATH:/opt/app/resource/strings/$LC_ALL"
$ export FGLRESOURCEPATH
$ echo $FGLRESOURCEPATH
/opt/app/forms:/opt/app/resource/strings/jp_JP.utf8
```

## Localized string files on mobile devices

On mobile devices, the language is determined by the operating system regional settings.

- On iOS devices, go to `Settings >> General >> Language & Region`
- On Android™ devices, go to `Settings >>
  General Management >> Language and Input`

The selected language is identified by a locale code following the ISO 639 standard.
Below are some language code examples; see the mobile OS documentation for information about
available languages and their corresponding ISO 639-x codes.

- `en` - English (for all regions)
- `en_US` - English in the United States
- `en_GB` - English in the United Kingdom

On startup, the mobile app will search for localized string files (.42s)
in the following directories:

1. appdir/full-locale-code:
   the sub-directory with language and category/region codes. For example for
   an English-US locale:
   appdir/en\_US.
2. appdir/language-code:
   the sub-directory with language code. For example for an English-US locale:
   appdir/en.
3. appdir/defaults: the fallback
   directory where default string files are located.

In order to localize your application, you simply need to place your
.42s localized string files in the appropriate language
sub-directory.

If the .42s filenames do not match the main program name, define the list of
localized strings files in the app's fglprofile file.

If you want to distinguish language categories (Simplified/Tradition Chinese), or if
you want to use different texts for territories that share the same language
(English in USA or Great Britain), create language sub-directories with the exact OS
locale identifier:

- For English in the USA, use "en\_US"
- For English in the United Kingdom, use "en\_GB"
- For English in Canada, use "en\_CA"
- etc...

```
appdir/en_US/mystrings.42s
appdir/en_GB/mystrings.42s
appdir/en_CA/mystrings.42s
```

If the language category or region can be ignored, create language sub-directories
with names matching the language identifier only:

- For English, use "en"
- For French, use "fr"
- For German, use "de"
- etc...

```
appdir/en/mystrings.42s
appdir/fr/mystrings.42s
appdir/de/mystrings.42s
```

A default set of string files can be provided under
appdir/defaults, in case the
regional settings of the device do not match one of the locale directories you
provide, otherwise the application will stop with error -8006:

```
appdir/defaults/mystrings.42s
```

For more details about the mobile app directory structures
(appdir), see [Directory structure for GMA apps](../17_mobile-applications/5104-directory-structure-for-gma-apps.md "Platform-specific rules need to be considered when deploying on Android devices (GMA).") and
[Directory structure for GMI apps](../17_mobile-applications/5108-directory-structure-for-gmi-apps.md "Platform-specific rules need to be considered when deploying on iOS devices (GMI).").

## Selecting the application language at runtime

The program can reset the path to find localized strings at runtime with the [`base.Application.reloadResources(path)`](../15_library-reference/2979-base-application-reloadresources.md "Resets FGLRESOURCEPATH and reloads localized string resources.") method. This can
be used to implement a login dialog, where the end user can choose the application language.

However, the `base.Application.reloadResources(path)` method
will only have an impact on subsequent loaded resources: The localized strings of already loaded
.42m modules and .42f forms are left unchanged. Subsequent
loaded forms and modules will get localized strings from new resource lookup path. Therefore, this
method must be called at the beginning of the program, before loading the first application
form.

A typical pattern is to do display an initial form that allows the user to select the language,
build the path to the localized strings files corresponding to the selected language, call the
`base.Application.reloadResources(path)` method, and reload the
initial form:

```
IMPORT os
MAIN
    DEFINE done BOOLEAN
    DEFINE rec RECORD
                user STRING,
                pswd STRING,
                lang CHAR(2)
           END RECORD,
           path STRING

    -- Login dialog with language selection
    LET rec.lang = "en" -- Can be "en", "fr", "ge" ...
    WHILE NOT done
        LET path = os.Path.join(base.Application.getProgramDir(), rec.lang)
        CALL base.Application.reloadResources(path)
        OPEN FORM f FROM "main"
        DISPLAY FORM f
        INPUT BY NAME rec.* WITHOUT DEFAULTS
            ON CHANGE lang
               EXIT INPUT -- restarts the input with the new locale settings
            ON ACTION cancel
               EXIT PROGRAM
            AFTER INPUT
               LET done = TRUE
        END INPUT
    END WHILE
    -- Here starts the real application code
    ...
END MAIN
```

## Related links

**Related concepts**  

[Deploying mobile apps on Android devices](../17_mobile-applications/5103-deploying-mobile-apps-on-android-devices.md "This section contains information to create a mobile application to be deployed on Android devices.")

[Deploying mobile apps on iOS devices](../17_mobile-applications/5107-deploying-mobile-apps-on-ios-devices.md "This section contains information to create a mobile application to be deployed on iOS devices.")

[Locale and character set basics](0867-locale-and-character-set-basics.md "This section is an introduction to locale and character set basics.")
