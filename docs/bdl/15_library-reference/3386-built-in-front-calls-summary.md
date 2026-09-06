---
title: "Built-in front calls summary"
source: "fgl-topics/c_fgl_frontcalls_genero.html"
breadcrumb: "Library reference > Built-in front calls > Built-in front calls summary"
type: "concept"
---

# Built-in front calls summary

> Various front-end functions are implemented within Genero front-ends.

This section describes the front-end functions available for all types of front-ends.
Note that several front-end functions are specific to the type of front-end.

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

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("table", "autoFitAllColumns", [screen-record], [] ) | Adapts the width of table columns to the displayed data. |
| ui.Interface.frontCall("table", "fitToViewAllColumns", [screen-record], [] ) | Adapts the width of table columns to show all columns. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("webcomponent", "call", [aui-name, function-name [, param1, param2, ... ] ], [result] ) | Calls a JavaScript function through the web component. |
| ui.Interface.frontCall("webcomponent", "frontCallAPIVersion", [],[result]) | Returns the API version of web component front-end calls. |
| ui.Interface.frontCall("webcomponent", "getTitle", [aui-name], [result] ) | Returns the title of the HTML doc rendered by a web component. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("monitor", "update", [ path-to-update-file [,warning-text [,elevation-prompt] ] ], [ result ]) | Starts the GDC update. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("theme", "setTheme", [name], []) | Activates a specific theme. |
| ui.Interface.frontCall("theme", "getCurrentTheme", [], [result]) | Gets the active theme. |
| ui.Interface.frontCall("theme", "listThemes", [], [result]) | Lists all available themes. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("monitor", "update", [ path-to-update-file [,warning-text [,elevation-prompt] ] ], [ result ]) | Starts the GDC update. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("browser", "setApplicationState", [anchor], []) | Sets the `#` anchor of the URL in the browser address bar. |
| ui.Interface.frontCall("browser", "getApplicationState", [], [anchor]) | Gets the `#` anchor of the current URL in the browser address bar. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("mobile", "chooseContact", [], [result]) | Lets the user choose a contact from the mobile device contact list and returns the vCard. |
| ui.Interface.frontCall("mobile", "choosePhoto", [], [path]) | Lets the user select a picture from the mobile device's photo gallery and returns a picture identifier. |
| ui.Interface.frontCall("mobile", "chooseVideo", [], [path]) | Lets the user select a video from the mobile device's video gallery and returns a video identifier. |
| ui.Interface.frontCall("mobile", "createNotification", [options], [id] ) | Creates or updates a local notification to be displayed on the mobile device. |
| ui.Interface.frontCall("mobile", "clearNotifications", [options], [status] ) | Drops notifications displayed on the mobile device. |
| ui.Interface.frontCall("mobile", "composeMail", [to, subject, content, cc, bcc, attachments ...], [result]) | Invokes the user's default mail application for a new mail to send. |
| ui.Interface.frontCall("mobile", "composeSMS", [ recipients, content ], [ result ] ) | Sends an SMS text to one or more phone numbers. |
| ui.Interface.frontCall("mobile", "connectivity", [], [result] ) | Returns the type of network available for the mobile device. |
| ui.Interface.frontCall("mobile", "getGeolocation", [], [status, latitude, longitude] ) | Returns the Global Positioning System (GPS) location of a mobile device. |
| ui.Interface.frontCall("mobile", "getLastNotificationInteractions", [], [notifications] ) | Get the last user interactions on mobile app notifications. |
| ui.Interface.frontCall("mobile","getRemoteNotifications", [], [data] ) | This front call retrieves push notification messages. |
| ui.Interface.frontCall("mobile", "importContact", [vcard], [result] ) | Creates a new contact, or merges to an existing entry, the contact details passed in a vCard string. |
| ui.Interface.frontCall("mobile", "isEmulator", [], [result] ) | Indicates if the mobile app runs or displays forms on an emulator/simulator. |
| ui.Interface.frontCall("mobile", "isForeground", [], [result] ) | Indicates if the mobile app is in foreground mode. |
| ui.Interface.frontCall("mobile", "newContact", [defaults],[vcard]) | Opens contact input form to create a new entry in the contact database. |
| ui.Interface.frontCall("mobile","registerForRemoteNotifications", [], [registration-token] ) | This front call registers a mobile device for push notifications. |
| ui.Interface.frontCall("mobile", "runOnServer", [ appurl, timeout ], [] ) | Run an application from the Genero Application Server using the specified URL. |
| ui.Interface.frontCall("mobile", "scanBarCode", [options], [code, type] ) | Allow the user to scan a barcode with a mobile device |
| ui.Interface.frontCall("mobile", "takePhoto", [], [path] ) | Lets the user take a picture with the mobile device and returns the corresponding picture identifier. |
| ui.Interface.frontCall("mobile", "takeVideo", [], [path]) | Lets the user take a video with the mobile device and returns the corresponding video identifier. |
| ui.Interface.frontCall("mobile","unregisterFromRemoteNotifications", [], [] ) | This front call unregisters the mobile device from push notifications. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("android","askForPermission", [permission], [result]) | Ask the user to enable a dangerous feature on the Android device. |
| ui.Interface.frontCall("android", "showAbout", [],[]) | Shows the GMA about box displaying version information. |
| ui.Interface.frontCall("android", "showSettings", [], []) | Shows the GMA settings box controlling debug options. |
| ui.Interface.frontCall("android","startActivity", [action, data, category, type, component, extras], []) | Starts an external Android application (activity), and returns to the GMA application immediately. |
| ui.Interface.frontCall("android", "startActivityForResult", [action, data, category, type, component, extras], [outdata, outextras]) | Starts an external application (Android activity) and waits until the activity is closed. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("ios", "getBadgeNumber", [],[value]) | Returns the current badge number associated to the app. |
| ui.Interface.frontCall("ios", "setBadgeNumber", [value], []) | Sets the current badge number associated to the app. |

| Function Name | Description |
| --- | --- |
| ui.Interface.frontCall("cordova", "call", [plugin-name, function-name [, param1, param2, ... ] ], [result] ) | Calls a function in a Cordova plugin and returns a result. |
| ui.Interface.frontCall("cordova", "callWithoutWaiting", [plugin-name, function-name [, param1, param2, ... ] ], [callback-id]) | Calls a function asynchronously in a Cordova plugin, without waiting for a result. |
| ui.Interface.frontCall("cordova", "getAllCallbackData", [callback-id-filter], [results]) | Returns all results for asynchronous Cordova plugin front calls, based on a callback ID filter. |
| ui.Interface.frontCall("cordova", "getCallbackDataCount", [], [count]) | Returns the number of pending Cordova plugin results. |
| ui.Interface.frontCall("cordova", "getCallbackData", [], [result, callback-id]) | Returns the first Cordova plugin result from the result queue of all asynchronous Cordova plugin front calls, and removes it from the queue. |
| ui.Interface.frontCall("cordova", "getPluginInfo", [plugin-name], [result] ) | Returns details about a specific Cordova plugin. |
| ui.Interface.frontCall("cordova", "listPlugins", [ ], [plugins] ) | Returns the list of available Cordova plugins. |
