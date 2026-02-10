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

This is **junior DevOps / platform engineer-level experience**.
