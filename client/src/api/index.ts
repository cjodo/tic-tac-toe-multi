
export const connect = (address: string, callback: Function | null) => {

	const socket = new WebSocket(address);
	console.log(`Attempting Connection To ${address}...`);

	socket.addEventListener('open', () => {
		console.log("Connection Successful");
	})

	socket.addEventListener('message', (message) => {
		if(callback) {
			callback(message)
		}
	})

	socket.addEventListener('close', (e) => {
		console.log("Socket Connection Closed", e)
	})

	socket.addEventListener('error', (error) => {
		console.log("Socket error: ", error)
	})


	return socket
}
