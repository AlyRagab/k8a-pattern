# k8a-pattern
Generating a Random Numbers API and check its Predictable Compute Usage

### Usage:
```
kubectl port-forward $POD_NAME 8080:8080 
curl http://localhost:8080
```
### The output will be :
```
Generated Number is : 0.5710077816816121
Pattern : Predictable Demands
```

### the logs should be as below:
```
Generated Number        = 0.8370161192974027 
Totall Heap Memory      = 63 MB
Used Memory             = 1007616 Byte
Free Memory             = 62 MB
Processors              = 2
```

### Let's try to change the limits and requests and see the logs again :
```
patch=$(cat <<EOT
[
  {
    "op": "replace",
    "path": "/spec/template/spec/containers/0/resources/requests/memory",
    "value": "300Mi"
  },
  {
    "op": "replace",
    "path": "/spec/template/spec/containers/0/resources/limits/memory",
    "value": "300Mi"
  }
]
EOT
)
kubectl patch deploy random-generator --type=json -p $patch
```
