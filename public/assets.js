const socket = new WebSocket('ws://localhost:3000/payments');

socket.addEventListener('open', ev => {
  document.monetization.addEventListener(
    'monetizationprogress',
    e => {
      socket.send(JSON.stringify(e.detail))
    }
  )
})
