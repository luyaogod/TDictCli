---
title: "gwa.app.TAppVersions type"
source: "fgl-topics/c_gwa_app_api_type_TAppVersions.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.TAppVersions type"
type: "concept"
---

# gwa.app.TAppVersions type

> The TAppVersions type defines a record for describing server and application version information.

## Syntax

```
TYPE TAppVersions RECORD
   serverVersion STRING
   version STRING
   serverBuild gwa.app.TBuildDateTime
   build gwa.app.TBuildDateTime
END RECORD
```

1. `serverVersion` holds the server version
2. `version` holds the application version.
3. `serverBuild` holds the server version's build date-time stamp.
4. `build` holds the current version's build date-time stamp.

## Usage

It provides an API to get server and application version information. A variable of the type
`TAppVersions` must be defined.
