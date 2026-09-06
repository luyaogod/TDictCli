---
title: "Front-end termination"
source: "fgl-topics/c_fgl_programs_010.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime) > Front-end termination"
type: "concept"
---

# Front-end termination

> The OPTIONS CLOSE APPLICATION instruction defines the callback function in case of front-end termination.

## Syntax

```
OPTIONS ON CLOSE APPLICATION CALL function
```

1. function is the name of a function with no parameters or return value.

## Usage

The `OPTIONS ON CLOSE APPLICATION CALL function` can be used to
execute specific code when the front-end sends an application close event. This event can occur in
the following cases:

1. With the Genero Desktop Client (GDC) in direct mode:
   - When the specific application connection in the GDC control panel is closed.
   - When the GDC front-end program is closed. This occurs also when the OS user sessions is ended,
     or when workstation OS is shutdown.
2. With the Genero Application Server (GAS) and a web browser:
   - When the Genero Application Server is shutdown.
   - When the browser tab containing application forms is closed or when the browser is closed. This
     occurs also when the OS user sessions is ended, or when workstation OS is shutdown.
     > **Note:**
     >
     > With GAS, depending on the type of the browser, when closing the tab or window, in best case the
     > program termination occurs immediately. However, the program termination may occur only after a
     > timeout defined by the `USER_AGENT` .xcf configuration
     > parameter. In this configuration, best practice for the end users is to trigger the application
     > action to terminate the program ("Quit" or "Close" button)
3. With Genero Mobile for Android (GMA), in development direct mode, runOnServer and embedded apps:
   - When the front-end or embedded app is closed. This occurs also when mobile device OS is
     shutdown.
   - The event is not send when the front-end or embedded app goes in background mode.
4. With Genero Mobile for iOS (GMI), this callback feature is not supported.

Before stopping, the front-end sends an internal event that is trapped by the runtime system.
When a callback function is specified with this program option command, the application code that
was executing is canceled, and the callback function is executed before the program stops.

Use the `OPTIONS ON CLOSE APPLICATION CALL function`
instruction with care, and do not execute complex code in the callback function. The code is expected to contain only simple and short cleanup operations; any interactive
instruction is avoided.

> **Important:**
>
> A front-end program crash or network failure is not detected and cannot be
> handled by the `ON CLOSE APPLICATION` option.

## Example

```
MAIN
    OPTIONS ON CLOSE APPLICATION CALL end_app
    MENU "test"
        COMMAND "quit" EXIT MENU
    END MENU
END MAIN

FUNCTION end_app()
    DISPLAY "Program terminated by front-end..."
END FUNCTION
```

## Related links

**Related concepts**  

[GUI front-end connection](../11_user-interface/1520-gui-front-end-connection.md "This section explains runtime to front-end connection in its simplest form.")
