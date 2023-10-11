# admission-controller-scaffold

Source: https://developers.redhat.com/articles/2021/09/17/kubernetes-admission-control-validating-webhooks#bootstrap_with_the_operator_sdk

![Kubernetes Admission controllers flow diagram](img/Kubernetes-Admission-controllers-01-flow-diagram.jpg "Kubernetes Admission controllers flow diagram")

### Generate controller scaffold
```bash
$ operator-sdk init --domain cfy.cz --repo github.com/Cloud-for-You/admission-controller-scaffold
```

### Add new admission webhook
- Configure router / Add path (routes can be configured in the *main.go* file)
- Add file to directory admissions

#### Example for route
```go
hookServer.Register("/sample", &webhook.Admission{Handler: &webhookadmission.Sample{Client: mgr.GetClient()}})
```

#### Examples for admission file
```go
package admission

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type Sample struct {
	Client  client.Client
	decoder *admission.Decoder
}

func (v *Sample) Handle(ctx context.Context, req admission.Request) admission.Response {
	// If DENY use:
	// return admission.Denied(fmt.Sprintf("Missing one or more of minimum required parameters. severity: %v, example_response_code: %v, example_content: %v", "DEBUG", "Response Code", "Content"))
	// If ALLOW use:
	//return admission.Allowed("Rule admitted by Generic validating webhook.")
}

func (v *Sample) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}

```