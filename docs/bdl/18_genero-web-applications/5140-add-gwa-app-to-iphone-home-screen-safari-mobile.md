---
title: "Add GWA app to iPhone home screen (Safari Mobile)"
source: "fgl-topics/c_fgl_gwa_safari_mobile_homescreen.html"
breadcrumb: "Genero Web applications > Deploying GWA apps > Add GWA app to iPhone® home screen (Safari® Mobile)"
type: "concept"
description: "There are special considerations when adding a Genero Web Application to the home screen from Safari Mobile."
---

# Add GWA app to iPhone home screen (Safari Mobile)

> There are special considerations when adding a Genero Web Application to the home screen from Safari® Mobile.

When you add a GWA app to the iPhone®
home screen from Safari Mobile, iOS
launches it in a separate browser instance. The first time you open the app from the home screen it
hasn’t been cached there yet, so you need network access to load the same GWA app you previously
opened in Safari Mobile. After it finishes
loading, the app will run offline from the home screen — and on all subsequent launches from the
home screen — even if the network is turned off.

> **Note:**
>
> This behavior occurs only when you add the GWA app to the iPhone home screen from Safari Mobile. It does not occur when you add the app to Safari on desktop or to other browsers.
