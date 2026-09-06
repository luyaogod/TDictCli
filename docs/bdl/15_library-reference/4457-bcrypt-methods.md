---
title: "security.BCrypt methods"
source: "fgl-topics/r_gws_SecurityBCrypt_methods.html"
breadcrumb: "Library reference > Extension packages > The security package > The BCrypt class > BCrypt methods"
type: "reference"
---

# security.BCrypt methods

> Methods of the security.BCrypt class.

| Name | Description |
| --- | --- |
| security.BCrypt.GenerateSalt( cost INTEGER ) RETURNS STRING | Generates the encoded value needed as input to the `HashPassword` method. |
| security.BCrypt.HashPassword( password STRING, salt STRING ) RETURNS STRING | Creates a hash password. |
| security.BCrypt.CheckPassword( password STRING, hashedPass STRING ) RETURNS INTEGER | Checks the hash password. |
