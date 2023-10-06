package main

import (
	operator "ABCAI-moderator/openai"
)

func main() {
	println("=== Testing moderator ===")
	score := operator.ComputeScoreProposal("title", "description")
	println("score:", score)
}