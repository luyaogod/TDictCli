---
title: "Handle channel exceptions"
source: "fgl-topics/c_fgl_ClassChannel_handle_exception.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Usage > Handle channel exceptions"
type: "concept"
description: "Channel errors can be trapped with the WHENEVER ERROR exception handler: MAIN DEFINE ch base.Channel DEFINE num INT, label CHAR(50) LET ch = base.Channel.create() --CALL ch.openFile(\"data.txt\",\"w\") ..."
---

# Handle channel exceptions

Channel errors can be trapped with the [`WHENEVER ERROR`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") exception handler:

```
MAIN
    DEFINE ch base.Channel
    DEFINE num INT, label CHAR(50)
    LET ch = base.Channel.create()
    --CALL ch.openFile("data.txt","w")
    WHENEVER ERROR CONTINUE
    CALL ch.write([num,label])
    IF status<0 THEN
        DISPLAY SFMT("Could not write to channel: status=%1",status)
    END IF
END MAIN
```

Or with a [`TRY/CATCH`](../09_advanced-features/0850-whenever-directive.md "Use the WHENEVER directive to define how exceptions must be handled for the rest of the module.") block:

```
MAIN
    DEFINE ch base.Channel
    DEFINE num INT, label CHAR(50)
    LET ch = base.Channel.create()
    --CALL ch.openFile("data.txt","w")
    TRY
        CALL ch.write([num,label])
    CATCH
        DISPLAY SFMT("Could not write to channel: status=%1",status)
    END TRY
END MAIN
```

Both programs will produce the following
output:

```
Could not write to channel: status=-6344
```
