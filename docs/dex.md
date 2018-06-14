# Connecting Gangway to Dex

[Dex](https://github.com/coreos/dex) is a handy tool created by CoreOS that provides a common OIDC endpoint for multiple identity providers.
To configure Gangway to communicate with Dex some information will need to be collected from Dex.
Following OIDC standard, Dex provides a URL where its OIDC configuration can be gathered.
This URL is located at `.well-known/openid-configuration`.
If Dex is configured with an Issuer URL of `http://app.example.com` its OpenID config can be found at `http://app.example.com/.well-known/openid-configuration`.
An example of the OpenID Configuration provided by Dex:

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

 Using the Gangway example Dex's `authorization_endpoint` can be used for `authorize_url` and `token_endpoint` can be used for `token_url`.
 The Dex configuration provides a list named `claims_supported` which can be chosen from when defining both `username_claim` and `email_claim`.
 The correct claim to use depends on the upstream identity provider that dex is configured for.
 `client_id` and `client_secret` are strings that can be any value, but they must match the Client ID and Secret in your Dex configuration.
