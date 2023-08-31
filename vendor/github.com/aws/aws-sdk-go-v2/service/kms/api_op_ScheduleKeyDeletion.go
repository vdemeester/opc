// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Schedules the deletion of a KMS key. By default, KMS applies a waiting period
// of 30 days, but you can specify a waiting period of 7-30 days. When this
// operation is successful, the key state of the KMS key changes to PendingDeletion
// and the key can't be used in any cryptographic operations. It remains in this
// state for the duration of the waiting period. Before the waiting period ends,
// you can use CancelKeyDeletion to cancel the deletion of the KMS key. After the
// waiting period ends, KMS deletes the KMS key, its key material, and all KMS data
// associated with it, including all aliases that refer to it. Deleting a KMS key
// is a destructive and potentially dangerous operation. When a KMS key is deleted,
// all data that was encrypted under the KMS key is unrecoverable. (The only
// exception is a multi-Region replica key , or an asymmetric or HMAC KMS key with
// imported key material .) To prevent the use of a KMS key without deleting it,
// use DisableKey . You can schedule the deletion of a multi-Region primary key and
// its replica keys at any time. However, KMS will not delete a multi-Region
// primary key with existing replica keys. If you schedule the deletion of a
// primary key with replicas, its key state changes to PendingReplicaDeletion and
// it cannot be replicated or used in cryptographic operations. This status can
// continue indefinitely. When the last of its replicas keys is deleted (not just
// scheduled), the key state of the primary key changes to PendingDeletion and its
// waiting period ( PendingWindowInDays ) begins. For details, see Deleting
// multi-Region keys (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html)
// in the Key Management Service Developer Guide. When KMS deletes a KMS key from
// an CloudHSM key store (https://docs.aws.amazon.com/kms/latest/developerguide/delete-cmk-keystore.html)
// , it makes a best effort to delete the associated key material from the
// associated CloudHSM cluster. However, you might need to manually delete the
// orphaned key material (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups. Deleting a KMS key from an external key store (https://docs.aws.amazon.com/kms/latest/developerguide/delete-xks-key.html)
// has no effect on the associated external key. However, for both types of custom
// key stores, deleting a KMS key is destructive and irreversible. You cannot
// decrypt ciphertext encrypted under the KMS key by using only its associated
// external key or CloudHSM key. Also, you cannot recreate a KMS key in an external
// key store by creating a new KMS key with the same key material. For more
// information about scheduling a KMS key for deletion, see Deleting KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html)
// in the Key Management Service Developer Guide. The KMS key that you use for this
// operation must be in a compatible key state. For details, see Key states of KMS
// keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in
// the Key Management Service Developer Guide. Cross-account use: No. You cannot
// perform this operation on a KMS key in a different Amazon Web Services account.
// Required permissions: kms:ScheduleKeyDeletion (key policy) Related operations
//   - CancelKeyDeletion
//   - DisableKey
func (c *Client) ScheduleKeyDeletion(ctx context.Context, params *ScheduleKeyDeletionInput, optFns ...func(*Options)) (*ScheduleKeyDeletionOutput, error) {
	if params == nil {
		params = &ScheduleKeyDeletionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ScheduleKeyDeletion", params, optFns, c.addOperationScheduleKeyDeletionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ScheduleKeyDeletionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ScheduleKeyDeletionInput struct {

	// The unique identifier of the KMS key to delete. Specify the key ID or key ARN
	// of the KMS key. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey .
	//
	// This member is required.
	KeyId *string

	// The waiting period, specified in number of days. After the waiting period ends,
	// KMS deletes the KMS key. If the KMS key is a multi-Region primary key with
	// replica keys, the waiting period begins when the last of its replica keys is
	// deleted. Otherwise, the waiting period begins immediately. This value is
	// optional. If you include a value, it must be between 7 and 30, inclusive. If you
	// do not include a value, it defaults to 30. You can use the
	// kms:ScheduleKeyDeletionPendingWindowInDays (https://docs.aws.amazon.com/kms/latest/developerguide/conditions-kms.html#conditions-kms-schedule-key-deletion-pending-window-in-days)
	// condition key to further constrain the values that principals can specify in the
	// PendingWindowInDays parameter.
	PendingWindowInDays *int32

	noSmithyDocumentSerde
}

type ScheduleKeyDeletionOutput struct {

	// The date and time after which KMS deletes the KMS key. If the KMS key is a
	// multi-Region primary key with replica keys, this field does not appear. The
	// deletion date for the primary key isn't known until its last replica key is
	// deleted.
	DeletionDate *time.Time

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the KMS key whose deletion is scheduled.
	KeyId *string

	// The current status of the KMS key. For more information about how key state
	// affects the use of a KMS key, see Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
	// in the Key Management Service Developer Guide.
	KeyState types.KeyState

	// The waiting period before the KMS key is deleted. If the KMS key is a
	// multi-Region primary key with replicas, the waiting period begins when the last
	// of its replica keys is deleted. Otherwise, the waiting period begins
	// immediately.
	PendingWindowInDays *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationScheduleKeyDeletionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpScheduleKeyDeletion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpScheduleKeyDeletion{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpScheduleKeyDeletionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opScheduleKeyDeletion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opScheduleKeyDeletion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ScheduleKeyDeletion",
	}
}
