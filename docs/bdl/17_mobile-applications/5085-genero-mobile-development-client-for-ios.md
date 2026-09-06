---
title: "Genero Mobile Development Client for iOS"
source: "fgl-topics/c_fgl_mobile_devel_mode_gmi.html"
breadcrumb: "Mobile applications > Mobile development mode > Genero Mobile Development Client for iOS"
type: "concept"
---

# Genero Mobile Development Client for iOS

> Set up a development environment to display app forms on an iOS device.

The Genero Mobile Development Client provides a development environment to display apps executing on a computer to an
iOS device or emulator. A ready-to-use Genero Mobile Development Client can be downloaded from the App Store, or you can
create your own Genero Mobile Development Client.

## Types of GMI front-end apps for iOS

To display Genero application forms on a device during developoment, the GMI front-end must be installed on the iOS
device or emulator.

Genero supports two types of GMI front-end Genero apps for iOS:

- The ready-to-use "Genero Mobile Development Client" for iOS, available through the App Store.
  > **Important:**
  >
  > Due to Apple® limitations, the Genero Mobile Development Client app is not allowed to
  > listen to a TCP port to provide a GUI service. In order to establish the GUI front-end connection,
  > the front-end must connect to the runtime system running on the development machine.
- A self-made GMI front-end, using the gmibuildtool to build a GMI front-end app with your own Apple developer account. With this configuration, you perform a
  direct GUI connection based on FGLSERVER.

## Install and use the Genero Mobile Development Client for iOS

Go to the App Store. Search for "Genero dev". Select and install the "Genero Development Client".

Once the Genero Mobile Development Client is installed on the device, make sure that WIFI is enabled and start the
GMI app.

Because of iOS app limitations defined by Apple, an app shipped on the App Store cannot listen to a TCP port to
provide a GUI service. In order to display Genero forms on the Genero Mobile Development Client, you will have to
establish the GUI connection from the device to the server, after starting the fglrun process with
the `--gui-listen` option. The home page of the Genero Mobile Development Client includes a URL field,
where you enter the IP address/hostname and the TCP port of the runtime system listening for incoming connections.

On the development server, start the application using the `fglrun
--gui-listen=portnum`
command:

```
fglrun --gui-listen=6500 main.42m
```

Make sure that the firewall on the development machine allows incoming connections for the TCP
port number specified with the `--gui-listen` fglrun option.

On the iOS device, enter the following in the URL
field:

```
fgl://dev-server-hostname:6500
```

Then tap the Connect button to establish the GUI connection.

## Build your own GMI front-end

It is possible to create your own Genero Mobile Development Client, with your own Apple certificate and provisioning profile. To display Genero forms on a self-made Genero Mobile
Development Client, you must use the direct GUI connection mode with fglrun connecting to the mobile
front-end via [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.").

A self-made GMI front-end can only be created on a macOS® computer.

The self-made GMI front-end can be deployed on your device or emulator. The GMI front-end will listen on the port
6400, displaying applications running on a server through the FGLSERVER setting.

Before creating your own GMI front-end, fulfill the prerequisites to build an iOS app as
described in [Building iOS apps with Genero](5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.").

To build your own GMI front-end:

1. Your macOS computer must be setup with
   Xcode® and Genero BDL environment as when
   creating of Genero Mobile for iOS apps. For more details, see [Install Genero Mobile for iOS (single version)](../04_installation/0051-install-genero-mobile-for-ios-single-version.md "To build and package Genero Mobile for iOS (GMI) applications, you must first install GMI. This topic explains how to install a unique GMI version/package into FGLDIR.").
2. Make sure that the installed fjs-gmi\*.zip archive is unzipped into
   FGLDIR.
3. Go to the $FGLDIR/demo/MobileDemo/gmiclient directory.
4. Delete the complete build directory, if it exists. (This can be done with a
   make clean command.)
5. Compile the GMI front-end program files (main.42m,
   main.42f, etc):

   ```
   $ make
   ```

   See
   $FGLDIR/demo/MobileDemo/gmiclient/Makefile for details.
6. Build the GMI front-end:
   - To build and install the GMI front-end on the simulator, first make sure that the simulator is
     started (open -a Simulator command), then execute the make
     command with the `gmi.install`
     rule:

     ```
     $ open -a Simulator
     $ make gmi.install
     ```
   - To build and install the GMI front-end on the device plugged to your Mac, get a development
     certificate (for the `IDENTITY` makefile variable) and provisioning profile (for the
     `PROVISIONING_PROFILE` makefile variable), then execute the make
     command with the `gmi.install` rule, by specifying the `TARGET`, the
     IDENTITY, and the PROVISIONING\_PROFILE
     variables:

     ```
     $ make TARGET=phone IDENTITY=WKRRJZ999 \
         PROVISIONING_PROFILE="~/Library/MobileDevice/Provisioning Profiles/myapp.mobileprovision" \
         gmi.install
     ```

Once the GMI front-end is installed on the device, make sure that WIFI is enabled and start the
GMI app.

The home page of the GMI shows the IP address of the device and the TCP port it is listening to (0=6400)."

On the development machine, set the [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application.") environment
variable with the IP address of the device.

You are now ready to run your app on the server and display on the iOS device.

## The Genero Mobile Development Client user interface

This section describes the interface of the Genero Mobile Development Client for iOS available through the App Store.
Once installed, you can access the front-end by tapping on the Genero icon in the home screen of the device. If using
iOS 14 or later, you may find the icon in the App Library, which appears as the last page in the
home screen. Once launched, you are taken to the home page, as shown in Figure 1.

![Screen of the Genero Mobile Development Client for iOS](../_images/gmi_dev_client_home.png)

*Genero Mobile Development Client screen*

Settings on the home page allow you to perform the functions listed.

- Enter URL: Enter the IP address/hostname and the TCP port for the server
  application. The URL can be the URL of an application running on a Genero Application Server (GAS)
  or an application started with `fglrun
  --gui-listen`:

  ```
  http://dev-server-hostname:tcp-port/gas/ua/r/myApp
  ```

  Alternatively,
  it can be a URL connecting directly to the FGL runtime with a URL using the
  fgl://
  format:

  ```
  fgl://dev-server-hostname:tcp-port
  ```

  For
  more information on starting an application on the server, see Install and use the Genero Mobile Development Client for iOS.
- Options/Timeout: This setting allows you to set a timeout (in seconds) of
  when the device stops trying to connect to the runtime. The default is 5 seconds. After the timeout,
  the button Connect to Dev URL becomes active again. The timeout can not be
  interrupted. When using a long timeout, you need to wait until the timeout is reached if the
  connection cannot be made.
- Options/Permanent retry: This setting means the GMI front-end will
  indefinitely try to reconnect to an application that is stopped. If you stop the application on the
  server, and then restart it, the device automatically reconnects. While the application is
  disconnected, a dialog displays a message with the number of seconds to the next reconnect attempt.
  You can cancel the permanent retry by tapping the cancel button on the
  dialog. Clear the option, if you wish to disable the permanent retry feature.
- HTTP viewing URL (AUI tree): On your server you can enter the URL
  displayed in a browser page to view the application AUI tree. For example, enter the device IP
  address and the TCP port displayed in the field. If using the emulator, you can enter
  localhost and TCP port.
- Results: This field displays the GBC version embedded in the installation
  used to render the application forms on the device. All connecting apps use the GBC on the server
  side.

## The self-made Genero Mobile Development Client user interface

This section describes the interface of the self-made Genero Mobile Development Client for
iOS—also known as the GMI client. For more information on creating the application, see Build your own GMI front-end. Once installed, you can
access the front-end by tapping on the GC (GMI client) icon in the home screen of the device. If
using iOS 14 or later, you may find the icon in the App Library, which appears
as the last page in the home screen. Once launched, you are taken to the home page as shown in Figure 2.

![Screen of the self-made Genero Mobile Development Client for iOS](../_images/gmi_selfmade_dev_client_ui.png)

*Genero Mobile Development Client screen*

The home page display details for the remote GUI connection and the HTTP URL for viewing the AUI
tree of the current application:

- On your server you can set FGLSERVER with the IP address/hostname and the TCP port to start an
  application with `fglrun --gui-listen` and display its forms to the device.
- On your server you can enter the URL displayed in a browser page to view the application AUI
  tree. For example, enter the device IP address and the TCP port displayed. If using the emulator,
  you can enter localhost and TCP port.

Details of the GBC version embedded in the installation used to render the application forms on
the device is displayed.

All connecting apps use the GBC on the server side.

## Related links

**Related concepts**  

[Deploying mobile apps on iOS devices](5107-deploying-mobile-apps-on-ios-devices.md "This section contains information to create a mobile application to be deployed on iOS devices.")
