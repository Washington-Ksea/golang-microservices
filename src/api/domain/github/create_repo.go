package github

/*
{
  "name": "golang-tutorial",
  "description": "This is your first repository",
  "homepage": "https://github.com",
  "private": false,
  "has_issues": true,
  "has_projects": true,
  "has_wiki": true
}
*/

type CreateRepoRequest struct {
	Name        string
	Description string
	Homepage    string
	Private     bool
	HasIssues   bool
	HasProjects bool
	HasWiki     bool
}

/*
func () CreateRepo() {
	req := map[string]interface{}{
		"name":    "Hello-World",
		"private": false,
	}
}
*/
