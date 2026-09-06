---
title: "Web Services on mobile devices"
source: "fgl-topics/c_fgl_mobile_web_services.html"
breadcrumb: "Mobile applications > Web Services on mobile devices"
type: "concept"
---

# Web Services on mobile devices

> Web Services can be used within mobile applications.

## Web Services and Genero Mobile for Android

Points to consider when using Web Services on Android™ platforms:

- V3 SSL Certificates are required.

For complete details about Web Services on GMA, see [Web Services on GMA (Android)](../16_web-services/4505-web-services-on-gma-android.md "Requirements to use Web services on Android platforms.")

## Web Services and Genero Mobile for iOS

Points to consider when using Web Services on iOS platforms:

- GWS configuration FGLPROFILE entries related to SSL/TLS keys (`security.*`)
  are not supported (uses iOS native SSL/TLS).
- When executing a long-running HTTP request, the app may go into background mode, and raise the
  runtime error -15553 when switching back to foreground mode.

For complete details about Web Services on GMI, see [Web Services on GMI (iOS)](../16_web-services/4504-web-services-on-gmi-ios.md "Requirements to use Web services on iOS platforms (GMI).").
