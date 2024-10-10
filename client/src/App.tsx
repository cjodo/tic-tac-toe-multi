'use strict'

import './App.css'
import { connect, sendMsg } from './api';

function App() {
  connect();

  const send = () => {
    sendMsg("Hello")
  }

  return (
    <>
      <div className="card">
        <button onClick={send}>Send Message</button>
      </div>
    </>
  )
}

export default App
