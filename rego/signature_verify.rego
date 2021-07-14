package signature

default verified = false

verified {
    # send HTTP POST request to cosign-wrapper
    body := {
    	"image": input.image,
    }
    headers_json := {"Content-Type": "application/json"}
    cosignHTTPWrapperURL := "http://localhost:8080/verify"
    output := http.send({"method": "post", "url": cosignHTTPWrapperURL, "headers": headers_json, "body": body})
    output.body.verified
}