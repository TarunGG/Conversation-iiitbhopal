let popup = document.getElementById("popup");
let form = document.getElementById("form");
let formDiv = document.getElementById('formDiv')

form.addEventListener('submit', (e) => {
  e.preventDefault()

  let formData = new Object()
  formData={
    'username': document.querySelector('#usernameInp').value,
    'satisfaction': document.querySelector('input[type="radio"]').value,
    'feedback': document.querySelector('#feedback').value,
  }
  // formData['bug']=document.querySelector('#bug').files[0]

  fetch('https://formspree.io/f/mgeqzgke',
    {
      method: "POST",
      // body: formData,
      body:JSON.stringify(formData),
      mode: "cors"
    }
  )
  form.reset()
  formDiv.style.display = 'none'
  popup.style.display = 'block'
})