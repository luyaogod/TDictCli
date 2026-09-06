---
title: "The model-view-controller paradigm"
source: "fgl-topics/c_fgl_prog_dialogs_mvc_paradigm.html"
breadcrumb: "User interface > User interface programming > Dialog programming basics > The model-view-controller paradigm"
type: "concept"
---

# The model-view-controller paradigm

> The dynamic user interface architecture is based on the Model-View-Controller (MVC) paradigm.

The model defines the object to be displayed (typically the application data that is stored in
program variables). The view defines the decoration of the model (how the model must be displayed to
the screen, this is typically the form). The controller is the interactive instruction that
implements the program code to handle the model.

Views are defined in the abstract user interface tree from compiled [.42f forms](1662-form-specification-files.md "Form specification files are the source files defining the layout and content of application forms.") loaded by programs. The [program variables](../08_language-basics/0686-variables.md "Explains how to define program variables.") act as models, and you implement the
controllers with [interactive instructions](2220-what-are-dialog-controllers.md "Application forms are controlled by interactive instruction blocks called dialogs. These blocks perform the common tasks associated with the form, such as field input and action handling."),
such as `DIALOG` or `INPUT`. Controllers also define action handlers
that contain the program code to be executed when an action view is triggered.

Normally the controllers are not intended to provide any decoration information, as that is the
purpose of views. Over the course of the history of the language, however, some interactive
instructions such as `MENU` define both the
controller and some presentation information such as menu title, command labels, and comments. In
this case, the runtime system automatically creates the view with that information; you can still
associate other views to the same controller.

## Related links

**Related concepts**  

[The dynamic user interface](1509-the-dynamic-user-interface.md "The dynamic user interface is the base concept of the Genero user interaction components.")

[The abstract user interface tree](1510-the-abstract-user-interface-tree.md "The abstract user interface tree is the XML representation of the application forms displayed to the end user.")
