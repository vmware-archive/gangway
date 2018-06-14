# Deploying Gangway

Deploying Gangway consists of writing a config file and then deploying the service.
The service is stateless so it is relatively easy to manage on Kubernetes.
How you provide access to the service is going to be dependent on your specific configuration.

A good starting point for yaml manifests for deploying to Kubernetes is in the [yaml](./yaml) directory.
This creates a namespace, configmap, deployment and service.
There is also an example ingress config that is set up to work with [Heptio Contour](https://github.com/heptio/contour), [JetStack cert-manager](https://github.com/jetstack/cert-manager) and [Let's Encrypt](https://letsencrypt.org/).
See [Dave Cheney's blog post](https://blog.heptio.com/how-to-deploy-web-applications-on-kubernetes-with-heptio-contour-and-lets-encrypt-d58efbad9f56) for an example of getting all of that configured.

You will probably have to adjust the service and ingress configs to match your environment as there is no one true way to reach services in Kubernetes that will work for everyone.

The "client secret" is embedded in this config file.
While this is called a secret, based on the way that OAuth2 works with command line tools, this secret won't be all secret.
This will be divulged to any client that is configured through gangway.
As such, it is probably acceptable to keep that secret in the config file and not worry about managing it as a true secret.

## Docker image

A recent release of Gangway is available at

```
gcr.io/heptio-images/gangway:<tag>
```

## Config file

```yaml
# The address to listen on. Defaults to 0.0.0.0 to listen on all interfaces.
# Env var: GANGWAY_HOST
# host: 0.0.0.0

# The port to listen on. Defaults to 8080.
# Env var: GANGWAY_PORT
# port: 8080

# Should Gangway serve TLS vs. plain HTTP? Default: false
# Env var: GANGWAY_SERVE_TLS
# serveTLS: false

# The public cert file (including root and intermediates) to use when serving
# TLS.
# Env var: GANGWAY_CERT_FILE
# certFile: /etc/gangway/tls/tls.crt

# The private key file when serving TLS.
# Env var: GANGWAY_KEY_FILE
# keyFile: /etc/gangway/tls/tls.key

# The cluster name. Used in UI and kubectl config instructions.
# Env var: GANGWAY_CLUSTER_NAME
clusterName: "YourCluster"

# OAuth2 URL to start authorization flow.
# Env var: GANGWAY_AUTHORIZE_URL
authorizeURL: "https://oauth.example.com/authorize"

# OAuth2 URL to obtain access tokens.
# Env var: GANGWAY_TOKEN_URL
tokenURL: "https://oauth.example.com/oauth/token"

# Endpoint that provides user profile information [optional]. Not all providers
# will require this.
# Env var: GANGWAY_AUDIENCE
audience: "https://oauth.example.com/userinfo"

# Used to specify the scope of the requested Oauth authorization.
# scopes: ["openid", "profile", "email", "offline_access"]

# Where to redirect back to. This should be a URL where gangway is reachable.
# Typically this also needs to be registered as part of the oauth application
# with the oAuth provider.
# Env var: GANGWAY_REDIRECT_URL
redirectURL: "https://gangway.yourcluster.com/callback"

# API client ID as indicated by the identity provider
# Env var: GANGWAY_CLIENT_ID
clientID: "<your client ID>"

# API client secret as indicated by the identity provider
# Env var: GANGWAY_CLIENT_SECRET
clientSecret: "<your client secret>"

# The JWT claim to use as the username. This is used in UI.
# Default is "nickname".
# Env var: GANGWAY_USERNAME_CLAIM
usernameClaim: "sub"

# The JWT claim to use as the email claim. This is used to name the
# "user" part of the config. Default is "email".
# Env var: GANGWAY_EMAIL_CLAIM
emailClaim: "email"

# The API server endpoint used to configure kubectl
# Env var: GANGWAY_APISERVER_URL
apiServerURL: "mycluster.example.com"

# The path to find the CA bundle for the API server. Used to configure kubectl.
# This is typically mounted into the default location for workloads running on
# a Kubernetes cluster and doesn't need to be set.
# Env var: GANGWAY_CLUSTER_CA_PATH
# cluster_ca_path: "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
```

A starting point for this config embedded in a configmap is at [yaml/02-config.yaml].

## Identity Provider Configs

Gangway can be used with a variety of OAuth2 identity providers.
Here are some instructions for common ones.

* [Auth0](auth0.md)
* [Google](google.md)
* [Dex](dex.md)