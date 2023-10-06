package openai

import (
	"context"
	"fmt"
	"os"

	"ABCAI-moderator/types"

	openai "github.com/sashabaranov/go-openai"
)

func ComputeScoreBatchProposals(proposals []types.Proposal) ([]float64, []error) {
	var scores []float64
	var errors []error
	scoreChan := make(chan float64)
	errorChan := make(chan error)

	for _, proposal := range proposals {
		go func(proposal types.Proposal) {
			score, err := ComputeScoreProposal(proposal)
			if err != nil {
				errorChan <- err
			} else {
				scoreChan <- score
			}
		}(proposal)
	}

	for i := 0; i < len(proposals); i++ {
		select {
		case score := <-scoreChan:
			scores = append(scores, score)
		case err := <-errorChan:
			errors = append(errors, err)
		}
	}

	return scores, errors
}

func ComputeScoreProposal(proposal types.Proposal) (float64, error) {
	client := openai.NewClient(os.Getenv("OPEN_AI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Temperature: 0.0,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Given a document, you are a blockchain governance proposal scorer. Your task is to read the document and determine whether it represents a legit blockchain governance proposal or a scam proposal. Assign a score between 0 and 1, where 1 represents a legit proposal and 0 represents a scam proposal.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: `Given a document, you are a blockchain governance proposal scorer. Your task is to read the document and determine whether it represents a legit blockchain governance proposal or a scam proposal. Assign a score between 0 and 1, where 1 represents a legit proposal and 0 represents a scam proposal.

					## Document:
					Title: {{proposal_title}}
					Description: {{proposal_description}}
					
					## Scoring Criteria:
					A legit blockchain governance proposal typically includes:
					- Clear and concise title and description
					- Author information with verifiable links
					- Detailed proposal content related to blockchain ecosystem, development, or community improvement
					- References to official discussion forums or documents
					- Links to relevant technical details or supporting documentation
					
					A scam proposal typically includes:
					- Vague or misleading title and description
					- Absence of reliable author information or links
					- Lack of detailed content or relevance to blockchain development or ecosystem
					- References to unofficial or suspicious discussion forums
					- Missing or broken links to relevant technical details or supporting documentation
					
					## Instructions:
					Read the given document and assign a score between 0 and 1 to indicate whether it is a legit blockchain governance proposal or a scam proposal. Use the scoring criteria provided above as guidelines.
					Given a prompt, return a score between 0 and 1 indicating the likelihood that the proposal is a legitimate blockchain governance proposal. A score of 1 indicates a high likelihood of legitimacy, while a score of 0 indicates a high likelihood of being a scam proposal.
					
					## Prompt:

					Title:` + proposal.Title + `
					Description:` + proposal.Description + `
					
					Expected Output: A float between 0 and 1 only.`,
				},
			},
		},
	)

	fmt.Printf("ChatCompletion response: %v\n", resp)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return 0.0, err
	}

	var score float64
	_, err = fmt.Sscanf(resp.Choices[0].Message.Content, "%f", &score)

	if err != nil {
		fmt.Printf("Error parsing score: %v\n", err)
		return 0.0, err
	}

	return score, nil
}