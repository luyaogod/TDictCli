---
title: "Example 5: Writing to STDERR"
source: "fgl-topics/c_fgl_ClassChannel_example_5.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Examples > Example 5: Writing to STDERR"
type: "concept"
description: "The following code implements a simple user function to write messages to the stderr stream: FUNCTION to_stderr(s) DEFINE s STRING DEFINE c base.Channel LET c = base.Channel.create() CALL ..."
---

# Example 5: Writing to STDERR

The following code implements a simple user function to write messages to the stderr stream:

```
FUNCTION to_stderr(s)
   DEFINE s STRING
   DEFINE c base.Channel
   LET c = base.Channel.create()
   CALL c.openFile("<stderr>", "w")
   CALL c.writeLine(s)
   CALL c.close()
END FUNCTION
```
