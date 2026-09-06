---
title: "mobile.takePhoto"
source: "fgl-topics/c_fgl_frontcall_mobile_takephoto.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.takePhoto"
type: "concept"
---

# mobile.takePhoto

> Lets the user take a picture with the mobile device and returns the corresponding picture identifier.

## Syntax

```
ui.Interface.frontCall("mobile", "takePhoto",
   [], [path] )
```

1. path - Holds the device opaque path to the picture that has been taken.

## Usage

When displaying the application a mobile device, the "`takePhoto`" front call
invokes the mobile device's camera to let the user take a picture and returns the local path/URL on
the mobile device to the picture.

> **Important:**
>
> To be usable with GBC, this front call needs a secure context / front-end
> connection, in addition to mobile device permissions, when required. A secure context
> is achieved by using the GAS via HTTPS, via localhost on the same machine, or running direct via
> GDC-UR. When using GDC-UR, additional security settings need to be enabled in the GDC configuration
> panel. See GDC documentation for more details.

On desktop, this front call will open the file selector, to chose a
file from the storage unit.

If the photo is canceled by the user, the front call returns `NULL`.

> **Important:**
>
> For GMA / Android™,
> using the `takePhoto` front call needs the
> `android.permission.CAMERA`,
> `android.permission.READ_EXTERNAL_STORAGE` and
> `android.permission.WRITE_EXTERNAL_STORAGE` Dangerous Permissions to be specified
> when building the APK. See [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more
> details.

The value returned in the path variable contains a reference to the system
location of the picture on the mobile device. This path is platform dependent, and may change in
future versions. Consider the path returned by this front call as an opaque local file identifier,
and do not use this path as a persistent filename for the picture.

Once the photo identifier/path is known, it is possible to fetch the photo file from the device
to the program context with the `fgl_getfile()` API. The procedure is similar to
fetching photos from the device. For more details about mobile image handling, see [images handling on
mobile devices](../11_user-interface/1588-runtime-images.md).

## Related links

**Related concepts**  

[Using images](../11_user-interface/1583-using-images.md "Describes how to use pictures in the forms of your application.")

[fgl\_getfile()](2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.")

[mobile.choosePhoto](3443-mobile-choosephoto.md "Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier.")
