---
title: "fgldb"
source: "fgl-topics/c_fgl_tools_fgldb.html"
breadcrumb: "Programming tools > Command reference > fgldb"
type: "concept"
---

# fgldb

> The fgldb tool is an interface program for remote debugging.

## Syntax 1: Debugging an application running on the computer

```
fgldb -p process-id
```

1. process-id is the process identifier of the
   `fglrun` process.

## Syntax 2: Debugging an app running on a mobile device

```
fgldb -m host[:port] ]
```

1. host is the hostname or IP address of the mobile device
   where the program executes.
2. port is the TCP debug port number to connect to, default
   is 6400.

## Options

| Option | Description |
| --- | --- |
| `-V` | Displays version information. |
| `-h` | Displays options for the tool. |
| `-p process-id` | Attach to an `fglrun` process running on the same computer, by using its process id. |
| `-m host[:port]` | Attach to a running mobile app to debug, using the mobile device hostname/IP and debug port. |

## Usage

The fgldb command line tool is an interface tool to attach
to the FGL integrated debugger remotely.

The fgldb tool can be used to:

- Debug an fglrun process currently running on the same computer, by using its
  process id. For more details, see [Attaching to a running program](2577-attaching-to-a-running-program.md "It is possible to start the debugger for a program running on the same computer.").
- Debug an application on a mobile device, by using the mobile device IP adress and its TCP debug
  port. For more details, see [Debugging on a mobile device](2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.")

## Related links

**Related concepts**  

[Integrated debugger](2573-integrated-debugger.md "Describes the command-line debugger you can use to find bugs in your programs.")
