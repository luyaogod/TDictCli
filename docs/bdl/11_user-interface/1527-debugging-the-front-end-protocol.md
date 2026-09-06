---
title: "Debugging the front-end protocol"
source: "fgl-topics/c_fgl_feconn_fglguidebug.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Debugging the front-end protocol"
type: "concept"
description: "When setting the FGLGUIDEBUG environment variable to 1, information about GUI communication will be written to stderr by the runtime system, and the GUI protocol exchange will be indented for a better ..."
---

# Debugging the front-end protocol

When setting the [FGLGUIDEBUG](../07_configuration/0526-fglguidebug.md "Defines the debug level in GUI mode.") environment
variable to 1, information about GUI communication will be written to stderr by
the runtime system, and the GUI protocol exchange will be indented for a better readability in the
front-end log window.

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.

UNIX™ (shell) example:

```
$ FGLGUIDEBUG=1
$ export FGLGUIDEBUG
$ fglrun myprog 2>guidbg.txt
```

The output format of FGLGUIDEBUG is for debug purpose only and can change in next product
releases.

In TUI mode, displayed screens can be dumped by setting the [DBSCREENDUMP or DBSCREENOUT](1536-text-mode-screen-dump.md) environment variables. This can be used to take a snapshot of
the current TUI screen, for debugging or testing purpose.

## Related links

**Related concepts**  

[GUI log events](1529-gui-log-events.md "GUI log events")
