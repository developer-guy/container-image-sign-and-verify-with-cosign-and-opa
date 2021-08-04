module github.com/developer-guy/container-image-sign-and-verify-with-cosign-and-opa

go 1.16

require (
	github.com/google/go-containerregistry v0.6.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/sigstore/cosign v1.0.1
)

replace github.com/prometheus/common => github.com/prometheus/common v0.26.0
