package admission

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	logger = ctrl.Log.WithName("sample-admission")
)

type Sample struct {
	Client  client.Client
	decoder *admission.Decoder
}

func (v *Sample) Handle(ctx context.Context, req admission.Request) admission.Response {
	logger.Info("Handling admission webhook ...")

	// Return func for DENY
	return admission.Denied(fmt.Sprintf("Missing one or more of minimum required parameters. severity: %v, example_response_code: %v, example_content: %v", "DEBUG", "Response Code", "Content"))
	// Return func for ALLOW
	//return admission.Allowed("Rule admitted by Generic validating webhook.")
}

func (v *Sample) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
