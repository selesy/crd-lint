package internal

import apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

type CRDMap map[string]apiextensionsv1beta1.CustomResourceDefinition
