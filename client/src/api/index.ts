const socket = new WebSocket("ws://localhost:4040/ws");

export const connect = () => {
		console.log("Attempting Connection...");

		socket.addEventListener('open', () => {
				console.log("Connection Successful");
		})

		socket.addEventListener('message', (message) => {
				console.log(message)
		})

		socket.addEventListener('close', (e) => {
				console.log("Socket Connection Closed", e)
		})

		socket.addEventListener('error', (error) => {
				console.log("Socket error: ", error)
		})
}

export const sendMsg = (msg: string) => {
		console.log("sending message: ", msg)
		socket.send(msg)
}

