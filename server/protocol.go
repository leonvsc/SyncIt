package main

import (
	"fmt"
)

// MessageType defines the type of message
type MessageType string

// ProtocolType defines the type of protocol
type ProtocolType string

// FileSystemType defines the type of file system
type FileSystemType string

// ContentType defines the type of content
type ContentType string

// Header defines the header of the message
type Header struct {
	RequestType     MessageType    `json:"request_type"`
	ProtocolType    ProtocolType   `json:"protocol_type"`
	ProtocolVersion string         `json:"protocol_version"`
	ContentType     ContentType    `json:"content_type"`
	ContentLength   int            `json:"content_length"`
	Path            string         `json:"path,omitempty"`
	GUID            string         `json:"guid,omitempty"`
	FileSystem      FileSystemType `json:"file_system,omitempty"`
	FileName        string         `json:"file_name,omitempty"`
	FileExtension   string         `json:"file_extension,omitempty"`
	MimeType        string         `json:"mime_type,omitempty"`
	Authorization   string         `json:"authorization,omitempty"`
}

// Message defines the structure of the message
type Message struct {
	Headers Header `json:"headers"`
	Body    []byte `json:"body"`
}

// protocolHandler handles the message based on the request type
func protocolHandler(msg Message) error {
	switch msg.Headers.RequestType {
	case "GET":
		fmt.Println("Handling GET request...")
	case "POST":
		fmt.Println("Handling POST request...")
	case "PUT":
		fmt.Println("Handling PUT request...")
	case "DELETE":
		fmt.Println("Handling DELETE request...")
	default:
		return fmt.Errorf("unknown request type: %s", msg.Headers.RequestType)
	}

	return nil
}

func sync() {
	fmt.Println("Executing synchroniseren...")
	receivedMessage := Message{
		Headers: Header{
			RequestType:  "GET",
			ProtocolType: "TCP",
		},
		Body: []byte(""),
	}

	// Process the received message
	if err := protocolHandler(receivedMessage); err != nil {
		fmt.Println("Error handling message:", err)
		return
	}
}
