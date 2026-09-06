---
title: "Accessing device functions"
source: "fgl-topics/c_fgl_mobile_device_functions.html"
breadcrumb: "Mobile applications > Accessing device functions"
type: "concept"
---

# Accessing device functions

> Mobile apps can access device functions by using front calls.

## Accessing device functions using frontcalls

Mobile applications typically want to acces device functions such as geolocation, multi-media
content (photos, videos), messaging (contacts database, email, sms).

This can be easily achieved by using front calls dedicated to mobile features. Note that some
functions are platform specific, for example to launch an Android activity, or access to iOS device
settings.

As a general rule, execute your front call in a `TRY / CATCH` block to catch
errors:

```
DEFINE status STRING,
       latitude, longitude FLOAT
TRY
    CALL ui.Interface.frontCall("standard", "getGeolocation",
         [], [status, latitude, longitude] )
CATCH
    ERROR "Could not get coordinates..."
END TRY
```

For more details, see [Genero Mobile common front calls](../15_library-reference/3441-genero-mobile-common-front-calls.md "This section describes common front calls provided by all mobile front-ends."), [Genero Mobile Android front calls](../15_library-reference/3463-genero-mobile-android-front-calls.md "This section describes front calls specific to the Android platform."), [Genero Mobile iOS front calls](../15_library-reference/3469-genero-mobile-ios-front-calls.md "This section describes front calls specific to the iOS platform.").

## Accessing Android device functions using the Java Interface

On Android
devices, some system functions can only be accessed in the context of a JVM. Use the [Java Interface](../14_extending-the-language/2655-the-java-interface.md "The Java interface allows you to import Java classes and instantiate Java objects in your programs.") with the
`com.fourjs.gma.vm.FglRun` class to access such system specifics. Custom Java classes
need to be part of the .apk package and can be used without any further
configuration.
