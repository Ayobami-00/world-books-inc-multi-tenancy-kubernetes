# Implementing Multi-Tenancy in Kubernetes

Welcome to the repository for the article (currently in draft): "Implementing Multi-Tenancy in Kubernetes." This repository contains code examples, resources, and insights discussed in the article.

## Introduction

As a Devops Professional, consider that you have been offered a project by WorldBooks Inc (The company name used in this article is entirely fictitious and has been deliberately chosen for creative and illustrative purposes. Any similarity to any actual company name is not intentional and should be construed as coincidental) to convert their book stores inventory management across multiple locations around the world from a manual inventory solution to a cloud solution in which every location around the world accesses a similar platform but sees only books stored and managed in their location. 

Now there might be plenty of ways to solve this problem, some may be less complex or more straightforward howver for the purpose of this project we are going to be looking at Multitenancy in kubernetes and how it can be applied to solve our current problem.

## Concepts

The following section gives a summary of the concepts used in the article and links that can take you to their usage in the repo.

1. [Kubernetes](https://kubernetes.io/)
2. [REST API NAMING CONVENTIONS](https://restfulapi.net/resource-naming/)
3. [gRPC](https://grpc.io/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/main/backend-service-go/pb/book_grpc.pb.go)
4. [GitHub Actions](https://github.com/features/actions). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/.github/workflows/ci.yml)
5. [Amazon Elastic Container Registry](https://aws.amazon.com/ecr/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/.github/workflows/ci.yml)
6. [Amazon Elastic Kubernetes Service](https://aws.amazon.com/eks/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/tree/deploy/deploy/eks)
7. [Kubernetes Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/workload.yaml)
8. [Kubernetes Service](https://kubernetes.io/docs/concepts/services-networking/service/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/service.yaml)
9. [Kubernetes Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/ingress.yaml)
10. [DRY (Don’t repeat yourself)](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself).
11. [Helm Charts](https://helm.sh/docs/topics/charts/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/tree/deploy/deploy/eks/helm-chart)
12. [Storage in Kubernetes](https://kubernetes.io/docs/concepts/storage/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/storage.yaml)
13. [Kubernets Namespaces](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/.github/workflows/ci.yml)
14. [Network Policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/network-policy.yaml)
15. [Resource Quotas](https://kubernetes.io/docs/concepts/policy/resource-quotas/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/resource-quota.yaml)
16. [LimitRanges](https://kubernetes.io/docs/concepts/policy/limit-range/). [See usage](https://github.com/Ayobami-00/world-books-inc-multi-tenancy-kubernetes/blob/deploy/deploy/eks/helm-chart/templates/limit-range.yaml)
17. [Soft and Hard Multi Tenancy](https://www.alibabacloud.com/blog/practices-of-kubernetes-multi-tenant-clusters_596178)
18. Cluster-as-a-service tools:  [Cluster API](https://github.com/kubernetes-sigs/cluster-api),  [HyperShift](https://github.com/openshift/hypershift) e.t.c

## Prerequisites

- Familiarity with Kubernetes concepts.
- Understanding of REST API conventions, gRPC, etc.
- The Go language installed on your local machine. See https://go.dev/ for details on how to install based the OS of your local machine.
- A running kubernetes cluster (we used Amazon Elastic Kubernetes Service in this article). See https://www.vmware.com/topics/glossary/content/kubernetes-cluster.html for details.

## Code Examples

This repository contains code examples used in the presentation. You can find them under this directory.

You can also follow me on [github](https://github.com/Ayobami-00) and [twitter](https://twitter.com/_enfinity). Looking forward to connecting and hearing how this article has helped you.

## Contributing

Contributions are welcome! If you have improvements, suggestions, or want to add relevant resources, feel free to submit a pull request.

## License

This project is licensed under the [MIT License](/LICENSE).
