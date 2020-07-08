package baiduai

import (
	"testing"
)

func TestTextCensoringV2(t *testing.T) {
	res, err := TextCensoringV2(&Config{
		ApiKey:    "",
		SecretKey: "",
		Content:   "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("res:", res)
	for i := range res.Data {
		t.Log(*res.Data[i])
		for j := range res.Data[i].Hits {
			t.Log(*res.Data[i].Hits[j])
		}
	}
}
