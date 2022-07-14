// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeleteSnapshotInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the EBS snapshot.
	//
	// SnapshotId is a required field
	SnapshotId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSnapshotInput"}

	if s.SnapshotId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSnapshotOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSnapshot = "DeleteSnapshot"

// DeleteSnapshotRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified snapshot.
//
// When you make periodic snapshots of a volume, the snapshots are incremental,
// and only the blocks on the device that have changed since your last snapshot
// are saved in the new snapshot. When you delete a snapshot, only the data
// not needed for any other snapshot is removed. So regardless of which prior
// snapshots have been deleted, all active snapshots will have access to all
// the information needed to restore the volume.
//
// You cannot delete a snapshot of the root device of an EBS volume used by
// a registered AMI. You must first de-register the AMI before you can delete
// the snapshot.
//
// For more information, see Deleting an Amazon EBS Snapshot (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-deleting-snapshot.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DeleteSnapshotRequest.
//    req := client.DeleteSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteSnapshot
func (c *Client) DeleteSnapshotRequest(input *DeleteSnapshotInput) DeleteSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteSnapshotOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteSnapshotRequest{Request: req, Input: input, Copy: c.DeleteSnapshotRequest}
}

// DeleteSnapshotRequest is the request type for the
// DeleteSnapshot API operation.
type DeleteSnapshotRequest struct {
	*aws.Request
	Input *DeleteSnapshotInput
	Copy  func(*DeleteSnapshotInput) DeleteSnapshotRequest
}

// Send marshals and sends the DeleteSnapshot API request.
func (r DeleteSnapshotRequest) Send(ctx context.Context) (*DeleteSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSnapshotResponse{
		DeleteSnapshotOutput: r.Request.Data.(*DeleteSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSnapshotResponse is the response type for the
// DeleteSnapshot API operation.
type DeleteSnapshotResponse struct {
	*DeleteSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSnapshot request.
func (r *DeleteSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
