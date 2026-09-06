---
title: "Implement front call modules for GMA"
source: "fgl-topics/c_fgl_frontcalls_user_ext_gma.html"
breadcrumb: "Extending the language > User-defined front calls > Implement front call modules for GMA"
type: "concept"
description: "Custom front call modules for the Android front-end are implemented by using the API for GMA front calls in Java."
---

# Implement front call modules for GMA

> Custom front call modules for the Android™ front-end are implemented by using the API for GMA front calls in Java.

## GMA custom front call basics

In order to extend the GMA with your own front calls, you must be familiar with Java
programming concepts, and if you want to interface with Android apps, understand concepts such as Android Activity and Intent.

The API for GMA front calls is based on the following Java interfaces:

- `com.fourjs.gma.extension.v1.IFunctionCallController`
- `com.fourjs.gma.extension.v1.IFunctionCall`

The front call function controller (`IFunctionCallController`) is implemented
by the GMA, it is used to notify function call results, raise runtime exceptions and invoke
activities.

The front call function body (`IFunctionCall`) implements the actual custom
front call code.

The steps to implement an `IFunctionCall` class are:

1. Create a Java source file with the name of the front call function, for example:
   "`getPhoneId.java`", that implements the `IFunctionCall`
   interface.
2. Define the Java package name identifying the front call module, for example: "`package
   com.mycompany.utilities;`".
3. Define a private `IFunctionCallController` object reference to handle the
   function controller.
4. Implement the `setFunctionCallController()` method for the function controller
   registration.
5. Implement the `invoke()` method to perform the actual front call task. In this
   method, use the controller's `returnValues()` method to return values from the front
   call. If needed, you can raise runtime errors with controller's `raiseError()`
   method. It is also possible to start an Android Activity with the `startActivity*` controller methods.
6. If an activity is started with controller's `startActivityForResult` method,
   implement the `onActivityResult()` method in the function body class, to handle the
   end of the activity, and call controller's `returnValues()` method to return values
   from the front call.
7. If needed, implement the `onSaveInstanceState()` and the
   `onRestoreInstanceState()` methods, to respectively save and restore information when
   Android has to suspend the
   application.

In any case, the `IFunctionCall` class must either call the controller's
`returnValues()` or `raiseError()` methods to give the control back to
the Genero program.

## The `com.fourjs.gma.extension.v1.IFunctionCall` interface

| Method | Description |
| --- | --- |
| void setFunctionCallController(IFunctionCallController controller) | This method binds the front call function controller object to the function body object.The controller parameter is the `IFunctionCallController` object to bind with the front call function body object. |
| abstract void invoke(Object[] args) throws IllegalArgumentException | This method performs the front call. It will be called when the front call is executed from the Genero program.The args parameter is a variable list of parameters passed to the front call. This corresponds to the third argument of [ui.Interface.frontCall](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.") |
| void onSaveInstanceState(Bundle state) | Saves the state of an ongoing function call when Android needs to suspend the application.The state parameter is the bundle to save the state to. |
| void onRestoreInstanceState(Bundle state) | Restores the state of an ongoing function call, when Android needs to restore the application.The state parameter is the bundle to restore the state from. |
| void onActivityResult(int resultCode, Intent data) | Callback invoked when an activity started through `IFunctionCallController.startActivityForResult` finishes.The resultCode parameter is the integer result code returned by the child activity through its `setResult()` method.The data parameter is an Intent object, which can return result data to the caller (various data can be attached to Intent "extras"). |

## The `com.fourjs.gma.extension.v1.IFunctionCallController` interface

| Method | Description |
| --- | --- |
| void returnValues(IFunctionCall functionCall, Object...values) | Notifies the controller that the front call function call has finished successfully. To be called typically at the end of the `IFunctionCall.invoke()` method.The functionCall parameter is the current `IFunctionCall` object invoked.The values parameter defines the variable list of front call function return values. This corresponds to the fourth parameter of [ui.Interface.frontCall](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end."). |
| void raiseError(IFunctionCall functionCall, String message) | Notifies the controller of an error in the front call function call. This leads to a BDL runtime exception. To be called if needed within the `IFunctionCall.invoke()` method.The functionCall parameter is the current `IFunctionCall` object invoked.The message parameter holds the error message to be returned to the Genero program in the second part of the error [-6333](../15_library-reference/4483-genero-bdl-errors.md) message (see front call error handling in [ui.Interface.frontCall](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.")). |
| void startActivity(IFunctionCall functionCall, Intent intent) | Starts a new activity. The function call won't be notified of the end of the activity. The Genero program will run in parallel of this activity. The behavior is similar to a `RUN WITHOUT WAITING`.The functionCall parameter is the current `IFunctionCall` object invoked.The intent parameter describes the activity to start. |
| void startActivityForResult(IFunctionCall functionCall, Intent intent) | Starts a new activity. The function call won't be notified of the end of the activity. The Genero program will remain blocked as long as the started activity isn't finished. The behavior is similar to a `RUN`.The method `IFunctionCall.onActivityResult` will be called once the activity finishes.The functionCall parameter is the current `IFunctionCall` object invoked.The intent parameter describes the activity to start. |
| IClientHandler getClientHandler() | Returns an `IClientHandler` object, which is able to interact with Genero applications. A client handler can post action execution, cancel action future execution and check if an action execution is going to be executed. It can also define an application state change listener which triggers callbacks, when the app goes to pause, is resumed or is stopped.For more details about the `IClientHandler` interface, see [The com.fourjs.gma.extension.v1.IClientHandler interface](2720-implement-front-call-modules-for-gma.md) |
| Activity getCurrentActivity() | Returns the current `Activity` object. Provided in case if you need to pass the current activity to an Android API requiring this object.Don't use the returned activity to start other activities (don't call `Activity.startActivity` or `Activity.startActivityForResult`), use the helpers of the current interface instead. |

## The `com.fourjs.gma.extension.v1.IClientHandler` interface

| Method | Description |
| --- | --- |
| void postDialogAction(String action) | Post a dialog action to be triggered as soon as a BDL dialog is active and the given action name fits one of the active actions of the current dialog.The action parameter defines the name of the dialog action. |
| boolean isDialogActionPending(String action) | Checks if a given action is still pending (the action was posted using `postDialogAction()`, but the action was still not activated by the Genero application and thus was not sent to the runtime)The action parameter defines the name of the dialog action. |

## Calling the custom front call from BDL

In the Genero program, use the [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.") API to call the front-end function. This method takes
the front call module name as first parameter and the front call function name as second parameter. The
front call module name is defined by the Java package name of the custom class implementing the
`IFunctionCall` interface, and the front call function name is defined by the name of
the class.

For example, if you implement the following front call function:

```
package com.mycompany.utilities;
...
public class GetPhoneId implements IFunctionCall {
...
```

The Genero program code must pass the Java package name "`com.mycompany.utilities`" as
front call module name and the class name "`GetPhoneId`" as front call function name:

```
CALL ui.Interface.frontCall("com.mycompany.utilities", "GetPhoneId", ["John DOE"], [msg])
```

## Deploying the custom front call

The compiled Java classes implementing the front calls must be included in the mobile application
Android package
(.apk), which is created in the Genero Studio deployment procedure.
The same GMA package building rules apply for front calls and for simple Java extensions. See [Packaging custom Java extensions for GMA](2697-packaging-custom-java-extensions-for-gma.md "Custom Java extension must be integrated in the GMA to run on Android devices.") for more details.

## Example

The example implements a HelloWorld call as a front call module.

HelloWorld.java:

```
package com.mycompany.testmodule;

import android.content.Intent;
import android.os.Bundle;

import com.fourjs.gma.extension.v1.IFunctionCall;
import com.fourjs.gma.extension.v1.IFunctionCallController;

public class HelloWorld implements IFunctionCall {

    private IFunctionCallController mController;

    @Override
    public void setFunctionCallController(IFunctionCallController controller) {
        mController = controller;
    }

    @Override
    public void invoke(Object[] args) throws IllegalArgumentException {
        if (args.length != 1) {
            throw new IllegalArgumentException("HelloWorld takes one argument");
        }

        mController.returnValues(this, "Hello " + args[0].toString());
    }

    @Override
    public void onSaveInstanceState(Bundle state) {
    }

    @Override
    public void onRestoreInstanceState(Bundle state) {
    }

    @Override
    public void onActivityResult(int returnCode, Intent data) {
    }
}
```

In order to invoke the HelloWorld front-end function, use the `ui.Interface.frontCall()`
API in the Genero program:

```
MAIN
  DEFINE msg STRING
  MENU
    ON ACTION frontcall ATTRIBUTES(TEXT="Call custom front call")
      CALL ui.Interface.frontCall("com.mycompany.testmodule", "HelloWorld", ["John DOE"], [msg])
    ON ACTION quit
      EXIT MENU
  END MENU
END MAIN
```

## Related links

**Related concepts**  

[Executing Java code with GMA](2693-executing-java-code-with-gma.md "Executing Java code with GMA")
