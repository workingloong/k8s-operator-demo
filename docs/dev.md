# Tutorial to Develop a Controller

The document describes how to use kubebuilder to customize
a Kubernetes controller and CRD.

## 1. Install kubebuilder

We can install kubebuilder by the [tutorial](https://book.kubebuilder.io/quick-start.html)

## 2. Initialize a Repo

### Define a Custom Resource

In the example, we define a `DemoJob` with workers.

```yaml
apiVersion: demo.example.com/v1
kind: DemoJob
metadata:
  name: DemoJob-sample
  namespace: default
spec:
  worker: 
    count: 1
    resource:
      cpu: 1
      memory: 1024
```

`apiVersion` if the format `{group}.{domain}/{version}`

Then, we use `kubebuilder` to build a scaffolding for `DemoJob`.

```shell
go mod init k8s-operator-demo
kubebuilder init --domain example.com
```

After the scaffolding is finished, we create `api` for `DemoJob`.

```
$ kubebuilder create api --group demo --version v1 --kind DemoJob
Create Resource [y/n]
y
Create Controller [y/n]
y
...
```

Now, we complete the initialization of the project.

