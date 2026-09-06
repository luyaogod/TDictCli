---
title: "android.askForPermission"
source: "fgl-topics/c_fgl_frontcall_android_askforpermission.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile Android™ front calls > android.askForPermission"
type: "concept"
---

# android.askForPermission

> Ask the user to enable a dangerous feature on the Android device.

## Syntax

```
ui.Interface.frontCall("android","askForPermission",
  [permission], [result])
```

1. permission - Identifies the Android™ permission to enable.
2. result - Holds the execution status of the front call:
   - `"ok"` : the user accepted the permission.
   - `"rejected"` : the user refused the permission.

## Usage

The "`askForPermission`" front call opens a message box, to let the
end user confirm the access to a "Dangerous Permission" on Android, to enable a risky feature of the mobile device.

> **Important:**
>
> The `askForPermission`
> front call has been introduced for Android
> 6: Since this version of Android, permissions to access dangerous mobile functions are no longer
> asked during app installation: The app code must explicitly ask the user for dangerous permissions
> when needed, with an `askForPermission` front call.

The permissions parameter defines the Android permission to be asked. It must be a string
representing one of the permission constants, as defined in [Android's Manifest permissions](http://developer.android.com/reference/android/Manifest.permission.html), prefixed by the
`"android.permission."` string. For example, the
`"android.permission.WRITE_EXTERNAL_STORAGE"` string can be used to identify the
permission to access the SDCARD storage unit.

> **Important:**
>
> Android Dangerous Permissions required by the app also need to be specified when building the
> app. For more details, see [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md).

The front call will raise a runtime exception if the permission identifier is not
valid.

## Example

The following code example asks the user to access the SDCARD, and handles the user
choice:

```
DEFINE result STRING
CALL ui.Interface.frontCall(
        "android", "askForPermission", 
        ["android.permission.WRITE_EXTERNAL_STORAGE"],
        [result] )
CASE result
   WHEN "ok"
      CALL os.Path.mkDir("/sdcard/myfiles")
   WHEN "rejected"
      ERROR "SDCARD access was denied by user"
END CASE
```

## Related links

**Related concepts**  

[Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.")
