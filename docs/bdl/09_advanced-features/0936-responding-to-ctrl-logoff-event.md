---
title: "Responding to CTRL_LOGOFF_EVENT"
source: "fgl-topics/c_fgl_programs_016.html"
breadcrumb: "Advanced features > Configuration options > Runtime configuration in FGLPROFILE > Responding to CTRL_LOGOFF_EVENT"
type: "concept"
description: "FGLPROFILE fglrun.ignoreLogoffEvent controls program behavior in case of logoff events on Windows platforms."
---

# Responding to CTRL_LOGOFF_EVENT

> FGLPROFILE fglrun.ignoreLogoffEvent controls program behavior in case of logoff events on Windows™ platforms.

## Syntax

```
fglrun.ignoreLogoffEvent = true
```

## Usage

On Windows platforms,
when the user disconnects, the system sends a CTRL\_LOGOFF\_EVENT
event to all console applications. When the runtime system receives
this event, it stops immediately.

On a Windows Terminal Server, if an Administrator
user closes his session, a CTRL\_LOGOFF\_EVENT is sent to all console
applications started by ANY user connected to the machine (even
if these applications were not started by the administrator).

To
prevent the runtime system from stopping on a logoff event, you can
use the `fglrun.ignoreLogoffEvent` entry in the
FGLPROFILE configuration file. If this entry is set to `true`,
the CTRL\_LOGOFF\_EVENT event is ignored by the runtime system.

As a result, when the administrator user disconnects on a Windows Terminal Server, programs
started by remote users would not stop.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
