const socket = new WebSocket('ws://localhost:3000/payments');

// socket.addEventListener('open', ev => {
// })

document.monetization.addEventListener(
  'monetizationprogress',
  e => {
    if (socket.readyState === 1) {
      socket.send(JSON.stringify(e.detail))
    } else if (socket.readyState === 2 || socket.readyState === 3) {
      console.log("WebSocket Connection is closed");
    } else {
      console.log("Connection is not ready", socket.readyState)
    }
  }
)
