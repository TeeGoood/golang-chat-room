// elements
const form = document.querySelector('form');
const chat = document.querySelector('#chat');

// websocket
socket = new WebSocket(`ws://localhost:8000/ws`);
console.log('Attempting websocket connection');

socket.onopen = () => {
	console.log('Connection successful');
};

socket.onclose = () => {
	console.log('socket close connection');
};

socket.onerror = (error) => {
	console.error('error socket: ', error);
};

socket.onmessage = (event) => {
	chat.innerHTML += `<div>${event.data}</div>`;
};

// event handlers
function sendMsg(e) {
	e.preventDefault();
	let message = e.target[0].value;
	socket.send(message);
}

form.onsubmit = sendMsg;
