# Configuring Gangway

Gangway reads a configuration file on startup. The path to the configuration file must be set using the `--config` flag.

The configuration file must be in YAML format, and contain a dictionary (aka. hash or map) of key/value pairs. The available options are described below.

## Configuration Options

The following table describes the options that can be set via the YAML configuration file.

| Key | Description                                                                |
|------|----------------------------------------------------------------------------|
| `host` | The address to listen on. Defaults to `0.0.0.0` (All interfaces). |
| `port` | The port to listen on. Defaults to `8080`. |
| `serveTLS` | Should Gangway serve TLS vs. plain HTTP? Defaults to `false`.|
| `certFile` | The public cert file (including root and intermediates) to use when serving TLS. Defaults to `/etc/gangway/tls/tls.crt`. |
| `keyFile` | The private key file when serving TLS. Defaults to `/etc/gangway/tls/tls.key`. |
| `clusterName` | The cluster name. Used in the UI and kubectl config instructions |
| `authorizeURL` | OAuth2 URL to start authorization flow.|
| `tokenURL` | OAuth2 URL to obtain access tokens. |
| `audience` | Endpoint that provides user profile information [optional]. Not all providers require this. |
| `scopes` | Used to specify the scope of the requested Oauth authorization. Defaults to `["openid", "profile", "email", "offline_access"]` |
| `redirectURL` | Where to redirect back to. This should be a URL where gangway is reachable. Typically this also needs to be registered as part of the oauth application with the oAuth provider. |
| `clientID` | API client ID as indicated by the identity provider |
| `clientSecret` | API client secret as indicated by the identity provider |
| `allowEmptyClientSecret` | Some identity providers accept an empty client secret, this is not generally considered a good idea. If you have to use an empty secret and accept the risks that come with that then you can set this to true. Defaults to `false`. |
| `usernameClaim` | The JWT claim to use as the username. This is used in UI. This is combined with the clusterName for the "user" portion of the kubeconfig. Defaults to `nickname`. |
| `emailClaim` | Deprecated. Defaults to `email`. |
| `apiServerURL` | The API server endpoint used to configure kubectl |
| `clusterCAPath` | The path to find the CA bundle for the API server. Used to configure kubectl. This is typically mounted into the default location for workloads running on a Kubernetes cluster and doesn't need to be set. Defaults to `/var/run/secrets/kubernetes.io/serviceaccount/ca.crt` |
| `trustedCAPath` | The path to a root CA to trust for self signed certificates at the Oauth2 URLs |
| `httpPath` | The path gangway uses to create urls. Defaults to `""`. |
| `customHTMLTemplatesDir` | The path to a directory that contains custom HTML templates. |
