---
title: "Genero Mobile for iOS (GMI) 1.30 changes"
source: "fgl-topics/c_fgl_Migrate_to_310_gmi_changes.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 3.10 upgrade guide > Genero Mobile for iOS (GMI) 1.30 changes"
type: "concept"
---

# Genero Mobile for iOS (GMI) 1.30 changes

> Modifications to consider when using Genero Mobile for iOS.

> **Note:**
>
> This topic describes features changes in the GMI 1.30 product. See also the Mobile section in
> [Genero BDL 3.10 New Features page](0063-bdl-3-10-new-features.md "Features added in 3.10 releases of the Genero Business Development Language.").

## GMI 1.30 with FGLGWS 3.10

> **Important:**
>
> The GMI version 1.30 is built on FGLGWS 3.10 and therefore, strongly tied to
> this Genero BDL version.

## Default directory for localized strings in GMI apps

The .42s localized string files for the default language can be provided in
the appdir/defaults directory.

Resource files such as .42s provided in
pwd are always loaded by the runtime
system. As GMA, pwd and
appdir are the same, it was not
possible to provide default strings files in
appdir. It was required to provide
localized string files for any language, in the corresponding
appdir/locale-code
directories.

For more details, see [Localized string files on mobile devices](../09_advanced-features/0909-loading-localized-strings-at-runtime.md) and [Deploying mobile apps](../17_mobile-applications/5102-deploying-mobile-apps.md "This section describes how to build and deploy mobile apps with Genero.").

## Building GMI apps with C extensions

Before GMI 1.30, in order to build an iOS app with C extensions, it was required to create your
own makefile based on the generic $GMIDIR/lib/Makefile-gmi
file, to add your own C extension libraries or custom front calls with the USEREXTENSION
variable:

```
...

GMI_OPTIONS = \
   APPNAME=MyApp \
   BUNDLE_IDENTIFIER=com.mycomany.myapp \
   IDENTITY="iPhone Developer" \
   PROVISIONING_PROFILE=~/Library/MobileDevice/Provisioning\ Profiles/myapp.mobileprovision \
   USEREXTENSION=userextension.o \
   TARGET=phone

GMI_MAKE = make -f $(GMIDIR)/lib/Makefile-gmi $(GMI_OPTIONS)

gmi.all: all
      $(GMI_MMAKE) all

...
```

Starting with GMI 1.30, it is now possible to build your static library with the
`staticlib` target of $GMIDIR/lib/Makefile-gmi, and pass your
library to the gmibuildtool with the `--extension-libs`
option:

```
$ gmibuildtool ... -extension-libs "-lz libBPush.a" ...
```

For more details, see [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.").

## Registering custom front calls in GMI

Before GMI 1.30, it was required to implement a `frontCalls()` function, in order
to declare the Objective-C class implementing your custom front
calls:

```
NSArray* frontCalls()
{
   return @[ [ExtensionFrontCall class] ];
}
```

Starting with GMI 1.30, upon startup, the GMI detects the extension by enumerating all
descendants of the Frontcall classes.

Thus, there is no need to implement this function anymore.

For more details, see [Implement front call modules for GMI](../14_extending-the-language/2721-implement-front-call-modules-for-gmi.md "Custom front call modules for the iOS front-end are implemented by using the API for GMI front calls in Objective-C.").

## Managing plugins in GMI

Starting with GMI 1.30, the gmibuildtool allows you to handle plugins with the
[`--install-plugins` and
`--list-plugins` options](../17_mobile-applications/5110-gmibuildtool.md "The gmibuildtool is a utility to create and test applications for an iOS devices.").

For more details, see [Cordova plugins](../17_mobile-applications/5117-cordova-plugins.md "This section describes how to use Cordova plugins.").
