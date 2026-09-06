---
title: "Set up and test your environment"
source: "fgl-topics/t_gws_restful_high_level_quick_start_set_test_env.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 1: RESTful server application > Set up and test your environment"
type: "task"
---

# Set up and test your environment

> Before running the server make sure that the environment variables FGLAPPSERVER and FGLWSDEBUG are properly set.

Complete these steps to set up your system to execute a Genero Web Services server in direct
mode.

1. Set FGLAPPSERVER to 8090

   ```
   set FGLAPPSERVER=8090
   ```

   This will start the web service server application on port 8090. If this environment variable
   is not set, port number 80 is used.

   > **Warning:**
   >
   > Do not set the FGLAPPSERVER variable in production environments. In production,
   > the Genero Application Server selects the port number. For more information on deploying Web services
   > see [Web services server program deployment](4908-web-services-server-program-deployment.md "The Genero Application Server (GAS) manages web services. You must consider GAS configuration when deploying your web service in a production environment.").
2. Set FGLWSDEBUG to 3

   ```
   set FGLWSDEBUG=3
   ```

   This ensures there is detail
   in the output of the HTTP requests and responses showing the interaction between the server and its
   clients. It is useful for debugging purposes. See [FGLWSDEBUG](../07_configuration/0537-fglwsdebug.md "The FGLWSDEBUG environment variable enables web services library debugging.").
