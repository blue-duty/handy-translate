package deepl

import (
	"handy-translate/config"
	"testing"
)

func TestTranslate(t *testing.T) {
	var deepl = &DeepL{
		Translate: config.Translate{
			Key: "",
		},
	}

	source := "Join Senior Developer Advocate Christina Warren for all the latest developer. open source, and GitHub news. This week, we look at the new GitHub Actions for Google Cloud, how to use GitHub Actions with your R projects, GitHub Sponsors for companies, and more."
	target, _ := deepl.PostQuery(source, "auto", "zh")

	t.Log(target)
}
