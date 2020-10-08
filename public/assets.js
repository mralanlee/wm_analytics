const socket = new WebSocket('ws://localhost:3000/payments');

// socket.addEventListener('open', ev => {
// })

const urlPath = window.location.pathname;

document.monetization.addEventListener(
  'monetizationprogress',
  e => {
    if (socket.readyState === 1) {
      const { detail } = e;
      const payments = {
        ...detail,
        urlPath,
      };

      socket.send(JSON.stringify(payments))
    } else if (socket.readyState === 2 || socket.readyState === 3) {
      console.log("WebSocket Connection is closed");
    } else {
      console.log("Connection is not ready", socket.readyState)
    }
  }
)
