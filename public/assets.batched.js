const urlPath = window.location.pathname;
const BATCH = 5;
const INTERVAL = 5000;
let paymentStore = [];
let SEND_PAYMENTS = false;

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

