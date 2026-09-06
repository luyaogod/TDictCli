---
title: "Overview"
source: "fgl-topics/c_fgl_section_gwa_overview.html"
breadcrumb: "Genero Web applications > Overview"
type: "concept"
---

# Overview

> Genero Web Application (GWA) is a browser-based solution that combines Genero Browser Client, p-code modules, and a WebAssembly-based runtime for executing applications offline as progressive web apps. It supports full Genero GUI features, requires HTTPS for secure deployment, and is ideal for scalable use cases, including offline functionality and large user bases.

The GWA is an end-to-end web application solution that bundles together a Genero Browser Client,
Genero p-code modules (42m), a Genero VM, and resources into a single page
application (SPA).

The GWA ships a [WebAssembly](https://webassembly.org/) (external link) (`wasm`) version of the VM for your application —
a format that enables deployment on the web and execution of the application in a browser. Like on
desktop, this VM interprets .42m p-code modules.

The VM is created with the help of the [emscripten](https://emscripten.org) (external link) compiler that is bundled with each GWA
application. The processing speed is almost as fast on the browser as native code on the
desktop.

With a GWA application, you do not need a runtime system (fglrun) process on
the server side. In fact, there is no other software server side except a web server. A GWA
application is by default a "progressive web app" (PWA); meaning once loaded from a web
server, it can continue to run offline without any connection to the internet.

All GUI features available in the Genero Browser Client are fully supported in a GWA application,
so Genero forms and styles run as usual together with Genero interaction statements such as
`INPUT/INPUT ARRAY/CONSTRUCT/MENU/DIALOG`, and so on.

A GWA application has certain similarities with Genero Mobile for Android™(GMA) and Genero Mobile for iOS (GMI) application;
however, the main difference is that the Genero runtime is not compiled into native code, instead
the runtime system is compiled to [WebAssembly](https://webassembly.org/) (external link) modules (wasm) for
execution on the browser.

Another noticeable difference between GMA/GMI and GWA, is that a GWA application emulates a [file system in the browser](5125-file-system.md "When the GWA application is launched in the browser, a virtual UNIX-like file system emulation of your GWA application is created in memory."), whereas in a GMA/GMI application
the file systems of Android/IOS® can be
used.

## Prerequisites

GWA has requirements for delivering applications in a [secure context](https://www.w3.org/TR/secure-contexts/)
(external link) as defined by the W3C consortium:

- The GWA uses Service Workers that require you to run applications with HTTPS activated, or to
  use localhost. Even if in development localhost (http://localhost or
  http://127.0.0.1) can be used, it is mainly necessary to develop with HTTPS
  activated. Of course, in production, HTTPS is mandatory.
- It is recommended to only use X.509 certificates issued by a recognized certificate authority,
  such as [Let's Encrypt](https://letsencrypt.org/),
  in both development and in production.
  > **Note:**
  >
  > **Self-signed certificates**
  >
  > It is possible to use
  > self-signed certificates for development as long as your browser allows the usage of those
  > certificates.

## GWA use cases

The following are examples of some common GWA use cases:

- **Run application from a URL**: Build a GWA application like you would a mobile application
  for GMI/GMA except with the advantage that you do not have to managed it through App Store; the
  installation is an ordinary web URL. Once the URL is selected by the end user, the application is
  downloaded, starts immediately, and is also immediately available for offline use.
- **Launch from mobile home screen icon**: For mobiles with Chrome/Edge browsers, an
  installation menu appears, prompting the end user to save this application under a mobile home
  screen icon. Starting the application from this icon does not require an internet connection. With
  other browsers, users can bookmark the application with the same offline functionality. In IOS
  Safari® for example, add to home
  screen.
- **Launch from desktop icon**: For desktops with Chrome/Edge browsers, an install button
  appears in the address bar, which allows the user to save the application to a desktop icon, turning
  the application into a desktop application with offline functionality.
- **Facilitate large numbers of simultaneous end users**: GWA applications are the medium of
  choice when it comes to giving access to large number of simultaneous end users. Standard Genero
  desktop applications are server based and scale well up to 1000 and even more end users, but if the
  potential simultaneous end user count is much higher (for example, 100,000 or more) then it is time
  to switch to GWA architecture – with web services server side and the program logic running entirely
  client side.
- **Create GWA applications written in Genero BDL**: This is typically the case for a GWA
  application: code written entirely in Genero runs client side, communication with the server is done
  via Genero `com` API/REST services. GWA applications can therefore compete with other
  client-side JavaScript technologies; with the added advantage that you can code your applications in
  Genero BDL.

## Recommendations for use of Genero Web application

GWA has these best practice recommendations:

- **Install the GWA and the REST services on the same host**: As GWA is actually a web
  application running in a browser, you can only communicate via REST web services to exchange data
  with a third party server. Therefore, we recommend you install the GWA and the REST services on the
  same host and port; otherwise, you may face a lot of Cross-Origin Resource Sharing (CORS) issues.
- **Set CORS headers on external servers providing web services**: If you need to access
  another REST service located on another host and port machine, and if the REST requests are basic –
  simple `GET` or `x-WWW-form-urlencoded` type requests – you will have
  to set the following CORS header in the second host machine:
  `Access-Control-Allow-Origin:"https://host:port"`

## Related links

**Related concepts**  

[Creating GWA apps with Genero](5126-creating-gwa-apps-with-genero.md "Prepare the environment to build GWA applications.")

**Related tasks**  

[Install Genero Web Application](../04_installation/0053-install-genero-web-application.md "To build and package Genero Web applications, you must first install Genero Web Application (GWA).")
