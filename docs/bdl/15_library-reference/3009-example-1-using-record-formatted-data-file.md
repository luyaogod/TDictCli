---
title: "Example 1: Using record-formatted data file"
source: "fgl-topics/c_fgl_ClassChannel_example_1.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Channel class > Examples > Example 1: Using record-formatted data file"
type: "concept"
description: "This program reads data from file.txt , which contains two columns separated by a pipe ( | ) character. It writes this data to the end of fileout.txt , using a percent sign ( % ) as the delimiter. ..."
---

# Example 1: Using record-formatted data file

This program reads data from file.txt, which contains two columns separated
by a pipe (`|`) character. It writes this data to the end of
fileout.txt, using a percent sign (`%`) as the delimiter.

```
MAIN
  DEFINE custinfo RECORD
            cust_num INTEGER,
            cust_name VARCHAR(40)
         END RECORD
  DEFINE ch_in, ch_out base.Channel 
  LET ch_in = base.Channel.create()
  CALL ch_in.setDelimiter("|")
  LET ch_out = base.Channel.create()
  CALL ch_out.setDelimiter("%")
  CALL ch_in.openFile("file.txt","r")
  CALL ch_out.openFile("fileout.txt","w")
  WHILE ch_in.read(custinfo)
    CALL ch_out.write(custinfo)
  END WHILE
  CALL ch_in.close()
  CALL ch_out.close()
END MAIN
```
