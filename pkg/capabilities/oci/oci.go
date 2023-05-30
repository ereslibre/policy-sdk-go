package oci

import (
	"fmt"

	cap "github.com/kubewarden/policy-sdk-go/pkg/capabilities"

	"github.com/mailru/easyjson"
)

type HostOCIVerifyVersion int64

const (
	V1 HostOCIVerifyVersion = iota
	V2
)

func (s HostOCIVerifyVersion) String() string {
	switch s {
	case V1:
		return "v1/verify"
	case V2:
		return "v2/verify"
	}
	return "unknown"
}

func Verify(h *cap.Host, requestObj easyjson.Marshaler, operation HostOCIVerifyVersion) (VerificationResponse, error) {
	// failsafe return response
	vr := VerificationResponse{
		IsTrusted: false,
		Digest:    "",
	}

	payload, err := easyjson.Marshal(requestObj)
	if err != nil {
		return vr, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "oci", operation.String(), payload)
	if err != nil {
		return vr, err
	}

	responseObj := VerificationResponse{}
	if err := easyjson.Unmarshal(responsePayload, &responseObj); err != nil {
		return vr, fmt.Errorf("cannot unmarshall response object: %w", err)
	}

	return responseObj, nil
}