---
title: "Example 6: Closing the output stream"
source: "fgl-topics/c_fgl_ClassChannel_example_6.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Examples > Example 6: Closing the output stream"
type: "concept"
description: "The following code shows how to send data lines to the wc command and terminate the processing by closing the channel output stream: MAIN DEFINE ch base.Channel LET ch = base.Channel.create() CALL ..."
---

# Example 6: Closing the output stream

The following code shows how to send data lines to the wc command and terminate
the processing by closing the channel output stream:

```
MAIN
    DEFINE ch base.Channel
    LET ch = base.Channel.create()
    CALL ch.openPipe("wc -l", "u") -- Unix command to count lines
    CALL ch.writeLine("one")
    CALL ch.writeLine("two")
    CALL ch.writeLine("three")
    CALL ch.closeOut() -- closes the writer output stream
    DISPLAY "result:", ch.readLine() -- reader input stream is still alive
    CALL ch.close()
END MAIN
```
