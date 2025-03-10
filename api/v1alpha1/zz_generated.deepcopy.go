// +build !ignore_autogenerated

/*
Copyright 2021 The cert-manager authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	certmanagerv1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	metav1 "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicy) DeepCopyInto(out *CertificateRequestPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicy.
func (in *CertificateRequestPolicy) DeepCopy() *CertificateRequestPolicy {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequestPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyCondition) DeepCopyInto(out *CertificateRequestPolicyCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyCondition.
func (in *CertificateRequestPolicyCondition) DeepCopy() *CertificateRequestPolicyCondition {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyList) DeepCopyInto(out *CertificateRequestPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateRequestPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyList.
func (in *CertificateRequestPolicyList) DeepCopy() *CertificateRequestPolicyList {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRequestPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicySpec) DeepCopyInto(out *CertificateRequestPolicySpec) {
	*out = *in
	if in.AllowedSubject != nil {
		in, out := &in.AllowedSubject, &out.AllowedSubject
		*out = new(PolicyX509Subject)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedCommonName != nil {
		in, out := &in.AllowedCommonName, &out.AllowedCommonName
		*out = new(string)
		**out = **in
	}
	if in.MinDuration != nil {
		in, out := &in.MinDuration, &out.MinDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.MaxDuration != nil {
		in, out := &in.MaxDuration, &out.MaxDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.AllowedDNSNames != nil {
		in, out := &in.AllowedDNSNames, &out.AllowedDNSNames
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedIPAddresses != nil {
		in, out := &in.AllowedIPAddresses, &out.AllowedIPAddresses
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedURIs != nil {
		in, out := &in.AllowedURIs, &out.AllowedURIs
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedEmailAddresses != nil {
		in, out := &in.AllowedEmailAddresses, &out.AllowedEmailAddresses
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedIssuers != nil {
		in, out := &in.AllowedIssuers, &out.AllowedIssuers
		*out = new([]metav1.ObjectReference)
		if **in != nil {
			in, out := *in, *out
			*out = make([]metav1.ObjectReference, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedIsCA != nil {
		in, out := &in.AllowedIsCA, &out.AllowedIsCA
		*out = new(bool)
		**out = **in
	}
	if in.AllowedUsages != nil {
		in, out := &in.AllowedUsages, &out.AllowedUsages
		*out = new([]certmanagerv1.KeyUsage)
		if **in != nil {
			in, out := *in, *out
			*out = make([]certmanagerv1.KeyUsage, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedPrivateKey != nil {
		in, out := &in.AllowedPrivateKey, &out.AllowedPrivateKey
		*out = new(PolicyPrivateKey)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalPolicyServers != nil {
		in, out := &in.ExternalPolicyServers, &out.ExternalPolicyServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicySpec.
func (in *CertificateRequestPolicySpec) DeepCopy() *CertificateRequestPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRequestPolicyStatus) DeepCopyInto(out *CertificateRequestPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateRequestPolicyCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRequestPolicyStatus.
func (in *CertificateRequestPolicyStatus) DeepCopy() *CertificateRequestPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateRequestPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyPrivateKey) DeepCopyInto(out *PolicyPrivateKey) {
	*out = *in
	if in.AllowedAlgorithm != nil {
		in, out := &in.AllowedAlgorithm, &out.AllowedAlgorithm
		*out = new(certmanagerv1.PrivateKeyAlgorithm)
		**out = **in
	}
	if in.MinSize != nil {
		in, out := &in.MinSize, &out.MinSize
		*out = new(int)
		**out = **in
	}
	if in.MaxSize != nil {
		in, out := &in.MaxSize, &out.MaxSize
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyPrivateKey.
func (in *PolicyPrivateKey) DeepCopy() *PolicyPrivateKey {
	if in == nil {
		return nil
	}
	out := new(PolicyPrivateKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyX509Subject) DeepCopyInto(out *PolicyX509Subject) {
	*out = *in
	if in.AllowedOrganizations != nil {
		in, out := &in.AllowedOrganizations, &out.AllowedOrganizations
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedCountries != nil {
		in, out := &in.AllowedCountries, &out.AllowedCountries
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedOrganizationalUnits != nil {
		in, out := &in.AllowedOrganizationalUnits, &out.AllowedOrganizationalUnits
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedLocalities != nil {
		in, out := &in.AllowedLocalities, &out.AllowedLocalities
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedProvinces != nil {
		in, out := &in.AllowedProvinces, &out.AllowedProvinces
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedStreetAddresses != nil {
		in, out := &in.AllowedStreetAddresses, &out.AllowedStreetAddresses
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedPostalCodes != nil {
		in, out := &in.AllowedPostalCodes, &out.AllowedPostalCodes
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.AllowedSerialNumber != nil {
		in, out := &in.AllowedSerialNumber, &out.AllowedSerialNumber
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyX509Subject.
func (in *PolicyX509Subject) DeepCopy() *PolicyX509Subject {
	if in == nil {
		return nil
	}
	out := new(PolicyX509Subject)
	in.DeepCopyInto(out)
	return out
}
