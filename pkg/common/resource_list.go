package common

import v1 "k8s.io/api/core/v1"

type On struct {
	Success Action
	Fail    Action
}

type WrappedConfigMap struct {
	Ref v1.ConfigMap
}

func (i WrappedConfigMap) Exists(msg string) Action {
	return &ExistsConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i WrappedConfigMap) Update(msg string) Action {
	return &UpdateConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i WrappedConfigMap) Create(msg string) Action {
	return &CreateConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i WrappedConfigMap) EnsureReady(msg string) Action {
	return &EnsureReadyConfigMapAction{
		ref: &i.Ref,
		msg: msg,
	}
}

func (i WrappedConfigMap) Branch(on On) Action {
	return &OnAction{
		ref:     &i.Ref,
		success: on.Success,
		fail:    on.Fail,
	}
}
