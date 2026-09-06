---
title: "util.Channels.accept"
source: "fgl-topics/c_fgl_ext_util_Channels_accept.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.accept"
type: "concept"
---

# util.Channels.accept

> Returns the accepted base.Channel object for a given a server-socket.

## Syntax

```
util.Channels.accept(
     server base.Channel )
RETURNS base.Channel
```

1. server is a `base.Channel` object opened with
   `base.Channel.openServerSocket()`.

## Usage

The `util.Channels.accept()` takes a `base.Channel` server-socket
object that has been returned from a call to the [`util.Channels.select()`](3499-util-channels-select.md "Waits for activity on a set of TCP socket listening base.Channel objects.") method.

The resulting `base.Channel` object can be used to read data from.

For a complete description and code example, see [`util.Channels.select()`](3499-util-channels-select.md "Waits for activity on a set of TCP socket listening base.Channel objects.").

## Related links

**Related concepts**  

[util.Channels.select](3499-util-channels-select.md "Waits for activity on a set of TCP socket listening base.Channel objects.")

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
