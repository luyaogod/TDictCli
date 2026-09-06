---
title: "Installing Cordova plugins"
source: "fgl-topics/c_fgl_cordova_github_plugins.html"
breadcrumb: "Mobile applications > Cordova plugins > Installing Cordova plugins"
type: "concept"
---

# Installing Cordova plugins

> Before usage, Cordova plugins need to be installed in the GMA or GMI development environment.

## Requirements for Cordova plugins installation

The following tools and utilities are required to install Cordova plugins:

1. Cordova plugins are available on a GitHub repository. Therefore, the git
   command line tool is required to install Cordova plugins.
2. The gmabuildtool uses the make utility to compile FGL
   wrappers and install Cordova plugins. Therefore, the make utility must be
   installed in the computer.

## Cordova plugin installation basics

GitHub contains a large number of Cordova plugins that can be used in your app, like those from
the [Apache GitHub](https://github.com/apache) for
example, or the [FOURJS Cordova GitHub](https://github.com/FourjsGenero-Cordova-Plugins).

Before installing the Cordova plugins in the GMA / GMI installation directories, you must first
clone the GitHub repository to your local disk.

## Upgrading GMA, GMI or FGL with GMI with installed Cordova plugins

Cordova plugins need to be re-installed in the following cases:

- With iOS/GMI: after upgrading the GMI package or the FGL package (the GMI package can be
  installed into the [FGLDIR installation directory](../04_installation/0051-install-genero-mobile-for-ios-single-version.md "To build and package Genero Mobile for iOS (GMI) applications, you must first install GMI. This topic explains how to install a unique GMI version/package into FGLDIR."), or in a
  separate [GMIDIR installation directory](../04_installation/0052-install-genero-mobile-for-ios-multiple-versions.md "To build and package Genero Mobile for iOS (GMI) applications, you must first install GMI. This topic explains how to install multiple GMI versions/packages in parallel into different GMIDIR directories."))
- With Android/GMA: after upgrading the GMA package (the GMA package is always to be installed in
  a dedicated [GMADIR installation directory](../04_installation/0050-install-genero-mobile-for-android.md "To build and package Genero Mobile for Android (GMA) applications, you must first install GMA."))

> **Important:**
>
> In all cases, after upgrading an FGL, GMI or GMA environment where plugins
> have been installed, you need to re-execute the plugin installation with a gmabuildtool
> scaffold --install-plugin or gmibuildtool --install-plugin
> command.

## Installing a Cordova plugin from a local clone of the GitHub repository

First clone the GitHub repository into a local
directory:

```
$ cd /usr/dev/myplugins
$ git clone https://github.com/apache/cordova-plugin-network-information.git
```

Then install the plugin:

- For GMI / iOS:

  ```
  $ gmibuildtool --install-plugins=/usr/dev/myplugins/cordova-plugin-network-information
  ```
- For GMA / Android (note that you
  need to use the "scaffold" command):

  ```
  $ gmabuildtool scaffold \
      --install-plugins /usr/dev/myplugins/cordova-plugin-network-information
  ```

  > **Note:**
  >
  > The gmabuildtool compiles cordova plugins FGL wrappers at installation with
  > the make command. Compiled .42m modules will be bundled
  > inside the scaffold archive. Before building an apk, FGL wrapper binaries will be copied inside the
  > application directory.

## Checking out a particular version from GitHub

If you want to test a particular version of a Cordova plugin, make a clone of the repository, and
check out that version, then install the plugin:

```
$ cd cordova-plugin-network-information
$ git checkout specific-version
$ cd -
```

It is now possible to use the gmibuildtool or gmabuildtool
to install this specific version.

> **Note:**
>
> If a plugin with the same name is already installed, it is replaced by the new
> installation.

## Checking Cordova plugin installation

Use the `--list-plugins` option of the GMA or GMI build tool, to check that the
plugin is properly installed:

- For GMI / iOS:

  ```
  $ gmibuildtool --list-plugins
  ...
  cordova-plugin-network-information
  ...
  ```
- For GMA / Android (note that you
  need to use the "scaffold" command):

  ```
  $ gmabuildtool scaffold --list-plugins
  ...
  cordova-plugin-network-information
  ...
  ```

The `cordova-plugin-network-information` name can now be used for bundling an
application using that plugin.

## Finding Cordova plugin wrapper libs in development environment

After installing the Cordova plugin, if the plugin repository includes a BDL wrapper library, the
fglrun runtime system needs to know where the .42m p-code
modules can be found.

If there is a local clone of the GitHub repository, search for \*.4gl files
(typically prefixed with "fglcdv"), make sure they are compiled to .42m
modules, and set the [FGLLDPATH](../07_configuration/0528-fglldpath.md "Defines a list of paths to find program modules.") to the
directory where the .42m files are located.

For example (on Unix/Linux), with the [`cordova-plugin-media`](https://github.com/FourjsGenero-Cordova-Plugins/cordova-plugin-media) repository
clone:

```
$ cd cordova-plugin-media
$ find . -name "fglcdv*.4gl"
./fgl/fglcdvMedia.4gl
$ cd fgl
$ fglcomp fglcdvMedia.4gl      (or make, if makefile is present)
$ export FGLLDPATH=$PWD        (or export FGLLDPATH=$PWD:$FGLLDPATH)
```
