---
title: "Implement front call modules for GBC"
source: "fgl-topics/c_fgl_frontcalls_user_ext_gwc4js.html"
breadcrumb: "Extending the language > User-defined front calls > Implement front call modules for GBC"
type: "concept"
---

# Implement front call modules for GBC

> Custom front call modules for the Genero Browser Client (GBC) front-end are implemented by using JavaScript.

## GBC custom front call basics

This topic is only an introduction to custom GBC front call implementation. For more details, see
the Genero Browser Client User Guide.

To extend the GBC with your own front calls, you must be familiar with JavaScript
programming concepts.

Custom front call module and function names must be registered in lowercase for the GBC
front-end.

With GBC, front-end calls are JavaScript functions executed locally on the workstation where the
browser is running.

Executing front calls in the context of a web browser is limited to the OS functions a web browser
can do. For example, it will not be possible to delete a file on the computer where the browser
executes.

## Customizing the GBC front-end

To integrate your custom front calls in the GBC front-end, you need to set up the GBC customization
environment.

You will have to:

1. Setup GBC customization (install [Node.js](http://nodejs.org)).
2. Extract the GBC front-end archive into a project-dir directory,
3. Copy your custom front calls JavaScript modules in the
   project-dir/customization,
4. Rebuild the GBC front-end with the gbc build utility.
5. Configure the GAS to use the customized GBC front-end.

For more details, see the Genero Browser Client User Guide.

## Structure of a custom front call JavaScript module

One JavaScript module will define a front call module implementing several front call functions.

The .js file must be copied into the
project-dir/customization directory.

A custom front call JavaScript module must have the following structure:

```
"use strict";

modulum('FrontCallService.modules.module-name', ['FrontCallService'],
  /**
   * @param {gbc} context
   * @param {classes} cls
   */
  function(context, cls) {
    context.FrontCallService.modules.module-name = {

      function-name: function (param1, ...) {
        
           ... user code ...

      return [ values ... ]

      },

      [...]   /* More functions can be defined for this module */

    };
  }
);
```

Where:

1. module-name is the name of the front call module, and corresponds to the first parameter of
   `ui.Interface.frontCall()`.
2. function-name is the name of the front
   call function, and corresponds to the second parameter of `ui.Interface.frontCall()`.
3. param1, param2 ... are the input values provided as third parameter of
   `ui.Interface.frontCall()`.
4. values is a JavaScript array containing the values to be returned in the last parameter of
   `ui.Interface.frontCall()`.

## GBC custom front call API

The following JavaScript functions are provided to implement your custom front calls:

| Method | Description |
| --- | --- |
| this.parametersError( [ message ] ) | This function can be invoked when an invalid number of parameters is passed to the front call, in order to raise on exception in the BDL program.The message parameter holds the error message to be returned to the Genero program in the second part of the error [-6333](../15_library-reference/4483-genero-bdl-errors.md) message (see front call error handling in [ui.Interface.frontCall](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")). |
| this.runtimeError( [ message ] ) | This function can be used to raise an exception in the BDL program, when the front call needs to warn the program that an error occurred.The message parameter holds the error message to be returned to the Genero program in the second part of the error [-6333](../15_library-reference/4483-genero-bdl-errors.md) message (see front call error handling in [ui.Interface.frontCall](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")). |

## Example

The JavaScript code in this example implements a GBC custom front call function `"myfunc"`
for the module `"mymodule"`:

```
"use strict";

modulum('FrontCallService.modules.mymodule', ['FrontCallService'],
  /**
   * @param {gbc} context
   * @param {classes} cls
   */
  function(context, cls) {
    context.FrontCallService.modules.mymodule = {

      myfunc: function (name) {
        if (name === undefined) {
          this.parametersError();
          return;
        }
        if (name.length === 0) {
          this.runtimeError("name shouldn't be empty");
          return;
        }
        return ["Hello, " + name + " !"];
      }

    };
  }
);
```

From the Genero BDL program:

```
DEFINE res STRING
CALL ui.Interface.frontcall( "mymodule","myfunc", ["world"] , [res] )
DISPLAY res
```
