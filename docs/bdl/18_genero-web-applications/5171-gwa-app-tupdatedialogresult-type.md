---
title: "gwa.app.TUpdateDialogResult type"
source: "fgl-topics/c_gwa_app_api_type_TUpdateDialogResult.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module > gwa.app.TUpdateDialogResult type"
type: "concept"
---

# gwa.app.TUpdateDialogResult type

> The TUpdateDialogResult type defines a record for returning the result of the up-to-date status of the application on the server.

## Syntax

```
TYPE TUpdateDialogResult RECORD
  upToDate BOOLEAN
  updatePerformed BOOLEAN
  error STRING
END RECORD
```

1. `upToDate` holds the boolean value for the current
   status of the application.
2. `updatePerformed` holds the boolean value for the status of the update.
3. `error` holds an error message.

## Usage

It provides a type for the result in [gwa.app.updateDialog](5182-gwa-app-updatedialog.md "Checks the version of the loaded application against the server version and prompts the user to update if required.").
