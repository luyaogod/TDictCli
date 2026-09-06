---
title: "com.WebOperation methods"
source: "fgl-topics/c_gws_ComWebOperation_methods.html"
breadcrumb: "Library reference > Extension packages > The com package > Web services classes > The WebOperation class > WebOperation methods"
type: "concept"
---

# com.WebOperation methods

> Methods for the com.WebOperation class.

| Name | Description |
| --- | --- |
| com.WebOperation.CreateDOCStyle( function STRING, operation STRING, inputVar any-type, outputVar any-type ) RETURNS com.WebOperation | Creates a new Web Operation object with Document style. |
| com.WebOperation.CreateOneWayDOCStyle( function STRING, operation STRING, inputVar any-type ) RETURNS com.WebOperation | Creates a new Web Operation object with One-Way Document style. |
| com.WebOperation.CreateOneWayRPCStyle( function STRING, operation STRING, inputVar any-type ) RETURNS com.WebOperation | Creates a new Web Operation object with One-Way RPC style. |
| com.WebOperation.CreateRPCStyle( function STRING, operation STRING, inputVar any-type, outputVar any-type ) RETURNS com.WebOperation | Creates a new Web Operation object with RPC style. |

| Name | Description |
| --- | --- |
| addFault( fault any-type, wsaaction STRING ) | Adds a fault to the current Web Operation definition. |
| addInputHeader( header any-type ) | Adds an input header for the current Web Operation definition. |
| addOutputHeader( header any-type ) | Adds an output header for the current Web Operation definition. |
| initiateSession( ok INTEGER ) | Defines the Web Operation as session initiator. |
| setComment( comment STRING ) | Sets the comment for the Web Operation object. |
| setInputAction( action STRING ) | Sets the WS-Addressing action identifier of the input operation. |
| setInputEncoded( val INTEGER ) | Defines the encoding mechanism for Web Operation input parameters. |
| setOutputAction( action STRING ) | Sets the WS-Addressing action identifier of the output operation. |
| setOutputEncoded( val INTEGER ) | Defines the encoding mechanism for Web Operation output parameters. |
