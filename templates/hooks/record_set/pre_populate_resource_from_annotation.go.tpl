	// id maps to Status.ID (ChangeInfo ID). It is intentionally optional for
	// adoption of pre-existing records, which have no associated ChangeInfo.
	// Inject an empty string so the generated required-field check passes;
	// the post hook clears Status.ID when it finds the injected empty value.
	if _, ok := fields["id"]; !ok {
		fields["id"] = ""
	}
