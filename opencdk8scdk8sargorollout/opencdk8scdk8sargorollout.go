// @opencdk8s/cdk8s-argo-rollout
package opencdk8scdk8sargorollout

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/opencdk8s/cdk8s-argo-rollouts-go/opencdk8scdk8sargorollout/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-go/cdk8s"
	"github.com/opencdk8s/cdk8s-argo-rollouts-go/opencdk8scdk8sargorollout/internal"
	"github.com/opencdk8s/cdk8s-argo-rollouts-go/opencdk8scdk8sargorollout/k8s"
)

// Experimental.
type AnalysisArgs struct {
	// Experimental.
	Name *string `json:"name"`
	// Experimental.
	Value *string `json:"value"`
	// Experimental.
	ValueFrom *ValueFrom `json:"valueFrom"`
}

// Experimental.
type AnalysisSpec struct {
	// Experimental.
	Args *[]*AnalysisArgs `json:"args"`
	// Experimental.
	Templates *[]*AnalysisTemplate `json:"templates"`
}

// Experimental.
type AnalysisTemplate struct {
	// Experimental.
	TemplateName *string `json:"templateName"`
}

// Experimental.
type ArgoRollout interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for ArgoRollout
type jsiiProxy_ArgoRollout struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ArgoRollout) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoRollout) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoRollout) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoRollout) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoRollout) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArgoRollout) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewArgoRollout(scope constructs.Construct, id *string, props *ArgoRolloutProps) ArgoRollout {
	_init_.Initialize()

	j := jsiiProxy_ArgoRollout{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRollout",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewArgoRollout_Override(a ArgoRollout, scope constructs.Construct, id *string, props *ArgoRolloutProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRollout",
		[]interface{}{scope, id, props},
		a,
	)
}

// Renders a Kubernetes manifest for an ingress object. https://github.com/kubernetes-sigs/aws-load-balancer-controller.
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func ArgoRollout_Manifest(props *ArgoRolloutProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRollout",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
// Experimental.
func ArgoRollout_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRollout",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ArgoRollout_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRollout",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
// Experimental.
func (a *jsiiProxy_ArgoRollout) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
// Experimental.
func (a *jsiiProxy_ArgoRollout) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_ArgoRollout) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_ArgoRollout) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (a *jsiiProxy_ArgoRollout) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
// Experimental.
func (a *jsiiProxy_ArgoRollout) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_ArgoRollout) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type ArgoRolloutProps struct {
	// Standard object's metadata.
	// Experimental.
	Metadata *k8s.ObjectMeta `json:"metadata"`
	// Spec defines the behavior of the ingress.
	// Experimental.
	Spec *ArgoSpecs `json:"spec"`
}

// Experimental.
type ArgoSpecs struct {
	// Experimental.
	Selector *k8s.LabelSelector `json:"selector"`
	// Experimental.
	Strategy *StrategySpecs `json:"strategy"`
	// Experimental.
	Template *k8s.PodTemplateSpec `json:"template"`
	// Experimental.
	MinReadySeconds *float64 `json:"minReadySeconds"`
	// Experimental.
	Paused *bool `json:"paused"`
	// Experimental.
	ProgressDeadlineSeconds *float64 `json:"progressDeadlineSeconds"`
	// Experimental.
	Replicas *float64 `json:"replicas"`
	// Experimental.
	RevisionHistoryLimit *float64 `json:"revisionHistoryLimit"`
}

// Experimental.
type BlueGreenStrategySpecs struct {
	// Experimental.
	ActiveService *string `json:"activeService"`
	// Experimental.
	AntiAffinity *k8s.PodAntiAffinity `json:"antiAffinity"`
	// Experimental.
	AutoPromotionEnabled *bool `json:"autoPromotionEnabled"`
	// Experimental.
	AutoPromotionSeconds *float64 `json:"autoPromotionSeconds"`
	// Experimental.
	PostPromotionAnalysis *AnalysisSpec `json:"postPromotionAnalysis"`
	// Experimental.
	PrePromotionAnalysis *AnalysisSpec `json:"prePromotionAnalysis"`
	// Experimental.
	PreviewReplicaCount *float64 `json:"previewReplicaCount"`
	// Experimental.
	PreviewService *string `json:"previewService"`
	// Experimental.
	ScaleDownDelayRevisionLimit *float64 `json:"scaleDownDelayRevisionLimit"`
	// Experimental.
	ScaleDownDelaySeconds *float64 `json:"scaleDownDelaySeconds"`
}

// Experimental.
type StrategySpecs struct {
	// Experimental.
	BlueGreen *BlueGreenStrategySpecs `json:"blueGreen"`
}

// Experimental.
type ValueFrom struct {
	// Experimental.
	FieldRef *k8s.ObjectFieldSelector `json:"fieldRef"`
	// Experimental.
	PodTemplateHashValue *string `json:"podTemplateHashValue"`
}

