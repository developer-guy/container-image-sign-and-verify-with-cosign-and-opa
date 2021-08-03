package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/julienschmidt/httprouter"
	"github.com/sigstore/cosign/cmd/cosign/cli"
	"github.com/sigstore/cosign/cmd/cosign/cli/fulcio"
	"github.com/sigstore/cosign/pkg/cosign"
)

type VerificationReq struct {
	Image string
}

type VerificationResp struct {
	Verified            bool   `json:"verified"`
	VerificationMessage string `json:"verification_message"`
}

func VerifySig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var body VerificationReq
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.TODO()

	wDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key, err := cosign.LoadPublicKey(ctx, filepath.Join(wDir, "cosign.pub"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ref, err := name.ParseReference(body.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	co := &cosign.CheckOpts{
		RootCerts:   fulcio.Roots,
		SigVerifier: key,
		Claims:      true,
	}

	var resp VerificationResp
	if _, err = cosign.Verify(ctx, ref, co); err != nil {
		resp = VerificationResp{
			Verified:            false,
			VerificationMessage: err.Error(),
		}
	} else {
		resp = VerificationResp{
			Verified:            true,
			VerificationMessage: fmt.Sprintf("valid signatures found for an image: %s", body.Image),
		}
	}

	respAsByte, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respAsByte)
}
func VerifyAttest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var body VerificationReq
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.TODO()

	wDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	verifier, err := cli.LoadPublicKey(ctx, filepath.Join(wDir, "cosign.pub"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ref, err := name.ParseReference(body.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	co := &cosign.CheckOpts{
		RootCerts:            fulcio.Roots,
		RegistryClientOpts:   cli.DefaultRegistryClientOpts(ctx),
	}
	co.SigVerifier = verifier
	var resp VerificationResp
	if _, err = cosign.Verify(ctx, ref, co); err != nil {
		resp = VerificationResp{
			Verified:            false,
			VerificationMessage: err.Error(),
		}
	} else {
		resp = VerificationResp{
			Verified:            true,
			VerificationMessage: fmt.Sprintf("valid signatures found for an image: %s", body.Image),
		}
	}

	respAsByte, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respAsByte)
}

func main() {
	router := httprouter.New()
	router.POST("/verify-signature", VerifySig)
	router.POST("/verify-attestation", VerifyAttest)
	log.Fatal(http.ListenAndServe(":8080", router))
}
