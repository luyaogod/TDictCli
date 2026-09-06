---
title: "Example 3: Reading lines from a text file"
source: "fgl-topics/c_fgl_ClassChannel_example_3.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Examples > Example 3: Reading lines from a text file"
type: "concept"
description: "MAIN DEFINE i INTEGER DEFINE s STRING DEFINE ch base.Channel LET ch = base.Channel.create() CALL ch.openFile(\"file.txt\",\"r\") LET i = 1 WHILE TRUE LET s = ch.readLine() IF ch.isEof() THEN EXIT WHILE ..."
---

# Example 3: Reading lines from a text file

```
MAIN
  DEFINE i INTEGER
  DEFINE s STRING
  DEFINE ch base.Channel 
  LET ch = base.Channel.create()
  CALL ch.openFile("file.txt","r")
  LET i = 1
  WHILE TRUE
    LET s = ch.readLine()
    IF ch.isEof() THEN EXIT WHILE END IF
    DISPLAY i, " ", s 
    LET i = i + 1
  END WHILE
  CALL ch.close()
END MAIN
```
