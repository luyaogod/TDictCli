---
title: "OPTIONS (Runtime)"
source: "fgl-topics/c_fgl_programs_OPTIONS_runtime.html"
breadcrumb: "Advanced features > Configuration options > OPTIONS (Runtime)"
type: "concept"
---

# OPTIONS (Runtime)

> The OPTIONS instruction inside program blocks controls program behavior at runtime.

## Syntax

```
OPTIONS options-clause [,...]
```

## Usage

Use the `OPTIONS` instruction inside a function block to control the
behavior of the runtime system for the rest of the program execution.

A program can execute successive `OPTIONS` statements at different places in
the code.

The runtime `OPTIONS` statement allows for control of the following runtime
features:

- [Defining the position of reserved lines](0925-defining-the-position-of-reserved-lines.md "The OPTIONS element LINE defines position of dedicated screen lines.")
- [Defining default TTY attributes](0926-defining-default-tty-attributes.md "The OPTIONS {INPUT|DISPLAY} ATTRIBUTES defines default TTY attributes for dialogs and display statements.")
- [Defining field tabbing order method](0928-defining-field-tabbing-order-method.md)
- [Defining the field input loop](0927-defining-the-field-input-loop.md "The OPTIONS INPUT [NO] WRAP instructions defines field wrapping in dialogs.")
- [Application termination](0929-application-termination.md "The OPTIONS TERMINATE SIGNAL defines a callback function in case of SIGTERM signal.")
- [Front-end termination](0930-front-end-termination.md "The OPTIONS CLOSE APPLICATION instruction defines the callback function in case of front-end termination.")
- [Defining the message file](0931-defining-the-message-file.md "The OPTIONS HELP FILE instruction defines the name of the message file.")
- [Defining control keys](0932-defining-control-keys.md "The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).")
- [Setting default screen modes for sub-programs](0933-setting-default-screen-modes-for-sub-programs.md "The OPTIONS RUN IN instruction defines the TTY mode to run sub-programs.")
- [Enabling/disabling SQL interruption](0934-enabling-disabling-sql-interruption.md "The OPTIONS SQL INTERRUPT instruction enables or disables SQL statement interruption.")

## Child topics

- [Defining the position of reserved lines](0925-defining-the-position-of-reserved-lines.md): The OPTIONS element LINE defines position of dedicated screen lines.
- [Defining default TTY attributes](0926-defining-default-tty-attributes.md): The OPTIONS {INPUT|DISPLAY} ATTRIBUTES defines default TTY attributes for dialogs and display statements.
- [Defining the field input loop](0927-defining-the-field-input-loop.md): The OPTIONS INPUT [NO] WRAP instructions defines field wrapping in dialogs.
- [Defining field tabbing order method](0928-defining-field-tabbing-order-method.md)
- [Application termination](0929-application-termination.md): The OPTIONS TERMINATE SIGNAL defines a callback function in case of SIGTERM signal.
- [Front-end termination](0930-front-end-termination.md): The OPTIONS CLOSE APPLICATION instruction defines the callback function in case of front-end termination.
- [Defining the message file](0931-defining-the-message-file.md): The OPTIONS HELP FILE instruction defines the name of the message file.
- [Defining control keys](0932-defining-control-keys.md): The OPTIONS logical-key KEY physical-key instruction defines physical keys for logical keys (TUI mode).
- [Setting default screen modes for sub-programs](0933-setting-default-screen-modes-for-sub-programs.md): The OPTIONS RUN IN instruction defines the TTY mode to run sub-programs.
- [Enabling/disabling SQL interruption](0934-enabling-disabling-sql-interruption.md): The OPTIONS SQL INTERRUPT instruction enables or disables SQL statement interruption.
