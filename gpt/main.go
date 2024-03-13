package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

func main() {
	log.Println(os.Args)
	ctx := context.Background()
	token := os.Args[1]
	if token == "" {
		log.Fatalf("missing GPT auth token")
	}
	client := openai.NewClient(token)
	ReuseAssistantResource(ctx, client)
}

func getInputText() string {
	// Create a new Scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user for input
	fmt.Print("(end with '##')-> ")

	// Read lines from stdin until the input contains '##'
	var inputLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, "##") {
			// Remove '##' from the end of the line and add it to the inputLines
			line = strings.TrimSuffix(line, "##")
			inputLines = append(inputLines, line)
			break
		}
		inputLines = append(inputLines, line)
	}

	return strings.TrimSpace(strings.Join(inputLines, "\n"))
}

func strToPoint(s string) *string {
	return &s
}

func ChatCompletionMessage() {
	client := openai.NewClient("")
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT4TurboPreview,
				Messages: messages,
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		content := resp.Choices[0].Message.Content
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println(content)
	}
}

func CreateAssistantCompletionMessage(ctx context.Context, client *openai.Client) {
	// step 1: upload original file and translated file
	fileOri, err := client.CreateFile(ctx, openai.FileRequest{
		FileName: "AI Translation Sample File Original.docx",
		FilePath: "./AI Translation Sample File Original.docx",
		Purpose:  "assistants",
	})
	if err != nil {
		log.Fatalf("step 1: upload file failed err: %v", err)
	}
	fileOriJSON, _ := json.Marshal(fileOri)
	log.Printf("step 1: upload file succ fileOri: %s", fileOriJSON)

	fileTrans, err := client.CreateFile(ctx, openai.FileRequest{
		FileName: "AI Translation Sample File Translated.docx",
		FilePath: "./AI Translation Sample File Translated.docx",
		Purpose:  "assistants",
	})
	if err != nil {
		log.Fatalf("step 1.1: upload file failed err: %v", err)
	}
	fileTransJSON, _ := json.Marshal(fileTrans)
	log.Printf("step 1.1: upload file succ fileTrans: %s", fileTransJSON)

	// step 2: create assistant
	assistant, err := client.CreateAssistant(ctx, openai.AssistantRequest{
		Model:        openai.GPT4TurboPreview,
		Name:         strToPoint("Ethan-doc-demo"),
		Description:  nil,
		Instructions: strToPoint("You are a document translation assistant, and you need to help users with all operations related to the documents they upload."),
		Tools: []openai.AssistantTool{
			{
				Type: openai.AssistantToolTypeRetrieval,
			},
		},
		FileIDs:  []string{fileOri.ID, fileTrans.ID},
		Metadata: map[string]any{},
	})
	if err != nil {
		log.Fatalf("step 2: create assistant failed err: %v", err)
	}
	assistantJSON, _ := json.Marshal(assistant)
	log.Printf("step 2: create assistant succ assistant: %s", assistantJSON)

	RunAssistantCompletionMessage(ctx, client, assistant, "")
}

func RunAssistantCompletionMessage(ctx context.Context, client *openai.Client, assistant openai.Assistant, threadID string) {
	if threadID == "" {
		// step 3: create a thread without message
		thread, err := client.CreateThread(ctx, openai.ThreadRequest{
			Messages: []openai.ThreadMessage{},
			Metadata: map[string]any{},
		})
		if err != nil {
			log.Fatalf("step 3: create thread failed err: %v", err)
		} else {
			threadJSON, _ := json.Marshal(thread)
			log.Printf("step 3: create thread succ thread: %s", threadJSON)
		}
		threadID = thread.ID
	}
	fmt.Println("Enter text (end with '##'): ")
	fmt.Println("Conversation")
	fmt.Println("---------------------")

	for {
		text := getInputText()
		if text == "break" {
			break
		}
		if text == "" {
			continue
		}
		// Step 4: Add a Message to a Thread
		sendMsg, err := client.CreateMessage(ctx, threadID, openai.MessageRequest{
			Role:    string(openai.ThreadMessageRoleUser),
			Content: text,
			// FileIds:  []string{},
			// Metadata: map[string]any{},
		})
		if err != nil {
			log.Printf("step 4: create message failed err: %v", err)
		} else {
			sendMsgJSON, _ := json.Marshal(sendMsg)
			log.Printf("step 4: create message succ sendMsg: %s", sendMsgJSON)
		}

		// Step 5: Run the Assistant
		run, err := client.CreateRun(ctx, threadID, openai.RunRequest{
			AssistantID: assistant.ID,
		})
		if err != nil {
			log.Printf("Step 5: create run failed err: %v", err)
		} else {
			runJSON, _ := json.Marshal(run)
			log.Printf("Step 5: create run succ run: %s", runJSON)
		}
		// the displayed message are not appearing again
		displayedMsgIdMap := map[string]bool{}
		for {
			// Step 6: loop Check the Run status
			runResult, err := client.RetrieveRun(ctx, threadID, run.ID)
			if err != nil {
				log.Printf("Step 6: retrieve run failed err: %v", err)
			} else {
				runResultJSON, _ := json.Marshal(runResult)
				log.Printf("Step 6: retrieve run succ runResult: %s", runResultJSON)
			}
			// Step 6.1: list run steps
			stepList, err := client.ListRunSteps(ctx, threadID, run.ID, openai.Pagination{})
			if err != nil {
				log.Printf("Step 6.1: list run steps failed err: %v", err)
			} else {
				stepListJSON, _ := json.Marshal(stepList)
				log.Printf("Step 6.1: list run steps succ stepList: %s", stepListJSON)
			}
			for _, step := range stepList.RunSteps {
				if step.Type == openai.RunStepTypeMessageCreation &&
					step.Status == openai.RunStepStatusCompleted &&
					!displayedMsgIdMap[step.StepDetails.MessageCreation.MessageID] {
					receivedMsg, err := client.RetrieveMessage(ctx, threadID, step.StepDetails.MessageCreation.MessageID)
					if err != nil {
						log.Printf("Assistant Answer failed err: %v", err)
					} else {
						receivedMsgJSON, _ := json.Marshal(receivedMsg)
						log.Printf("\n\n***************Assistant Answer***************:\n\n%s", receivedMsgJSON)
					}
					displayedMsgIdMap[step.StepDetails.MessageCreation.MessageID] = true
				}
			}

			if runResult.Status == openai.RunStatusInProgress {
				time.Sleep(2 * time.Second)
			} else if runResult.Status == openai.RunStatusQueued ||
				runResult.Status == openai.RunStatusRequiresAction ||
				runResult.Status == openai.RunStatusCancelling {
				time.Sleep(2 * time.Second)
			} else if runResult.Status == openai.RunStatusCompleted ||
				runResult.Status == openai.RunStatusFailed ||
				runResult.Status == openai.RunStatusExpired {
				break
			} else {
				break
			}
		}

	}

	runList, err := client.ListRuns(ctx, threadID, openai.Pagination{})
	if err != nil {
		log.Printf("list runs failed err: %v", err)
	} else {
		runListJSON, _ := json.Marshal(runList)
		log.Printf("list runs succ runList: %s", runListJSON)
	}

	msgList, err := client.ListMessage(ctx, threadID, nil, nil, nil, nil)
	if err != nil {
		log.Printf("list message failed err: %v", err)
	} else {
		msgListJSON, _ := json.Marshal(msgList)
		log.Printf("list message succ msgList: %s", msgListJSON)
	}
}

func ListAllAssistantResource(ctx context.Context, client *openai.Client) openai.AssistantsList {
	// list assistants
	assList, err := client.ListAssistants(ctx, nil, nil, nil, nil)
	if err != nil {
		log.Fatalf("list assistants failed err: %v", err)
	} else {
		assListJSON, _ := json.Marshal(assList)
		log.Printf("list assistants succ assList: %s", assListJSON)
	}
	assistantID := assList.FirstID

	// list assistants files
	assistantFilesList, err := client.ListAssistantFiles(ctx, *assistantID, nil, nil, nil, nil)
	if err != nil {
		log.Fatalf("list assistants failed err: %v", err)
	} else {
		assistantFilesListJSON, _ := json.Marshal(assistantFilesList)
		log.Printf("list assistants succ assistantFilesList: %s", assistantFilesListJSON)
	}

	// list uploaded files
	fileList, err := client.ListFiles(ctx)
	if err != nil {
		log.Fatalf("list uploaded files failed err: %v", err)
	} else {
		fileListJSON, _ := json.Marshal(fileList)
		log.Printf("list uploaded files succ fileList: %s", fileListJSON)
	}
	return assList
}

func ReuseAssistantResource(ctx context.Context, client *openai.Client) {
	assList := ListAllAssistantResource(ctx, client)
	RunAssistantCompletionMessage(ctx, client, assList.Assistants[0], "")
}
