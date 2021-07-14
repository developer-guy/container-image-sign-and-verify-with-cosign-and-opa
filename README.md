# Sign Container Images with cosign and Verify signature by using Open Policy Agent (OPA)

In the beginning, I believe it is worth saying that this project is just a proof-of-concept project that shows people how they can use cosign and OPA (Open Policy Agent) together to implement the signing and verifying container image process together.

In most basic form, [cosign](https://github.com/sigstore/cosign) is a container signing tool; it helps us to sign and verify container images by using the signature algorithm (ECDSA-P256) and payload format ([Red Hat Simple Signing](https://www.redhat.com/en/blog/container-image-signing)).

[Dan Lorenc](https://twitter.com/lorenc_dan), who is the maintainer of the project, wrote an excellent article about what cosign is and the motivation behind it; you can follow the [link](https://blog.sigstore.dev/cosign-signed-container-images-c1016862618) to access it.

On the other hand side, the Open Policy Agent (OPA, pronounced "oh-pa") is an open-source, general-purpose policy engine that unifies policy enforcement across the stack. So, the motivation behind using this kind of policy engine is providing an easy way of enforcing organizational policies across the stack.

## What is the motivation for combining both cosign and OPA?

Let's assume that we have to ensure only the images that have valid signatures can be deployed into production-grade Kubernetes clusters. So, to implement this kind of scenario is that we can use OPA's [http.send](https://www.openpolicyagent.org/docs/latest/policy-reference/#http) built-in function to call some external service an HTTP server that exposes `/verify` endpoint and uses `cosign` under the hood to verify the signature of an image.