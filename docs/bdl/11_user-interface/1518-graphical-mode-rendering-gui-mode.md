---
title: "Graphical mode rendering (GUI mode)"
source: "fgl-topics/c_fgl_DynamicUI_004.html"
breadcrumb: "User interface > User interface basics > Genero user interface modes > Graphical mode rendering (GUI mode)"
type: "concept"
description: "What is the graphical mode (GUI)? Genero supports the Graphical User Interface (GUI) mode to display application windows and forms with a real graphical look and feel, for desktop workstation, web ..."
---

# Graphical mode rendering (GUI mode)

## What is the graphical mode (GUI)?

Genero supports the Graphical User Interface (GUI) mode to display application [windows and forms](1561-windows-and-forms.md "The section describes the concept of windows and forms in the language.") with a real graphical look and
feel, for desktop workstation, web browsers and mobile front-end platforms.

The graphical mode is the default with Genero. The [FGLGUI](../07_configuration/0525-fglgui.md "Defines the user interface mode to be used by the program.") environment variable can be set to 0 (zero) to run the application in [text mode](1517-text-mode-rendering-tui-mode.md).

![Universal Rendering screenshot](../_images/Screen005_gbc.jpg)

*Graphical rendering window example*

## Form specification files for GUI mode

When designing a .per form file, use a [`LAYOUT`](1715-layout-section.md "The LAYOUT section defines the graphical alignment of the form by using a tree of layout containers.") section to define the
layout of the form:

```
SCHEMA stores
LAYOUT
...
END
ATTRIBUTES
...
END
```

Forms defined with the `SCREEN` layout section can be displayed in GUI mode.
However, to benefit from Genero BDL graphical enhancements, use a `LAYOUT`
section.

## Defining the target front-end

In graphical mode, the application forms are displayed on the front-end workstation identified
with the [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.") environment variable.

If this variable is not defined, the runtime system (fglrun) assumes that the
front-end executes on the same computer.

## Traditional GUI mode

To simplify migration from text mode to graphical mode with legacy applications, Genero supports
the [Traditional GUI mode](1519-graphical-mode-with-traditional-display.md) option to render all
application windows in a single front-end GUI window.

## Checking for graphical mode in programs

In the program code, use the [`ui.Interface.getFrontEndName()`](../15_library-reference/3103-ui-interface-getfrontendname.md "Returns the type of the front-end currently in use.") method to query for the front-end type.

When this method return a value different from `"console"`, the program executes
in graphical mode.

## Defining the GBC to be used in direct mode

In [direct mode](1521-connecting-with-a-front-end.md), it is possible to configure
which GBC version has to be transmitted by the runtime system to the front-end.

The GBC component files must be located on the computer where the fglrun
program executes, and will be transferred when the application starts.

The GBC component will be searched in the following directories:

1. The appdir/gbc directory, where
   appdir is the directory where the [program
   file](../09_advanced-features/0829-executing-programs.md "There are different ways to execute compiled programs, depending on the configuration and the development or production context.") is located,
2. The directory defined in the [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode.")
   environment variable,
3. The [$FGLDIR/web\_utilities/gbc/gbc](../07_configuration/0523-fgldir.md "Defines the installation directory of Genero Business Development Language.") directory.

> **Tip:**
>
> Set the [FGLGUIDEBUG](1527-debugging-the-front-end-protocol.md) environment
> variable, to enable GUI protocol debug logging and verify which GBC is transmitted to the front-end
> by the runtime system. Search for log lines
> like:
>
> ```
> *** uic_FT_processGet.1188: requestedName=gbc://index.html realName=/app/gbc/index.html
> ```

## Defining the GBC to be used with the GAS

When executing applications through the Genero Application Server, displaying on the GDC, GMA or
GMI front-ends, the GBC component files are found with the mechanism available in the GAS.

The $APPDIR/gbc directory is the recommended location for a production
environment, when a specific GBC is required by an application. This is typical when an application
requires a specific GBC customization.

Otherwise, if no specific GBC is required, the default GBC set in the GAS configuration will be
used (check the `GBC_LOOKUP_PATH` .xcf parameter)

The [FGLGBCDIR](../07_configuration/0524-fglgbcdir.md "Defines the GBC component to be used in GUI direct mode.") environment variable is not
taken into account by the GAS to define the GBC.

For more details about GBC usage with the GAS, see the Configuring GBC client for
applications topic in the Genero Application Server User Guide.

## Building mobile apps with a specific GBC

When building a GMA or GMI app, the GBC component needs to be bundled with the app package.

The gmabuildtool and gmibuildtool commands support an
option to specify the GBC to be bundled with the app.

For more details, see [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices."), [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.").

## Child topics

- [Graphical mode with Traditional Display](1519-graphical-mode-with-traditional-display.md)
