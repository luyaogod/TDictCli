---
title: "GUI log events"
source: "fgl-topics/c_fgl_feconn_logevent.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > GUI log events"
type: "concept"
description: "The GUI protocol supports a special front-end event named \" LogEvent \", to add a marker in the GUI log, with some debug data. When a log event is fired, it is written to the GUI logs as follows ..."
---

# GUI log events

The GUI protocol supports a special front-end event named "`LogEvent`", to add a
marker in the GUI log, with some debug data.

When a log event is fired, it is written to the GUI logs as follows (FGLGUIDEBUG
output):

```
... _om 123{}{{LogEvent 41 {{data " ... free-text ... "}}}}
```

The typical usage is with the Genero Ghost Client (GGC) tool, to implement non-regression tests:
When running a program with a GUI front-end, a `LogEvent` can be triggered with the
Alt-F12 key. The `LogEvent` events in the GUI log are then reused by GGC, for example
to generate snapshots of the AUI tree.

For more details, see the documentation of the Genero Ghost Client (GGC)

## Related links

**Related concepts**  

[Debugging the front-end protocol](1527-debugging-the-front-end-protocol.md "Debugging the front-end protocol")

[Front-end protocol logging](1528-front-end-protocol-logging.md "Front-end protocol logging")
