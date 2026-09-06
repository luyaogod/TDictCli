---
title: "mobile.chooseVideo"
source: "fgl-topics/c_fgl_frontcall_mobile_choosevideo.html"
breadcrumb: "Library reference > Built-in front calls > Genero Mobile common front calls > mobile.chooseVideo"
type: "concept"
---

# mobile.chooseVideo

> Lets the user select a video from the mobile device's video gallery and returns a video identifier.

## Syntax

```
ui.Interface.frontCall("mobile", "chooseVideo",
  [], [path])
```

1. path - Holds the device opaque path to the selected video.

## Usage

When displaying the application a mobile device, the "`chooseVideo`" front call
starts the system's video chooser (the device's video gallery), allows the user to choose a video,
and returns the path/URL on the mobile device of the selected video.

On desktop, this front call will open the file selector, to chose a
file from the storage unit.

If the user cancels the video chooser, `NULL` is returned.

> **Important:**
>
> For GMA / Android™,
> using the `chooseVideo` front call needs the
> `android.permission.READ_EXTERNAL_STORAGE` Dangerous Permission to be specified when
> building the APK. See [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md) for more
> details.

The value returned in the path variable contains a reference to the system
location of the video on the mobile device. This path is platform dependent, and may change in
future versions. Consider the path returned by this front call as an opaque local file identifier,
and do not use this path as a persistent filename for the video.

Once the video identifier/path is known, it is possible to fetch the video file from the device
to the program context with the `fgl_getfile()` API. The procedure is similar to
fetching photos from the device. For more details, see the section about [video handling on
mobile devices](../11_user-interface/1588-runtime-images.md).

To play the video, you can perform a ["`launchURL`"](3406-standard-launchurl.md "Opens a URL with the default URL handler of the front-end.") front call, with the opaque path returned by this front call.

## Related links

**Related concepts**  

[Using images](../11_user-interface/1583-using-images.md "Describes how to use pictures in the forms of your application.")

[fgl\_getfile()](2762-fgl-getfile.md "Retrieves a file from the front-end context to the virtual machine context.")

[mobile.takeVideo](3461-mobile-takevideo.md "Lets the user take a video with the mobile device and returns the corresponding video identifier.")
