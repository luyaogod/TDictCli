---
title: "Automatic front-end startup"
source: "fgl-topics/c_fgl_DynamicUI_027.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Automatic front-end startup"
type: "concept"
description: "Important: This feature is desupported. It may still be available, but can be removed in a next version. The Genero can start a graphical front-end automatically, when the runtime system and the ..."
---

# Automatic front-end startup

> **Important:**
>
> This feature is desupported. It may still be
> available, but can be removed in a next version.

The Genero can start a graphical front-end automatically, when the runtime system and the
front-end reside on the same computer.

The GUI server autostart feature is supported for backward compatibility, and should be avoided.
Consider using other solutions to start your Genero application, such as GDC shortcuts for desktop
applications, the Genero Application Server (GAS) for web applications, or the ["Run On Server"](../17_mobile-applications/5111-running-mobile-apps-on-an-application-server.md "From the mobile device, programs can be started remotely on an application server, and displayed on the device.") solution for mobile applications.

When a program starts in graphical mode, the runtime system tries to open a connection to
the graphical front-end set by the [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.") environment
variable. This requires having the front-end already started and listening to the TCP port specified
by FGLSERVER.

In some configurations, such as *X11 workstations* or *METAFRAME/Citrix
Winframe* or *Microsoft™
Windows® Terminal Server*, each user may want to start
his own front-end to have a dedicated process. This can be done by starting the front-end
automatically when the program executes, based on settings in the DISPLAY (X11) or
SESSIONNAME/CLIENTNAME (WTSE) environment variables.

Automatic front-end startup settings are defined with
`gui.server.autostart.*` entries in [FGLPROFILE](../07_configuration/0483-the-fglprofile-file-s.md "FGLPROFILE environment variable defines Genero BDL configuration files"). In these FGLPROFILE entries, the term "GUI server" refers to the graphical front
end.

In a first time, the runtime system tries to establish the connection without starting
the front-end (in a normal usage, it is already started). The front-end is identified by the
FGLSERVER environment variable. If FGLSERVER is not defined, it defaults to
`localhost:0`, except if `gui.server.autostart.wsmap` entries are
defined in FGLPROFILE. When `wsmap` entries are defined, workstation id to GUI server
id mapping takes place and FGLSERVER defaults to `localhost:n`, where n is the GUI
server number found from `wsmap` entries.

If this first connection fails and the `gui.server.autostart.cmd` entry is
defined, the runtime system executes the command to start the GUI server, then waits for n seconds
as defined by `gui.server.autostart.wait` entry, and after this delay tries to
connect to the front-end. If the connection fails, it tries again for a number of attempts defined
by the `gui.server.autostart.repeat` entry. Finally, if the last attempt fails, the
runtime system stops with a GUI connection error [-6300](../15_library-reference/4483-genero-bdl-errors.md).

If the `gui.server.autostart.cmd` entry is not defined, neither
workstation id to GUI id mapping, nor automatic front-end startup is done.

Here is a detailed description of each `gui.server.autostart` FGLPROFILE
entry:

The `cmd` entry is used to define the command to be executed to start the
front-end:

```
gui.server.autostart.cmd = "/opt/app/gdc-2.30/bin/gdc -p %d -q -M"
```

Here, `%d` will be replaced by the TCP port the front-end must listen
to.

By default the runtime system waits for two seconds before it tries to connect to the
front-end. You can change this delay with the **wait** entry:

```
gui.server.autostart.wait = 5   -- wait five seconds
```

The runtime system tries to connect to the front-end ten times. You can change this with
the `repeat` entry:

```
gui.server.autostart.repeat = 3 -- repeat three times
```

The following FGLPROFILE entries can be used to define workstation id to front-end id
mapping:

```
gui.server.autostart.wsmap.max = 3
gui.server.autostart.wsmap.0.names = "fox:1.0,fox.sxb.4js.com:1.0"
gui.server.autostart.wsmap.1.names = "wolf:1.0,wolf.sxb.4js.com:1.0"
gui.server.autostart.wsmap.2.names = "wolf:2.0,wolf.sxb.4js.com:2.0"
```

The first `wsmap.max` entry defines the maximum number of front-end
identifiers to look for. The `wsmap.N.names` entries define a mapping for
each GUI server, where **N** is the front-end identifier. The value of those entries
defines a comma-separated list of workstation names to match. If no wsmap entries are
defined, the GUI server number will default to zero.

For `gui.server.autostart.wsmap` entries, the first GUI server number
starts at zero.

On X11 configurations, a workstation is identified by the DISPLAY environment variable.
In this example, `fox:1.0` identifies a workstation that will make the
runtime start a front-end with the number 1.

On Windows Terminal Server, the CLIENTNAME
environment variable identifies the workstation. If no corresponding front-end id can be found in
the `wsmap` entries, the front-end number defaults to the id of the session defined
by the SESSIONNAME environment variable, plus one. The value of this variable has the form
`protocol#id`; for example, `RDP-Tcp#4` would
automatically define a front-end id of 5 (4+1).

If the front-end processes are started on the same machine as the runtime system, you do
not need to set the FGLSERVER environment variable. This will then default to
`localhost:id`, where *id* will be detected from the
`wsmap` workstation mapping entries.

If the front-end is executed on a middle-tier machine that is different from the
application server, MIDHOST for example, you can set FGLSERVER to MIDHOST without a GUI server id.
The workstation mapping will automatically find the id from the `wsmap` settings.

> **Note:**
>
> The Genero Desktop Client (GDC), raises the control panel to the top of the window stack when
> you try to restart it. In this case, the program window might be hidden by the GDC control panel. To
> avoid this problem, use the -M option to start the GDC in minimized mode.
