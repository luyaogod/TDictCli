---
title: "Web Services changes"
source: "fgl-topics/c_fgl_Migrate_to_221_web_services.html"
breadcrumb: "Upgrading > Upgrade Guides for Genero BDL > BDL 2.21 upgrade guide > Web Services changes"
type: "concept"
---

# Web Services changes

> There are changes in support of web services in Genero 2.21.

## Operation publication restrictions on server side

If you use a variable as the name of the function to publish, you will get an error message at
compile time.

For
example:

```
com.WebOperation.CreateRPCStyle(test,"Add",add_in,add_out)
```

Where **test** is a string variable, `add_in` and `add_out` are
input and output records.

At compile time, you get the error
message:

```
error:(-9054) Web service function must be a string
```

The function name in the parameter can only be a string literal not a string variable.

Since version 2.21, FGL has introduced the concept of PUBLIC/PRIVATE functions, there is a risk
for a user publishing private functions. Private functions are not always available at runtime.

As a workaround you can add a switch based on the function name value in order to call the
appropriate publication API with the name in a string literal as shown in the following
sample:

```
CASE function_name
  WHEN "Operation1"
    LET op = com.WebOperation.CreateDocStyle(
               "Operation1","Operation1",op1_in,op1_out)
  WHEN "Operation2"
    LET op = com.WebOperation.CreateDocStyle(
               "Operation2","Operation2",op2_in,op2_out)
  OTHERWISE
    DISPLAY "ERROR"
END CASE
```

In Java or in .NET you cannot publish a different number of operations for the same service,
everything is done at compile time. For instance, when you publish a Web service in Java, only the
public methods are published as operations of the service. There is no way to add or remove some
methods at runtime. The only way you have is to create another Java class.

Be aware that if you dynamically change the service operations names, you are creating a
different service, which might be confusing for the Web service client.
