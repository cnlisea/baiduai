package baiduai

import "testing"

func TestTextCensoring(t *testing.T) {
	res, err := TextCensoring(&Config{
		ApiKey:    "",
		SecretKey: "",
		Content:   "",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("span", res.Result.Spam)
	t.Log("reject:")
	for i := range res.Result.Reject {
		t.Log(res.Result.Reject[i].Hit)
	}

	t.Log("review:")
	for i := range res.Result.Review {
		t.Log(res.Result.Review[i].Hit)
	}

	t.Log("pass")
	for i := range res.Result.Pass {
		t.Log(res.Result.Pass[i].Hit)
	}
}
