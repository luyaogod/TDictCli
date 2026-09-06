---
title: "GUI connection timeout"
source: "fgl-topics/c_fgl_feconn_timeout.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > GUI connection timeout"
type: "concept"
description: "When initiating the connection to the front-end, if the front-end software is stopped, the host machine is down, or a firewall drops connections for the TCP port used for the GUI connection, the ..."
---

# GUI connection timeout

When initiating the connection to the front-end, if the front-end software is stopped, the host
machine is down, or a firewall drops connections for the TCP port used for the GUI connection, the
program will stop with an error after a given timeout.

This timeout can be specified with the following FGLPROFILE
entry:

```
gui.connection.timeout = seconds
```

The default timeout is 30 seconds.

## Related links

**Related concepts**  

[The FGLPROFILE file(s)](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files")
