# Example Manifests

You can run your function locally and test it using `crossplane render`
with these example manifests.

```shell
# Run the function locally
$ go run . --insecure --debug
```

```shell
# Then, in another terminal, call it with these example manifests
$ crossplane render xr.yaml composition.yaml functions.yaml -r
---
apiVersion: example.crossplane.io/v1beta1
kind: XR
metadata:
  name: example
status:
  dummy: cool-status
---
apiVersion: iam.aws.upbound.io/v1beta1
kind: User
metadata:
  annotations:
    crossplane.io/composition-resource-name: test-user-0
  generateName: example-
  labels:
    crossplane.io/composite: example
    dummy: foo
    testing.upbound.io/example-name: test-user-0
  name: test-user-0
  ownerReferences:
  - apiVersion: example.crossplane.io/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: XR
    name: example
    uid: ""
spec:
  forProvider: {}
---
apiVersion: iam.aws.upbound.io/v1beta1
kind: AccessKey
metadata:
  annotations:
    crossplane.io/composition-resource-name: sample-access-key-0
  generateName: example-
  labels:
    crossplane.io/composite: example
  name: sample-access-key-0
  ownerReferences:
  - apiVersion: example.crossplane.io/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: XR
    name: example
    uid: ""
spec:
  forProvider:
    userSelector:
      matchLabels:
        testing.upbound.io/example-name: test-user-0
  writeConnectionSecretToRef:
    name: sample-access-key-secret-0
    namespace: crossplane-system
---
dcds:
  sample-access-key-0:
    Ready: "True"
    Resource:
      apiVersion: iam.aws.upbound.io/v1beta1
      kind: AccessKey
      metadata:
        annotations: {}
        name: sample-access-key-0
      spec:
        forProvider:
          userSelector:
            matchLabels:
              testing.upbound.io/example-name: test-user-0
        writeConnectionSecretToRef:
          name: sample-access-key-secret-0
          namespace: crossplane-system
  test-user-0:
    Ready: "False"
    Resource:
      apiVersion: iam.aws.upbound.io/v1beta1
      kind: User
      metadata:
        annotations: {}
        labels:
          dummy: foo
          testing.upbound.io/example-name: test-user-0
        name: test-user-0
      spec:
        forProvider: {}
metadata:
  annotations:
    crossplane.io/composition-resource-name: dcds
  generateName: example-
  labels:
    crossplane.io/composite: example
  name: dcds
  ownerReferences:
  - apiVersion: example.crossplane.io/v1beta1
    blockOwnerDeletion: true
    controller: true
    kind: XR
    name: example
    uid: ""
```
