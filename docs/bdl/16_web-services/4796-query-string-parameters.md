---
title: "Query string parameters"
source: "fgl-topics/r_gws_restful_high_level_url.html"
breadcrumb: "Web services > RESTful web services > RESTful Web services (high-level framework) > Reference > Query string parameters"
type: "reference"
---

# Query string parameters

> You can append query string parameters to a REST service URL to request OpenAPI documents or retrieve version information about the service.

Different parameters support tasks such as selecting an OpenAPI format or version, listing all
available service versions, or providing authentication details when required.

The following table lists the supported parameters and explains how each one is used.

| Name | Value | Description |
| --- | --- | --- |
| ?openapi.json | None | Request the service's OpenAPI document in JSON. Example: http://myServer:port/gas/ws/r/group/xcf/myService?openapi.json. For more details, go to [Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.") |
| ?openapi.yaml | None | Request the service's OpenAPI document in YAML. Example: http://myServer:port/gas/ws/r/group/xcf/myService?openapi.yaml. For more details, go to [Get the OpenAPI service description](4776-get-the-openapi-description.md "Retrieve the OpenAPI description for a RESTful web service in JSON or YAML.") |
| access\_token | OAuth access token | Include this parameter to retrieve protected OpenAPI documents or access services that require authentication. Example: ?openapi.json&access\_token=token. For more details, go to [Get OpenAPI description for a secure service](4779-get-openapi-secure.md "Retrieve the OpenAPI description from a secure RESTful service by obtaining an access token and including it in the service URL.") |
| version | OpenAPI version identifier | When combined with ?openapi.json, selects the OpenAPI version to return. Example: ?openapi.json&version=v3. For more details, go to [Get the OpenAPI description for a specific version](4777-get-openapi-description-for-a-version.md "Retrieve the OpenAPI description for a specific version of a RESTful web service.") |
| version | Response format suffix | When used alone as ?version.json or ?version.yaml, triggers version-listing mode. The response lists all declared versions of the service and identifies the default version if one is defined. Any value assigned to the parameter is ignored. For more details, go to [Get available versions of a service](4778-get-available-versions-of-a-service.md "Retrieve the list of available versions of a REST service and identify the default version if defined.") |
