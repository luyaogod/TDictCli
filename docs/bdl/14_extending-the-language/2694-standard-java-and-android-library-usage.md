---
title: "Standard Java and Android library usage"
source: "fgl-topics/c_fgl_JavaBridge_gma_java.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Executing Java code with GMA > Standard Java and Android library usage"
type: "concept"
description: "You can use Java classes that are part of the standard Java library and Android Java library."
---

# Standard Java and Android library usage

> You can use Java classes that are part of the standard Java library and Android™ Java library.

## Using standard Java within the GMA

Java classes provided in the standard Java library and in the Android Java library can be used
directly by including the `IMPORT JAVA classname` keywords in the
Genero code:

```
IMPORT JAVA java.lang.Runtime
IMPORT JAVA android.os.Build

MAIN
  DEFINE rtm Runtime, msg STRING

  LET rtm = java.lang.Runtime.getRuntime()

  LET msg = SFMT("Device:[%1] %2 - %3 (%4 procs)",
                 android.os.Build.MANUFACTURER,
                 android.os.Build.MODEL,
                 android.os.Build.SERIAL,
                 rtm.availableProcessors() )

  MENU "Test" ATTRIBUTES(STYLE="dialog", COMMENT=msg)
    ON ACTION ok
      EXIT MENU
  END MENU

END MAIN
```

The Android Java library does not include all the classes of a regular JRE. User interface
classes are specific to the Android user interface framework. The list of standard Android Java
packages can be found at <http://developer.android.com/reference/packages.html>.

Only non-interactive classes can be used in this context. To get a graphical user interface, you
must implement an Android Activity, as described in [Implement Android activities in GMA](2696-implement-android-activities-in-gma.md "Android activities can be bundled with your GMA app and called from the Genero code.").

Because Android apps are Java-based, the JVM and standard Java library is directly available.
There is no need to bundle the Java library with your Genero program files when you deploy your app
as .apk package.

When executing the Genero program on a computer in development mode, it is not possible to
use classes that are specific to the Android Java library, because the Android Java library is not
available in development mode at runtime.

You must compile your app code and deploy it on an Android device for execution. To compile your
app code on the development platform, you need to setup the Java SDK environment and the CLASSPATH
to the Android SDK library (android.jar).

## JVM context-dependent Android API calls

On an Android device, the GMA executes a Genero program in a JVM process. Some Android system
APIs cannot be directly called from the Genero runtime system context; they must be called from the
JVM context.

In order to call such APIs, you must import the `com.fourjs.gma.vm.FglRun` class
and get the Android JVM thread context by calling the `getContext()` method of the
`FglRun` class.

The `getContext()` method will return an instance of the
`android.content.Context` class. For more details, see <http://developer.android.com/reference/android/content/Context.html>

To use this Android JVM interface, you must add the android.jar library
(from the Android SDK) to the class path.

The `com.fourjs.gma.vm.FglRun` class implements the following methods:

| Method | Description |
| --- | --- |
| `Context getContext()` | Returns the Android JVM context object of the runtime system. |

In the program code, use the `getContext()` method to get the JVM context and call
specific Android APIs:

```
IMPORT JAVA android.app.Service
IMPORT JAVA android.content.Context
IMPORT JAVA android.util.DisplayMetrics
IMPORT JAVA android.view.WindowManager

IMPORT JAVA com.fourjs.gma.vm.FglRun

MAIN
    DEFINE w, h, d INT
    MENU "Java"
       ON ACTION test
          CALL android_screen_metrics() RETURNING w, h, d
          MESSAGE "Width: ", w, "\nHeight: ", h, "\nDensity: ", d
    END MENU
END MAIN

FUNCTION android_screen_metrics()
    DEFINE ctx android.content.Context,
           dm android.util.DisplayMetrics,
           wm android.view.WindowManager

    LET ctx = com.fourjs.gma.vm.FglRun.getContext()
    LET dm = android.util.DisplayMetrics.create()
    LET wm = CAST ( ctx.getSystemService("window")
                    AS android.view.WindowManager )
    CALL wm.getDefaultDisplay().getMetrics(dm)

    RETURN dm.widthPixels,
           dm.heightPixels,
           dm.densityDpi
END FUNCTION
```

## Using front calls instead of pure Java

For maximum portability, consider implementing Android-specific extensions as custom front calls.
When using the front call technology, apps can be executed in development (app running on the
server) and in deployed mode (app running on the mobile device) with the same Genero code.

For more details about custom front calls with GMA, see [Implement front call modules for GMA](2720-implement-front-call-modules-for-gma.md "Custom front call modules for the Android front-end are implemented by using the API for GMA front calls in Java.").
