
# How get tls.crt and tls.key
```bash
$ oc get secret -n local-testing-operator-sdk generic-kubernetes-webhook -o jsonpath="{.data.tls\.crt}" | base64 -d > /var/folders/pc/4wlrj1112b77159kcy1n138r0000gn/T/k8s-webhook-server/serving-certs/tls.crt
$ oc get secret -n local-testing-operator-sdk generic-kubernetes-webhook -o jsonpath="{.data.tls\.key}" | base64 -d > /var/folders/pc/4wlrj1112b77159kcy1n138r0000gn/T/k8s-webhook-server/serving-certs/tls.key
```

# How to get ca.crt for Webhook caBundle

```bash
$ oc get secret -n openshift-service-ca signing-key -o jsonpath="{.data.tls\.crt}"
```