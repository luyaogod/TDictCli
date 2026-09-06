---
title: "cordova.getPluginInfo"
source: "fgl-topics/c_fgl_frontcall_cordova_getPluginInfo.html"
breadcrumb: "Library reference > Built-in front calls > Cordova plugin front calls > cordova.getPluginInfo"
type: "concept"
---

# cordova.getPluginInfo

> Returns details about a specific Cordova plugin.

## Syntax

```
ui.Interface.frontCall("cordova", "getPluginInfo",
  [plugin-name], [result] )
```

1. plugin-name - This is the name of the Cordova plugin.
2. result - Holds details about the Cordova plugin, in JSON format.

## Usage

The `getPluginInfo` front call returns details about the Cordova plugin passed as
parameter.

The result variable is of type `STRING` filled with a JSON
string containing the details about the plugin such as:

1. The `id` of the plugin
2. The `version` of the plugin
3. The `git_version` of the plugin

## Example

```
DEFINE result STRING
CALL ui.interface.frontcall( "cordova", "getPluginInfo", [], [result] )
DISPLAY util.JSON.format(result)
```
