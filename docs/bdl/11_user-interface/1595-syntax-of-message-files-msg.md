---
title: "Syntax of message files (.msg)"
source: "fgl-topics/c_fgl_message_files_003.html"
breadcrumb: "User interface > Form definitions > Message files > Syntax of message files (.msg)"
type: "concept"
---

# Syntax of message files (.msg)

> A message file contains a set of messages identified by an integer number.

```
filename.msg
```

1. *filename* is the name of the message source file.

## Syntax of a message file

```
{
  message-definition
| include-directive 
}[...]
```

where message-definition
is:

```
.message-number
message-line | new-page
[...]
```

where include-directive
is:

```
.include filename
```

And where new-page is:

```
^L (Control-L, ASCII 12)
```

1. message-number is an integer in the range -2147483648 to 2147483647.
2. You can split the message into pages by adding the ^L (Control-L / ASCII 12) in a line.
3. Note that multi-line messages will include the newline (ASCII 10) characters.
