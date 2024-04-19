package git

import "config_tools/tools/lifecycle"

type GitRegister struct{}

func (g *GitRegister) Start() {
	cloneAll()
}

func (g *GitRegister) Priority() uint32 {
	return lifecycle.NormalPriority
}

func (g *GitRegister) Stop() {

}

func NewGitRegister() *GitRegister {
	return &GitRegister{}
}
