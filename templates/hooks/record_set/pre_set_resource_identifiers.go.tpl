	// ChangeInfo ID (NameOrID) is optional for adoption of pre-existing records.
	// Inject a sentinel so the generated required-field guard passes; the post
	// hook will clear Status.ID when it sees the sentinel.
	if identifier.NameOrID == "" {
		identifier.NameOrID = "\x00"
	}
