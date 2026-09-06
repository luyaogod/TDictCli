---
title: "gwa.app.TCheckUpToDateResult type"
source: "fgl-topics/c_gwa_app_api_type_TCheckUpToDateResult.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.TCheckUpToDateResult type"
type: "concept"
---

# gwa.app.TCheckUpToDateResult type

> The TCheckUpToDateResult type defines a record for checking the up-to-date status of the application.

## Syntax

```
TYPE TCheckUpToDateResult RECORD
  upToDate BOOLEAN
  versions gwa.app.TAppVersions
  error STRING
END RECORD
```

1. `upToDate` holds the boolean value for the current
   status of the application.
2. `versions` holds the application version.
3. `error` holds an error message.

## Usage

It provides a type for checking the up-to-date status of the application.
