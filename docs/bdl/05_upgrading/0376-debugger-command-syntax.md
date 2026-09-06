---
title: "Debugger command syntax"
source: "fgl-topics/c_fgl_MigI4GL_024.html"
breadcrumb: "Upgrading > Migrating from IBM® Informix® 4GL to Genero BDL > 4GL programming topics > Debugger command syntax"
type: "concept"
---

# Debugger command syntax

> While I4GL and Genero BDL both provide a program debugger, the commands used and how it is used can differ.

IBM® Informix® 4GL
(I4GL) provides a program debugger. Genero Business Development Language provides a program debugger
with a different set of commands as I4GL, compatible with the well-known gdb tool. This debugger can
be used alone in command line mode, or with a graphical shell compatible with gdb, such as
ddd:

```
ddd --debugger "fglrun -d myprog"
```

## Related links

**Related concepts**  

[Starting fglrun in debug mode](../13_programming-tools/2576-starting-fglrun-in-debug-mode.md "The runtime system can be started in debug mode with the -d option.")
