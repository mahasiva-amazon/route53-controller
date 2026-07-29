	// If the pre hook injected the empty-NameOrID sentinel, clear Status.ID so
	// sdkFind skips the ChangeInfo lookup for adopted pre-existing records.
	if r.ko.Status.ID != nil && *r.ko.Status.ID == "\x00" {
		r.ko.Status.ID = nil
	}

	f1, f1ok := identifier.AdditionalKeys["recordType"]
	if f1ok {
		r.ko.Spec.RecordType = &f1
	}

	if f2, f2ok := identifier.AdditionalKeys["name"]; f2ok {
		r.ko.Spec.Name = &f2
	}
