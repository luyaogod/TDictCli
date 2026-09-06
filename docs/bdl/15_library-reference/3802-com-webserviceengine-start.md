---
title: "com.WebServiceEngine.Start"
source: "fgl-topics/c_gws_ComWebServiceEngine_Start.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebServiceEngine class > WebServiceEngine methods > com.WebServiceEngine.Start"
type: "concept"
---

# com.WebServiceEngine.Start

> Starts the Web Service engine.

## Syntax

```
com.WebServiceEngine.Start()
```

## Usage

The `com.WebServiceEngine.Start()` class method starts the engine for all
registered Web Services.

If you run the web services server program in standalone mode, port 80 is used unless
FGLAPPSERVER is set. We recommend that you set FGLAPPSERVER, as port 80 may already be in use by
other web servers. If you run the web services server program through the Genero Application Server,
the FGLAPPSERVER variable is automatically set by the Genero Application Server. Do NOT manually set
FGLAPPSERVER in this case.

In case of error, the method throws an exception and sets the
`status` variable. Depending on the error, a human-readable description of the
problem is available in the `sqlca.sqlerrm` register. See [Error handling in GWS calls (status)](../16_web-services/5064-error-handling-in-gws-calls-status.md "When errors are encountered, the methods of GWS classes can throw exceptions and set the status variable with the appropriate error number.").

## Related links

**Related concepts**  

[FGLAPPSERVER](../07_configuration/0520-fglappserver.md "Defines the listening TCP port of the Web service in development context.")
