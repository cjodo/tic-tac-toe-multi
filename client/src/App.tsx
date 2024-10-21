import './App.css'

import { Board } from './components/Board/Board'
import { useEffect, useState } from 'react'

import { connect } from './api';

function App() {

  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const ws = connect("ws://localhost:4040/ws", null)
    setSocket(ws)

    return () => {
      ws.close()
    }
  }, [])

  return (
    <>
      <div>
        <Board socket={socket}/>
      </div>
    </>
  )
}

export default App
