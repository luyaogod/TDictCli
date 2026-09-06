---
title: "System packages"
source: "fgl-topics/c_fgl_installation_syspacks.html"
breadcrumb: "Installation > Software requirements > System packages"
type: "concept"
---

# System packages

> Some Genero BDL features need specific operating system packages.

## Security note: OpenSSL 3.0 LTS support

Starting with FGLGWS 3.21.01, 4.01.05 and 5.00.00, OpenSSL 3.0 LTS is required for
encryption and security.

Because [OpenSSL 1.1.1 goes EOL in September 2023](https://www.openssl.org/blog/blog/2023/03/28/1.1.1-EOL/) (external link), it is now
mandatory to use OpenSSL 3.0 LTS to get the latest security fixes.

When installing an FGLGWS package, OpenSSL 3.0 libs will be provided in FGLDIR, if no OpenSSL 3.0
exists on the system.

Starting with OpenSSL 3.0, the SHA-1 digest algorithm is no longer supported by default. The
OpenSSL 3.0 libs provided in FGLDIR still have SHA-1 digest activated by default. If you want to
enable SHA-1 with the system OpenSSL 3.0 libs, use a command such as update-crypto-policies
--set DEFAULT:SHA1 in order to use SHA-1. However, the SHA-1 digest algorithm is no longer
recommended, because it is increasingly vulnerable as computers become more and more powerful. If
you are using SHA-1 with GWS crypto APIs, consider moving to SHA-256 or to a stronger secure hash
algorithm.

See [GWS Security](../16_web-services/4556-security.md "These topics cover security for Genero Web Services.") for more details about security
and encryption with GWS.

## Curses (or NCurses) library

On Unix-like platforms, the Genero runtime system (fglrun) needs the Curses
library to be installed on the system.

The Curses library is required for the text mode (for INFORMIXTERM=terminfo). However, fglrun is
linked directly to the Curses library so it needs to be installed on the system even when the GUI
mode is used.

Depending on the operating system, fglrun will require
libcurses, libncurses, libncurses5,
libncursesw5, libncurses6, or
libncursesw6 to be installed on the system.

For UTF-8 support in text mode, you need the wide-char Curses shared library
(libncursesw.so.5 for example). Make sure that the Curses software package
installed on your system includes this library. See [Platform specific notes](0040-platform-specific-notes.md "These notes provide operating system specific information to use Genero BDL on the platform.").

If the required Curses library is not installed, executing fglrun will report
a shared library not found error.

## Related links

**Related concepts**  

[TERMINFO terminal capabilities](../11_user-interface/1532-terminfo-terminal-capabilities.md "TERMINFO terminal capabilities")
