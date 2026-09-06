---
title: "gwa.app module"
source: "fgl-topics/r_gwa_app_api_functions.html"
breadcrumb: "Genero Web applications > GWA Reference > Genero Web Application API > The gwa.app module > gwa.app module"
type: "reference"
---

# gwa.app module

> Functions and types of the gwa.app module for querying version numbers and for working with application updates.

| Types | Description |
| --- | --- |
| TYPE TAppInformation RECORD title STRING version STRING build gwa.app.TBuildDateTime gwaVersion STRING embeddedVMVersion STRING fgl_getversion STRING gbcVersion STRING server STRING END RECORD | The TAppInformation type defines a record for returning application information. |
| TYPE TAppVersions RECORD serverVersion STRING version STRING serverBuild gwa.app.TBuildDateTime build gwa.app.TBuildDateTime END RECORD | The TAppVersions type defines a record for describing server and application version information. |
| TYPE TBuildDateTime DATETIME YEAR TO FRACTION | The TBuildDateTime type defines a record for the build date time stamp. |
| TYPE TCheckUpToDateResult RECORD upToDate BOOLEAN versions gwa.app.TAppVersions error STRING END RECORD | The TCheckUpToDateResult type defines a record for checking the up-to-date status of the application. |
| TYPE TServerVersionAndBuildResult RECORD serverVersion STRING serverBuild gwa.app.TBuildDateTime error STRING END RECORD | The TServerVersionAndBuildResult type defines a record for checking the up-to-date status of the application on the server. |
| TYPE TUpdateDialogResult RECORD upToDate BOOLEAN updatePerformed BOOLEAN error STRING END RECORD | The TUpdateDialogResult type defines a record for returning the result of the up-to-date status of the application on the server. |

| Function | Description |
| --- | --- |
| FUNCTION checkUpToDate() RETURNS gwa.app.TUpdateDialogResult | Checks if the application is up to date by comparing versions and date time stamps. |
| FUNCTION getBuild() RETURNS gwa.app.TBuildDateTime | Returns the date time stamp of when the application was built. |
| FUNCTION getGBCVersion() RETURNS STRING | Returns the version of the GBC embedded in the application. |
| FUNCTION getInformation() RETURNS gwa.app.TAppInformation | Returns the application version information. |
| FUNCTION getServerVersionAndBuild() RETURNS gwa.app.TServerVersionAndBuildResult | Returns server application version information. |
| FUNCTION getTitle() RETURNS STRING | Returns the application title set by `gwabuildtool --title`. |
| FUNCTION getVersion() RETURNS STRING | Returns the application version set by `gwabuildtool --app-version`. |
| FUNCTION performUpdate() RETURNS STRING | Update the application. |
| FUNCTION set_href_to_new_index() RETURNS STRING | Set the href of the index page. |
| FUNCTION simulateBrowserUpdate() RETURNS STRING | Simulate a browser update (service worker update). |
| FUNCTION updateDialog(reportWhenUpToDate BOOLEAN) RETURNS gwa.app.TUpdateDialogResult | Checks the version of the loaded application against the server version and prompts the user to update if required. |
| FUNCTION versionsAndBuildsEqual( version1 STRING, version2 STRING, build1 gwa.app.TBuildDateTime, build2 gwa.app.TBuildDateTime) RETURNS BOOLEAN | Checks if the versions and build date-time stamps are equal. If versions are `NULL`, only build date-time stamps are compared for equality. |

## Child topics

- [gwa.app.TAppInformation type](5166-gwa-app-tappinformation-type.md): The TAppInformation type defines a record for returning application information.
- [gwa.app.TAppVersions type](5167-gwa-app-tappversions-type.md): The TAppVersions type defines a record for describing server and application version information.
- [gwa.app.TBuildDateTime type](5168-gwa-app-tbuilddatetime-type.md): The TBuildDateTime type defines a record for the build date time stamp.
- [gwa.app.TCheckUpToDateResult type](5169-gwa-app-tcheckuptodateresult-type.md): The TCheckUpToDateResult type defines a record for checking the up-to-date status of the application.
- [gwa.app.TServerVersionAndBuildResult type](5170-gwa-app-tserverversionandbuildresult-type.md): The TServerVersionAndBuildResult type defines a record for checking the up-to-date status of the application on the server.
- [gwa.app.TUpdateDialogResult type](5171-gwa-app-tupdatedialogresult-type.md): The TUpdateDialogResult type defines a record for returning the result of the up-to-date status of the application on the server.
- [gwa.app.checkUpToDate](5172-gwa-app-checkuptodate.md): Checks if the application is up to date by comparing versions and date time stamps.
- [gwa.app.getBuild](5173-gwa-app-getbuild.md): Returns the date time stamp of when the application was built.
- [gwa.app.getGBCVersion](5174-gwa-app-getgbcversion.md): Returns the version of the GBC embedded in the application.
- [gwa.app.getInformation](5175-gwa-app-getinformation.md): Returns the application version information.
- [gwa.app.getServerVersionAndBuild](5176-gwa-app-getserverversionandbuild.md): Returns server application version information.
- [gwa.app.getTitle](5177-gwa-app-gettitle.md): Returns the application title set by gwabuildtool --title.
- [gwa.app.getVersion](5178-gwa-app-getversion.md): Returns the application version set by gwabuildtool --app-version.
- [gwa.app.performUpdate](5179-gwa-app-performupdate.md): Update the application.
- [gwa.app.set_href_to_new_index](5180-gwa-app-set-href-to-new-index.md): Set the href of the index page.
- [gwa.app.simulateBrowserUpdate](5181-gwa-app-simulatebrowserupdate.md): Simulate a browser update (service worker update).
- [gwa.app.updateDialog](5182-gwa-app-updatedialog.md): Checks the version of the loaded application against the server version and prompts the user to update if required.
- [gwa.app.versionsAndBuildsEqual](5183-gwa-app-versionsandbuildsequal.md): Checks if the versions and build date-time stamps are equal. If versions are NULL, only build date-time stamps are compared for equality.
