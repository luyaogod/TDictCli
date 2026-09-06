---
title: "FGLWSDEBUG"
source: "fgl-topics/c_fgl_EnvVariables_FGLWSDEBUG.html"
breadcrumb: "Configuration > Environment variables > Genero environment variables > FGLWSDEBUG"
type: "concept"
---

# FGLWSDEBUG

> The FGLWSDEBUG environment variable enables web services library debugging.

When FGLWSDEBUG environment variable is set to 1, debug information is written to
stderr for each operation done by the web services library.

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.

UNIX™ (shell) example:

```
$ FGLWSDEBUG=1
$ export FGLWSDEBUG
$ fglrun myprog 2>wsdbg.txt
```

The FGLWSDEBUG debug log is to be used in development context only. The output format can change
in future product releases.

| Value | Definition |
| --- | --- |
| 0 | No data displayed; debug turned off. |
| 1 | Display socket errors. |
| 2 | Display HTTP bodies of incoming and outgoing requests (the XML content) |
| 3 | Display all information about incoming and outgoing requests (HTTP headers + HTTP bodies) |

## Related links

**Related concepts**  

[Debugging](../16_web-services/4501-debugging.md "Turn on the debug mode to log the data sent or received by your Web service application.")
