---
title: "standard.getGeolocation"
source: "fgl-topics/c_fgl_frontcall_standard_getgeolocation.html"
breadcrumb: "Library reference > Built-in front calls > Standard front calls > standard.getGeolocation"
type: "concept"
---

# standard.getGeolocation

> Returns the Global Positioning System (GPS) location of a device.

## Syntax

```
ui.Interface.frontCall("standard", "getGeolocation",
   [], [status, latitude, longitude] )
```

1. status - Holds the status of the front call execution.
2. latitude - Holds the current latitude.
3. longitude - Holds the current longitude.

## Usage

The "`getGeolocation`" front call returns the current location of the device,
based on the current GPS information.

> **Important:**
>
> To be usable with GBC, this front call needs a secure context / front-end
> connection, in addition to mobile device permissions, when required. A secure context
> is achieved by using the GAS via HTTPS, via localhost on the same machine, or running direct via
> GDC-UR. When using GDC-UR, additional security settings need to be enabled in the GDC configuration
> panel. See GDC documentation for more details.

> **Important:**
>
> For GMA / Android™,
> using the `getGeolocation` front call needs the
> `android.permission.ACCESS_FINE_LOCATION` and
> `android.permission.ACCESS_COARSE_LOCATION` Dangerous Permissions to be specified
> when building the APK.
> See [Android permissions](../17_mobile-applications/5105-building-android-apps-with-genero.md)
> for more details.

The possible values returned in the status parameter are:

- "`ok`": The device location was found.
- In case of failure, the status variable contains the error description, for example,
  "`location services not enabled`".

It is recommended that the returned coordinates are stored in `FLOAT`
variables.

If the device's location cannot be found within a given period, the front call returns an
error status.

## Example

```
DEFINE status STRING, latitude, longitude FLOAT
CALL ui.Interface.frontCall("standard", "getGeolocation",
     [], [status, latitude, longitude] )
MESSAGE SFMT(
  "Geo location: (status=%1) Latitude=%2 Longitude=%3",
  status, latitude, longitude )
```

## Related links

**Related concepts**  

[mobile.getGeolocation](3451-mobile-getgeolocation.md "Returns the Global Positioning System (GPS) location of a mobile device.")
