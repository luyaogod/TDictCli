---
title: "Implement front call modules for GMI"
source: "fgl-topics/c_fgl_frontcalls_user_ext_gmi.html"
breadcrumb: "Extending the language > User-defined front calls > Implement front call modules for GMI"
type: "concept"
---

# Implement front call modules for GMI

> Custom front call modules for the iOS front-end are implemented by using the API for GMI front calls in Objective-C.

## GMI custom front call basics

In order to extend the GMI with your own front calls, you must be familiar with Objective-C
programming, and if you want to interface with iOS Apps, have a knowledge of the iOS
API.

> **Important:**
>
> Before starting with GMI front call implementation, you need to get the GMI
> package and unzip the archive into the FGLDIR directory, as described in the prerequisites sections
> of [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.").

The API for GMI front calls is based on the `FrontCall` class and the
`FrontCallHelper` and `FunctionCall` protocols. You can find these in
the file frontcall.h in the in the $FGLDIR/include/gmi
diretory.

To implement custom front calls, write a class which extends `FrontCall` and
implements the "`moduleName`" and "`execute:retCount:params`" methods
as well as the "`initWithFunctionModuleHelper:`" initializer.

Follow these steps to implement a custom front call module for the GMI:

1. Import the `frontcall.h` header file in your source.
2. Define an interface (`MyFrontCall`) which extends `FrontCall`.
3. Create the class (`MyFrontCall`) which implements this interface:
   1. Implement the `- (instancetype) initWithFunctionModuleHelper:(id)aHelper` initializer,
      calling `[superinitWithFunctionModuleHelper:aHelper]` to pass the
      `FrontCallHelper` to the base implementation.
   2. Implement the `- (NSString*) moduleName` method, returning the name of the front call
      module.
   3. Implement the `- (void)execute:(NSString)name retCount:(int)retCount params:(NSArray)params`
      method, defining the body of your front calls. See below for details about the `execute`
      method.

## API to implement custom front calls in GMI

To get parameters passed from the Genero program to the front call, and return values from
the front call to the Genero program, use the following macros and methods of the
`FrontCall` class:

| Macro / Method | Description |
| --- | --- |
| (void) FC_REQUIRED_PARAMS(count) | Checks that the number of parameters passed by the Genero program equals count. This macro will raise an error in the Genero program if not enough parameters were passed. |
| (NSString *) FC_PARAM(index) | Get the string parameter passed to the front call, at the given position. If the parameters are of a different type, use the `doubleValue`, `floatValue` and `integerValue` methods on `NSString` or a `NSScanner`, to convert the parameter to the expected type. |
| (int) FC_PARAM_INT(index) | Get the int parameter passed to the front call, at the given position. |
| (void) intResult:(int) intValue | Ends the front call by returning one integer to Genero. |
| (void) doubleResult:(double) doubleValue | Ends the front call by returning one double to Genero. |
| (void) stringResult:(NSString * ):stringValue* | Ends the front call by returning one string to Genero. |
| (void) startResult | Initiate setting multiple result values.Must be followed by `add*` function calls and ended with `endResult`. |
| (void) addIntResult:(int)intValue | Add an integer to the list of results returned.To be used after a `startResult` call. |
| (void) addDoubleResult:(double) doubleValue | Add a double to the list of results returned.To be used after a `startResult` call. |
| (void) addStringResult:(NSString * ) stringValue* | Add a string to the list of results returned.To be used after a `startResult` call. |
| (void) endResult | Finalize the setting of multiple result values and return the results to the Genero program, with front call error code zero (indicating success). |
| (void) ok | Ends the front call without returning any value to Genero, indicating that the front call execution was successful. |
| (void) error(FCErrorCode):error | Ends the front call with a specific front call error code defined in FCErrorCode enum in frontcall.h, to indicate that front call execution failed, typically because of invalid parameters or invalid function name. |
| (void) errorWithMessage(NSString * )message* | Ends the front call with front call return code -4 (maps to BDL error [-6333](../15_library-reference/4483-genero-bdl-errors.md)), and a user-defined error message, that can be read with `ERR_GET()` in the Genero program. |
| (void) willSetResultLater | To be called at the end of the `execute` function, if result values are intended to be set after the `execute` function does a return.If the `willSetResultLater` function is used, the current front call will not end until one of the result functions is called.For example, if your front call opens a message box, the `execute` function will return before one of the message box buttons are selected. Once a button is pressed, the front call result value is set. |

## Calling the custom front call from BDL

In the Genero program, use the [`ui.Interface.frontCall()`](../09_advanced-features/0964-ui-interface-frontcall.md "ui.Interface.frontCall performs a function call to the current front-end.") API to call the front-end function.
This method takes the front call module name as first parameter and the front call function name as
second parameter.

The front call module name is defined by the string value returned from the
`-(NSString * ) moduleName*` method of your front call implementation, and
the front call function name is passed to the `execute` method you
implemented as first parameter (name).

For example, if you implement the following class:

```
#import <gmi/frontcall.h>
...

@interface MyFrontCall : FrontCall
   ...
@end

@class MyFrontCall

-(instancetype) initWithFunctionModuleHelper:(id)aHelper
{
   if (self = [superinitWithFunctionModuleHelper:aHelper]) {
      ...
   }
   return self;
}

-(NSString*) moduleName{
   return @"MyModule";
}

-(void)execute:(NSString)name
               retCount:(int)retCount
               params:(NSArray)params
{
   [super execute:name retCount:retCount params:params];
   if ([[name lowercaseString]isEqualToString:@"myfrontcall"]) {
      ...
```

The Genero program code must pass the module name "`MyModule`" as front call
module name and the class name "`MyFrontCall`" as front call function
name:

```
CALL ui.Interface.frontCall("MyModule", "MyFrontCall", ["John DOE"],[msg])
```

## Custom front call implementation details (execute method)

First of all, call the execute method of the parent `FrontCall` class, right
at the top of the `execute`
method:

```
   [super execute:name retCount:retCount params:params];
```

The `execute` method must check the name of the front call function passed
as parameter, to perform the expected code. This is the function name passed to the
`ui.Interface.frontCall()` call in the Genero program:

```
   if([[name lowercaseString] isEqualToString:@"myfunction"]) {
```

Implement
the body of the front call function in the `if()` block as follows:

Add an `assert()` line, to make sure that the number of return values
match:

```
      assert(retCount == 2);
```

In order to get the parameters passed from the Genero program, use the FC\_\* macros in the
body of your front call function.

First, check that the number of parameters passed is correct, with the
`FC_REQUIRED_PARAMS(count)`
macro:

```
      FC_REQUIRED_PARAMS(3);
      ...
```

Retrieve the parameters passed to the front call with the `FC_PARAM(index)`
or `FC_PARAM_INT(index)` macros, which return a `NSString*`
and an `int` respectively. If needed, use the `doubleValue`,
`floatValue` and `integerValue` methods on `NSString` or a
`NSScanner`, to convert the parameter to the expected type:

```
      NSString * info = FC_PARAM(0);
      int v1 = [FC_PARAM(1) integerValue];
      double v2 = [FC_PARAM(2) doubleValue];
```

Implement the actual code of the front call.

To return values to Genero, use one of the helper methods such as
`intResult:value`, if a single value must be returned
to the Genero program. If more than one value must be returned, build a return set with the
`startResult`, `add*Result` and `endResult`
methods:

```
    [self startResult];
    [self addIntResult:isIpad];
    [self addIntResult:canLocate];
    [self endResult];
```

If the front call displays a UI (for example, an `UIAlertController` or
displays a customer `UIViewController`), call the `willSetResultLater`
method of the `FrontCall` class, to avoid having the control flow returned to the
Genero program upon exit of the `execute`
method:

```
    [self willSetResultLater];
```

Additionally, if you call the `willSetResultLater` method, you need to call one of
the result methods like `stringResult` at a later time.

## Deploying the custom front call

The complied Objective-C classes must be included in the iOS app build process.

The same app building rules apply for custom front calls as for C extensions.

See [Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.") for more details.

## Example

In this example, the `ExtensionFrontCall` class implements two front calls:
"isipad" and "logindialog".

We start by defining the interface for the custom front call module:

```
@interface ExtensionFrontCall : FrontCall<UIAlertViewDelegate>
@end
```

The `ExtensionFrontCall` class extends `FrontCall`, and
implements the `UIAlertViewDelegate` protocol which is used by the
"logindialog" front call.

Next, we start the implementation of the interface:

```
@implementation ExtensionFrontCall

-(instancetype) initWithFunctionModuleHelper:(id)aHelper
{
   if (self = [superinitWithFunctionModuleHelper:aHelper]) {
   }
   return self;
}

-(NSString*) moduleName{
   return @"ExtensionFrontCall";
}

-(void)execute:(NSString)name retCount:(int)retCount params:(NSArray)params {
   [super execute:name retCount:retCount params:params];
   ...
}

...
@end
```

We use the standard initializer which will be called by GMI on start-up and define
"`ExtensionFrontCall`" as module name by returning it from the
`moduleName` method.

We also start the implementation of the `execute` method by calling the super
method.

Front call modules are compiled with the help of Makefile-gmi to a static IOS library and can be
tested instantly on the command line with the help of Makefile-gmi. The static library can be either
passed as an command line argument to gmibuildtool or placed in the
gmi sub directory of a project directory. In both cases GMI detects the
extension by enumerating all front call descendant classes upon start-up.

When using more than one extension module, take care that class names and object/library file
names are distinct.

**The `isipad` front call example**

This front call simply returns the information on which device GMI is running. If it is an iPad,
the integer 1 will be returned to the Genero
program:

```
if ([[name lowercaseString] isEqualToString:@"isipad"]) {
    assert(retCount == 1);
    BOOL isIpad = UI_USER_INTERFACE_IDIOM() == UIUserInterfaceIdiomPad;
    [self startResult];
    [self addIntResult:isIpad];
    [self endResult];
}
```

After checking that only one return parameter was defined in Genero, the code identifies the
platform with the `UI_USER_INTERFACE_IDIOM()` API and stores the result in the
`isIpad` variable.

The next three lines return the result value to Genero, by starting a result block with
`startResult`, adding an int to the return set with `addIntResult`,
and finally calling `endResult` to send the result to the Genero program.

The same behavior can be achieved with a single line: `[self intResult:isIpad];`,
since we only return one result value.

The Genero program calls the `isIPad` front call as follows:

```
DEFINE res INTEGER
CALL ui.Interface.frontCall( "ExtensionFrontCall", "isipad", [ ], [res] )
```

**The
`logindialog` front call example**

This front call displays a log-in dialog to the user. It expects two parameters (the title and
the message for the log-in dialog), and returns the log-in name and the password entered by the end
user:

```
if([[name lowercaseString] isEqualToString:@"logindialog"]) {
    assert(retCount == 2);
    FC_REQUIRED_PARAMS(2);
    NSString *title = FC_PARAM(0);
    NSString *message = FC_PARAM(1);
    UIAlertView *alert = [[UIAlertView alloc]
      initWithTitle:title
      message:message
      delegate:self
      cancelButtonTitle:NSLocalizedString(@"Cancel",@"Cancel")
      otherButtonTitles:NSLocalizedString(@"OK",@"OK"),nil];
    alert.alertViewStyle = UIAlertViewStyleLoginAndPasswordInput;
    [alert show];
    [self willSetResultLater];
}
```

We first check that two result values were set in Genero and that two parameters were supplied to
the front call.

Then we use the `FC_PARAM` macro to fetch the parameters and assign them to
`NSStrings`.

Then we allocate and initialize an `UIAlertView` with the given message and title
and set the `alertViewStyle` to
"`UIAlertViewStyleLoginAndPasswordInput`", so that one plain text field and one
password field will be displayed on the alert.

In the `initWithTitle` call we also set "`self`" as the delegate of
the alert so that we receive callbacks after user input (we had added the
`UIAlertViewDelegate` protocol to our `ExtensionFrontCall` interface
definition).

Finally, we call `willSetResultLater`, to keep the control flow in iOS. If we
don’t call this function, GMI concludes the front call was not handled by the
`execute` function (as none of the `xxxResult` functions was called
inside), and the front call will fail with a "Frontcall not found" error message.

The `ExtensionFrontCall` class implements the
`alertView:didDismissWithButtonIndex:` method from the `UIAlertViewDelegate`
protocol:

```
pragma mark UIAlertViewDelegate(void)

alertView:(UIAlertView *)alertViewdidDismissWithButtonIndex:(NSInteger)buttonIndex {
   [self startResult];
   if (buttonIndex != alertView.cancelButtonIndex) {
      [self addStringResult:[alertViewtextFieldAtIndex:0].text];
      [self addStringResult:[alertViewtextFieldAtIndex:1].text];
   } else {
      [self addStringResult:nil];
      [self addStringResult:nil];
   }
   [self endResult];
}
```

This method is called after the user has tapped on one of the buttons and the view has been
dismissed. Inside this method, we first call `startResult` to enable adding
more than one return value.

If the tapped button was not the Cancel button, we add the values of the log-in and password
fields as strings to the results and then call `endResult` to return the control flow
to the Genero program.

The Genero program calls the log-in dialog front call as follows:

```
DEFINE ul, up STRING
CALL ui.Interface.frontCall( "ExtensionFrontCall", "logindialog",
                             ["MyApp","User login"],
                             [un, up] )
IF up IS NULL THEN
   ERROR "Login canceled"
   EXIT PROGRAM
END IF
```

> **Tip:**
>
> The file userextension.m of the GMI Extension project contains a complete
> example on how to write custom front calls.

## Related links

**Related concepts**  

[Building iOS apps with Genero](../17_mobile-applications/5109-building-ios-apps-with-genero.md "Genero provides a command-line tool to build applications for iOS devices.")
