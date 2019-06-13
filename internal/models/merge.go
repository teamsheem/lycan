package models


type LastCommit struct {
	Id string `json:"id"`
	Message string `json:"message"`
}

type ObjectAttributes struct {
	Id int `json:"id"`
	TargetBranch string `json:"target_branch"` //later we can create filters based on target and source branch
	SourceBranch string `json:"source_branch"`
	MergeStatus string `json:"merge_status"`
	State string `json:"state"`
	LastCommit LastCommit `json:"last_commit"`
}

type Merge struct {
	ObjectKind string `json:"object_kind"`
	ObjectAttributes ObjectAttributes `json:"object_attributes" bindings:"required"`
}