---
title: "mobile.runOnServer"
source: "fgl-topics/c_fgl_frontcall_mobile_runonserver.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.runOnServer"
type: "concept"
---

# mobile.runOnServer

> Run an application from the Genero Application Server using the specified URL.

## Syntax

```
ui.Interface.frontCall("mobile", "runOnServer",
   [ appurl, timeout ], [] )
```

1. appurl - The GAS URL to the Genero application (this must be a
   `ua/r` URL).
2. timeout - The timeout (in seconds) to wait for the first connection to the
   web server and GAS answer. This does not include the time to start the server-side application.
   A timeout of zero indicates an
   infinite timeout.

## Usage

The `runOnServer` front call allows you to start an application in the Genero
Application Server (GAS) from an embedded/local application running on the mobile device. The remote
application's graphical user interface displays on the mobile device.

The front call returns when the called application ends and control goes back to the initial
application executing on the mobile device.

The applications executed on the GAS server must use the UTF-8 encoding. Mobile front-ends will
reject any attempt to display forms of an application that uses an encoding other than UTF-8.

The remote application must use `RUN` to start child programs. `RUN WITHOUT
WAITING` cannot be used by the remote application; it is not supported.

The appurl parameter identifies the remote application to be started. It must
contain an "`ua/r`" URL syntax (the UA protocol introduced with the GAS 3.00).

For example: `http://myappserver:6394/ua/r/myapp`.

This URL may contain a query string with parameters for the application, to be executed by the
GAS.

The timeout parameter can be used to give the control back to the local app
when the web server or GAS takes too long to respond. This timeout applies to the first connection
to the web server or GAS and does not include the time to start the server program.
A timeout of zero indicates an
infinite timeout.

In case of failure (such as application not found or timeout expired), the front call raises the
runtime error [-6333](4483-genero-bdl-errors.md) and the
HTTP status code of the request can be found in the error message details.

> **Note:**
>
> The application running on the GAS can only access the data-directory
> directory, in the sandbox of the embedded application that executes the `runOnServer`
> front call. File handling APIs like `fgl_getfile()` and
> `fgl_putfile()` can only access this directory on the mobile device. If no absolute
> path is specified in the file path for the mobile device, the
> data-directory is used.

## Example

```
TRY
    CALL ui.interface.frontcall("mobile","runOnServer",["http://santana:6394/ua/r/orders"],[])
CATCH
    ERROR err_get(status)
END TRY
```

## Related links

**Related concepts**  

[Running mobile apps on an application server](../17_mobile-applications/5111-running-mobile-apps-on-an-application-server.md "From the mobile device, programs can be started remotely on an application server, and displayed on the device.")
