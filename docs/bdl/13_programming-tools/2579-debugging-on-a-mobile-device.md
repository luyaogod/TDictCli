---
title: "Debugging on a mobile device"
source: "fgl-topics/c_fgl_Debugger_003.html"
breadcrumb: "Programming tools > Integrated debugger > Debugging on a mobile device"
type: "concept"
---

# Debugging on a mobile device

> It is possible to remotely start the debugger for an app running on a mobile device.

## Basics

When an app was created with debug mode and is running on a device, it is possible to attach to
the runtime system on the mobile for debug purpose, by using the
`fgldb -m host[:port]` command.

The app must have been created in debug mode: Check out how to build mobile apps with debug mode
in the [Deploying mobile apps](../17_mobile-applications/5102-deploying-mobile-apps.md "This section describes how to build and deploy mobile apps with Genero.") section.

On iOS devices, after installing the app, make sure that the debug port is enabled in the app
settings. Otherwise, the app will not listen to the debug port.

The fgldb command line tool takes two arguments: The host (or IP address) of
the mobile device, and an optional TCP port number to connect to. For mobile devices, the debug TCP
port is `6400`. Note that this is the same port the mobile front-end is listening to
for GUI connection, when working in GUI client/server mode.

> **Tip:**
>
> When attaching to the mobile app with fgldb, the app will
> typically be executing a dialog instruction waiting for user interaction. If you need to debug app
> code before the first dialog instruction, add a temporary `MENU` statement before the
> code to be inspected. When the code is fixed, remove the `MENU` statement.

Before starting a debug session, make sure that you fulfill the [prerequisites for debugging](2575-prerequisites-to-run-the-debugger.md "Some requirements are needed before running the integrated debugger.").

## Debugging an app running on a physical device

Considering the mobile device IP address is "`192.168.1.23`", and the
application is running locally on a physical mobile device, you can open a debug
session from the development machine as
follows:

```
$ fgldb -m 192.168.1.23:6400
108	    DISPLAY ARRAY contlist TO sr.*
(fgldb)
```

The `(fgldb)` prompt indicates that you are now connected to the
fglrun process on mobile device, and the program flow is suspended.
To continue with the program flow, enter the "`continue`" debugger
command:

```
(fgldb) continue
Continuing.
```

The application will then resume on the mobile device. To suspend the program again
and enter debugger commands, press CTRL-C in the debug console; fgldb
will show the interrupt message and give you the control back:

```
...
Continuing.
^CINTERRUPT
108	    DISPLAY ARRAY contlist TO sr.*
(fgldb)
```

At this point, you can for example set a break point and continue until the break
point is reached:

```
(fgldb) b 427
Breakpoint 2 at 0x00000000: file contacts.4gl, line 427.
(fgldb) continue 
Continuing.
Breakpoint 2, edit_contact() at contacts.4gl:427
427	    IF new THEN
(fgldb)
```

To finish the debug session, close the connection with the "`detach`"
debugger command:

```
(fgldb) detach
Connection closed by foreign host.
```

## Debugging an app running on an Android device emulator

When the mobile application is executing on an Android™ device emulator in the same machine as the development
environment, you must first forward the 6400 TCP port.

To forward the port 6400, use the following Android Debug Bridge tool command:

```
$ adb forward tcp:6400 tcp:6400
```

To be able to show GMA service debug information from a browser with the <http://localhost:6480>
URL, you need also to forward the port 6480:

```
$ adb forward tcp:6480 tcp:6480
```

## Related links

**Related concepts**  

[fgldb](2520-fgldb.md "The fgldb tool is an interface program for remote debugging.")
