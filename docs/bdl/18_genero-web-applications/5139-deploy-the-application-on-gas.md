---
title: "Deploy the application on GAS"
source: "fgl-topics/t_fgl_gwa_deploy_launch_app.html"
breadcrumb: "Genero Web applications > Deploying GWA apps > Deploy the application on GAS"
type: "task"
---

# Deploy the application on GAS

> Deploy the application on your GAS.

It is assumed you have built your application. For more information on building the application,
go to [Build and test the application](5138-build-and-test-the-application.md "Build a GWA application."). In this task you will package the
`gwa_dist` in a Genero archive (gwa) file using [fglgar gwa](../13_programming-tools/2526-fglgar.md).

Once you have created an archive for your application, you can now deploy it on the GAS and/or
uploaded it to a web server.

**Steps**

1. Execute the fglgar gwa command and related options to create a Genero
   archive (gwa) file. 

   ```
   fglgar gwa --gwa gwa_dist --output myapp.gwa
   ```

   Where:

   - The `--gwa` option is mandatory. It must be set to the path of the gwa directory
     created by the gwabuildtool.
   - The `--output` option is also mandatory. You must set it to the name of the
     gwa file to generate and the relative or absolute path to where it is to be
     output.

   An archive file (`gwa`) with all the necessary files is created in the output
   path.
2. Execute the gasadmin gwa command and related options to deploy the Genero
   archive (gwa) file on the GAS. 

   For details, refer to the gasadmin topic in Genero Application Server User Guide.

   ```
   gasadmin gwa --deploy-archive myapp.gwa
   ```

   The application should now be accessible on the server where the GAS is running at a URL like
   this:
   https://host:port/gas/gwa/myapp/index.html

## Related links

**Related concepts**  

[Troubleshooting GWA apps](5142-troubleshooting-gwa-apps.md "What steps can you take if you have trouble with a Genero web application (GWA)?")

[Debugging GWA apps](5143-debugging-gwa-apps.md "Different solutions are available to debug a GWA application.")

[Creating GWA apps with Genero](5126-creating-gwa-apps-with-genero.md "Prepare the environment to build GWA applications.")

[Packaging gwa files](../13_programming-tools/2650-packaging-gwa-files.md "Using the fglgar tool to build a Genero Web Application (gwa) file allows you to deploy applications that are ready to run in a browser.")
