package signature

default allow_attestation = false

allow_attestation {
    # send HTTP POST request to cosign-wrapper
    body := {
    	"image": input.image,
    }
    headers_json := {"Content-Type": "application/json"}
    cosignHTTPWrapperURL := "http://localhost:8080/verify-attestation"
    output := http.send({"method": "post", "url": cosignHTTPWrapperURL, "headers": headers_json, "body": body})
    p := json.unmarshal(base64.decode(output.body.payload))
    contains(p.predicate.Data, "{ \"id\": \"internal ci/cd platform\"}")
}
