---
title: "mobile.choosePhoto"
source: "fgl-topics/c_fgl_frontcall_mobile_choosephoto.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.choosePhoto"
type: "concept"
---

# mobile.choosePhoto

> Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier.

## Syntax

```
ui.Interface.frontCall("mobile", "choosePhoto",
   [], [path])
```

1. path - Holds the device opaque path to the chosen photo.

## Usage

When displaying the application a mobile device, the "`choosePhoto`" front call
starts the system's photo chooser (the device's photo gallery), allows the user to choose a photo,
and returns the path/URL on the mobile device of the chosen photo.

On desktop, this front call will open the file selector, to chose a
file from the storage unit.

If the user cancels the photo chooser, `NULL` is returned.

> **Important:**
>
> For GMA / Android™,
> using the `choosePhoto` front call needs the
> `android.permission.READ_EXTERNAL_STORAGE` Dangerous Permission to be specified when
> building the APK. See
> [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more
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

[mobile.takePhoto](3460-mobile-takephoto.md "Lets the user take a picture with the mobile device and returns the corresponding picture identifier.")
