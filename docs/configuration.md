# Configuring Gangway

## Configuration Options

The following list shows the configuration options available:

### host

**Description**: The address to listen on

**Default**: `0.0.0.0` (all interfaces)

**Env Var Override**: `GANGWAY_HOST`


| Name | Description                                                                | Default | Env Var Override |
|------|----------------------------------------------------------------------------|---------|------------------|
| `host` | The address to listen on | `0.0.0.0` (All interfaces) | `GANGWAY_HOST`    |
| `port` | The port to listen on | `8080` | `GANGWAY_PORT ` |
| `serveTLS ` | Should Gangway serve TLS vs. plain HTTP? | `false ` | `GANGWAY_SERVE_TLS ` |
| `certFile ` | The public cert file (including root and intermediates) to use when serving TLS. | `` | `GANGWAY_CERT_FILE` |
| `keyFile ` | The private key file when serving TLS. | `` | `GANGWAY_KEY_FILE` |
| `clusterName` | The cluster name. Used in the UI and kubectl config instructions | `` | `GANGWAY_CLUSTER_NAME` |
| `authorizeURL` | OAuth2 URL to start authorization flow.| `` | `GANGWAY_AUTHORIZE_URL` |
| `tokenURL` | OAuth2 URL to obtain access tokens. | `` | `GANGWAY_TOKEN_URL` |
| `audience` | Endpoint that provides user profile information [optional]. Not all providers require this. | `` | `GANGWAY_AUDIENCE` |
| `scopes` | Used to specify the scope of the requested Oauth authorization. | `["openid", "profile", "email", "offline_access"]` | N/A |
| `redirectURL` | Where to redirect back to. This should be a URL where gangway is reachable. Typically this also needs to be registered as part of the oauth application with the oAuth provider. | `` | `GANGWAY_REDIRECT_URL` |
| `clientID` | API client ID as indicated by the identity provider | `` | `GANGWAY_CLIENT_ID` |
| `clientSecret` | API client secret as indicated by the identity provider | `GANGWAY_CLIENT_SECRET` |
| `allowEmptyClientSecret` | Some identity providers accept an empty client secret, this is not generally considered a good idea. If you have to use an empty secret and accept the risks that come with that then you can set this to true. | `false` | N/A |
| `usernameClaim` | The JWT claim to use as the username. This is used in UI. Default is "nickname". This is combined with the clusterName for the "user" portion of the kubeconfig. | `sub` | `GANGWAY_USERNAME_CLAIM` |
| `apiServerURL` | The API server endpoint used to configure kubectl | `` | `GANGWAY_APISERVER_URL` |
| `clusterCAPath` | The path to find the CA bundle for the API server. Used to configure kubectl. This is typically mounted into the default location for workloads running on a Kubernetes cluster and doesn't need to be set. | `/var/run/secrets/kubernetes.io/serviceaccount/ca.crt` | `GANGWAY_CLUSTER_CA_PATH` | 
| `trustedCAPath` | The path to a root CA to trust for self signed certificates at the Oauth2 URLs | `` | `GANGWAY_TRUSTED_CA_PATH` |
| `httpPath` | The path gangway uses to create urls | `""` | `GANGWAY_HTTP_PATH` |