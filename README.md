
gangway
=======

_(noun): An opening in the bulwark of the ship to allow passengers to board or leave the ship._

a collection of utilities that can be used to enable authn/authz flows for a kubernetes cluster.

## Deploy

### gangway manifest

gangway is comprised of a deployment and a service. The service may be exposed in any form you prefer be it Loadbalancer directly to the service or an Ingress.

[auth0 example](examples/auth0-gangway-example.yaml)


### Creating Config Secret

The [gangway config](#gangway Config) should be stored as a secret in Kubernetes becaue it contains sensitive information. Create a gangway config based on the example, assuming you named the file `gangway-config.yaml` run the following command: `kubectl create secret generic gangway --from-file=gangway.yaml=gangway-config.yaml`


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

