# Go Web Application

This is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.

## Running the server

To run the server, execute the following command:

```bash
go run main.go
```

The server will start on port 8080. You can access it by navigating to `http://localhost:8080/courses` in your web browser.

## Looks like this

![Website](static/images/golang-website.png)

---------------------------------------------------------------------------

## What I Actually Learned by Building This Kubernetes Setup

This wasn’t just “deploying a Go app”.
It was a full end-to-end platform exercise.

---

### 1 I built real Kubernetes networking intuition

I didn’t just “make it work” — I learned *why* it works.

I now understand the full traffic chain:

Browser
- NodePort
- Ingress Controller
- Ingress rule (host/path)
- Service (ClusterIP)
- Pod (container:8080)

This is the core mental model of Kubernetes networking.

---

### 2 I worked with multiple Ingress Controllers in one cluster

This is **not beginner-level** Kubernetes.

Ingress controllers used:
- Traefik (k3s default)
- NGINX Ingress (industry standard)

Key learnings:
- `ingressClassName` decides which controller owns an Ingress
- Multiple ingress controllers can coexist
- One Ingress resource is always handled by exactly one controller

This mirrors real production clusters.

---

### 3 I handled WSL ↔ Kubernetes ↔ Windows networking

This was one of the hardest and most valuable parts.

I learned:
- Why NodePort isn’t reachable on `localhost` by default
- The difference between:
  - WSL VM IP
  - Windows loopback
- Why `nip.io` works
- Why browser behavior can differ from `curl`

These are real-world networking issues that stop many people.

---

### 4 I diagnosed image pull failures like a real SRE

I debugged and fixed:
- `ErrImagePull`
- `ImagePullBackOff`
- TLS handshake timeouts
- `401 UNAUTHORIZED` errors

How I fixed them:
- Switching container registries
- Pre-pulling images using `k3s ctr`
- Understanding kubelet vs user-space networking

This is production incident troubleshooting.

---

### 5 I learned that Kubernetes ≠ Docker

This was a huge conceptual shift.

Key realizations:
- Docker builds images
- `containerd` runs them
- Kubernetes does not care if Docker images are deleted locally
- Runtime image cache ≠ build-time tools

Many Docker users never fully understand this separation.

---

### 6 I created a local AWS-like environment

Without using AWS, I effectively simulated:

| Local setup        | AWS equivalent        |
|-------------------|----------------------|
| k3s               | EKS                  |
| NodePort          | LoadBalancer         |
| Traefik / NGINX   | ALB / NGINX Ingress  |
| nip.io            | Route53              |
| containerd        | EKS runtime          |

So when I eventually use AWS, nothing will feel “magical”.

---

### 7 I practiced safe cleanup and coexistence

I didn’t blindly delete things.

I:
- Kept both ingress controllers while testing
- Removed only outdated ingress objects
- Verified cluster state after every change

This is real cluster hygiene.

---

## Additional Skills I Gained

---

### 8 I learned tool boundaries

I learned where each tool belongs:
- `k3s kubectl` vs `kubectl`
- `helm` vs `k3s`
- Chart root vs `templates/`

I stopped mixing responsibilities — a key operational skill.

---

### 9 I internalized filesystem & execution context

I dealt with:
- Running Helm from the wrong directory
- Relative paths (`.` vs `./chart-name`)
- WSL filesystem (`/mnt/d`) vs native Linux paths

I learned that many “Kubernetes problems” are actually environment problems.

---

### 10 I practiced incremental debugging

My debugging pattern was consistent:
1. Inspect current state (`kubectl get ...`)
2. Change one thing
3. Re-verify

This is exactly how production clusters are debugged.

---

### 11 I learned why Helm exists

I understood:
- Why raw YAML doesn’t scale
- Why `kubectl apply` becomes painful
- Why `helm upgrade --install` is the standard workflow

This shifted my mindset to **release management**.

---

### 12 I learned Kubernetes failure modes

I worked through:
- `<pending>` LoadBalancers
- Broken `kubectl` binaries
- Kubeconfig permission issues
- WSL networking edge cases

Understood the difference between tutorials and real operations.

---

### 13 I learned to respect defaults before overriding them

I didn’t immediately remove Traefik.

I:
- Observed defaults
- Understood why they exist
- Chose when NGINX made more sense
- Adapted the cluster instead of fighting it

Platform thinking.

---

### 14 I built transferable, production-ready skills

Nothing here is WSL-specific.

Everything transfers to:
- EKS
- GKE
- AKS
- On-prem Kubernetes
- Production ingress setups

---

## Final takeaway

I didn’t just deploy an application.

I proved that I can:
- Package software
- Ship versions
- Run applications on Kubernetes
- Expose services safely
- Debug failures under pressure
- Reason about the system end-to-end

---

## Continuous Integration (CI): From Code to Container

### 15 I built a real CI pipeline for a Go application

I implemented a CI pipeline that runs on every push.

The pipeline:
- Runs unit tests (`go test ./...`)
- Enforces code quality with `golangci-lint`
- Builds a container image
- Pushes the image to a container registry

This ensured that **only tested and linted code** can be shipped.

---

### 16 I learned CI toolchain compatibility the hard way

I encountered and fixed:
- Go version vs `golangci-lint` binary incompatibilities
- Breaking config changes between `golangci-lint` v1 and v2
- CI failures caused by invalid or outdated lint configuration

Key learnings:
- CI tools must match the language runtime version
- Pinning versions matters
- Lint configuration is **part of the codebase**, not an afterthought

---

### 17 I implemented fail-fast quality gates

My CI pipeline fails immediately when:
- Tests fail
- Linting rules are violated

This prevents broken or low-quality code from ever reaching Kubernetes.

---

### 18 I prepared the project for GitOps-based CD

The CI pipeline produces **immutable container images**.

Instead of deploying directly, CI:
- Publishes a versioned image
- Leaves deployment decisions to the CD layer (Argo CD)

This cleanly separates:
- **CI responsibilities** (build & verify)
- **CD responsibilities** (deploy & reconcile)

---

### CI Summary

By implementing CI before CD, I ensured that:
- Kubernetes only runs **validated artifacts**
- Debugging is easier because failures are caught early
- The system scales naturally into GitOps workflows

---

## Continuous Delivery (CD): GitOps with Argo CD

### 19 I implemented GitOps-based Continuous Delivery

Instead of deploying directly from CI, I implemented **GitOps-style CD** using Argo CD.

In this model:
- Git is the **single source of truth**
- Kubernetes state is reconciled continuously
- Deployments happen automatically when Git changes

This mirrors how real production platforms operate.

---

### 20 I separated CI and CD responsibilities clearly

I learned that:
- CI is responsible for **building and validating artifacts**
- CD is responsible for **deploying and reconciling state**

My pipeline now follows this flow:

Code push
→ CI (tests, lint, build image)
→ Image pushed to registry
→ Git repository updated (Helm values / manifests)
→ Argo CD detects change
→ Kubernetes reconciles desired state


This separation reduced coupling and improved reliability.

---

### 21 I deployed applications declaratively using Argo CD

I used Argo CD to:
- Deploy Helm charts declaratively
- Continuously monitor application health
- Automatically recover from drift

If a deployment was modified manually:
- Argo CD detected the difference
- Reconciled the cluster back to the declared state

This enforced **platform-level consistency**.

---

### 22 I debugged real-world CD failures

During CD setup, I debugged:
- Application sync failures
- Unhealthy deployments caused by missing dependencies
- Misconfigured ingress and service definitions
- Environment-specific issues in local clusters (k3s + WSL)

I learned how to:
- Read Argo CD application status meaningfully
- Distinguish between **application errors** and **platform errors**
- Fix issues without bypassing GitOps principles

---

### 23 I validated production-style delivery locally

Even on a local k3s cluster, I achieved:
- Automated deployments via Git commits
- Controlled rollouts through declarative config
- Full visibility into application state and health

This proved that **GitOps is not cloud-dependent** — it is a mindset and workflow.

---

### CD Summary

By implementing GitOps-based CD, I demonstrated that I can:
- Operate Kubernetes declaratively
- Trust Git as the deployment interface
- Debug delivery pipelines end-to-end
- Run Argo CD as a real platform component, not a demo tool
