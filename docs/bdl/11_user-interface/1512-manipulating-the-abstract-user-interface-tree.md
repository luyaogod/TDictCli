---
title: "Manipulating the abstract user interface tree"
source: "fgl-topics/c_fgl_DynamicUI_019.html"
breadcrumb: "User interface > User interface basics > The abstract user interface tree > Manipulating the abstract user interface tree"
type: "concept"
description: "Modifying the AUI tree with user interface specific built-in classes The objects of the abstract user interface tree can be queried and modified at runtime with specific built-in classes like ui.Form ..."
---

# Manipulating the abstract user interface tree

## Modifying the AUI tree with user interface specific built-in classes

The objects of the abstract user interface tree can be queried and modified at runtime with
specific built-in classes like [`ui.Form`](../15_library-reference/3143-the-form-class.md "The ui.Form class provides an interface to form objects created by an OPEN WINDOW WITH FORM or DISPLAY FORM instruction."),
provided to manipulate form elements.

The code in the example gets the current window object, then gets the current form in that
window, and hides a group-box form element identified by the name
"`gb1`":

```
DEFINE w ui.Window 
DEFINE f ui.Form 
LET w = ui.Window.getCurrent()
LET f = w.getForm()
CALL f.setElementHidden("gb1",1)
```

Using the user interface specific built-in classes is the recommended way to modify the AUI tree
in your programs.

## Using low-level APIs to modify the AUI tree

In very special cases, you can also directly access the nodes of the AUI tree by using DOM
built-in classes like `om.DomDocument` and `om.DomNode`.

> **Important:**
>
> As FOUR Js continues to add new features to Genero, we encounter situations
> that forces us to modify the AUI Tree to add new element types and attributes. If you are using the
> low level API to directly modify user interface elements, your code may be impacted when changes are
> made in the AUI tree definition. To minimize the impact, consider the following course of action
> with regards to use of the DOM/SAX API:
>
> 1. During a dialog instruction, do not create or delete AUI tree nodes, or change attributes that
>    are not for decoration only (text, color and styles can be changed during dialog). When possible,
>    consider using `ui.Dialog` and
>    `ui.Form` built-in class methods instead of
>    the low-level DOM/SAX API.
> 2. Place all custom changes to the DOM/SAX API within centralized library functions that are
>    accessible to all modules, as opposed to scattering DOM/SAX calls throughout your code base.
> 3. Do not create nodes or change attributes that are not explicitly documented as modifiable. For
>    example, `TopMenu` or `ToolBar` nodes can be created and configured
>    dynamically, but it is not recommended to add `FormField` nodes to existing forms, or
>    modify the `active` attribute of fields or actions. Do not create AUI nodes, that
>    would not be created by a .42f file produced by the fglform
>    compiler from a .per source file.
> 4. AUI tree nodes that have already been sent to the front-end can only be modified by updating,
>    appending or deleting children nodes. Inserting nodes with the [`om.DomNode.insertBefore()`](../15_library-reference/3299-om-domnode-insertbefore.md "Inserts an existing node before the existing node specified.") method
>    is not supported in the AUI tree synchronization protocol, once the parent node has been sent.

To get the user interface nodes at runtime, the language provides different
kinds of API functions or methods, depending on the context. For example, to
get the root of the AUI tree, call the [`ui.Interface.getRootNode()`](../15_library-reference/3107-ui-interface-getrootnode.md "Get the root DOM node of the abstract user interface.") method. You can also get the current form node with
[`ui.Form.getNode()`](../15_library-reference/3150-ui-form-getnode.md "Get the DOM node of the form.")
or search for an element by name with the [`ui.Form.findNode()`](../15_library-reference/3151-ui-form-findnode.md "Search for a child node in the form.") method.

## Related links

**Related concepts**  

[Dialog instructions](1875-dialog-instructions.md "This section describes the dialog instructions to control application forms and the concepts related to dialog implementation.")
