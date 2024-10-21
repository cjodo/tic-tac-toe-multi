import { Message } from "./Message/Message";

interface Message {
  data: string;
  timeStamp: string;
}

interface ChatHistoryProps {
  chatHistory: Message[]
}

export const ChatHistory = ({chatHistory}: ChatHistoryProps) => {

  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      { chatHistory.map((msg) => {
        return <Message key={msg.timeStamp} message={msg.data}/>
      })
      }
    </div>
  )
}
