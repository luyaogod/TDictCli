---
title: "android.startActivity"
source: "fgl-topics/c_fgl_frontcall_android_startactivity.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile Android™ front calls > android.startActivity"
type: "concept"
description: "Starts an external Android application (activity), and returns to the GMA application immediately."
---

# android.startActivity

> Starts an external Android™ application (activity), and returns to the GMA application immediately.

## Syntax

```
ui.Interface.frontCall("android","startActivity",
  [action, data, category, type, component, extras],
  [])
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

## Usage

The "`startActivity`" front call starts an external application (Android activity), and returns
to the GMA application immediately after invoking the activity.

> **Important:**
>
> This front call is only available for an
> application running on an Android device.

This front call is similar to the `RUN WITHOUT WAITING` statement: It allows
the user to switch between the GMA and the launched application.

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

## Example

The following code example starts the VIEW Android activity to show an image. The Genero program flow will continue
after this call, but the started activity will be shown. Note that such action is better performed
with a [`launchurl`](3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") front
call.

```
CALL ui.Interface.frontCall(
        "android", "startActivity", 
        [ "android.intent.action.VIEW",
          "file:///storage/path_to_image_file",
          NULL, "image/*" ],
        [ ] )
```

## Related links

**Related concepts**  

[android.startActivityForResult](3468-android-startactivityforresult.md "Starts an external application (Android activity) and waits until the activity is closed.")
