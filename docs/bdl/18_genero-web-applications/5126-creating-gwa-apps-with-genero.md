---
title: "Creating GWA apps with Genero"
source: "fgl-topics/c_fgl_gwa_build_app.html"
breadcrumb: "Genero Web applications > Creating GWA apps with Genero"
type: "concept"
---

# Creating GWA apps with Genero

> Prepare the environment to build GWA applications.

Before creating the GWA application:

1. The Genero BDL development environment (FGLDIR) must be installed to compile your program
   files.
2. The GWA must be installed. For more details, go to [Install Genero Web Application](../04_installation/0053-install-genero-web-application.md "To build and package Genero Web applications, you must first install Genero Web Application (GWA).").
3. Refer to the topics in the [Overview](5124-overview.md "Genero Web Application (GWA) is a browser-based solution that combines Genero Browser Client, p-code modules, and a WebAssembly-based runtime for executing applications offline as progressive web apps. It supports full Genero GUI features, requires HTTPS for secure deployment, and is ideal for scalable use cases, including offline functionality and large user bases.") section to become
   familiar with features of the GWA.

Follow the sequence of topics below as needed to create and customize your GWA app.

## Child topics

- [Create a program directory](5127-create-a-program-directory.md): To build a GWA application, you must create a directory.
- [Set environment with fglprofile](5128-set-environment-with-fglprofile.md): Configure environment settings with FGLPROFILE entries.
- [Specify the GBC](5129-specify-the-gbc.md): A GBC must be bundled with the GWA.
- [Package web components](5130-package-web-components.md): Web components must be packaged with the GWA application because they are preloaded and must also be available offline.
- [Specify the start-up module](5131-specify-the-start-up-module.md): Specify the module you want to use as the main entry point.
- [Customizing GWA apps](5132-customizing-gwa-apps.md): This section describes ways you may customize your GWA apps, such as adding a favicon, customizing the index page, and internationalizing your app.
