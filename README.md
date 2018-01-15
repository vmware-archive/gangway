gangway
=======
_(noun): An opening in the bulwark of the ship to allow passengers to board or leave the ship._

a collection of utilities that can be used to enable authn/authz flows for a kubernetes cluster.


## Deploy


## Gangway Config
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

# ?
audience: "https://example.auth0.com/userinfo"

# Where to redirect back to. This should be a URL
# Where gangway is reachable
redirect_url: "https://gangway.yourcluster.com/callback" 

# ?
scopes: ["openid", "profile", "email", "offline_access"] 

# What field to look at in the token to pull the username from
username_claim: "sub"

# What field to look at in the token to pull the email from
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


