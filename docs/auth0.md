# Connecting Gangway to Auth0
Make sure you configure RS256 signing.

** Rule for adding group metadata**

```go
function (user, context, callback) {
    if (user.app_metadata && 'groups' in user.app_metadata) {
        context.idToken.groups = user.app_metadata.groups;
    } else {
        context.idToken.groups = [];
    }

  callback(null, user, context);
}
```

A typical gangway config for Auth0:

```yaml
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
apiServerURL: "https://kube-apiserver.yourcluster.com"
```
