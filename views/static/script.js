const getOtpBtn = document.getElementById('get-otp');
const otpInput = document.getElementById('otp-input');
const sendOtpBtn = document.getElementById('send-otp');

getOtpBtn.addEventListener('click', async () => {
  await fetch('/api/get-otp', {method: 'POST'})
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      localStorage.setItem('otpKey', data.data.key);
      alert('OTP sent to your email')
    }).catch((error) => {
      console.log(error);
    })
});

let otpValueInput = '';
otpInput.addEventListener('change', (event) => {
  otpValueInput = event.target.value;
})

sendOtpBtn.addEventListener('click', async (event) => {
  event.preventDefault();
  const otpKey = localStorage.getItem('otpKey');
  await fetch('/api/send-otp', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      key: otpKey,
      otp: otpValueInput
    })
  })
  .then((response) => response.json())
  .then((data) => {
    console.log(data);
    if (data.errors) {
      alert(data.errors);
      return;
    }
    alert(data.data.message)
  }).catch((error) => {
    console.log(error);
  })
})