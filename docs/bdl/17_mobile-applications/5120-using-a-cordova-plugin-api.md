---
title: "Using a Cordova plugin API"
source: "fgl-topics/c_fgl_cordova_plugin_api.html"
breadcrumb: "Mobile applications > Cordova plugins > Using a Cordova plugin API"
type: "concept"
---

# Using a Cordova plugin API

> Cordova plugin features can be used by invoking cordova front calls.

## How to access Cordova plugin APIs

Cordova plugin APIs are available to the front-end and therefore need to be called through [Cordova plugin front calls](../15_library-reference/3472-cordova-plugin-front-calls.md "Genero provides a set of Cordova plugin front calls that make use of the Cordova plugins.").

Before calling a Cordova plugin function, you need to identify what native APIs are available. If
provided, check the Cordova plugin API documentation for available functions.

It is a good practice to implement BDL wrapper functions that encapsulate the Cordova front calls
specific to this plugin. BDL wrapper libraries are provided with the plugins on the [FOURJS Cordova
GitHub](https://github.com/FourjsGenero-Cordova-Plugins).

## Identifying and verifying the version of available Cordova plugins at runtime

At runtime, the code of an app can query available Cordova plugins with the [`cordova.listPlugins`](../15_library-reference/3479-cordova-listplugins.md "Returns the list of available Cordova plugins.") and get
details of the plugin with [`cordova.getPluginInfo`](../15_library-reference/3478-cordova-getplugininfo.md "Returns details about a specific Cordova plugin.") front calls.

To verify the git version in the wrapper modules for usage in direct mode and GAS mode, implement
an initialization function in the wrapper module, to check that the `GIT_VERSION`
string defined in the git\_version.txt file (found nearby the wrapper library
sources) corresponds to the current available plugin version returned by the
`cordova.getPluginInfo` front
call:

```
&include "git_version.txt"  -- defines GIT_VERSION
...
DEFINE initialized BOOLEAN
...
PUBLIC FUNCTION initialize()
    IF initialized THEN
        -- exclusive library usage
        DISPLAY "The library is already in use."
        EXIT PROGRAM 1
    END IF
    CALL checkGitVersion(GIT_VERSION)
    LET initialized = TRUE
END FUNCTION
...
PRIVATE FUNCTION checkGitVersion(version STRING)
  DEFINE rec RECORD
    git_version STRING
  END RECORD
  CALL ui.interface.frontcall("cordova","getPluginInfo",["Accelerometer"],[rec])
  IF NOT version.equals(rec.git_version) THEN
      DISPLAY sfmt("fglcdvMotion.4gl version mismatch, git version:%1, expecting:%2.",
                   rec.git_version, version)
      EXIT PROGRAM 1
  END IF
END FUNCTION
```

## Identifying Cordova plugin native APIs in JS sources

Cordova plugins are primarily targeted for JavaScript. Each plugin source directory has a
www directory with JavaScript wrapper code calling into the native code.

For example, with the `cordova-plugin-network-information` plugin, the
`www/network.js` file implements the following `"getInfo"` JavaScript
function, that performs an `exec()` call of the
`"NetworkStatus.getConnectionInfo"` native Cordova function:

```
NetworkConnection.prototype.getInfo =
    function(successCallback, errorCallback) {
        exec(successCallback, errorCallback,
             "NetworkStatus", "getConnectionInfo", []);
    };
```

The BDL equivalent of the above `exec()` call is done with the generic Cordova
front call ["`cordova.call`"](../15_library-reference/3473-cordova-call.md "Calls a function in a Cordova plugin and returns a result."):

```
CALL ui.Interface.frontCall( "cordova", "call",
        ["NetworkStatus","getConnectionInfo"], [result] )
```

> **Note:**
>
> In the README.md files of the plugins, the available methods for the
> plugins are documented for the JavaScript wrappers and not the underlying native calls.

In order to find available native Cordova APIs, such as `getConnectionInfo`,
search for the `exec()` calls in the js files of the www
directory. The native Cordova function name is the fourth parameter of the `exec()`
function.

## Implementing an app that uses the Cordova plugin

After installing a plugin, you can create a Genero program using a function of this plugin, by
using the generic front call for Cordova APIs.

For example:

```
IMPORT FGL fgldialog
MAIN
 DEFINE result STRING
 MENU "network"
   COMMAND "Info"
     CALL ui.Interface.frontCall( "cordova", "call",
             ["NetworkStatus","getConnectionInfo"], [result] )
     CALL fgl_winMessage("Result",result,"info")
   COMMAND "Exit"
     EXIT MENU
 END MENU
END MAIN
```

Compile and build your app as usual (the installed Cordova plugins will be included).

## Cordova plugin front calls overview

Cordova plugin front calls are generic functions that give access to the native Cordova APIs, in
conjunction with the predefined `cordovacallback` action.

The [`cordova.call`](../15_library-reference/3473-cordova-call.md "Calls a function in a Cordova plugin and returns a result.") function
issues a synchronous Cordova function call. In this case, the underlying native function is executed
and the program waits until it returns. Results are provided in the front call output
parameters.

The [`cordova.callWithoutWaiting`](../15_library-reference/3474-cordova-callwithoutwaiting.md "Calls a function asynchronously in a Cordova plugin, without waiting for a result.") function performs an asynchronous Cordova
function call. The program continues while the Cordova function executes in parallel. The program
can implement a trigger for the `cordovacallback` predefined action, to detect when
results are available. To retrieve results of asynchronous calls, use the [`cordova.getCallbackDataCount`](../15_library-reference/3476-cordova-getcallbackdatacount.md "Returns the number of pending Cordova plugin results."), [`cordova.getCallbackData`](../15_library-reference/3477-cordova-getcallbackdata.md "Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue.")
and [`cordova.getAllCallbackData`](../15_library-reference/3475-cordova-getallcallbackdata.md "Returns all results for asynchronous Cordova plugin front calls, based on a callback ID filter.") front calls.

See these front calls in action in the demos provided on the [FOURJS Cordova
GitHub](https://github.com/FourjsGenero-Cordova-Plugins).

## Implement BDL wrappers on top of Cordova front calls

For better code readability and maintenance, we strongly suggest you implement BDL functions that
encapsulate Cordova front calls, as done with plugins available on the [FOURJS Cordova
GitHub](https://github.com/FourjsGenero-Cordova-Plugins).

## Android™ Cordova plugin API internals

The Android code is located in the
src/Android directory of a plugin.

The code in src/android/NetworkManager.java contains an
`execute` method for handling the `getConnectionInfo` string:

```
public boolean execute(String action, JSONArray args, CallbackContext callbackContext) {
 if (action.equals("getConnectionInfo")) {
```

## IOS® Cordova plugin API internals

The IOS code is located in the
src/ios directory of a plugin.

The code in src/ios/CDVConnection.m contains a method called

```
(void)getConnectionInfo:(CDVInvokedUrlCommand*)command
```

This is the native objective-C method called by the front call.

## Debug Cordova plugins

To debug your application using a Cordova plugin, build your project with the plugin enabled.
This process is outlined in [Including Cordova plugins in your Android APK package](5121-embed-cordova-plugins-in-a-gma-app.md).
After building your project, move the APK file to the external file directory on your device. This
path is outlined in the development client. Alternatively, to find the path to the external file
directory, use the `getexternalstoragepath` frontcall command. Moving the APK file to
this directory allows the development client to recognize Cordova plugins packaged in it. This makes
it possible to launch and develop your application in remote mode. With a device running on Android 14, for security reasons, this APK
file needs to be read-only (file permissions `"-r--r--r--"`). If the APK file is not
read only, the Cordova plugin cannot be loaded.

To verify the plugins are recognized, use the `ListPlugins` and
`GetPluginInfo` frontcall commands to return lists and information on externally
found plugins.
> **Note:**
>
> You must rebuild the project when your application is ready to get the final APK. This contains
> both the plugin and the final code.
