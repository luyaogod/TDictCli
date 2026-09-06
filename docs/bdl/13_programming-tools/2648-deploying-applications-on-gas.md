---
title: "Deploying applications on GAS"
source: "fgl-topics/t_fgl_deploying_applications_on_gas.html"
breadcrumb: "Programming tools > Packaging web applications > Deploying applications on GAS"
type: "task"
---

# Deploying applications on GAS

> Follow these procedures to build archives for deploying applications and services with fglgar.

As a prerequisite, you must have your applications and/or services created, compiled, and
tested. Consolidate all the necessary files for your archive under a root directory.

Included in this page are some typical examples of how to build archives and deploy
applications:

- Overview of the main tasks
- Build a simple archive file
- Build an archive file with many applications
- Build an archive with public resources
- Build an archive with deployment triggers
- Deploy your application on your machine
- Enable your application on your machine
- Run the deployed application

For a full understanding of what Genero archiving offers, please read all archiving topics in
the Genero Application Server User Guide.

## Overview of the main tasks

This procedure provides you with a quick overview of the main steps for archiving and
deploying an application using fglgar. For more information, see [fglgar](2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).").

1. Create the archive file using the fglgar gar command.

   If you do not have application configuration files created, they are created automatically for
   you. See [Understanding packaging with GAS](2647-understanding-packaging-with-gas.md "Before deploying applications on a GAS installation, you need to package the required compiled files and resources in a Genero archive file. The fglgar command provides this function and you can find out here what is involved in this process.").

   ```
   fglgar gar --application helloWorld.42r --service welcomeService.42m
   ```

   This example shows the same command referencing xcf files you have
   already created for your application and
   service.

   ```
   fglgar gar --application helloWorld.xcf --service welcomeService.xcf
   ```

   A gar file is created that has the same name as your current directory. If
   you wish to specify a name for your archive, use the `--output` option.
2. Deploy the archive file.

   - If you are on the application server, you can deploy the archive with the
     gasadmin tool or the FGL web services PublishGAR tool.
     Examples are shown for each
     method:

     ```
     gasadmin gar --deploy-archive myApp.gar
     ```

     ```
     fglrun 
     $FGLDIR/web_utilities/services/deployment/bin/PublishGAR
     http://localhost:6394 deploy myApp.gar
     ```

     This example assumes the standalone dispatcher on
     your GAS installation is already started on you local machine and you have not changed the default
     port number.
   - If the GAS is on a external web server, you can deploy applications with
     PublishGAR. In this example, the gar is deployed at the base
     URL of the GAS (for example
     http://zeus:8090/gas)
     on the Web server.

     ```
     fglrun 
     $FGLDIR/web_utilities/services/deployment/bin/PublishGAR
     http://zeus:8090/gas deploy myApp.gar
     ```

   The archive is unpacked in the
   $(res.appdata.path)/deployment directory.

   The application
   deployed is named with the xcf, not the gar file
   name.
3. Enable the archive.

   This makes applications in the archive available to end users.
   - If you are on the application server, you can enable the archive with the
     gasadmin tool or the FGL Web services tool, PublishGAR.
     Examples are shown for each
     method:

     ```
     gasadmin gar --enable-archive myApp.gar
     ```

     ```
     fglrun 
     $FGLDIR/web_utilities/services/deployment/bin/PublishGAR
     http://localhost:6394 enable myApp.gar
     ```
   - If the GAS is on a remote server, you can enable applications with
     PublishGAR.

     ```
     fglrun 
     $FGLDIR/web_utilities/services/deployment/bin/PublishGAR
     http://zeus:8090/gas enable myApp.gar
     ```

The applications, services, and resources included in the archive are available for your end
users.

## Build a simple archive file

The fglgar gar command used with the `--application` option
builds an archive for your application. Using fglgar as shown in this
procedure creates this type of archive.

Create the archive file (gar).

For example, at the command line type the
following:

```
fglgar gar --application myApp.42r
```

A gar file is created that has the same name as your current directory. A
MANIFEST file and a configuration file is created automatically and included in
the archive.
> **Warning:**
>
> If a MANIFEST
> file already exists when you run the command, errors will be raised. For more information on the
> use of these options, see [fglgar](2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).").

## Build an archive file with many applications

If you have many applications to package, you may find it easier to create the
MANIFEST by hand instead of referencing each application at the command
line. The fglgar checks if a MANIFEST file is present
in the archive directory and uses it to create the gar.

1. Create your own MANIFEST file. See the MANIFEST file topic in
   the Genero Application Server User Guide.
2. Create your own application configuration files (xcf).

   Make sure to set the PATH element to `$(res.deployment.root)`.
   For example, if your compiled files (forms, modules, and so on) were in the
   /bin directory of your archive, you would specify the PATH
   element as shown:

   ```
   <PATH>$(res.deployment.root)/bin</PATH>
   ```

   For more
   information, see the Configuring applications on GAS chapter in the Genero Application Server User Guide.
3. Use the `fglgar` tool to create a Genero Archive.

   If you are in the directory containing your MANIFEST file and your program
   files:

   ```
   fglgar gar
   ```

   This creates an archive (gar) file with the same name as your current
   directory.

   If you need to specify the directory where the archive content is located, include the
   `--input-source`
   option:

   ```
   fglgar gar --input-source ./myArchiveDir
   ```

   This creates an archive file with the same name as the archive directory, drawing its content
   from the ./myArchiveDir directory.

   If you wish to specify a name for your archive, use the `--output`
   option:

   ```
   fglgar gar --input-source ./myArchiveDir --output myApp.gar
   ```

   This creates an archive file with the name myApp.gar, drawing its content
   from the ./myArchiveDir directory.

## Build an archive with public resources

An archive can contain common or public resources such as images, reports, etc., that all
deployed applications on the GAS can use. Using fglgar as shown in this
procedure creates this type of archive.

1. Put your application's public images in a dedicated directory of your archive directory. 

   You can name it, for example, "myAppPublicImages".
2. Create the archive file (gar).

   For example, at the command line type the following command with the
   `--resource` option specifying the name of the resource directory in the archive
   directory:

   ```
   fglgar gar --resource myAppPublicImages --application myApp.xcf
   ```

   A general knowledge of how the Genero Browser Client (GBC) operates can be helpful in the
   planning and deploying of Web applications. For more information, refer to the Genero Browser Client User Guide.

## Build an archive with deployment triggers

An archive can be defined with deployment parameters. These are commands that execute when
deploying and undeploying an application on the GAS. Using fglgar as shown in
this procedure creates the archive with deployment trigger options.

Deployment triggers are typically not required, you can deploy your applications without
them.

Create an archive file (gar).

For
example:

```
fglgar gar --application myApp.xcf --trigger-component cpn.gar.execution.local 
--deploy-trigger "fglrun mydeploy.42r" --undeploy-trigger "fglrun myundeploy.42r"
```

Where:

- The --trigger-component option references a trigger component in the GAS
  as.xcf that defines the runtime environment where triggers are run. If not set,
  it defaults to "cpn.gar.execution.local".
- The --deploy-trigger specifies your DEPLOY command.
- The --undeploy-trigger  options specifies the UNDEPLOY
  command.

These commands are saved in a MANIFEST file in the gar.
See the Triggers page in the Genero Application Server User Guide.
> **Warning:**
>
> If a MANIFEST
> file already exists when you run the command, errors will be raised. For more information on the
> use of these options, see [fglgar](2526-fglgar.md "The fglgar is a tool for packaging applications for deployment on any web server with Genero Application Server (GAS).").

## Deploy your application on your machine

**Before your begin:**

Once you have created an archive for your application, you can
now deploy it locally on your machine where the GAS is installed. If the
standalone dispatcher on your GAS installation is not already started, run it from the command line
using httpdispatch.

Deploy your gar file.

For example, to deploy an archive named myApp.gar using the
PublishGar tool, type the
following:

```
fglrun 
$FGLDIR/web_utilities/services/deployment/bin/PublishGAR
http://localhost:6394 deploy myApp.gar
```

Or if using the gasadmin tool, type the following
command:

```
gasadmin gar --deploy-archive  myApp.gar
```

The archive is unpacked in the
$(res.appdata.path)/deployment directory.

The application
deployed is named with the xcf, not the gar file
name.

## Enable your application on your machine

Once you have deployed your application, you can now make it available for end users by enabling
it on the machine where the GAS is installed. If the
standalone dispatcher on your GAS installation is not already started, run it from the command line
using httpdispatch.

Enable your application.

For example, using the PublishGAR tool, type the
command

```
fglrun 
$FGLDIR/web_utilities/services/deployment/bin/PublishGAR
http://localhost:6394  enable myApp.gar
```

Or if using the gasadmin tool, type the
command:

```
gasadmin gar --enable-archive myApp.gar
```

## Run the deployed application

Once the application is enabled, you can now run it. If the
standalone dispatcher on your GAS installation is not already started, run it from the command line
using httpdispatch.

In a browser enter the address of your deployed application.

http://localhost:6394/ua/r/myApp

In this example, the URL is looking for a configuration file named
myApp.xcf.

If your application is displayed and you can interact with it, you have successfully
deployed an application.
