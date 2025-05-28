package relaymode

const (
	Unknown = iota
	ChatCompletions
	BatchInference
	Completions
	Embeddings
	Moderations
	ImagesGenerations
	Edits
	AudioSpeech
	AudioTranscription
	AudioTranslation
	// Proxy is a special relay mode for proxying requests to custom upstream
	Proxy
)
