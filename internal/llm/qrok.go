package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

const (
	MESSAGE_ROLE_USER      = "user"
	MESSAGE_ROLE_ASSISTANT = "assistant"

	GORQ_REQ_URL = "https://api.groq.com/openai/v1/chat/completions"

	PROMPT = `process DIFF and suggest commit message.


### Output Structure explained
replace <type> with: feature, fix, docstest
replace <description> with: short description about commit in less than 50 characters
replace <body> with: long description about commit if needed between 0 and 8 lines

### Requirements:
- Break long lines
- Just send plain-text commit message without any extra words/explaintation
- Strongly comply Output Structure
- Commit message should be standard

### DIFF
%s

### Output Structure
<type>: <description>

body`
)

type ResponseFormat struct {
	Type string `json:"type"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GorqRequest struct {
	Messages       []Message       `json:"messages"`
	Model          string          `json:"model"`
	MaxTokens      int             `json:"max_tokens,omitempty"`
	Temperature    float64         `json:"temperature,omitempty"`
	TopP           int             `json:"top_p,omitempty"`
	Stream         bool            `json:"stream,omitempty"`
	Stop           string          `json:"stop,omitempty"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Delta        Message `json:"delta"`
	FinishReason string  `json:"finish_reason"`
}

type GorqResponseError struct {
	Message         string `json:"message"`
	Type            string `json:"type"`
	Code            string `json:"code"`
	FailedGenerated string `json:"failed_generation"`
}

type GorqResponse struct {
	Error             GorqResponseError `json:"error"`
	ID                string            `json:"id"`
	Object            string            `json:"object"`
	Created           int64             `json:"created"`
	Model             string            `json:"model"`
	SystemFingerprint string            `json:"system_fingerprint"`
	Choices           []Choice          `json:"choices"`
}

type Qrok struct {
}

func (q Qrok) GenerateCommitByDiff(diff string) (string, error) {
	body := GorqRequest{
		Messages: []Message{{
			Role:    MESSAGE_ROLE_USER,
			Content: fmt.Sprintf(PROMPT, diff),
		}},
		Model:       viper.GetString("LLM_MODEL"),
		Temperature: 1,
		MaxTokens:   1024,
		TopP:        1,
	}
	commitMessage, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, GORQ_REQ_URL, bytes.NewBuffer(commitMessage))
	if err != nil {
		return "", fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+viper.GetString("GROK_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	resBytes, _ := io.ReadAll(res.Body)

	gorqRes := GorqResponse{}
	err = json.Unmarshal(resBytes, &gorqRes)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to do http request: %s", gorqRes.Error.Message)
	}

	return gorqRes.Choices[0].Message.Content, nil
}
