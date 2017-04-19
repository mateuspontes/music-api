package music

type Music struct {
	ID       string `bson:"_id" json:"id"`
	Name     string `json:"name"`
	Duration string `json:"durantion,omitempty"`
	Author   string `json:"author,omitempty"`
}
