package validate

import (
	"github.com/drlau/akashi/pkg/ruleset"
)

type InvalidResource struct {
	ResourceID *ruleset.ResourceIdentifier
}

type ValidateResult struct {
	Valid                   bool
	InvalidCreatedResources []InvalidResource
	InvalidDeletedResources []InvalidResource
	InvalidChangedResources []InvalidResource
}

func getUnnamedResources(rs []ruleset.Resource) []*ruleset.ResourceIdentifier {
	var res []*ruleset.ResourceIdentifier
	for _, r := range rs {
		id := r.ID()
		if id.Name == "" {
			res = append(res, id)
		}
	}
	return res
}

func Validate(rs ruleset.Ruleset) *ValidateResult {
	var res *ValidateResult
	if rs.CreatedResources.RequireName {
		ids := getUnnamedResources(rs.CreatedResources.Resources)
		&res.InvalidChangedResources = ids
	}
	return res
}
