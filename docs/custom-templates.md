# Custom Templates

To customize the HTML pages rendered by Gangway, you may provide a set of custom templates to use instead of the built-in ones.

:exclamation: **Important: The data passed to the templates might change between versions, and we do not guarantee that we will maintain backwards compatibility. If using custom templates, extra care must be taken when upgrading Gangway.**

To enable this feature, set the `customHTMLTemplatesDir` option in Gangway's configuration file to a directory that contains the following custom templates:

* home.tmpl: Home page template.
* commandline.tmpl: Post-login template that typically lists the commands needed to configure `kubectl`.

The templates are processed using Go's `html/template` [package][0].

[0]: https://golang.org/pkg/html/template/
