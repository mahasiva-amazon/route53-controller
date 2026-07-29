// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/

package record_set

import (
	"testing"

	svcapitypes "github.com/aws-controllers-k8s/route53-controller/apis/v1alpha1"
)

func newTestResource() *resource {
	return &resource{
		ko: &svcapitypes.RecordSet{},
	}
}

func Test_PopulateResourceFromAnnotation_IDOptional(t *testing.T) {
	// id absent: Status.ID must remain nil, adoption succeeds
	r := newTestResource()
	fields := map[string]string{
		"hostedZoneID": "Z123456",
		"recordType":   "A",
		"name":         "www",
	}
	if err := r.PopulateResourceFromAnnotation(fields); err != nil {
		t.Fatalf("expected no error when id absent, got: %v", err)
	}
	if r.ko.Status.ID != nil {
		t.Errorf("expected Status.ID to be nil when id absent, got %q", *r.ko.Status.ID)
	}
	if r.ko.Spec.HostedZoneID == nil || *r.ko.Spec.HostedZoneID != "Z123456" {
		t.Errorf("expected HostedZoneID=Z123456, got %v", r.ko.Spec.HostedZoneID)
	}
}

func Test_PopulateResourceFromAnnotation_IDEmptyString(t *testing.T) {
	// id present but empty string: Status.ID must remain nil
	r := newTestResource()
	fields := map[string]string{
		"id":           "",
		"hostedZoneID": "Z123456",
		"recordType":   "A",
	}
	if err := r.PopulateResourceFromAnnotation(fields); err != nil {
		t.Fatalf("expected no error when id is empty string, got: %v", err)
	}
	if r.ko.Status.ID != nil {
		t.Errorf("expected Status.ID to be nil when id is empty string, got %q", *r.ko.Status.ID)
	}
}

func Test_PopulateResourceFromAnnotation_IDPresent(t *testing.T) {
	// id present with a real value: Status.ID must be set
	r := newTestResource()
	changeID := "/change/C027057324F8QNL1ERD2R"
	fields := map[string]string{
		"id":           changeID,
		"hostedZoneID": "Z123456",
		"recordType":   "A",
		"name":         "www",
	}
	if err := r.PopulateResourceFromAnnotation(fields); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.ko.Status.ID == nil || *r.ko.Status.ID != changeID {
		t.Errorf("expected Status.ID=%q, got %v", changeID, r.ko.Status.ID)
	}
}

func Test_PopulateResourceFromAnnotation_MissingHostedZoneID(t *testing.T) {
	// hostedZoneID is required: must return terminal error
	r := newTestResource()
	fields := map[string]string{
		"recordType": "A",
		"name":       "www",
	}
	if err := r.PopulateResourceFromAnnotation(fields); err == nil {
		t.Fatal("expected error when hostedZoneID missing, got nil")
	}
}

func Test_PopulateResourceFromAnnotation_OptionalFields(t *testing.T) {
	// recordType and name are optional
	r := newTestResource()
	fields := map[string]string{
		"hostedZoneID": "Z999",
	}
	if err := r.PopulateResourceFromAnnotation(fields); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.ko.Spec.RecordType != nil {
		t.Errorf("expected RecordType nil, got %q", *r.ko.Spec.RecordType)
	}
	if r.ko.Spec.Name != nil {
		t.Errorf("expected Name nil, got %q", *r.ko.Spec.Name)
	}
}
