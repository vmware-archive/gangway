
gangway
=======

_(noun): An opening in the bulwark of the ship to allow passengers to board or leave the ship._

a collection of utilities that can be used to enable authn/authz flows for a kubernetes cluster.

## Deploy

### Docker image
`gcr.io/heptio-images/gangway:<tag>`

### gangway manifest

gangway is comprised of a deployment and a service. The service may be exposed in any form you prefer be it Loadbalancer directly to the service or an Ingress.

[auth0 example](examples/auth0-gangway-example.yaml)


### Creating Config Secret

The [gangway config](#gangway Config) should be stored as a secret in Kubernetes becaue it contains sensitive information. Create a gangway config based on the example, assuming you named the file `gangway-config.yaml` run the following command: `kubectl create secret generic gangway --from-file=gangway.yaml=gangway-config.yaml`


## gangway Config

```
# The cluster name
clusterName: "YourCluster"

# The URL to send authorize requests to
authorizeURL: "https://example.auth0.com/authorize"

# URL to get a token from
tokenURL: "https://example.auth0.com/oauth/token"

# API client ID as indicated by the identity provider
clientID: "<your client ID>"

# API client secret as indicated by the identity provider
clientSecret: "<your client secret>"

# Endpoint that provides user profile information
audience: "https://example.auth0.com/userinfo"

# Where to redirect back to. This should be a URL
# Where gangway is reachable
redirectURL: "https://gangway.example.com/callback"

# Used to specify the scope of the requested Oauth authorization.
scopes: ["openid", "profile", "email", "offline_access"]

# The JWT claim to use as the Kubnernetes username
usernameClaim: "sub"

# The JWT claim to use as the email claim
emailClaim: "email"
```

Example config file

```
---
clusterName: "YourCluster"
authorizeURL: "https://example.auth0.com/authorize"
tokenURL: "https://example.auth0.com/oauth/token"
clientID: "<your client ID>"
clientSecret: "<your client secret>"
audience: "https://example.auth0.com/userinfo"
redirectURL: "https://gangway.example.com/callback"
scopes: ["openid", "profile", "email", "offline_access"]
usernameClaim: "sub"
emailClaim: "email"
```

## Connecting Gangway to Dex
  		  
[Dex](https://github.com/coreos/dex) is a handy tool created by CoreOS that provides a common OIDC endpoint for multiple identity providers. To configure Gangway to communicate with Dex some information will need to be collected from Dex. Following OIDC standard, Dex provides a URL where its OIDC configuration can be gathered. This URL is located at `.well-known/openid-configuration`.If Dex is configured with an Issuer URL of `http://app.example.com` its OpenID config can be found at `http://app.example.com/.well-known/openid-configuration`. An example of the OpenID Configuration provided by Dex:
 
 ```
 {
   "issuer": "http://app.example.com",
   "authorization_endpoint": "http://app.example.com/auth",
   "token_endpoint": "http://app.example.com/token",
   "jwks_uri": "http:/app.example.com/keys",
   "response_types_supported": [
     "code"
   ],
   "subject_types_supported": [
     "public"
   ],
   "id_token_signing_alg_values_supported": [
     "RS256"
   ],
   "scopes_supported": [
     "openid",
     "email",
     "groups",
     "profile",
     "offline_access"
   ],
   "token_endpoint_auth_methods_supported": [
     "client_secret_basic"
   ],
   "claims_supported": [
     "aud",
     "email",
     "email_verified",
     "exp",
     "iat",
     "iss",
     "locale",
     "name",
     "sub"
   ]
 }
 ```
 
 Using the Gangway example Dex's `autorization_endpoint` can be used for `authorize_url` and `token_endpoint` can be used for `token_url`. The Dex configuration provides a list named `claims_supported` which can be chosen from when defining both `username_claim` and `email_claim`. `client_id` and `client_secret` are strings that can be any value, but they must match the Client ID and Secret in your Dex configuration. 

## API-Server flags

gangway requires that the Kubernetes API server is configured for OIDC:

https://kubernetes.io/docs/admin/authentication/#configuring-the-api-server

```
kube-apiserver
...
--oidc-issuer-url "https://example.auth0.com"
--oidc-client-id 3YM4ue8MoXgBkvCIHh00000000000
--oidc-username-claim sub
--oidc-groups-claim "https://example.auth0.com/groups"
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

- Go (built with 1.10)
- [go-bindata](https://github.com/jteeuwen/go-bindata)
- [dep](https://github.com/golang/dep)

A Makefile is provided for building tasks. The options are as follows

Getting started is as simple as:
```
$ make setup
$ make
```
