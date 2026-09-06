---
title: "util.Channels.selectWithTimeout"
source: "fgl-topics/c_fgl_ext_util_Channels_selectWithTimeout.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods > util.Channels.selectWithTimeout"
type: "concept"
---

# util.Channels.selectWithTimeout

> Waits for activity on a set of TCP socket listening base.Channel objects and returns after a given period of inactivity.

## Syntax

```
util.Channels.selectWithTimeout(
     channels DYNAMIC ARRAY OF base.Channel,
     timeout INTEGER )
RETURNS INTEGER
```

1. channels is an array of `base.Channel` objects opened with
   [`base.Channel.openServerSocket()`](2992-base-channel-openserversocket.md "Open a TCP server socket channel.") or provided by [`util.Channels.accept()`](3497-util-channels-accept.md "Returns the accepted base.Channel object for a given a server-socket.").
2. timeout is a number of seconds.
3. Returns the index of the channel that can be used to read data from, or zero if the timeout has
   expired.

## Usage

The `util.Channels.selectWithTimeout()` is equivalent to [`util.Channels.select()`](3499-util-channels-select.md "Waits for activity on a set of TCP socket listening base.Channel objects."), except
that it takes a timeout as parameter.

The purpose of this method is to give the control back to the program, after a given period of
inactivity.

The timeout is specified as a number of seconds.

The method returns zero, when none of the TCP socket channels have got data to read from, and the
timeout period has expired.

## Related links

**Related concepts**  

[util.Channels.select](3499-util-channels-select.md "Waits for activity on a set of TCP socket listening base.Channel objects.")

[util.Channels.accept](3497-util-channels-accept.md "Returns the accepted base.Channel object for a given a server-socket.")

[The Channel class](2980-the-channel-class.md "The base.Channel class is a built-in class providing basic input/output functions.")
