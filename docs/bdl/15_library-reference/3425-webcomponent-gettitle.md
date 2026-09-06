---
title: "webcomponent.getTitle"
source: "fgl-topics/c_fgl_frontcall_webcomponent_gettitle.html"
breadcrumb: "Library reference > Built-in front calls > Web component front calls > webcomponent.getTitle"
type: "concept"
---

# webcomponent.getTitle

> Returns the title of the HTML doc rendered by a web component.

## Syntax

```
ui.Interface.frontCall("webcomponent", "getTitle",
  [aui-name], [result] )
```

1. aui-name - This is the name of the web component in the AUI tree.
2. result - Holds the title of the HTML document.

## Usage

This front call can be used to get the title of the HTML document that is rendered by the
web component identified by the aui-name. For more details refer to <http://www.w3schools.com/tags/tag_title.asp>.

A typical usage of this front call is when implementing a web component based on the O-Auth
mechanism to identify the current user. For example, with the Google accounts authentication
service, after the login and password have been validated by Google, the authentication token is
returned in the title of the HTML document. This token is typically used by the application to
identify the user in remote API calls.
