// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateIamInstanceProfileInput struct {
	_ struct{} `type:"structure"`

	// The ID of the IAM instance profile association.
	//
	// AssociationId is a required field
	AssociationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateIamInstanceProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateIamInstanceProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateIamInstanceProfileInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateIamInstanceProfileOutput struct {
	_ struct{} `type:"structure"`

	// Information about the IAM instance profile association.
	IamInstanceProfileAssociation *IamInstanceProfileAssociation `locationName:"iamInstanceProfileAssociation" type:"structure"`
}

// String returns the string representation
func (s DisassociateIamInstanceProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDisassociateIamInstanceProfile = "DisassociateIamInstanceProfile"

// DisassociateIamInstanceProfileRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates an IAM instance profile from a running or stopped instance.
//
// Use DescribeIamInstanceProfileAssociations to get the association ID.
//
//    // Example sending a request using DisassociateIamInstanceProfileRequest.
//    req := client.DisassociateIamInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateIamInstanceProfile
func (c *Client) DisassociateIamInstanceProfileRequest(input *DisassociateIamInstanceProfileInput) DisassociateIamInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opDisassociateIamInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateIamInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &DisassociateIamInstanceProfileOutput{})

	return DisassociateIamInstanceProfileRequest{Request: req, Input: input, Copy: c.DisassociateIamInstanceProfileRequest}
}

// DisassociateIamInstanceProfileRequest is the request type for the
// DisassociateIamInstanceProfile API operation.
type DisassociateIamInstanceProfileRequest struct {
	*aws.Request
	Input *DisassociateIamInstanceProfileInput
	Copy  func(*DisassociateIamInstanceProfileInput) DisassociateIamInstanceProfileRequest
}

// Send marshals and sends the DisassociateIamInstanceProfile API request.
func (r DisassociateIamInstanceProfileRequest) Send(ctx context.Context) (*DisassociateIamInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateIamInstanceProfileResponse{
		DisassociateIamInstanceProfileOutput: r.Request.Data.(*DisassociateIamInstanceProfileOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateIamInstanceProfileResponse is the response type for the
// DisassociateIamInstanceProfile API operation.
type DisassociateIamInstanceProfileResponse struct {
	*DisassociateIamInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateIamInstanceProfile request.
func (r *DisassociateIamInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
