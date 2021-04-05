# Deploying Gangway

Deploying Gangway consists of writing a config file and then deploying the service.
The service is stateless so it is relatively easy to manage on Kubernetes.
How you provide access to the service is going to be dependent on your specific configuration.

A good starting point for yaml manifests for deploying to Kubernetes is in the [yaml](./yaml) directory.
This creates a namespace, configmap, deployment and service.
There is also an example ingress config that is set up to work with [Project Contour](https://github.com/projectcontour/contour), [JetStack cert-manager](https://github.com/jetstack/cert-manager) and [Let's Encrypt](https://letsencrypt.org/).
See [this guide](https://projectcontour.io/guides/cert-manager/) for help getting started.

You will probably have to adjust the service and ingress configs to match your environment as there is no one true way to reach services in Kubernetes that will work for everyone.

The "client secret" is embedded in this config file.
While this is called a secret, based on the way that OAuth2 works with command line tools, this secret won't be all secret.
This will be divulged to any client that is configured through gangway.
As such, it is probably acceptable to keep that secret in the config file and not worry about managing it as a true secret.

We also have a secret string that is used to as a way to encrypt the cookies that are returned to the users.
If using the example YAML, create a secret to hold this value with the following command line:

```
kubectl -n gangway create secret generic gangway-key \
  --from-literal=sessionkey=$(openssl rand -base64 32)
```

## Path Prefix

Gangway takes an optional path prefix if you want to host it at a url other than '/' (e.g. `https://example.com/gangway`).
By configuring this parameter, all redirects will have the proper path appended to the url parameters.

This variable can be configured via the [ConfigMap](https://github.com/heptiolabs/gangway/blob/master/docs/yaml/02-config.yaml#L81) or via environment variable (`GANGWAY_HTTP_PATH`).

## Detailed Instructions

The following guide is a more detailed review of how to get Gangway and other components configured in an AWS environment.
AWS is not a requirement, rather, is just an example of how a fully functional system can be constructed and components used may vary depending on the specific environment that you are targeting.

It's also important to note that this guide does not include an Auth provider, that comes in the next section:

We will use the following components:

- [gangway](https://github.com/heptiolabs/gangway): OIDC client application
- [contour](https://github.com/projectcontour/contour): Kubernetes Ingress controller
- [cert-manager](https://github.com/jetstack/cert-manager): Controller for managing TLS certificates with Let's Encrypt.

### Cert-Manager

Run:

```sh
kubectl apply -f https://raw.githubusercontent.com/jetstack/cert-manager/v0.4.1/contrib/manifests/cert-manager/with-rbac.yaml
```

Apply the following config to create the staging and production `ClusterIssuer` making sure to update the `email` field below with your email address:

```sh
cat <<EOF | kubectl create -f -
apiVersion: certmanager.k8s.io/v1alpha1
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
  namespace: cert-manager
spec:
  acme:
    email: <enter email here>
    http01: {}
    privateKeySecretRef:
      name: letsencrypt-prod
    server: https://acme-v02.api.letsencrypt.org/directory
---
apiVersion: certmanager.k8s.io/v1alpha1
kind: ClusterIssuer
metadata:
  name: letsencrypt-staging
  namespace: cert-manager
spec:
  acme:
    email: <enter email here>
    http01: {}
    privateKeySecretRef:
      name: letsencrypt-staging
    server: https://acme-staging-v02.api.letsencrypt.org/directory
EOF
```

### Contour

Run:

```sh
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
```

This will deploy Contour in the `projectcontour` namespace, and expose it using a service of type `LoadBalancer`.

### Configure DNS Record

Get the hostname of the ELB that Kubernetes created for the contour service:

```sh
kubectl get svc -n projectcontour contour -o jsonpath='{.status.loadBalancer.ingress[0].hostname}'
```

Create a wildcard CNAME record that aliases the domain under your control to the hostname of the ELB obtained above.

For instance, if you own `example.com`, create a CNAME record for `*.example.com`, so that you can access gangway at `https://gangway.example.com`.

### Deploy Gangway

#### Update config

Update the following files before deploying replacing all env variables (e.g. ${VAR}):

*Important: Each placeholder can show up multiple times in the same document. Make sure to update all occurrences.*

- docs/yaml/05-ingress.yaml

#### Deploy

Once the placeholders have been updated, deploy:

```sh
# This command will make sure that you have updated the placeholders
grep -r '\${' docs/yaml/ || kubectl apply -f docs/yaml/01-namespace.yaml -f 03-deployment.yaml -f 04-service.yaml -f 05-ingress.yaml
```

#### Create secret

Create the gangway cookies that are used to encrypt gangway cookies:

```sh
kubectl -n gangway create secret generic gangway-key \
  --from-literal=sessionkey=$(openssl rand -base64 32)
```

### Validate Certs

At this point Contour, Gangway and Cert-Manager should all be deployed.
The default Ingress example included first uses the `staging` issuer for Let's Encrypt.
This provisioner is intended to be used while the application is being setup and has no rate-limiting.
If all is successful there should be a secret named `gangway` in the gangway namespace:

```sh
$ kubectl describe secret gangway -n gangway
Name:         gangway
Namespace:    gangway
Labels:       certmanager.k8s.io/certificate-name=gangway
Annotations:  certmanager.k8s.io/alt-names=gangway.example.com
              certmanager.k8s.io/common-name=gangway.example.com
              certmanager.k8s.io/issuer-kind=ClusterIssuer
              certmanager.k8s.io/issuer-name=letsencrypt-staging

Type:  kubernetes.io/tls

Data
====
tls.crt:  3818 bytes
tls.key:  1675 bytes
```

To move to a real certificate, update the Ingress object `gangway` in the gangway namespace and change the annotation from `letsencrypt-staging` to `letsencrypt-prod`, then delete the gangway secret to have Cert-Manager request the new certificate.

At this point there should be a valid certificate for gangway serving over TLS.
Continue on to the next section to configure an Identity Provider.

## Identity Provider Configs

Gangway can be used with a variety of OAuth2 identity providers.
Here are some instructions for common ones.

* [Auth0](auth0.md)
* [Google](google.md)
* [Dex](dex.md)

## Configure Role Binding

Once Gangway is deployed and functional and the Identity Provider is functional you'll need to get a token.

1. Open up gangway and auth to your provider
2. Gangway will then return a command to run which will configure your kubectl locally
3. Configure RBAC permissions for the user. A simple example is included in this repo (`docs/yaml/role/rolebinding.yaml`). Update the user field and then apply which will make that user a `cluster-admin`.

## Docker image

A recent release of Gangway is available at

```
gcr.io/heptio-images/gangway:<tag>
```
