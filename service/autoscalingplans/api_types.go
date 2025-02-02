// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents an application source.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/ApplicationSource
type ApplicationSource struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of a AWS CloudFormation stack.
	CloudFormationStackARN *string `type:"string"`

	// A set of tags (up to 50).
	TagFilters []TagFilter `type:"list"`
}

// String returns the string representation
func (s ApplicationSource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplicationSource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplicationSource"}
	if s.TagFilters != nil {
		for i, v := range s.TagFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a CloudWatch metric of your choosing that can be used for predictive
// scaling.
//
// For predictive scaling to work with a customized load metric specification,
// AWS Auto Scaling needs access to the Sum and Average statistics that CloudWatch
// computes from metric data. Statistics are calculations used to aggregate
// data over specified time periods.
//
// When you choose a load metric, make sure that the required Sum and Average
// statistics for your metric are available in CloudWatch and that they provide
// relevant data for predictive scaling. The Sum statistic must represent the
// total load on the resource, and the Average statistic must represent the
// average load per capacity unit of the resource. For example, there is a metric
// that counts the number of requests processed by your Auto Scaling group.
// If the Sum statistic represents the total request count processed by the
// group, then the Average statistic for the specified metric must represent
// the average request count processed by each instance of the group.
//
// For information about terminology, available metrics, or how to publish new
// metrics, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html)
// in the Amazon CloudWatch User Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/CustomizedLoadMetricSpecification
type CustomizedLoadMetricSpecification struct {
	_ struct{} `type:"structure"`

	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify
	// the same dimensions in your customized load metric specification.
	Dimensions []MetricDimension `type:"list"`

	// The name of the metric.
	//
	// MetricName is a required field
	MetricName *string `type:"string" required:"true"`

	// The namespace of the metric.
	//
	// Namespace is a required field
	Namespace *string `type:"string" required:"true"`

	// The statistic of the metric. Currently, the value must always be Sum.
	//
	// Statistic is a required field
	Statistic MetricStatistic `type:"string" required:"true" enum:"true"`

	// The unit of the metric.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s CustomizedLoadMetricSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CustomizedLoadMetricSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CustomizedLoadMetricSpecification"}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if len(s.Statistic) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Statistic"))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a CloudWatch metric of your choosing that can be used for dynamic
// scaling as part of a target tracking scaling policy.
//
// To create your customized scaling metric specification:
//
//    * Add values for each required parameter from CloudWatch. You can use
//    an existing metric, or a new metric that you create. To use your own metric,
//    you must first publish the metric to CloudWatch. For more information,
//    see Publish Custom Metrics (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)
//    in the Amazon CloudWatch User Guide.
//
//    * Choose a metric that changes proportionally with capacity. The value
//    of the metric should increase or decrease in inverse proportion to the
//    number of capacity units. That is, the value of the metric should decrease
//    when capacity increases.
//
// For more information about CloudWatch, see Amazon CloudWatch Concepts (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/CustomizedScalingMetricSpecification
type CustomizedScalingMetricSpecification struct {
	_ struct{} `type:"structure"`

	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify
	// the same dimensions in your customized scaling metric specification.
	Dimensions []MetricDimension `type:"list"`

	// The name of the metric.
	//
	// MetricName is a required field
	MetricName *string `type:"string" required:"true"`

	// The namespace of the metric.
	//
	// Namespace is a required field
	Namespace *string `type:"string" required:"true"`

	// The statistic of the metric.
	//
	// Statistic is a required field
	Statistic MetricStatistic `type:"string" required:"true" enum:"true"`

	// The unit of the metric.
	Unit *string `type:"string"`
}

// String returns the string representation
func (s CustomizedScalingMetricSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CustomizedScalingMetricSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CustomizedScalingMetricSpecification"}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if len(s.Statistic) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Statistic"))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a single value in the forecast data used for predictive scaling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/Datapoint
type Datapoint struct {
	_ struct{} `type:"structure"`

	// The time stamp for the data point in UTC format.
	Timestamp *time.Time `type:"timestamp"`

	// The value of the data point.
	Value *float64 `type:"double"`
}

// String returns the string representation
func (s Datapoint) String() string {
	return awsutil.Prettify(s)
}

// Represents a dimension for a customized metric.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/MetricDimension
type MetricDimension struct {
	_ struct{} `type:"structure"`

	// The name of the dimension.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The value of the dimension.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s MetricDimension) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricDimension) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricDimension"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a predefined metric that can be used for predictive scaling.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/PredefinedLoadMetricSpecification
type PredefinedLoadMetricSpecification struct {
	_ struct{} `type:"structure"`

	// The metric type.
	//
	// PredefinedLoadMetricType is a required field
	PredefinedLoadMetricType LoadMetricType `type:"string" required:"true" enum:"true"`

	// Identifies the resource associated with the metric type. You can't specify
	// a resource label unless the metric type is ALBRequestCountPerTarget and there
	// is a target group for an Application Load Balancer attached to the Auto Scaling
	// group.
	//
	// The format is app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>,
	// where:
	//
	//    * app/<load-balancer-name>/<load-balancer-id> is the final portion of
	//    the load balancer ARN.
	//
	//    * targetgroup/<target-group-name>/<target-group-id> is the final portion
	//    of the target group ARN.
	ResourceLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PredefinedLoadMetricSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PredefinedLoadMetricSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PredefinedLoadMetricSpecification"}
	if len(s.PredefinedLoadMetricType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PredefinedLoadMetricType"))
	}
	if s.ResourceLabel != nil && len(*s.ResourceLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceLabel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a predefined metric that can be used for dynamic scaling as part
// of a target tracking scaling policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/PredefinedScalingMetricSpecification
type PredefinedScalingMetricSpecification struct {
	_ struct{} `type:"structure"`

	// The metric type. The ALBRequestCountPerTarget metric type applies only to
	// Auto Scaling groups, Spot Fleet requests, and ECS services.
	//
	// PredefinedScalingMetricType is a required field
	PredefinedScalingMetricType ScalingMetricType `type:"string" required:"true" enum:"true"`

	// Identifies the resource associated with the metric type. You can't specify
	// a resource label unless the metric type is ALBRequestCountPerTarget and there
	// is a target group for an Application Load Balancer attached to the Auto Scaling
	// group, Spot Fleet request, or ECS service.
	//
	// The format is app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>,
	// where:
	//
	//    * app/<load-balancer-name>/<load-balancer-id> is the final portion of
	//    the load balancer ARN.
	//
	//    * targetgroup/<target-group-name>/<target-group-id> is the final portion
	//    of the target group ARN.
	ResourceLabel *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PredefinedScalingMetricSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PredefinedScalingMetricSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PredefinedScalingMetricSpecification"}
	if len(s.PredefinedScalingMetricType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PredefinedScalingMetricType"))
	}
	if s.ResourceLabel != nil && len(*s.ResourceLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceLabel", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a scaling instruction for a scalable resource.
//
// The scaling instruction is used in combination with a scaling plan, which
// is a set of instructions for configuring dynamic scaling and predictive scaling
// for the scalable resources in your application. Each scaling instruction
// applies to one resource.
//
// AWS Auto Scaling creates target tracking scaling policies based on the scaling
// instructions. Target tracking scaling policies adjust the capacity of your
// scalable resource as required to maintain resource utilization at the target
// value that you specified.
//
// AWS Auto Scaling also configures predictive scaling for your Amazon EC2 Auto
// Scaling groups using a subset of parameters, including the load metric, the
// scaling metric, the target value for the scaling metric, the predictive scaling
// mode (forecast and scale or forecast only), and the desired behavior when
// the forecast capacity exceeds the maximum capacity of the resource. With
// predictive scaling, AWS Auto Scaling generates forecasts with traffic predictions
// for the two days ahead and schedules scaling actions that proactively add
// and remove resource capacity to match the forecast.
//
// We recommend waiting a minimum of 24 hours after creating an Auto Scaling
// group to configure predictive scaling. At minimum, there must be 24 hours
// of historical data to generate a forecast.
//
// For more information, see Getting Started with AWS Auto Scaling (https://docs.aws.amazon.com/autoscaling/plans/userguide/auto-scaling-getting-started.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/ScalingInstruction
type ScalingInstruction struct {
	_ struct{} `type:"structure"`

	// The customized load metric to use for predictive scaling. This parameter
	// or a PredefinedLoadMetricSpecification is required when configuring predictive
	// scaling, and cannot be used otherwise.
	CustomizedLoadMetricSpecification *CustomizedLoadMetricSpecification `type:"structure"`

	// Controls whether dynamic scaling by AWS Auto Scaling is disabled. When dynamic
	// scaling is enabled, AWS Auto Scaling creates target tracking scaling policies
	// based on the specified target tracking configurations.
	//
	// The default is enabled (false).
	DisableDynamicScaling *bool `type:"boolean"`

	// The maximum capacity of the resource. The exception to this upper limit is
	// if you specify a non-default setting for PredictiveScalingMaxCapacityBehavior.
	//
	// MaxCapacity is a required field
	MaxCapacity *int64 `type:"integer" required:"true"`

	// The minimum capacity of the resource.
	//
	// MinCapacity is a required field
	MinCapacity *int64 `type:"integer" required:"true"`

	// The predefined load metric to use for predictive scaling. This parameter
	// or a CustomizedLoadMetricSpecification is required when configuring predictive
	// scaling, and cannot be used otherwise.
	PredefinedLoadMetricSpecification *PredefinedLoadMetricSpecification `type:"structure"`

	// Defines the behavior that should be applied if the forecast capacity approaches
	// or exceeds the maximum capacity specified for the resource. The default value
	// is SetForecastCapacityToMaxCapacity.
	//
	// The following are possible values:
	//
	//    * SetForecastCapacityToMaxCapacity - AWS Auto Scaling cannot scale resource
	//    capacity higher than the maximum capacity. The maximum capacity is enforced
	//    as a hard limit.
	//
	//    * SetMaxCapacityToForecastCapacity - AWS Auto Scaling may scale resource
	//    capacity higher than the maximum capacity to equal but not exceed forecast
	//    capacity.
	//
	//    * SetMaxCapacityAboveForecastCapacity - AWS Auto Scaling may scale resource
	//    capacity higher than the maximum capacity by a specified buffer value.
	//    The intention is to give the target tracking scaling policy extra capacity
	//    if unexpected traffic occurs.
	//
	// Only valid when configuring predictive scaling.
	PredictiveScalingMaxCapacityBehavior PredictiveScalingMaxCapacityBehavior `type:"string" enum:"true"`

	// The size of the capacity buffer to use when the forecast capacity is close
	// to or exceeds the maximum capacity. The value is specified as a percentage
	// relative to the forecast capacity. For example, if the buffer is 10, this
	// means a 10 percent buffer, such that if the forecast capacity is 50, and
	// the maximum capacity is 40, then the effective maximum capacity is 55.
	//
	// Only valid when configuring predictive scaling. Required if the PredictiveScalingMaxCapacityBehavior
	// is set to SetMaxCapacityAboveForecastCapacity, and cannot be used otherwise.
	//
	// The range is 1-100.
	PredictiveScalingMaxCapacityBuffer *int64 `type:"integer"`

	// The predictive scaling mode. The default value is ForecastAndScale. Otherwise,
	// AWS Auto Scaling forecasts capacity but does not create any scheduled scaling
	// actions based on the capacity forecast.
	PredictiveScalingMode PredictiveScalingMode `type:"string" enum:"true"`

	// The ID of the resource. This string consists of the resource type and unique
	// identifier.
	//
	//    * Auto Scaling group - The resource type is autoScalingGroup and the unique
	//    identifier is the name of the Auto Scaling group. Example: autoScalingGroup/my-asg.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension associated with the resource.
	//
	//    * autoscaling:autoScalingGroup:DesiredCapacity - The desired capacity
	//    of an Auto Scaling group.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// Controls whether a resource's externally created scaling policies are kept
	// or replaced.
	//
	// The default value is KeepExternalPolicies. If the parameter is set to ReplaceExternalPolicies,
	// any scaling policies that are external to AWS Auto Scaling are deleted and
	// new target tracking scaling policies created.
	//
	// Only valid when configuring dynamic scaling.
	//
	// Condition: The number of existing policies to be replaced must be less than
	// or equal to 50. If there are more than 50 policies to be replaced, AWS Auto
	// Scaling keeps all existing policies and does not create new ones.
	ScalingPolicyUpdateBehavior ScalingPolicyUpdateBehavior `type:"string" enum:"true"`

	// The amount of time, in seconds, to buffer the run time of scheduled scaling
	// actions when scaling out. For example, if the forecast says to add capacity
	// at 10:00 AM, and the buffer time is 5 minutes, then the run time of the corresponding
	// scheduled scaling action will be 9:55 AM. The intention is to give resources
	// time to be provisioned. For example, it can take a few minutes to launch
	// an EC2 instance. The actual amount of time required depends on several factors,
	// such as the size of the instance and whether there are startup scripts to
	// complete.
	//
	// The value must be less than the forecast interval duration of 3600 seconds
	// (60 minutes). The default is 300 seconds.
	//
	// Only valid when configuring predictive scaling.
	ScheduledActionBufferTime *int64 `type:"integer"`

	// The namespace of the AWS service.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`

	// The structure that defines new target tracking configurations (up to 10).
	// Each of these structures includes a specific scaling metric and a target
	// value for the metric, along with various parameters to use with dynamic scaling.
	//
	// With predictive scaling and dynamic scaling, the resource scales based on
	// the target tracking configuration that provides the largest capacity for
	// both scale in and scale out.
	//
	// Condition: The scaling metric must be unique across target tracking configurations.
	//
	// TargetTrackingConfigurations is a required field
	TargetTrackingConfigurations []TargetTrackingConfiguration `type:"list" required:"true"`
}

// String returns the string representation
func (s ScalingInstruction) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ScalingInstruction) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ScalingInstruction"}

	if s.MaxCapacity == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxCapacity"))
	}

	if s.MinCapacity == nil {
		invalidParams.Add(aws.NewErrParamRequired("MinCapacity"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if len(s.ScalableDimension) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ScalableDimension"))
	}
	if len(s.ServiceNamespace) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ServiceNamespace"))
	}

	if s.TargetTrackingConfigurations == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetTrackingConfigurations"))
	}
	if s.CustomizedLoadMetricSpecification != nil {
		if err := s.CustomizedLoadMetricSpecification.Validate(); err != nil {
			invalidParams.AddNested("CustomizedLoadMetricSpecification", err.(aws.ErrInvalidParams))
		}
	}
	if s.PredefinedLoadMetricSpecification != nil {
		if err := s.PredefinedLoadMetricSpecification.Validate(); err != nil {
			invalidParams.AddNested("PredefinedLoadMetricSpecification", err.(aws.ErrInvalidParams))
		}
	}
	if s.TargetTrackingConfigurations != nil {
		for i, v := range s.TargetTrackingConfigurations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TargetTrackingConfigurations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents a scaling plan.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/ScalingPlan
type ScalingPlan struct {
	_ struct{} `type:"structure"`

	// The application source.
	//
	// ApplicationSource is a required field
	ApplicationSource *ApplicationSource `type:"structure" required:"true"`

	// The Unix time stamp when the scaling plan was created.
	CreationTime *time.Time `type:"timestamp"`

	// The scaling instructions.
	//
	// ScalingInstructions is a required field
	ScalingInstructions []ScalingInstruction `type:"list" required:"true"`

	// The name of the scaling plan.
	//
	// ScalingPlanName is a required field
	ScalingPlanName *string `min:"1" type:"string" required:"true"`

	// The version number of the scaling plan.
	//
	// ScalingPlanVersion is a required field
	ScalingPlanVersion *int64 `type:"long" required:"true"`

	// The status of the scaling plan.
	//
	//    * Active - The scaling plan is active.
	//
	//    * ActiveWithProblems - The scaling plan is active, but the scaling configuration
	//    for one or more resources could not be applied.
	//
	//    * CreationInProgress - The scaling plan is being created.
	//
	//    * CreationFailed - The scaling plan could not be created.
	//
	//    * DeletionInProgress - The scaling plan is being deleted.
	//
	//    * DeletionFailed - The scaling plan could not be deleted.
	//
	//    * UpdateInProgress - The scaling plan is being updated.
	//
	//    * UpdateFailed - The scaling plan could not be updated.
	//
	// StatusCode is a required field
	StatusCode ScalingPlanStatusCode `type:"string" required:"true" enum:"true"`

	// A simple message about the current status of the scaling plan.
	StatusMessage *string `type:"string"`

	// The Unix time stamp when the scaling plan entered the current status.
	StatusStartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ScalingPlan) String() string {
	return awsutil.Prettify(s)
}

// Represents a scalable resource.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/ScalingPlanResource
type ScalingPlanResource struct {
	_ struct{} `type:"structure"`

	// The ID of the resource. This string consists of the resource type and unique
	// identifier.
	//
	//    * Auto Scaling group - The resource type is autoScalingGroup and the unique
	//    identifier is the name of the Auto Scaling group. Example: autoScalingGroup/my-asg.
	//
	//    * ECS service - The resource type is service and the unique identifier
	//    is the cluster name and service name. Example: service/default/sample-webapp.
	//
	//    * Spot Fleet request - The resource type is spot-fleet-request and the
	//    unique identifier is the Spot Fleet request ID. Example: spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE.
	//
	//    * DynamoDB table - The resource type is table and the unique identifier
	//    is the resource ID. Example: table/my-table.
	//
	//    * DynamoDB global secondary index - The resource type is index and the
	//    unique identifier is the resource ID. Example: table/my-table/index/my-table-index.
	//
	//    * Aurora DB cluster - The resource type is cluster and the unique identifier
	//    is the cluster name. Example: cluster:my-db-cluster.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The scalable dimension for the resource.
	//
	//    * autoscaling:autoScalingGroup:DesiredCapacity - The desired capacity
	//    of an Auto Scaling group.
	//
	//    * ecs:service:DesiredCount - The desired task count of an ECS service.
	//
	//    * ec2:spot-fleet-request:TargetCapacity - The target capacity of a Spot
	//    Fleet request.
	//
	//    * dynamodb:table:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:table:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB table.
	//
	//    * dynamodb:index:ReadCapacityUnits - The provisioned read capacity for
	//    a DynamoDB global secondary index.
	//
	//    * dynamodb:index:WriteCapacityUnits - The provisioned write capacity for
	//    a DynamoDB global secondary index.
	//
	//    * rds:cluster:ReadReplicaCount - The count of Aurora Replicas in an Aurora
	//    DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible
	//    edition.
	//
	// ScalableDimension is a required field
	ScalableDimension ScalableDimension `type:"string" required:"true" enum:"true"`

	// The name of the scaling plan.
	//
	// ScalingPlanName is a required field
	ScalingPlanName *string `min:"1" type:"string" required:"true"`

	// The version number of the scaling plan.
	//
	// ScalingPlanVersion is a required field
	ScalingPlanVersion *int64 `type:"long" required:"true"`

	// The scaling policies.
	ScalingPolicies []ScalingPolicy `type:"list"`

	// The scaling status of the resource.
	//
	//    * Active - The scaling configuration is active.
	//
	//    * Inactive - The scaling configuration is not active because the scaling
	//    plan is being created or the scaling configuration could not be applied.
	//    Check the status message for more information.
	//
	//    * PartiallyActive - The scaling configuration is partially active because
	//    the scaling plan is being created or deleted or the scaling configuration
	//    could not be fully applied. Check the status message for more information.
	//
	// ScalingStatusCode is a required field
	ScalingStatusCode ScalingStatusCode `type:"string" required:"true" enum:"true"`

	// A simple message about the current scaling status of the resource.
	ScalingStatusMessage *string `type:"string"`

	// The namespace of the AWS service.
	//
	// ServiceNamespace is a required field
	ServiceNamespace ServiceNamespace `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ScalingPlanResource) String() string {
	return awsutil.Prettify(s)
}

// Represents a scaling policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/ScalingPolicy
type ScalingPolicy struct {
	_ struct{} `type:"structure"`

	// The name of the scaling policy.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`

	// The type of scaling policy.
	//
	// PolicyType is a required field
	PolicyType PolicyType `type:"string" required:"true" enum:"true"`

	// The target tracking scaling policy. Includes support for predefined or customized
	// metrics.
	TargetTrackingConfiguration *TargetTrackingConfiguration `type:"structure"`
}

// String returns the string representation
func (s ScalingPolicy) String() string {
	return awsutil.Prettify(s)
}

// Represents a tag.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/TagFilter
type TagFilter struct {
	_ struct{} `type:"structure"`

	// The tag key.
	Key *string `min:"1" type:"string"`

	// The tag values (0 to 20).
	Values []string `type:"list"`
}

// String returns the string representation
func (s TagFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagFilter"}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a target tracking configuration to use with AWS Auto Scaling. Used
// with ScalingInstruction and ScalingPolicy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/TargetTrackingConfiguration
type TargetTrackingConfiguration struct {
	_ struct{} `type:"structure"`

	// A customized metric. You can specify either a predefined metric or a customized
	// metric.
	CustomizedScalingMetricSpecification *CustomizedScalingMetricSpecification `type:"structure"`

	// Indicates whether scale in by the target tracking scaling policy is disabled.
	// If the value is true, scale in is disabled and the target tracking scaling
	// policy doesn't remove capacity from the scalable resource. Otherwise, scale
	// in is enabled and the target tracking scaling policy can remove capacity
	// from the scalable resource.
	//
	// The default value is false.
	DisableScaleIn *bool `type:"boolean"`

	// The estimated time, in seconds, until a newly launched instance can contribute
	// to the CloudWatch metrics. This value is used only if the resource is an
	// Auto Scaling group.
	EstimatedInstanceWarmup *int64 `type:"integer"`

	// A predefined metric. You can specify either a predefined metric or a customized
	// metric.
	PredefinedScalingMetricSpecification *PredefinedScalingMetricSpecification `type:"structure"`

	// The amount of time, in seconds, after a scale in activity completes before
	// another scale in activity can start. This value is not used if the scalable
	// resource is an Auto Scaling group.
	//
	// The cooldown period is used to block subsequent scale in requests until it
	// has expired. The intention is to scale in conservatively to protect your
	// application's availability. However, if another alarm triggers a scale-out
	// policy during the cooldown period after a scale-in, AWS Auto Scaling scales
	// out your scalable target immediately.
	ScaleInCooldown *int64 `type:"integer"`

	// The amount of time, in seconds, after a scale-out activity completes before
	// another scale-out activity can start. This value is not used if the scalable
	// resource is an Auto Scaling group.
	//
	// While the cooldown period is in effect, the capacity that has been added
	// by the previous scale-out event that initiated the cooldown is calculated
	// as part of the desired capacity for the next scale out. The intention is
	// to continuously (but not excessively) scale out.
	ScaleOutCooldown *int64 `type:"integer"`

	// The target value for the metric. The range is 8.515920e-109 to 1.174271e+108
	// (Base 10) or 2e-360 to 2e360 (Base 2).
	//
	// TargetValue is a required field
	TargetValue *float64 `type:"double" required:"true"`
}

// String returns the string representation
func (s TargetTrackingConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TargetTrackingConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TargetTrackingConfiguration"}

	if s.TargetValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetValue"))
	}
	if s.CustomizedScalingMetricSpecification != nil {
		if err := s.CustomizedScalingMetricSpecification.Validate(); err != nil {
			invalidParams.AddNested("CustomizedScalingMetricSpecification", err.(aws.ErrInvalidParams))
		}
	}
	if s.PredefinedScalingMetricSpecification != nil {
		if err := s.PredefinedScalingMetricSpecification.Validate(); err != nil {
			invalidParams.AddNested("PredefinedScalingMetricSpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
