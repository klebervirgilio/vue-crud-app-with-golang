package core

// Kudo represents a oos kudo.
type Kudo struct {
	UserID   string `json:"user_id" bson:"userId"`
	RepoID   string `json:"repo_id" bson:"repoId"`
	RepoName string `json:"repo_name" bson:"repoName"`
	RepoURL  string `json:"repo_url" bson:"repoUrl"`
	Language string `json:"language" bson:"language"`
	Notes    string `json:"notes" bson:"notes"`
}
