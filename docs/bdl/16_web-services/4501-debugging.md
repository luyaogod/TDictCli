---
title: "Debugging"
source: "fgl-topics/c_gws_debugging.html"
breadcrumb: "Web services > General > Debugging"
type: "concept"
---

# Debugging

> Turn on the debug mode to log the data sent or received by your Web service application.

Debug information is written to the standard error stream of the console. If needed, it can be
redirected to a file.

To turn on the debugging feature, set the FGLWSDEBUG environment variable before starting the
application.

The level of debugging depends on the value set for the FGLWSDEBUG variable. See [FGLWSDEBUG](../07_configuration/0537-fglwsdebug.md "The FGLWSDEBUG environment variable enables web services library debugging.").

To debug a Web Service application managed by the Genero Application Server (GAS), you have to
modify the value of the FGLWSDEBUG environment variable in the GAS configuration file. To understand
how environment variables are set within the GAS configuration file, see the
ENVIRONMENT\_VARIABLE topic in the Genero Application Server User Guide.
