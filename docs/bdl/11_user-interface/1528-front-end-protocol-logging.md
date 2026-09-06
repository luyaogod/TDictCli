---
title: "Front-end protocol logging"
source: "fgl-topics/c_fgl_feconn_replay.html"
breadcrumb: "User interface > User interface basics > GUI front-end connection > Front-end protocol logging"
type: "concept"
description: "GUI protocol exchanges can be logged to a file with the --start-guilog= filename option of fglrun , that can be replayed on another computer with the --run-guilog= filename option, without program ..."
---

# Front-end protocol logging

GUI protocol exchanges can be logged to a file with the
`--start-guilog=filename` option of fglrun,
that can be replayed on another computer with the
`--run-guilog=filename` option, without program files.

> **Important:**
>
> Sensitive and personal data may be written to the output. Make sure that the log output is
> written to files that can only be read by application administrators.

The `--start-guilog/--run-guilog` options are used to simulate a program execution
by connecting with a front-end, without the program files. Basically, the runtime system connects
with the front-end and replays the abstract user interface exchanges.

In order to use the `--start-guilog/--run-guilog` options, the front-end must be
started and listening for direct connections.

The `--start-guilog=filename` instructs
fglrun to start the program and log into filename, all user
interface exchanges with the front-end. This option is typically used on site, with the program
files.

With the `--run-guilog=filename` option, the runtime system
replays abstract user interface exchanges, as when the program was executed with the
`--start-guilog` option. This step can be done without the program files.

This GUI log feature can be used to set up non-regression tests for front-ends, and to provide a
log file to your support center in order to replay a specific application issue.

This solution can also be used to implement tests with the GGC tool. For more details, see the
Genero Ghost Client User Guide.

The options take the log file as parameter:

UNIX™ (shell) example:

```
$ fglrun --start-guilog=mylog.txt myprogram
```

All user interaction and AUI tree updates will be logged into the mylog.txt
file.

The log file can then be replayed with the `--run-guilog` option, to mimic the
user interaction and program, to reproduce potential issues in front-ends:

```
$ fglrun --run-guilog=mylog.txt myprogram
```

If the parent program starts other child programs with `RUN cmd [WITHOUT
WAITING]`, the parent program and each child program will write into the same log file. When
replaying the GUI log file, the runtime system is able to identify parent and child program logs in
order to restart individual processes.

When replaying the GUI log file, the runtime system prompts you for each user interaction. Type
`ENTER` to replay each event step by step, or type `c` to continue by
playing the full log:

```
$ fglrun --run-guilog=mylog.txt myprogram
enter: one step; c: continue
*** new process pid:1787036
<< 32351 3 i:meta Client{{name "GDC"}{version "3.00.06-152753"}{host  ...
>> 32351 56 o:om 0 {{an 0 UserInterface 0 {{name "guilog"} {text "guilog"} { ...
$
<< 32351 4181 i:event _om 0{}{{ActionEvent 0{{idRef "96"}}}}
*** new process pid:1787047
<< 32355 4186 i:meta Client{{name "GDC"}{version "3.00.06-152753"}{ho ...
>> 32355 4209 o:om 0 {{an 0 UserInterface 0 {{name "guilog"} {text "guilog"} ...
$c
<< 32355 5428 i:event _om 0{}{{ActionEvent 0{{idRef "96"}}}}
>> 32355 5429 o:om 1 {{rn 0}}
*** process terminated
*** process changed
<< 32351 6278 i:event _om 1{}{{ActionEvent 0{{idRef "97"}}}}
>> 32351 6278 o:om 2 {{un 95 {{selection "97"}}} {un 0 {{runtimeStatus "chil ...
>> 32351 6278 o:om 3 {{un 0 {{runtimeStatus "processing"}}}}
>> 32351 6283 o:om 4 {{un 0 {{focus "97"} {runtimeStatus "interactive"}}}}
*** new process
<< 32360 6296 i:meta Client{{name "GDC"}{version "3.00.06-152753-testEN"}{ho ...
>> 32360 6344 o:om 0 {{an 0 UserInterface 0 {{name "guilog"} {text "guilog"} ...
<< 32360 7341 i:event _om 0{}{{ActionEvent 0{{idRef "96"}}}}
>> 32360 7342 o:om 1 {{rn 0}}
*** process terminated
*** process changed
>> 32351 8099 o:om 5 {{rn 0}}
*** process terminated
bye!
```

The format of a guilog line is the
following:

```
ind pid time type: io-data
```

where:

- ind: in/out indicator 1: `<<` for input,
  `>>` for output (see below).
- pid: the process id of the fglrun process.
- time: the relative time since starting the parent fglrun, in
  milliseconds.
- type: in/out indicator 2:
  - `o` = data out (fglrun to front-end): This data is send unmodified to the
    front-end by the log player.
  - `i` = data in (front-end to fglrun): This data is informative and just displayed
    by the log player.
- io-data: Everything right of the colon (`:`) is AUI exchange
  data.

## Related links

**Related concepts**  

[fglrun](../13_programming-tools/2513-fglrun.md "The fglrun tool is the runtime system program that executes p-code programs.")

[GUI log events](1529-gui-log-events.md "GUI log events")
