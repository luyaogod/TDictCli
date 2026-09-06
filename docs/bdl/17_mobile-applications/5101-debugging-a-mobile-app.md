---
title: "Debugging a mobile app"
source: "fgl-topics/c_fgl_mobile_debugging.html"
breadcrumb: "Mobile applications > Debugging a mobile app"
type: "concept"
---

# Debugging a mobile app

> Different solutions are available to debug a mobile app.

## Debugging a mobile app in development mode

When executing a mobile app program on a server, displaying the user interface on a
mobile front-end defined by [FGLSERVER](../11_user-interface/1521-connecting-with-a-front-end.md),
it is possible to debug the BDL code with the fglrun -d option:

```
$ export FGLSERVER=device-ip-address
$ fglrun -d main.42m
```

For more details, see [Starting fglrun in debug mode](../13_programming-tools/2576-starting-fglrun-in-debug-mode.md "The runtime system can be started in debug mode with the -d option.").

## AUI protocol debugging

With app running on a server or on the device, it is possible to show AUI protocol
exchanges in the console running the program on the server, by setting the FGLGUIDEBUG
environment variable to 1. When this variable is set, you can watch user interface
events that occur during program execution and how they are treated by the runtime system.

To set the FGLGUIDEBUG environment variable for an app running on the device, use an
FGLPROFILE `fglrun.environment` entry. The output can be inspected
with the program logs as described later in this section.

For more details, see [FGLGUIDEBUG](../07_configuration/0526-fglguidebug.md "Defines the debug level in GUI mode.").

## AUI protocol logging in development mode

With an app running on a server, it can be useful to log AUI protocol exchanges between
the runtime system and the mobile front-end, to inspect the content, or replay a scenario.
This is possible with the `--start-guilog` and `--run-guilog`
options of fglrun:

```
$ fglrun --start-guilog=case1.log
```

The AUI protocol log file produced by the `--start-guilog` option can
then be shared for analysis.

For more details, see [Front-end protocol logging](../11_user-interface/1528-front-end-protocol-logging.md).

## Debugging a mobile app running on the device

When executing the mobile app on a device, and if the app has been created with debug
mode, it is possible to establish a connection to the runtime system executing on
the mobile device, by using the fgldb command line tool.

> **Important:**
>
> On iOS devices, after installing the app, you need to enable the
> debug port in the app settings, otherwise the app will not listen to the debug port.

For example:

```
$ fgldb -m 192.168.1.23:6400
108	    DISPLAY ARRAY contlist TO sr.*
(fgldb)
```

This way you can debug an app running on a device, by using the source code located
on the server where the fgldb command is executed.

For more details, see [Debugging on a mobile device](../13_programming-tools/2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.").

## Building mobile apps in debug mode

In order to enable debug features of an app running on a mobile device, you need to build the app
in debug mode:

- For Android:

  The gmabuildtool provides the `--mode debug`
  option, to create a debug version of the APK.

  For more details, see [Building Android apps with Genero](5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.").
- For iOS:

  The gmibuildtool provides the `--mode debug`
  option, to create a debug version of the IPA. The certificate defined in the provisioning profile
  must be a development certificate.

  After installing the debug version of the app on your iOS
  device, you need to enable the debug port in the app settings.

  For more details, see [Building iOS apps with Genero](5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.").

## Browse the AUI tree created on the mobile front-end side

The content of the [Abstract User Interface
tree](../11_user-interface/1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.") created on the mobile front-end side can be inspected from a web
browser, when the app has been created with debug mode, or in development mode by
executing the app on a server and displaying on the device.

To inspect AUI tree, open a web browser and enter the following
URL:

```
http://device-ip-address:6480 (or 6400)
```

For more details, see [Inspecting the AUI tree of a front-end](../11_user-interface/1515-inspecting-the-aui-tree-of-a-front-end.md).

## Viewing embedded app program logs

The program logs of an app running on a device can be viewed in a browser, if the app
was created in debug mode. VM messages (runtime errors, standard output and standard error)
are available. This feature is not available if the app is built in release mode.

To inspect program logs, open a web browser and enter the following
URL:

```
http://device-ip-address:6480 (or 6400)
```

A menu will then appear in the web page, where you can choose the VM output to be
inspected.

## Debugging a WEBCOMPONENT HTML/JavaScript

The HTML content and JavaScript code can be debugged on GMA and GMI mobile devices.

For more details, see [Debugging a web component](../11_user-interface/2384-debugging-a-web-component.md).

## Related links

**Related concepts**  

[Integrated debugger](../13_programming-tools/2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs.")

[Program profiler](../13_programming-tools/2622-program-profiler.md "Find out what function is causing the bottleneck in your program.")
