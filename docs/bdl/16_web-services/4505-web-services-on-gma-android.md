---
title: "Web Services on GMA (Android)"
source: "fgl-topics/c_gws_mobile_gma_requirements.html"
breadcrumb: "Web services > General > Platform-specific notes > Web Services on GMA (Android™)"
type: "concept"
description: "Requirements to use Web services on Android platforms."
---

# Web Services on GMA (Android)

> Requirements to use Web services on Android™ platforms.

## V3 SSL Certificates

The SSL certificates for secured servers must be of type V3: Android does not support other types of SSL certificates. When creating
your own self-signed certificates (to be installed in the "Install from storage" Keystore of
Android), make sure that type V3 is used.
