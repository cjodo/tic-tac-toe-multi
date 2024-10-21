export interface Message {
  data:       string;
  timeStamp:  string
}

export interface ChatHistory {
  messages: Message[]
}
