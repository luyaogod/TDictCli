---
title: "detach"
source: "fgl-topics/c_fgl_Debugger_detach.html"
breadcrumb: "Programming tools > Integrated debugger > Debugger commands > detach"
type: "concept"
---

# detach

> The detach command closes the TCP connection of a remove debug session.

## Syntax

```
detach
```

## Usage

The `detach` command must be used to terminate a remove debug session, by closing
the debug TCP connection.

## Example

```
(fgldb) detach
```

## Related links

**Related concepts**  

[Attaching to a running program](2577-attaching-to-a-running-program.md "It is possible to start the debugger for a program running on the same computer.")

[Debugging on a mobile device](2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.")
