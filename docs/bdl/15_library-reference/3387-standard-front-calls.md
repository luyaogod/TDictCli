---
title: "Standard front calls"
source: "fgl-topics/c_fgl_frontcalls_standard.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls"
type: "concept"
---

# Standard front calls

> Standard front call functions provide common utility APIs to control the front-end.

This table shows the functions implemented by the front-ends in the "`standard`"
module, available on all front-ends.

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("standard", "cbAdd", [text], [result]) | Adds to the content of the clipboard. |
| ui.Interface.frontCall("standard", "cbClear", [], [result]) | Clears the content of the clipboard. |
| ui.Interface.frontCall("standard", "cbGet", [], [text]) | Gets the content of the clipboard. |
| ui.Interface.frontCall("standard", "cbPaste", [], [result]) | Pastes the content of the clipboard to the current field. |
| ui.Interface.frontCall("standard", "cbSet", [text], [result]) | Set the content of the clipboard. |
| ui.Interface.frontCall("standard", "clearFileCache", [], [result]) | Clears the local file cache. |
| ui.Interface.frontCall("standard", "clearNotifications", [options], [status] ) | Drops notifications displayed on the host system of the front end. |
| ui.Interface.frontCall("standard", "composeMail", [to, subject, content, cc, bcc], [result]) | Invokes the user's default mail application for a new mail to send. |
| ui.Interface.frontCall("standard", "connectivity", [], [result] ) | Returns the type of network available for the device. |
| ui.Interface.frontCall("standard", "createNotification", [options], [id] ) | Creates a new local notification to be displayed on the host system of the front end. |
| ui.Interface.frontCall("standard", "execute", [cmd,wait], [result]) | Executes a command on the front-end platform, with or without waiting. |
| ui.Interface.frontCall("standard", "feInfo", [name], [result]) | Queries general front-end properties. |
| ui.Interface.frontCall("standard", "getEnv", [name], [value]) | Returns an environment variable set in the user session on the front end platform. |
| ui.Interface.frontCall("standard", "getGeolocation", [], [status, latitude, longitude] ) | Returns the Global Positioning System (GPS) location of a device. |
| ui.Interface.frontCall("standard", "getLastNotificationInteractions", [], [notifications] ) | Report all user interactions with notifications since last call to this function. |
| ui.Interface.frontCall("standard", "hardCopy", [pgsize], [result]) | Prints a screenshot of the current window |
| ui.Interface.frontCall("standard", "isForeground", [], [result] ) | Indicates if the app is in foreground mode. |
| ui.Interface.frontCall("standard", "launchURL", [ url [, mode ] ], [] ) | Opens a URL with the default URL handler of the front-end. |
| ui.Interface.frontCall("standard", "openDir", [path,caption], [result]) | Displays a file dialog window to get a directory path on the local file system. |
| ui.Interface.frontCall("standard", "openFile", [path,name,wildcards,caption], [result]) | Displays a file dialog window to let the user select a single file path on the local file system. |
| ui.Interface.frontCall("standard", "openFiles", [path,name,wildcards,caption], [result]) | Displays a file dialog window to let the user select a list of file paths on the local file system. |
| ui.Interface.frontCall("standard", "playSound", [ resource [, wait ] ], []) | Plays the sound file passed as parameter on the front-end platform. |
| ui.Interface.frontCall("standard", "restoreSize", [delay], [result]) | Asks GDC to restore the stored window container size. |
| ui.Interface.frontCall("standard", "storeSize", [], [result]) | Asks GDC to store the size of the current window container. |
| ui.Interface.frontCall("standard", "saveFile", [path,name,filetype,caption], [result]) | Displays a file dialog window to get a path to save a file on the local file system. |
| ui.Interface.frontCall("standard", "shellExec", [document, action], [result]) | Opens a file on the front-end platform with the program associated to the file extension. |
