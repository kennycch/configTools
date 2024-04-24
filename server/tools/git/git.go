package git

import (
	"config_tools/app/dao"
	"config_tools/app/request"
	"config_tools/config"
	"fmt"
	"os"
	"strings"

	"github.com/kennycch/gotools/general"
	"gopkg.in/src-d/go-git.v4"
	conf "gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

// GetTargetPath 获取目标目录路径
func GetTargetPath(gitPath string, envType EnvType) string {
	// 获取git后缀
	paths := strings.Split(gitPath, "/")
	last := paths[len(paths)-1]
	// 去除.git部分
	names := strings.Split(last, ".")
	names = general.DeleteValueByKey(names, len(names)-1)
	name := strings.Join(names, "-")
	// 拼接前缀
	fullName := fmt.Sprintf("./file/git_projects/%s_%s", name, envType)
	return fullName
}

// Clone 克隆项目
func Clone(gitPath string, envType EnvType) error {
	// SSH认证选项
	auth, err := ssh.NewPublicKeysFromFile("git", config.Git.SshPath, "")
	if err != nil {
		return err
	}
	// 克隆仓库
	targetPath := GetTargetPath(gitPath, envType)
	_, err = git.PlainClone(targetPath, false, &git.CloneOptions{
		URL:           gitPath,
		Auth:          auth,
		Progress:      os.Stdout,
		ReferenceName: plumbing.NewBranchReferenceName(string(envType)),
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}

// WriteFile 修改文件
func WriteFile(gitPath, fileName, content string, envType EnvType) error {
	targetPath := GetTargetPath(gitPath, envType)
	err := os.WriteFile(fmt.Sprintf("%s/%s", targetPath, fileName), []byte(content), 0644)
	return err
}

// Push 发布修改
func Push(gitPath, comment string, envType EnvType) error {
	// SSH认证选项
	auth, err := ssh.NewPublicKeysFromFile("git", config.Git.SshPath, "")
	if err != nil {
		return err
	}
	// 打开本地仓库
	targetPath := GetTargetPath(gitPath, envType)
	repo, err := git.PlainOpen(targetPath)
	if err != nil {
		return err
	}
	// 获取工作目录
	wt, err := repo.Worktree()
	if err != nil {
		return err
	}
	// 将修改添加到暂存区
	_, err = wt.Add(".")
	if err != nil {
		return err
	}
	// 提交更改
	_, err = wt.Commit(comment, &git.CommitOptions{
		Author: &object.Signature{
			Name:  config.Git.Name,
			Email: config.Git.Email,
			When:  general.Now(),
		},
	})
	if err != nil {
		return err
	}
	// 推送更改到远程仓库
	err = repo.Push(&git.PushOptions{
		Auth:     auth,
		RefSpecs: []conf.RefSpec{conf.RefSpec(fmt.Sprintf("refs/heads/%s:refs/heads/%s", envType, envType))},
	})
	return err
}

// CloneByGame 克隆游戏Git项目
func CloneByGame(game *dao.Game) {
	Clone(game.ClientGit, Dev)
	Clone(game.ClientGit, Master)
	Clone(game.ServerGit, Dev)
	Clone(game.ServerGit, Master)
	Clone(game.ExcelGit, Dev)
	Clone(game.ExcelGit, Master)
}

// cloneAll 克隆所有项目
func cloneAll() {
	// 获取所有游戏
	req := &request.ListBaseRequest{
		Page:     1,
		PageSize: 100000,
	}
	games, _ := dao.GetGameList(req)
	// 逐个克隆项目
	for _, game := range games {
		CloneByGame(game)
	}
}
