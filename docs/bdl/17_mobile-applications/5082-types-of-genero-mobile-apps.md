---
title: "Types of Genero Mobile apps"
source: "fgl-topics/c_gm_whatis_mobileapp.html"
breadcrumb: "Mobile applications > Types of Genero Mobile apps"
type: "concept"
---

# Types of Genero Mobile apps

> Genero supports different types of mobile app architectures: development mode, standalone apps, partially-connected apps, and client-server apps.

When you are developing your app, and you execute the app on your development machine for display
on a device or emulator, you are running the app in development mode.

When you follow the procedure to deploy your application to the device for testing or to
distribute your app to your end users, you have a deployed app. This type of app is
also referred to as an embedded app, as it has a DVM embedded in the app that runs on
the mobile device. A deployed app might be an app that executes irrespective of network
availability; it might be an app that accesses device peripherals; it might be an app that requires
network access.

## Development mode

In development mode, the app process is running on a server, and the forms are
displayed on a mobile front-end.

This architecture is typically used when developing your application, as you can easily modify,
recompile and execute your sources on the development machine, without the need to deploy the app on
the device.

![Diagram showing the Genero Mobile Client on the device and DVM running on the development machine.](../_images/mobile_development.png)

*Development mode*

## Standalone apps

A standalone app has the DVM and display client entirely on
the mobile device. This app executes irrespective of network availability and can access device
peripherals such as the camera, contacts, email, calendar, GPS, and storage via exposed APIs (front
calls). For database needs, this app can only connect to a local SQLite database.

![Diagram showing both Genero Mobile Client and DVM on mobile device.](../_images/mobile_standalone.png)

*Standalone app*

## Partially-connected apps

A partially-connected app has the DVM and display client
entirely on the mobile device, yet this app includes items that require a network connection. This
app must be able to run when no network connection is available. This app uses a network API to talk
to any back-end.

Examples include:

- Web Services performed with JSON over HTTP; use RESTful methods to write data synchronization
  routines. With this example, business logic executes within the device's Virtual Machine and the
  user is able to store captured data to a local SQLite database. When the network becomes available,
  the user synchronizes the stored data with the remote server's database.
- A web component that runs Google Maps.

This app first operates without a network connection, and must be able to run without a
network connection. Once network connectivity is restored, the app can perform network-dependent
tasks such as synchronizing with a remote database, make a web service call, or use a web
component.

If you are using GMI, and the device goes into standby mode, the application does not run in the
background and activities with the network are suspended.

![Diagram showing mobile device with both Genero Mobile Client and DVM on one side, web server with DVM and DBMS on the other side. Communication is bi-directional using JSON http over wireless.](../_images/mobile_datasync_json.png)

*Partially-connected app (example showing data synchronization via JSON)*

## Client-server apps

With a client-server app or connected app, the
bulk of the app runs on a remote server and the display client sits on the mobile device.

As with any deployed app, this app first starts on the mobile device; the DVM for the deployed
app runs on the mobile device. The role of the deployed app, however, is to connect to a remote
corporate server as an online terminal. It is the deployed app that launches the remote application
using the `runOnServer` frontcall. The remote application's DVM and business logic
reside on the remote server, somewhere in the network. The remote application is not limited to a
SQLite database.

In the event that the network is interrupted, the Genero Mobile client app is suspended until
service resumes.

![Diagram showing mobile device with Genero Mobile Client one side, web server with Genero Application Server and DVM on the other side. Communication is bi-directional using http over wireless.](../_images/mobile_connected.png)

*Client-server app*

## Licensing considerations

When deploying a Genero Mobile app, the type of mobile app architecture you need may have
additional requirements for Genero licenses as follows:

- Development mode needs a development license.
- Partially-connected app consumes a Genero runtime license when the app connects to the Web
  service to synchronize with the network database.
- Client-server app consumes a Genero runtime license on the Genero Application Server.

For specific mobile app license details, see your Software License Agreement, or contact your
nearest Four Js sales office.
