---
title: "Custom GMA binary archive build"
source: "fgl-topics/t_fgl_JavaBridge_gma_packaging_steps.html"
breadcrumb: "Extending the language > The Java interface > Advanced programming > Executing Java code with GMA > Packaging custom Java extensions for GMA > Custom GMA binary archive build"
type: "task"
description: "If you are planning to build Genero Mobile for Android extensions for your GMA project, you need to do this using Android Studio. Follow this procedure to extend GMA."
---

# Custom GMA binary archive build

> If you are planning to build Genero Mobile for Android™ extensions for your GMA project, you need to do this using Android Studio. Follow this procedure to extend GMA.

When using GMA user extensions, you need to build the Java binaries for the final custom GMA
.apk or .aab, as well as the
user-extension.jar library, to compile Genero sources using
`IMPORT JAVA user-extension`. For a complete overview, see [Packaging custom Java extensions for GMA](2697-packaging-custom-java-extensions-for-gma.md "Custom Java extension must be integrated in the GMA to run on Android devices.").

> **Important:**
>
> Android Studio
> must be installed, and minimum Android
> development skills are required to perform this task.

1. Locate the original GMA binary archive on your computer. If using Genero Studio, it is found in the GMADIR/gma/artifacts directory. The GMADIR
   variable is set in the Genero Studio configuration settings. If you are not
   using Genero Studio, the GMA binary archive is provided as a separate
   package.

   The GMA binary archive consists of a file named
   fjs-gma-version-build-android-scaffolding.zip.
   Version 1.30.00 or later is required.

   This file contains the original GMA core binary and the sub-project to build custom
   extensions.
2. Create a project directory to work in, for example extension-example, to
   where you extract the GMA scaffolding archive.
3. Unzip the
   fjs-gma-version-build-android-scaffolding.zip
   archive into extension-example/project/scaffold.
4. In Android Studio, open the
   project from extension-example/project/scaffold.

   In the Android Studio
   Project view, under the extension library, you can find
   the TestExtensionActivity.java source file: This is a sample activity to test
   the extension bundled in the scaffold project. This class is provided for testing purpose and
   sample. Rename it and change the code for your own purpose, or replace it with your own extension
   source.
5. Define the extension module package name in the Android manifest file.

   In the Android Studio
   Project view, under the extension directory, open the
   templates/src/extension/build.gradle.template file. Modify the value of
   `namespace`:

   ```
   ...
          release {
               minifyEnabled false
               proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
           }
       }   
       namespace 'com.gma.extension'
   ...
   ```

   The package name must be a Java-language-style package name, for example, in the format of
   `"com.gma.extension"`. The package name identifies your extension
   library, it is not the name used to build a final app in Genero Studio.
6. Add your Java sources to this Android Studio project, under the extension library.
7. If required, add external .jar libraries to the project: Copy the
   additional jar libraries into
   extension-example/project/scaffold/extension/libs. These jar libraries will be
   included when building the project.
8. In Android Studio, build the
   project in release or debug mode.

   This creates the custom extensions binaries to build the final .apk or
   .aab and the extension.jar file in
   extension-example/project/scaffold/extension/build/libs. The
   extension.jar file contains Java classes that are callable from Genero BDL. The
   extension.jar file is required to compile Genero source code using
   `IMPORT JAVA extension`.

   The project build in Android Studio
   will also create the customized GMA .apk or
   .aab. However, this step is done explicitly later, either from
   Genero Studio, or by using the gmabuildtool command.
9. Build your custom GMA .apk or .aab with your custom
   extensions.

   If using Genero Studio, perform the following steps:

   1. Modify the CLASSPATH environment variable to include the extension.jar
      file.

      This is required to let the Genero compiler find your Java classes.
   2. Modify the GMADIR configuration variable to point to the
      extension-example/project/scaffold directory. 

      This is required to let Genero Studio use your customized GMA binary to build apps. For more
      information, see the Configuration for extending Genero Mobile for Android topic in the
      Genero Studio User Guide.

   If using gmabuildtool to build apps from the command line:

   1. Change directory to extension-example and call the
      gmabuildtool to build and bundle your final .apk or
      .aab file with the required options provided by the tool:
      gmabuiltool will use by default find GMA binary files and resources in the
      current working directory. For more details about gmabuildtool usage, see [Building Android apps with Genero](../17_mobile-applications/5105-building-android-apps-with-genero.md "Genero provides a command-line tool to create applications for Android devices.").
10. Deploy the new GMA Android
    application on the device.

    - When using Genero Studio, the apk packages to be installed on the
      device are referenced from the GMADIR environment variable.
    - If not using Genero Studio, you can use the Android debug bridge (adb)
      command-line utility included with Android SDK to install the .apk
      package to your device.
    - The .aab package can be installed with the bundletool
      utility. See [Android documentation](https://developer.android.com/studio/command-line/bundletool) for more details.
