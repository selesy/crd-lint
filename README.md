# crd-lint

A Kubernetes tool that verifies that manifests for custom resource
definitions (CRDs) are valid according to the CRD's validation rules.  It
is also capable of performing a variety of "policy" checks.  By default
CRD validation rules are pulled from the current server.  Alternately,
a directory containing the CRDs may be specified for completely off-line
use.

## Usage

For normal operation against the current server, simply provide a list of
manifests to be checked:

```shell
crd-lint manifest-1.yaml manifest-2.yaml manifest-3.yaml
```

For off-line operation, also pass the path to a directory containing the
CRDs.  Note that for off-line mode no ``.kube/config`` file is necessary:

```shell
crd-lint -crds ./crds manifest-1.yaml manifest-2.yaml manifest-3.yaml
```

By default, ``crd-lint`` uses the user's ``.kube/config`` file to locate
the current server.  An absolute pate to an alternate location can be
specified using the ``-kubeconfig`` flag as follows:

```shell
crd-lint -kubeconfig /tmp/.kube/config manifest-1.yaml manifest-2.yaml
```

For more detailed usage information and a complete list of the available
command-line flags for ``crd-lint`` see the ``crd-list`` help.

```shell
crd-lint -help
```

## Policies

Currently implemented policies:

- The specifed CRD version must be loaded on the server (except off-line
  mode.)
