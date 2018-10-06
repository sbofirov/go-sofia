# go-sofia
# Workshop materials

This document: https://kubernetes.grahovac.pro

Elena Grahovac

- hello@grahovac.pro
- https://twitter.com/webdeva
- https://github.com/rumyantseva

## Agenda

- Morning coffee & Registration  - 9:30 - 10:00
- Introduction & Preparation     - 10:00 - 11:00
- Coffee break                   - 11:00 - 11:15
- Writing Go from scratch        - 11:15 - 12:15
- Lunch (Provided at the venue)  - 12:15 - 13:00
- Getting ready for production   - 13:00 - 14:45
- Coffee Break                   - 14:45 - 15:00
- Ship it!                       - 15:00 - 16:45
- Q&A + Optional topics          - 16:45 - 17:30

## Checklist

- Go: https://golang.org/dl/
- IDE or editor to write code
- Docker CE: https://www.docker.com/community-edition#/download
- If Windows: Cygwin - https://cygwin.com/install.html
- GitHub Account: https://github.com
- Git client

## Cleanup notes

- The cluster will be available for a couple of weeks.
- On October, 20th the cluster and all the data will be deleted.
- If you want us to delete your data earlier, please let us know.

## Additional slides
- http://gowayfest.grahovac.pro

## Example service

https://github.com/rumyantseva/go-sofia

## Check your Kubernetes-readines

Follow the instruction: https://ui.k8s.community
Kubernetes dashboard: https://dash.k8s.community

```
USER=your_github_user
kubectl run hello-app --image=gcr.io/google-samples/hello-app:1.0 --port=8080 -n ${USER}
kubectl expose deployment hello-app -n ${USER}
```

Ingress configuration (example):

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: hello-app-ingress
  namespace: rumyantseva
  annotations:
    kubernetes.io/ingress.class: nginx
    nginx.ingress.kubernetes.io/rewrite-target: "/"
spec:
  rules:
  - host: services.k8s.community
    http:
      paths:
      - path: /rumyantseva/hello
        backend:
          serviceName: hello-app
          servicePort: 8080
  tls:
  - hosts:
    - services.k8s.community
    secretName: tls-secret
```

```
kubectl apply -f ingress.yaml
```


## Build a Docker image and run container


```
docker build -t go-sofia .
docker run -p 8080:8080 -p 8585:8585 -t go-sofia
```


## MVP to complete the workshop

- Set up https://github.com/apps/k8s for your service
- Provide helm configuration and Makefile (you can use https://github.com/rumyantseva/go-sofia as an example)
- Prepare a new "release" branch (e.g. release/0.0.1, release/0.0.2 etc)
- Push changes to the release branch
- You changes should be triggered by the CI and you will see the message next to the commit as soon as the CI process completed
