module github.com/developer-guy/container-image-sign-and-verify-with-cosign-and-opa

go 1.16

require (
	github.com/docker/cli v20.10.0-beta1.0.20201117192004-5cc239616494+incompatible // indirect
	github.com/docker/docker v20.10.0-beta1.0.20201110211921-af34b94a78a1+incompatible // indirect
	github.com/google/go-containerregistry v0.5.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/sigstore/cosign v0.6.1-0.20210713005353-82d49dcf3b8b
	github.com/sigstore/rekor v0.2.1-0.20210712122031-1c30d2ff9518 // indirect
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gotest.tools/v3 v3.0.3 // indirect
)

replace github.com/prometheus/common => github.com/prometheus/common v0.26.0
