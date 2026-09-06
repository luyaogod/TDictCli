---
title: "Web Services on IBM AIX"
source: "fgl-topics/c_gws_installation_003.html"
breadcrumb: "Web services > General > Platform-specific notes > Web Services on IBM® AIX®"
type: "concept"
description: "Requirements for IBM AIX."
---

# Web Services on IBM AIX

> Requirements for IBM® AIX®.

## XL C++ Runtime for AIX version

Depending on the AIX version used, you must install the the corresponding "IBM XL C++
Runtime for AIX".

To check the installed C++ runtime version, execute the following
commands:

```
$ lslpp -l | grep "IBM XL C++ Runtime for AIX"
  xlC.aix61.rte             13.1.3.0  COMMITTED  IBM XL C++ Runtime for AIX 6.1
  xlC.rte                   13.1.3.0  COMMITTED  IBM XL C++ Runtime for AIX
```

If the C++ runtime is not installed, Web Services library usage will produce the following error
message:

```
Could not load C extension library 'com'.
Reason: A file or directory in the path name does not exist.
```

For more details about the C++ runtime package installation on AIX, see <https://www.ibm.com/support/pages/fileset-information-xlcrte>.

## OpenSSL and /deb/urandom PRNG numbers

Due to an IBM issue on 64-bit platforms, the OpenSSL
library is unable to open the system /dev/urandom device to generate a PRNG
number.

If you want to use security APIs in your GWS application (especially if you access a server in
HTTPS), install Entropy Gathering Daemon (EGD).
