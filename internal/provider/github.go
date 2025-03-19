package provider

type GithubProvider struct {
	Name string
}

func NewGithubProvider() *GithubProvider {
	return &GithubProvider{
		Name: github,
	}
}

func (gh *GithubProvider) Handle() {

}
