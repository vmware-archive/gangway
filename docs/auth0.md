# Connecting Gangway to Auth0

1. Create an account for Auth0 and login
2. From the dashboard, click "New Application"
3. Enter a name and choose "Single Page Web Applications"
4. Click on "Settings" and gather the relevant information and update the file `docs/yaml/02-configmap.yaml` and apply to cluster
5. Update the "Allowed Callback URLs" to match the "redirectURL" parameter in the configmap configured previously
6. Click "Save Changes"
7. Add Rule for adding group metadata by clicking on "Rules" from the menu
8. Give the rule a name and copy/paste the following:

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

9. Configure API Server with the following config replacing issuer-url & client-id values:

    ```
    --oidc-issuer-url=https://example.auth0.com/
    --oidc-client-id=<clientid>
    --oidc-username-claim=email
    --oidc-groups-claim=groups
    ```

## Example

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
