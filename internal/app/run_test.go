package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Yunsang-Jeong/terraforge/internal/app"
	"github.com/Yunsang-Jeong/terraforge/internal/utils"
)

func TestRun(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("fail")
		}
	}()

	dotGitDir, _ := utils.GetSomethingPathInParents(".", ".git", false)
	gitRoot := filepath.Dir(dotGitDir)
	if err := os.Chdir(gitRoot); err != nil {
		panic(err)
	}

	debug := true
	cf := "terraforge.hcl"

	testSets := []struct {
		wd string
	}{
		{
			wd: "example/dev",
		},
		{
			wd: "example/prod",
		},
	}

	for _, testSet := range testSets {
		app.NewTerraforge(testSet.wd, cf, debug).Run()
	}
}
