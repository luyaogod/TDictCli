---
title: "Debugging a web component"
source: "fgl-topics/c_fgl_webcomponent_debug.html"
breadcrumb: "User interface > User interface programming > Web components > Debugging a web component"
type: "concept"
description: "Debugging a Web Component in a web browser In order to debug a Web Component when the application forms are displayed in a web browser, start the debug mode of your browser (typically, F12 key), and ..."
---

# Debugging a web component

## Debugging a Web Component in a web browser

In order to debug a Web Component when the application forms are displayed in a web browser,
start the debug mode of your browser (typically, F12 key), and use the integrated HTML / JavaScript
debugger to inspect the Web Component contents.

## Debugging a Web Component with the GDC front-end

The WebEngine debugger is a debugging tool that opens in a separate window, providing access to
the QT WebEngine Developer Tools. These tools include a variety of tools. One such tool is the
inspector, which shows the HTML and JavaScript code when you click on an element. To learn more
about tools for Web developers, see [Chrome
DevTools](http://developers.google.com/web/tools/chrome-devtools).

To open the Chrome Developer Tools window, press `[Alt]+[Shift]` and right-click
the mouse pointer over the web component field. This shortcut will work in native mode and universal
rendering mode.

## Debugging a Web Component with the GMA front-end

In order to debug a web component displayed on GMA, you need a Chrome web browser, and an USB
cable to plug the device to your computer.

Steps to debug a web component on GMA:

1. Connect your device to your desktop with the USB cable.
2. Stop GMA on the device, if it is already running.
3. Enable debugging option in GMA parameters.
4. Start the application with a web component. This will start GMA with debug service enabled
   (please wait as the application could take longer to start)
5. Open Chrome on the desktop computer.
6. In the Chrome URL address bar, enter: `chrome://inspect/#devices`
7. On the device, accept the USB debugging.
8. Click on "inspect" link to open the HTML debug window.

## Debugging a Web Component with the GMI front-end

In order to debug a web component displayed on GMI, you need to setup the Safari® Web Inspector. Check the Apple® development site for more details about enabling the Safari Web Inspector: <http://developer.apple.com>.

Steps to debug a web component on GMI:

1. On you Mac, open Safari and make sure
   the "Develop" option is available. If this option is not available, go to Safari preferences and enable the option.
2. On the device, go to Settings. In Safari preferences, go to Advanced and enable the Web Inspector option.
3. Run the app on the device or simulator, or start the program in development mode to display on
   the GMI front-end.
4. On the Mac, find the connected device/simulator in the Develop menu of Safari and browse the debuggable components.
