package entities

import "github.com/etcinit/gonduit/util"

// DifferentialRevision represents a revision in Differential.
type DifferentialRevision struct {
	ID             string              `json:"id"`
	PHID           string              `json:"phid"`
	Title          string              `json:"title"`
	URI            string              `json:"uri"`
	DateCreated    util.UnixTimestamp  `json:"dateCreated"`
	DateModified   util.UnixTimestamp  `json:"dateModified"`
	AuthorPHID     string              `json:"authorPHID"`
	Status         string              `json:"status"`
	StatusName     string              `json:"statusName"`
	Branch         string              `json:"branch"`
	Summary        string              `json:"summary"`
	TestPlan       string              `json:"testPlan"`
	LineCount      string              `json:"lineCount"`
	ActiveDiffPHID string              `json:"activeDiffPHID"`
	Diffs          []string            `json:"diffs"`
	Commits        []string            `json:"commits"`
	Reviewers      []string            `json:"reviewers"`
	CCs            []string            `json:"ccs"`
	Hashes         [][]string          `json:"hashes"`
	Auxiliary      map[string][]string `json:"auxiliary"`
	RepositoryPHID string              `json:"repositoryPHID"`
}

// DifferentialDiff represents a specific diff within a Differential revision.
// A new diff is created every time you update a differential revision (that's
// what arc diff does duh).
//
// NOTE: Two fields are missing from this struct:
//
// - Changes (changes) is a list of a fairly complex data-structure with all
//   hunks contained in this diff along with some dynamically typed metadata;
// - Properties (properties) is another dynamically typed field which will be
//   an empty list on a closed diff (as far as I can tell) or a fairly complex
//   data-structure containing more metadata about the diff (info about the
//   local commits and about arc's interaction with the staging area if you
//   repository has one set up).
type DifferentialDiff struct {
	ID                        string             `json:"id"`
	RevisionID                string             `json:"revisionID"`
	DateCreated               util.UnixTimestamp `json:"dateCreated"`
	DateModified              util.UnixTimestamp `json:"dateModified"`
	SourceControlBaseRevision string             `json:"sourceControlBaseRevision"`
	SourceControlPath         string             `json:"sourceControlPath"`
	SourceControlSystem       string             `json:"sourceControlSystem"`
	Branch                    string             `json:"branch"`
	Bookmark                  string             `json:"bookmark"`
	CreationMethod            string             `json:"creationMethod"`
	Description               string             `json:"description"`
	UnitStatus                string             `json:"unitStatus"`
	LintStatus                string             `json:"lintStatus"`
	AuthorName                string             `json:"authorName"`
	AuthorEmail               string             `json:"authorEmail"`
}
