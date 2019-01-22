# Configuring Gangway

## Configuration Options

The following list shows the configuration options available:

### host

**Description**: The address to listen on

**Default**: `0.0.0.0` (all interfaces)

**Env Var Override**: `GANGWAY_HOST`

### port

**Description**: The port to listen on

**Default**: `8080`

**Env Var Override**: `GANGWAY_PORT

### serveTLS

**Description**: Should Gangway serve TLS vs. plain HTTP?

**Default**: `false`

**Env Var Override**: `GANGWAY_SERVE_TLS`

### certFile

**Description**: The public cert file (including root and intermediates) to use when serving TLS.

**Default**: ``

**Env Var Override**: `GANGWAY_CERT_FILE`

### keyFile

**Description**: The private key file when serving TLS.

**Default** ``

**Env Var Override**: `GANGWAY_KEY_FILE`

### clusterName

**Description**: The cluster name. Used in the UI and kubectl config instructions

**Default**: ``

**Env Var Override**: `GANGWAY_CLUSTER_NAME`

### authorizeURL

**Description**: OAuth2 URL to start authorization flow.

**Default**: ``

**Env Var Override**: `GANGWAY_AUTHORIZE_URL`

### tokenURL

**Description**: OAuth2 URL to obtain access tokens.

**Default**: ``

**Env Var Override**: `GANGWAY_TOKEN_URL`

### audience

**Description**: Endpoint that provides user profile information [optional]. Not all providers require this.

**Default**: ``

**Env Var Override**: `GANGWAY_AUDIENCE`

### scopes

**Description**: Used to specify the scope of the requested Oauth authorization.

**Default**: `["openid", "profile", "email", "offline_access"]`

**Env Var Override**:  N/A

### redirectURL

**Description**: Where to redirect back to. This should be a URL where gangway is reachable. Typically this also needs to be registered as part of the oauth application with the oAuth provider.

**Default**:  ``

**Env Var Override**: `GANGWAY_REDIRECT_URL`

### clientID

**Description**: API client ID as indicated by the identity provider

**Default**: ``

**Env Var Override**: `GANGWAY_CLIENT_ID`

### clientSecret

**Description**: API client secret as indicated by the identity provider

**Default**: ``

**Env Var Override**: `GANGWAY_CLIENT_SECRET`

### allowEmptyClientSecret

**Description**: Some identity providers accept an empty client secret, this is not generally considered a good idea. If you have to use an empty secret and accept the risks that come with that then you can set this to true.

**Default**: `false`

**Env Var Override**:  N/A

### usernameClaim

**Description**: The JWT claim to use as the username. This is used in UI. Default is "nickname". This is combined with the clusterName for the "user" portion of the kubeconfig.

**Default**: `sub`

**Env Var Override**:  `GANGWAY_USERNAME_CLAIM`

### apiServerURL

**Description**: The API server endpoint used to configure kubectl

**Default**: ``

**Env Var Override**: `GANGWAY_APISERVER_URL`

### clusterCAPath

**Description**: The path to find the CA bundle for the API server. Used to configure kubectl. This is typically mounted into the default location for workloads running on a Kubernetes cluster and doesn't need to be set.

**Default**: `/var/run/secrets/kubernetes.io/serviceaccount/ca.crt`

**Env Var Override**: `GANGWAY_CLUSTER_CA_PATH`

### trustedCAPath

**Description**: The path to a root CA to trust for self signed certificates at the Oauth2 URLs

**Default**: ``

**Env Var Override**: `GANGWAY_TRUSTED_CA_PATH`

### httpPath

***Description**: The path gangway uses to create urls

**Default**: `""`

**Env Var Override**: `GANGWAY_HTTP_PATH`
