package common

import (
	"encoding/json"
	"fmt"
	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	manifestName = "manifest"
	manifestKey  = "createdResources"
)

type Manifest struct {
	ref     *v1.ConfigMap
	content ManifestContent
}

type ManifestItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
	UUID string `json:"uuid"`
}

type ManifestContent struct {
	Items []ManifestItem `json:"items"`
}

func NewManifest(cr *kc.Keycloak) *Manifest {
	manifestConfigMap := &v1.ConfigMap{
		ObjectMeta: v12.ObjectMeta{
			Name:      manifestName,
			Namespace: cr.Namespace,
		},
		Data: map[string]string{
			manifestKey: "",
		},
	}

	return &Manifest{
		ref: manifestConfigMap,
		content: ManifestContent{
			Items: []ManifestItem{},
		},
	}
}

func (i *Manifest) Add(resourceType string, resource runtime.Object) {
	obj := resource.(v12.Object)

	item := ManifestItem{
		Name: obj.GetName(),
		Type: resourceType,
		UUID: fmt.Sprintf("%v", obj.GetUID()),
	}

	i.content.Items = append(i.content.Items, item)
}

func (i *Manifest) Commit() {
	str, err := json.Marshal(i.content)
	if err != nil {
		i.ref.Data[manifestKey] = string(str)
	}
}
