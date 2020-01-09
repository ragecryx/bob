package builder

// BBActor contains more details about the
// person that triggered the push
type BBActor struct {
	DisplayName string `json:"display_name"`
	UUID        string `json:"uuid"`
	AccountID   string `json:"account_id"`
	Nickname    string `json:"nickname"`
	Type        string `json:"type"`
}

// BBRepository contains really basic info
// about the BitBucket repository.
// It is contained in most webhook payloads.
type BBRepository struct {
	SCM       string  `json:"scm"`
	Website   *string `json:"website"`
	Name      string  `json:"name"`
	FullName  string  `json:"full_name"`
	Owner     BBActor `json:"owner"`
	Type      string  `json:"type"`
	IsPrivate bool    `json:"is_private"`
	UUID      string  `json:"uuid"`
}

// BBReferenceState contains state of the reference
// It is used in BBPushDetails to capture the state
// before (old) the push and after (new) it.
type BBReferenceState struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Target struct {
		Hash    string `json:"hash"`
		Type    string `json:"type"`
		Date    string `json:"date"`
		Message string `json:"message"`
		Author  struct {
			Type string `json:"type"`
			Raw  string `json:"raw"`
		} `json:"author"`
	} `json:"target"`
}

// BBPushDetails contains all the details
// and references of the push.
type BBPushDetails struct {
	Changes []struct {
		Forced    bool `json:"forced"` // Force Push
		Created   bool `json:"created"`
		Truncated bool `json:"truncated"` //  whether BB truncated the commits array in this payload.
		Closed    bool `json:"closed"`
		// An object containing information about the state
		// of the reference before the push. When a branch
		// is created, old is null.
		Old *BBReferenceState `json:"old"`
		// An object containing information about the state
		// of the reference after the push. When a branch is
		// deleted, new is null.
		New *BBReferenceState `json:"new"`
	} `json:"changes"`
}

// BBPushEvent represents a Push of changes
// either new branch or a merge etc
//
// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-Push
type BBPushEvent struct {
	Push       BBPushDetails `json:"push"`
	Actor      BBActor       `json:"actor"`
	Repository BBRepository  `json:"repository"`
}
