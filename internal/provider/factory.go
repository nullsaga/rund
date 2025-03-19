package provider

const (
	github    = "github"
	gitlab    = "gitlab"
	bitbucket = "bitbucket"
	gitea     = "gitea"
)

func New(name string) Provider {
	switch name {
	case github:
		return NewGithubProvider()
	case gitlab:
		return NewGitlabProvider()
	default:
		return nil // return generic
	}
}
