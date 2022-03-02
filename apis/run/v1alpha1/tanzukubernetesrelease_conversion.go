package v1alpha1

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *TanzuKubernetesRelease) ConvertTo(dest conversion.Hub) error {
	return nil
}

func (dest *TanzuKubernetesRelease) ConvertFrom(src conversion.Hub) error {
	return nil
}
