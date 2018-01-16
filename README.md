[TOC]

gangway
=======

_(noun): An opening in the bulwark of the ship to allow passengers to board or leave the ship._

a collection of utilities that can be used to enable authn/authz flows for a kubernetes cluster.

## Deploy

### gangway manifest

gangway is comprised of a deployment and a service. The service may be exposed in any form you prefer be it Loadbalancer directly to the service or an Ingress.

[auth0 example](examples/auth0-gangway-example.yaml)



### Creating Config Secret

The [gangway config](#gangway Config) should be stored as a secret in Kubernetes becaue it contains sensitive information. To do this the value of your config file must be base64 encoded. For example the command `base64 config.yaml`  will produce results similar to the following. 

```
ICAgIGNsdXN0ZXJfbmFtZTogIllvdXJDbHVzdGVyIgogICAgYXV0aG9yaXplX3VybDogImh0dHBzOi8vZXhhbXBsZS5hdXRoMC5jb20vYXV0aG9yaXplIiAgCiAgICB0b2tlbl91cmw6ICJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tL29hdXRoL3Rva2VuIiAKICAgIGNsaWVudF9pZDogIjx5b3VyIGNsaWVudCBJRD4iCiAgICBjbGllbnRfc2VjcmV0OiAiPHlvdXIgY2xpZW50IHNlY3JldD4iCiAgICBhdWRpZW5jZTogImh0dHBzOi8vZXhhbXBsZS5hdXRoMC5jb20vdXNlcmluZm8iCiAgICByZWRpcmVjdF91cmw6ICJodHRwczovL2dhbmd3YXkueW91cmNsdXN0ZXIuY29tL2NhbGxiYWNrIiAKICAgIHNjb3BlczogWyJvcGVuaWQiLCAicHJvZmlsZSIsICJlbWFpbCIsICJvZmZsaW5lX2FjY2VzcyJdIAogICAgdXNlcm5hbWVfY2xhaW06ICJzdWIiCiAgICBlbWFpbF9jbGFpbTogImVtYWlsIg==

```

This will need to then be placed in to your secret manifest like the following example

```
apiVersion: v1
kind: Secret
metadata:
  name: gangway
data:
  # The value of gangway.yaml is a base64 encoded version of a config file
  gangway.yaml: ICAgIGNsdXN0ZXJfbmFtZTogIllvdXJDbHVzdGVyIgogICAgYXV0aG9yaXplX3VybDogImh0dHBzOi8vZXhhbXBsZS5hdXRoMC5jb20vYXV0aG9yaXplIiAgCiAgICB0b2tlbl91cmw6ICJodHRwczovL2V4YW1wbGUuYXV0aDAuY29tL29hdXRoL3Rva2VuIiAKICAgIGNsaWVudF9pZDogIjx5b3VyIGNsaWVudCBJRD4iCiAgICBjbGllbnRfc2VjcmV0OiAiPHlvdXIgY2xpZW50IHNlY3JldD4iCiAgICBhdWRpZW5jZTogImh0dHBzOi8vZXhhbXBsZS5hdXRoMC5jb20vdXNlcmluZm8iCiAgICByZWRpcmVjdF91cmw6ICJodHRwczovL2dhbmd3YXkueW91cmNsdXN0ZXIuY29tL2NhbGxiYWNrIiAKICAgIHNjb3BlczogWyJvcGVuaWQiLCAicHJvZmlsZSIsICJlbWFpbCIsICJvZmZsaW5lX2FjY2VzcyJdIAogICAgdXNlcm5hbWVfY2xhaW06ICJzdWIiCiAgICBlbWFpbF9jbGFpbTogImVtYWlsIg==

```



## gangway Config

```
# Your Cluster Name
cluster_name: "YourCluster"

# The URL to send authorize requests to
authorize_url: "https://example.auth0.com/authorize"

# URL to get a token from 
token_url: "https://example.auth0.com/oauth/token" 

# API Client ID
client_id: "<your client ID>"

# API Client Secret
client_secret: "<your client secret>"

# Endpoint that provides user profile information
audience: "https://example.auth0.com/userinfo"

# Where to redirect back to. This should be a URL
# Where gangway is reachable
redirect_url: "https://gangway.yourcluster.com/callback" 

# Used to specify the scope of the requested authorisation in OAuth. 
scopes: ["openid", "profile", "email", "offline_access"] 

# What field to look at in the token to pull the username from
username_claim: "sub"

# What field to look at in the token to pull the email from
email_claim: "email"
```



Example config file

```
---
  cluster_name: "YourCluster"
  authorize_url: "https://example.auth0.com/authorize"  
  token_url: "https://example.auth0.com/oauth/token" 
  client_id: "<your client ID>"
  client_secret: "<your client secret>"
  audience: "https://example.auth0.com/userinfo"
  redirect_url: "https://gangway.yourcluster.com/callback" 
  scopes: ["openid", "profile", "email", "offline_access"] 
  username_claim: "sub"
  email_claim: "email"
```



## API-Server flags

https://kubernetes.io/docs/admin/authentication/#configuring-the-api-server

The needed Flags are as follows
```
oidc-issuer-url: https://example.auth0.com
oidc-client-id: 3YM4ue8MoXgBkvCIHh00000000000
oidc-username-claim: sub
oidc-groups-claim: "https://example.auth0.com/groups"
```
## Auth0 Config

- RS256 signing

**Rule for adding group metadata**
```
function (user, context, callback) {
    if (user.app_metadata && 'groups' in user.app_metadata) {
      context.idToken.groups = user.app_metadata.groups;
    } else {
        context.idToken.groups = [];
    }

  callback(null, user, context);
}
```



## Build

Requirements for building

- Go
- [go-bindata](https://github.com/jteeuwen/go-bindata)
- [dep](https://github.com/golang/dep)

A Makefile is provided for building tasks. The options are as follows

```
all: deps bindata
	go build ./...

deps:
	dep ensure -v

bindata:
	go-bindata -o cmd/gangway/bindata.go templates/

test:
	go test ./...

image:
	docker build .
```

