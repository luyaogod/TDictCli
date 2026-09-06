---
title: "util.Channels methods"
source: "fgl-topics/c_fgl_ext_util_Channels_methods.html"
breadcrumb: "Library reference > Extension packages > The util package > The util.Channels class > util.Channels methods"
type: "concept"
---

# util.Channels methods

> Methods for the util.ChannelsDate class.

| Name | Description |
| --- | --- |
| util.Channels.accept( server base.Channel ) RETURNS base.Channel | Returns the accepted `base.Channel` object for a given a server-socket. |
| util.Channels.copyN( in base.Channel, out base.Channel, n INTEGER ) RETURNS INTEGER | Copies a given number of bytes from one `base.Channel` to another. |
| util.Channels.select( channels DYNAMIC ARRAY OF base.Channel ) RETURNS INTEGER | Waits for activity on a set of TCP socket listening `base.Channel` objects. |
| util.Channels.selectWithTimeout( channels DYNAMIC ARRAY OF base.Channel, timeout INTEGER ) RETURNS INTEGER | Waits for activity on a set of TCP socket listening `base.Channel` objects and returns after a given period of inactivity. |
| util.Channels.readBinaryString( in base.Channel, n INTEGER ) RETURNS STRING | Reads a given number of bytes from a `base.Channel` as a character string. |
| util.Channels.readNetInt16( in base.Channel ) RETURNS SMALLINT | Reads next two bytes from a `base.Channel` as a 16 bit integer, in network byte order. |
| util.Channels.readNetInt32( in base.Channel ) RETURNS INTEGER | Reads next four bytes from a `base.Channel` as a 32 bit integer, in network byte order. |
| util.Channels.readNetInt8( in base.Channel ) RETURNS SMALLINT | Reads next byte from a `base.Channel` as a 8 bit integer. |
| util.Channels.writeBinaryString( out base.Channel, s STRING ) | Writes a character string to a `base.Channel`, without the trailing zero. |
| util.Channels.writeNetInt16( out base.Channel, v SMALLINT ) | Writes two bytes of a 16 bit integer to a `base.Channel`, in network byte order. |
| util.Channels.writeNetInt32( out base.Channel, v INTEGER ) | Writes four bytes of a 32 bit integer to a `base.Channel`, in network byte order. |
| util.Channels.writeNetInt8( out base.Channel, v SMALLINT ) | Writes the byte of an 8 bit integer to a `base.Channel`. |
