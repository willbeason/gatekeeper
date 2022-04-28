//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	templates "github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ByPodStatus)(nil), (*templates.ByPodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ByPodStatus_To_templates_ByPodStatus(a.(*ByPodStatus), b.(*templates.ByPodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.ByPodStatus)(nil), (*ByPodStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_ByPodStatus_To_v1beta1_ByPodStatus(a.(*templates.ByPodStatus), b.(*ByPodStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CRD)(nil), (*templates.CRD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CRD_To_templates_CRD(a.(*CRD), b.(*templates.CRD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.CRD)(nil), (*CRD)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_CRD_To_v1beta1_CRD(a.(*templates.CRD), b.(*CRD), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CRDSpec)(nil), (*templates.CRDSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CRDSpec_To_templates_CRDSpec(a.(*CRDSpec), b.(*templates.CRDSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.CRDSpec)(nil), (*CRDSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_CRDSpec_To_v1beta1_CRDSpec(a.(*templates.CRDSpec), b.(*CRDSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConstraintTemplate)(nil), (*templates.ConstraintTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ConstraintTemplate_To_templates_ConstraintTemplate(a.(*ConstraintTemplate), b.(*templates.ConstraintTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.ConstraintTemplate)(nil), (*ConstraintTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_ConstraintTemplate_To_v1beta1_ConstraintTemplate(a.(*templates.ConstraintTemplate), b.(*ConstraintTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConstraintTemplateList)(nil), (*templates.ConstraintTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ConstraintTemplateList_To_templates_ConstraintTemplateList(a.(*ConstraintTemplateList), b.(*templates.ConstraintTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.ConstraintTemplateList)(nil), (*ConstraintTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_ConstraintTemplateList_To_v1beta1_ConstraintTemplateList(a.(*templates.ConstraintTemplateList), b.(*ConstraintTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConstraintTemplateSpec)(nil), (*templates.ConstraintTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ConstraintTemplateSpec_To_templates_ConstraintTemplateSpec(a.(*ConstraintTemplateSpec), b.(*templates.ConstraintTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.ConstraintTemplateSpec)(nil), (*ConstraintTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_ConstraintTemplateSpec_To_v1beta1_ConstraintTemplateSpec(a.(*templates.ConstraintTemplateSpec), b.(*ConstraintTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConstraintTemplateStatus)(nil), (*templates.ConstraintTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ConstraintTemplateStatus_To_templates_ConstraintTemplateStatus(a.(*ConstraintTemplateStatus), b.(*templates.ConstraintTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.ConstraintTemplateStatus)(nil), (*ConstraintTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_ConstraintTemplateStatus_To_v1beta1_ConstraintTemplateStatus(a.(*templates.ConstraintTemplateStatus), b.(*ConstraintTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CreateCRDError)(nil), (*templates.CreateCRDError)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CreateCRDError_To_templates_CreateCRDError(a.(*CreateCRDError), b.(*templates.CreateCRDError), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.CreateCRDError)(nil), (*CreateCRDError)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_CreateCRDError_To_v1beta1_CreateCRDError(a.(*templates.CreateCRDError), b.(*CreateCRDError), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Names)(nil), (*templates.Names)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Names_To_templates_Names(a.(*Names), b.(*templates.Names), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.Names)(nil), (*Names)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_Names_To_v1beta1_Names(a.(*templates.Names), b.(*Names), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Target)(nil), (*templates.Target)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Target_To_templates_Target(a.(*Target), b.(*templates.Target), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.Target)(nil), (*Target)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_Target_To_v1beta1_Target(a.(*templates.Target), b.(*Target), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*templates.Validation)(nil), (*Validation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_templates_Validation_To_v1beta1_Validation(a.(*templates.Validation), b.(*Validation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*Validation)(nil), (*templates.Validation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Validation_To_templates_Validation(a.(*Validation), b.(*templates.Validation), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_ByPodStatus_To_templates_ByPodStatus(in *ByPodStatus, out *templates.ByPodStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.ObservedGeneration = in.ObservedGeneration
	out.Errors = *(*[]templates.CreateCRDError)(unsafe.Pointer(&in.Errors))
	return nil
}

// Convert_v1beta1_ByPodStatus_To_templates_ByPodStatus is an autogenerated conversion function.
func Convert_v1beta1_ByPodStatus_To_templates_ByPodStatus(in *ByPodStatus, out *templates.ByPodStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ByPodStatus_To_templates_ByPodStatus(in, out, s)
}

func autoConvert_templates_ByPodStatus_To_v1beta1_ByPodStatus(in *templates.ByPodStatus, out *ByPodStatus, s conversion.Scope) error {
	out.ID = in.ID
	out.ObservedGeneration = in.ObservedGeneration
	out.Errors = *(*[]CreateCRDError)(unsafe.Pointer(&in.Errors))
	return nil
}

// Convert_templates_ByPodStatus_To_v1beta1_ByPodStatus is an autogenerated conversion function.
func Convert_templates_ByPodStatus_To_v1beta1_ByPodStatus(in *templates.ByPodStatus, out *ByPodStatus, s conversion.Scope) error {
	return autoConvert_templates_ByPodStatus_To_v1beta1_ByPodStatus(in, out, s)
}

func autoConvert_v1beta1_CRD_To_templates_CRD(in *CRD, out *templates.CRD, s conversion.Scope) error {
	if err := Convert_v1beta1_CRDSpec_To_templates_CRDSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CRD_To_templates_CRD is an autogenerated conversion function.
func Convert_v1beta1_CRD_To_templates_CRD(in *CRD, out *templates.CRD, s conversion.Scope) error {
	return autoConvert_v1beta1_CRD_To_templates_CRD(in, out, s)
}

func autoConvert_templates_CRD_To_v1beta1_CRD(in *templates.CRD, out *CRD, s conversion.Scope) error {
	if err := Convert_templates_CRDSpec_To_v1beta1_CRDSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_templates_CRD_To_v1beta1_CRD is an autogenerated conversion function.
func Convert_templates_CRD_To_v1beta1_CRD(in *templates.CRD, out *CRD, s conversion.Scope) error {
	return autoConvert_templates_CRD_To_v1beta1_CRD(in, out, s)
}

func autoConvert_v1beta1_CRDSpec_To_templates_CRDSpec(in *CRDSpec, out *templates.CRDSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_Names_To_templates_Names(&in.Names, &out.Names, s); err != nil {
		return err
	}
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(templates.Validation)
		if err := Convert_v1beta1_Validation_To_templates_Validation(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Validation = nil
	}
	return nil
}

// Convert_v1beta1_CRDSpec_To_templates_CRDSpec is an autogenerated conversion function.
func Convert_v1beta1_CRDSpec_To_templates_CRDSpec(in *CRDSpec, out *templates.CRDSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CRDSpec_To_templates_CRDSpec(in, out, s)
}

func autoConvert_templates_CRDSpec_To_v1beta1_CRDSpec(in *templates.CRDSpec, out *CRDSpec, s conversion.Scope) error {
	if err := Convert_templates_Names_To_v1beta1_Names(&in.Names, &out.Names, s); err != nil {
		return err
	}
	if in.Validation != nil {
		in, out := &in.Validation, &out.Validation
		*out = new(Validation)
		if err := Convert_templates_Validation_To_v1beta1_Validation(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Validation = nil
	}
	return nil
}

// Convert_templates_CRDSpec_To_v1beta1_CRDSpec is an autogenerated conversion function.
func Convert_templates_CRDSpec_To_v1beta1_CRDSpec(in *templates.CRDSpec, out *CRDSpec, s conversion.Scope) error {
	return autoConvert_templates_CRDSpec_To_v1beta1_CRDSpec(in, out, s)
}

func autoConvert_v1beta1_ConstraintTemplate_To_templates_ConstraintTemplate(in *ConstraintTemplate, out *templates.ConstraintTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ConstraintTemplateSpec_To_templates_ConstraintTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ConstraintTemplateStatus_To_templates_ConstraintTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_ConstraintTemplate_To_templates_ConstraintTemplate is an autogenerated conversion function.
func Convert_v1beta1_ConstraintTemplate_To_templates_ConstraintTemplate(in *ConstraintTemplate, out *templates.ConstraintTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_ConstraintTemplate_To_templates_ConstraintTemplate(in, out, s)
}

func autoConvert_templates_ConstraintTemplate_To_v1beta1_ConstraintTemplate(in *templates.ConstraintTemplate, out *ConstraintTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_templates_ConstraintTemplateSpec_To_v1beta1_ConstraintTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_templates_ConstraintTemplateStatus_To_v1beta1_ConstraintTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_templates_ConstraintTemplate_To_v1beta1_ConstraintTemplate is an autogenerated conversion function.
func Convert_templates_ConstraintTemplate_To_v1beta1_ConstraintTemplate(in *templates.ConstraintTemplate, out *ConstraintTemplate, s conversion.Scope) error {
	return autoConvert_templates_ConstraintTemplate_To_v1beta1_ConstraintTemplate(in, out, s)
}

func autoConvert_v1beta1_ConstraintTemplateList_To_templates_ConstraintTemplateList(in *ConstraintTemplateList, out *templates.ConstraintTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]templates.ConstraintTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_ConstraintTemplate_To_templates_ConstraintTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_ConstraintTemplateList_To_templates_ConstraintTemplateList is an autogenerated conversion function.
func Convert_v1beta1_ConstraintTemplateList_To_templates_ConstraintTemplateList(in *ConstraintTemplateList, out *templates.ConstraintTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_ConstraintTemplateList_To_templates_ConstraintTemplateList(in, out, s)
}

func autoConvert_templates_ConstraintTemplateList_To_v1beta1_ConstraintTemplateList(in *templates.ConstraintTemplateList, out *ConstraintTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConstraintTemplate, len(*in))
		for i := range *in {
			if err := Convert_templates_ConstraintTemplate_To_v1beta1_ConstraintTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_templates_ConstraintTemplateList_To_v1beta1_ConstraintTemplateList is an autogenerated conversion function.
func Convert_templates_ConstraintTemplateList_To_v1beta1_ConstraintTemplateList(in *templates.ConstraintTemplateList, out *ConstraintTemplateList, s conversion.Scope) error {
	return autoConvert_templates_ConstraintTemplateList_To_v1beta1_ConstraintTemplateList(in, out, s)
}

func autoConvert_v1beta1_ConstraintTemplateSpec_To_templates_ConstraintTemplateSpec(in *ConstraintTemplateSpec, out *templates.ConstraintTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_CRD_To_templates_CRD(&in.CRD, &out.CRD, s); err != nil {
		return err
	}
	out.Targets = *(*[]templates.Target)(unsafe.Pointer(&in.Targets))
	return nil
}

// Convert_v1beta1_ConstraintTemplateSpec_To_templates_ConstraintTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_ConstraintTemplateSpec_To_templates_ConstraintTemplateSpec(in *ConstraintTemplateSpec, out *templates.ConstraintTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_ConstraintTemplateSpec_To_templates_ConstraintTemplateSpec(in, out, s)
}

func autoConvert_templates_ConstraintTemplateSpec_To_v1beta1_ConstraintTemplateSpec(in *templates.ConstraintTemplateSpec, out *ConstraintTemplateSpec, s conversion.Scope) error {
	if err := Convert_templates_CRD_To_v1beta1_CRD(&in.CRD, &out.CRD, s); err != nil {
		return err
	}
	out.Targets = *(*[]Target)(unsafe.Pointer(&in.Targets))
	return nil
}

// Convert_templates_ConstraintTemplateSpec_To_v1beta1_ConstraintTemplateSpec is an autogenerated conversion function.
func Convert_templates_ConstraintTemplateSpec_To_v1beta1_ConstraintTemplateSpec(in *templates.ConstraintTemplateSpec, out *ConstraintTemplateSpec, s conversion.Scope) error {
	return autoConvert_templates_ConstraintTemplateSpec_To_v1beta1_ConstraintTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_ConstraintTemplateStatus_To_templates_ConstraintTemplateStatus(in *ConstraintTemplateStatus, out *templates.ConstraintTemplateStatus, s conversion.Scope) error {
	out.Created = in.Created
	out.ByPod = *(*[]templates.ByPodStatus)(unsafe.Pointer(&in.ByPod))
	return nil
}

// Convert_v1beta1_ConstraintTemplateStatus_To_templates_ConstraintTemplateStatus is an autogenerated conversion function.
func Convert_v1beta1_ConstraintTemplateStatus_To_templates_ConstraintTemplateStatus(in *ConstraintTemplateStatus, out *templates.ConstraintTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_ConstraintTemplateStatus_To_templates_ConstraintTemplateStatus(in, out, s)
}

func autoConvert_templates_ConstraintTemplateStatus_To_v1beta1_ConstraintTemplateStatus(in *templates.ConstraintTemplateStatus, out *ConstraintTemplateStatus, s conversion.Scope) error {
	out.Created = in.Created
	out.ByPod = *(*[]ByPodStatus)(unsafe.Pointer(&in.ByPod))
	return nil
}

// Convert_templates_ConstraintTemplateStatus_To_v1beta1_ConstraintTemplateStatus is an autogenerated conversion function.
func Convert_templates_ConstraintTemplateStatus_To_v1beta1_ConstraintTemplateStatus(in *templates.ConstraintTemplateStatus, out *ConstraintTemplateStatus, s conversion.Scope) error {
	return autoConvert_templates_ConstraintTemplateStatus_To_v1beta1_ConstraintTemplateStatus(in, out, s)
}

func autoConvert_v1beta1_CreateCRDError_To_templates_CreateCRDError(in *CreateCRDError, out *templates.CreateCRDError, s conversion.Scope) error {
	out.Code = in.Code
	out.Message = in.Message
	out.Location = in.Location
	return nil
}

// Convert_v1beta1_CreateCRDError_To_templates_CreateCRDError is an autogenerated conversion function.
func Convert_v1beta1_CreateCRDError_To_templates_CreateCRDError(in *CreateCRDError, out *templates.CreateCRDError, s conversion.Scope) error {
	return autoConvert_v1beta1_CreateCRDError_To_templates_CreateCRDError(in, out, s)
}

func autoConvert_templates_CreateCRDError_To_v1beta1_CreateCRDError(in *templates.CreateCRDError, out *CreateCRDError, s conversion.Scope) error {
	out.Code = in.Code
	out.Message = in.Message
	out.Location = in.Location
	return nil
}

// Convert_templates_CreateCRDError_To_v1beta1_CreateCRDError is an autogenerated conversion function.
func Convert_templates_CreateCRDError_To_v1beta1_CreateCRDError(in *templates.CreateCRDError, out *CreateCRDError, s conversion.Scope) error {
	return autoConvert_templates_CreateCRDError_To_v1beta1_CreateCRDError(in, out, s)
}

func autoConvert_v1beta1_Names_To_templates_Names(in *Names, out *templates.Names, s conversion.Scope) error {
	out.Kind = in.Kind
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	return nil
}

// Convert_v1beta1_Names_To_templates_Names is an autogenerated conversion function.
func Convert_v1beta1_Names_To_templates_Names(in *Names, out *templates.Names, s conversion.Scope) error {
	return autoConvert_v1beta1_Names_To_templates_Names(in, out, s)
}

func autoConvert_templates_Names_To_v1beta1_Names(in *templates.Names, out *Names, s conversion.Scope) error {
	out.Kind = in.Kind
	out.ShortNames = *(*[]string)(unsafe.Pointer(&in.ShortNames))
	return nil
}

// Convert_templates_Names_To_v1beta1_Names is an autogenerated conversion function.
func Convert_templates_Names_To_v1beta1_Names(in *templates.Names, out *Names, s conversion.Scope) error {
	return autoConvert_templates_Names_To_v1beta1_Names(in, out, s)
}

func autoConvert_v1beta1_Target_To_templates_Target(in *Target, out *templates.Target, s conversion.Scope) error {
	out.Target = in.Target
	out.Rego = in.Rego
	out.Libs = *(*[]string)(unsafe.Pointer(&in.Libs))
	out.CELX = in.CELX
	return nil
}

// Convert_v1beta1_Target_To_templates_Target is an autogenerated conversion function.
func Convert_v1beta1_Target_To_templates_Target(in *Target, out *templates.Target, s conversion.Scope) error {
	return autoConvert_v1beta1_Target_To_templates_Target(in, out, s)
}

func autoConvert_templates_Target_To_v1beta1_Target(in *templates.Target, out *Target, s conversion.Scope) error {
	out.Target = in.Target
	out.Rego = in.Rego
	out.Libs = *(*[]string)(unsafe.Pointer(&in.Libs))
	out.CELX = in.CELX
	return nil
}

// Convert_templates_Target_To_v1beta1_Target is an autogenerated conversion function.
func Convert_templates_Target_To_v1beta1_Target(in *templates.Target, out *Target, s conversion.Scope) error {
	return autoConvert_templates_Target_To_v1beta1_Target(in, out, s)
}

func autoConvert_templates_Validation_To_v1beta1_Validation(in *templates.Validation, out *Validation, s conversion.Scope) error {
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		*out = new(v1.JSONSchemaProps)
		if err := v1.Convert_apiextensions_JSONSchemaProps_To_v1_JSONSchemaProps(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.OpenAPIV3Schema = nil
	}
	out.LegacySchema = (*bool)(unsafe.Pointer(in.LegacySchema))
	return nil
}

// Convert_templates_Validation_To_v1beta1_Validation is an autogenerated conversion function.
func Convert_templates_Validation_To_v1beta1_Validation(in *templates.Validation, out *Validation, s conversion.Scope) error {
	return autoConvert_templates_Validation_To_v1beta1_Validation(in, out, s)
}
