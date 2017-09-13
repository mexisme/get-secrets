// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudformation provides the client and types for making API
// requests to AWS CloudFormation.
//
// AWS CloudFormation allows you to create and manage AWS infrastructure deployments
// predictably and repeatedly. You can use AWS CloudFormation to leverage AWS
// products, such as Amazon Elastic Compute Cloud, Amazon Elastic Block Store,
// Amazon Simple Notification Service, Elastic Load Balancing, and Auto Scaling
// to build highly-reliable, highly scalable, cost-effective applications without
// creating or configuring the underlying AWS infrastructure.
//
// With AWS CloudFormation, you declare all of your resources and dependencies
// in a template file. The template defines a collection of resources as a single
// unit called a stack. AWS CloudFormation creates and deletes all member resources
// of the stack together and manages all dependencies between the resources
// for you.
//
// For more information about AWS CloudFormation, see the AWS CloudFormation
// Product Page (http://aws.amazon.com/cloudformation/).
//
// Amazon CloudFormation makes use of other AWS products. If you need additional
// technical information about a specific AWS product, you can find the product's
// technical documentation at docs.aws.amazon.com (http://docs.aws.amazon.com/).
//
// APIs for stacks
//
// When you use AWS CloudFormation, you manage related resources as a single
// unit called a stack. You create, update, and delete a collection of resources
// by creating, updating, and deleting stacks. All the resources in a stack
// are defined by the stack's AWS CloudFormation template.
//
// Actions
//
//    * CancelUpdateStack (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CancelUpdateStack.html)
//
//    * ContinueUpdateRollback (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ContinueUpdateRollback.html)
//
//    * CreateStack (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStack.html)
//
//    * DeleteStack (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteStack.html)
//
//    * DescribeStackEvents (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackEvents.html)
//
//    * DescribeStackResource (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackResource.html)
//
//    * DescribeStackResources (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackResources.html)
//
//    * DescribeStacks (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStacks.html)
//
//    * EstimateTemplateCost (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_EstimateTemplateCost.html)
//
//    * GetStackPolicy (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetStackPolicy.html)
//
//    * GetTemplate (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplate.html)
//
//    * GetTemplateSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html)
//
//    * ListExports (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListExports.html)
//
//    * ListImports (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListImports.html)
//
//    * ListStackResources (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackResources.html)
//
//    * ListStacks (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStacks.html)
//
//    * SetStackPolicy (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_SetStackPolicy.html)
//
//    * UpdateStack (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStack.html)
//
//    * ValidateTemplate (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ValidateTemplate.html)
//
// Data Types
//
//    * Export (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Export.html)
//
//    * Parameter (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
//
//    * ParameterConstraints (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ParameterConstraints.html)
//
//    * ParameterDeclaration (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ParameterDeclaration.html)
//
//    * Stack (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Stack.html)
//
//    * StackEvent (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackEvent.html)
//
//    * StackResource (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResource.html)
//
//    * StackResourceDetail (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceDetail.html)
//
//    * StackResourceSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html)
//
//    * StackSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSummary.html)
//
//    * Tag (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Tag.html)
//
//    * TemplateParameter (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_TemplateParameter.html)
//
// APIs for change sets
//
// If you need to make changes to the running resources in a stack, you update
// the stack. Before making changes to your resources, you can generate a change
// set, which is summary of your proposed changes. Change sets allow you to
// see how your changes might impact your running resources, especially for
// critical resources, before implementing them.
//
// Actions
//
//    * CreateChangeSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateChangeSet.html)
//
//    * DeleteChangeSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteChangeSet.html)
//
//    * DescribeChangeSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeChangeSet.html)
//
//    * ExecuteChangeSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ExecuteChangeSet.html)
//
//    * ListChangeSets (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListChangeSets.html)
//
// Data Types
//
//    * Change (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Change.html)
//
//    * ChangeSetSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ChangeSetSummary.html)
//
//    * ResourceChange (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ResourceChange.html)
//
//    * ResourceChangeDetail (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ResourceChangeDetail.html)
//
//    * ResourceTargetDefinition (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ResourceTargetDefinition.html)
//
// APIs for stack sets
//
// AWS CloudFormation StackSets lets you create a collection, or stack set,
// of stacks that can automatically and safely provision a common set of AWS
// resources across multiple AWS accounts and multiple AWS regions from a single
// AWS CloudFormation template. When you create a stack set, AWS CloudFormation
// provisions a stack in each of the specified accounts and regions by using
// the supplied AWS CloudFormation template and parameters. Stack sets let you
// manage a common set of AWS resources in a selection of accounts and regions
// in a single operation.
//
// Actions
//
//    * CreateStackInstances (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html)
//
//    * CreateStackSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackSet.html)
//
//    * DeleteStackInstances (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteStackInstances.html)
//
//    * DeleteStackSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DeleteStackSet.html)
//
//    * DescribeStackInstance (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackInstance.html)
//
//    * DescribeStackSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackSet.html)
//
//    * DescribeStackSetOperation (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_DescribeStackSetOperation.html)
//
//    * ListStackInstances (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackInstances.html)
//
//    * ListStackSetOperationResults (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSetOperationResults)
//
//    * ListStackSetOperations (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSetOperations)
//
//    * ListStackSets (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_ListStackSets)
//
//    * StopStackSetOperation (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StopStackSetOperation.html)
//
//    * UpdateStackSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
//
// Data Types
//
//    * Parameter (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
//
//    * StackInstance (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstance.html.html)
//
//    * StackInstanceSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackInstanceSummary.html.html)
//
//    * StackSet (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSet.html)
//
//    * StackSetOperation (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperation.html.html)
//
//    * StackSetOperationPreferences (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationPreferences.html.html)
//
//    * StackSetOperationResultSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationResultSummary.html.html)
//
//    * StackSetOperationSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetOperationSummary.html.html)
//
//    * StackSetSummary (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackSetSummary.html)
//
//    * Tag (http://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Tag.html)
//
// See https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15 for more information on this service.
//
// See cloudformation package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/
//
// Using the Client
//
// To AWS CloudFormation with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CloudFormation client CloudFormation for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/#New
package cloudformation
