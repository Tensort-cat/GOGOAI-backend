package test

import (
	"context"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func TestApiKey(t *testing.T) {
	// 构造 client
	client := openai.NewClient(
		option.WithAPIKey("sk-1YnJtTmClFr15ApMknwB7ND1VnsTsmpsEv8nSzZwhRmroFUi"), // 混元 APIKey
		option.WithBaseURL("https://api.hunyuan.cloud.tencent.com/v1/"),          // 混元 endpoint
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("你是谁？"),
			},
			Model: "hunyuan-turbos-latest",
		},
		option.WithJSONSet("enable_enhancement", true), // <- 自定义参数
	)
	if err != nil {
		panic(err.Error())
	}
	t.Logf("ChatCompletion: %v", chatCompletion)
}
