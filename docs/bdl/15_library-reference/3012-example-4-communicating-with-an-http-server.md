---
title: "Example 4: Communicating with an HTTP server"
source: "fgl-topics/c_fgl_ClassChannel_example_4.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Examples > Example 4: Communicating with an HTTP server"
type: "concept"
description: "MAIN DEFINE ch base.Channel, eof INTEGER LET ch = base.Channel.create() -- HTTP protocol forces every line to be terminate by \\r\\n -- So we use channel binary mode to avoid CR+LF translation on ..."
---

# Example 4: Communicating with an HTTP server

```
MAIN
  DEFINE ch base.Channel, eof INTEGER
  LET ch = base.Channel.create()
  -- HTTP protocol forces every line to be terminate by \r\n
  -- So we use channel binary mode to avoid CR+LF translation on Windows.
  -- In text mode, each line would be terminated by \r\r\n on Windows.
  WHENEVER ERROR CONTINUE
  CALL ch.openClientSocket("localhost", 80, "ub", 30)
  IF status != 0 THEN
    DISPLAY "Could not open socket: error ", status
    EXIT PROGRAM 1
  END IF
  WHENEVER ERROR STOP
  -- HTTP expects CR+LF: Note that LF is added by writeLine()!
  CALL ch.writeLine("GET / HTTP/1.0\r")
  -- No HTTP headers...
  -- Empty line = end of headers 
  CALL ch.writeLine("\r")
  WHILE NOT eof 
    DISPLAY ch.readLine()
    LET eof = ch.isEof()
  END WHILE
  CALL ch.close()
END MAIN
```
