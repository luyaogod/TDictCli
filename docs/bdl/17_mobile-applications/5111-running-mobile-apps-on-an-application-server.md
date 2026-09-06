---
title: "Running mobile apps on an application server"
source: "fgl-topics/c_fgl_programs_runonserver.html"
breadcrumb: "Mobile applications > Deploying mobile apps > Running mobile apps on an application server"
type: "concept"
---

# Running mobile apps on an application server

> From the mobile device, programs can be started remotely on an application server, and displayed on the device.

## Purpose of remote application execution for mobile devices

Remote applications displayed on a mobile device allow the use of the processor, memory, storage
and software resources available on a server, for mobile users.

Executing remote/server applications for display on a mobile device requires a reliable and
constant network connection. If the network connection fails, the application will stop, as with
other client/server Genero front-ends.

Server applications can only be started through the Genero Application Server (GAS), by using the
UA protocol available since version 3.00. You must set up and configure the GAS for the programs you
want to start remotely. See the GAS documentation for more details.

Applications executed on the GAS server must use the UTF-8 encoding. Mobile front-ends will
reject any attempt to display forms of an application using an encoding other than UTF-8.

## Implementing the embedded mobile app

Create a small application to be deployed on the mobile device, which then starts the
application(s) on an GAS server.

The server application is started from the embedded application through the [`runOnServer`](../15_library-reference/3458-mobile-runonserver.md "Run an application from the Genero Application Server using the specified URL.") front call. The
embedded mobile application can be a very simple `MAIN / END MAIN` program, only
performing the "`runOnServer`" front call.

For example, this is the very minimum embedded application, starting a program on the
GAS:

```
MAIN
    CALL ui.interface.frontcall("mobile","runOnServer",
                                ["http://myappserver:6394/ua/r/myapp"],[])
END MAIN
```

When the remote application starts, the graphical user interface displays on the mobile
device.

When the the called server-side application ends, the `runOnServer` front call
returns, and control goes back to the initial application executing on the mobile device.

In development context, it is possible to execute the parent starter app on a server, display on
a mobile device with [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.") set properly, and use the
`runOnServer` front call. Because starting remote GAS applications is done with a
front call, this configuration mimics an embedded starter app running on the device.

## Using the runOnServer front call

The application executed on the server-side is identified by the first parameter of the
`runOnServer` front call. This application must be delivered by the Genero
Application Server. The parameter must contain an "`ua/r`" URL syntax (the UA
protocol introduced with the GAS 3.00).

For example: `http://myappserver:6394/ua/r/myapp`

The URL may contain a query string, with parameters for the application to be executed by the
GAS.

If needed, you can add a second argument to define a timeout as a number of seconds. The embedded
application will wait for the remote application to start, until the timeout expired.
A timeout of zero indicates an
infinite timeout.

In case of failure (application not found, timeout expired), the front call raises the runtime
error [-6333](../15_library-reference/4483-genero-bdl-errors.md) and the HTTP
status code of the request can be found in the error message details. Use a
`TRY/CATCH` block, to check if the execution of the server application was
successful:

```
MAIN
  TRY
    CALL ui.interface.frontcall("mobile","runOnServer",
         ["http://myappserver:6394/ua/r/myapp"],[])
  CATCH
    ERROR err_get(status)
  END TRY
END MAIN
```

Subsequent server-side application runs are allowed; the last active application will display on
the device. However, it is not possible to navigate between started applications. Therefore, an
application started with the `runOnServer` front call must only use the [`RUN`](../09_advanced-features/0830-run.md "The RUN instruction executes the command passed as argument.") instruction to start sub-programs.
`RUN WITHOUT WAITING` is not supported.

## Passing parameters to the server application

If needed, the embedded app can pass arguments to the server application by using parameter
specification in the URL string, with the `?Arg=value1&Arg=value1&...`
notation:

```
DEFINE params, base, complete_url STRING
LET params = "Arg=verbose&Arg=5677"
LET url = "http://myappserver:6394/ua/r/myapp"
LET complete_url = base || "?" || params
```

The remote program can retrieve the parameters with the [`arg_val()`](../15_library-reference/2726-arg-val.md "Returns a command line argument by position.") built-in function.

It is not needed to URL-encode the string passed to the `runOnServer` front
call.

See the GAS documentation (`AllowUrlParameters` attribute) about passing
parameters in the application URL.

This is an example of an embedded application to be deployed on the mobile device, which passes
parameters to a server-side application:

```
IMPORT util

MAIN
    DEFINE arr DYNAMIC ARRAY OF STRING, x INT
    MENU "test"
        COMMAND "runOnServer"
           CALL arr.clear()
           LET arr[1] = "first argument"
           LET arr[2] = "second argument"
           LET x = do_run("http://10.0.40.29:6394/ua/r/test1", 10, arr)
        COMMAND "exit"
           EXIT MENU
    END MENU
END MAIN

FUNCTION do_run(url,timeout,params)
  DEFINE url STRING,
         timeout SMALLINT,
         params DYNAMIC ARRAY OF STRING
  DEFINE i, r INTEGER, tmp STRING
  LET r = 0
  LET tmp = url
  FOR i=1 TO params.getLength()
      LET tmp = tmp || IIF(i==1,"?","&") || "Arg=" || params[i]
  END FOR
  TRY
    CALL ui.interface.frontcall("mobile","runOnServer",[tmp,timeout],[])
  CATCH
    ERROR err_get(status)
    LET r = -1
  END TRY
  RETURN r
END FUNCTION
```

A sample server-side application:

```
MAIN
    MENU "Prog1"
        COMMAND "arg1" MESSAGE "Arg 1 = ", arg_val(1)
        COMMAND "arg2" MESSAGE "Arg 2 = ", arg_val(2)
        COMMAND "arg3" MESSAGE "Arg 3 = ", arg_val(3)
        COMMAND "Quit" EXIT MENU
    END MENU
END MAIN
```

## Sharing files between embedded and server app

If files need to be shared between the embedded application and the server application, the
application running on the GAS can only access the data-directory directory,
in the sandbox of the embedded application that executes the "`runOnServer`"
front call.

This matters when using file handling APIs such as [`fgl_putfile()`](../15_library-reference/2774-fgl-putfile.md "Transfers a file from the virtual machine context to the front-end context.") and [`fgl_getfile()`](../15_library-reference/2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.") or front
calls like [`takePhoto`](../15_library-reference/3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.")
and [`launchURL`](../15_library-reference/3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.").

The data-directory on the mobile device can be found with the [`feInfo/dataDirectory`](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.") front
call. In both the embedded app and the app running on the server, this front call will return
the same directory.

The following workflow can be used:

1. Before starting the server application with a `runOnServer` front call, the
   embedded app must copy files to the data-directory.
2. While executing, the server application can retrieve files from the
   data-directory with `fgl_getfile()`, and send its own files to the
   data-directory, with `fgl_putfile()`.
3. When the server application terminates, the embedded app can read files the server application
   left in the data-directory.

If several remote applications are started successively on the server with a `RUN`
instruction, make sure to not overwrite files written by other server programs.

In order to write code for the embedded app, that can be executed in development mode (running
on a server) and on the mobile device, you can adapt to the execution context: Make a simple file
copy when executing on the mobile device, or do an `fgl_putfile()` call, when
running on the development server. Check the execution context with the
`base.Application.isMobile()` method.

This example, in the embedded app on the mobile device, copies a file from the device private
directory to the data-directory:

```
IMPORT os
...
    CALL mobile_copy_to_data_dir("myfile.txt")
...
FUNCTION mobile_copy_to_data_dir(fn)
    DEFINE fn, dd, dst STRING, r INT
    CALL ui.interface.frontcall("standard","feInfo",["dataDirectory"],[dd])
    -- Always use / as path sep for Android/iOS dirs.
    LET dst = dd || "/" || os.Path.basename(fn)
    IF base.Application.isMobile() THEN
       -- Executing on device: make a simple copy to data-dir
       LET r = os.Path.copy(fn, dst)
       MESSAGE SFMT("COPY status = %1", r)
    ELSE
       -- Executing on dev server: make a file transfer to data-dir
       CALL fgl_putfile(fn, dst)
    END IF
END FUNCTION
```

We do not use the `os.Path.join()` method here, because it
would add the path separator for the operating system where the application is executed. This would
not be a problem when executing on the mobile device or Unix-like platforms. However, when running
on a Windows® platform, the
`os.Path.join()` method would join the directory and the filename with a backslash,
and the resulting path would not fit Android™ or iOS directory path specification for the data-directory.

In the server application, use the `fgl_getfile()` function, to transfer a file
from the mobile device data-directory to the local server
disk:

```
IMPORT os
...
  CALL server_get_from_data_dir("myfile.txt", "/tmp/server_file.txt")
...
FUNCTION server_get_from_data_dir(fn, dst)
    DEFINE fn, dst, dd, src STRING
    CALL ui.interface.frontcall("standard","feInfo",["dataDirectory"],[dd])
    -- Use / as path sep for Android/iOS dirs!
    LET src = dd || "/" || fn
    CALL fgl_getfile(src, dst)
END FUNCTION
```

Similarly, in the server application, use the `fgl_putfile()` function, to copy
a file from the server application to the data-directory of the embedded
app:

```
IMPORT os
...
  CALL server_put_to_data_dir("/tmp/server_file.txt", "myfile.txt")
...
FUNCTION server_put_to_data_dir(src, fn)
    DEFINE src, fn, dd, dst STRING
    CALL ui.interface.frontcall("standard","feInfo",["dataDirectory"],[dd])
    -- Use / as path sep for Android/iOS dirs!
    LET dst = dd || "/" || fn
    CALL fgl_putfile(src, dst)
END FUNCTION
```

## Related links

**Related concepts**  

[standard.feInfo](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.")

[mobile.runOnServer](../15_library-reference/3458-mobile-runonserver.md "Run an application from the Genero Application Server using the specified URL.")

[Debugging on a mobile device](../13_programming-tools/2579-debugging-on-a-mobile-device.md "It is possible to remotely start the debugger for an app running on a mobile device.")
