---
title: "Compile and start the service"
source: "fgl-topics/t_gws_restful_high_level_quick_start_compile_servive.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Quick starts > Quick start 2: RESTful server application, part 2 > Compile and start the service"
type: "task"
---

# Compile and start the service

> This describes the steps to setup the service.

If the server is running, you need to stop it. Stop a REST Web Services server running in direct
mode using the CTRL+C keys at the terminal.

1. Compile the [service](4715-add-functions-to-the-service-module.md "Define resources to create, update, and delete customers.") application:

   fglcomp myservice.4gl
2. Start the Genero RESTful Web Service [server](4710-create-the-server-module.md "The server module registers the Web Service application with the Genero Web Services (GWS) server that starts the Web service."):

   fglrun wsserver.42m
