---
title: "Executing programs"
source: "fgl-topics/c_fgl_programs_005.html"
breadcrumb: "Advanced features > Program execution > Executing programs"
type: "concept"
---

# Executing programs

> There are different ways to execute compiled programs, depending on the configuration and the development or production context.

## Prerequisites before executing a program

Make sure that all required environment variables are properly defined, such as: [FGLPROFILE](../07_configuration/0530-fglprofile.md "Defines the configuration files to be used by the runtime system.")
[FGLGUI](../07_configuration/0525-fglgui.md "Defines the user interface mode to be used by the program."), [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application."), [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules."), [LC\_ALL (or LANG)](../07_configuration/0497-lc-all-or-lang.md "Defines the current application locale on UNIX platforms.").

To display program forms in graphical mode, the GUI front-end must run on the computer
defined by [FGLSERVER](../07_configuration/0532-fglserver.md "Defines the graphical front-end for the application."), and all network security
components (such as firewalls) must allow TCP connections on the port defined by this environment
variable.

Verify the database client environment settings, and check that the
database server is running and can be accessed, for example by using a database vendor
specific tool to execute SQL commands.

## Starting a program from the command line on the server

A program can be executed with the fglrun tool from the server command
line:

```
fglrun progname arg1 arg2 ...
```

This method is typically used in development context. After compiling the programs and forms, for
example with the make utility, execute the programs with
fglrun.

The program name passed as parameter must be a .42m module implementing a
`MAIN` block, or a .42r program file. The file extension
(.42m or .42r) can be omitted: If no file extension is
specified, fglrun will try to load
progname.42r, then
progname.42m.

The progname argument can be a simple base name of the program file, or an
absolute or relative filename (including the path) to the program file. The FGLLDPATH environment
variable is not used to find the program file: This variable is used to find modules used by the
program.

For more details about the fglrun options and how the program file is found,
see [fglrun command reference](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.").

## Executing sub-programs from a parent program with RUN

Sub-programs can be executed from a parent program with the `RUN` instruction.

There can be limitations, depending on the platform where the parent program executes.

The `RUN` instruction takes a shell command as parameter. To start another BDL
program from the current program, specify [fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs."), followed by the program name and arguments.

The `RUN` instruction provides `IN LINE/FORM MODE` and
`WITHOUT WAITING` options, to be used depending on the purpose of the child program
or command, and if the parent program runs in TUI or GUI mode.

In GUI mode, the windows/forms of each application can be displayed in a common window container
(OS desktop window / web browser tab), or in a dedicated window container.

For more details, see the reference page of the [`RUN`](0830-run.md "The RUN instruction executes the command passed as argument.") instruction.

## Starting a program from the front-end

It is also possible to start programs on the application server from the platform where the
front-end resides.

This is actually the typical way to start applications in a production environment.

- For a desktop front-end (GDC) application, define application shortcuts and use
  rlogin/ssh network protocols to start programs on the server or by using HTTP
  through a web server (GAS).
- For a web-browser application (GAS), configure the application server to run
  applications from a URL.
- For a mobile device application (GMI/GMA), in a configuration where the programs
  run on a GAS application server, use the "`runOnServer`" front
  call, to start a program from the GAS.

## Starting programs on a mobile device

After deploying program files on a mobile device, it can be executed as a local
application, typically with a tap on the application icon.

- For a GMA (Android™) application, program
  files and GMA must be bundled together in an .apk
  Android
  package to be deployed. For more details, see [Deploying mobile apps on Android devices](../17_mobile-applications/5103-deploying-mobile-apps-on-android-devices.md "This section contains information to create a mobile application to be deployed on Android devices.").
- For a GMI (iOS) application, program files and GMI must be bundled together in an
  .ipa package to be deployed. For more details, see
  [Deploying mobile apps on iOS devices](../17_mobile-applications/5107-deploying-mobile-apps-on-ios-devices.md "This section contains information to create a mobile application to be deployed on iOS devices.").
- To start programs on an application server from a small embedded mobile
  application (starter), use the `runOnServer` front call. For more
  details, see [Running mobile apps on an application server](../17_mobile-applications/5111-running-mobile-apps-on-an-application-server.md "From the mobile device, programs can be started remotely on an application server, and displayed on the device.").

## Foreground and background modes

On mobile devices, apps can switch between the foreground and background mode. The app state can
be detected and managed in your Genero program code.

With GBC in a browser (using GAS), it is also possible to detect when the browser window or
current tab is hidden or shown.

Genero BDL provides two predefined actions, to detect when the state of a mobile app changes, or
when the browser window/tab gets hidden or shown to the end user:

1. The action `enterbackground` is fired, when the mobile app goes to background
   mode, or when the browser window or current tab gets hidden to the user, because a browser
   window/tab switch occurred, or the browser window got minimized. With GDC/UR, the
   `enterbackground` action is fired when the window container is minimized.
2. The action `enterforeground` is fired, when the mobile app goes to foreground
   mode, or when the browser window or current tab is shown to the user, because a browser window/tab
   switch occurred, or the browser window is restored. With GDC/UR, the
   `enterforeground` action is fired when the window container is restored.

The `enterforeground` action is not fired when the app starts. This action is only
fired when returning to foreground mode, after it was in background mode.

To execute code when the app goes to background or foreground mode, use on `ON
ACTION` handler:

```
   ON ACTION enterbackground
      LET skip_timers = TRUE
   ON ACTION enterforeground
      IF NOT ask_password() THEN
         EXIT PROGRAM
      ENF IF
```

For example, when the mobile app enters background mode, it is recommended that the program
suspend any activity, and skip code that would be executed in [`ON TIMER`](../11_user-interface/2227-get-program-control-on-a-regular-timed-basis.md "Execute some code after a given number of seconds, with or without user interaction with the program.") triggers.

On the other hand, when the mobile app enters foreground mode, the program can for example ask
the user's login/password again, for security reasons, to make sure that the mobile device did not
end up in other hands while the app was in background mode.

To control action view rendering defaults and current field validation behavior when the
`enterforeground` / `enterbackground` actions are used, consider
setting action default attributes for these action in your [.4ad file](../11_user-interface/2259-configuring-actions.md "Action attributes related to decoration, keyboard shortcuts, and behavior can be defined with action attributes.") as follows (like done in
FGLDIR/lib/default.4ad):

```
<ActionDefaultList>
  ...
  <ActionDefault name="enterbackground" validate="no" defaultView="no" contextMenu="no"/>
  <ActionDefault name="enterforeground" validate="no" defaultView="no" contextMenu="no"/>
  ...
</ActionDefaultList>
```

Another option is to define these action defaults attributes in the `ON ACTION`
handlers:

```
ON ACTION enterbackground ATTRIBUTES(VALIDATE=NO, DEFAULTVIEW=NO)
   ...
ON ACTION enterforeground ATTRIBUTES(VALIDATE=NO, DEFAULTVIEW=NO)
   ...
```

See also [List of predefined actions](../11_user-interface/2257-list-of-predefined-actions.md) and [Background/foreground modes](../17_mobile-applications/5094-background-foreground-modes.md "Describes how to handle background or foreground modes in mobile apps.").

## Common app directories on mobile platforms

On mobile devices, you can use the following APIs to get common directories:

1. `base.Application.getProgramDir` returns
   the directory path where the main .42m is located. Consider this location
   read-only and safe (no other app can access it). See also the [FGLAPPDIR](../07_configuration/0519-fglappdir.md "Contains the path to the application directory when executing on a mobile device.") environment variable.
2. `os.Path.pwd` returns the path to
   the current working directory. When a mobile application is started, the GMA and the GMI set the
   working directory to the default application directory. Consider this location read-write and safe
   (no other app can access it).
3. The front call [`standard.feInfo/dataDirectory`](../15_library-reference/3399-standard-feinfo.md "Queries general front-end properties.") returns the front-end
   side temporary directory. Storage on this directory may be erased by the OS. On
   an embedded mobile application, as the runtime and the front-end run on the same
   system, the program can use this front call to retrieve a temporary directory
   and use the path to store temporary files. Consider this location read-write and
   unsafe. Applications executed remotely through a `runOnServer`
   front call, can use the sandboxRunOnServer directory under
   the directory returned by the `feInfo/dataDirectory` front call,
   to exchange files with the embedded application.

## Related links

**Related concepts**  

[Application locale](0864-application-locale.md "The application locale defines the language and codeset for your application.")

[Genero environment variables](../07_configuration/0506-genero-environment-variables.md "Genero environment variables")

[Operating system environment variables](../07_configuration/0496-operating-system-environment-variables.md "Describes some well-known system environment variables that are used by Genero software components.")
