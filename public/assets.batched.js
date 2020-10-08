const urlPath = window.location.pathname;
const BATCH = 5;
const INTERVAL = 5000;
let paymentStore = [];

window.setInterval(handler, INTERVAL)

function calculateAmount(amount, assetScale) {
  const value = parseInt(amount)
  const scale = Math.pow(10, assetScale)

  return value / scale
}

document.monetization.addEventListener(
  'monetizationstart',
  e => {
    SEND_PAYMENTS = true
  }
)

document.monetization.addEventListener(
  'monetizationprogress',
  e => {
    const receivedAt = new Date().toJSON();
    const { detail } = e;
    const payload = {
      ...detail,
      urlPath,
      receivedAt,
    };

    paymentStore.push(payload);
    setInterval(sendPaymentDetails(true), 5000)
  }
)

document.monetization.addEventListener(
  'monetizationstop',
  e => {
    sendPaymentDetails()
  }
)

function sendPaymentDetails(batched = false) {
  let data = {
    payments: batched ? paymentStore.splice(0, BATCH) : paymentStore,
  }

  if (!data.payments.length) {
    return
  }

  fetch("http://localhost:3000/capture", {
    method: 'POST',
    body: JSON.stringify(data)
  })
  .then(res => res.json)
  .then(data => console.log(data))
  .catch(err => console.log(err.message));
}

