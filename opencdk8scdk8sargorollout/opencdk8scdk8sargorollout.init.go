package opencdk8scdk8sargorollout

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.AnalysisArgs",
		reflect.TypeOf((*AnalysisArgs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.AnalysisSpec",
		reflect.TypeOf((*AnalysisSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.AnalysisTemplate",
		reflect.TypeOf((*AnalysisTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRollout",
		reflect.TypeOf((*ArgoRollout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ArgoRollout{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.ArgoRolloutProps",
		reflect.TypeOf((*ArgoRolloutProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.ArgoSpecs",
		reflect.TypeOf((*ArgoSpecs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.BlueGreenStrategySpecs",
		reflect.TypeOf((*BlueGreenStrategySpecs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.StrategySpecs",
		reflect.TypeOf((*StrategySpecs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-argo-rollout.ValueFrom",
		reflect.TypeOf((*ValueFrom)(nil)).Elem(),
	)
}
