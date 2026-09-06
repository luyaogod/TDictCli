---
title: "android.startActivityForResult"
source: "fgl-topics/c_fgl_frontcall_android_startactivityforresult.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile Android™ front calls > android.startActivityForResult"
type: "concept"
description: "Starts an external application (Android activity) and waits until the activity is closed."
---

# android.startActivityForResult

> Starts an external application (Android™ activity) and waits until the activity is closed.

## Syntax

```
ui.Interface.frontCall("android", "startActivityForResult",
  [action, data, category, type, component, extras],
  [outdata, outextras])
```

1. action - Identifies the activity to be started on the Android device.
2. data - (optional) The data to operate on in the activity (URL,
   etc).
3. category - (optional) A comma separated list of categories.
4. type - (optional) Specifies the type of the data passed to the
   activity.
5. component - (optional) Specifies a component class to use for the
   intent.
6. extras - (optional) This is a JSON string containing parameters to
   pass to the activity.

Return values include:

1. outdata - holds the flat value returned by the invoked
   activity.
2. outextras - holds the JSON data of structured value returned by the
   invoked activity.

The return values depend entirely on the invoked activity.

## Usage

The "`startActivityForResult`" front call starts an external application
(Android activity), then waits for the
user to exit the external application prior to returning the GMA application.

> **Important:**
>
> This front call is only available for an
> application running on an Android device.

This front call is similar to the `RUN` statement: The user cannot return to
the GMA application while the activity is executing.

The parameters passed to this front call are used to build an Android "intent" object to start an "activity". For more
details about Android intent object,
refer to [the Android "Intent" definition](http://developer.android.com/reference/android/content/Intent.html).

The action parameter defines the Android activity to perform, such as
`"android.intent.action.MAIN"`, `"android.intent.action.VIEW"`, and so
on.

The data (optional) parameter contains the data to operate on. This is
the main parameter to transmit data to the activity. It can for example be a URL.

The category (optional) parameter contains a comma separated list of
categories, where a category gives additional information about the action to execute. For
example, `"android.intent.category.LAUNCHER"` means it appears in the
Launcher as a top-level application. See the Android documentation for details about possible categories for
a given activity.

The type (optional) parameter defines the type (in fact, a MIME type)
of the activity data. Normally the type is inferred from the data itself. By setting this
attribute, you disable that evaluation and force an explicit type.

The component (optional) parameter defines the name of a component
class to use for the intent. Normally this is determined by looking at the other
information in the intent. The component name typically specified as
"`apk-package-name/java-class-name`"
or "`java-class-name`" (the APK package name is
optional). If the APK package is not specified, GMA considers that the Java class is
included in the current APK.

The extras (optional) parameter specifies a JSON string containing
parameters to pass to the activity. This can be used to provide extended information to
the component. For example, with an action sending an e-mail message, the extra data can
include data to supply a subject, body, for the e-mail.

The outdata returning argument will contain the flag value returned
from the activity, typically when the data is simple and not structured.

The outextras returning argument can hold JSON data of any structured
value returned by the invoked activity, or `NULL` in case of error (for
example, when the application corresponding to the activity is not installed)

## Example

This example invokes the barcode scanner application, and returns the scanned
barcode.

```
IMPORT util
...
DEFINE data, extras STRING,
       json_object util.JSONObject,
       scanned_value STRING
...
CALL ui.Interface.frontCall(
        "android", "startActivityForResult", 
        [ "com.google.zxing.client.android.SCAN",
          NULL, "android.intent.category.DEFAULT" ],
        [ data, extras ])
IF extras IS NULL THEN
  -- If the application isn't installed invoke 
  -- the Play Store to give the user a chance to install it
  CALL ui.Interface.frontCall("standard", "launchurl",
          ["market://details?id=com.google.zxing.client.android"], [])
ELSE
  LET json_object = util.JSONObject.parse(extras)
  -- Fetch the scanned value
  LET scanned_value = json_object.get("SCAN_RESULT")
END IF
```

## Related links

**Related concepts**  

[android.startActivity](3467-android-startactivity.md "Starts an external Android application (activity), and returns to the GMA application immediately.")
