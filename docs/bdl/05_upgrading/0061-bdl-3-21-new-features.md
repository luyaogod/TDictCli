---
title: "BDL 3.21 new features"
source: "fgl-topics/fgl_whatsnew_321.html"
breadcrumb: "Upgrading > New features of Genero BDL > BDL 3.21 new features"
type: "topic"
---

# BDL 3.21 new features

> Features added in 3.21 releases of the Genero Business Development Language.

> **Important:**
>
> This page covers only those new features
> introduced with the Genero BDL version specified in the page title. Check prior new features pages
> if you migrate from an earlier version. Make sure to also read the upgrade guide corresponding to
> this Genero version.

Corresponding upgrade guide: [BDL 3.21 upgrade guide](0153-bdl-3-21-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.21.").

Prior new features guide: [BDL 3.20 new features](0062-bdl-3-20-new-features.md "Features added in 3.20 releases of the Genero Business Development Language.").

| Overview | Reference |
| --- | --- |
| Support for license manager version 6. | See Install and License your Genero Products. |

| Overview | Reference |
| --- | --- |
| **Starting at 3.21.02** |  |
| The FGLPROFILE has a new entry, `security.global.options`, added to support legacy OpenSSL 1 options. | See [New security.global.option in FGLPROFILE to allow legacy OpenSSL 1 options](0155-web-services-changes.md) |
| New fglwsdl option `-SSLOptions` added to support legacy OpenSSL 1 options. | See [fglwsdl option -SSLOptions to support legacy OpenSSL 1 options](0155-web-services-changes.md) and [fglwsdl](../13_programming-tools/2523-fglwsdl.md). |
| `STRING`-typed return parameter in OAUTH API function. | See [Change to OAuthAPI.GetIDSubject returns](0155-web-services-changes.md). |
| The FGLPROFILE has a new entry, `security.global.certificate.selfsigned.preload`, added to preload the global certificate and key at the start of the application instead of at the first HTTPS connection. | See [New security.global.certificate.selfsigned.preload entry in FGLPROFILE](0155-web-services-changes.md) |
| [***Related upgrade notes***](0153-bdl-3-21-upgrade-guide.md "These topics describe product changes you must be aware of when upgrading to version 3.21.") |  |
| Upgrade notes for web services. | See [Web Services changes](0155-web-services-changes.md "There are changes in support of web services in Genero 3.21.") |
