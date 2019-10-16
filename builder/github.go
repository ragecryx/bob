package builder

// GHPusher contains the most basic info of
// the person that triggered the push
type GHPusher struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GHSender contains more details about the
// person that triggered the push
type GHSender struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"html_url"`
	Type      string `json:"type"`
}

// GHRepository contains really basic info
// about a GitHub repository. It is contained
// in most webhook payloads.
type GHRepository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	URL      string `json:"html_url"`
	Created  bool   `json:"created"`
	Deleted  bool   `json:"deleted"`
	Forced   bool   `json:"forced"`
	BaseRef  string `json:"base_ref"`
}

// GHPushEvent represents a Push of changes
// either new branch or a merge etc
// https://developer.github.com/v3/activity/events/types/#pushevent
type GHPushEvent struct {
	Ref          string       `json:"ref"`
	SHAHead      string       `json:"head"`
	SHABefore    string       `json:"before"`
	Size         int          `json:"size"`
	DistinctSize int          `json:"distinct_size"`
	Repository   GHRepository `json:"repository"`
	Pusher       GHPusher     `json:"pusher"`
	Sender       GHSender     `json:"sender"`
}
