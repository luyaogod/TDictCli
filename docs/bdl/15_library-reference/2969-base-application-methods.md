---
title: "base.Application methods"
source: "fgl-topics/c_fgl_ClassApplication_methods.html"
breadcrumb: "Library reference > Built-in packages > The base package > The Application class > base.Application methods"
type: "concept"
description: "Table 1. Class methods Name Description base.Application.getArgument ( index INTEGER ) RETURNS STRING Returns the command line argument by position. base.Application.getArgumentCount () RETURNS ..."
---

# base.Application methods

| Name | Description |
| --- | --- |
| base.Application.getArgument( index INTEGER ) RETURNS STRING | Returns the command line argument by position. |
| base.Application.getArgumentCount() RETURNS INTEGER | Returns the total number of command line arguments. |
| base.Application.getProgramDir() RETURNS STRING | Returns the directory path of the current program. |
| base.Application.getProgramName() RETURNS STRING | Returns the name of the current program. |
| base.Application.getFglDir() RETURNS STRING | Returns the path to the FGLDIR installation directory. |
| base.Application.getResourceEntry( name STRING ) RETURNS STRING | Returns the value of a FGLPROFILE entry. |
| base.Application.getStackTrace() RETURNS STRING | Returns the function call stack trace. |
| base.Application.isGWA() RETURNS BOOLEAN | Indicates if the application runs in a browser. |
| base.Application.isMobile() RETURNS BOOLEAN | Indicates if the application runs on a mobile device. |
| base.Application.reloadResources( newResourcePath STRING) | Resets FGLRESOURCEPATH and reloads localized string resources. |
