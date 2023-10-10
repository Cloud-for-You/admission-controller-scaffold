package admission

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type GenericValidator struct {
	Client client.Client
	decoder *admission.Decoder
}

func (v *GenericValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	return admission.Allowed("Rule admitted by Generic validating webhook.")
}

func (v *GenericValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}