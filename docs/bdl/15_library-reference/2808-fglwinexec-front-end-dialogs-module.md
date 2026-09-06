---
title: "fglwinexec: Front-end dialogs module"
source: "fgl-topics/r_fgl_utility_functions_fglwinexec.html"
breadcrumb: "Library reference > Utility modules > fglwinexec: Front-end dialogs module"
type: "reference"
description: "Table 1. Front-end-side dialog functions (fglwinexec.4gl) (deprecated: use ui.Interface.frontCall() instead) Function Description Important: This feature is deprecated, its use is discouraged although ..."
---

# fglwinexec: Front-end dialogs module

| Function | Description |
| --- | --- |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION winopendir( dirname STRING, caption STRING ) RETURNS STRING | Opens a dialog window to get a directory path on the front-end workstation. |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION winopenfile( dirname STRING, typename STRING, extlist STRING, caption STRING ) RETURNS STRING | Opens a dialog window to get a file to be read on the front-end workstation. |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION winsavefile( dirname STRING, typename STRING, extlist STRING, caption STRING ) RETURNS STRING | Opens a dialog window to get a file path to save data on the front-end workstation. |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION winshellexec( filename STRING ) RETURNS INTEGER | Opens a document on the workstation where the Windows front-end runs.**Note:**Microsoft™ Windows only! |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION winexecwait( command STRING ) RETURNS INTEGER | Executes a program on the workstation where the Windows front-end runs and waits for termination.**Note:**Microsoft Windows only! |
| **Important:**This feature is deprecated, its use is discouraged although not prohibited.FUNCTION winexec( command STRING ) RETURNS INTEGER | Executes a program on the workstation where the Windows front-end runs and returns immediately.**Note:**Microsoft Windows only! |
